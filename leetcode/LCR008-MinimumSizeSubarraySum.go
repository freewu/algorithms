package main

// LCR 008. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
// 如果不存在符合条件的子数组，返回 0 。

// 示例 1：
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

// 示例 2：
// 输入：target = 4, nums = [1,4,4]
// 输出：1

// 示例 3：
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

// 提示：
//     1 <= target <= 10^9
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
 
// 进阶：
//     如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

import "fmt"

func minSubArrayLen(target int, nums []int) int {
    left, sum, res := 0, 0, len(nums) + 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for right, v := range nums { // 从头开始向尾部
        sum += v // 累加
        for sum >= target { // 如果超过目标值
            res = min(res, right - left + 1)  // 取个最小的长度
            sum -= nums[left] // 减掉开头的值
            left++ // 向尾部走一下
        }
    }
    if res == len(nums) + 1 { // 所有值累加都达不到目标值 返回 0
        return 0
    }
    return res
}

// best solution
func minSubArrayLen1(target int, nums []int) int {
    // minLength := len(nums)
    left, sum, res := 0, 0, 1 << 32 -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i:= 0; i < len(nums); i++ {
        sum += nums[i]
        for sum >= target {
            res = min(res, i + 1 - left)
            sum -= nums[left]
            left++
        }
    }
    if res != 1 << 32 -1 {
        return res
    } else {
        return 0
    }
}

func minSubArrayLen2(target int, nums []int) int {
    i, j, sum, res := 0, 0, 0, len(nums) + 1
    for ; j < len(nums); j++ {
        sum += nums[j]
        for sum >= target {
            if res > j - i + 1 {
                res = j - i + 1
            }
            sum -= nums[i]
            i++
        }
    }
    if res == len(nums) + 1 {
        return 0
    }
    return res
}

func main() {
    fmt.Printf("minSubArrayLen(7,[]int{2,3,1,2,4,3}) = %v\n",minSubArrayLen(7,[]int{2,3,1,2,4,3})) // 2  [4,3]
    fmt.Printf("minSubArrayLen(4,[]int{1,4,4}) = %v\n",minSubArrayLen(4,[]int{1,4,4})) // 1 [4]
    fmt.Printf("minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1}) = %v\n",minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1})) // 0

    fmt.Printf("minSubArrayLen1(7,[]int{2,3,1,2,4,3}) = %v\n",minSubArrayLen1(7,[]int{2,3,1,2,4,3})) // 2  [4,3]
    fmt.Printf("minSubArrayLen1(4,[]int{1,4,4}) = %v\n",minSubArrayLen1(4,[]int{1,4,4})) // 1 [4]
    fmt.Printf("minSubArrayLen1(11,[]int{1,1,1,1,1,1,1,1}) = %v\n",minSubArrayLen1(11,[]int{1,1,1,1,1,1,1,1})) // 0

    fmt.Printf("minSubArrayLen2(7,[]int{2,3,1,2,4,3}) = %v\n",minSubArrayLen2(7,[]int{2,3,1,2,4,3})) // 2  [4,3]
    fmt.Printf("minSubArrayLen2(4,[]int{1,4,4}) = %v\n",minSubArrayLen2(4,[]int{1,4,4})) // 1 [4]
    fmt.Printf("minSubArrayLen2(11,[]int{1,1,1,1,1,1,1,1}) = %v\n",minSubArrayLen2(11,[]int{1,1,1,1,1,1,1,1})) // 0
}