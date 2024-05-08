package main

// 506. Relative Ranks
// You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. 
// All the scores are guaranteed to be unique.

// The athletes are placed based on their scores, where the 1st place athlete has the highest score, 
// the 2nd place athlete has the 2nd highest score, and so on. 
// The placement of each athlete determines their rank:
//     The 1st place athlete's rank is "Gold Medal".
//     The 2nd place athlete's rank is "Silver Medal".
//     The 3rd place athlete's rank is "Bronze Medal".

// For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").
// Return an array answer of size n where answer[i] is the rank of the ith athlete.

// Example 1:
// Input: score = [5,4,3,2,1]
// Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
// Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].

// Example 2:
// Input: score = [10,3,8,9,4]
// Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
// Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].

// Constraints:
//     n == score.length
//     1 <= n <= 10^4
//     0 <= score[i] <= 10^6
//     All the values in score are unique.

import "fmt"
import "sort"
import "strconv"

func findRelativeRanks(score []int) []string {
    arr, m := make([]int, len(score)), make(map[int]int, len(score))
    copy(arr, score)
    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
    for i,v := range arr { // v 和 i做个映射
        m[v] = i + 1
    }
    res := make([]string,len(score))
    for i,v := range score {
        switch m[v] {
        case 1:
            res[i] = "Gold Medal"
        case 2:
            res[i] = "Silver Medal"
        case 3:
            res[i] = "Bronze Medal"
        default:
            res[i] = strconv.Itoa(m[v])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: score = [5,4,3,2,1]
    // Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
    // Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].
    fmt.Println(findRelativeRanks([]int{5,4,3,2,1})) // ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
    // Example 2:
    // Input: score = [10,3,8,9,4]
    // Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
    // Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].
    fmt.Println(findRelativeRanks([]int{10,3,8,9,4})) // ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
}

