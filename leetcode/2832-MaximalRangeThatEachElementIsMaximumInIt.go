package main

// 2832. Maximal Range That Each Element Is Maximum in It
// You are given a 0-indexed array nums of distinct integers.

// Let us define a 0-indexed array ans of the same length as nums in the following way:
//     ans[i] is the maximum length of a subarray nums[l..r], 
//     such that the maximum element in that subarray is equal to nums[i].

// Return the array ans.

// Note that a subarray is a contiguous part of the array.

// Example 1:
// Input: nums = [1,5,4,3,6]
// Output: [1,4,2,1,5]
// Explanation: For nums[0] the longest subarray in which 1 is the maximum is nums[0..0] so ans[0] = 1.
// For nums[1] the longest subarray in which 5 is the maximum is nums[0..3] so ans[1] = 4.
// For nums[2] the longest subarray in which 4 is the maximum is nums[2..3] so ans[2] = 2.
// For nums[3] the longest subarray in which 3 is the maximum is nums[3..3] so ans[3] = 1.
// For nums[4] the longest subarray in which 6 is the maximum is nums[0..4] so ans[4] = 5.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: [1,2,3,4,5]
// Explanation: For nums[i] the longest subarray in which it's the maximum is nums[0..i] so ans[i] = i + 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     All elements in nums are distinct.

import "fmt"

func maximumLengthOfRanges(nums []int) []int {
    n := len(nums)
    left, right := make([]int, n), make([]int, n)
    for i := range left {
        left[i], right[i] = -1, n
    }
    stack := []int{}
    for i, v := range nums {
        for len(stack) > 0 && nums[stack[len(stack) - 1]] <= v {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) > 0 {
            left[i] = stack[len(stack) - 1]
        }
        stack = append(stack, i)
    }
    stack = []int{}
    for i := n - 1; i >= 0; i-- {
        for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
            stack = stack[:len(stack) - 1]
        }
        if len(stack) > 0 {
            right[i] = stack[len(stack) - 1]
        }
        stack = append(stack, i)
    }
    res := make([]int, n)
    for i := range res {
        res[i] = right[i] - left[i] - 1
    }
    return res
}

// Monotonic Stack单调栈
func maximumLengthOfRanges1(nums []int) []int {
    // 每个元素为最大值:只要知道left,right左右两边 > candi的元素，即可确定candi为最大值的最大范围
    n := len(nums) 
    // candi's范围计算:[l, r] => (r - l - 1)，注意 nums[0],nums[n-1]首尾两元素的取值: 
    res, arr, top := make([]int, n), make([]int, 0), 0 
    for i := 0; i < n; i += 1 {
        cur := nums[i] 
        for top > 0 && nums[arr[top - 1]] < cur { // 如果栈不为空，并且栈顶存在比cur小的元素，pop出去小的元素:
            res[arr[top - 1]] += i  // 知道top元素的nextGreater元素位置
            arr = arr[:len(arr) - 1]
            top--
        }
        if top > 0 {
            res[i] -= (arr[top - 1] + 1)  //  道cur's LEFT NearestGreater元素位置
        }
        arr = append(arr, i)
        top++
    }
    for top > 0 {
        i := arr[top - 1]
        top--
        res[i] += n // 确定RIGHT NearestGreater元素位置
        arr = arr[:top]
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [1,5,4,3,6]
    // Output: [1,4,2,1,5]
    // Explanation: For nums[0] the longest subarray in which 1 is the maximum is nums[0..0] so ans[0] = 1.
    // For nums[1] the longest subarray in which 5 is the maximum is nums[0..3] so ans[1] = 4.
    // For nums[2] the longest subarray in which 4 is the maximum is nums[2..3] so ans[2] = 2.
    // For nums[3] the longest subarray in which 3 is the maximum is nums[3..3] so ans[3] = 1.
    // For nums[4] the longest subarray in which 6 is the maximum is nums[0..4] so ans[4] = 5.
    fmt.Println(maximumLengthOfRanges([]int{1,5,4,3,6})) // [1,4,2,1,5]
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: [1,2,3,4,5]
    // Explanation: For nums[i] the longest subarray in which it's the maximum is nums[0..i] so ans[i] = i + 1.
    fmt.Println(maximumLengthOfRanges([]int{1,2,3,4,5})) // [1,2,3,4,5]

    fmt.Println(maximumLengthOfRanges1([]int{1,5,4,3,6})) // [1,4,2,1,5]
    fmt.Println(maximumLengthOfRanges1([]int{1,2,3,4,5})) // [1,2,3,4,5]
}