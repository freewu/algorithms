package main

// 3232. Find if Digit Game Can Be Won
// You are given an array of positive integers nums.

// Alice and Bob are playing a game. 
// In the game, Alice can choose either all single-digit numbers or all double-digit numbers from nums, and the rest of the numbers are given to Bob. 
// Alice wins if the sum of her numbers is strictly greater than the sum of Bob's numbers.

// Return true if Alice can win this game, otherwise, return false.

// Example 1:
// Input: nums = [1,2,3,4,10]
// Output: false
// Explanation:
// Alice cannot win by choosing either single-digit or double-digit numbers.

// Example 2:
// Input: nums = [1,2,3,4,5,14]
// Output: true
// Explanation:
// Alice can win by choosing single-digit numbers which have a sum equal to 15.

// Example 3:
// Input: nums = [5,5,5,25]
// Output: true
// Explanation:
// Alice can win by choosing double-digit numbers which have a sum equal to 25.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 99

import "fmt"

func canAliceWin(nums []int) bool {
    one, two := 0, 0
    for _, v := range nums {
        if v >= 10 {
            two += v
        } else {
            one += v
        }
    }
    if one == two { return false }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,10]
    // Output: false
    // Explanation:
    // Alice cannot win by choosing either single-digit or double-digit numbers.
    fmt.Println(canAliceWin([]int{1,2,3,4,10})) // false
    // Example 2:
    // Input: nums = [1,2,3,4,5,14]
    // Output: true
    // Explanation:
    // Alice can win by choosing single-digit numbers which have a sum equal to 15.
    fmt.Println(canAliceWin([]int{1,2,3,4,5,14})) // true
    // Example 3:
    // Input: nums = [5,5,5,25]
    // Output: true
    // Explanation:
    // Alice can win by choosing double-digit numbers which have a sum equal to 25.
    fmt.Println(canAliceWin([]int{5,5,5,25})) // true
}