package main

// 3097. Shortest Subarray With OR at Least K II
// You are given an array nums of non-negative integers and an integer k.

// An array is called special if the bitwise OR of all of its elements is at least k.

// Return the length of the shortest special non-empty subarray of nums, or return -1 if no special subarray exists.

// Example 1:
// Input: nums = [1,2,3], k = 2
// Output: 1
// Explanation:
// The subarray [3] has OR value of 3. Hence, we return 1.

// Example 2:
// Input: nums = [2,1,8], k = 10
// Output: 3
// Explanation:
// The subarray [2,1,8] has OR value of 11. Hence, we return 3.

// Example 3:
// Input: nums = [1,2], k = 0
// Output: 1
// Explanation:
// The subarray [1] has OR value of 1. Hence, we return 1.

// Constraints:
//     1 <= nums.length <= 2 * 10^5
//     0 <= nums[i] <= 10^9
//     0 <= k <= 10^9

import "fmt"
import "math"

func minimumSubarrayLength(nums []int, k int) int {
    res, n := 1 << 31, len(nums)
    bits := make([]int, 32)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    addBits := func(num int, bits []int) {
        i := 0
        for i < 32 {
            bits[i] += (num & 1)
            num >>= 1
            i++
        }
    }
    removeBits := func(num int, bits []int) {
        i := 0
        for i < 32 {
            bits[i] -= (num & 1)
            num >>= 1
            i++
        }
    }
    getValue := func(bits []int) int {
        res := 0
        for i := 0; i < 32; i++ {
            if bits[i] != 0 {
                res += int(math.Pow(float64(2), float64(i)))
            }
        }
        return res
    }
    for left, right := 0, 0; right < n; right++ {
        addBits(nums[right], bits)
        for left <= right && getValue(bits) >= k {
            res = min(res, right - left + 1)
            removeBits(nums[left], bits)
            left++
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func minimumSubarrayLength1(nums []int, k int) int {
    res := 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    helper := func(count []int) int {
        res := 0
        for i, v := range count {
            if v > 0 {
                res += 1 << i
            }
        }
        return res
    }
    for i, start, count := 0, 0, make([]int, 32); i < len(nums); i++ {
        for j := range count {
            if 1 << j & nums[i] != 0 {
                count[j]++
            }
        }
        for ; helper(count) >= k && start <= i; start++ {
            res = min(res, i - start + 1)
            for j := range count {
                if 1 << j & nums[start] != 0 {
                    count[j]--
                }
            }
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2
    // Output: 1
    // Explanation:
    // The subarray [3] has OR value of 3. Hence, we return 1.
    fmt.Println(minimumSubarrayLength([]int{1,2,3}, 2)) // 1
    // Example 2:
    // Input: nums = [2,1,8], k = 10
    // Output: 3
    // Explanation:
    // The subarray [2,1,8] has OR value of 11. Hence, we return 3.
    fmt.Println(minimumSubarrayLength([]int{2,1,8}, 10)) // 3
    // Example 3:
    // Input: nums = [1,2], k = 0
    // Output: 1
    // Explanation:
    // The subarray [1] has OR value of 1. Hence, we return 1.
    fmt.Println(minimumSubarrayLength([]int{1,2}, 0)) // 1

    fmt.Println(minimumSubarrayLength1([]int{1,2,3}, 2)) // 1
    fmt.Println(minimumSubarrayLength1([]int{2,1,8}, 10)) // 3
    fmt.Println(minimumSubarrayLength1([]int{1,2}, 0)) // 1
}