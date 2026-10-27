package chapter

import "time"

// chapter02 teaches arrays, hash maps and sets, counting, and the
// seen-so-far pattern for the arrays-hashing ladder.
// Storyboards (IDs registered in app.js): hash-table.
func chapter02() Chapter {
	return Chapter{
		Number:        2,
		Slug:          "arrays-hashing",
		Title:         "Arrays and Hash Maps",
		Category:      "arrays-hashing",
		Prerequisites: []int{1},
		Summary: "Arrays find an item by its position in O(1); hash maps and sets find an item by its value in O(1) on average. " +
			"Learn how hashing makes that possible and what it costs, then the two patterns behind most of the Arrays & Hashing ladder: " +
			"counting with a dictionary, and remembering what you've seen so far.",
		Motivation: `You're given 100,000 numbers and asked whether any value appears twice. Chapter 1's recipe starts with a brute force: compare every pair. That is about 5 billion comparisons, minutes of work. Sorting first and comparing neighbours brings it down to O(n log n).

Now look at what the brute force is really asking at each number: "have I seen this value before?" If that question could be answered in one step, a single pass of 100,000 steps would be enough. A Python list can't do it: "x in my_list" walks the list from the start, so it is O(n) per question and O(n^2) overall, no better than the brute force. A set can, because it doesn't search at all. It runs x through a hash function that computes where x would be stored, and looks only there.

This chapter shows how that works, why it is O(1) on average but not always, and how to turn questions like "have I seen it?", "how many times?" and "where was its partner?" into dictionary lookups. That one move, trading O(n) memory for O(1) lookups, turns many O(n^2) brute forces into O(n) solutions.`,
		Objectives: []string{
			"Explain why reading nums[i] is O(1) but inserting at the front of a list is O(n).",
			"Explain how a hash table finds a key in O(1) on average, what a collision is, and why the worst case is O(n).",
			"Choose between a set (seen or not), a dict of counts (how many) and a dict of positions (where) for a given question.",
			"Implement the seen-so-far pattern (look up before you store) and trace it on a small input.",
			"Count items with a dict or Counter, and decide whether two collections hold the same items by comparing counts.",
		},
		WhenToUse: []string{
			"The task asks whether something has appeared before: duplicates, repeats, \"already seen\", \"first repeated\".",
			"For each item the partner you need is fully determined, like target - x or x + k: look it up instead of searching for it.",
			"Order doesn't matter but how many times does: anagrams, \"same letters\", frequencies, \"most common\".",
			"Items must be grouped by a shared property (same sorted letters, same signature): use the property as a dictionary key.",
			"Your brute force has an inner loop that only searches for a known value: replace that loop with a set or dict lookup.",
			"The input is unsorted and n is up to 10^5 or more, so O(n^2) is too slow and you may use O(n) extra memory.",
		},
		Concepts: []Concept{
			{
				Name:        "Arrays (Python lists)",
				Description: "A block of numbered slots. Reaching a slot by its index is O(1); inserting or deleting anywhere but the end shifts the rest, which is O(n).",
				WhyItWorks: `An array stores its slots side by side in memory, all the same size (in a Python list, each slot holds a reference to an object). To find slot i, the computer computes start + i * slot_size: one multiplication and one addition, whatever i and n are. That is why nums[i] is O(1).

The same layout makes inserting expensive. To put a value at the front, every existing value must move one slot to the right to make room: n moves, O(n). Deleting from the middle shifts the tail left. Appending at the end is cheap: Python reserves spare slots at the end and, when they run out, copies the list into a block about twice as big, so the occasional O(n) copy is spread over many O(1) appends ("amortised O(1)"). Searching for a value with x in nums has no shortcut: the array knows positions, not values, so it checks every slot, O(n).`,
				Intuition: "A row of numbered lockers. To open locker 47 you walk straight to it, because the numbers tell you where it is (O(1) by index). But to find the locker holding a red bag you have to open them one by one (O(n) by value). And to make room for a new locker at the front, everyone's belongings have to shift down one locker (O(n) insert).",
				CoreIdeas: []string{
					"nums[i] and nums[i] = v are O(1): the address is computed, not searched for.",
					"nums.append(v) and nums.pop() at the end are O(1) (amortised).",
					"nums.insert(0, v), nums.pop(0) and deleting from the middle are O(n): the rest must shift.",
					"x in nums, nums.index(x) and nums.count(x) are O(n): arrays know positions, not values.",
					"Slicing nums[a:b] copies b - a items, so it costs O(b - a) time and memory.",
				},
				CommonMistakes: []string{
					"Using x in some_list inside a loop and calling it O(n). Each membership test scans the list, so the loop is O(n^2); use a set when you need membership.",
					"Using nums.pop(0) or nums.insert(0, x) in a loop to treat a list as a queue. Each call is O(n); use collections.deque, whose popleft is O(1).",
					"Slicing inside a loop (nums[i:] or nums[1:]) and forgetting that each slice copies the rest of the list.",
				},
				Examples: []CodeExample{
					{
						ID:       "array-costs",
						Title:    "What each list operation costs",
						Language: "python",
						Snippet: `nums = [4, 8, 15, 16, 23]
x = nums[2]          # O(1): jump to slot 2 -> 15
nums[2] = 99         # O(1): overwrite in place
nums.append(42)      # O(1) amortised: spare room at the end
nums.insert(0, 7)    # O(n): every value shifts one slot right
found = 23 in nums   # O(n): checks slots one by one`,
						Notes: "Lists are fast by position and slow by value. When you need fast by value, reach for a set or a dict.",
					},
				},
			},
			{
				Name:        "Hash maps and sets",
				Description: "A dict maps keys to values and a set stores keys alone. Both find a key in O(1) on average by hashing it to a bucket instead of searching.",
				WhyItWorks: `A hash table is an array of buckets plus a hash function. To store key k, Python computes hash(k), a big integer that is always the same for equal keys, and reduces it to a bucket index such as hash(k) % number_of_buckets. To look k up later it computes the same index and checks only that bucket. No other bucket can hold k, so nothing else needs to be searched.

Two different keys can land in the same bucket: a collision. It's unavoidable, since there are far more possible keys than buckets. The table handles it by keeping several entries per bucket (chaining) or by probing the next free slot (open addressing, which Python uses); either way a lookup compares k against the few keys in its spot. The table keeps those groups small by resizing: when it gets about two-thirds full (its load factor), it moves everything into a bigger array. With a good hash function the keys spread evenly, each bucket holds O(1) keys on average, and lookups, inserts and deletes are O(1) on average. The worst case, every key colliding into one bucket, is O(n), but it doesn't happen with Python's built-in hashes on ordinary data. Keys must be hashable: an immutable value whose hash can never change, like an int, string or tuple. Lists and dicts can't be keys, because changing one after storing it would leave it in the wrong bucket.`,
				Intuition: "A coat check with 100 numbered hooks. When you hand over your coat, the attendant turns your name into a hook number with a fixed rule (the hash function) and hangs it there. When you come back, the same rule gives the same hook, so they go straight to it instead of checking all 100. Sometimes two guests' names give the same hook (a collision); then that hook holds two coats and the attendant checks both, which is still quick. If the hooks get crowded, the club installs more hooks and rehangs every coat (resizing).",
				CoreIdeas: []string{
					"x in some_set, d[key], d[key] = v and del d[key] are O(1) on average.",
					"A set answers \"seen or not?\"; a dict answers \"what do I know about this key?\" (a count, an index, a list).",
					"Collisions are normal; resizing keeps buckets small, so the average stays O(1). The worst case is O(n).",
					"Keys must be hashable (int, str, tuple of hashables). Convert a list to a tuple to use it as a key.",
					"The O(1) lookups cost O(n) extra memory: you trade space for time.",
					"Dicts keep insertion order in Python 3.7+, but sets have no meaningful order.",
				},
				CommonMistakes: []string{
					"Using a list as a dict key or set element, which raises TypeError: unhashable type: 'list'. Use tuple(the_list).",
					"Reading d[key] for a key that may be missing, which raises KeyError. Use d.get(key, default), key in d, or collections.defaultdict.",
					"Assuming a set or dict is sorted. Iterate over sorted(d) if order matters.",
					"Claiming O(1) worst case. Hash operations are O(1) on average; say so when stating complexity.",
				},
				Examples: []CodeExample{
					{
						ID:       "set-and-dict",
						Title:    "Set and dict basics",
						Language: "python",
						Snippet: `seen = set()
seen.add(5)
print(5 in seen, 7 in seen)      # True False, each O(1) on average

index_of = {}
index_of["apple"] = 0            # store key -> value
print(index_of.get("pear", -1))  # -1: a safe lookup for a missing key

key = tuple([1, 0, 2])           # lists can't be keys; tuples can
groups = {key: ["bca"]}`,
						Notes: "Pick a set when presence is all you need, and a dict when you need to remember something about each key.",
					},
				},
			},
			{
				Name:        "Counting with a dict or Counter",
				Description: "Build a value -> how-many-times table in one pass, then answer questions by comparing or reading counts.",
				WhyItWorks: `Many questions don't care where items are, only how many of each there are: are two words anagrams, which value is most common, can these letters spell that note. A count table keeps exactly that information and throws away the order. Building it touches each item once with an O(1) dict update, so it costs O(n) time and O(k) space for k distinct items.

Two collections are rearrangements of each other exactly when their count tables are equal: same keys, same counts. Comparing tables costs O(k), not O(n^2) of pairwise matching. When the keys come from a small fixed alphabet (26 lowercase letters), a list of 26 counters indexed by ord(ch) - ord('a') works too, and its size doesn't grow with the input, so it is O(1) space. A tuple of those counts is hashable, so it can itself be a dict key, which is how you group items that share the same counts.`,
				Intuition: "Counting votes after an election. Nobody re-reads the ballots in order to compare two polling stations; they keep a tally sheet with one row per candidate and add a mark for each ballot. Two stations voted identically exactly when their tally sheets match, whatever order the ballots came in.",
				CoreIdeas: []string{
					"counts[x] = counts.get(x, 0) + 1 is the basic counting line; collections.Counter(items) does the whole loop.",
					"Equal count tables means the same items in a different order (anagrams).",
					"Counting forgets order: use it when the question doesn't depend on positions.",
					"For a fixed small alphabet, a 26-slot list of counts is O(1) space and its tuple can be a dict key.",
					"Counts can also be spent: decrement as you match items from a second collection, and fail when a count would go below 0.",
				},
				CommonMistakes: []string{
					"Using a set where counts matter: set(\"aab\") == set(\"abb\") is True. Keep the counts.",
					"Writing counts[x] += 1 before the key exists, which raises KeyError. Use counts.get(x, 0) + 1, Counter or defaultdict(int).",
					"Comparing only that every item of one collection appears in the other. \"ab\" and \"a\" pass that test; compare full counts or check lengths first.",
				},
				Examples: []CodeExample{
					{
						ID:       "counting",
						Title:    "Counting two ways",
						Language: "python",
						Snippet: `from collections import Counter

def letter_counts(word):
    counts = {}
    for ch in word:
        counts[ch] = counts.get(ch, 0) + 1
    return counts

print(letter_counts("banana"))           # {'b': 1, 'a': 3, 'n': 2}
print(Counter("listen") == Counter("silent"))  # True: same counts`,
						Notes: "Counter is a dict subclass, so everything you know about dicts applies to it.",
					},
				},
			},
			{
				Name:        "The seen-so-far pattern",
				Description: "Walk the input once; for each item, look up what you need among the items already passed, then store the current item for the ones still to come.",
				WhyItWorks: `A brute force that compares each item with every earlier item does O(n) work per item because it searches. The seen-so-far pattern keeps the earlier items in a set or dict, so that search becomes one O(1) lookup. The invariant that makes it correct: when you reach index i, the structure describes exactly nums[0..i-1]. Any pair (i, j) with i < j is then considered when the loop reaches j, because nums[i] is already stored. So nothing is missed.

The order inside the loop matters. Looking up before storing means the current item is never matched with itself, which is what you want when the two items must be at different positions. What you store depends on the question: a set for "has it appeared?", value -> index for "where was it?", value -> count for "how many earlier items qualify?".`,
				Intuition: "Meeting people at a party and looking for someone from your home town. You don't re-interview every guest each time someone new arrives. You keep a guest list in your head (the seen-so-far set). Each new arrival, you first check the list for a match (look up), then add them to it (store). You'd never match a person with themselves, because you check before adding them.",
				CoreIdeas: []string{
					"Choose what to remember: a set (presence), value -> index (position) or value -> count (how many).",
					"Look up before you store: check for the partner of the current item among earlier items.",
					"Record the current item so later items can find it.",
					"Return the answer as soon as it is known, or after the loop.",
					"Turns an O(n^2) pair search into O(n) time with O(n) extra space.",
				},
				CommonMistakes: []string{
					"Storing the current item before looking up its partner, so an item can pair with itself: Two Sum on [3, 2, 4] with target 6 would return [0, 0].",
					"Storing earlier items in a list. The lookup x in list is O(n), and the pattern is O(n^2) again.",
					"Storing the wrong thing: a set when the answer needs an index, or overwriting a count instead of adding to it.",
				},
				Examples: []CodeExample{
					{
						ID:       "seen-so-far",
						Title:    "First value that repeats",
						Language: "python",
						Snippet: `def first_repeat(nums):
    seen = set()                 # choose what to remember
    for x in nums:
        if x in seen:            # look up before you store
            return x
        seen.add(x)              # record the current item
    return None                  # no value repeats`,
						Notes: "For [3, 1, 4, 1, 5, 3] this returns 1: the second 1 comes before the second 3.",
					},
				},
			},
		},
		Animations: []Storyboard{
			{
				ID:    "hash-table",
				Title: "Hash Table Operations",
				Goal:  "See a key turned into a bucket index, a collision handled, and a lookup that checks one bucket instead of the whole table.",
				MemoryTips: []string{
					"Choose what to remember: a set for presence, a dict for a count or an index.",
					"Hash the key: hash(key) % buckets names the only bucket the key can be in.",
					"Look up before you store: check that bucket for the key (or its partner) first.",
					"Record the current item: add it to its bucket; a shared bucket is a collision, and that's fine.",
					"Expect O(1) on average: resizing keeps buckets small; the worst case, everything in one bucket, is O(n).",
				},
				Steps: []StoryboardStep{
					{Cue: "hash", Narration: "To store a key, compute hash(key) and reduce it to a bucket index, for example hash % 8. Equal keys always get the same index, so we will know where to look later.", VisualHint: "Key -> hash -> bucket index", Duration: Duration(3 * time.Second), CodeRef: "hash_table: hash(key) -> bucket"},
					{Cue: "collision", Narration: "Two different keys can land in the same bucket. That's a collision; the bucket simply holds both entries. Resizing when the table gets full keeps such groups small.", VisualHint: "Linked list in bucket", Duration: Duration(3 * time.Second), CodeRef: "hash_table: collision handling"},
					{Cue: "insert", Narration: "Inserting adds the key and its value to its bucket. Only that bucket is touched, so it costs O(1) on average no matter how many keys the table holds.", VisualHint: "Append to chain", Duration: Duration(2 * time.Second), CodeRef: "hash_table: bucket.append"},
					{Cue: "lookup", Narration: "Looking a key up hashes it again, goes straight to its bucket and compares only the few keys there. That's why x in some_set is O(1) on average while x in some_list is O(n).", VisualHint: "O(1) average", Duration: Duration(2 * time.Second), CodeRef: "hash_table: lookup"},
				},
			},
		},
		WorkedExamples: []WorkedExample{
			{
				Title:   "Seen-so-far: is there a pair with difference k?",
				Problem: "Given a list of integers nums and an integer k >= 0, return True if two different positions i and j have nums[i] - nums[j] == k, and False otherwise. Example: nums = [7, 1, 5, 3], k = 4 gives True (5 - 1 = 4).",
				Steps: []WorkedStep{
					{
						Subgoal:     "Understand with examples",
						Explanation: "Work cases by hand, as in Chapter 1. [7, 1, 5, 3] with k = 4: 5 - 1 = 4 and 7 - 3 = 4, so True. With k = 0 we need two equal values at different positions: [2, 9, 2] is True but [2, 9] is False, because a value can't pair with itself.",
						State:       "[7,1,5,3], k=4 -> True; [2,9,2], k=0 -> True; [2,9], k=0 -> False",
					},
					{
						Subgoal:     "Brute force first",
						Explanation: "Check every pair in both orders. That is O(n^2): for n = 100,000, about 5 billion pairs. The waste: at each j we already know the only values that could pair with nums[j] (nums[j] - k and nums[j] + k), yet we scan for them.",
						Code: `for i in range(len(nums)):
    for j in range(len(nums)):
        if i != j and nums[i] - nums[j] == k:
            return True
return False`,
						State: "n = 4 -> 12 ordered pairs checked",
					},
					{
						Subgoal:     "Choose what to remember",
						Explanation: "For the current x, an earlier partner y works if x - y == k (y = x - k) or y - x == k (y = x + k). We only need to know whether such a y has appeared, not where, so a set of earlier values is enough.",
						Code:        `seen = set()`,
						State:       "seen = {}",
					},
					{
						Subgoal:     "Look up before you store",
						Explanation: "Check both possible partners among the earlier values. Looking up before adding x is what makes k = 0 correct: x is not yet in seen, so it can only match an earlier copy of itself.",
						Code: `for x in nums:
    if x - k in seen or x + k in seen:
        return True`,
						State: "x=7: 3? 11? no | x=1: -3? 5? no | x=5: 1? yes -> return True",
					},
					{
						Subgoal:     "Record the current item",
						Explanation: "If x had no partner, store it so later values can pair with it.",
						Code:        `    seen.add(x)`,
						State:       "after 7: {7}; after 1: {7, 1}",
					},
					{
						Subgoal:     "Return the answer",
						Explanation: "Reaching the end means no value ever found a partner among the values before it. Every pair (i, j) was considered when the later of the two was processed, so the answer is False. One pass with O(1) average lookups: O(n) time, O(n) space.",
						Code:        `return False`,
						State:       "[1, 2, 3], k = 5 -> seen grows to {1, 2, 3}, returns False",
					},
				},
				Code: `def has_pair_with_difference(nums, k):
    seen = set()
    for x in nums:
        if x - k in seen or x + k in seen:
            return True
        seen.add(x)
    return False`,
			},
			{
				Title:   "Counting: first unique character",
				Problem: "Given a string s of lowercase letters, return the index of the first character that appears exactly once in s, or -1 if there is none. Example: s = \"loveleetcode\" gives 2 (the 'v').",
				Steps: []WorkedStep{
					{
						Subgoal:     "Understand with examples",
						Explanation: "\"leetcode\": 'l' appears once, so 0. \"loveleetcode\": 'l' and 'o' repeat, 'v' doesn't, so 2. \"aabb\": every letter repeats, so -1. \"first\" means smallest index, so the answer depends on order, even though counting itself doesn't.",
						State:       "\"leetcode\" -> 0; \"loveleetcode\" -> 2; \"aabb\" -> -1",
					},
					{
						Subgoal:     "Brute force first",
						Explanation: "For each position, count how often its character appears with s.count(ch), which scans the whole string. n positions times an O(n) scan is O(n^2). The waste: the count of 'e' is recomputed every time an 'e' is reached.",
						Code: `for i, ch in enumerate(s):
    if s.count(ch) == 1:
        return i
return -1`,
						State: "\"loveleetcode\": count('l')=2, count('o')=2, count('v')=1 -> 2, after 3 full scans",
					},
					{
						Subgoal:     "Count each item",
						Explanation: "Compute every character's count once, in a single pass, and keep them in a dict. At most 26 keys exist, so the table is O(1) space.",
						Code: `counts = {}
for ch in s:
    counts[ch] = counts.get(ch, 0) + 1`,
						State: "\"loveleetcode\" -> {l:2, o:2, v:1, e:4, t:1, c:1, d:1}",
					},
					{
						Subgoal:     "Scan in original order",
						Explanation: "Counting threw away the order, so walk s again from the left and stop at the first character whose count is 1. Each check is an O(1) dict lookup.",
						Code: `for i, ch in enumerate(s):
    if counts[ch] == 1:
        return i`,
						State: "i=0 'l' count 2 | i=1 'o' count 2 | i=2 'v' count 1 -> return 2",
					},
					{
						Subgoal:     "Return the answer",
						Explanation: "If no character has count 1, return -1. Two passes of O(n) each give O(n) time, and the table is O(1) space because the alphabet is fixed.",
						Code:        `return -1`,
						State:       "\"aabb\" -> {a:2, b:2} -> no count of 1 -> -1",
					},
				},
				Code: `def first_unique_char(s):
    counts = {}
    for ch in s:
        counts[ch] = counts.get(ch, 0) + 1
    for i, ch in enumerate(s):
        if counts[ch] == 1:
            return i
    return -1`,
			},
		},
		Completions: []CompletionProblem{
			{
				Title:  "Count the pairs that sum to target",
				Prompt: "Return how many index pairs i < j have nums[i] + nums[j] == target. Use the seen-so-far pattern, remembering how many times each value has appeared. Example: nums = [1, 5, 3, 3, 3], target = 6 gives 4 (1 + 5, and three ways to pick two of the 3s). Target: O(n).",
				Code: `def count_pairs(nums, target):
    count_of = {}  # value -> how many times seen so far
    pairs = 0
    for x in nums:
        pairs += count_of.get(____, 0)
        count_of[x] = ____
    return pairs`,
				Blanks:      []string{"target - x", "count_of.get(x, 0) + 1"},
				Explanation: "Look up before you store: every earlier copy of target - x forms one pair with x, so add how many there were. Then record x by adding one to its count. For [1, 5, 3, 3, 3]: 1 adds 0, 5 adds 1 (the 1), the first 3 adds 0, the second adds 1, the third adds 2, for a total of 4. Storing x before the lookup would wrongly pair a 3 with itself.",
			},
			{
				Title:  "Can the note be built from the magazine?",
				Prompt: "You may cut letters out of the string magazine, using each letter at most once. Return True if you can spell the string note. Example: note = \"aab\", magazine = \"baa\" gives True; note = \"aa\", magazine = \"ab\" gives False. Count the available letters, then spend them.",
				Code: `def can_construct(note, magazine):
    available = {}
    for ch in magazine:
        available[ch] = ____
    for ch in note:
        if available.get(ch, 0) == ____:
            return False
        ____
    return True`,
				Blanks:      []string{"available.get(ch, 0) + 1", "0", "available[ch] -= 1"},
				Explanation: "Count each item in the magazine, then spend one letter for each letter of the note. If a needed letter's count is already 0, the magazine has run out of it. This is counting with a dict: order never matters, only how many of each letter are left. O(n + m) time, O(1) space for a fixed alphabet.",
			},
		},
		CheckQuestions: []CheckQuestion{
			{
				Prompt:      "Why is x in s fast when s is a set but slow when s is a list?",
				Choices:     []string{"Sets are kept sorted, so Python binary-searches them", "A set hashes x to compute the one bucket where x could be, and checks only there; a list has to be scanned slot by slot", "Sets hold fewer items than lists", "Lists are linked lists, so each step follows a pointer"},
				Answer:      1,
				Explanation: "Sets aren't sorted at all; they use hashing. hash(x) points to the only place x could be stored, so the check is O(1) on average. A list knows positions, not values, so x in list compares x with each slot: O(n).",
			},
			{
				Prompt:      "Trace this code with nums = [4, 1, 3, 4] and target = 8. What does it return?\n\nseen = {}\nfor i, x in enumerate(nums):\n    if target - x in seen:\n        return [seen[target - x], i]\n    seen[x] = i\nreturn None",
				Choices:     []string{"[0, 0]", "[0, 3]", "[3, 0]", "None"},
				Answer:      1,
				Explanation: "i=0: 8-4=4 not in {} -> seen={4:0}. i=1: 7 no -> {4:0, 1:1}. i=2: 5 no -> add 3. i=3: 8-4=4 is in seen at index 0 -> return [0, 3]. [0, 0] is what you'd get if x were stored before the lookup, pairing the first 4 with itself.",
			},
			{
				Prompt:      "Which approach fits best: \"Return True if some value appears at least 3 times in an unsorted list of 10^5 numbers\"?",
				Choices:     []string{"Compare every triple of positions (Chapter 1 brute force)", "Count each value with a dict, then check whether any count is at least 3", "Keep a set of values seen so far", "Selection sort the list, then scan it"},
				Answer:      1,
				Explanation: "The question is \"how many times?\", so count with a dict: O(n). A set only remembers seen-or-not and can't tell a second copy from a third. Triples are O(n^3), and selection sort is O(n^2) before the scan even starts; sorting with sorted() then scanning runs would work, but at O(n log n).",
			},
			{
				Prompt:      "Two tasks: (A) find the smallest difference between any two values in a list; (B) decide whether some value equals another value plus k. Which tool fits each, and why: sorting (Chapter 1) or a hash set (this chapter)?",
				Explanation: "(A) needs sorting: the closest partner of a value is its neighbour in sorted order, and a hash set can only answer exact questions (\"is this value present?\"), not \"what's nearest?\". (B) needs a hash set: for each x the partner is exactly x - k or x + k, a known value, so one O(1) lookup per item gives O(n). Rule of thumb: exact partner -> hash; nearest or ordered partner -> sort.",
			},
			{
				Prompt:      "A loop runs n times and does one dict lookup and one dict insert per iteration. What is its average time complexity? (Big O from Chapter 1.)",
				Choices:     []string{"O(1)", "O(n)", "O(n log n)", "O(n^2)"},
				Answer:      1,
				Explanation: "Each iteration does O(1) average work, and n iterations in sequence add up to n * O(1) = O(n). The dict's size doesn't make each lookup slower on average, because hashing goes straight to the right bucket.",
			},
			{
				Prompt:      "What is a hash collision, and why is a hash table still O(1) on average despite collisions?",
				Explanation: "A collision is two different keys landing in the same bucket. It can't be avoided, since there are more possible keys than buckets. The table stores colliding keys together and compares against each on lookup. It stays fast because a good hash spreads keys evenly and the table resizes (to roughly double) when it gets about two-thirds full, so each bucket holds only a constant number of keys on average. Only if almost every key collided would a lookup degrade to O(n).",
			},
			{
				Prompt:      "Which of these can be used as a dictionary key?",
				Choices:     []string{"[1, 2]", "(1, 2)", "{1: 2}", "{1, 2}"},
				Answer:      1,
				Explanation: "Keys must be hashable, meaning immutable so their hash can never change. Tuples of hashable values qualify. Lists, dicts and sets are mutable, so Python refuses to hash them (TypeError: unhashable type). Convert a list with tuple(...) to use it as a key.",
			},
			{
				Prompt:      "What does nums.insert(0, x) cost on a Python list of n elements?",
				Choices:     []string{"O(1)", "O(log n)", "O(n)", "O(n^2)"},
				Answer:      2,
				Explanation: "Lists store items side by side, so putting something at the front means shifting all n items one slot right: O(n). Only appending at the end is O(1) (amortised). Use collections.deque if you need fast inserts at the front.",
			},
		},
	}
}
