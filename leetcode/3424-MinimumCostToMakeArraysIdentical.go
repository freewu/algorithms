package main

// 3424. Minimum Cost to Make Arrays Identical
// You are given two integer arrays arr and brr of length n, and an integer k. 
// You can perform the following operations on arr any number of times:
//     1. Split arr into any number of contiguous subarrays and rearrange these subarrays in any order. 
//        This operation has a fixed cost of k.
//     2. Choose any element in arr and add or subtract a positive integer x to it. 
//        The cost of this operation is x.

// Return the minimum total cost to make arr equal to brr.

// Example 1:
// Input: arr = [-7,9,5], brr = [7,-2,-5], k = 2
// Output: 13
// Explanation:
// Split arr into two contiguous subarrays: [-7] and [9, 5] and rearrange them as [9, 5, -7], with a cost of 2.
// Subtract 2 from element arr[0]. The array becomes [7, 5, -7]. The cost of this operation is 2.
// Subtract 7 from element arr[1]. The array becomes [7, -2, -7]. The cost of this operation is 7.
// Add 2 to element arr[2]. The array becomes [7, -2, -5]. The cost of this operation is 2.
// The total cost to make the arrays equal is 2 + 2 + 7 + 2 = 13.

// Example 2:
// Input: arr = [2,1], brr = [2,1], k = 0
// Output: 0
// Explanation:
// Since the arrays are already equal, no operations are needed, and the total cost is 0.

// Constraints:
//     1 <= arr.length == brr.length <= 10^5
//     0 <= k <= 2 * 10^10
//     -10^5 <= arr[i] <= 10^5
//     -10^5 <= brr[i] <= 10^5

import "fmt"
import "sort"

func minCost(arr []int, brr []int, k int64) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    calc := func(a, b []int) int64 {
        res := int64(0)
        for i := range a {
            res += int64(abs(a[i] - b[i]))
        }
        return res
    }
    c1 := calc(arr, brr)
    sort.Ints(arr)
    sort.Ints(brr)
    c2 := calc(arr, brr) + k
    return min(c1, c2)
}

// 贪心
func minCost1(arr, brr []int, k int64) int64 {
    res1 := int64(0)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    for i, v := range arr {
        res1 += int64(abs(v - brr[i]))
    }
    if res1 <= k {
        return res1
    }
    sort.Ints(arr)
    sort.Ints(brr)
    res2 := k
    for i, v := range arr {
        res2 += int64(abs(v - brr[i]))
    }
    return min(res1, res2)
}

func main() {
    // Example 1:
    // Input: arr = [-7,9,5], brr = [7,-2,-5], k = 2
    // Output: 13
    // Explanation:
    // Split arr into two contiguous subarrays: [-7] and [9, 5] and rearrange them as [9, 5, -7], with a cost of 2.
    // Subtract 2 from element arr[0]. The array becomes [7, 5, -7]. The cost of this operation is 2.
    // Subtract 7 from element arr[1]. The array becomes [7, -2, -7]. The cost of this operation is 7.
    // Add 2 to element arr[2]. The array becomes [7, -2, -5]. The cost of this operation is 2.
    // The total cost to make the arrays equal is 2 + 2 + 7 + 2 = 13.
    fmt.Println(minCost([]int{-7,9,5}, []int{7,-2,-5}, 2)) // 13
    // Example 2:
    // Input: arr = [2,1], brr = [2,1], k = 0
    // Output: 0
    // Explanation:
    // Since the arrays are already equal, no operations are needed, and the total cost is 0.
    fmt.Println(minCost([]int{2,1}, []int{2,1}, 0)) // 0

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 2
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 2

    fmt.Println(minCost1([]int{-7,9,5}, []int{7,-2,-5}, 2)) // 13
    fmt.Println(minCost1([]int{2,1}, []int{2,1}, 0)) // 0
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 2
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 2
}