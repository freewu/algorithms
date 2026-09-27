package main

// 4064. Longest Subarray Divisible by K with At Most One Negation II
// You are given an integer array nums and an integer k.

// A subarray is valid if its sum is divisible by k, or can become divisible by k by negating one element within that subarray.

// Negating an element means replacing its value x with -x.

// Return the length of the longest valid subarray. If no valid subarray exists, return 0.

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
//     1 <= nums.length <= 10^5​​​​​​​
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= 3000​​​​​​​

import "fmt"
import "sort"

func longestSubarray(nums []int, k int) int {
    numPos := make([][]int, k) // 2*nums[i] % k 出现的所有位置
    firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
    for i := range firstPos {
        firstPos[i] = -1
    }
    lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
    lastPos[0], firstPos[0] = 0, 0
    res, sum := 0, 0 // 前缀和
    for i, x := range nums {
        x = x%k + k // 保证 x 非负
        y := x * 2 % k
        numPos[y] = append(numPos[y], i)
        sum = (sum + x) % k
        r := i + 1
        l := firstPos[sum]
        if l < 0 {
            firstPos[sum] = r
        } else {
            // 不取反时的最大长度
            res = max(res, r-l)
        }
        lastPos[sum] = r
    }
    // 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
    for sl, l := range firstPos {
        if l < 0 {
            continue
        }
        for sr, r := range lastPos {
            if r-l <= res { // 最优性优化：res 无法增大
                continue
            }
            // 2*nums[i]%k = (sr-sl)%k
            pos := numPos[(sr-sl+k)%k] // +k 保证结果非负
            index := sort.SearchInts(pos, l)
            if index < len(pos) && pos[index] < r {
                res = r - l
            }
        }
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