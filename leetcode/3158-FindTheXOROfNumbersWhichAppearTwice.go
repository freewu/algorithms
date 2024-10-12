package main

// 3158. Find the XOR of Numbers Which Appear Twice
// You are given an array nums, where each number in the array appears either once or twice.

// Return the bitwise XOR of all the numbers that appear twice in the array, or 0 if no number appears twice.

// Example 1:
// Input: nums = [1,2,1,3]
// Output: 1
// Explanation:
// The only number that appears twice in nums is 1.

// Example 2:
// Input: nums = [1,2,3]
// Output: 0
// Explanation:
// No number appears twice in nums.

// Example 3:
// Input: nums = [1,2,2,1]
// Output: 3
// Explanation:
// Numbers 1 and 2 appeared twice. 1 XOR 2 == 3.

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 50
//     Each number in nums appears either once or twice.

import "fmt"

func duplicateNumbersXOR(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    arr := []int{}
    for k, v := range mp {
        if v == 2 {
            arr = append(arr, k)
        }
    }
    res := 0
    for _, v := range arr {
        res ^= v
    }
    return res
}

func duplicateNumbersXOR1(nums []int) int {
    res, cnt1,cnt2 := 0, 0, 0
    for _, v := range nums {
        if cnt1 & ( 1 << v) == 0 {
            cnt1 |= (1 << v)
        } else {
            cnt2 |= (1 << v)
        }
    }
    for i := 0; cnt2 > 0; i++ {
        if cnt2 & 1 != 0 {
            res ^= i
        }
        cnt2 >>= 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,3]
    // Output: 1
    // Explanation:
    // The only number that appears twice in nums is 1.
    fmt.Println(duplicateNumbersXOR([]int{1,2,1,3})) // 1
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation:
    // No number appears twice in nums.
    fmt.Println(duplicateNumbersXOR([]int{1,2,3})) // 0
    // Example 3:
    // Input: nums = [1,2,2,1]
    // Output: 3
    // Explanation:
    // Numbers 1 and 2 appeared twice. 1 XOR 2 == 3.
    fmt.Println(duplicateNumbersXOR([]int{1,2,2,1})) // 3

    fmt.Println(duplicateNumbersXOR1([]int{1,2,1,3})) // 1
    fmt.Println(duplicateNumbersXOR1([]int{1,2,3})) // 0
    fmt.Println(duplicateNumbersXOR1([]int{1,2,2,1})) // 3
}