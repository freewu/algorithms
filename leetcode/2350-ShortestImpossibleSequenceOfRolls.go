package main

// 2350. Shortest Impossible Sequence of Rolls
// You are given an integer array rolls of length n and an integer k. 
// You roll a k sided dice numbered from 1 to k, n times, where the result of the ith roll is rolls[i].

// Return the length of the shortest sequence of rolls so that there's no such subsequence in rolls.

// A sequence of rolls of length len is the result of rolling a k sided dice len times.

// Example 1:
// Input: rolls = [4,2,1,2,3,3,2,4,1], k = 4
// Output: 3
// Explanation: Every sequence of rolls of length 1, [1], [2], [3], [4], can be taken from rolls.
// Every sequence of rolls of length 2, [1, 1], [1, 2], ..., [4, 4], can be taken from rolls.
// The sequence [1, 4, 2] cannot be taken from rolls, so we return 3.
// Note that there are other sequences that cannot be taken from rolls.

// Example 2:
// Input: rolls = [1,1,2,2], k = 2
// Output: 2
// Explanation: Every sequence of rolls of length 1, [1], [2], can be taken from rolls.
// The sequence [2, 1] cannot be taken from rolls, so we return 2.
// Note that there are other sequences that cannot be taken from rolls but [2, 1] is the shortest.

// Example 3:
// Input: rolls = [1,1,3,2,2,2,3,3], k = 4
// Output: 1
// Explanation: The sequence [4] cannot be taken from rolls, so we return 1.
// Note that there are other sequences that cannot be taken from rolls but [4] is the shortest.

// Constraints:
//     n == rolls.length
//     1 <= n <= 10^5
//     1 <= rolls[i] <= k <= 10^5

import "fmt"

func shortestSequence(rolls []int, k int) int {
    res, curr, seen := 1, 0, make(map[int]bool)
    for _, v := range rolls {
        if !seen[v] {
            curr++
            seen[v] = true
        }
        // when we've collected all k nums, ++
        if curr == k {
            res++
            curr = 0
            seen = make(map[int]bool)
        }
    }
    return res
}

func shortestSequence1(rolls []int, k int) int {
    res, left, mark := 1, k, make([]int, k + 1)
    for _, v := range rolls {
        if mark[v] < res {
            mark[v] = res
            if left--; left == 0 {
                left = k
                res++
            }
        }
    }
    return res
}

func shortestSequence2(rolls []int, k int) int {
    count, res, exists := 0, 0, make([]bool, k)
    for i := 0; i < len(rolls); i++ {
        if !exists[rolls[i]-1] {
            count++ // count distinct rolls in sequence
            exists[rolls[i]-1] = true // mark roll as encountered
        }
        if count == k { // we have encountered all possible rolls in current sequence
            res++
            count = 0 // reset sequence counter
            for j := 0; j < k; j++ {
                exists[j] = false
            }
        }
    }
    return res + 1
}

func main() {
    // Example 1:
    // Input: rolls = [4,2,1,2,3,3,2,4,1], k = 4
    // Output: 3
    // Explanation: Every sequence of rolls of length 1, [1], [2], [3], [4], can be taken from rolls.
    // Every sequence of rolls of length 2, [1, 1], [1, 2], ..., [4, 4], can be taken from rolls.
    // The sequence [1, 4, 2] cannot be taken from rolls, so we return 3.
    // Note that there are other sequences that cannot be taken from rolls.
    fmt.Println(shortestSequence([]int{4,2,1,2,3,3,2,4,1}, 4)) // 3
    // Example 2:
    // Input: rolls = [1,1,2,2], k = 2
    // Output: 2
    // Explanation: Every sequence of rolls of length 1, [1], [2], can be taken from rolls.
    // The sequence [2, 1] cannot be taken from rolls, so we return 2.
    // Note that there are other sequences that cannot be taken from rolls but [2, 1] is the shortest.
    fmt.Println(shortestSequence([]int{1,1,2,2}, 2)) // 2
    // Example 3:
    // Input: rolls = [1,1,3,2,2,2,3,3], k = 4
    // Output: 1
    // Explanation: The sequence [4] cannot be taken from rolls, so we return 1.
    // Note that there are other sequences that cannot be taken from rolls but [4] is the shortest.
    fmt.Println(shortestSequence([]int{1,1,3,2,2,2,3,3}, 4)) // 1

    fmt.Println(shortestSequence1([]int{4,2,1,2,3,3,2,4,1}, 4)) // 3
    fmt.Println(shortestSequence1([]int{1,1,2,2}, 2)) // 2
    fmt.Println(shortestSequence1([]int{1,1,3,2,2,2,3,3}, 4)) // 1

    fmt.Println(shortestSequence2([]int{4,2,1,2,3,3,2,4,1}, 4)) // 3
    fmt.Println(shortestSequence2([]int{1,1,2,2}, 2)) // 2
    fmt.Println(shortestSequence2([]int{1,1,3,2,2,2,3,3}, 4)) // 1
}