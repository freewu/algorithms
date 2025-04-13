package main

// 3514. Number of Unique XOR Triplets II
// You are given an integer array nums.

// A XOR triplet is defined as the XOR of three elements nums[i] XOR nums[j] XOR nums[k] where i <= j <= k.

// Return the number of unique XOR triplet values from all possible triplets (i, j, k).

// Example 1:
// Input: nums = [1,3]
// Output: 2
// Explanation:
// The possible XOR triplet values are:
// (0, 0, 0) → 1 XOR 1 XOR 1 = 1
// (0, 0, 1) → 1 XOR 1 XOR 3 = 3
// (0, 1, 1) → 1 XOR 3 XOR 3 = 1
// (1, 1, 1) → 3 XOR 3 XOR 3 = 3
// The unique XOR values are {1, 3}. Thus, the output is 2.

// Example 2:
// Input: nums = [6,7,8,9]
// Output: 4
// Explanation:
// The possible XOR triplet values are {6, 7, 8, 9}. Thus, the output is 4.

// Constraints:
//     1 <= nums.length <= 1500
//     1 <= nums[i] <= 1500

import "fmt"
import "slices"
import "math/bits"

func uniqueXorTriplets(nums []int) int {
    res, arr, visit := 0, make([]bool, 1 << 11), make([]bool, 1 << 11)
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            visit[nums[i]^nums[j]] = true
        }
    }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(visit); j++ {
            if !visit[j] { continue }
            if arr[nums[i]^j] { continue }
            arr[nums[i]^j] = true
            res++
        }
    }
    return res
}

func uniqueXorTriplets1(nums []int) int {
    res, n := 0, 1 << bits.Len(uint(slices.Max(nums)))
    has := make([]bool, n)
    for i, x := range nums {
        for _, y := range nums[i:] {
            has[x^y] = true
        }
    }
    has3 := make([]bool, n)
    for xy, b := range has {
        if !b { continue }
        for _, z := range nums {
            has3[xy^z] = true
        }
    }
    for _, b := range has3 {
        if b {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2]
    // Output: 2
    // Explanation:
    // The possible XOR triplet values are:
    // (0, 0, 0) → 1 XOR 1 XOR 1 = 1
    // (0, 0, 1) → 1 XOR 1 XOR 2 = 2
    // (0, 1, 1) → 1 XOR 2 XOR 2 = 1
    // (1, 1, 1) → 2 XOR 2 XOR 2 = 2
    // The unique XOR values are {1, 2}, so the output is 2.
    fmt.Println(uniqueXorTriplets([]int{1,2})) // 2
    // Example 2:
    // Input: nums = [6,7,8,9]
    // Output: 4
    // Explanation:
    // The possible XOR triplet values are {6, 7, 8, 9}. Thus, the output is 4.
    fmt.Println(uniqueXorTriplets([]int{6,7,8,9})) // 4

    fmt.Println(uniqueXorTriplets([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(uniqueXorTriplets([]int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(uniqueXorTriplets1([]int{1,2})) // 2
    fmt.Println(uniqueXorTriplets1([]int{6,7,8,9})) // 4
    fmt.Println(uniqueXorTriplets1([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(uniqueXorTriplets1([]int{9,8,7,6,5,4,3,2,1})) // 16
}