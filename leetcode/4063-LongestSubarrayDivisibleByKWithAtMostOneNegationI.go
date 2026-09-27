package main

// 4063. Longest Subarray Divisible by K with At Most One Negation I
// You are given an integer array nums and an integer k.

// A subarray is valid if its sum is divisible by k, or can become divisible by k by negating one element within that subarray.

// Negating an element means replacing its value x with -x.

// Return the length of the longest valid subarray. 
// If no valid subarray exists, return 0.

// A subarray is a contiguous, non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [4,1,2], k = 3
// Output: 3
// Explanation:
// The sum of the entire array is 7, and 7 % 3 = 1, so it is not divisible by k = 3.
// Negating nums[2] = 2 changes the sum to 4 + 1 − 2 = 3, which is divisible by k.
// Therefore, the entire array is a valid subarray, giving a length of 3.

// Example 2:
// Input: nums = [5,3,4], k = 7
// Output: 2
// Explanation:
// The sum of the entire array is 12, and negating any one of its elements does not make its sum divisible by 7.
// However, the subarray [3, 4] has a sum of 7, which is divisible by k = 7 without any negation.
// Therefore, the longest valid subarray has a length of 2.

// Example 3:
// Input: nums = [2,2,5], k = 6
// Output: 2
// Explanation:
// The sum of the entire array is 9, and negating any one of its elements does not make its sum divisible by 6.
// The subarray [2, 2] has a sum of 4. Negating either element changes it to [-2, 2] or [2, -2], both of which have a sum of 0.
// Therefore, the longest valid subarray has a length of 2.

// Constraints:
//     1 <= nums.length <= 1000
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= 10^5

import "fmt"

func longestSubarray(nums []int, k int) int {
    // 可被 K 整除的子数组
    longestSubarrayDivByK := func(nums []int, k int) int {
        pos := map[int]int{0: -1} // 前缀和 % k 首次出现的下标
        res, sum := 0, 0 // 前缀和
        for r, x := range nums {
            sum = (sum + x%k + k) % k // 保证 sum 非负
            l, ok := pos[sum]
            if ok {
                res = max(res, r-l)
            } else {
                pos[sum] = r
            }
        }
        return res
    }
    // 不取反
    res := longestSubarrayDivByK(nums, k)
    // 枚举取反元素
    for i := range nums {
        nums[i] *= -1 // 取反
        res = max(res, longestSubarrayDivByK(nums, k))
        nums[i] *= -1 // 复原
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,1,2], k = 3
    // Output: 3
    // Explanation:
    // The sum of the entire array is 7, and 7 % 3 = 1, so it is not divisible by k = 3.
    // Negating nums[2] = 2 changes the sum to 4 + 1 − 2 = 3, which is divisible by k.
    // Therefore, the entire array is a valid subarray, giving a length of 3.
    fmt.Println(longestSubarray([]int{4,1,2}, 3)) // 3 
    // Example 2:
    // Input: nums = [5,3,4], k = 7
    // Output: 2
    // Explanation:
    // The sum of the entire array is 12, and negating any one of its elements does not make its sum divisible by 7.
    // However, the subarray [3, 4] has a sum of 7, which is divisible by k = 7 without any negation.
    // Therefore, the longest valid subarray has a length of 2.
    fmt.Println(longestSubarray([]int{5,3,4}, 7)) // 2
    // Example 3:
    // Input: nums = [2,2,5], k = 6
    // Output: 2
    // Explanation:
    // The sum of the entire array is 9, and negating any one of its elements does not make its sum divisible by 6.
    // The subarray [2, 2] has a sum of 4. Negating either element changes it to [-2, 2] or [2, -2], both of which have a sum of 0.
    // Therefore, the longest valid subarray has a length of 2.
    fmt.Println(longestSubarray([]int{2,2,5}, 6)) // 2

    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9}, 3)) // 9 
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1}, 3)) // 9
}