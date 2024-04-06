package main

// 410. Split Array Largest Sum
// Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.
// Return the minimized largest sum of the split.
// A subarray is a contiguous part of the array.

// Example 1:
// Input: nums = [7,2,5,10,8], k = 2
// Output: 18
// Explanation: There are four ways to split nums into two subarrays.
// The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18.

// Example 2:
// Input: nums = [1,2,3,4,5], k = 2
// Output: 9
// Explanation: There are four ways to split nums into two subarrays.
// The best way is to split it into [1,2,3] and [4,5], where the largest sum among the two subarrays is only 9.
 
// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 10^6
//     1 <= k <= min(50, nums.length)

import "fmt"

func splitArray(nums []int, m int) int {
    maxNum, sum := 0, 0
    for _, num := range nums {
        sum += num
        if num > maxNum {
            maxNum = num
        }
    }
    if m == 1 {
        return sum
    }
    calSum := func (mid, m int, nums []int) bool {
        sum, count := 0, 0
        for _, v := range nums {
            sum += v
            if sum > mid {
                sum = v
                count++
                // 分成 m 块，只需要插桩 m -1 个
                if count > m - 1 {
                    return false
                }
            }
        }
        return true
    }
    low, high := maxNum, sum
    for low < high {
        mid := low + (high-low)>>1
        if calSum(mid, m, nums) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func main() {
    fmt.Println(splitArray([]int{7,2,5,10,8},2)) // 18 [10,8] - 14 [7,2,5] = 4
    fmt.Println(splitArray([]int{1,2,3,4,5},2)) // 9 [4,5] - 6 [1,2,3] = 3
}