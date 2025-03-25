package main

// 2524. Maximum Frequency Score of a Subarray
// You are given an integer array nums and a positive integer k.

// The frequency score of an array is the sum of the distinct values in the array raised to the power of their frequencies, taking the sum modulo 109 + 7.
//     For example, the frequency score of the array [5,4,5,7,4,4] is (43 + 52 + 71) modulo (10^9 + 7) = 96.

// Return the maximum frequency score of a subarray of size k in nums. 
// You should maximize the value under the modulo and not the actual value.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,1,1,2,1,2], k = 3
// Output: 5
// Explanation: The subarray [2,1,2] has a frequency score equal to 5. It can be shown that it is the maximum frequency score we can have.

// Example 2:
// Input: nums = [1,1,1,1,1,1], k = 4
// Output: 1
// Explanation: All the subarrays of length 4 have a frequency score equal to 1.

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

func maxFrequencyScore(nums []int, k int) (ans int) {
    res, score, mod := 0, 0, 1_000_000_007
    mp := make(map[int][]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        stack, ok := mp[v]  // 每个 nums[i] 创建一个栈，每遇到一个 v 就把 v 与栈顶的乘积入栈
        if !ok {
            score += v
            mp[v] = []int{v}
        } else {
            last := stack[len(stack) - 1]
            cur := last * v % mod
            score += cur - last
            mp[v] = append(stack, cur)
        }
        if i >= k-1 {
            res = max(res, (score % mod + mod) % mod)
            v = nums[i - k + 1]
            stack = mp[v]
            score -= stack[len(stack) - 1]
            if len(stack) == 1 {
                delete(mp, v)
            } else {
                score += stack[len(stack) - 2]
                mp[v] = stack[:len(stack) - 1]
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,2,1,2], k = 3
    // Output: 5
    // Explanation: The subarray [2,1,2] has a frequency score equal to 5. It can be shown that it is the maximum frequency score we can have.
    fmt.Println(maxFrequencyScore([]int{1,1,1,2,1,2}, 3)) // 5
    // Example 2:
    // Input: nums = [1,1,1,1,1,1], k = 4
    // Output: 1
    // Explanation: All the subarrays of length 4 have a frequency score equal to 1.
    fmt.Println(maxFrequencyScore([]int{1,1,1,1,1,1}, 4)) // 1

    fmt.Println(maxFrequencyScore([]int{1,2,3,4,5,6,7,8,9}, 4)) // 30
    fmt.Println(maxFrequencyScore([]int{9,8,7,6,5,4,3,2,1}, 4)) // 30
}