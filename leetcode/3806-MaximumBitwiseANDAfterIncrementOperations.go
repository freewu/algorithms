package main

// 3806. Maximum Bitwise AND After Increment Operations
// You are given an integer array nums and two integers k and m.

// You may perform at most k operations. In one operation, you may choose any index i and increase nums[i] by 1.

// Return an integer denoting the maximum possible bitwise AND of any subset of size m after performing up to k operations optimally.

// Example 1:
// Input: nums = [3,1,2], k = 8, m = 2
// Output: 6
// Explanation:
// We need a subset of size m = 2. Choose indices [0, 2].
// Increase nums[0] = 3 to 6 using 3 operations, and increase nums[2] = 2 to 6 using 4 operations.
// The total number of operations used is 7, which is not greater than k = 8.
// The two chosen values become [6, 6], and their bitwise AND is 6, which is the maximum possible.

// Example 2:
// Input: nums = [1,2,8,4], k = 7, m = 3
// Output: 4
// Explanation:
// We need a subset of size m = 3. Choose indices [0, 1, 3].
// Increase nums[0] = 1 to 4 using 3 operations, increase nums[1] = 2 to 4 using 2 operations, and keep nums[3] = 4.
// The total number of operations used is 5, which is not greater than k = 7.
// The three chosen values become [4, 4, 4], and their bitwise AND is 4, which is the maximum possible.​​​​​​​

// Example 3:
// Input: nums = [1,1], k = 3, m = 2
// Output: 2
// Explanation:
// We need a subset of size m = 2. Choose indices [0, 1].
// Increase both values from 1 to 2 using 1 operation each.
// The total number of operations used is 2, which is not greater than k = 3.
// The two chosen values become [2, 2], and their bitwise AND is 2, which is the maximum possible.
 
// Constraints:
//     1 <= n == nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9
//     1 <= m <= n

import "fmt"
import "slices"
import "math/bits"

func maximumAND(nums []int, k int, m int) int {
    ops := make([]int, len(nums)) // Number of operations for each number
    res, maxWidth := 0, bits.Len(uint(slices.Max(nums) + k))
    for bit := maxWidth - 1; bit >= 0; bit-- {
        target := res | 1 << bit // Note: target includes the bits already set in ans
        for i, v := range nums {
            j := bits.Len(uint(target &^ v))
            // j - 1 is the highest bit where target is 1 and v is 0
            mask := 1 << j - 1
            ops[i] = target & mask - v & mask
        }
        // Greedy: pick the smallest m operation counts
        slices.Sort(ops)
        sum := 0
        for _, v := range ops[:m] {
            sum += v
        }
        if sum <= k {
            res = target // This bit of the answer can be set to 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2], k = 8, m = 2
    // Output: 6
    // Explanation:
    // We need a subset of size m = 2. Choose indices [0, 2].
    // Increase nums[0] = 3 to 6 using 3 operations, and increase nums[2] = 2 to 6 using 4 operations.
    // The total number of operations used is 7, which is not greater than k = 8.
    // The two chosen values become [6, 6], and their bitwise AND is 6, which is the maximum possible.
    fmt.Println(maximumAND([]int{3,1,2}, 8, 2)) // 6
    // Example 2:
    // Input: nums = [1,2,8,4], k = 7, m = 3
    // Output: 4
    // Explanation:
    // We need a subset of size m = 3. Choose indices [0, 1, 3].
    // Increase nums[0] = 1 to 4 using 3 operations, increase nums[1] = 2 to 4 using 2 operations, and keep nums[3] = 4.
    // The total number of operations used is 5, which is not greater than k = 7.
    // The three chosen values become [4, 4, 4], and their bitwise AND is 4, which is the maximum possible.​​​​​​​
    fmt.Println(maximumAND([]int{1,2,8,4}, 7, 3)) // 4
    // Example 3:
    // Input: nums = [1,1], k = 3, m = 2
    // Output: 2
    // Explanation:
    // We need a subset of size m = 2. Choose indices [0, 1].
    // Increase both values from 1 to 2 using 1 operation each.
    // The total number of operations used is 2, which is not greater than k = 3.
    // The two chosen values become [2, 2], and their bitwise AND is 2, which is the maximum possible.
    fmt.Println(maximumAND([]int{1,1}, 3, 2)) // 2

    fmt.Println(maximumAND([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 10
    fmt.Println(maximumAND([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 10
}