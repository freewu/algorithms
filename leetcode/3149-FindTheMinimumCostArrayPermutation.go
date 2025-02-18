package main

// 3149. Find the Minimum Cost Array Permutation
// You are given an array nums which is a permutation of [0, 1, 2, ..., n - 1]. 
// The score of any permutation of [0, 1, 2, ..., n - 1] named perm is defined as:
//     score(perm) = |perm[0] - nums[perm[1]]| + |perm[1] - nums[perm[2]]| + ... + |perm[n - 1] - nums[perm[0]]|

// Return the permutation perm which has the minimum possible score. 
// If multiple permutations exist with this score, return the one that is lexicographically smallest among them.

// Example 1:
// Input: nums = [1,0,2]
// Output: [0,1,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/04/example0gif.gif" />
// The lexicographically smallest permutation with minimum cost is [0,1,2]. 
// The cost of this permutation is |0 - 0| + |1 - 2| + |2 - 1| = 2.

// Example 2:
// Input: nums = [0,2,1]
// Output: [0,2,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/04/example1gif.gif" />
// The lexicographically smallest permutation with minimum cost is [0,2,1]. 
// The cost of this permutation is |0 - 1| + |2 - 2| + |1 - 0| = 2.

// Constraints:
//     2 <= n == nums.length <= 14
//     nums is a permutation of [0, 1, 2, ..., n - 1].

import "fmt"
import "math/bits"

func findPermutation(numbers []int) []int {
    n := len(numbers)
    nums, memo, bestPick := numbers, make([][]int, n), make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, 1 << n)
    }
    for i := range bestPick {
        bestPick[i] = make([]int, 1 << n)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var getScore func(last int, mask uint16) int
    getScore = func(last int, mask uint16) int {
        if bits.OnesCount16(mask) == len(nums) { return abs(last - nums[0]) }
        if memo[last][mask] > 0 { return memo[last][mask] }
        minScore := 1 << 31
        for i, n := range nums {
            if (mask & (1 << i)) != 0 { continue }
            nextMinScore := abs(n - last) + getScore(i, mask | (1 << i))
            if nextMinScore < minScore {
                minScore = nextMinScore
                bestPick[last][mask] = i
            }
        }
        memo[last][mask] = minScore
        return memo[last][mask] 
    }
    getScore(0, uint16(1))
    res := make([]int, n)
    for i, last, mask := 0, 0, 1; i < len(bestPick); i++ {
        res[i] = last
        last = bestPick[last][mask]
        mask |= 1 << last
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,2]
    // Output: [0,1,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/04/example0gif.gif" />
    // The lexicographically smallest permutation with minimum cost is [0,1,2]. 
    // The cost of this permutation is |0 - 0| + |1 - 2| + |2 - 1| = 2.
    fmt.Println(findPermutation([]int{1,0,2})) // [0,1,2]
    // Example 2:
    // Input: nums = [0,2,1]
    // Output: [0,2,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/04/example1gif.gif" />
    // The lexicographically smallest permutation with minimum cost is [0,2,1]. 
    // The cost of this permutation is |0 - 1| + |2 - 2| + |1 - 0| = 2.
    fmt.Println(findPermutation([]int{0,2,1})) // [0,2,1]

    fmt.Println(findPermutation([]int{1,2,3,4,5,6,7,8,9})) // [0 8 7 6 5 4 3 2 1]
    fmt.Println(findPermutation([]int{9,8,7,6,5,4,3,2,1})) // [0 4 5 3 6 2 7 1 8]
}