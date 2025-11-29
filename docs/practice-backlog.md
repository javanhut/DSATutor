# Practice Backlog (Blind 75 Gaps)

This file tracks problems that are still missing from the Blind 75 / NeetCode-style coverage in `internal/practice/embedded_problems.go`.

**Status Update**: Major progress made - 67 new problems added with full hints, visualizations, and linked tutorials.

- Current embedded problems: **96** (was 29)
- Remaining in this backlog: ~45
- All problems include 3-level hints, solutions with walkthroughs, and chapter links

## Implemented Categories (Complete or Nearly Complete)

### Stack (6/6) - COMPLETE
All implemented with hints, solutions, and chapter links:
- [x] Min Stack
- [x] Evaluate Reverse Polish Notation
- [x] Generate Parentheses
- [x] Daily Temperatures
- [x] Car Fleet
- [x] Largest Rectangle in Histogram

### Binary Search (6/6) - COMPLETE
- [x] Search a 2D Matrix
- [x] Koko Eating Bananas
- [x] Find Minimum in Rotated Sorted Array
- [x] Search in Rotated Sorted Array
- [x] Time Based Key-Value Store
- [x] Median of Two Sorted Arrays

### Linked List (7/7) - COMPLETE
- [x] Merge Two Sorted Lists
- [x] Reorder List
- [x] Remove Nth Node From End
- [x] Copy List with Random Pointer
- [x] Add Two Numbers
- [x] Merge K Sorted Lists
- [x] LRU Cache

### Trees (9/10) - NEARLY COMPLETE
- [x] Same Tree
- [x] Subtree of Another Tree
- [x] Lowest Common Ancestor of a BST
- [x] Binary Tree Level Order Traversal
- [x] Validate BST
- [x] Kth Smallest in BST
- [x] Construct Binary Tree from Preorder and Inorder Traversal
- [x] Binary Tree Maximum Path Sum
- [x] Serialize and Deserialize Binary Tree
- [ ] Word Search II (moved to Tries)

### Tries (3/3) - COMPLETE
- [x] Implement Trie
- [x] Design Add and Search Words
- [x] Word Search II

### Heap / Priority Queue (5/6) - NEARLY COMPLETE
- [x] Kth Largest Element in a Stream
- [x] Last Stone Weight
- [x] K Closest Points to Origin
- [x] Task Scheduler
- [x] Find Median from Data Stream
- [ ] Design Twitter

### Backtracking (5/9) - PARTIAL
Implemented:
- [x] Subsets
- [x] Word Search
- [x] N-Queens
- [x] Combination Sum (was in original)
- [x] Permutations (was in original)

Remaining:
- [ ] Subsets II
- [ ] Combination Sum II
- [ ] Palindrome Partitioning
- [ ] Letter Combinations of a Phone Number

### Graphs (6/11) - PARTIAL
Implemented:
- [x] Clone Graph
- [x] Pacific Atlantic Water Flow
- [x] Rotting Oranges
- [x] Course Schedule II
- [x] Number of Islands (was in original)
- [x] Course Schedule (was in original)

Remaining:
- [ ] Surrounded Regions
- [ ] Walls and Gates
- [ ] Redundant Connection
- [ ] Word Ladder
- [ ] Graph Valid Tree
- [ ] Number of Connected Components
- [ ] Alien Dictionary

### 1D DP (6/9) - PARTIAL
Implemented:
- [x] House Robber II
- [x] Longest Palindromic Substring
- [x] Coin Change
- [x] Longest Increasing Subsequence
- [x] Word Break
- [x] Climbing Stairs (was in original)

Remaining:
- [ ] Palindromic Substrings
- [ ] Decode Ways
- [ ] Maximum Product Subarray
- [ ] Partition Equal Subset Sum

### 2D DP (1/11) - NEEDS WORK
Implemented:
- [x] Unique Paths (was in original)

Remaining:
- [ ] Longest Common Subsequence
- [ ] Best Time to Buy and Sell Stock with Cooldown
- [ ] Coin Change II
- [ ] Target Sum
- [ ] Interleaving String
- [ ] Longest Increasing Path in Matrix
- [ ] Distinct Subsequences
- [ ] Edit Distance
- [ ] Burst Balloons
- [ ] Regular Expression Matching

### Greedy (3/8) - PARTIAL
Implemented:
- [x] Maximum Subarray
- [x] Jump Game (was in original)
- [x] Jump Game II

Remaining:
- [ ] Gas Station
- [ ] Hand of Straights
- [ ] Merge Triplets to Form Target
- [ ] Partition Labels
- [ ] Valid Parenthesis String

### Intervals (2/5) - PARTIAL
Implemented:
- [x] Merge Intervals
- [x] Non-overlapping Intervals

Remaining:
- [ ] Insert Interval
- [ ] Meeting Rooms
- [ ] Meeting Rooms II

### Math & Geometry (2/8) - PARTIAL
Implemented:
- [x] Rotate Image
- [x] Spiral Matrix

Remaining:
- [ ] Set Matrix Zeroes
- [ ] Happy Number
- [ ] Plus One
- [ ] Pow(x, n)
- [ ] Multiply Strings
- [ ] Detect Squares

### Bit Manipulation (3/6) - PARTIAL
Implemented:
- [x] Single Number
- [x] Counting Bits
- [x] Missing Number

Remaining:
- [ ] Number of 1 Bits
- [ ] Reverse Bits
- [ ] Sum of Two Integers

## Advanced Graphs (0/5) - NOT STARTED
- [ ] Reconstruct Itinerary
- [ ] Min Cost to Connect All Points
- [ ] Network Delay Time
- [ ] Swim in Rising Water
- [ ] Cheapest Flights Within K Stops

---

## Problem Structure Reference

Each implemented problem includes:

1. **Core Fields**: ID, Number, Title, Difficulty, Category, Tags
2. **Related Chapters**: Links to tutorial chapters (e.g., [2, 7] for algorithms explained in chapters 2 and 7)
3. **Description & Constraints**: Full problem statement
4. **Examples & Test Cases**: Human-readable examples and machine test cases
5. **Time/Space Complexity**: Big-O analysis
6. **Starter Code**: Python template
7. **Hints (3 levels)**:
   - Level 1: Approach hint (general strategy)
   - Level 2: Algorithm hint (specific technique)
   - Level 3: Code hint (implementation details)
8. **Solution**: Complete code with explanation
9. **Walkthrough Steps**: Step-by-step code breakdown with line references

## Chapter Cross-References

Problems are linked to these tutorial chapters:
- Chapter 1: Introduction to Algorithms (Big-O, binary search basics)
- Chapter 2: Selection Sort, Arrays, Linked Lists
- Chapter 3: Recursion, Stack operations
- Chapter 4: Quicksort, Divide and Conquer, Backtracking
- Chapter 5: Hash Tables
- Chapter 6: BFS, Graph traversal
- Chapter 7: Dijkstra, Heaps
- Chapter 8: Greedy Algorithms
- Chapter 9: Dynamic Programming
- Chapter 10: K-Nearest Neighbors, Trees
- Chapter 11: Advanced Data Structures (Tries, Heaps, BSTs)
- Chapter 12: Design Patterns and Advanced Problems
