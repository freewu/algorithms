package main

// 4047. Minimum Operations to Make XOR of All Elements Zero
// You are given an integer array nums consisting of positive integers.

// You may perform the following operation any number of times:
//     1. Choose two distinct indices i and j such that nums[i] != nums[j], 
//        and replace either nums[i] or nums[j] with nums[i] ^ nums[j], where ^ denotes the bitwise XOR.

// Return the minimum number of operations required to make the bitwise XOR of all elements in nums equal to 0. 
// If it is impossible, return -1.

// Example 1:
// Input: nums = [8,1,4,8,2]
// Output: 3
// Explanation:
// One optimal sequence of operations is:
// Choose indices 0 and 1, and replace nums[0] with 8 ^ 1 = 9. The array becomes [9, 1, 4, 8, 2].
// Choose indices 2 and 3, and replace nums[3] with 4 ^ 8 = 12. The array becomes [9, 1, 4, 12, 2].
// Choose indices 0 and 4, and replace nums[0] with 9 ^ 2 = 11. The array becomes [11, 1, 4, 12, 2].
// The XOR of all elements of nums is 11 ^ 1 ^ 4 ^ 12 ^ 2 = 0, so the answer is 3.

// Example 2:
// Input: nums = [1,2,3]
// Output: 0
// Explanation:
// The XOR of all elements of nums is 1 ^ 2 ^ 3 = 0, so no operations are required.

// Example 3:
// Input: nums = [1,2,4]
// Output: -1
// Explanation:
// It is impossible to make the XOR of all elements of nums equal to 0, so the answer is -1.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 2000

import "fmt"

func minOperations(nums []int) int {
    count := make(map[int]int)
    mx := 0
    for _, v := range nums {
        count[v]++
        if v > mx {
            mx = v
        }
    }
    // compute mask size
    bitLen, tmp := 0, mx
    for tmp > 0 {
        tmp >>= 1
        bitLen++
    }
    n, inf := 1 << bitLen, 1 << 61
    // dp[s]: minimal distinct numbers to XOR to mask
    s := make([]int, n)
    for i := range s {
        s[i] = inf
    }
    s[0] = 0
    // only one unique number and count odd, impossible
    if len(count) == 1 {
        for _, v := range count {
            if v % 2 == 1 {
                return -1
            }
        }
    }
    ms := 0
    for v, freq := range count {
        if freq % 2 == 1 {
            ms ^= v
        }
        ps := make([]int, n)
        copy(ps, s)
        for t := 1; t < n; t++ {
            if s[t^v] != inf {
                if ps[t] > s[t^v]+1 {
                    ps[t] = s[t^v] + 1
                }
            }
        }
        s = ps
    }
    res := s[ms]
    if res >= len(nums) {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [8,1,4,8,2]
    // Output: 3
    // Explanation:
    // One optimal sequence of operations is:
    // Choose indices 0 and 1, and replace nums[0] with 8 ^ 1 = 9. The array becomes [9, 1, 4, 8, 2].
    // Choose indices 2 and 3, and replace nums[3] with 4 ^ 8 = 12. The array becomes [9, 1, 4, 12, 2].
    // Choose indices 0 and 4, and replace nums[0] with 9 ^ 2 = 11. The array becomes [11, 1, 4, 12, 2].
    // The XOR of all elements of nums is 11 ^ 1 ^ 4 ^ 12 ^ 2 = 0, so the answer is 3.
    fmt.Println(minOperations([]int{8,1,4,8,2})) // 3
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation:
    // The XOR of all elements of nums is 1 ^ 2 ^ 3 = 0, so no operations are required.
    fmt.Println(minOperations([]int{1,2,3})) // 0
    // Example 3:
    // Input: nums = [1,2,4]
    // Output: -1
    // Explanation:
    // It is impossible to make the XOR of all elements of nums equal to 0, so the answer is -1.
    fmt.Println(minOperations([]int{1,2,4})) // -1
    // Example 4:
    // Input: nums = [15,4]
    // Output: -1
    fmt.Println(minOperations([]int{15,4})) // -1
    // Example 5:
    // Input: nums = [7,7,7]
    // Output: -1
    fmt.Println(minOperations([]int{7,7,7})) // -1
    // Example 6:
    // Input: nums = [6,2,7,9]
    // Output: -1
    fmt.Println(minOperations([]int{6,2,7,9})) // -1
    // Example 7:
    // Input: nums = [8,12,11,4]
    // Output: 1
    fmt.Println(minOperations([]int{8,12,11,4})) // 1
    // Example 8:
    // Input: nums = [6,15,15,15]
    // Output: 2
    fmt.Println(minOperations([]int{6,15,15,15})) // 2
    // Example 9:
    // Input: nums = [5,2,9,7,6]
    // Output: 2
    fmt.Println(minOperations([]int{5,2,9,7,6})) // 2

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 1
}