package main

// 1524. Number of Sub-arrays With Odd Sum
// Given an array of integers arr, return the number of subarrays with an odd sum.

// Since the answer can be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: arr = [1,3,5]
// Output: 4
// Explanation: All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
// All sub-arrays sum are [1,4,9,3,8,5].
// Odd sums are [1,9,3,5] so the answer is 4.

// Example 2:
// Input: arr = [2,4,6]
// Output: 0
// Explanation: All subarrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
// All sub-arrays sum are [2,6,12,4,10,6].
// All sub-arrays have even sum and the answer is 0.

// Example 3:
// Input: arr = [1,2,3,4,5,6,7]
// Output: 16

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= arr[i] <= 100

import "fmt"

func numOfSubarrays(arr []int) int {
    index, tmp := 0, []int{ 1, 0 }
    for _, v := range arr {
        index = ((index + v) % 2 + 2) % 2
        tmp[index]++
    }
    return (tmp[0] * tmp[1]) % 1_000_000_007
}

func numOfSubarrays1(arr []int) int {
    sum, odd, even := 0, 0, 1 // empty list is even number
    for _, v := range arr {
        sum += v
        if sum & 1 == 1 {
            odd++
        } else {
            even++
        }
    }
    return (odd * even) % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: arr = [1,3,5]
    // Output: 4
    // Explanation: All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
    // All sub-arrays sum are [1,4,9,3,8,5].
    // Odd sums are [1,9,3,5] so the answer is 4.
    fmt.Println(numOfSubarrays([]int{1,3,5})) // 4
    // Example 2:
    // Input: arr = [2,4,6]
    // Output: 0
    // Explanation: All subarrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
    // All sub-arrays sum are [2,6,12,4,10,6].
    // All sub-arrays have even sum and the answer is 0.
    fmt.Println(numOfSubarrays([]int{2,4,6})) // 0
    // Example 3:
    // Input: arr = [1,2,3,4,5,6,7]
    // Output: 16
    fmt.Println(numOfSubarrays([]int{1,2,3,4,5,6,7})) // 16

    fmt.Println(numOfSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(numOfSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 25

    fmt.Println(numOfSubarrays1([]int{1,3,5})) // 4
    fmt.Println(numOfSubarrays1([]int{2,4,6})) // 0
    fmt.Println(numOfSubarrays1([]int{1,2,3,4,5,6,7})) // 16
    fmt.Println(numOfSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(numOfSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 25
}