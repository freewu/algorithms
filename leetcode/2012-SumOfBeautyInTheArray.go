package main

// 2012. Sum of Beauty in the Array
// You are given a 0-indexed integer array nums. 
// For each index i (1 <= i <= nums.length - 2) the beauty of nums[i] equals:
//     2, if nums[j] < nums[i] < nums[k], for all 0 <= j < i and for all i < k <= nums.length - 1.
//     1, if nums[i - 1] < nums[i] < nums[i + 1], and the previous condition is not satisfied.
//     0, if none of the previous conditions holds.

// Return the sum of beauty of all nums[i] where 1 <= i <= nums.length - 2.

// Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation: For each index i in the range 1 <= i <= 1:
// - The beauty of nums[1] equals 2.

// Example 2:
// Input: nums = [2,4,6,4]
// Output: 1
// Explanation: For each index i in the range 1 <= i <= 2:
// - The beauty of nums[1] equals 1.
// - The beauty of nums[2] equals 0.

// Example 3:
// Input: nums = [3,2,1]
// Output: 0
// Explanation: For each index i in the range 1 <= i <= 1:
// - The beauty of nums[1] equals 0.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func sumOfBeauties(nums []int) int {
    res, n := 0, len(nums)
    rmin := make([]int, n + 1)
    rmin[n] = 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        rmin[i] = min(rmin[i + 1], nums[i])
    }
    for lmax, i := nums[0], 1; i < n - 1; i++ {
        if lmax < nums[i] && nums[i] < rmin[i+1] { // 2, if nums[j] < nums[i] < nums[k], for all 0 <= j < i and for all i < k <= nums.length - 1.
            res += 2
        } else if nums[i-1] < nums[i] && nums[i] < nums[i + 1] { // 1, if nums[i - 1] < nums[i] < nums[i + 1], and the previous condition is not satisfied.
            res++
        }
        lmax = max(lmax, nums[i])
    }
    return res
}

func sumOfBeauties1(nums []int) int {
    res, cur, n := 0, 0, len(nums)
    pre := make([]bool, n)
    for i, v := range nums {
        if i > 0 && i < n - 1 && v > nums[i - 1] && v < nums[i + 1] {
            res++
        }
        if cur < v {
            pre[i], cur = true, v
        }
    }
    cur = nums[n-1]
    for i := n - 1; i > 0; i-- {
        if nums[i] < cur {
            cur = nums[i]
            if pre[i] {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 2
    // Explanation: For each index i in the range 1 <= i <= 1:
    // - The beauty of nums[1] equals 2.
    fmt.Println(sumOfBeauties([]int{1,2,3})) // 2
    // Example 2:
    // Input: nums = [2,4,6,4]
    // Output: 1
    // Explanation: For each index i in the range 1 <= i <= 2:
    // - The beauty of nums[1] equals 1.
    // - The beauty of nums[2] equals 0.
    fmt.Println(sumOfBeauties([]int{2,4,6,4})) // 1
    // Example 3:
    // Input: nums = [3,2,1]
    // Output: 0
    // Explanation: For each index i in the range 1 <= i <= 1:
    // - The beauty of nums[1] equals 0.
    fmt.Println(sumOfBeauties([]int{3,2,1})) // 0

    fmt.Println(sumOfBeauties1([]int{1,2,3})) // 2
    fmt.Println(sumOfBeauties1([]int{2,4,6,4})) // 1
    fmt.Println(sumOfBeauties1([]int{3,2,1})) // 0
}