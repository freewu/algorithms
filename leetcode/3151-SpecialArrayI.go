package main

// 3151. Special Array I
// An array is considered special if every pair of its adjacent elements contains two numbers with different parity.
// You are given an array of integers nums. 
// Return true if nums is a special array, otherwise, return false.

// Example 1:
// Input: nums = [1]
// Output: true
// Explanation:
// There is only one element. So the answer is true.

// Example 2:
// Input: nums = [2,1,4]
// Output: true
// Explanation:
// There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.

// Example 3:
// Input: nums = [4,3,1,6]
// Output: false
// Explanation:
// nums[1] and nums[2] are both odd. So the answer is false.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func isArraySpecial(nums []int) bool {
    odd := func(n int) bool {
        if n % 2 != 0 { return true }
        return false
    }
    for i := 0; i < len(nums) - 1; i++ {
        if (odd(nums[i]) && odd(nums[i+1])) || 
           (!odd(nums[i+1]) && !odd(nums[i])) { // 两个同为奇数或偶数 就不是 特殊数组 
            return false
        }
    }
    return true
}

func isArraySpecial1(nums []int) bool {
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] % 2 == nums[i] % 2 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [1]
    // Output: true
    // Explanation:
    // There is only one element. So the answer is true.
    fmt.Println(isArraySpecial([]int{1})) // true
    // Example 2:
    // Input: nums = [2,1,4]
    // Output: true
    // Explanation:
    // There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.
    fmt.Println(isArraySpecial([]int{2,1,4})) // true
    // Example 3:
    // Input: nums = [4,3,1,6]
    // Output: false
    // Explanation:
    // nums[1] and nums[2] are both odd. So the answer is false.
    fmt.Println(isArraySpecial([]int{4,3,1,6})) // false

    fmt.Println(isArraySpecial1([]int{1})) // true
    fmt.Println(isArraySpecial1([]int{2,1,4})) // true
    fmt.Println(isArraySpecial1([]int{4,3,1,6})) // false
}