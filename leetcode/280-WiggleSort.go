package main 

// 280. Wiggle Sort
// Given an integer array nums, reorder it such that nums[0] <= nums[1] >= nums[2] <= nums[3]....
// You may assume the input array always has a valid answer.

// Example 1:
// Input: nums = [3,5,2,1,6,4]
// Output: [3,5,1,6,2,4]
// Explanation: [1,6,2,5,3,4] is also accepted.

// Example 2:
// Input: nums = [6,6,5,6,3,8]
// Output: [6,6,5,6,3,8]
 
// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     0 <= nums[i] <= 10^4
//     It is guaranteed that there will be an answer for the given input nums.
 
// Follow up: Could you solve the problem in O(n) time complexity?

import "fmt"

func wiggleSort(nums []int)  {
    l := len(nums)
    for i := 1; i < l; i+=2 {
        if nums[i-1] > nums[i] { // 前一个数比当前数大，则交换
            nums[i], nums[i-1] = nums[i-1], nums[i]
        }
        if i == l - 1 {
            return
        }
        if nums[i+1] > nums[i] { // 后一个数比当前数大，则交换
            nums[i], nums[i+1] = nums[i+1], nums[i]
        }
    }
}

func main() {
    // Example 1:
    // Input: nums = [3,5,2,1,6,4]
    // Output: [3,5,1,6,2,4]
    // Explanation: [1,6,2,5,3,4] is also accepted.
    arr1 := []int{3,5,2,1,6,4}
    fmt.Println("before: ", arr1)
    wiggleSort(arr1)
    fmt.Println("after: ", arr1) // [3,5,1,6,2,4]
    // Example 2:
    // Input: nums = [6,6,5,6,3,8]
    // Output: [6,6,5,6,3,8]
    arr2 := []int{6,6,5,6,3,8}
    fmt.Println("before: ", arr2)
    wiggleSort(arr2)
    fmt.Println("after: ", arr2) // [6,6,5,6,3,8]
}