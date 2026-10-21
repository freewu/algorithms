package main

// 4087. Maximum Subarray Sum of Fixed Length After at Most K Swaps
// You are given an integer array nums and two integers length and k.

// You may perform at most k swap operations on nums.

// In one swap operation, you may choose two distinct indices i and j and swap nums[i] and nums[j].

// Return the maximum possible sum of a subarray of length exactly length after performing the swaps.

// Example 1:
// Input: nums = [1,-1,0,2], length = 2, k = 1
// Output: 3
// Explanation:
// - Swap the elements at indices 1 and 3, resulting in [1, 2, 0, -1].
// - The subarray [1, 2] has length 2 and sum 3.
// - No subarray of length 2 can have a larger sum, since 1 and 2 are the two largest elements.

// Example 2:
// Input: nums = [4,3,2,4], length = 3, k = 2
// Output: 11
// Explanation:
// - Swap the elements at indices 2 and 3, resulting in [4, 3, 4, 2].
// - The subarray [4, 3, 4] has length 3 and sum 11.
// - This is the sum of the three largest elements, so no larger answer is possible. Only one swap is needed.

// Example 3:
// Input: nums = [-1,-2], length = 2, k = 0
// Output: -3
// Explanation:
// No swaps are allowed. The only subarray of length 2 is [-1, -2], whose sum is -3.

// Constraints:
//     - 1 <= nums.length <= 300
//     - -10^{9} <= nums[i] <= 10^{9}
//     - 1 <= length <= nums.length
//     - 0 <= k <= nums.length

import "fmt"
import "sort"

//贪心
func maxSum(nums []int, length int, k int) int64 {
    res, n := int64(-1 << 61), len(nums)
    // 枚举所有连续窗口 [l, l+length-1]
    for l := 0; l+length <= n; l++ {
        inside, outside := []int{}, []int{}
        for i := 0; i < n; i++ {
            if i >= l && i < l+length {
                inside = append(inside, nums[i])
            } else {
                outside = append(outside, nums[i])
            }
        }
        // 窗口内升序，从小到大
        sort.Ints(inside)
        // 窗口外降序，从大到小
        sort.Slice(outside, func(i, j int) bool {
            return outside[i] > outside[j]
        })
        sum := int64(0)
        for _, v := range inside {
            sum += int64(v)
        }
        // 最多能交换的次数：k、窗口外元素数量，取小值
        limit := min(k, len(outside))
        for t := 0; t < limit; t++ {
            if outside[t] <= inside[t] {
                break
            }
            sum += int64(outside[t] - inside[t])
        }
        if sum > res {
            res = sum
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,0,2], length = 2, k = 1
    // Output: 3
    // Explanation:
    // - Swap the elements at indices 1 and 3, resulting in [1, 2, 0, -1].
    // - The subarray [1, 2] has length 2 and sum 3.
    // - No subarray of length 2 can have a larger sum, since 1 and 2 are the two largest elements.
    fmt.Println(maxSum([]int{1, -1, 0, 2}, 2, 1)) // 3
    // Example 2:
    // Input: nums = [4,3,2,4], length = 3, k = 2
    // Output: 11
    // Explanation:
    // - Swap the elements at indices 2 and 3, resulting in [4, 3, 4, 2].
    // - The subarray [4, 3, 4] has length 3 and sum 11.
    // - This is the sum of the three largest elements, so no larger answer is possible. Only one swap is needed.
    fmt.Println(maxSum([]int{4, 3, 2, 4}, 3, 2)) // 11
    // Example 3:
    // Input: nums = [-1,-2], length = 2, k = 0
    // Output: -3
    // Explanation:
    // No swaps are allowed. The only subarray of length 2 is [-1, -2], whose sum is -3.
    fmt.Println(maxSum([]int{-1, -2}, 2, 0)) // -3

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // 17
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // 17
}