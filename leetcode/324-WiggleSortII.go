package main

// 324. Wiggle Sort II
// Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
// You may assume the input array always has a valid answer.

// Example 1:
// Input: nums = [1,5,1,1,6,4]
// Output: [1,6,1,5,1,4]
// Explanation: [1,4,1,5,1,6] is also accepted.

// Example 2:
// Input: nums = [1,3,2,2,3,1]
// Output: [2,3,1,3,1,2]
 
// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     0 <= nums[i] <= 5000
//     It is guaranteed that there will be an answer for the given input nums.
    
// Follow Up: Can you do it in O(n) time and/or in-place with O(1) extra space?

import "fmt"
import "sort"

// 双指针
func wiggleSort(nums []int) {
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    sort.Ints(sorted)

    mid, right := (len(nums) - 1) / 2, len(nums) - 1
    for i := range nums {
        if i % 2 == 0 { // 如果为偶数需要从中间开位置值交换到当前
            nums[i] = sorted[mid]
            mid--
        } else {// 为奇数从末尾
            nums[i] = sorted[right]
            right--
        }
    }
}

func main() {
    // Explanation: [1,4,1,5,1,6] is also accepted.
    arr1 := []int{1,5,1,1,6,4}
    fmt.Println("before arr1 = ",arr1)
    wiggleSort(arr1)
    fmt.Println("after arr1 = ",arr1) // [1,6,1,5,1,4]

    arr2 := []int{1,3,2,2,3,1}
    fmt.Println("before arr2 = ",arr2)
    wiggleSort(arr2)
    fmt.Println("after arr2 = ",arr2) // [2,3,1,3,1,2]
}