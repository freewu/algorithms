package main

// 3495. Minimum Operations to Make Array Elements Zero
// You are given a 2D array queries, where queries[i] is of the form [l, r]. 
// Each queries[i] defines an array of integers nums consisting of elements ranging from l to r, both inclusive.

// In one operation, you can:
//     1. Select two integers a and b from the array.
//     2. Replace them with floor(a / 4) and floor(b / 4).

// Your task is to determine the minimum number of operations required to reduce all elements of the array to zero for each query. 
// Return the sum of the results for all queries.

// Example 1:
// Input: queries = [[1,2],[2,4]]
// Output: 3
// Explanation:
// For queries[0]:
// The initial array is nums = [1, 2].
// In the first operation, select nums[0] and nums[1]. The array becomes [0, 0].
// The minimum number of operations required is 1.
// For queries[1]:
// The initial array is nums = [2, 3, 4].
// In the first operation, select nums[0] and nums[2]. The array becomes [0, 3, 1].
// In the second operation, select nums[1] and nums[2]. The array becomes [0, 0, 0].
// The minimum number of operations required is 2.
// The output is 1 + 2 = 3.

// Example 2:
// Input: queries = [[2,6]]
// Output: 4
// Explanation:
// For queries[0]:
// The initial array is nums = [2, 3, 4, 5, 6].
// In the first operation, select nums[0] and nums[3]. The array becomes [0, 3, 4, 1, 6].
// In the second operation, select nums[2] and nums[4]. The array becomes [0, 3, 1, 1, 1].
// In the third operation, select nums[1] and nums[2]. The array becomes [0, 0, 0, 1, 1].
// In the fourth operation, select nums[3] and nums[4]. The array becomes [0, 0, 0, 0, 0].
// The minimum number of operations required is 4.
// The output is 4.

// Constraints:
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     queries[i] == [l, r]
//     1 <= l < r <= 10^9

import "fmt"
import "math/bits"

func minOperations(queries [][]int) int64 {
    res := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, q := range queries {
        l, r, ops := q[0], q[1], 0
        for i, step := 1, 1; i <= 15; i++ {
            low, high := max(step, l), min(4 * step, r + 1)
            if low < high {
                ops += i * (high - low)
            }
            step *= 4
        }
        res += (ops + 1) / 2
    }
    return int64(res)
}

func minOperations1(queries [][]int) int64 {
    res := 0
    helper := func(n int) int { // 返回 [1,n] 的单个元素的操作次数之和
        m := bits.Len(uint(n))
        k := (m - 1) / 2 * 2
        res := k / 2 << k - (1 << k - 1) / 3
        return res + (m + 1) / 2 * (n + 1 - 1 << k)
    }
    for _, q := range queries {
        l, r := q[0], q[1]
        sum, mx := helper(r) - helper(l - 1), (bits.Len(uint(r)) + 1) / 2
        res += max((sum + 1)/2, mx)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: queries = [[1,2],[2,4]]
    // Output: 3
    // Explanation:
    // For queries[0]:
    // The initial array is nums = [1, 2].
    // In the first operation, select nums[0] and nums[1]. The array becomes [0, 0].
    // The minimum number of operations required is 1.
    // For queries[1]:
    // The initial array is nums = [2, 3, 4].
    // In the first operation, select nums[0] and nums[2]. The array becomes [0, 3, 1].
    // In the second operation, select nums[1] and nums[2]. The array becomes [0, 0, 0].
    // The minimum number of operations required is 2.
    // The output is 1 + 2 = 3.
    fmt.Println(minOperations([][]int{{1, 2}, {2, 4}})) // 3
    // Example 2:
    // Input: queries = [[2,6]]
    // Output: 4
    // Explanation:
    // For queries[0]:
    // The initial array is nums = [2, 3, 4, 5, 6].
    // In the first operation, select nums[0] and nums[3]. The array becomes [0, 3, 4, 1, 6].
    // In the second operation, select nums[2] and nums[4]. The array becomes [0, 3, 1, 1, 1].
    // In the third operation, select nums[1] and nums[2]. The array becomes [0, 0, 0, 1, 1].
    // In the fourth operation, select nums[3] and nums[4]. The array becomes [0, 0, 0, 0, 0].
    // The minimum number of operations required is 4.
    // The output is 4.
    fmt.Println(minOperations([][]int{{2, 6}})) // 4

    fmt.Println(minOperations1([][]int{{1, 2}, {2, 4}})) // 3
    fmt.Println(minOperations1([][]int{{2, 6}})) // 4
}