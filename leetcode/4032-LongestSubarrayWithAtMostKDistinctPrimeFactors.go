package main

// 4032. Longest Subarray With at Most K Distinct Prime Factors
// You are given an integer array nums consisting of positive integers and an integer k.

// The prime factor set of a subarray is the union of the distinct prime factors of all its elements.

// Return the length of the longest subarray whose prime factor set contains at most k distinct prime factors. 
// If no such subarray exists, return 0.

// Example 1:
// Input: nums = [7,6,10,12,11], k = 3
// Output: 3
// Explanation:
// Consider the subarray [6, 10, 12]:
// The distinct prime factors of 6 are {2, 3}.
// The distinct prime factors of 10 are {2, 5}.
// The distinct prime factors of 12 are {2, 3}.
// The union of these sets is {2, 3, 5}, which contains 3 distinct prime factors.
// No longer subarray satisfies the condition. Therefore, the answer is 3.

// Example 2:
// Input: nums = [4,6,9,18], k = 4
// Output: 4
// Explanation:
// Consider the entire array [4, 6, 9, 18]:
// The distinct prime factors of 4 are {2}.
// The distinct prime factors of 6 are {2, 3}.
// The distinct prime factors of 9 are {3}.
// The distinct prime factors of 18 are {2, 3}.
// The union of these sets is {2, 3}, which contains 2 distinct prime factors.
// Since 2 <= 4, the entire array is valid. Therefore, the answer is 4.

// Example 3:
// Input: nums = [6,10,15], k = 2
// Output: 1
// Explanation:
// Every subarray of length at least 2 has prime factor set {2, 3, 5}, which contains 3 distinct prime factors.
// Since 3 > 2, only subarrays of length 1 are valid. Therefore, the answer is 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     2 <= nums[i] <= 10^5
//     1 <= k <= 10^4

import "fmt"

const MX = 100_001
var prime = [MX][]int{}

func init() {
    for i := 2; i < MX; i++ {
        if prime[i] == nil { // i 是质数
            for j := i; j < MX; j += i { // i 的倍数 j 有质因子 i
                prime[j] = append(prime[j], i)
            }
        }
    }
}

func longestSubarray(nums []int, k int) int {
    res, left, count := 0, 0, make(map[int]int)
    for i, v := range nums {
        for _, p := range prime[v] {
            count[p]++
        }
        for len(count) > k {
            for _, p := range prime[nums[left]] {
                if count[p] > 1 {
                    count[p]--
                } else {
                    delete(count, p) // 保证 len(cnt) 是窗口内的不同质因子个数
                }
            }
            left++
        }
        res = max(res, i-left+1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [7,6,10,12,11], k = 3
    // Output: 3
    // Explanation:
    // Consider the subarray [6, 10, 12]:
    // The distinct prime factors of 6 are {2, 3}.
    // The distinct prime factors of 10 are {2, 5}.
    // The distinct prime factors of 12 are {2, 3}.
    // The union of these sets is {2, 3, 5}, which contains 3 distinct prime factors.
    // No longer subarray satisfies the condition. Therefore, the answer is 3.
    fmt.Println(longestSubarray([]int{7,6,10,12,11}, 3)) // 3
    // Example 2:
    // Input: nums = [4,6,9,18], k = 4
    // Output: 4
    // Explanation:
    // Consider the entire array [4, 6, 9, 18]:
    // The distinct prime factors of 4 are {2}.
    // The distinct prime factors of 6 are {2, 3}.
    // The distinct prime factors of 9 are {3}.
    // The distinct prime factors of 18 are {2, 3}.
    // The union of these sets is {2, 3}, which contains 2 distinct prime factors.
    // Since 2 <= 4, the entire array is valid. Therefore, the answer is 4.
    fmt.Println(longestSubarray([]int{4,6,9,18}, 4)) // 4
    // Example 3:
    // Input: nums = [6,10,15], k = 2
    // Output: 1
    // Explanation:
    // Every subarray of length at least 2 has prime factor set {2, 3, 5}, which contains 3 distinct prime factors.
    // Since 3 > 2, only subarrays of length 1 are valid. Therefore, the answer is 1.
    fmt.Println(longestSubarray([]int{6,10,15}, 2)) // 1

    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9}, 2)) // 4
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1}, 2)) // 4
}