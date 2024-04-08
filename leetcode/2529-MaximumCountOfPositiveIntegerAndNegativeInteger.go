package main

// 2529. Maximum Count of Positive Integer and Negative Integer
// Given an array nums sorted in non-decreasing order, 
// return the maximum between the number of positive integers and the number of negative integers.
//     In other words, if the number of positive integers in nums is pos and the number of negative integers is neg, then return the maximum of pos and neg.

// Note that 0 is neither positive nor negative.

// Example 1:
// Input: nums = [-2,-1,-1,1,2,3]
// Output: 3
// Explanation: There are 3 positive integers and 3 negative integers. The maximum count among them is 3.

// Example 2:
// Input: nums = [-3,-2,-1,0,0,1,2]
// Output: 3
// Explanation: There are 2 positive integers and 3 negative integers. The maximum count among them is 3.

// Example 3:
// Input: nums = [5,20,66,1314]
// Output: 4
// Explanation: There are 4 positive integers and 0 negative integers. The maximum count among them is 4.
 
// Constraints:
//     1 <= nums.length <= 2000
//     -2000 <= nums[i] <= 2000
//     nums is sorted in a non-decreasing order.

import "fmt"
import "sort"

// 暴力解法
func maximumCount(nums []int) int {
    pos, neg := 0, 0
    for _, v := range nums {
        if v > 0 { 
            pos++
        } else if v < 0 {
            neg++
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(pos, neg)
}

// use lib
func maximumCount1(nums []int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(sort.SearchInts(nums, 0), len(nums) - sort.SearchInts(nums, 1))
}

// 二分法
func maximumCount2(nums []int) int {
    binarySearch := func (nums[]int,target int, flag bool) int {
        low, high := 0, len(nums)
        pos := func(mid int, target int, flag bool) bool {
            if flag {
                return mid < target
            } else {
                return mid <= target
            }
        }
        for low < high {
            mid := low + (high - low) >> 1
            if pos(nums[mid], target, flag) {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(binarySearch(nums, 0, true), len(nums) - binarySearch(nums, 0, false))
}

func main() {
    // Explanation: There are 3 positive integers and 3 negative integers. The maximum count among them is 3.
    fmt.Println(maximumCount([]int{-2,-1,-1,1,2,3})) // 3
    // Explanation: There are 2 positive integers and 3 negative integers. The maximum count among them is 3.
    fmt.Println(maximumCount([]int{-3,-2,-1,0,0,1,2})) // 3
    // Explanation: There are 4 positive integers and 0 negative integers. The maximum count among them is 4.
    fmt.Println(maximumCount([]int{5,20,66,1314})) // 4

    fmt.Println(maximumCount1([]int{-2,-1,-1,1,2,3})) // 3
    fmt.Println(maximumCount1([]int{-3,-2,-1,0,0,1,2})) // 3
    fmt.Println(maximumCount1([]int{5,20,66,1314})) // 4

    fmt.Println(maximumCount2([]int{-2,-1,-1,1,2,3})) // 3
    fmt.Println(maximumCount2([]int{-3,-2,-1,0,0,1,2})) // 3
    fmt.Println(maximumCount2([]int{5,20,66,1314})) // 4
}