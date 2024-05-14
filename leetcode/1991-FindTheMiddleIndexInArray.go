package main

// 1991. Find the Middle Index in Array
// Given a 0-indexed integer array nums, find the leftmost middleIndex (i.e., the smallest amongst all the possible ones).
// A middleIndex is an index where nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1].
// If middleIndex == 0, the left side sum is considered to be 0. 
// Similarly, if middleIndex == nums.length - 1, the right side sum is considered to be 0.
// Return the leftmost middleIndex that satisfies the condition, or -1 if there is no such index.

// Example 1:
// Input: nums = [2,3,-1,8,4]
// Output: 3
// Explanation: The sum of the numbers before index 3 is: 2 + 3 + -1 = 4
// The sum of the numbers after index 3 is: 4 = 4

// Example 2:
// Input: nums = [1,-1,4]
// Output: 2
// Explanation: The sum of the numbers before index 2 is: 1 + -1 = 0
// The sum of the numbers after index 2 is: 0

// Example 3:
// Input: nums = [2,5]
// Output: -1
// Explanation: There is no valid middleIndex.
 
// Constraints:
//     1 <= nums.length <= 100
//     -1000 <= nums[i] <= 1000

import "fmt"

func findMiddleIndex(nums []int) int {
    leftPrefixSum, rightPrefixSum := make([]int, len(nums)), make([]int, len(nums))
    leftPrefixSum[0] = nums[0]
    rightPrefixSum[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1; i < len(nums); i++ {
        leftPrefixSum[i] = nums[i] + leftPrefixSum[i - 1]
        endIdx := len(nums) - i - 1 
        rightPrefixSum[endIdx] = nums[endIdx] + rightPrefixSum[endIdx + 1]
    }
    for i := 0; i < len(nums); i++ {
        if leftPrefixSum[i] - nums[i] == rightPrefixSum[i] - nums[i] {
            return i
        }
    }
    return -1
}

func findMiddleIndex1(nums []int) int {
    right, left := 0, 0
    for _, num := range nums {
        right += num
    }   
    for i := 0; i < len(nums); i++ {
        right -= nums[i]
        if left == right {
            return i
        }
        left += nums[i]
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,3,-1,8,4]
    // Output: 3
    // Explanation: The sum of the numbers before index 3 is: 2 + 3 + -1 = 4
    // The sum of the numbers after index 3 is: 4 = 4
    fmt.Println(findMiddleIndex([]int{2,3,-1,8,4})) // 3
    // Example 2:
    // Input: nums = [1,-1,4]
    // Output: 2
    // Explanation: The sum of the numbers before index 2 is: 1 + -1 = 0
    // The sum of the numbers after index 2 is: 0
    fmt.Println(findMiddleIndex([]int{1,-1,4})) // 2
    // Example 3:
    // Input: nums = [2,5]
    // Output: -1
    // Explanation: There is no valid middleIndex.
    fmt.Println(findMiddleIndex([]int{2, 5})) // -1

    fmt.Println(findMiddleIndex1([]int{2,3,-1,8,4})) // 3
    fmt.Println(findMiddleIndex1([]int{1,-1,4})) // 2
    fmt.Println(findMiddleIndex1([]int{2, 5})) // -1
}