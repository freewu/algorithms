package main

// 646. Maximum Length of Pair Chain
// You are given an array of n pairs pairs where pairs[i] = [lefti, righti] and lefti < righti.
// A pair p2 = [c, d] follows a pair p1 = [a, b] if b < c. A chain of pairs can be formed in this fashion.
// Return the length longest chain which can be formed.
// You do not need to use up all the given intervals. You can select pairs in any order.

// Example 1:
// Input: pairs = [[1,2],[2,3],[3,4]]
// Output: 2
// Explanation: The longest chain is [1,2] -> [3,4].

// Example 2:
// Input: pairs = [[1,2],[7,8],[4,5]]
// Output: 3
// Explanation: The longest chain is [1,2] -> [4,5] -> [7,8].

// Constraints:
//     n == pairs.length
//     1 <= n <= 1000
//     -1000 <= lefti < righti <= 1000

import "fmt"
import "sort"

func findLongestChain(pairs [][]int) int {
    res := 1
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    prev := pairs[0][1]
    for i := 1; i < len(pairs); i++ {
        // 如果不能被包含则 多一链 res++
        if prev < pairs[i][0] {
            res++
            prev = pairs[i][1]
        }
    }
    return res
}

func main() {
    // Explanation: The longest chain is [1,2] -> [3,4].
    fmt.Println(findLongestChain([][]int{{1,2},{2,3},{3,4}})) // 2
    // Explanation: The longest chain is [1,2] -> [4,5] -> [7,8].
    fmt.Println(findLongestChain([][]int{{1,2},{7,8},{4,5}})) // 3
}