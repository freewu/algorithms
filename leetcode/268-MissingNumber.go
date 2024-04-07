package main

// 268. Missing Number
// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

// Example 1:
// Input: nums = [3,0,1]
// Output: 2
// Explanation:
//     n = 3 since there are 3 numbers,
//     so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

// Example 2:
// Input: nums = [0,1]
// Output: 2
// Explanation:
//     n = 2 since there are 2 numbers,
//     so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.

// Example 3:
// Input: nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
// Explanation: 
//     n = 9 since there are 9 numbers,
//     so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^4
//     0 <= nums[i] <= n
//     All the numbers of nums are unique.

import "fmt"

// 异或  O(n)  O(1) 
func missingNumber(nums []int) int {
    xor, i := 0, 0
    for i = 0; i < len(nums); i++ {
        xor = xor ^ i ^ nums[i] // i 和 nums[i]  值出现两次的都会抵消掉
    }
    return xor ^ i
}

// 存在 map 中 O(n)  O(n) 
func missingNumber1(nums []int) int {
    m := make([]bool, len(nums)+1) 
    for _, n := range nums {
        m[n] = true
    }
    for i, exists := range m {
        if !exists {
            return i
        }
    }
    return 0
}

func main() {
    // n = 3 since there are 3 numbers,
    // so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
    fmt.Printf("missingNumber([]int{ 3,0,1 }) = %v\n",missingNumber([]int{ 3,0,1 })) // 2
    // n = 2 since there are 2 numbers,
    // so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
    fmt.Printf("missingNumber([]int{ 0,1 }) = %v\n",missingNumber([]int{ 0,1})) // 2
    // n = 9 since there are 9 numbers,
    // so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
    fmt.Printf("missingNumber([]int{ 9,6,4,2,3,5,7,0,1 }) = %v\n",missingNumber([]int{ 9,6,4,2,3,5,7,0,1 })) // 8

    fmt.Printf("missingNumber1([]int{ 3,0,1 }) = %v\n",missingNumber1([]int{ 3,0,1 })) // 2
    fmt.Printf("missingNumber1([]int{ 0,1 }) = %v\n",missingNumber1([]int{ 0,1})) // 2
    fmt.Printf("missingNumber1([]int{ 9,6,4,2,3,5,7,0,1 }) = %v\n",missingNumber1([]int{ 9,6,4,2,3,5,7,0,1 })) // 8
}