package main

// LCR 010. 和为 K 的子数组
// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。

// 示例 1：
// 输入:nums = [1,1,1], k = 2
// 输出: 2
// 解释: 此题 [1,1] 与 [1,1] 为两种不同的情况

// 示例 2：
// 输入:nums = [1,2,3], k = 3
// 输出: 2
 
// 提示:
//     1 <= nums.length <= 2 * 10^4
//     -1000 <= nums[i] <= 1000
//     -10^7 <= k <= 10^7

import "fmt"

func subarraySum(nums []int, k int) int {
    res, l := 0, len(nums)
    for i := 0; i < l; i++ {
        sum := 0
        for j:= i; j < l; j++ {
            sum += nums[j]
            if sum == k {
                res++
            }
        }
    }
    return res
}

// best solution
// 要求找到连续区间和为 k 的子区间总数，即区间 [i,j] 内的和为 K ⇒ prefixSum[j] - prefixSum[i-1] == k。
// 所以 prefixSum[j] == k - prefixSum[i-1] 
func subarraySum1(nums []int, k int) int {
    res, pre := 0,0
    m := map[int]int{}
    m[0] = 1
    for i:= 0; i <len(nums); i++ {
        pre += nums[i]
        if _,ok := m[pre - k]; ok {
            res += m[pre - k]
        }
        m[pre] += 1 // 用 map 存储累加过的结果
    }
    return res
}

func main() {
    fmt.Println(subarraySum([]int{1,1,1}, 2)) // 2
    fmt.Println(subarraySum([]int{1,2,3}, 3)) // 2
    fmt.Println(subarraySum([]int{1,2,3,4,5,6,7,8,0}, 22)) // 1

    fmt.Println(subarraySum1([]int{1,1,1}, 2)) // 2
    fmt.Println(subarraySum1([]int{1,2,3}, 3)) // 2
    fmt.Println(subarraySum1([]int{1,2,3,4,5,6,7,8,0}, 22)) // 1
}