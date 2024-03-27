package main

// 713. Subarray Product Less Than K
// Given an array of integers nums and an integer k, 
// return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.

// Example 1:
// Input: nums = [10,5,2,6], k = 100
// Output: 8
// Explanation: The 8 subarrays that have product less than 100 are:
// [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
// Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

// Example 2:
// Input: nums = [1,2,3], k = 0
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 3 * 10^4
//     1 <= nums[i] <= 1000
//     0 <= k <= 10^6

import "fmt"

// 滑动窗口
func numSubarrayProductLessThanK(nums []int, k int) int {
    if k == 0 {
        return 0
    }
    res, left, right, prod := 0, 0, 0, 1
    // 在窗口滑动的过程中不断累乘，直到乘积大于 k，大于 k 的时候就缩小左窗口
    for left < len(nums) {
        if right < len(nums) && prod*nums[right] < k {
            prod = prod * nums[right]
            right++
        } else if left == right { // 左边窗口等于右窗口，这个时候需要左窗口和右窗口同时右移
            left++
            right++
        } else {
            res += right - left
            prod = prod / nums[left]
            left++ // // 大于 k 的时候就缩小左窗口
        }
    }
    return res
}

func numSubarrayProductLessThanK1(nums []int, k int) int {
	left, res, cur := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		cur = cur * nums[right]
		for left <= right && cur >= k {
			cur = cur / nums[left]
			left++
		}
		res = res + right - left + 1 // 每次移动到 right 位置都有 right - left + 1 个数组符合要求
	}
	return res
}

func main() {
    // Explanation: The 8 subarrays that have product less than 100 are:
    // [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
    // Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
    fmt.Println(numSubarrayProductLessThanK([]int{10,5,2,6},100)) // 8
    fmt.Println(numSubarrayProductLessThanK([]int{1,2,3},0)) // 0

    fmt.Println(numSubarrayProductLessThanK1([]int{10,5,2,6},100)) // 8
    fmt.Println(numSubarrayProductLessThanK1([]int{1,2,3},0)) // 0
}