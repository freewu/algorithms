package main

// 1437. Check If All 1's Are at Least Length K Places Away
// Given an binary array nums and an integer k, 
// return true if all 1's are at least k places away from each other, otherwise return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/15/sample_1_1791.png" />
// Input: nums = [1,0,0,0,1,0,0,1], k = 2
// Output: true
// Explanation: Each of the 1s are at least 2 places away from each other.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/04/15/sample_2_1791.png" />
// Input: nums = [1,0,0,1,0,1], k = 2
// Output: false
// Explanation: The second 1 and third 1 are only one apart from each other.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= k <= nums.length
//     nums[i] is 0 or 1

import "fmt"

func kLengthApart(nums []int, k int) bool {
    t := -1
    for i, v := range nums {
        if v == 1 {
            if t != -1 && nums[t] == 1 && i - t - 1 < k { // i - t - 1 计算与上一个 1 的间隔
                return false
            }
            t = i
        }
    }
    return true
}

func kLengthApart1(nums []int, k int) bool {
    j := -k - 1
    for i, v := range nums {
        if v == 1 {
            if i - j - 1  < k {
                return false
            }
            j = i
        }
    } 
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/04/15/sample_1_1791.png" />
    // Input: nums = [1,0,0,0,1,0,0,1], k = 2
    // Output: true
    // Explanation: Each of the 1s are at least 2 places away from each other.
    fmt.Println(kLengthApart([]int{1,0,0,0,1,0,0,1}, 2)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/04/15/sample_2_1791.png" />
    // Input: nums = [1,0,0,1,0,1], k = 2
    // Output: false
    // Explanation: The second 1 and third 1 are only one apart from each other.
    fmt.Println(kLengthApart([]int{1,0,0,1,0,1}, 2)) // false

    fmt.Println(kLengthApart([]int{0,0,0,0,0,0,0,0,0,0}, 2)) // true
    fmt.Println(kLengthApart([]int{1,1,1,1,1,1,1,1,1,1}, 2)) // false
    fmt.Println(kLengthApart([]int{0,0,0,0,0,1,1,1,1,1}, 2)) // false
    fmt.Println(kLengthApart([]int{1,1,1,1,1,0,0,0,0,0}, 2)) // false
    fmt.Println(kLengthApart([]int{1,0,1,0,1,0,1,0,1,0}, 2)) // false
    fmt.Println(kLengthApart([]int{0,1,0,1,0,1,0,1,0,1}, 2)) // false

    fmt.Println(kLengthApart1([]int{1,0,0,0,1,0,0,1}, 2)) // true
    fmt.Println(kLengthApart1([]int{1,0,0,1,0,1}, 2)) // false
    fmt.Println(kLengthApart1([]int{0,0,0,0,0,0,0,0,0,0}, 2)) // true
    fmt.Println(kLengthApart1([]int{1,1,1,1,1,1,1,1,1,1}, 2)) // false
    fmt.Println(kLengthApart1([]int{0,0,0,0,0,1,1,1,1,1}, 2)) // false
    fmt.Println(kLengthApart1([]int{1,1,1,1,1,0,0,0,0,0}, 2)) // false
    fmt.Println(kLengthApart1([]int{1,0,1,0,1,0,1,0,1,0}, 2)) // false
    fmt.Println(kLengthApart1([]int{0,1,0,1,0,1,0,1,0,1}, 2)) // false
}