package main

// 1980. Find Unique Binary String
// Given an array of strings nums containing n unique binary strings each of length n, 
// return a binary string of length n that does not appear in nums. 
// If there are multiple answers, you may return any of them.

// Example 1:
// Input: nums = ["01","10"]
// Output: "11"
// Explanation: "11" does not appear in nums. "00" would also be correct.

// Example 2:
// Input: nums = ["00","01"]
// Output: "11"
// Explanation: "11" does not appear in nums. "10" would also be correct.

// Example 3:
// Input: nums = ["111","011","001"]
// Output: "101"
// Explanation: "101" does not appear in nums. "000", "010", "100", and "110" would also be correct.

// Constraints:
//     n == nums.length
//     1 <= n <= 16
//     nums[i].length == n
//     nums[i] is either '0' or '1'.
//     All the strings of nums are unique.

import "fmt"

func findDifferentBinaryString(nums []string) string {
    res := make([]byte, len(nums))
    for i := range nums {
        if nums[i][i] == '0' {
            res[i] = '1'
        } else {
            res[i] = '0'
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: nums = ["01","10"]
    // Output: "11"
    // Explanation: "11" does not appear in nums. "00" would also be correct.
    fmt.Println(findDifferentBinaryString([]string{"01","10"})) // "11"
    // Example 2:
    // Input: nums = ["00","01"]
    // Output: "11"
    // Explanation: "11" does not appear in nums. "10" would also be correct.
    fmt.Println(findDifferentBinaryString([]string{"00","01"})) // "10"
    // Example 3:
    // Input: nums = ["111","011","001"]
    // Output: "101"
    // Explanation: "101" does not appear in nums. "000", "010", "100", and "110" would also be correct.
    fmt.Println(findDifferentBinaryString([]string{"111","011","001"})) // "000"
}