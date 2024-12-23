package main

// 3299. Sum of Consecutive Subsequences
// We call an array arr of length n consecutive if one of the following holds:
//     arr[i] - arr[i - 1] == 1 for all 1 <= i < n.
//     arr[i] - arr[i - 1] == -1 for all 1 <= i < n.

// The value of an array is the sum of its elements.

// For example, [3, 4, 5] is a consecutive array of value 12 and [9, 8] is another of value 17. 
// While [3, 4, 3] and [8, 6] are not consecutive.

// Given an array of integers nums, return the sum of the values of all consecutive non-empty subsequences.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that an array of length 1 is also considered consecutive.

// Example 1:
// Input: nums = [1,2]
// Output: 6
// Explanation:
// The consecutive subsequences are: [1], [2], [1, 2].

// Example 2:
// Input: nums = [1,4,2,3]
// Output: 31
// Explanation:
// The consecutive subsequences are: [1], [4], [2], [3], [1, 2], [2, 3], [4, 3], [1, 2, 3].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func getSum(nums []int) int {
    const mod = 1_000_000_007
    calc := func(nums []int) int64 {
        res, n := int64(0), len(nums)
        left, right, count := make([]int64, n), make([]int64, n), make(map[int]int64)
        for i := 1; i < n; i++ {
            count[nums[i-1]] += 1 + count[nums[i-1]-1]
            left[i] = count[nums[i]-1]
        }
        count = make(map[int]int64)
        for i := n - 2; i >= 0; i-- {
            count[nums[i+1]] += 1 + count[nums[i+1]+1]
            right[i] = count[nums[i]+1]
        }
        for i, v := range nums {
            res = (res + (left[i] + right[i] + (left[i] * right[i] % mod)) * int64(v) % mod) % mod
        }
        return res
    }
    x := calc(nums)
    for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
    y := calc(nums)
    sum := int64(0)
    for _, v := range nums {
        sum += int64(v)
    }
    return int((x + y + sum) % mod)
}

func main() {
    // Example 1:
    // Input: nums = [1,2]
    // Output: 6
    // Explanation:
    // The consecutive subsequences are: [1], [2], [1, 2].
    fmt.Println(getSum([]int{1,2})) // 6
    // Example 2:
    // Input: nums = [1,4,2,3]
    // Output: 31
    // Explanation:
    // The consecutive subsequences are: [1], [4], [2], [3], [1, 2], [2, 3], [4, 3], [1, 2, 3].
    fmt.Println(getSum([]int{1,4,2,3})) // 31
}