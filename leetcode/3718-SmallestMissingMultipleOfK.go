package main

// 3718. Smallest Missing Multiple of K
// Given an integer array nums and an integer k, return the smallest positive multiple of k that is missing from nums.

// A multiple of k is any positive integer divisible by k.

// Example 1:
// Input: nums = [8,2,3,4,6], k = 2
// Output: 10
// Explanation:
// The multiples of k = 2 are 2, 4, 6, 8, 10, 12... and the smallest multiple missing from nums is 10.

// Example 2:
// Input: nums = [1,4,7,10,15], k = 5
// Output: 5
// Explanation:
// The multiples of k = 5 are 5, 10, 15, 20... and the smallest multiple missing from nums is 5.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     1 <= k <= 100

import "fmt"

func missingMultiple(nums []int, k int) int {
    mp := make(map[int]bool)
    for _, v := range nums {
        mp[v] = true
    }
    for i := 1; i <= 100; i++ {
        if !mp[i*k] {
            return i*k
        }
    }
    return 101
}

func main() {
    // Example 1:
    // Input: nums = [8,2,3,4,6], k = 2
    // Output: 10
    // Explanation:
    // The multiples of k = 2 are 2, 4, 6, 8, 10, 12... and the smallest multiple missing from nums is 10.
    fmt.Println(missingMultiple([]int{8,2,3,4,6}, 2)) // 10
    // Example 2:
    // Input: nums = [1,4,7,10,15], k = 5
    // Output: 5
    // Explanation:
    // The multiples of k = 5 are 5, 10, 15, 20... and the smallest multiple missing from nums is 5.
    fmt.Println(missingMultiple([]int{1,4,7,10,15}, 5)) // 5

    fmt.Println(missingMultiple([]int{1,2,3,4,5,6,7,8,9}, 2)) // 10
    fmt.Println(missingMultiple([]int{9,8,7,6,5,4,3,2,1}, 2)) // 10
    fmt.Println(missingMultiple([]int{9,8,7,6,5,4,3,2,1}, 2)) // 10
    fmt.Println(missingMultiple([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100}, 1)) // 101
}