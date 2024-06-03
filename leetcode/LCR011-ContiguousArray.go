package main

// LCR 011. 连续数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

// 示例 1：
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

// 示例 2：
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
 
// 提示：
//     1 <= nums.length <= 10^5
//     nums[i] 不是 0 就是 1

import "fmt"

func findMaxLength(nums []int) int {
    max := func (a, b int) int { if a > b { return a; }; return b; }
    dict := map[int]int{}
    // 0 和 1 的数量相同可以转化为两者数量相差为 0，如果将 0 看作为 -1，那么原题转化为求最长连续子数组，其元素和为 0 。
    // 又变成了区间内求和的问题，自然而然转换为前缀和来处理。
        // 假设连续子数组是 [i,j] 区间，这个区间内元素和为 0 意味着 prefixSum[j] - prefixSum[i] = 0，也就是 prefixSum[i] = prefixSum[j]
    dict[0] = -1
    count, res := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            count--
        } else {
            count++
        }
        // 一旦某个 key 存在了，代表之前某个下标的前缀和和当前下标构成的区间，这段区间内的元素和为 0 。
        // 这个区间是所求。扫完整个数组，扫描过程中动态更新最大区间长度，扫描完成便可得到最大区间长度，即最长连续子数组
        if index, ok := dict[count]; ok {
            res = max(res, i - index)
        } else {
            dict[count] = i
        }
        //fmt.Println("dict: ",dict)
    }
    return res
}

// best solution
func findMaxLength1(nums []int) int {
    max := func (a, b int) int { if a > b { return a; }; return b; }
    // 0，-1，1，+1，相同前缀和出现的下标，只记录第一个，每多出现一次更新res
    res, pre, cnt := 0, 0, map[int]int{ 0: -1 }
    for i,x := range nums {
        if x == 0 {
            pre--
        } else {
            pre++
        }
        nums[i] = pre
        if j,ok := cnt[pre]; ok {
            res = max(res,i-j)
        } else {
            cnt[pre] = i
        }
    }
    return res
}

func main() {
    // Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
    fmt.Println(findMaxLength([]int{0,1})) // 2
    // Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
    fmt.Println(findMaxLength([]int{0,1,0})) // 2
    fmt.Println(findMaxLength([]int{0,1,0,1,1,0,0,1,1,0,1,0,0,1,0})) // 14
    fmt.Println(findMaxLength([]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 0
    fmt.Println(findMaxLength([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})) // 0

    fmt.Println(findMaxLength1([]int{0,1})) // 2
    fmt.Println(findMaxLength1([]int{0,1,0})) // 2
    fmt.Println(findMaxLength1([]int{0,1,0,1,1,0,0,1,1,0,1,0,0,1,0})) // 14
    fmt.Println(findMaxLength1([]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 0
    fmt.Println(findMaxLength1([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})) // 0
}