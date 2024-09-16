package main

// 1144. Decrease Elements To Make Array Zigzag
// Given an array nums of integers, a move consists of choosing any element and decreasing it by 1.

// An array A is a zigzag array if either:
//     Every even-indexed element is greater than adjacent elements, ie. A[0] > A[1] < A[2] > A[3] < A[4] > ...
//     OR, every odd-indexed element is greater than adjacent elements, ie. A[0] < A[1] > A[2] < A[3] > A[4] < ...

// Return the minimum number of moves to transform the given array nums into a zigzag array.

// Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation: We can decrease 2 to 0 or 3 to 1.

// Example 2:
// Input: nums = [9,6,1,6,2]
// Output: 4

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 1000

import "fmt"

func movesToMakeZigzag(nums []int) int {
    nums1 := make([]int, len(nums))
    copy(nums1, nums)
    even := func (nums []int) int {
        step := 0
        for i := 0; i < len(nums)-1; i++ {
            if i % 2 == 0 {
                if nums[i] <= nums[i+1] {
                    step += nums[i+1] - nums[i] + 1
                    nums[i+1] = nums[i] - 1
                }
            } else {
                if nums[i] >= nums[i+1] {
                    step += nums[i] - nums[i+1] + 1
                    nums[i] = nums[i+1] - 1
                }
            }
        }
        return step
    }
    odd := func (nums []int) int {
        step := 0
        for i := 0; i < len(nums)-1; i++ {
            if i % 2 != 0 {
                if nums[i] <= nums[i+1] {
                    step += nums[i+1] - nums[i] + 1
                    nums[i+1] = nums[i] - 1
                }
            } else {
                if nums[i] >= nums[i+1] {
                    step += nums[i] - nums[i+1] + 1
                    nums[i] = nums[i+1] - 1
                }
            }
        }
        return step
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(even(nums), odd(nums1))
}

func movesToMakeZigzag1(nums []int) int {
    odd, even := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(nums); i += 2 {
        target := nums[i]
        if i - 1 >= 0 {
            target = min(target, nums[i-1]-1)
        }
        if i+1 < len(nums) {
            target = min(target, nums[i+1]-1)
        }
        even += nums[i] - target
    }
    for i := 0; i < len(nums); i += 2 {
        target := nums[i]
        if i-1 >= 0 {
            target = min(target, nums[i-1]-1)
        }
        if i+1 < len(nums) {
            target = min(target, nums[i+1]-1)
        }
        odd += nums[i] - target
    }
    return min(even, odd)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 2
    // Explanation: We can decrease 2 to 0 or 3 to 1.
    fmt.Println(movesToMakeZigzag([]int{1,2,3})) // 2
    // Example 2:
    // Input: nums = [9,6,1,6,2]
    // Output: 4
    fmt.Println(movesToMakeZigzag([]int{9,6,1,6,2})) // 4

    fmt.Println(movesToMakeZigzag1([]int{1,2,3})) // 2
    fmt.Println(movesToMakeZigzag1([]int{9,6,1,6,2})) // 4
}