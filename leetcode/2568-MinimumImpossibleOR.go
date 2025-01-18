package main

// 2568. Minimum Impossible OR
// You are given a 0-indexed integer array nums.

// We say that an integer x is expressible from nums 
// if there exist some integers 0 <= index1 < index2 < ... < indexk < nums.length 
// for which nums[index1] | nums[index2] | ... | nums[indexk] = x. 
// In other words, an integer is expressible if it can be written as the bitwise OR of some subsequence of nums.

// Return the minimum positive non-zero integer that is not expressible from nums.

// Example 1:
// Input: nums = [2,1]
// Output: 4
// Explanation: 1 and 2 are already present in the array. We know that 3 is expressible, since nums[0] | nums[1] = 2 | 1 = 3. Since 4 is not expressible, we return 4.

// Example 2:
// Input: nums = [5,3,2]
// Output: 1
// Explanation: We can show that 1 is the smallest number that is not expressible.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func minImpossibleOR(nums []int) int {
    set := make(map[int]bool)
    for _, v := range nums {
        set[v] = true
    }
    for i := 0; i < 31; i++ {
        if !set[1 << i] {
            return 1 << i
        }
    }
    return -1
}

func minImpossibleOR1(nums []int) int {
    mask := 0 
    for _, v := range nums {
        if v & (v - 1) == 0 {
            mask |= v 
        }
    }
    mask = ^mask 
    return mask & -mask
}

func main() {
    // Example 1:
    // Input: nums = [2,1]
    // Output: 4
    // Explanation: 1 and 2 are already present in the array. We know that 3 is expressible, since nums[0] | nums[1] = 2 | 1 = 3. Since 4 is not expressible, we return 4.
    fmt.Println(minImpossibleOR([]int{2, 1})) // 4
    // Example 2:
    // Input: nums = [5,3,2]
    // Output: 1
    // Explanation: We can show that 1 is the smallest number that is not expressible.
    fmt.Println(minImpossibleOR([]int{5,3,2})) // 1

    fmt.Println(minImpossibleOR([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(minImpossibleOR([]int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(minImpossibleOR1([]int{2, 1})) // 4
    fmt.Println(minImpossibleOR1([]int{5,3,2})) // 1
    fmt.Println(minImpossibleOR1([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(minImpossibleOR1([]int{9,8,7,6,5,4,3,2,1})) // 16
}