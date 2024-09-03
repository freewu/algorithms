package main

// 2708. Maximum Strength of a Group
// You are given a 0-indexed integer array nums representing the score of students in an exam. 
// The teacher would like to form one non-empty group of students with maximal strength, 
// where the strength of a group of students of indices i0, i1, i2, ... , ik is defined as nums[i0] * nums[i1] * nums[i2] * ... * nums[ik​].

// Return the maximum strength of a group the teacher can create.

// Example 1:
// Input: nums = [3,-1,-5,2,5,-9]
// Output: 1350
// Explanation: One way to form a group of maximal strength is to group the students at indices [0,2,3,4,5]. Their strength is 3 * (-5) * 2 * 5 * (-9) = 1350, which we can show is optimal.

// Example 2:
// Input: nums = [-4,-5,-4]
// Output: 20
// Explanation: Group the students at indices [0, 1] . Then, we’ll have a resulting strength of 20. We cannot achieve greater strength.

// Constraints:
//     1 <= nums.length <= 13
//     -9 <= nums[i] <= 9

import "fmt"
import "sort"

func maxStrength(nums []int) int64 {
    if len(nums) == 1 { return int64(nums[0]) }
    sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
    res, i := 0, 0
    for i < len(nums) && nums[i] > 0 {
        if res == 0 {
            res = 1
        }
        res *= nums[i] // 累乘
        i++
    }
    for i < len(nums) && nums[i] == 0 {
        i++
    }
    neg := nums[i:]
    n := len(neg)
    for j := n % 2; j < n; j++ {
        if res == 0 {
            res = neg[n % 2]
            continue
        }
        res *= neg[j]
    }
    return int64(res)
}

func maxStrength1(nums []int) int64 {
    mn, mx := nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums[1:] {
        mn, mx = min(min(mn, v), min(mn * v, mx * v)), max(max(mx, v), max(mn * v, mx * v))
    }
    return int64(mx)
}

func main() {
    // Example 1:
    // Input: nums = [3,-1,-5,2,5,-9]
    // Output: 1350
    // Explanation: One way to form a group of maximal strength is to group the students at indices [0,2,3,4,5]. Their strength is 3 * (-5) * 2 * 5 * (-9) = 1350, which we can show is optimal.
    fmt.Println(maxStrength([]int{3,-1,-5,2,5,-9})) // 1350
    // Example 2:
    // Input: nums = [-4,-5,-4]
    // Output: 20
    // Explanation: Group the students at indices [0, 1] . Then, we’ll have a resulting strength of 20. We cannot achieve greater strength.
    fmt.Println(maxStrength([]int{-4,-5,-4})) // 20

    fmt.Println(maxStrength1([]int{3,-1,-5,2,5,-9})) // 1350
    fmt.Println(maxStrength1([]int{-4,-5,-4})) // 20
}