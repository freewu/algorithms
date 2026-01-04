package main

// 3795. Minimum Subarray Length With Distinct Sum At Least K
// You are given an integer array nums and an integer k.

// Return the minimum length of a subarray whose sum of the distinct values present in that subarray (each value counted once) is at least k. 
// If no such subarray exists, return -1.
 
// Example 1:
// Input: nums = [2,2,3,1], k = 4
// Output: 2
// Explanation:
// The subarray [2, 3] has distinct elements {2, 3} whose sum is 2 + 3 = 5, which is ​​​​​​​at least k = 4. Thus, the answer is 2.

// Example 2:
// Input: nums = [3,2,3,4], k = 5
// Output: 2
// Explanation:
// The subarray [3, 2] has distinct elements {3, 2} whose sum is 3 + 2 = 5, which is ​​​​​​​at least k = 5. Thus, the answer is 2.

// Example 3:
// Input: nums = [5,5,4], k = 5
// Output: 1
// Explanation:
// The subarray [5] has distinct elements {5} whose sum is 5, which is at least k = 5. Thus, the answer is 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= 10^9

import "fmt"

func minLength(nums []int, k int) int {
    mp := make(map[int]int)
    res, sum, left := 1 << 31, 0, 0
    for i, v := range nums {
        mp[v]++ // 1. 入
        if mp[v] == 1 {
            sum += v
        }
        for sum >= k {
            res = min(res, i - left + 1)  // 2. 更新答案
            out := nums[left]
            mp[out]-- // 3. 出
            if mp[out] == 0 {
                sum -= out
            }
            left++
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func minLength1(nums []int, k int) int {
    mp := make([]int, 100_001)
    res, count, i :=  1 << 31, 0, 0
    for j, n := range nums {
        if mp[n] == 0 {
            count += n
        }
        mp[n]++
        for ; count >= k && i <= j; i++ {
            res = min(res, j-i+1)
            old := nums[i]
            if mp[old] == 1 {
                count -= nums[i]
            }
            mp[nums[i]]--
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3,1], k = 4
    // Output: 2
    // Explanation:
    // The subarray [2, 3] has distinct elements {2, 3} whose sum is 2 + 3 = 5, which is ​​​​​​​at least k = 4. Thus, the answer is 2.
    fmt.Println(minLength([]int{2,2,3,1}, 4)) // 2
    // Example 2:
    // Input: nums = [3,2,3,4], k = 5
    // Output: 2
    // Explanation:
    // The subarray [3, 2] has distinct elements {3, 2} whose sum is 3 + 2 = 5, which is ​​​​​​​at least k = 5. Thus, the answer is 2.
    fmt.Println(minLength([]int{3,2,3,4}, 5)) // 2
    // Example 3:
    // Input: nums = [5,5,4], k = 5
    // Output: 1
    // Explanation:
    // The subarray [5] has distinct elements {5} whose sum is 5, which is at least k = 5. Thus, the answer is 1.
    fmt.Println(minLength([]int{5,5,4}, 5)) // 1

    fmt.Println(minLength([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(minLength([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1

    fmt.Println(minLength1([]int{2,2,3,1}, 4)) // 2
    fmt.Println(minLength1([]int{3,2,3,4}, 5)) // 2
    fmt.Println(minLength1([]int{5,5,4}, 5)) // 1
    fmt.Println(minLength1([]int{1,2,3,4,5,6,7,8,9}, 4)) // 1
    fmt.Println(minLength1([]int{9,8,7,6,5,4,3,2,1}, 4)) // 1
}