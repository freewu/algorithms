package main

// 3152. Special Array II
// An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

// You are given an array of integer nums and a 2D integer matrix queries, 
// where for queries[i] = [fromi, toi] your task is to check that subarray nums[fromi..toi] is special or not.

// Return an array of booleans answer such that answer[i] is true if nums[fromi..toi] is special.

// Example 1:
// Input: nums = [3,4,1,2,6], queries = [[0,4]]
// Output: [false]
// Explanation:
// The subarray is [3,4,1,2,6]. 2 and 6 are both even.

// Example 2:
// Input: nums = [4,3,1,6], queries = [[0,2],[2,3]]
// Output: [false,true]
// Explanation:
// The subarray is [4,3,1]. 3 and 1 are both odd. So the answer to this query is false.
// The subarray is [1,6]. There is only one pair: (1,6) and it contains numbers with different parity. So the answer to this query is true.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     0 <= queries[i][0] <= queries[i][1] <= nums.length - 1

import "fmt"

// 超出时间限制 535 / 536 个通过的测试用例
func isArraySpecial(nums []int, queries [][]int) []bool {
    res := []bool{}
    check := func(nums []int) bool {
        for i := 1; i < len(nums); i++ {
            if nums[i - 1] % 2 == nums[i] % 2 {
                return false
            }
        }
        return true
    }
    for _, v := range queries {
        res = append(res, check(nums[v[0]:v[1] + 1]))
    }
    return res
}

// Prefix Sum
func isArraySpecial1(nums []int, queries [][]int) []bool {
    n, n1, count := len(nums), len(queries), 0
    res, violations := []bool{}, make([]int, n)
    for i := 1; i < n; i++ {
        if nums[i] % 2 == nums[i-1] % 2 {
            count++
        }
        violations[i] = count
    }
    for i := 0; i < n1; i++ {
        start := queries[i][0]
        end := queries[i][1]
        res = append(res, violations[end] - violations[start] == 0)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,4,1,2,6], queries = [[0,4]]
    // Output: [false]
    // Explanation:
    // The subarray is [3,4,1,2,6]. 2 and 6 are both even.
    fmt.Println(isArraySpecial([]int{3,4,1,2,6},[][]int{{0,4}})) // [false]
    // Example 2:
    // Input: nums = [4,3,1,6], queries = [[0,2],[2,3]]
    // Output: [false,true]
    // Explanation:
    // The subarray is [4,3,1]. 3 and 1 are both odd. So the answer to this query is false.
    // The subarray is [1,6]. There is only one pair: (1,6) and it contains numbers with different parity. So the answer to this query is true.
    fmt.Println(isArraySpecial([]int{4,3,1,6},[][]int{{0,2},{2,3}})) // [false,true]

    fmt.Println(isArraySpecial1([]int{3,4,1,2,6},[][]int{{0,4}})) // [false]
    fmt.Println(isArraySpecial1([]int{4,3,1,6},[][]int{{0,2},{2,3}})) // [false,true]
}