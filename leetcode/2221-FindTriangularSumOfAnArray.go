package main

// 2221. Find Triangular Sum of an Array
// You are given a 0-indexed integer array nums, where nums[i] is a digit between 0 and 9 (inclusive).

// The triangular sum of nums is the value of the only element present in nums after the following process terminates:
//     1. Let nums comprise of n elements. 
//        If n == 1, end the process. Otherwise, create a new 0-indexed integer array newNums of length n - 1.
//     2. For each index i, where 0 <= i < n - 1, 
//        assign the value of newNums[i] as (nums[i] + nums[i+1]) % 10, where % denotes modulo operator.
//     3. Replace the array nums with newNums.
//     4. Repeat the entire process starting from step 1.

// Return the triangular sum of nums.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/22/ex1drawio.png" />
// Input: nums = [1,2,3,4,5]
// Output: 8
// Explanation:
// The above diagram depicts the process from which we obtain the triangular sum of the array.

// Example 2:
// Input: nums = [5]
// Output: 5
// Explanation:
// Since there is only one element in nums, the triangular sum is the value of that element itself.

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 9

import "fmt" 

// Greedy
func triangularSum(nums []int) int {
    for len(nums) != 1 {
        arr := make([]int, len(nums) - 1)
        for i := 0; i < len(nums) - 1; i++ {
            arr[i] = (nums[i] + nums[i + 1]) % 10
        }
        nums = arr
    }
    return nums[0]
}

func triangularSum1(nums []int) int {
    n := len(nums)
    for n != 1 {
        for i := 0 ;i < n - 1; i++ {
            nums[i] = (nums[i] + nums[i + 1]) % 10
        }
        n--
    }
    return nums[0]
}

func triangularSum2(nums []int) int {
    // 每循环一轮，数组长度就减一
    for n := len(nums) - 1; n > 0; n-- {
        for i := range n {
            nums[i] = (nums[i] + nums[i + 1]) % 10
        }
    }
    return nums[0]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/22/ex1drawio.png" />
    // Input: nums = [1,2,3,4,5]
    // Output: 8
    // Explanation:
    // The above diagram depicts the process from which we obtain the triangular sum of the array.
    fmt.Println(triangularSum([]int{1,2,3,4,5})) // 8
    // Example 2:
    // Input: nums = [5]
    // Output: 5
    // Explanation:
    // Since there is only one element in nums, the triangular sum is the value of that element itself.
    fmt.Println(triangularSum([]int{5})) // 5

    fmt.Println(triangularSum([]int{9,9,9,9,9,9,9,9,9})) // 4
    fmt.Println(triangularSum([]int{1,1,1,1,1,1,1,1,1})) // 6
    fmt.Println(triangularSum([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(triangularSum([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(triangularSum1([]int{1,2,3,4,5})) // 8
    fmt.Println(triangularSum1([]int{5})) // 5
    fmt.Println(triangularSum1([]int{9,9,9,9,9,9,9,9,9})) // 4
    fmt.Println(triangularSum1([]int{1,1,1,1,1,1,1,1,1})) // 6
    fmt.Println(triangularSum1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(triangularSum1([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(triangularSum2([]int{1,2,3,4,5})) // 8
    fmt.Println(triangularSum2([]int{5})) // 5
    fmt.Println(triangularSum2([]int{9,9,9,9,9,9,9,9,9})) // 4
    fmt.Println(triangularSum2([]int{1,1,1,1,1,1,1,1,1})) // 6
    fmt.Println(triangularSum2([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(triangularSum2([]int{9,8,7,6,5,4,3,2,1})) // 0
}