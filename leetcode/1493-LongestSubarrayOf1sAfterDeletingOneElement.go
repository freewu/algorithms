package main

// 1493. Longest Subarray of 1's After Deleting One Element
// Given a binary array nums, you should delete one element from it.
// Return the size of the longest non-empty subarray containing only 1's in the resulting array. 
// Return 0 if there is no such subarray.

// Example 1:
// Input: nums = [1,1,0,1]
// Output: 3
// Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.

// Example 2:
// Input: nums = [0,1,1,1,0,1,1,0,1]
// Output: 5
// Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].

// Example 3:
// Input: nums = [1,1,1]
// Output: 2
// Explanation: You must delete one element.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.

import "fmt"

func longestSubarray(nums []int) int {
    res,l := 0, len(nums)
    for i := 0; i < l; i++ {
        delet, count := 1, 0
        for j := i; j < l; j++ {
            if nums[j] == 1 {// 是 1 就累加
                count++
            } else if delet == 1 { // 遇到 0，有1次删除机会使用掉
                delet--
                continue
            } else { // 遇到 0，没有删除机会了
                break
            }
        }
        if count > res { // 取最大的
            res = count
        }
    }
    if res == l { // 处理全是 1的情况(必须做一上删除操作)
        return l - 1
    }
    return res
}

func longestSubarray1(nums []int) int {
    sum, l, r := 0, 0, 0
    for ; r <len(nums); r++ {
        sum += nums[r]
        if sum < r - l {
            sum -= nums[l]
            l++
        }
    }
    return r - l - 1
}

func longestSubarray2(nums []int) int {
    res, left, right, k, isKUsed := 0, 0, 0, 0, false
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right < len(nums) {   
        if nums[right] == 1 {
            res = max(res, right - left + 1 - k)
            right++
        } else {
            if k == 0 { // use k
                k++
                res = max(res, right - left + 1 - k)
                right++
                isKUsed = true
            } else {
                // skip the first non zero
                for nums[left] == 1 {
                    left++
                }
                left++
                k--
            }
        }
    }
    if !isKUsed {
        res--
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,0,1]
    // Output: 3
    // Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
    fmt.Println(longestSubarray([]int{1,1,0,1})) // 3
    // Example 2:
    // Input: nums = [0,1,1,1,0,1,1,0,1]
    // Output: 5
    // Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
    fmt.Println(longestSubarray([]int{0,1,1,1,0,1,1,0,1})) // 5
    // Example 3:
    // Input: nums = [1,1,1]
    // Output: 2
    // Explanation: You must delete one element.
    fmt.Println(longestSubarray([]int{1,1,1})) // 2

    fmt.Println(longestSubarray1([]int{1,1,0,1})) // 3
    fmt.Println(longestSubarray1([]int{0,1,1,1,0,1,1,0,1})) // 5
    fmt.Println(longestSubarray1([]int{1,1,1})) // 2

    fmt.Println(longestSubarray2([]int{1,1,0,1})) // 3
    fmt.Println(longestSubarray2([]int{0,1,1,1,0,1,1,0,1})) // 5
    fmt.Println(longestSubarray2([]int{1,1,1})) // 2
}