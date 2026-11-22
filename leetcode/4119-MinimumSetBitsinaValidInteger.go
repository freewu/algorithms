package main

// 4119. Minimum Set Bits in a Valid Integer
// You are given an integer array nums of length n and an integer k.

// Let x_{1}, x_{2}, ..., x_{n-k+1} denote the bitwise AND of each subarray of length k in nums, in order from left to right.

// A positive integer y is called valid if (y & x_{i}) != 0 for every 1 <= i <= n - k + 1.

// Return the minimum number of set bits in any valid integer y. 
// If no valid integer exists, return -1.

// Example 1:
// Input: nums = [6,7,3], k = 2
// Output: 1
// Explanation:
//     The two subarrays of length k = 2 are [6, 7] and [7, 3].
//     Their bitwise ANDs are x_{1} = 6 & 7 = 6 and x_{2} = 7 & 3 = 3.
//     y = 2 is valid, since 2 & 6 = 2 and 2 & 3 = 2, both non-zero.
//     y = 2 has 1 set bit ("010"), so the answer is 1.

// Example 2:
// Input: nums = [3,5,12], k = 1
// Output: 2
// Explanation:
//     Since k = 1, each subarray has one element, so x_{1} = 3, x_{2} = 5, and x_{3} = 12.
//     y = 9 is valid, since 9 & 3 = 1, 9 & 5 = 1, and 9 & 12 = 8, all non-zero.
//     y = 9 has 2 set bits ("1001"), so the answer is 2.

// Example 3:
// Input: nums = [2,1], k = 2
// Output: -1
// Explanation:
//     The only subarray of length k = 2 is [2, 1], with bitwise AND x_{1} = 2 & 1 = 0.
//     Since x_{1} = 0, no y is valid. Thus, the answer is -1.

// Constraints:
//     1 <= n == nums.length <= 10^{5}
//     1 <= k <= n
//     0 <= nums[i] <= 10^{5}

import "fmt"
import "math/bits"
import "slices"

func minSetBits(nums []int, k int) int {
    w := bits.Len(uint(slices.Max(nums)))
    count, f := make([]int, w), make([]bool, 1 << w)
    for r, v := range nums {
        // 1. 入
        for ; v > 0; v &= v - 1 {
            count[bits.TrailingZeros(uint(v))]++
        }
        l := r + 1 - k
        if l < 0 { // 尚未形成第一个窗口
            continue
        }
        // 2. 计算 x
        x := 0
        for i, c := range count {
            if c == k {
                x |= 1 << i
            }
        }
        f[x] = true
        // 3. 出
        for v = nums[l]; v > 0; v &= v - 1 {
            count[bits.TrailingZeros(uint(v))]--
        }
    }
    for i := range w {
        for s := 0; s < 1<<w; s++ {
            s |= 1 << i // 优化：快速跳到 i 位是 1 的 s
            f[s] = f[s] || f[s^1<<i]
        }
    }
    mx := -1
    // 先计算 y 的补集的二进制 1 的个数的最大值
    for z, hasSubset := range f[ : 1 << w] {
        if !hasSubset {
            mx = max(mx, bits.OnesCount(uint(z)))
        }
    }
    if mx < 0 {
        return -1
    }
    // 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
    return w - mx
}

func main() {
    // Example 1:
    // Input: nums = [6,7,3], k = 2
    // Output: 1
    // Explanation:
    //     The two subarrays of length k = 2 are [6, 7] and [7, 3].
    //     Their bitwise ANDs are x_{1} = 6 & 7 = 6 and x_{2} = 7 & 3 = 3.
    //     y = 2 is valid, since 2 & 6 = 2 and 2 & 3 = 2, both non-zero.
    //     y = 2 has 1 set bit ("010"), so the answer is 1.
    fmt.Println(minSetBits([]int{6,7,3}, 2)) // 1
    // Example 2:
    // Input: nums = [3,5,12], k = 1
    // Output: 2
    // Explanation:
    //     Since k = 1, each subarray has one element, so x_{1} = 3, x_{2} = 5, and x_{3} = 12.
    //     y = 9 is valid, since 9 & 3 = 1, 9 & 5 = 1, and 9 & 12 = 8, all non-zero.
    //     y = 9 has 2 set bits ("1001"), so the answer is 2.
    fmt.Println(minSetBits([]int{3,5,12}, 1)) // 2
    // Example 3:
    // Input: nums = [2,1], k = 2
    // Output: -1
    // Explanation:
    //     The only subarray of length k = 2 is [2, 1], with bitwise AND x_{1} = 2 & 1 = 0.
    //     Since x_{1} = 0, no y is valid. Thus, the answer is -1.
    fmt.Println(minSetBits([]int{2,1}, 2)) // -1

    fmt.Println(minSetBits([]int{1,2,3,4,5,6,7,8,9}, 2)) // -1
    fmt.Println(minSetBits([]int{9,8,7,6,5,4,3,2,1}, 2)) // -1
}