package main

// 3919. Minimum Cost to Move Between Indices
// You are given an integer array nums where nums is strictly increasing.

// For each index x, let closest(x) be the adjacent index such that abs(nums[x] - nums[y]) is minimized. 
// If both adjacent indices exist and give the same difference, choose the smaller index.

// From any index x, you can move in two ways:
//     1. To any index y with cost abs(nums[x] - nums[y]), or
//     2. To closest(x) with cost 1.

// You are also given a 2D integer array queries, where each queries[i] = [li, ri].

// For each query, calculate the minimum total cost to move from index li to index ri.

// Return an integer array ans, where ans[i] is the answer for the ith query.

// An array is said to be strictly increasing if each element is strictly greater than its previous one.

// The absolute difference between two values x and y is defined as abs(x - y).

// Example 1:
// Input: nums = [-5,-2,3], queries = [[0,2],[2,0],[1,2]]
// Output: [6,2,5]
// Explanation:​​​​​​​​​​​​​​​​​​​​
// The closest indices are [1, 0, 1] respectively.
// For [0, 2], the path 0 → 1 → 2 uses a closest move from index 0 to 1 with cost 1 and a move from index 1 to 2 with cost |-2 - 3| = 5, giving total 1 + 5 = 6.
// For [2, 0], the path 2 → 1 → 0 uses two closest moves from index 2 to 1 and from index 1 to 0, each with cost 1, giving total 2.
// For [1, 2], the direct move from index 1 to index 2 has cost |-2 - 3| = 5, which is optimal.
// Thus, ans = [6, 2, 5].

// Example 2:
// Input: nums = [0,2,3,9], queries = [[3,0],[1,2],[2,0]]
// Output: [4,1,3]
// Explanation:
// The closest indices are [1, 2, 1, 2] respectively.
// For [3, 0], the path 3 → 2 → 1 → 0 uses closest moves from index 3 to 2 and from 2 to 1, each with cost 1, and a move from 1 to 0 with cost |2 - 0| = 2, giving total 1 + 1 + 2 = 4.
// For [1, 2], the closest move from index 1 to 2 has cost 1.
// For [2, 0], the path 2 → 1 → 0 uses a closest move from index 2 to 1 with cost 1 and a move from 1 to 0 with cost |2 - 0| = 2, giving total 1 + 2 = 3.
// Thus, ans = [4, 1, 3].

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9
//     nums is strictly increasing
//     1 <= queries.length <= 10^5
//     queries[i] = [li, ri]​​​​​​​
//     0 <= li, ri < nums.length

import "fmt"

func minCost(nums []int, queries [][]int) []int {
    res, n := make([]int, len(queries)), len(nums)
    left, right := make([]int, n), make([]int, n) // left[i] 等于从 i 移动到 0 的代价和, right[i] 等于从 0 移动到 i 的代价和
    for i := 1; i < n; i++ {
        // 往左走 i -> i-1
        cost := 1
        if i < n-1 && nums[i]-nums[i-1] > nums[i+1]-nums[i] { // closest(i) = i+1
            cost = nums[i] - nums[i-1] // 只能用方式一往左走
        }
        left[i] = left[i-1] + cost  
        // 往右走 i-1 -> i
        cost = 1
        if i > 1 && nums[i-1]-nums[i-2] <= nums[i]-nums[i-1] { // closest(i-1) = i-2
            cost = nums[i] - nums[i-1] // 只能用方式一往右走
        }
        right[i] = right[i-1] + cost
    }
    for i, q := range queries {
        l, r := q[0], q[1]
        if l < r {
            // cost(0 -> r) - cost(0 -> l) = cost(l -> r)
            res[i] = right[r] - right[l]
        } else {
            // cost(l -> 0) - cost(r -> 0) = cost(l -> r)
            res[i] = left[l] - left[r]
        }
    }
    return res
}

func minCost1(nums []int, queries [][]int) []int {
    res, n := make([]int, len(queries)), len(nums)
    left, right := make([]int, n), make([]int, n)
    for i := 1; i < n; i++ {
        var cost int
        if i == 1 {
            cost = 1
        } else {
            dl := nums[i-1] - nums[i-2]
            dr := nums[i] - nums[i-1]
            if dr < dl {
                cost = 1
            } else {
                cost = dr
            }
        }
        left[i] = left[i-1] + cost
    }
    for i := n-2; i >= 0; i-- {
        cost := 0
        if i == n - 2 {
            cost = 1
        } else {
            dr := nums[i+2] - nums[i+1]
            dl := nums[i+1] - nums[i]
            if dl <= dr {
                cost = 1
            } else {
                cost = dl
            }
        }
        right[i] = right[i+1] + cost
    }
    for i, q := range queries {
        u, v := q[0], q[1]
        if u < v {
            res[i] = left[v] - left[u]        
        } else {
            res[i] = right[v] - right[u]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-5,-2,3], queries = [[0,2],[2,0],[1,2]]
    // Output: [6,2,5]
    // Explanation:​​​​​​​​​​​​​​​​​​​​
    // The closest indices are [1, 0, 1] respectively.
    // For [0, 2], the path 0 → 1 → 2 uses a closest move from index 0 to 1 with cost 1 and a move from index 1 to 2 with cost |-2 - 3| = 5, giving total 1 + 5 = 6.
    // For [2, 0], the path 2 → 1 → 0 uses two closest moves from index 2 to 1 and from index 1 to 0, each with cost 1, giving total 2.
    // For [1, 2], the direct move from index 1 to index 2 has cost |-2 - 3| = 5, which is optimal.
    // Thus, ans = [6, 2, 5].
    fmt.Println(minCost([]int{-5,-2,3}, [][]int{{0,2},{2,0},{1,2}})) // [6,2,5]
    // Example 2:
    // Input: nums = [0,2,3,9], queries = [[3,0],[1,2],[2,0]]
    // Output: [4,1,3]
    // Explanation:
    // The closest indices are [1, 2, 1, 2] respectively.
    // For [3, 0], the path 3 → 2 → 1 → 0 uses closest moves from index 3 to 2 and from 2 to 1, each with cost 1, and a move from 1 to 0 with cost |2 - 0| = 2, giving total 1 + 1 + 2 = 4.
    // For [1, 2], the closest move from index 1 to 2 has cost 1.
    // For [2, 0], the path 2 → 1 → 0 uses a closest move from index 2 to 1 with cost 1 and a move from 1 to 0 with cost |2 - 0| = 2, giving total 1 + 2 = 3.
    // Thus, ans = [4, 1, 3].
    fmt.Println(minCost([]int{0,2,3,9}, [][]int{{3,0},{1,2},{2,0}})) // [4,1,3]

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, [][]int{{3,0},{1,2},{2,0}})) // [3 1 2]
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, [][]int{{3,0},{1,2},{2,0}})) // [3 -1 2]

    fmt.Println(minCost1([]int{-5,-2,3}, [][]int{{0,2},{2,0},{1,2}})) // [6,2,5]
    fmt.Println(minCost1([]int{0,2,3,9}, [][]int{{3,0},{1,2},{2,0}})) // [4,1,3]
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, [][]int{{3,0},{1,2},{2,0}})) // [3 1 2]
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, [][]int{{3,0},{1,2},{2,0}})) // [3 -1 2]
}