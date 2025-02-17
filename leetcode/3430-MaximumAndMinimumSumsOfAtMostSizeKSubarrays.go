package main

// 3430. Maximum and Minimum Sums of at Most Size K Subarrays
// You are given an integer array nums and a positive integer k. 
// Return the sum of the maximum and minimum elements of all subarrays with at most k elements.

// Example 1:
// Input: nums = [1,2,3], k = 2
// Output: 20
// Explanation:
// The subarrays of nums with at most 2 elements are:
// Subarray	Minimum	Maximum	Sum
// [1]             1       1   2
// [2]             2       2   4
// [3]             3       3   6
// [1, 2]          1       2   3
// [2, 3]          2       3   5
// Final Total         20
// The output would be 20.

// Example 2:
// Input: nums = [1,-3,1], k = 2
// Output: -6
// Explanation:
// The subarrays of nums with at most 2 elements are:
// Subarray	Minimum	Maximum	Sum
// [1]             1       1   2
// [-3]            -3      -3  -6
// [1]             1       1   2
// [1, -3]         -3      1   -2
// [-3, 1]	        -3      1   -2
// Final Total         -6
// The output would be -6.

// Constraints:
//     1 <= nums.length <= 80000
//     1 <= k <= nums.length
//     -10^6 <= nums[i] <= 10^6

import "fmt"
import "math"

func minMaxSubarraySum(nums []int, k int) int64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 计算最小值的贡献
    calc := func() int {
        res, stack := 0, []int{ -1 } // 哨兵
        for r, x := range nums {
            r0 := r
            for len(stack) > 1 && nums[stack[len(stack) - 1]] >= x {
                i := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]
                l := stack[len(stack) - 1]
                if r - l- 1 <= k {
                    res += nums[i] * ((i - l) * (r - i)) // 累加贡献
                } else {
                    l, r = max(l, i - k), min(r, i + k)
                    a := (r - i) * (i - (r - k)) // 左端点 > r-k 的子数组个数
                    b := (l + r + k - i*2 + 1) * (r - l - k) / 2 // 左端点 <= r-k 的子数组个数
                    res += nums[i] * (a + b) // 累加贡献
                }
            }
            stack = append(stack, r0)
        }
        return res
    }
    nums = append(nums, math.MinInt)
    res := calc()
    for i := range nums { // 所有元素取反（求最大值），就可以复用同一份代码了
        nums[i] = -nums[i]
    }
    res -= calc()
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2
    // Output: 20
    // Explanation:
    // The subarrays of nums with at most 2 elements are:
    // Subarray	Minimum	Maximum	Sum
    // [1]             1       1   2
    // [2]             2       2   4
    // [3]             3       3   6
    // [1, 2]          1       2   3
    // [2, 3]          2       3   5
    // Final Total         20
    // The output would be 20.
    fmt.Println(minMaxSubarraySum([]int{1,2,3}, 2)) // 20
    // Example 2:
    // Input: nums = [1,-3,1], k = 2
    // Output: -6
    // Explanation:
    // The subarrays of nums with at most 2 elements are:
    // Subarray	Minimum	Maximum	Sum
    // [1]             1       1   2
    // [-3]            -3      -3  -6
    // [1]             1       1   2
    // [1, -3]         -3      1   -2
    // [-3, 1]	        -3      1   -2
    // Final Total         -6
    // The output would be -6.
    fmt.Println(minMaxSubarraySum([]int{1,-3,1}, 2)) // -6

    fmt.Println(minMaxSubarraySum([]int{1,2,3,4,5,6,7,8,9}, 2)) // -6
    fmt.Println(minMaxSubarraySum([]int{9,8,7,6,5,4,3,2,1}, 2)) // -6
}