package chapter

// chapter21 is the bonus chapter: k-nearest neighbours (distance, scaling,
// voting) and MapReduce (map, shuffle, reduce), showing how the patterns of
// the course reappear outside interviews, and ending with what to study next.
// Storyboards to include (IDs registered in app.js): none.
func chapter21() Chapter {
	return Chapter{
		Number:        21,
		Slug:          "beyond",
		Title:         "Beyond Interviews: KNN and MapReduce",
		Category:      "",
		Prerequisites: []int{1},
		Summary:       "Two ideas from outside the interview canon that are built from tools you already have: k-nearest neighbours classifies a point by measuring distances and taking a vote (a heap picks the k closest), and MapReduce counts or aggregates data too big for one machine by mapping, grouping by key (a hash map), and reducing. The chapter ends with a map of what to study next.",
		Motivation: `A fruit stand has 1,000 labelled examples of apples and oranges, each described by weight (grams) and diameter (cm). A new fruit weighs 150 g and is 7 cm across. Writing rules by hand ("if weight > 140 and diameter < 7.5 ...") gets messy fast and breaks as soon as a new variety arrives. A simpler idea: find the few known fruits that look most like the new one and let them vote. That is k-nearest neighbours, and it needs nothing more than a distance formula, a way to pick the k smallest of 1,000 distances (a heap, Chapter 11) and a counter for the vote (Chapter 2).

Now scale up: count how often each word appears in 10 billion web pages. One machine reading 100 MB per second needs about a month just to read 250 TB. Split the pages across 1,000 machines and each needs well under an hour, but then every machine holds partial counts for the same words. MapReduce is the recipe for that: each machine maps its pages to (word, 1) pairs, the system shuffles pairs so that all pairs for one word meet on one machine (a distributed hash map), and each machine reduces its words by adding up the 1s.

Neither idea is an interview pattern, but both are everyday tools in data work, and both are compositions of the data structures from this course.`,
		Objectives: []string{
			"Compute Euclidean and Manhattan distances between feature vectors and explain why features must be scaled first",
			"Trace a k-nearest-neighbours classification on a small data set, including the choice of k and the vote",
			"Implement KNN with heapq.nsmallest and state its cost per query in Big O",
			"Trace a MapReduce word count through the map, shuffle and reduce stages and explain what each stage can run in parallel",
			"Choose a next topic to study and connect it to a pattern from an earlier chapter",
		},
		WhenToUse: []string{
			"KNN: you have labelled examples and want to label a new one by similarity, with little data and no time to train a model",
			"KNN: 'find the k closest points / most similar items' to a query, which is a top-k selection with a heap",
			"MapReduce: the data is too large for one machine, and the answer is an aggregate per key (counts, sums, maximums, lists)",
			"MapReduce: the work splits into independent pieces whose partial results can be combined by key",
		},
		Concepts: []Concept{
			{
				Name:        "k-nearest neighbours (KNN)",
				Description: "To label a new point, measure its distance to every labelled example, keep the k closest, and let them vote (classification) or average their values (regression). There is no training phase: the data set itself is the model.",
				WhyItWorks:  `KNN rests on one assumption: points that are close in feature space tend to have the same label. If that holds, the labels of the nearest examples are good evidence for the query's label. Using k neighbours instead of 1 makes the vote robust: a single mislabelled or unusual example can decide a 1-NN prediction, but it is outvoted when k = 5. Making k too large has the opposite problem: with k = n every query gets the overall majority label. Picking the k closest of n distances is a top-k selection, so a heap of size k does it in O(n log k) instead of sorting all n distances in O(n log n). Each query must compute every distance, so the cost per query is O(n * d) for n examples with d features, which is why large systems use spatial indexes (k-d trees, approximate nearest-neighbour search).`,
				Intuition:   `Moving to a new neighbourhood and wondering which football team to support: you ask the five closest neighbours and go with the majority. Asking only the one next door risks a fan of a tiny club (k too small); asking the whole city just gives you the city's most popular team (k too big). Neighbours are training examples, closeness is distance in feature space, and the majority is the predicted label.`,
				CoreIdeas: []string{
					"Euclidean distance: sqrt(sum((a_i - b_i)^2)); compare squared distances to skip the square root, since it does not change the order",
					"Manhattan distance: sum(|a_i - b_i|), often better on grids or when features are counts",
					"Keep the k closest with heapq.nsmallest(k, distances): O(n log k) instead of O(n log n)",
					"Classification: majority vote of the k labels (use an odd k for two classes to avoid ties); regression: average of their values",
					"Cost per query is O(n * d) for n examples with d features; no training cost",
				},
				CommonMistakes: []string{
					"Choosing k = 1 and trusting it: one noisy example decides everything. Try several k values on held-out examples and pick the most accurate.",
					"Using an even k with two classes, which can tie 2-2. Use an odd k or break ties by the nearest neighbour's label.",
					"Sorting all distances for every query when only k are needed. Use a heap of size k (Chapter 11).",
				},
				Examples: []CodeExample{
					{
						ID:       "knn-classify",
						Title:    "KNN classification with a heap",
						Language: "python",
						Snippet: `import heapq
from collections import Counter

def knn_classify(points, labels, query, k):
    distances = []
    for point, label in zip(points, labels):
        d2 = sum((p - q) ** 2 for p, q in zip(point, query))
        distances.append((d2, label))
    nearest = heapq.nsmallest(k, distances)
    votes = Counter(label for _, label in nearest)
    return votes.most_common(1)[0][0]`,
						Notes: "Squared distances give the same order as real distances, so the square root is never needed.",
					},
				},
			},
			{
				Name:        "Feature scaling",
				Description: "Distances add up differences across features, so a feature measured in large units (dollars, grams) can drown out one measured in small units (years, centimetres). Rescale every feature to a comparable range before computing distances.",
				WhyItWorks:  `Euclidean distance squares and adds the per-feature differences. If income ranges over 20,000-200,000 and age over 18-90, a 1,000-dollar income difference contributes 1,000,000 to the squared distance while a 30-year age difference contributes only 900, so age is effectively ignored even if it matters more. Min-max scaling maps each feature to 0..1 with (x - min) / (max - min); standardisation maps it to mean 0 and standard deviation 1 with (x - mean) / std. After either, a one-unit difference means 'one full range' or 'one standard deviation' in every feature, so each feature gets a fair say. The scaling parameters must come from the training data and be applied unchanged to queries; otherwise the query lives in a different coordinate system.`,
				Intuition:   `Comparing two runners by 'distance' made of their marathon time in seconds and their age in decades: the seconds column is thousands of times larger, so the comparison is really just about time. Converting both columns to percent-of-range is like putting both on the same ruler before adding them up.`,
				CoreIdeas: []string{
					"Min-max: (x - min) / (max - min) maps the training range to 0..1",
					"Standardise: (x - mean) / std gives mean 0 and spread 1",
					"Compute min, max, mean and std on the training data only, then apply them to queries",
					"After scaling, choose the distance (Euclidean or Manhattan) and k",
				},
				CommonMistakes: []string{
					"Skipping scaling, so the feature with the largest numbers decides every neighbour. Scale first, then measure distance.",
					"Scaling the query with its own min and max (or rescaling the whole data set including the query), which shifts it into different coordinates. Reuse the training data's parameters.",
					"Dividing by zero when a feature is constant (max == min). Drop that feature or treat its scaled value as 0.",
				},
				Examples: []CodeExample{
					{
						ID:       "min-max-scale",
						Title:    "Min-max scaling per feature",
						Language: "python",
						Snippet: `def fit_min_max(points):
    columns = list(zip(*points))
    return [(min(col), max(col)) for col in columns]

def scale(point, ranges):
    return [(x - lo) / (hi - lo) if hi > lo else 0.0
            for x, (lo, hi) in zip(point, ranges)]

ranges = fit_min_max([[20000, 18], [200000, 90], [65000, 40]])
print(scale([110000, 54], ranges))   # [0.5, 0.5]`,
						Notes: "Fit the ranges once on the training data, then scale both the training points and every query with them.",
					},
				},
			},
			{
				Name:        "MapReduce: map, shuffle, reduce",
				Description: "A recipe for computing per-key aggregates over data spread across many machines: map each record to (key, value) pairs, shuffle so all pairs with the same key reach the same reducer, then reduce each key's list of values to one result.",
				WhyItWorks:  `Map runs on each chunk of input independently, so 1,000 machines can map 1,000 chunks at the same time with no communication. The shuffle is a distributed group-by: it sends each pair to the reducer chosen by hash(key) % number_of_reducers, so every value for a given key ends up in one place, which is the same job a dictionary of lists does on one machine (Chapter 2). Reduce then also runs in parallel, one key at a time, because keys are independent. The answer is correct as long as the reduce function only depends on the key's values and not on their order, which holds for sums, counts, maximums and minimums. When reduce is also associative (like +), each mapper can pre-reduce its own pairs (a 'combiner'), sending ('the', 5000) instead of five thousand ('the', 1) pairs across the network.`,
				Intuition:   `Counting votes in a national election. Each polling station counts its own ballots (map, in parallel), the per-candidate tallies are sent to one desk per candidate (shuffle, grouped by key), and each desk adds up its tallies (reduce, in parallel across candidates). Polling stations are mappers, candidates are keys, and the desks are reducers.`,
				CoreIdeas: []string{
					"Map: record -> list of (key, value) pairs, independent per record, so it parallelises perfectly",
					"Shuffle: group pairs by key, usually by sending each pair to reducer hash(key) % R",
					"Reduce: (key, list of values) -> result, independent per key",
					"A combiner pre-reduces on each mapper when reduce is associative and commutative (sums, counts, max)",
					"On one machine the same pipeline is a loop, a dictionary of lists and a final loop over the dictionary",
				},
				CommonMistakes: []string{
					"Writing a reduce that depends on the order of values (e.g. 'take the first value'). The shuffle gives no ordering guarantee; use order-independent operations or sort inside reduce.",
					"Using a combiner for an operation that is not associative, such as an average. Average of averages is wrong when groups differ in size; emit (sum, count) pairs and divide at the end.",
					"Choosing a key with very few distinct values (or one huge one), so one reducer gets almost all the data and the job waits for it. Pick keys that spread the work.",
				},
				Examples: []CodeExample{
					{
						ID:       "mapreduce",
						Title:    "Map, shuffle and reduce on one machine",
						Language: "python",
						Snippet: `from collections import defaultdict

def run_mapreduce(records, map_fn, reduce_fn):
    pairs = []
    for record in records:            # map: one record at a time
        pairs.extend(map_fn(record))
    groups = defaultdict(list)        # shuffle: group by key
    for key, value in pairs:
        groups[key].append(value)
    return {key: reduce_fn(key, values) for key, values in groups.items()}  # reduce`,
						Notes: "In a real cluster the map loop and the reduce loop run on many machines, and the dictionary becomes network traffic.",
					},
				},
			},
			{
				Name:        "Where to go next",
				Description: "The patterns from this course are the building blocks of more advanced topics. Pick the next topic by asking which pattern it extends, and learn it the same way: worked example first, then implement it yourself, then practise.",
				WhyItWorks:  `Advanced data structures are mostly the course's ideas with one more twist. Balanced BSTs (AVL, red-black trees) are Chapter 9's BSTs plus rotations that keep the height O(log n). Segment trees and Fenwick trees answer range-sum queries in O(log n) by combining divide and conquer (Chapter 8) with trees. k-d trees speed up the nearest-neighbour search of this chapter by splitting space like a BST splits numbers. Bloom filters combine hashing (Chapter 2) with a bit array (Chapter 20) to answer 'have I seen this?' in tiny memory, with occasional false positives. String algorithms such as KMP and rolling hashes extend sliding windows (Chapter 4). Because each one reuses a pattern you already know, learning it is a small step rather than a fresh start, and the evidence-based routine from Chapter 1 (worked examples, retrieval, spacing, interleaving) still applies.`,
				Intuition:   `The course is a set of familiar cities on a map. Each advanced topic is a road leading out of one of them: from Trees to balanced trees and segment trees, from Hash Maps and Bits to Bloom filters, from Sliding Window to string matching. Travel the road that starts at a city you know well.`,
				CoreIdeas: []string{
					"Trees -> balanced BSTs (AVL, red-black), segment trees, Fenwick trees, k-d trees",
					"Hashing + bits -> Bloom filters, consistent hashing, count-min sketch",
					"Sliding window and two pointers -> string matching (KMP, Rabin-Karp rolling hash)",
					"Graphs -> network flow and bipartite matching; DP -> bitmask DP and DP on trees",
					"Systems -> MapReduce, distributed key-value stores, approximate nearest-neighbour search",
				},
				CommonMistakes: []string{
					"Jumping to advanced topics while the basics are shaky. Keep the Review queue going so earlier patterns stay fresh; they are what the new topics are built on.",
					"Only reading about a new structure. Implement it from scratch once and trace it on a small input, the same way the chapters of this course did.",
				},
				Examples: []CodeExample{
					{
						ID:       "bloom-filter",
						Title:    "A tiny Bloom filter: hashing plus a bit array",
						Language: "python",
						Snippet: `class BloomFilter:
    def __init__(self, size=1024, hashes=3):
        self.size = size
        self.hashes = hashes
        self.bits = 0                    # one int used as a bit array (Chapter 20)

    def _positions(self, item):
        return [hash((i, item)) % self.size for i in range(self.hashes)]

    def add(self, item):
        for pos in self._positions(item):
            self.bits |= 1 << pos        # set bit pos

    def might_contain(self, item):
        return all((self.bits >> pos) & 1 for pos in self._positions(item))`,
						Notes: "might_contain never says False for an added item, but can say True for one never added (a false positive). That trade buys very small memory.",
					},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Classify a point with 3-nearest neighbours",
				Problem: "Training points (already on comparable scales): (1,1) A, (2,1) A, (4,4) B, (5,4) B, (4,5) B. Classify the query (2,2) with k = 3.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Measure every distance",
						Explanation: "Compute the squared Euclidean distance from the query to each training point. Squaring keeps the same order as the true distance and avoids square roots.",
						Code: `distances = []
for point, label in zip(points, labels):
    d2 = sum((p - q) ** 2 for p, q in zip(point, query))
    distances.append((d2, label))`,
						State: "(1,1): 1+1=2 A; (2,1): 0+1=1 A; (4,4): 4+4=8 B; (5,4): 9+4=13 B; (4,5): 4+9=13 B",
					},
					{
						Subgoal:     "Keep the k closest",
						Explanation: "Only the 3 smallest distances matter. heapq.nsmallest keeps a heap of size k while scanning, which costs O(n log k) instead of sorting everything.",
						Code:        "nearest = heapq.nsmallest(k, distances)",
						State:       "nearest = [(1, A), (2, A), (8, B)]",
					},
					{
						Subgoal:     "Vote",
						Explanation: "Count the labels of the k neighbours. With k = 3 and two classes a tie is impossible.",
						Code:        "votes = Counter(label for _, label in nearest)",
						State:       "votes = {A: 2, B: 1}",
					},
					{
						Subgoal:     "Return the winning label",
						Explanation: "The majority label is the prediction. With k = 5 the vote would be A 2, B 3 and the answer would flip to B, which is why k is chosen by testing on held-out data.",
						Code:        "return votes.most_common(1)[0][0]",
						State:       "returns 'A'",
					},
				},
				Code: `import heapq
from collections import Counter

def knn_classify(points, labels, query, k):
    distances = []
    for point, label in zip(points, labels):
        d2 = sum((p - q) ** 2 for p, q in zip(point, query))
        distances.append((d2, label))
    nearest = heapq.nsmallest(k, distances)
    votes = Counter(label for _, label in nearest)
    return votes.most_common(1)[0][0]`,
			},
			{
				Title:   "Count words with MapReduce",
				Problem: "Count how often each word appears in the lines [\"the cat\", \"the dog the end\"], structured as map, shuffle and reduce so that each stage could run on many machines.",
				Steps: []WorkedStep{
					{
						Subgoal:     "Map: emit pairs",
						Explanation: "Each line is turned into (word, 1) pairs on its own, without looking at any other line. That independence is what lets thousands of machines map at once.",
						Code: `def map_words(line):
    return [(word, 1) for word in line.split()]`,
						State: "\"the cat\" -> (the,1) (cat,1); \"the dog the end\" -> (the,1) (dog,1) (the,1) (end,1)",
					},
					{
						Subgoal:     "Shuffle: group by key",
						Explanation: "Bring all values for the same key together. On one machine this is a dictionary of lists; in a cluster each pair is sent to the reducer numbered hash(key) % R.",
						Code: `groups = defaultdict(list)
for key, value in pairs:
    groups[key].append(value)`,
						State: "the: [1, 1, 1]; cat: [1]; dog: [1]; end: [1]",
					},
					{
						Subgoal:     "Reduce: combine each group",
						Explanation: "Each key's values are combined independently of other keys. Adding is order-independent, so it does not matter in which order the shuffle delivered the 1s.",
						Code: `def reduce_count(key, values):
    return sum(values)`,
						State: "the -> 3; cat -> 1; dog -> 1; end -> 1",
					},
					{
						Subgoal:     "Collect the results",
						Explanation: "The output is one result per key. A combiner could have summed the two (the, 1) pairs from the second line on its mapper, sending (the, 2) instead.",
						Code:        "return {key: reduce_count(key, values) for key, values in groups.items()}",
						State:       "{'the': 3, 'cat': 1, 'dog': 1, 'end': 1}",
					},
				},
				Code: `from collections import defaultdict

def map_words(line):
    return [(word, 1) for word in line.split()]

def reduce_count(key, values):
    return sum(values)

def word_count(lines):
    pairs = []
    for line in lines:
        pairs.extend(map_words(line))
    groups = defaultdict(list)
    for key, value in pairs:
        groups[key].append(value)
    return {key: reduce_count(key, values) for key, values in groups.items()}`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "KNN regression: predict a price",
				Prompt: "Predict a house price as the average price of the k nearest houses. Reuse the subgoals Measure every distance and Keep the k closest, then replace the vote with an average.",
				Code: `import heapq

def knn_predict(points, prices, query, k):
    distances = []
    for point, price in zip(points, prices):
        d2 = sum((p - q) ** 2 for p, q in zip(point, query))
        distances.append((d2, price))
    nearest = heapq.nsmallest(____, distances)
    total = 0
    for _, price in nearest:
        total += ____
    return total / ____`,
				Blanks:      []string{"k", "price", "k"},
				Explanation: "Regression differs from classification only in the last step: instead of counting labels, average the neighbours' values. nsmallest(k, ...) keeps the k closest (distance, price) pairs, the loop adds their prices, and dividing by k gives the mean.",
			},
			{
				Title:  "MapReduce: hottest temperature per city",
				Prompt: "Each record is a line \"city,temperature\". Fill in the map so the city is the key, and the reduce so each city's result is its maximum temperature.",
				Code: `def map_reading(line):
    city, temp = line.split(",")
    return [(____, int(temp))]

def reduce_max(city, temps):
    return ____`,
				Blanks:      []string{"city", "max(temps)"},
				Explanation: "The key decides what gets grouped together, so it must be the city. max is order-independent and associative, so it is a safe reduce and could also be used as a combiner on each mapper. For [\"oslo,3\", \"rome,21\", \"oslo,7\"] the result is {oslo: 7, rome: 21}.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt: "Which earlier pattern fits? \"Find the 5 stored points closest to a query among 1,000,000 points, once.\"",
				Choices: []string{
					"Sort all 1,000,000 distances and take the first 5: O(n log n)",
					"Keep a max-heap of size 5 while scanning the distances (Chapter 11): O(n log k)",
					"Binary search (Chapter 6) over the distances",
					"Two pointers from both ends of the list (Chapter 3)",
				},
				Answer:      1,
				Explanation: "Picking the k smallest of n values is the top-k pattern from Chapter 11: a heap of size k is updated once per point, O(n log k), roughly n * 2.3 steps here. Sorting works but does O(n log n) work, about n * 20. Binary search needs a sorted order that the distances do not have, and two pointers need a sorted sequence and a pair condition.",
			},
			{
				Prompt: "Features are income in dollars (20,000 to 200,000) and age in years (18 to 90), unscaled. Which feature mostly decides who the nearest neighbours are?",
				Choices: []string{
					"Age, because it is listed second",
					"Income, because its differences are thousands of times larger and dominate the squared distance",
					"Both equally, because Euclidean distance treats all features the same",
					"Neither: KNN ignores feature values",
				},
				Answer:      1,
				Explanation: "Euclidean distance adds squared differences, so a 1,000-dollar gap adds 1,000,000 while a 30-year gap adds only 900. Distance treats the numbers equally, not the features, which is exactly why each feature must be scaled to a comparable range first.",
			},
			{
				Prompt: "Trace the word-count MapReduce on [\"a b a\", \"b a\"]. What is the result?",
				Choices: []string{
					"{'a': 3, 'b': 2}",
					"{'a': 2, 'b': 1}",
					"{'a': 1, 'b': 1}",
					"{'a b a': 1, 'b a': 1}",
				},
				Answer:      0,
				Explanation: "Map emits (a,1) (b,1) (a,1) from the first line and (b,1) (a,1) from the second. The shuffle groups a: [1,1,1] and b: [1,1], and reduce sums each list: a = 3, b = 2. {'a': 2, 'b': 1} is the count for the first line only, which is what you would get if reducers saw only one mapper's output.",
			},
			{
				Prompt: "Chapter 1 review: a KNN classifier stores n training examples with d features each and answers one query by computing every distance. What is the cost of one query, and what happens when n doubles?",
				Choices: []string{
					"O(d): the query time does not depend on n",
					"O(n * d): computing each distance costs O(d), so doubling n doubles the time",
					"O(log n): like binary search, doubling n adds one step",
					"O(n^2 * d): every pair of training points is compared",
				},
				Answer:      1,
				Explanation: "Each of the n distances needs d subtractions and squares, so one query is O(n * d) (plus O(n log k) for the heap). The cost grows linearly with n: twice the data means twice the time per query. That is why large systems build indexes such as k-d trees or approximate nearest-neighbour structures.",
			},
			{
				Prompt:      "Why does MapReduce need the shuffle step, and which data structure from Chapter 2 does it imitate on a single machine?",
				Explanation: "After the map step, the pairs for one key (say every ('the', 1)) are scattered across all mappers, but reduce needs all values of a key in one place to combine them. The shuffle routes every pair with the same key to the same reducer, typically by hash(key) % R. On one machine this is a hash map from key to a list of values (a defaultdict(list)), the group-by pattern from Chapter 2.",
			},
			{
				Prompt:      "Name one topic you would study next, the earlier chapter whose pattern it extends, and how you would learn it using the routine from this course.",
				Explanation: "Model answer: segment trees, which extend divide and conquer (Chapter 8) and binary trees (Chapter 9) to answer range-sum or range-minimum queries in O(log n) with O(log n) updates. Learn it by studying one worked example (build a segment tree for [2, 1, 5, 3] and trace a query on a sub-range), implementing it from scratch, explaining why each node stores the combination of its children's ranges, then solving two or three practice problems spaced over a few days. Any topic works as long as you name the pattern it builds on.",
			},
		},
	}
}
