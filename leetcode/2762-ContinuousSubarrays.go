package main

// 2762. Continuous Subarrays
// You are given a 0-indexed integer array nums. A subarray of nums is called continuous if:
//     Let i, i + 1, ..., j be the indices in the subarray. 
//     Then, for each pair of indices i <= i1, i2 <= j, 0 <= |nums[i1] - nums[i2]| <= 2.

// Return the total number of continuous subarrays.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [5,4,2,4]
// Output: 8
// Explanation: 
// Continuous subarray of size 1: [5], [4], [2], [4].
// Continuous subarray of size 2: [5,4], [4,2], [2,4].
// Continuous subarray of size 3: [4,2,4].
// Thereare no subarrys of size 4.
// Total continuous subarrays = 4 + 3 + 1 = 8.
// It can be shown that there are no more continuous subarrays.

// Example 2:
// Input: nums = [1,2,3]
// Output: 6
// Explanation: 
// Continuous subarray of size 1: [1], [2], [3].
// Continuous subarray of size 2: [1,2], [2,3].
// Continuous subarray of size 3: [1,2,3].
// Total continuous subarrays = 3 + 2 + 1 = 6.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func continuousSubarrays(nums []int) int64 {
    res, left := 0, 0
    count := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right, v := range nums {
        count[v]++
        for {
            mx, mn := v, v
            for i := range count {
                mx, mn = max(mx, i), min(mn, i)
            }
            if mx - mn <= 2 { break }
            y := nums[left]
            if count[y]--; count[y] == 0 {
                delete(count, y)
            }
            left++
        }
        res += (right - left + 1)
    }
    return int64(res)
}

func continuousSubarrays1(nums []int) int64 {
    res, mx, mn := 0, []int{}, []int{}
    for l, r := 0, 0; r < len(nums); r++ {
        for len(mx) > 0 && nums[r] >= nums[mx[len(mx) - 1]] {
            mx = mx[:len(mx) - 1]
        }
        mx = append(mx, r)
        for len(mn) > 0 && nums[r] <= nums[mn[len(mn) - 1]] {
            mn = mn[:len(mn) - 1]
        }
        mn = append(mn, r)
        for nums[mx[0]] - nums[mn[0]] > 2 {
            if l == mx[0] { mx = mx[1:] }
            if l == mn[0] { mn = mn[1:] }
            l++
        }
        res += (r - l + 1)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [5,4,2,4]
    // Output: 8
    // Explanation: 
    // Continuous subarray of size 1: [5], [4], [2], [4].
    // Continuous subarray of size 2: [5,4], [4,2], [2,4].
    // Continuous subarray of size 3: [4,2,4].
    // Thereare no subarrys of size 4.
    // Total continuous subarrays = 4 + 3 + 1 = 8.
    // It can be shown that there are no more continuous subarrays.
    fmt.Println(continuousSubarrays([]int{5,4,2,4})) // 8
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: 6
    // Explanation: 
    // Continuous subarray of size 1: [1], [2], [3].
    // Continuous subarray of size 2: [1,2], [2,3].
    // Continuous subarray of size 3: [1,2,3].
    // Total continuous subarrays = 3 + 2 + 1 = 6.
    fmt.Println(continuousSubarrays([]int{1,2,3})) // 6

    fmt.Println(continuousSubarrays1([]int{5,4,2,4})) // 8
    fmt.Println(continuousSubarrays1([]int{1,2,3})) // 6
}