package main

// LCR 009. 乘积小于 K 的子数组
// 给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

// 示例 1:
// 输入: nums = [10,5,2,6], k = 100
// 输出: 8
// 解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
// 需要注意的是 [10,5,2] 并不是乘积小于100的子数组。

// 示例 2:
// 输入: nums = [1,2,3], k = 0
// 输出: 0

// 提示: 
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

func numSubarrayProductLessThanK2(nums []int, k int) int {
    res, i, j, multiply := 0, 0, 0, 1
    for j < len(nums) {
        multiply *= nums[j]
        for multiply >= k && i <= j {
            multiply /= nums[i]
            i++
        }
        res += j - i + 1
        j++
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

    fmt.Println(numSubarrayProductLessThanK2([]int{10,5,2,6},100)) // 8
    fmt.Println(numSubarrayProductLessThanK2([]int{1,2,3},0)) // 0
}