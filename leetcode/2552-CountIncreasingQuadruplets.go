package main

// 2552. Count Increasing Quadruplets
// Given a 0-indexed integer array nums of size n containing all numbers from 1 to n, 
// return the number of increasing quadruplets.

// A quadruplet (i, j, k, l) is increasing if:
//     0 <= i < j < k < l < n, and
//     nums[i] < nums[k] < nums[j] < nums[l].

// Example 1:
// Input: nums = [1,3,2,4,5]
// Output: 2
// Explanation: 
// - When i = 0, j = 1, k = 2, and l = 3, nums[i] < nums[k] < nums[j] < nums[l].
// - When i = 0, j = 1, k = 2, and l = 4, nums[i] < nums[k] < nums[j] < nums[l]. 
// There are no other quadruplets, so we return 2.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 0
// Explanation: There exists only one quadruplet with i = 0, j = 1, k = 2, l = 3, but since nums[j] < nums[k], we return 0.

// Constraints:
//     4 <= nums.length <= 4000
//     1 <= nums[i] <= nums.length
//     All the integers of nums are unique. nums is a permutation.

import "fmt"

// 暴力法 brute force 超出时间限制 70 / 121 
func countQuadruplets(nums []int) int64 {
    res, n := 0, len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                for l := k + 1; l < n; l++ {
                    if nums[i] < nums[k] && nums[k] < nums[j] && nums[j] < nums[l] { 
                        res++ 
                    } 
                }
            }
        }
    }
    return int64(res)
}

func countQuadruplets1(nums []int) int64 {
    res, n := 0, len(nums)
    countLarger, countSmaller := make([][]int, n), make([][]int, n)
    for i := range countLarger {
        countLarger[i] = make([]int, n)
        countSmaller[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        c := 0
        for j := i + 1; j < n; j++ {
            if nums[j] > nums[i] { c++ }
            countLarger[i][j] = c
        }
    }
    for i := n - 1; i >= 0; i-- {
        c := 0
        for j := i - 1; j >= 0; j-- {
            if nums[j] < nums[i] { c++ }
            countSmaller[j][i] = c
        }
    }    
    for i := 1; i < n - 2; i++ {
        for j := i + 1; j < n - 1; j++ {
            if (nums[j] > nums[i]) { continue }
            // find the count how many numbers smaller than nums[k] and index smaller than j, and how many numbers larger than nums[j] and index larger than k， than add the product of them into res
            res += (countSmaller[0][j] - countSmaller[i][j]) * (countLarger[i][n - 1] - countLarger[i][j])
        }
    }
    return int64(res)
}

func countQuadruplets2(nums []int) int64 {
    cnt4, cnt3 := 0, make([]int, len(nums))
    for i := 2; i < len(nums); i++ {
        cnt2 := 0
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] { // 3 < 4
                cnt4 += cnt3[j]
                // 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
                cnt2++
            } else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
                cnt3[j] += cnt2
            }
        }
    }
    return int64(cnt4)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,4,5]
    // Output: 2
    // Explanation: 
    // - When i = 0, j = 1, k = 2, and l = 3, nums[i] < nums[k] < nums[j] < nums[l].
    // - When i = 0, j = 1, k = 2, and l = 4, nums[i] < nums[k] < nums[j] < nums[l]. 
    // There are no other quadruplets, so we return 2.
    fmt.Println(countQuadruplets([]int{1,3,2,4,5})) // 2
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 0
    // Explanation: There exists only one quadruplet with i = 0, j = 1, k = 2, l = 3, but since nums[j] < nums[k], we return 0.
    fmt.Println(countQuadruplets([]int{1,2,3,4})) // 0

    fmt.Println(countQuadruplets1([]int{1,3,2,4,5})) // 2
    fmt.Println(countQuadruplets1([]int{1,2,3,4})) // 0

    fmt.Println(countQuadruplets2([]int{1,3,2,4,5})) // 2
    fmt.Println(countQuadruplets2([]int{1,2,3,4})) // 0
}