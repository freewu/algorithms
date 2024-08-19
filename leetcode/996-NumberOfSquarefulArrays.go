package main

// 996. Number of Squareful Arrays
// An array is squareful if the sum of every pair of adjacent elements is a perfect square.
// Given an integer array nums, return the number of permutations of nums that are squareful.
// Two permutations perm1 and perm2 are different if there is some index i such that perm1[i] != perm2[i].

// Example 1:
// Input: nums = [1,17,8]
// Output: 2
// Explanation: [1,8,17] and [17,8,1] are the valid permutations.

// Example 2:
// Input: nums = [2,2,2]
// Output: 1

// Constraints:
//     1 <= nums.length <= 12
//     0 <= nums[i] <= 10^9

import "fmt"
import "math"

func numSquarefulPerms(nums []int) int {
    isSquare := func (n int) bool {
        f := float64(n)
        x := int(math.Sqrt(f))
        if x * x != n { return false }
        return true
    }
    var permutate func(nums []int, k int) int
    permutate = func(nums []int, k int) int {
        if k >= len(nums) {
            return 1
        }
        count, set := 0, map[int]bool{}
        for i := k; i < len(nums); i++ {
            if b := set[nums[i]]; b {
                continue
            }
            set[nums[i]] = true
            nums[i], nums[k] = nums[k], nums[i]
            if k == 0 || isSquare(nums[k] + nums[k-1]) {
                count += permutate(nums, k+1)
            }
            nums[i], nums[k] = nums[k], nums[i]
        }
        return count
    }
    return permutate(nums, 0)
}

func main() {
    // Example 1:
    // Input: nums = [1,17,8]
    // Output: 2
    // Explanation: [1,8,17] and [17,8,1] are the valid permutations.
    fmt.Println(numSquarefulPerms([]int{1,17,8})) // 2
    // Example 2:
    // Input: nums = [2,2,2]
    // Output: 1
    fmt.Println(numSquarefulPerms([]int{2,2,2})) // 1
}