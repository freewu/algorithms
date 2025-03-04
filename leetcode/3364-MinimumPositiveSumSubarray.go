package main

// 3364. Minimum Positive Sum Subarray
// You are given an integer array nums and two integers l and r. 
// Your task is to find the minimum sum of a subarray whose size is between l and r (inclusive) and whose sum is greater than 0.

// Return the minimum sum of such a subarray. If no such subarray exists, return -1.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [3, -2, 1, 4], l = 2, r = 3
// Output: 1
// Explanation:
// The subarrays of length between l = 2 and r = 3 where the sum is greater than 0 are:
// [3, -2] with a sum of 1
// [1, 4] with a sum of 5
// [3, -2, 1] with a sum of 2
// [-2, 1, 4] with a sum of 3
// Out of these, the subarray [3, -2] has a sum of 1, which is the smallest positive sum. Hence, the answer is 1.

// Example 2:
// Input: nums = [-2, 2, -3, 1], l = 2, r = 3
// Output: -1
// Explanation:
// There is no subarray of length between l and r that has a sum greater than 0. So, the answer is -1.

// Example 3:
// Input: nums = [1, 2, 3, 4], l = 2, r = 4
// Output: 3
// Explanation:
// The subarray [1, 2] has a length of 2 and the minimum sum greater than 0. So, the answer is 3.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= l <= r <= nums.length
//     -1000 <= nums[i] <= 1000

import "fmt"

func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    prefix:= make([]int, n + 1)
    for i, v := range nums { // calculate prefix sum
        prefix[i + 1] = prefix[i] + v
    }
    res := 1 << 31
    // iterate for all window size
    for i := l; i <= r; i++ {
        for j := 0; j + i <= n; j++ {
            sum := prefix[i + j] - prefix[j]
            if 0 < sum && sum < res {
                res = sum
            }
        }
    }
    if res == 1 << 31 {
        return -1
    }
    return res
}

func minimumSumSubarray1(nums []int, l int, r int) int {
    res, n := 1 << 31, len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n - l + 1; i++ {
        sum := 0
        for j := i; j < n && j -i + 1 <= r; j++ {
            sum += nums[j]
            if sum > 0 && j - i + 1 >= l {
                res = min(res, sum)
            }
        }
    }
    if res == 1 << 31 {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3, -2, 1, 4], l = 2, r = 3
    // Output: 1
    // Explanation:
    // The subarrays of length between l = 2 and r = 3 where the sum is greater than 0 are:
    // [3, -2] with a sum of 1
    // [1, 4] with a sum of 5
    // [3, -2, 1] with a sum of 2
    // [-2, 1, 4] with a sum of 3
    // Out of these, the subarray [3, -2] has a sum of 1, which is the smallest positive sum. Hence, the answer is 1.
    fmt.Println(minimumSumSubarray([]int{3, -2, 1, 4}, 2, 3)) // 1
    // Example 2:
    // Input: nums = [-2, 2, -3, 1], l = 2, r = 3
    // Output: -1
    // Explanation:
    // There is no subarray of length between l and r that has a sum greater than 0. So, the answer is -1.
    fmt.Println(minimumSumSubarray([]int{-2, 2, -3, 1}, 2, 3)) // -1
    // Example 3:
    // Input: nums = [1, 2, 3, 4], l = 2, r = 4
    // Output: 3
    // Explanation:
    // The subarray [1, 2] has a length of 2 and the minimum sum greater than 0. So, the answer is 3.
    fmt.Println(minimumSumSubarray([]int{1, 2, 3, 4}, 2, 4)) // 3

    fmt.Println(minimumSumSubarray([]int{1,2,3,4,5,6,7,8,9}, 2, 4)) // 3
    fmt.Println(minimumSumSubarray([]int{9,8,7,6,5,4,3,2,1}, 2, 4)) // 3

    fmt.Println(minimumSumSubarray1([]int{3, -2, 1, 4}, 2, 3)) // 1
    fmt.Println(minimumSumSubarray1([]int{-2, 2, -3, 1}, 2, 3)) // -1
    fmt.Println(minimumSumSubarray1([]int{1, 2, 3, 4}, 2, 4)) // 3
    fmt.Println(minimumSumSubarray1([]int{1,2,3,4,5,6,7,8,9}, 2, 4)) // 3
    fmt.Println(minimumSumSubarray1([]int{9,8,7,6,5,4,3,2,1}, 2, 4)) // 3
}