package main

// 1558. Minimum Numbers of Function Calls to Make Target Array
// You are given an integer array nums. 
// You have an integer array arr of the same length with all values set to 0 initially. 
// You also have the following modify function:
// <img src="https://assets.leetcode.com/uploads/2020/07/10/sample_2_1887.png" />

// You want to use the modify function to convert arr to nums using the minimum number of calls.

// Return the minimum number of function calls to make nums from arr.

// The test cases are generated so that the answer fits in a 32-bit signed integer.

// Example 1:
// Input: nums = [1,5]
// Output: 5
// Explanation: Increment by 1 (second element): [0, 0] to get [0, 1] (1 operation).
// Double all the elements: [0, 1] -> [0, 2] -> [0, 4] (2 operations).
// Increment by 1 (both elements)  [0, 4] -> [1, 4] -> [1, 5] (2 operations).
// Total of operations: 1 + 2 + 2 = 5.

// Example 2:
// Input: nums = [2,2]
// Output: 3
// Explanation: Increment by 1 (both elements) [0, 0] -> [0, 1] -> [1, 1] (2 operations).
// Double all the elements: [1, 1] -> [2, 2] (1 operation).
// Total of operations: 2 + 1 = 3.

// Example 3:
// Input: nums = [4,2,5]
// Output: 6
// Explanation: (initial)[0,0,0] -> [1,0,0] -> [1,0,1] -> [2,0,2] -> [2,1,2] -> [4,2,4] -> [4,2,5](nums).

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "math/bits"

func minOperations(nums []int) int {
    sort.Ints(nums)
    var solve func(nums []int) int
    solve = func (nums []int) int {
        if nums[len(nums)-1] == 0 { return 0 }
        prod, add := 0, 0
        for i := 0; i < len(nums); i++ {
            if nums[i] % 2 == 1 {
                nums[i] = nums[i]-1
                add++
            }
            if nums[i] != 0 {
                nums[i] = nums[i] / 2
                if prod == 0 { prod = 1 }
            }
        }
        return solve(nums) + prod + add
    }
    return solve(nums)
}

func minOperations1(nums []int) int {
    res, mx := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        mx = max(mx, v)
        res += bits.OnesCount(uint(v))
    }
    if mx > 0 {
        res += bits.Len(uint(mx)) - 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,5]
    // Output: 5
    // Explanation: Increment by 1 (second element): [0, 0] to get [0, 1] (1 operation).
    // Double all the elements: [0, 1] -> [0, 2] -> [0, 4] (2 operations).
    // Increment by 1 (both elements)  [0, 4] -> [1, 4] -> [1, 5] (2 operations).
    // Total of operations: 1 + 2 + 2 = 5.
    fmt.Println(minOperations([]int{1,5})) // 5
    // Example 2:
    // Input: nums = [2,2]
    // Output: 3
    // Explanation: Increment by 1 (both elements) [0, 0] -> [0, 1] -> [1, 1] (2 operations).
    // Double all the elements: [1, 1] -> [2, 2] (1 operation).
    // Total of operations: 2 + 1 = 3.
    fmt.Println(minOperations([]int{2,2})) // 3
    // Example 3:
    // Input: nums = [4,2,5]
    // Output: 6
    // Explanation: (initial)[0,0,0] -> [1,0,0] -> [1,0,1] -> [2,0,2] -> [2,1,2] -> [4,2,4] -> [4,2,5](nums).
    fmt.Println(minOperations([]int{4,2,5})) // 6

    fmt.Println(minOperations1([]int{1,5})) // 5
    fmt.Println(minOperations1([]int{2,2})) // 3
    fmt.Println(minOperations1([]int{4,2,5})) // 6
}