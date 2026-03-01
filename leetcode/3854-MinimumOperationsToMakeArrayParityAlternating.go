package main

// 3854. Minimum Operations to Make Array Parity Alternating
// You are given an integer array nums.

// An array is called parity alternating if for every index i where 0 <= i < n - 1, nums[i] and nums[i + 1] have different parity (one is even and the other is odd).

// In one operation, you may choose any index i and either increase nums[i] by 1 or decrease nums[i] by 1.

// Return an integer array answer of length 2 where:
//     1. answer[0] is the minimum number of operations required to make the array parity alternating.
//     2. answer[1] is the minimum possible value of max(nums) - min(nums) taken over all arrays that are parity alternating and can be obtained by performing exactly answer[0] operations.

// An array of length 1 is considered parity alternating.

// Example 1:
// Input: nums = [-2,-3,1,4]
// Output: [2,6]
// Explanation:
// Applying the following operations:
// Increase nums[2] by 1, resulting in nums = [-2, -3, 2, 4].
// Decrease nums[3] by 1, resulting in nums = [-2, -3, 2, 3].
// The resulting array is parity alternating, and the value of max(nums) - min(nums) = 3 - (-3) = 6 is the minimum possible among all parity alternating arrays obtainable using exactly 2 operations.

// Example 2:
// Input: nums = [0,2,-2]
// Output: [1,3]
// Explanation:
// Applying the following operation:
// Decrease nums[1] by 1, resulting in nums = [0, 1, -2].
// The resulting array is parity alternating, and the value of max(nums) - min(nums) = 1 - (-2) = 3 is the minimum possible among all parity alternating arrays obtainable using exactly 1 operation.

// Example 3:
// Input: nums = [7]
// Output: [0,0]
// Explanation:
// No operations are required. The array is already parity alternating, and the value of max(nums) - min(nums) = 7 - 7 = 0, which is the minimum possible.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func makeParityAlternating(nums []int) []int {
    const MX = 1000000001
    mx, mn,mx0, mn0, mx1, mn1, n := nums[0], nums[0],  -MX, MX, -MX, MX, len(nums)
    count0, count1 :=  0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        mx, mn = max(mx, nums[i]), min(mn, nums[i])
    }
    for i := 0; i < n; i++ {
        if (nums[i] % 2 + 2) % 2 == i % 2 {
            mx0, mn0 = max(mx0, nums[i]), min(mn0, nums[i])
            if nums[i] == mn {
                mx1,mn1 = max(mx1, mn + 1), min(mn1, mn + 1)
            } else if nums[i] == mx {
                mx1,mn1 = max(mx1, mx - 1), min(mn1, mx - 1)
            }
            count1++
        } else{
            mx1,mn1 = max(mx1, nums[i]), min(mn1, nums[i])
            if nums[i] == mn {
                mx0, mn0 = max(mx0, mn + 1), min(mn0, mn + 1)
            } else if nums[i] == mx {
                mx0, mn0 = max(mx0, mx - 1), min(mn0, mx - 1)
            }
            count0++
        }
    }
    if count0 == count1 {
        return []int{ count0, min(mx0 - mn0, mx1 - mn1) }
    } 
    if count0 > count1 {
        return []int{ count1, mx1 - mn1 }
    }
    return []int{ count0, mx0 - mn0 }
}

func main() {
    // Example 1:
    // Input: nums = [-2,-3,1,4]
    // Output: [2,6]
    // Explanation:
    // Applying the following operations:
    // Increase nums[2] by 1, resulting in nums = [-2, -3, 2, 4].
    // Decrease nums[3] by 1, resulting in nums = [-2, -3, 2, 3].
    // The resulting array is parity alternating, and the value of max(nums) - min(nums) = 3 - (-3) = 6 is the minimum possible among all parity alternating arrays obtainable using exactly 2 operations.
    fmt.Println(makeParityAlternating([]int{-2,-3,1,4})) // [2,6]
    // Example 2:
    // Input: nums = [0,2,-2]
    // Output: [1,3]
    // Explanation:
    // Applying the following operation:
    // Decrease nums[1] by 1, resulting in nums = [0, 1, -2].
    // The resulting array is parity alternating, and the value of max(nums) - min(nums) = 1 - (-2) = 3 is the minimum possible among all parity alternating arrays obtainable using exactly 1 operation.
    fmt.Println(makeParityAlternating([]int{0,2,-2})) // [1,3]
    // Example 3:
    // Input: nums = [7]
    // Output: [0,0]
    // Explanation:
    // No operations are required. The array is already parity alternating, and the value of max(nums) - min(nums) = 7 - 7 = 0, which is the minimum possible.
    fmt.Println(makeParityAlternating([]int{7})) // [0,0]

    fmt.Println(makeParityAlternating([]int{1,2,3,4,5,6,7,8,9})) // [0 8]
    fmt.Println(makeParityAlternating([]int{9,8,7,6,5,4,3,2,1})) // [0 8]
}