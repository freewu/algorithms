package main

// 2150. Find All Lonely Numbers in the Array
// You are given an integer array nums. A number x is lonely when it appears only once, 
// and no adjacent numbers (i.e. x + 1 and x - 1) appear in the array.

// Return all lonely numbers in nums. You may return the answer in any order.

// Example 1:
// Input: nums = [10,6,5,8]
// Output: [10,8]
// Explanation: 
// - 10 is a lonely number since it appears exactly once and 9 and 11 does not appear in nums.
// - 8 is a lonely number since it appears exactly once and 7 and 9 does not appear in nums.
// - 5 is not a lonely number since 6 appears in nums and vice versa.
// Hence, the lonely numbers in nums are [10, 8].
// Note that [8, 10] may also be returned.

// Example 2:
// Input: nums = [1,3,5,3]
// Output: [1,5]
// Explanation: 
// - 1 is a lonely number since it appears exactly once and 0 and 2 does not appear in nums.
// - 5 is a lonely number since it appears exactly once and 4 and 6 does not appear in nums.
// - 3 is not a lonely number since it appears twice.
// Hence, the lonely numbers in nums are [1, 5].
// Note that [5, 1] may also be returned.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^6

import "fmt"
import "sort"

func findLonely(nums []int) []int {
    res, mp := []int{}, make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    for k, v := range mp {
        if v == 1 && mp[k - 1] == 0 && mp[k + 1] == 0 { // 在数组中仅出现 一次 ，且没有 相邻 数字（即，x + 1 和 x - 1）出现在数组中
            res = append(res, k) // 则认为数字 x 是 孤独数字
        }
    }
    return res
}

func findLonely1(nums []int) []int {
    if len(nums) == 1 { return nums }
    res, mn, mx  := make([]int, 0), nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        mn, mx = min(mn, v), max(mx, v) 
    } 
    count := make([]int, 1 + mx) 
    for _, v := range nums {
        count[v]++
    }
    for i := mn; i <= mx; i++ {
        if count[i] == 1 {
            if i == 0 {
                if count[1] == 0 {
                    res = append(res, 0) 
                }
                continue
            } else if i == mx && count[i - 1] == 0 {
                res = append(res, mx)
                break
            }
            if count[i - 1] == 0 && count[i + 1] == 0 {
                res = append(res, i) 
            }
        }
    }
    return res 
}

func findLonely2(nums []int) []int {
    n := len(nums)
    sort.Ints(nums)
    res := make([]int, 0, n)
    if n == 1 || nums[0] + 1 < nums[1] {
        res = append(res, nums[0])
    }
    for i := 1; i < n-1; i++ {
        if nums[i]-1 > nums[i-1] && nums[i]+1 < nums[i+1] {
            res = append(res, nums[i])
        }
    }
    if n > 1 && nums[n-2]+1 < nums[n-1] {
        res = append(res, nums[n-1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,6,5,8]
    // Output: [10,8]
    // Explanation: 
    // - 10 is a lonely number since it appears exactly once and 9 and 11 does not appear in nums.
    // - 8 is a lonely number since it appears exactly once and 7 and 9 does not appear in nums.
    // - 5 is not a lonely number since 6 appears in nums and vice versa.
    // Hence, the lonely numbers in nums are [10, 8].
    // Note that [8, 10] may also be returned.
    fmt.Println(findLonely([]int{10,6,5,8})) // [10,8]
    // Example 2:
    // Input: nums = [1,3,5,3]
    // Output: [1,5]
    // Explanation: 
    // - 1 is a lonely number since it appears exactly once and 0 and 2 does not appear in nums.
    // - 5 is a lonely number since it appears exactly once and 4 and 6 does not appear in nums.
    // - 3 is not a lonely number since it appears twice.
    // Hence, the lonely numbers in nums are [1, 5].
    // Note that [5, 1] may also be returned.
    fmt.Println(findLonely([]int{1,3,5,3})) // [1,5]

    fmt.Println(findLonely1([]int{10,6,5,8})) // [10,8]
    fmt.Println(findLonely1([]int{1,3,5,3})) // [1,5]

    fmt.Println(findLonely2([]int{10,6,5,8})) // [10,8]
    fmt.Println(findLonely2([]int{1,3,5,3})) // [1,5]
}