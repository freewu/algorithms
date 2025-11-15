package main

// 2031. Count Subarrays With More Ones Than Zeros
// You are given a binary array nums containing only the integers 0 and 1. 
// Return the number of subarrays in nums that have more 1's than 0's. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: nums = [0,1,1,0,1]
// Output: 9
// Explanation:
// The subarrays of size 1 that have more ones than zeros are: [1], [1], [1]
// The subarrays of size 2 that have more ones than zeros are: [1,1]
// The subarrays of size 3 that have more ones than zeros are: [0,1,1], [1,1,0], [1,0,1]
// The subarrays of size 4 that have more ones than zeros are: [1,1,0,1]
// The subarrays of size 5 that have more ones than zeros are: [0,1,1,0,1]

// Example 2:
// Input: nums = [0]
// Output: 0
// Explanation:
// No subarrays have more ones than zeros.

// Example 3:
// Input: nums = [1]
// Output: 1
// Explanation:
// The subarrays of size 1 that have more ones than zeros are: [1]

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 1

import "fmt"

// 哈希 + 前缀和
func subarraysWithMoreOnesThanZeroes(nums []int) int {
    res, sum, count := 0, 0, 0
    mp := make(map[int]int)
    mp[0] = 1
    for _, v := range nums {
        if v == 1 { // 判断出以当前元素结尾时，满足要求的子数组的数量
            count += mp[sum]
            sum++
        } else {
            count -= mp[sum-1]
            sum--
        }
        mp[sum]++ // 累计数量
        res += count // 新增的满足要求的数组数量
        res %= 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1,0,1]
    // Output: 9
    // Explanation:
    // The subarrays of size 1 that have more ones than zeros are: [1], [1], [1]
    // The subarrays of size 2 that have more ones than zeros are: [1,1]
    // The subarrays of size 3 that have more ones than zeros are: [0,1,1], [1,1,0], [1,0,1]
    // The subarrays of size 4 that have more ones than zeros are: [1,1,0,1]
    // The subarrays of size 5 that have more ones than zeros are: [0,1,1,0,1]
    fmt.Println(subarraysWithMoreOnesThanZeroes([]int{0,1,1,0,1})) // 9
    // Example 2:
    // Input: nums = [0]
    // Output: 0
    // Explanation:
    // No subarrays have more ones than zeros.
    fmt.Println(subarraysWithMoreOnesThanZeroes([]int{0})) // 0
    // Example 3:
    // Input: nums = [1]
    // Output: 1
    // Explanation:
    // The subarrays of size 1 that have more ones than zeros are: [1]
    fmt.Println(subarraysWithMoreOnesThanZeroes([]int{1})) // 1

    fmt.Println(subarraysWithMoreOnesThanZeroes([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(subarraysWithMoreOnesThanZeroes([]int{9,8,7,6,5,4,3,2,1})) // 1
}