package main

// 473. Matchsticks to Square
// You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. 
// You want to use all the matchsticks to make one square. 
// You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

// Return true if you can make this square and false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/matchsticks1-grid.jpg" />
// Input: matchsticks = [1,1,2,2,2]
// Output: true
// Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.

// Example 2:
// Input: matchsticks = [3,3,3,3,4]
// Output: false
// Explanation: You cannot find a way to form a square with all the matchsticks.
 
// Constraints:
//     1 <= matchsticks.length <= 15
//     1 <= matchsticks[i] <= 10^8

import "fmt"

func makesquare(matchsticks []int) bool {
    sum, mx, k, vis := 0, 0, 4, make([]bool, len(matchsticks))
    for _, n := range matchsticks {
        sum += n
        if n > mx {
            mx = n
        }
    }
    if sum % k != 0 || mx > sum / k {
        return false
    }
    var backtracking func(nums []int, vis *[]bool, k, targetSubsetSum, curSubsetSum, nextIndexToCheck int) bool
    backtracking = func(nums []int, vis *[]bool, k, targetSubsetSum, curSubsetSum, nextIndexToCheck int) bool {
        if k == 0 {
            return true
        }
        if curSubsetSum == targetSubsetSum {
            return backtracking(nums, vis, k-1, targetSubsetSum, 0, 0)
        }
        for i := nextIndexToCheck; i < len(nums); i++ {
            if !(*vis)[i] && curSubsetSum+nums[i] <= targetSubsetSum {
                (*vis)[i] = true
                if backtracking(nums, vis, k, targetSubsetSum, curSubsetSum+nums[i], i+1) {
                    return true
                }
                (*vis)[i] = false
            }
        }
        return false
    }
    return backtracking(matchsticks, &vis, k, sum/k, 0, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/matchsticks1-grid.jpg" />
    // Input: matchsticks = [1,1,2,2,2]
    // Output: true
    // Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.
    fmt.Println(makesquare([]int{1,1,2,2,2})) // true
    // Example 2:
    // Input: matchsticks = [3,3,3,3,4]
    // Output: false
    // Explanation: You cannot find a way to form a square with all the matchsticks.
    fmt.Println(makesquare([]int{3,3,3,3,4})) // true
}