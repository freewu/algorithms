package main

// 3542. Minimum Operations to Convert All Elements to Zero
// You are given an array nums of size n, consisting of non-negative integers. 
// Your task is to apply some (possibly zero) operations on the array so that all elements become 0.

// In one operation, you can select a subarray [i, j] (where 0 <= i <= j < n) and set all occurrences of the minimum non-negative integer in that subarray to 0.

// Return the minimum number of operations required to make all elements in the array 0.

// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: nums = [0,2]
// Output: 1
// Explanation:
// Select the subarray [1,1] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0].
// Thus, the minimum number of operations required is 1.

// Example 2:
// Input: nums = [3,1,2,1]
// Output: 3
// Explanation:
// Select subarray [1,3] (which is [1,2,1]), where the minimum non-negative integer is 1. Setting all occurrences of 1 to 0 results in [3,0,2,0].
// Select subarray [2,2] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [3,0,0,0].
// Select subarray [0,0] (which is [3]), where the minimum non-negative integer is 3. Setting all occurrences of 3 to 0 results in [0,0,0,0].
// Thus, the minimum number of operations required is 3.

// Example 3:
// Input: nums = [1,2,1,2,1,2]
// Output: 4
// Explanation:
// Select subarray [0,5] (which is [1,2,1,2,1,2]), where the minimum non-negative integer is 1. Setting all occurrences of 1 to 0 results in [0,2,0,2,0,2].
// Select subarray [1,1] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,2,0,2].
// Select subarray [3,3] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,0,0,2].
// Select subarray [5,5] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,0,0,0].
// Thus, the minimum number of operations required is 4.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

// 单调栈
func minOperations(nums []int) int {
    res, stack := 0, nums[:0] // 原地
    for _, v := range nums {
        for len(stack) > 0 && v < stack[len(stack) - 1] {
            stack = stack[:len(stack) - 1] // pop
            res++
        }
        // 如果 v 与栈顶相同，那么 v 与栈顶可以在同一次操作中都变成 0，v 无需入栈
        if len(stack) == 0 || v != stack[len(stack) - 1] {
            stack = append(stack, v)
        }
    }
    if stack[0] == 0 { // 0 不需要操作
        res--
    }
    return res + len(stack)
}

func minOperations1(nums []int) int {
    res , top := 0 , -1
    for i := 0 ; i < len(nums) ; i++ {
        for top >= 0 && nums[i] < nums[top] {
            top-- 
            res++
        }
        if top < 0 || nums[i] != nums[top] {
            top++ 
            nums[top] = nums[i]
        }
    }
    if nums[0] > 0 {
        return res + top + 1
    }
    return res + top
}

func main() {
    // Example 1:
    // Input: nums = [0,2]
    // Output: 1
    // Explanation:
    // Select the subarray [1,1] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0].
    // Thus, the minimum number of operations required is 1.
    fmt.Println(minOperations([]int{0,2})) // 1
    // Example 2:
    // Input: nums = [3,1,2,1]
    // Output: 3
    // Explanation:
    // Select subarray [1,3] (which is [1,2,1]), where the minimum non-negative integer is 1. Setting all occurrences of 1 to 0 results in [3,0,2,0].
    // Select subarray [2,2] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [3,0,0,0].
    // Select subarray [0,0] (which is [3]), where the minimum non-negative integer is 3. Setting all occurrences of 3 to 0 results in [0,0,0,0].
    // Thus, the minimum number of operations required is 3.
    fmt.Println(minOperations([]int{3,1,2,1})) // 3
    // Example 3:
    // Input: nums = [1,2,1,2,1,2]
    // Output: 4
    // Explanation:
    // Select subarray [0,5] (which is [1,2,1,2,1,2]), where the minimum non-negative integer is 1. Setting all occurrences of 1 to 0 results in [0,2,0,2,0,2].
    // Select subarray [1,1] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,2,0,2].
    // Select subarray [3,3] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,0,0,2].
    // Select subarray [5,5] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [0,0,0,0,0,0].
    // Thus, the minimum number of operations required is 4.
    fmt.Println(minOperations([]int{1,2,1,2,1,2})) // 4

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(minOperations1([]int{0,2})) // 1
    fmt.Println(minOperations1([]int{3,1,2,1})) // 3
    fmt.Println(minOperations1([]int{1,2,1,2,1,2})) // 4
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // 9
}