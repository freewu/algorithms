package main

// 3354. Make Array Elements Equal to Zero
// You are given an integer array nums.

// Start by selecting a starting position curr such that nums[curr] == 0, and choose a movement direction of either left or right.

// After that, you repeat the following process:
//     1. If curr is out of the range [0, n - 1], this process ends.
//     2. If nums[curr] == 0, move in the current direction by incrementing curr if you are moving right, or decrementing curr if you are moving left.
//     3. Else if nums[curr] > 0:
//         3.1 Decrement nums[curr] by 1.
//         3.2 Reverse your movement direction (left becomes right and vice versa).
//         3.3 Take a step in your new direction.

// A selection of the initial position curr and movement direction is considered valid if every element in nums becomes 0 by the end of the process.

// Return the number of possible valid selections.

// Example 1:
// Input: nums = [1,0,2,0,3]
// Output: 2
// Explanation:
// The only possible valid selections are the following:
// Choose curr = 3, and a movement direction to the left.
// [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,0].
// Choose curr = 3, and a movement direction to the right.
// [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0].

// Example 2:
// Input: nums = [2,3,4,0,4,1,0]
// Output: 0
// Explanation:
// There are no possible valid selections.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 100
//     There is at least one element i where nums[i] == 0.

import "fmt"

func countValidSelections(nums []int) int {
    res, curr, sum := 0, 0, 0
    for _, v := range nums {
        sum += v
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        if v == 0 {
            s := abs(curr * 2 - sum)
            if s == 0 { // both left and right direction for moving is okay
                res += 2
            } else if s == 1 { // either move left or right
                res++
            }
        }
        curr += v
    }
    return res
}

func countValidSelections1(nums []int) int {
    res, suffix, prefix := 0, 0, 0
    for _, v := range nums {
        suffix += v
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        prefix += v
        suffix -= v
        diff := abs(prefix - suffix)
        if v > 0 || diff > 1 { continue }
        if diff == 0 {
            res++
        }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,2,0,3]
    // Output: 2
    // Explanation:
    // The only possible valid selections are the following:
    // Choose curr = 3, and a movement direction to the left.
    // [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,0].
    // Choose curr = 3, and a movement direction to the right.
    // [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0].
    fmt.Println(countValidSelections([]int{1,0,2,0,3})) // 2
    // Example 2:
    // Input: nums = [2,3,4,0,4,1,0]
    // Output: 0
    // Explanation:
    // There are no possible valid selections.
    fmt.Println(countValidSelections([]int{2,3,4,0,4,1,0})) // 0

    fmt.Println(countValidSelections([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countValidSelections([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(countValidSelections1([]int{1,0,2,0,3})) // 2
    fmt.Println(countValidSelections1([]int{2,3,4,0,4,1,0})) // 0
    fmt.Println(countValidSelections1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countValidSelections1([]int{9,8,7,6,5,4,3,2,1})) // 0
}