package main

// 862. Shortest Subarray with Sum at Least K
// Given an integer array nums and an integer k, 
// return the length of the shortest non-empty subarray of nums with a sum of at least k.  
// If there is no such subarray, return -1.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1], k = 1
// Output: 1

// Example 2:
// Input: nums = [1,2], k = 4
// Output: -1

// Example 3:
// Input: nums = [2,-1,2], k = 3
// Output: 3

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= 10^9

import "fmt"

func shortestSubarray(nums []int, k int) int {
    d, n := [][2]int{{0, -1}}, len(nums) // d - monotonic deque of [a[0]+...+a[i], i]
    res := n + 1
    for i, s := 0, 0; i < n; i++ {
        s += nums[i]
        if nums[i] > 0 {
            if s - d[0][0] < k {
                d = append(d, [2]int{s, i})
            } else {
                // try to reduce len
                j := 0
                for ; j < len(d) && s - d[j][0] >= k; j++ {}
                // update minimum
                if i - d[j-1][1] < res {
                    res = i - d[j-1][1]
                }
                // reduce d, because used elements are not needed
                d = append(d[j:], [2]int{s, i})
            }
        } else {
            // remove elements >= s from d
            j := len(d) - 2
            for ; j >= 0 && d[j][0] >= s; j-- {}
            d = append(d[:j+1], [2]int{s, i})
        }
    }
    if res == n + 1 {
        return -1
    }
    return res
}

func shortestSubarray1(nums []int, k int) int {
    n := len(nums)
    prefix, res, queue := make([]int, n+1), n + 1, []int{}
    for i, num := range nums {
        prefix[i+1] = prefix[i] + num
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, sum := range prefix {
        for len(queue) > 0 && sum - prefix[queue[0]] >= k {
            res = min(res, i - queue[0])
            queue = queue[1:] // pop
        }
        for len(queue) > 0 && prefix[queue[len(queue) - 1]] >= sum {
            queue = queue[:len(queue) - 1]
        }
        queue = append(queue, i) // push
    }
    if res < n + 1 {
        return res
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1], k = 1
    // Output: 1
    fmt.Println(shortestSubarray([]int{1}, 1)) // 1
    // Example 2:
    // Input: nums = [1,2], k = 4
    // Output: -1
    fmt.Println(shortestSubarray([]int{1,2}, 4)) // -1
    // Example 3:
    // Input: nums = [2,-1,2], k = 3
    // Output: 3
    fmt.Println(shortestSubarray([]int{2,-1,2}, 3)) // 3

    fmt.Println(shortestSubarray1([]int{1}, 1)) // 1
    fmt.Println(shortestSubarray1([]int{1,2}, 4)) // -1
    fmt.Println(shortestSubarray1([]int{2,-1,2}, 3)) // 3
}