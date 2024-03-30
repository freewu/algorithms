package main

// 992. Subarrays with K Different Integers
// Given an integer array nums and an integer k, return the number of good subarrays of nums.
// A good array is an array where the number of different integers in that array is exactly k.
//     For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,2,1,2,3], k = 2
// Output: 7
// Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]

// Example 2:
// Input: nums = [1,2,1,3,4], k = 3
// Output: 3
// Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
 
// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i], k <= nums.length

import "fmt"

// 双指针
func subarraysWithKDistinct(nums []int, k int) int {
    most := func (nums []int, k int) int {
        left, right, l := 0, 0, len(nums)
        m := make(map[int]int,0)
        res := 0
        for right < l {
            m[nums[right]]++
            for left < l && len(m) > k {
                m[nums[left]]--
                if m[nums[left]] == 0{
                    delete(m,nums[left])
                }
                left++
            }
            res += (right - left)
            right++
        }
        return res
    }
    return most(nums, k) - most(nums, k-1)
}

func subarraysWithKDistinct1(nums []int, k int) int {
    f := func(k int) []int {
        n := len(nums)
        pos := make([]int, n)
        cnt := make([]int, n + 1)
        s, j := 0, 0
        for i, x := range nums {
            cnt[x]++
            if cnt[x] == 1 {
                s++
            }
            for ; s > k; j++ {
                cnt[nums[j]]--
                if cnt[nums[j]] == 0 {
                    s--
                }
            }
            pos[i] = j
        }
        return pos
    }
    res := 0
    left, right := f(k), f(k-1)
    for i := range left {
        res += right[i] - left[i]
    }
    return res
}

func main() {
    // Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
    fmt.Println(subarraysWithKDistinct([]int{1,2,1,2,3}, 2)) // 7
    // Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
    fmt.Println(subarraysWithKDistinct([]int{1,2,1,3,4}, 3)) // 3

    fmt.Println(subarraysWithKDistinct1([]int{1,2,1,2,3}, 2)) // 7
    fmt.Println(subarraysWithKDistinct1([]int{1,2,1,3,4}, 3)) // 3
}