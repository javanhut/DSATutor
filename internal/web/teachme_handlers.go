package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
	"strings"

	"dsatutor/internal/practice"
	"dsatutor/internal/tutor"
)

// handleTeachLesson builds a Teach Me lesson on any topic the learner names
// and attaches practice-bank problems that use it.
func (s *Server) handleTeachLesson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Model string `json:"model"`
		Topic string `json:"topic"`
		Level string `json:"level"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 64<<10)).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Topic) == "" {
		http.Error(w, "topic is required", http.StatusBadRequest)
		return
	}
	lesson, model, err := s.tutor.TeachTopic(r.Context(), req.Model, req.Topic, req.Level)
	if err != nil {
		writeTeachError(w, err)
		return
	}
	lesson.Problems = s.teachProblems(req.Topic, lesson.Keywords, 5)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"model": model, "lesson": lesson})
}

// handleTeachFeedback grades the learner's explain-it-back answer.
func (s *Server) handleTeachFeedback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Model     string `json:"model"`
		Topic     string `json:"topic"`
		Prompt    string `json:"prompt"`
		Reference string `json:"reference"`
		Answer    string `json:"answer"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 64<<10)).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Answer) == "" {
		http.Error(w, "answer is required", http.StatusBadRequest)
		return
	}
	fb, model, err := s.tutor.GradeExplanation(r.Context(), req.Model, req.Topic, req.Prompt, req.Reference, req.Answer)
	if err != nil {
		writeTeachError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"model": model, "feedback": fb})
}

func writeTeachError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusBadGateway
	if errors.Is(err, tutor.ErrUnavailable) {
		status = http.StatusServiceUnavailable
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
}

// teachProblems ranks practice problems by how well their tags, category
// and title match the topic and the lesson's keywords. The learner's own
// words count most: the whole topic, then each two-word phrase in it (so
// "binary search on the answer" finds the Binary Search tag), then the
// model's keywords. Easier problems come first among equal matches.
func (s *Server) teachProblems(topic string, keywords []string, n int) []tutor.TeachProblem {
	if s.practiceServer == nil {
		return nil
	}
	norm := strings.NewReplacer("-", " ", "_", " ")
	clean := func(k string) string { return strings.Join(strings.Fields(norm.Replace(strings.ToLower(k))), " ") }
	terms := map[string]int{}
	add := func(k string, w int) {
		// Singular, so "prefix sums" finds the "Prefix Sum" tag and the
		// plural still matches by substring.
		if k = strings.TrimSuffix(clean(k), "s"); len(k) >= 3 && w > terms[k] {
			terms[k] = w
		}
	}
	add(topic, 6)
	words := strings.Fields(clean(topic))
	stop := map[string]bool{"a": true, "an": true, "the": true, "on": true, "of": true, "to": true, "in": true, "for": true, "with": true, "and": true, "or": true, "vs": true, "how": true, "what": true, "is": true, "me": true}
	for i := 0; i+1 < len(words); i++ {
		if !stop[words[i]] && !stop[words[i+1]] {
			add(words[i]+" "+words[i+1], 4)
		}
	}
	for _, k := range keywords {
		add(k, 3)
	}
	type scored struct {
		p     *practice.Problem
		score int
	}
	var hits []scored
	for _, p := range s.practiceServer.loader.GetAllProblems() {
		title := clean(p.Title)
		hay := clean(p.Category + " " + strings.Join(p.Tags, " "))
		score := 0
		for t, w := range terms {
			if strings.Contains(hay, t) {
				score += w
			}
			if strings.Contains(title, t) {
				score += w / 2
			}
		}
		if score > 0 {
			hits = append(hits, scored{p, score})
		}
	}
	rank := map[string]int{"Easy": 0, "Medium": 1, "Hard": 2}
	sort.SliceStable(hits, func(i, j int) bool {
		if hits[i].score != hits[j].score {
			return hits[i].score > hits[j].score
		}
		return rank[hits[i].p.Difficulty] < rank[hits[j].p.Difficulty]
	})
	out := []tutor.TeachProblem{}
	for _, h := range hits {
		if len(out) == n {
			break
		}
		out = append(out, tutor.TeachProblem{ID: h.p.ID, Title: h.p.Title, Difficulty: h.p.Difficulty, Category: h.p.Category})
	}
	// Show them easiest first.
	sort.SliceStable(out, func(i, j int) bool { return rank[out[i].Difficulty] < rank[out[j].Difficulty] })
	return out
}
