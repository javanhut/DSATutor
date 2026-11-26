package judge

import (
	"testing"

	"dsatutor/internal/practice"
)

// TestGreedyDrillsRejectMistakes checks that each greedy foundations drill
// fails the mistake its tests and notes promise to catch. In greedy the
// mistakes that matter are wrong greedy choices (the wrong sort key, the wrong
// thing to give away), so these cases are plausible-but-wrong rules that must
// be defeated by a counterexample in the tests.
func TestGreedyDrillsRejectMistakes(t *testing.T) {
	if testing.Short() {
		t.Skip("runs Python for every test case")
	}
	cases := []struct{ id, why, code string }{
		{"drill-assign-cookies", "sorts only the cookies, not the children", `def findContentChildren(g, s):
    child = 0
    for cookie in sorted(s):
        if child < len(g) and cookie >= g[child]:
            child += 1
    return child`},
		{"drill-assign-cookies", "sorts only the children, not the cookies", `def findContentChildren(g, s):
    g = sorted(g)
    child = 0
    for cookie in s:
        if child < len(g) and cookie >= g[child]:
            child += 1
    return child`},
		{"drill-assign-cookies", "gives each child the biggest cookie", `def findContentChildren(g, s):
    cookies = sorted(s)
    count = 0
    for need in sorted(g):
        if cookies and cookies[-1] >= need:
            cookies.pop()
            count += 1
    return count`},
		{"drill-assign-cookies", "pairs sorted children and cookies by index", `def findContentChildren(g, s):
    return sum(1 for need, cookie in zip(sorted(g), sorted(s)) if cookie >= need)`},
		{"drill-lemonade-change", "gives three 5s for a 20 before a 10 and a 5", `def lemonadeChange(bills):
    fives = tens = 0
    for bill in bills:
        if bill == 5:
            fives += 1
        elif bill == 10:
            if fives == 0:
                return False
            fives -= 1
            tens += 1
        else:
            if fives >= 3:
                fives -= 3
            elif tens > 0 and fives > 0:
                tens -= 1
                fives -= 1
            else:
                return False
    return True`},
		{"drill-lemonade-change", "never gives three 5s for a 20", `def lemonadeChange(bills):
    fives = tens = 0
    for bill in bills:
        if bill == 5:
            fives += 1
        elif bill == 10:
            if fives == 0:
                return False
            fives -= 1
            tens += 1
        else:
            if tens > 0 and fives > 0:
                tens -= 1
                fives -= 1
            else:
                return False
    return True`},
		{"drill-lemonade-change", "does not keep the 10s it receives", `def lemonadeChange(bills):
    fives = 0
    for bill in bills:
        if bill == 5:
            fives += 1
        elif bill == 10:
            if fives == 0:
                return False
            fives -= 1
        else:
            if fives >= 3:
                fives -= 3
            else:
                return False
    return True`},
		{"drill-max-meetings", "sorts by start time", `def maxMeetings(meetings):
    count = 0
    last_end = 0
    for start, end in sorted(meetings):
        if start >= last_end:
            count += 1
            last_end = end
    return count`},
		{"drill-max-meetings", "takes the shortest meeting first", `def maxMeetings(meetings):
    chosen = []
    for start, end in sorted(meetings, key=lambda m: m[1] - m[0]):
        if all(end <= s or start >= e for s, e in chosen):
            chosen.append((start, end))
    return len(chosen)`},
		{"drill-max-meetings", "uses > so back-to-back meetings clash", `def maxMeetings(meetings):
    count = 0
    last_end = -1
    for start, end in sorted(meetings, key=lambda m: m[1]):
        if start > last_end:
            count += 1
            last_end = end
    return count`},
		{"drill-max-meetings", "moves the cutoff for a skipped meeting", `def maxMeetings(meetings):
    count = 0
    last_end = 0
    for start, end in sorted(meetings, key=lambda m: m[1]):
        if start >= last_end:
            count += 1
        last_end = end
    return count`},
		{"drill-greedy-coins-check", "returns the greedy count for both entries", `def compareGreedyCoins(coins, amount):
    greedy = 0
    remaining = amount
    for coin in sorted(coins, reverse=True):
        greedy += remaining // coin
        remaining %= coin
    return [greedy, greedy]`},
		{"drill-greedy-coins-check", "returns the best count for both entries", `def compareGreedyCoins(coins, amount):
    best = [0] + [amount + 1] * amount
    for total in range(1, amount + 1):
        for coin in coins:
            if coin <= total:
                best[total] = min(best[total], best[total - coin] + 1)
    return [best[amount], best[amount]]`},
		{"drill-greedy-coins-check", "takes the coins in the order given", `def compareGreedyCoins(coins, amount):
    greedy = 0
    remaining = amount
    for coin in coins:
        greedy += remaining // coin
        remaining %= coin
    best = [0] + [amount + 1] * amount
    for total in range(1, amount + 1):
        for coin in coins:
            if coin <= total:
                best[total] = min(best[total], best[total - coin] + 1)
    return [greedy, best[amount]]`},
		{"drill-boats-to-save-people", "pairs sorted neighbours", `def numRescueBoats(people, limit):
    people = sorted(people)
    boats = 0
    i = 0
    while i < len(people):
        if i + 1 < len(people) and people[i] + people[i + 1] <= limit:
            i += 1
        i += 1
        boats += 1
    return boats`},
		{"drill-boats-to-save-people", "does not sort", `def numRescueBoats(people, limit):
    lightest = 0
    heaviest = len(people) - 1
    boats = 0
    while lightest <= heaviest:
        if people[lightest] + people[heaviest] <= limit:
            lightest += 1
        heaviest -= 1
        boats += 1
    return boats`},
		{"drill-boats-to-save-people", "stops when the pointers meet", `def numRescueBoats(people, limit):
    people = sorted(people)
    lightest = 0
    heaviest = len(people) - 1
    boats = 0
    while lightest < heaviest:
        if people[lightest] + people[heaviest] <= limit:
            lightest += 1
        heaviest -= 1
        boats += 1
    return boats`},
		{"drill-boats-to-save-people", "moves the lightest when the pair does not fit", `def numRescueBoats(people, limit):
    people = sorted(people)
    lightest = 0
    heaviest = len(people) - 1
    boats = 0
    while lightest <= heaviest:
        if people[lightest] + people[heaviest] <= limit:
            heaviest -= 1
        lightest += 1
        boats += 1
    return boats`},
		{"drill-max-units-on-truck", "sorts by the number of boxes", `def maximumUnits(boxTypes, truckSize):
    total = 0
    room = truckSize
    for boxes, units in sorted(boxTypes, key=lambda b: b[0], reverse=True):
        take = min(boxes, room)
        total += take * units
        room -= take
    return total`},
		{"drill-max-units-on-truck", "sorts by total units in the pile", `def maximumUnits(boxTypes, truckSize):
    total = 0
    room = truckSize
    for boxes, units in sorted(boxTypes, key=lambda b: b[0] * b[1], reverse=True):
        take = min(boxes, room)
        total += take * units
        room -= take
    return total`},
		{"drill-max-units-on-truck", "sorts ascending by units", `def maximumUnits(boxTypes, truckSize):
    total = 0
    room = truckSize
    for boxes, units in sorted(boxTypes, key=lambda b: b[1]):
        take = min(boxes, room)
        total += take * units
        room -= take
    return total`},
		{"drill-max-units-on-truck", "does not cap the boxes at the room left", `def maximumUnits(boxTypes, truckSize):
    total = 0
    room = truckSize
    for boxes, units in sorted(boxTypes, key=lambda b: b[1], reverse=True):
        if room <= 0:
            break
        total += boxes * units
        room -= boxes
    return total`},
	}
	for _, c := range cases {
		c := c
		t.Run(c.id+"/"+c.why, func(t *testing.T) {
			t.Parallel()
			var p *practice.Problem
			for _, q := range practice.EmbeddedProblems {
				if q.ID == c.id {
					p = q
				}
			}
			if p == nil {
				t.Fatalf("no problem %s", c.id)
			}
			for i := range p.TestCases {
				res, err := Run(p, c.code, &p.TestCases[i], Options{})
				if err != nil {
					t.Fatalf("case %d: %v", i+1, err)
				}
				if !res.Passed {
					return
				}
			}
			t.Errorf("code that %s passed every case", c.why)
		})
	}
}
