package main

// 1755. Closest Subsequence Sum
// You are given an integer array nums and an integer goal.

// You want to choose a subsequence of nums such that the sum of its elements is the closest possible to goal. 
// That is, if the sum of the subsequence's elements is sum, then you want to minimize the absolute difference abs(sum - goal).

// Return the minimum possible value of abs(sum - goal).

// Note that a subsequence of an array is an array formed by removing some elements (possibly all or none) of the original array.

// Example 1:
// Input: nums = [5,-7,3,5], goal = 6
// Output: 0
// Explanation: Choose the whole array as a subsequence, with a sum of 6.
// This is equal to the goal, so the absolute difference is 0.

// Example 2:
// Input: nums = [7,-9,15,-2], goal = -5
// Output: 1
// Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
// The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.

// Example 3:
// Input: nums = [1,2,3], goal = -7
// Output: 7

// Constraints:
//     1 <= nums.length <= 40
//     -10^7 <= nums[i] <= 10^7
//     -10^9 <= goal <= 10^9

import "fmt"
import "sort"
import "math/bits"

func minAbsDifference(nums []int, goal int) int {
    res, half := 1 << 31, len(nums) / 2
    a, b := nums[0:half], nums[half:] // Split nums array in two halves
    asums, bsums := make(map[int]bool), make(map[int]bool) // Use hash like a set to remove duplicates
    var dfs func(i int, arr []int, sums map[int]bool, cur int) // Generate up to 2^n possible unique sums given n elements in arr
    dfs = func(i int, arr []int, sums map[int]bool, cur int) {
        if i == len(arr) {
            sums[cur] = true
            return
        }
        dfs(i + 1, arr, sums, cur)
        dfs(i + 1, arr, sums, cur + arr[i])
    }
    flatten := func(hash map[int]bool) []int {
        res, i := make([]int, len(hash)), 0
        for k := range hash {
            res[i] = k
            i++
        }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    dfs(0, a, asums, 0)
    dfs(0, b, bsums, 0)
    arr, brr := flatten(asums), flatten(bsums)
    sort.Ints(brr) // Sort brr so we can bisect in it later
    for _, v := range arr {  // For each x in arr, search for y in brr so x + y is closest to goal
        y := goal - v
        i := sort.SearchInts(brr, y)
        if i < len(brr) { // If y is greater than the max in brr, i would be equal to len(brr).
            res = min(res, abs(y - brr[i]))
        }
        if i > 0 {  // We must have the follow-up if i > 0 clause to handle this case.
            res = min(res, abs(y - brr[i-1]))
        }
    }
    return res
}

var ltsums, rtsums [1<<20]int

func minAbsDifference1(nums []int, goal int) int {
    n := len(nums)
    ltsums, rtsums := ltsums[:1<<(n/2)], rtsums[:1<<(n-n/2)]
    sumside := func(sums []int, k, offset int) {
        sums[0] = 0
        for i := 1; i < 1<<k; i++ {
            sums[i] = sums[i&(i-1)] + nums[offset + bits.TrailingZeros(uint(i))]
        }
        sort.Ints(sums)
    }
    sumside(ltsums, n/2, 0)
    sumside(rtsums, n-n/2, n/2)
    rtsums = append(rtsums, 1<<40)
    res, rt := 1 << 31, 1<<(n-n/2)-1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for lt := 0; lt < 1<<(n/2); lt++ {
        for rt > 0 && ltsums[lt] + rtsums[rt] > goal { rt-- }
        res = min(res, min(abs(goal - (ltsums[lt]+rtsums[rt])), abs(goal - (ltsums[lt]+rtsums[rt+1])) ))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,-7,3,5], goal = 6
    // Output: 0
    // Explanation: Choose the whole array as a subsequence, with a sum of 6.
    // This is equal to the goal, so the absolute difference is 0.
    fmt.Println(minAbsDifference([]int{5,-7,3,5}, 6)) // 0
    // Example 2:
    // Input: nums = [7,-9,15,-2], goal = -5
    // Output: 1
    // Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
    // The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.
    fmt.Println(minAbsDifference([]int{7,-9,15,-2}, -5)) // 1
    // Example 3:
    // Input: nums = [1,2,3], goal = -7
    // Output: 7
    fmt.Println(minAbsDifference([]int{1,2,3}, -7)) // 7
    fmt.Println(minAbsDifference([]int{1556913,-259675,-7667451,-4380629,-4643857,-1436369,7695949,-4357992,-842512,-118463}, -9681425)) // 10327

    fmt.Println(minAbsDifference1([]int{5,-7,3,5}, 6)) // 0
    fmt.Println(minAbsDifference1([]int{7,-9,15,-2}, -5)) // 1
    fmt.Println(minAbsDifference1([]int{1,2,3}, -7)) // 7
    fmt.Println(minAbsDifference1([]int{1556913,-259675,-7667451,-4380629,-4643857,-1436369,7695949,-4357992,-842512,-118463}, -9681425)) // 10327

}