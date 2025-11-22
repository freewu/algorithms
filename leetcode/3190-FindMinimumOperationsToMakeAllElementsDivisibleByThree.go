package main

// 3190. Find Minimum Operations to Make All Elements Divisible by Three
// You are given an integer array nums. In one operation, you can add or subtract 1 from any element of nums.
// Return the minimum number of operations to make all elements of nums divisible by 3.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 3
// Explanation:
// All array elements can be made divisible by 3 using 3 operations:
// Subtract 1 from 1.
// Add 1 to 2.
// Subtract 1 from 4.

// Example 2:
// Input: nums = [3,6,9]
// Output: 0

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 50

import "fmt"

func minimumOperations(nums []int) int {
    res := 0
    for _, v := range nums {
        if v % 3 != 0 { // 只会有两种 余数 余1 余2 余1则 +1 余2则 -1 都是1步,所有找出不被3整除的数量即可
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 3
    // Explanation:
    // All array elements can be made divisible by 3 using 3 operations:
    // Subtract 1 from 1.
    // Add 1 to 2.
    // Subtract 1 from 4.
    fmt.Println(minimumOperations([]int{1,2,3,4})) // 3
    // Example 2:
    // Input: nums = [3,6,9]
    // Output: 0
    fmt.Println(minimumOperations([]int{3,6,9})) // 0

    fmt.Println(minimumOperations([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(minimumOperations([]int{9,8,7,6,5,4,3,2,1})) // 6
}