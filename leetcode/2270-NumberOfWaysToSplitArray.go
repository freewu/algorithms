package main

// 2270. Number of Ways to Split Array
// You are given a 0-indexed integer array nums of length n.

// nums contains a valid split at index i if the following are true:
//     1. The sum of the first i + 1 elements is greater than or equal to the sum of the last n - i - 1 elements.
//     2. There is at least one element to the right of i. That is, 0 <= i < n - 1.

// Return the number of valid splits in nums.

// Example 1:
// Input: nums = [10,4,-8,7]
// Output: 2
// Explanation: 
// There are three ways of splitting nums into two non-empty parts:
// - Split nums at index 0. Then, the first part is [10], and its sum is 10. The second part is [4,-8,7], and its sum is 3. Since 10 >= 3, i = 0 is a valid split.
// - Split nums at index 1. Then, the first part is [10,4], and its sum is 14. The second part is [-8,7], and its sum is -1. Since 14 >= -1, i = 1 is a valid split.
// - Split nums at index 2. Then, the first part is [10,4,-8], and its sum is 6. The second part is [7], and its sum is 7. Since 6 < 7, i = 2 is not a valid split.
// Thus, the number of valid splits in nums is 2.

// Example 2:
// Input: nums = [2,3,1,0]
// Output: 2
// Explanation: 
// There are two valid splits in nums:
// - Split nums at index 1. Then, the first part is [2,3], and its sum is 5. The second part is [1,0], and its sum is 1. Since 5 >= 1, i = 1 is a valid split. 
// - Split nums at index 2. Then, the first part is [2,3,1], and its sum is 6. The second part is [0], and its sum is 0. Since 6 >= 0, i = 2 is a valid split.

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func waysToSplitArray(nums []int) int {
    sum, left, right, res := 0, 0, 0, 0
    for _, v := range nums {
        sum += v
    }
    for i := 0; i < len(nums) - 1; i++ {
        // the following is the left and right subarray sum
        left += nums[i]
        right = sum - left
        // if the given constrain is satisfied increment answer
        if left >= right {
            res++
        }
        
    }
    return res
}

func waysToSplitArray1(nums []int) int {
    res, cur, sum := 0, 0, 0
    for _, v := range nums {
        sum += v
    }
    for _, v := range nums[:len(nums) - 1] {
        cur += v 
        sum -= v
        if cur >= sum {
            res++
        }
    }
    return res
}

func waysToSplitArray2(nums []int) int {
    res, cur, sum := 0, 0, 0
    for _, v := range nums {
        sum += v
    }
    for _, v := range nums[:len(nums) - 1] {
        cur += v
        if cur >= sum - cur {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,4,-8,7]
    // Output: 2
    // Explanation: 
    // There are three ways of splitting nums into two non-empty parts:
    // - Split nums at index 0. Then, the first part is [10], and its sum is 10. The second part is [4,-8,7], and its sum is 3. Since 10 >= 3, i = 0 is a valid split.
    // - Split nums at index 1. Then, the first part is [10,4], and its sum is 14. The second part is [-8,7], and its sum is -1. Since 14 >= -1, i = 1 is a valid split.
    // - Split nums at index 2. Then, the first part is [10,4,-8], and its sum is 6. The second part is [7], and its sum is 7. Since 6 < 7, i = 2 is not a valid split.
    // Thus, the number of valid splits in nums is 2.
    fmt.Println(waysToSplitArray([]int{10,4,-8,7})) // 2
    // Example 2:
    // Input: nums = [2,3,1,0]
    // Output: 2
    // Explanation: 
    // There are two valid splits in nums:
    // - Split nums at index 1. Then, the first part is [2,3], and its sum is 5. The second part is [1,0], and its sum is 1. Since 5 >= 1, i = 1 is a valid split. 
    // - Split nums at index 2. Then, the first part is [2,3,1], and its sum is 6. The second part is [0], and its sum is 0. Since 6 >= 0, i = 2 is a valid split.
    fmt.Println(waysToSplitArray([]int{2,3,1,0})) // 2

    fmt.Println(waysToSplitArray1([]int{10,4,-8,7})) // 2
    fmt.Println(waysToSplitArray1([]int{2,3,1,0})) // 2

    fmt.Println(waysToSplitArray2([]int{10,4,-8,7})) // 2
    fmt.Println(waysToSplitArray2([]int{2,3,1,0})) // 2
}
