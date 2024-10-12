package main

// 2035. Partition Array Into Two Arrays to Minimize Sum Difference
// You are given an integer array nums of 2 * n integers. 
// You need to partition nums into two arrays of length n to minimize the absolute difference of the sums of the arrays. 
// To partition nums, put each element of nums into one of the two arrays.

// Return the minimum possible absolute difference.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/02/ex1.png" />
// Input: nums = [3,9,7,3]
// Output: 2
// Explanation: One optimal partition is: [3,9] and [7,3].
// The absolute difference between the sums of the arrays is abs((3 + 9) - (7 + 3)) = 2.

// Example 2:
// Input: nums = [-36,36]
// Output: 72
// Explanation: One optimal partition is: [-36] and [36].
// The absolute difference between the sums of the arrays is abs((-36) - (36)) = 72.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/10/02/ex3.png" />
// Input: nums = [2,-1,0,4,-2,-9]
// Output: 0
// Explanation: One optimal partition is: [2,4,-9] and [-1,0,-2].
// The absolute difference between the sums of the arrays is abs((2 + 4 + -9) - (-1 + 0 + -2)) = 0.
 
// Constraints:
//     1 <= n <= 15
//     nums.length == 2 * n
//     -10^7 <= nums[i] <= 10^7

import "fmt"
import "sort"

func minimumDifference(nums []int) int {
    sum, n := 0, len(nums) / 2
    left, right := nums[:n], nums[n:]
    leftSums, rightSums := make([][]int, n+1), make([][]int, n+1)
    for _, v := range nums {
        sum += v
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var collectPossibleSums func(i, n, nitems, currSum int, nums []int, output *[][]int)
    collectPossibleSums = func(i, n, nitems, currSum int, nums []int, output *[][]int) {
        if i == n {
            (*output)[nitems] = append((*output)[nitems], currSum)
            return
        }
        collectPossibleSums(i+1, n, nitems+1, currSum+nums[i], nums, output) // pick
        collectPossibleSums(i+1, n, nitems, currSum, nums, output)           // skip
    }
    collectPossibleSums(0, n, 0, 0, left, &leftSums)
    collectPossibleSums(0, n, 0, 0, right, &rightSums)
    for k := range leftSums {
        sort.Ints(leftSums[k])
    }
    res, halfSum := 1 << 31, sum / 2
    // pick k numbers from rightSums
    for k := range rightSums {
        for _, rightSum := range rightSums[k] {
            // and n-k numbers from the right side
            options := leftSums[n - k]
            if len(options) == 0 { continue }
            // Search for closest match
            // recall that searchInts finds the place where the given number would
            // be inserted. That position can be both before and after the optimal
            // match, and beyond the size of the slice.
            idx := sort.SearchInts(options, halfSum-rightSum)
            if idx != len(options) {
                res = min(res, 2*abs(rightSum+options[idx]-halfSum))
            }
            if idx != 0 {
                res = min(res, 2*abs(rightSum+options[idx-1]-halfSum))
            }
        }
    }
    if abs(sum) % 2 == 1 {
        res += 1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/02/ex1.png" />
    // Input: nums = [3,9,7,3]
    // Output: 2
    // Explanation: One optimal partition is: [3,9] and [7,3].
    // The absolute difference between the sums of the arrays is abs((3 + 9) - (7 + 3)) = 2.
    fmt.Println(minimumDifference([]int{3,9,7,3})) // 2
    // Example 2:
    // Input: nums = [-36,36]
    // Output: 72
    // Explanation: One optimal partition is: [-36] and [36].
    // The absolute difference between the sums of the arrays is abs((-36) - (36)) = 72.
    fmt.Println(minimumDifference([]int{-36,36})) // 72
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/10/02/ex3.png" />
    // Input: nums = [2,-1,0,4,-2,-9]
    // Output: 0
    // Explanation: One optimal partition is: [2,4,-9] and [-1,0,-2].
    // The absolute difference between the sums of the arrays is abs((2 + 4 + -9) - (-1 + 0 + -2)) = 0.
    fmt.Println(minimumDifference([]int{2,-1,0,4,-2,-9})) // 0
}