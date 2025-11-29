package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"dsatutor/internal/chapter"
	"dsatutor/internal/practice"
	"dsatutor/internal/storage"
	"dsatutor/internal/ui"
	"dsatutor/internal/web"
)

func main() {
	exportPath := flag.String("export", "", "export chapters to a JSON file")
	loadPath := flag.String("load", "", "load chapters from a JSON file (overrides built-in templates)")
	playChapter := flag.Int("play", 0, "play the first storyboard for a chapter number")
	serve := flag.Bool("serve", false, "start the web UI server")
	addr := flag.String("addr", ":8080", "address for web UI (use with -serve)")
	flag.Parse()

	chapters := chapter.DefaultChapters()
	if *loadPath != "" {
		var err error
		chapters, err = storage.LoadJSON(*loadPath)
		if err != nil {
			log.Fatalf("load chapters: %v", err)
		}
	}

	registry := chapter.NewRegistry(chapters)

	if *exportPath != "" {
		if err := storage.SaveJSON(*exportPath, registry.All()); err != nil {
			log.Fatalf("export chapters: %v", err)
		}
		fmt.Fprintf(os.Stdout, "Exported %d chapters to %s\n", len(chapters), *exportPath)
	}

	fmt.Println("Interactive DSA Tutor: Chapter Templates")
	fmt.Println("----------------------------------------")
	for _, ch := range registry.All() {
		fmt.Printf("%02d. %s\n   slug: %s\n   summary: %s\n\n", ch.Number, ch.Title, ch.Slug, ch.Summary)
	}

	if *serve {
		// Initialize practice problem loader
		problemLoader := practice.NewProblemLoader(practice.ProblemsFS, practice.ProblemsBasePath)
		if err := problemLoader.Load(); err != nil {
			log.Printf("Warning: Could not load practice problems: %v", err)
			// Continue without practice problems
			server := web.NewServer(registry.All())
			if err := server.Listen(*addr); err != nil {
				log.Fatalf("serve: %v", err)
			}
			return
		}

		stats := problemLoader.GetStats()
		log.Printf("Loaded %d practice problems (Easy: %d, Medium: %d, Hard: %d)",
			stats["total"], stats["easy"], stats["medium"], stats["hard"])

		server := web.NewServerWithPractice(registry.All(), problemLoader)
		if err := server.Listen(*addr); err != nil {
			log.Fatalf("serve: %v", err)
		}
		return
	}

	if *playChapter > 0 {
		ch, ok := registry.Chapter(*playChapter)
		if !ok {
			log.Fatalf("chapter %d not found", *playChapter)
		}
		if len(ch.Animations) == 0 {
			log.Fatalf("chapter %d has no storyboard to play", *playChapter)
		}
		renderer := ui.ConsoleRenderer{Out: os.Stdout}
		fmt.Printf("Playing storyboard for chapter %d: %s\n", ch.Number, ch.Title)
		if err := renderer.Render(ch.Animations[0]); err != nil {
			log.Fatalf("render storyboard: %v", err)
		}
	}
}
