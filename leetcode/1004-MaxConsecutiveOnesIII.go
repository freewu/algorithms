package main

// 1004. Max Consecutive Ones III
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

// Example 1:
// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Example 2:
// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.
//     0 <= k <= nums.length

import "fmt"

func longestOnes(nums []int, k int) int {
    res, curSum, start := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for end := 0; end < len(nums); end++ {
        // that variable is to know how many 1s an 0s we on window
        curSum += nums[end]
        // end-start+1 that is length of the window and by
        //substracting curSum we know exactly how many zeros we
        //have on that window and it can exceed from k cuz we 
        //can turn k 0s to 1
        if end - start + 1 - curSum <= k {
            res = max(res, end - start + 1)
        } else {
            // condition doesnt met it is gonna move left 
            // pointer which is start by substracting the value
            // of start from curSum and adding one
            curSum -= nums[start]
            start++
        }
    }
    return res
}

func longestOnes1(nums []int, k int) int {
    res, left, lsum, rsum := 0, 0, 0, 0 // lsum 左前缀和、right 右前缀和
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right, v := range nums {
        rsum += 1 - v
        for lsum < rsum - k {
            lsum += 1 - nums[left]
            left++
        }
        res = max(res, right - left + 1)
    }
    return res
}

func longestOnes2(nums []int, k int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for l, r, nLen, zCnt := 0, 0, len(nums), 0; r < nLen; r++ {
        if nums[r] == 0 { // 统计0有多少
            zCnt++
        }
        for ; zCnt > k; l++ { // 填 k 个 0 后
            if nums[l] == 0 {
                zCnt--
            }
        }
        res = max(res, r - l + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
    // Output: 6
    // Explanation: [1,1,1,0,0,1,1,1,1,1,1]
    // Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
    fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0},2)) // 6
    // Example 2:
    // Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
    // Output: 10
    // Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
    // Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
    fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3)) // 10

    fmt.Println(longestOnes1([]int{1,1,1,0,0,0,1,1,1,1,0},2)) // 6
    fmt.Println(longestOnes1([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3)) // 10

    fmt.Println(longestOnes2([]int{1,1,1,0,0,0,1,1,1,1,0},2)) // 6
    fmt.Println(longestOnes2([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3)) // 10
}