package main

// 915. Partition Array into Disjoint Intervals
// Given an integer array nums, partition it into two (contiguous) subarrays left and right so that:
//     Every element in left is less than or equal to every element in right.
//     left and right are non-empty.
//     left has the smallest possible size.

// Return the length of left after such a partitioning.

// Test cases are generated such that partitioning exists.

// Example 1:
// Input: nums = [5,0,3,8,6]
// Output: 3
// Explanation: left = [5,0,3], right = [8,6]

// Example 2:
// Input: nums = [1,1,1,0,6,12]
// Output: 4
// Explanation: left = [1,1,1,0], right = [6,12]

// Constraints:
//     2 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^6
//     There is at least one valid answer for the given input.

import "fmt"

func partitionDisjoint(nums []int) int {
    localmx, mx, index := nums[0], nums[0], 0
    for i := range nums {
        if localmx > nums[i] {
            localmx = mx
            index = i
        } else {
            if nums[i] > mx {
                mx = nums[i]
            }
        }
    }
    return index + 1
}

func partitionDisjoint1(nums []int) int {
    leftMax, curMax, index := nums[0], nums[0], 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums) - 1; i++ {
        curMax = max(curMax, nums[i])
        if nums[i] < leftMax {
            leftMax = curMax
            index = i
        }
    }
    return index + 1
}

func main() {
    // Example 1:
    // Input: nums = [5,0,3,8,6]
    // Output: 3
    // Explanation: left = [5,0,3], right = [8,6]
    fmt.Println(partitionDisjoint([]int{5,0,3,8,6})) // 3
    // Example 2:
    // Input: nums = [1,1,1,0,6,12]
    // Output: 4
    // Explanation: left = [1,1,1,0], right = [6,12]
    fmt.Println(partitionDisjoint([]int{1,1,1,0,6,12})) // 4

    fmt.Println(partitionDisjoint1([]int{5,0,3,8,6})) // 3
    fmt.Println(partitionDisjoint1([]int{1,1,1,0,6,12})) // 4
}