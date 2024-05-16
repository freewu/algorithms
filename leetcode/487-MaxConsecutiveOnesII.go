package main

// 487. Max Consecutive Ones II
// Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.

// Example 1:
// Input: nums = [1,0,1,1,0]
// Output: 4
// Explanation: 
// - If we flip the first zero, nums becomes [1,1,1,1,0] and we have 4 consecutive ones.
// - If we flip the second zero, nums becomes [1,0,1,1,1] and we have 3 consecutive ones.
// The max number of consecutive ones is 4.

// Example 2:
// Input: nums = [1,0,1,1,0,1]
// Output: 4
// Explanation: 
// - If we flip the first zero, nums becomes [1,1,1,1,0,1] and we have 4 consecutive ones.
// - If we flip the second zero, nums becomes [1,0,1,1,1,1] and we have 4 consecutive ones.
// The max number of consecutive ones is 4.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.
 
// Follow up: What if the input numbers come in one by one as an infinite stream? In other words, you can't store all numbers coming from the stream as it's too large to hold in memory. Could you solve it efficiently?

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    res, left, right := 0, 0, 0
    pos := []int{} // 记录0的位置
    cnt0, cnt1 := 0, 0 // 0的数量，1的数量
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right < len(nums) {
        if nums[right] == 0 {
            cnt0++
            pos = append(pos, right)
        } else {
            cnt1++
        }
        right++
        if cnt0 > 1 { // 如果0的数量大于1，则窗口缩小，直到0的数量<1
            cnt1 =  cnt1 - (pos[0]-left) //1的数量减少从left到第一个0出现的位置个
            cnt0--
            left = pos[0] + 1  // 将left直接跳跃到第一个0出现的位置+1
            pos = pos[1:]
        }
        res = max(res, cnt0 + cnt1)
    }
    return res
}

func findMaxConsecutiveOnes1(nums []int) int {
    res, one_before_zero, cursor := 0, 0, 0
    zero_flap := false
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, n := range nums {
        if n == 1 {
            cursor++
        }
        if n == 0 {
            if zero_flap == false {
                zero_flap = true
            } else {
                res = max(res, one_before_zero+cursor + 1)
            }
            one_before_zero = cursor
            cursor = 0
        }
    }
    if zero_flap {
        return max(res, one_before_zero + cursor + 1)
    }
    return max(res, one_before_zero + cursor)
}

func main() {
    // Example 1:
    // Input: nums = [1,0,1,1,0]
    // Output: 4
    // Explanation: 
    // - If we flip the first zero, nums becomes [1,1,1,1,0] and we have 4 consecutive ones.
    // - If we flip the second zero, nums becomes [1,0,1,1,1] and we have 3 consecutive ones.
    // The max number of consecutive ones is 4.
    fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0})) // 4
    // Example 2:
    // Input: nums = [1,0,1,1,0,1]
    // Output: 4
    // Explanation: 
    // - If we flip the first zero, nums becomes [1,1,1,1,0,1] and we have 4 consecutive ones.
    // - If we flip the second zero, nums becomes [1,0,1,1,1,1] and we have 4 consecutive ones.
    // The max number of consecutive ones is 4.
    fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0,1})) // 4

    fmt.Println(findMaxConsecutiveOnes1([]int{1,0,1,1,0})) // 4
    fmt.Println(findMaxConsecutiveOnes1([]int{1,0,1,1,0,1})) // 4
}