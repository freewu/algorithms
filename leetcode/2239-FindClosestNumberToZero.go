package main

// 2239. Find Closest Number to Zero
// Given an integer array nums of size n, return the number with the value closest to 0 in nums. 
// If there are multiple answers, return the number with the largest value.

// Example 1:
// Input: nums = [-4,-2,1,4,8]
// Output: 1
// Explanation:
// The distance from -4 to 0 is |-4| = 4.
// The distance from -2 to 0 is |-2| = 2.
// The distance from 1 to 0 is |1| = 1.
// The distance from 4 to 0 is |4| = 4.
// The distance from 8 to 0 is |8| = 8.
// Thus, the closest number to 0 in the array is 1.

// Example 2:
// Input: nums = [2,-1,1]
// Output: 1
// Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.

// Constraints:
//     1 <= n <= 1000
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func findClosestNumber(nums []int) int {
    res, diff := 0, 100_001
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        p := abs(v)
        if p < diff {
            diff, res = p, v
        } else if p == diff {
            if v > res {
                res = v
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-4,-2,1,4,8]
    // Output: 1
    // Explanation:
    // The distance from -4 to 0 is |-4| = 4.
    // The distance from -2 to 0 is |-2| = 2.
    // The distance from 1 to 0 is |1| = 1.
    // The distance from 4 to 0 is |4| = 4.
    // The distance from 8 to 0 is |8| = 8.
    // Thus, the closest number to 0 in the array is 1.
    fmt.Println(findClosestNumber([]int{-4,-2,1,4,8})) // 1
    // Example 2:
    // Input: nums = [2,-1,1]
    // Output: 1
    // Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.
    fmt.Println(findClosestNumber([]int{2,-1,1})) // 1
}