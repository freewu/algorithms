package main

// 152. Maximum Product Subarray
// Given an integer array nums, find a subarray that has the largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:
// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.

// Example 2:
// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 
// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     -10 <= nums[i] <= 10
//     The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// # 解题思路
//     给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）
//     最大值是 Max(f(n)) = Max( Max(f(n-1)) * n, Min(f(n-1)) * n)；最小值是 Min(f(n)) = Min( Max(f(n-1)) * n, Min(f(n-1)) * n)。
//     只要动态维护这两个值，如果最后一个数是负数，最大值就在负数 * 最小值中产生，如果最后一个数是正数，最大值就在正数 * 最大值中产生。

import "fmt"

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    low, high, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 { // 如果当前为负数，大小互换
            high, low = low, high
        }
        high = max(nums[i], high * nums[i]) // 取最大值
        low = min(nums[i], low * nums[i]) // 取最小值 如有负数  负数 * 最小值中产生  负负得正
        res = max(res, high)
    }
    return res
}

func maxProduct1(nums []int) int {
    high, low, res := nums[0], nums[0], nums[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        mx, mn := high, low
        high = max(mx * nums[i], max(nums[i], mn * nums[i]))
        low = min(mn * nums[i], min(nums[i], mx * nums[i]))
        res = max(high, res)
    }
    return res
}

func main() {
    fmt.Printf("maxProduct([]int{2,3,-2,4}) = %v\n",maxProduct([]int{2,3,-2,4})) // 6 = 2 * 3
    fmt.Printf("maxProduct([]int{-2,0,-1}) = %v\n",maxProduct([]int{-2,0,-1})) // 0

    fmt.Printf("maxProduct1([]int{2,3,-2,4}) = %v\n",maxProduct1([]int{2,3,-2,4})) // 6 = 2 * 3
    fmt.Printf("maxProduct1([]int{-2,0,-1}) = %v\n",maxProduct1([]int{-2,0,-1})) // 0
}
