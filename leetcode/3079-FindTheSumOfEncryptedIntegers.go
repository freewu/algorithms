package main

// 3079. Find the Sum of Encrypted Integers
// You are given an integer array nums containing positive integers. 
// We define a function encrypt such that encrypt(x) replaces every digit in x with the largest digit in x. 
// For example, encrypt(523) = 555 and encrypt(213) = 333.

// Return the sum of encrypted elements.

// Example 1:
// Input: nums = [1,2,3]
// Output: 6
// Explanation: The encrypted elements are [1,2,3]. The sum of encrypted elements is 1 + 2 + 3 == 6.

// Example 2:
// Input: nums = [10,21,31]
// Output: 66
// Explanation: The encrypted elements are [11,22,33]. The sum of encrypted elements is 11 + 22 + 33 == 66.

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 1000

import "fmt"

func sumOfEncryptedInt(nums []int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        largest, multiply := 0, 0
        for v > 0 {
            largest, multiply = max(largest, v % 10), multiply * 10 + 1
            v /= 10
        }
        res += largest * multiply
    }
    return res
}

func sumOfEncryptedInt1(nums []int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    findEncryptedInt := func(num int) int {
        largest, digits := 1, fmt.Sprint(num)
        for _, v := range digits {
            largest = max(largest, int(v - '0'))
        }
        res := largest
        for i := 1; i < len(digits); i++ {
            res = res * 10 + largest
        }
        return res
    }
    for _, v := range nums {
        res += findEncryptedInt(v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 6
    // Explanation: The encrypted elements are [1,2,3]. The sum of encrypted elements is 1 + 2 + 3 == 6.
    fmt.Println(sumOfEncryptedInt([]int{1,2,3})) // 6
    // Example 2:
    // Input: nums = [10,21,31]
    // Output: 66
    // Explanation: The encrypted elements are [11,22,33]. The sum of encrypted elements is 11 + 22 + 33 == 66.
    fmt.Println(sumOfEncryptedInt([]int{10,21,31})) // 66

    fmt.Println(sumOfEncryptedInt1([]int{1,2,3})) // 6
    fmt.Println(sumOfEncryptedInt1([]int{10,21,31})) // 66
}