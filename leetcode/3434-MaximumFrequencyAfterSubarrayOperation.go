package main

// 3434. Maximum Frequency After Subarray Operation
// You are given an array nums of length n. 
// You are also given an integer k.

// You perform the following operation on nums once:
//     1. Select a subarray nums[i..j] where 0 <= i <= j <= n - 1.
//     2. Select an integer x and add x to all the elements in nums[i..j].

// Find the maximum frequency of the value k after the operation.

// Example 1:
// Input: nums = [1,2,3,4,5,6], k = 1
// Output: 2
// Explanation:
// After adding -5 to nums[2..5], 1 has a frequency of 2 in [1, 2, -2, -1, 0, 1].

// Example 2:
// Input: nums = [10,2,3,4,5,5,4,3,2,2], k = 10
// Output: 4
// Explanation:
// After adding 8 to nums[1..9], 10 has a frequency of 4 in [10, 10, 11, 12, 13, 13, 12, 11, 10, 10].

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 50
//     1 <= k <= 50

import "fmt"

func maxFrequency(nums []int, k int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    kadane := func(num int) int {
        res, count := 0, 0
        for _, v := range nums {
            if v == k    { count-- }
            if v == num  { count++ }
            if count < 0 { count = 0 }
            res = max(res, count)
        }
        return res
    }
    res := 0
    for i := range mp {
        res = max(res, kadane(i))
    }
    return res + mp[k]
}

func maxFrequency1(nums []int, k int) int {
    res, count := 0, 0
    arr := [51]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        arr[v] = max(arr[v], count) + 1
        if v == k {
            count++
            res++
        }
        res = max(res, arr[v])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6], k = 1
    // Output: 2
    // Explanation:
    // After adding -5 to nums[2..5], 1 has a frequency of 2 in [1, 2, -2, -1, 0, 1].
    fmt.Println(maxFrequency([]int{1,2,3,4,5,6}, 1)) // 2
    // Example 2:
    // Input: nums = [10,2,3,4,5,5,4,3,2,2], k = 10
    // Output: 4
    // Explanation:
    // After adding 8 to nums[1..9], 10 has a frequency of 4 in [10, 10, 11, 12, 13, 13, 12, 11, 10, 10].
    fmt.Println(maxFrequency([]int{10,2,3,4,5,5,4,3,2,2}, 10)) // 4

    fmt.Println(maxFrequency([]int{1,2,3,4,5,6,7,8,9}, 1)) // 2
    fmt.Println(maxFrequency([]int{9,8,7,6,5,4,3,2,1}, 1)) // 2

    fmt.Println(maxFrequency1([]int{1,2,3,4,5,6}, 1)) // 2
    fmt.Println(maxFrequency1([]int{10,2,3,4,5,5,4,3,2,2}, 10)) // 4
    fmt.Println(maxFrequency1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 2
    fmt.Println(maxFrequency1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 2
}