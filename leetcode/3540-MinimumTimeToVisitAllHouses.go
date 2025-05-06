package main

// 3540. Minimum Time to Visit All Houses
// You are given two integer arrays forward and backward, both of size n. 
// You are also given another integer array queries.

// There are n houses arranged in a circle. 
// The houses are connected via roads in a special arrangement:
//     1. For all 0 <= i <= n - 2, house i is connected to house i + 1 via a road with length forward[i] metres. 
//        Additionally, house n - 1 is connected back to house 0 via a road with length forward[n - 1] metres, completing the circle.
//     2. For all 1 <= i <= n - 1, house i is connected to house i - 1 via a road with length backward[i] metres. 
//        Additionally, house 0 is connected back to house n - 1 via a road with length backward[n - 1] metres, completing the circle.

// You can walk at a pace of one metre per second. 
// Starting from house 0, find the minimum time taken to visit each house in the order specified by queries.

// Return the minimum total time taken to visit the houses.

// Example 1:
// Input: forward = [1,4,4], backward = [4,1,2], queries = [1,2,0,2]
// Output: 12
// Explanation:
// The path followed is 0(0) → 1(1) →​​​​​​​ 2(5) → 1(7) →​​​​​​​ 0(8) → 2(12).
// Note: The notation used is node(total time), → represents forward road, and → represents backward road.

// Example 2:
// Input: forward = [1,1,1,1], backward = [2,2,2,2], queries = [1,2,3,0]
// Output: 4
// Explanation:
// The path travelled is 0 →​​​​​​​ 1 →​​​​​​​ 2 →​​​​​​​ 3 → 0. Each step is in the forward direction and requires 1 second.

// Constraints:
//     2 <= n <= 10^5
//     n == forward.length == backward.length
//     1 <= forward[i], backward[i] <= 10^5
//     1 <= queries.length <= 10^5
//     0 <= queries[i] < n
//     queries[i] != queries[i + 1]
//     queries[0] is not 0.

import "fmt"

// func minTotalTime(forward []int, backward []int, queries []int) int64 {
//     accumulate := func (nums []int) []int { // accumulate 返回一个切片，其中包含了原切片每个位置之前所有元素的累积和
//         if len(nums) == 0 { return []int{} }
//         res := make([]int, len(nums))
//         res[0] = nums[0] // 第一个元素就是其自身
//         for i := 1; i < len(nums); i++ {
//             res[i] = res[i-1] + nums[i]
//         }
//         return res
//     }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     prefixForward := accumulate(forward)
//     prefixBackward := accumulate(backward[1:])
//     res, last, f, b := 0, 0, 0, 0
//     for _, v := range queries {
//         if v > last {
//             f = prefixForward[v] - prefixForward[last]
//             b = prefixBackward[last] + backward[0] + prefixBackward[len(prefixBackward) - 1] - prefixBackward[v]
//         } else if v < last {
//             f = prefixForward[len(prefixForward) - 1] - prefixForward[last] + forward[len(forward) - 1] + prefixForward[v]
//             b = prefixBackward[last] - prefixBackward[v]
//         }
//         res += min(f, b)
//         last = v
//     }
//     return int64(res)
// }

// class Solution:
//     def minTotalTime(
//         self, forward: List[int], backward: List[int], queries: List[int]) -> int:
//         forpresum = list(accumulate(forward[:-1], initial=0))
//         bacpresum = list(accumulate(backward[1:], initial=0))
//         ans, last = 0, 0
//         for q in queries:
//             if q > last:
//                 f = forpresum[q] - forpresum[last]
//                 b = bacpresum[last] + backward[0] + bacpresum[-1] - bacpresum[q]
//             elif q < last:
//                 f = forpresum[-1] - forpresum[last] + forward[-1] + forpresum[q]
//                 b = bacpresum[last] - bacpresum[q]
//             ans += min(f, b)
//             last = q
//         return ans

func minTotalTime(forward []int, backward []int, queries []int) int64 {
    prefixForward, prefixBackward := make([]int, len(forward)), make([]int, len(backward))
    res, last := 0, 0
    for i := range forward {
        prefixForward[i], prefixBackward[i] = forward[i], backward[i]
        if i != 0 {
            prefixForward[i] += prefixForward[i - 1]
            prefixBackward[i] += prefixBackward[i - 1]
        }
    }
    get1 := func(mn, mx int) int {
        c1, c2 := prefixForward[mx - 1], prefixBackward[mn] + prefixBackward[len(prefixBackward) - 1] - prefixBackward[mx]
        if mn != 0 { c1 -= prefixForward[mn - 1] }
        if c1 < c2 { return c1 }
        return c2
    }
    get2 := func(mn, mx int) int {
        c1, c2 := prefixBackward[mx] - prefixBackward[mn], prefixForward[len(prefixForward) - 1]
        if mx != 0 { c2 -= prefixForward[mx - 1] }
        if mn != 0 { c2 += prefixForward[mn - 1] }
        if c1 < c2 { return c1 }
        return c2
    }
    helper := func(last, v int) int {
        if last == v { return 0 }
        if last < v  { return get1(last, v) }
        return get2(v, last)
    }
    for _, v := range queries {
        res += helper(last, v)
        last = v
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: forward = [1,4,4], backward = [4,1,2], queries = [1,2,0,2]
    // Output: 12
    // Explanation:
    // The path followed is 0(0) → 1(1) →​​​​​​​ 2(5) → 1(7) →​​​​​​​ 0(8) → 2(12).
    // Note: The notation used is node(total time), → represents forward road, and → represents backward road.
    fmt.Println(minTotalTime([]int{1,4,4}, []int{4,1,2}, []int{1,2,0,2})) // 12
    // Example 2:
    // Input: forward = [1,1,1,1], backward = [2,2,2,2], queries = [1,2,3,0]
    // Output: 4
    // Explanation:
    // The path travelled is 0 →​​​​​​​ 1 →​​​​​​​ 2 →​​​​​​​ 3 → 0. Each step is in the forward direction and requires 1 second.
    fmt.Println(minTotalTime([]int{1,1,1,1}, []int{2,2,2,2}, []int{1,2,3,0})) // 4

    fmt.Println(minTotalTime([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,0})) // 15
    fmt.Println(minTotalTime([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,0})) // 27
    fmt.Println(minTotalTime([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,0})) // 33
    fmt.Println(minTotalTime([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,0})) // 45
}