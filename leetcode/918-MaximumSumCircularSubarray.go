package main

// 918. Maximum Sum Circular Subarray
// Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.

// A circular array means the end of the array connects to the beginning of the array. 
// Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].

// A subarray may only include each element of the fixed buffer nums at most once. 
// Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.

// Example 1:
// Input: nums = [1,-2,3,-2]
// Output: 3
// Explanation: Subarray [3] has maximum sum 3.

// Example 2:
// Input: nums = [5,-3,5]
// Output: 10
// Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.

// Example 3:
// Input: nums = [-3,-2,-3]
// Output: -2
// Explanation: Subarray [-2] has maximum sum -2.
 
// Constraints:
//     n == nums.length
//     1 <= n <= 3 * 10^4
//     -3 * 10^4 <= nums[i] <= 3 * 10^4

import "fmt"

func maxSubarraySumCircular(nums []int) int {
    sum, cmx, cmn, gmx, gmn := nums[0], nums[0], nums[0], nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        cmx = max(cmx + nums[i], nums[i])
        cmn = min(cmn + nums[i], nums[i])
        gmx = max(gmx, cmx)
        gmn = min(gmn, cmn)
        sum += nums[i]
    }
    if gmx < 0 {
        return gmx
    } else {
        return max(gmx, sum - gmn)
    }
}

func maxSubarraySumCircular1(nums []int) int {
    max := func (nums []int, s int, e int) int {
        l, mx, ending := len(nums), nums[s], nums[s]
        if e < s {
            e = e + l
        }
        for i := s + 1; i <= e; i++ {
            k := i % l
            if nums[k] > nums[k]+ending {
                ending = nums[k]
            } else {
                ending = nums[k] + ending
            }
            if mx < ending {
                mx = ending
            }
        }
        return mx
    }
    min := func(nums []int) (int, int, int) {
        mn, ending, left, right := nums[0], nums[0], 0, 0
        for i := 1; i < len(nums); i++ {
            if nums[i] < nums[i]+ending {
                ending = nums[i]
                left = i
            } else {
                ending = nums[i] + ending
            }
            if mn >= ending {
                mn = ending
                right = i
            }
        }
        if left > right {
            left = right
        }
        return mn, left, right
    }
    _, left, right := min(nums)
    l := len(nums)
    fullMax := max(nums, 0, l-1)
    gapMax := fullMax
    if left == 0 {
        if right == l-1 {
            gapMax = max(nums, left, right)
        } else {
            gapMax = max(nums, right+1, l-1)
        }
    } else {
        if right == l-1 {
            gapMax = max(nums, 0, left-1)
        } else {
            gapMax = max(nums, right+1, left-1)
        }
    }
    if fullMax > gapMax {
        return fullMax
    }
    return gapMax
}

func main() {
    // Example 1:
    // Input: nums = [1,-2,3,-2]
    // Output: 3
    // Explanation: Subarray [3] has maximum sum 3.
    fmt.Println(maxSubarraySumCircular([]int{1,-2,3,-2})) // 3
    // Example 2:
    // Input: nums = [5,-3,5]
    // Output: 10
    // Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
    fmt.Println(maxSubarraySumCircular([]int{5,-3,5})) // 10
    // Example 3:
    // Input: nums = [-3,-2,-3]
    // Output: -2
    // Explanation: Subarray [-2] has maximum sum -2.
    fmt.Println(maxSubarraySumCircular([]int{-3,-2,-3})) // -2

    fmt.Println(maxSubarraySumCircular1([]int{1,-2,3,-2})) // 3
    fmt.Println(maxSubarraySumCircular1([]int{5,-3,5})) // 10
    fmt.Println(maxSubarraySumCircular1([]int{-3,-2,-3})) // -2
}