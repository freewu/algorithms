package main

// 368. Largest Divisible Subset
// Given a set of distinct positive integers nums, 
// return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:
//     answer[i] % answer[j] == 0, or
//     answer[j] % answer[i] == 0

// If there are multiple solutions, return any of them.

// Example 1:
// Input: nums = [1,2,3]
// Output: [1,2]
// Explanation: [1,3] is also accepted.

// Example 2:
// Input: nums = [1,2,4,8]
// Output: [1,2,4,8]

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 2 * 10^9
//     All the integers in nums are unique.

import "fmt"
import "sort"

func largestDivisibleSubset(nums []int) []int {
    // 先将集合排序
    sort.Ints(nums)
    dp, res := make([]int, len(nums)), []int{}
    for i := range dp {
        dp[i] = 1
    }
    maxSize, maxVal := 1, 1
    // 以某一个小的数作为基准，不断的选择能整除的数加入集合
    for i := 1; i < len(nums); i++ {
        for j, v := range nums[:i] {
            // 能整除 则 + 1
            if nums[i]%v == 0 && dp[j]+1 > dp[i] {
                dp[i] = dp[j] + 1
            }
        }
        if dp[i] > maxSize {
            maxSize, maxVal = dp[i], nums[i]
        }
    }
    if maxSize == 1 {
        return []int{nums[0]}
    }
    // 通过得到的 最大元素(maxVal) 反推出最大集合
    for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
        if dp[i] == maxSize && maxVal%nums[i] == 0 {
            res = append(res, nums[i])
            maxVal = nums[i]
            maxSize--
        }
    }
    return res
}

func largestDivisibleSubset1(nums []int) []int {
    sort.Ints(nums)
    n, mx := len(nums), 0
    res, dp := []int{}, make([][2]int, n + 1)
    for j := 1; j < n; j++ {
        exist := false
        for i := 0; i < j; i++ {
            if nums[j] % nums[i] != 0 { continue }
            exist = true
            if dp[i+1][1] >= dp[dp[j+1][0]][1] {
                dp[j+1][0] = i + 1
            }
        }
        if exist {
            dp[j+1][1] = dp[dp[j+1][0]][1] + 1
        }
        if dp[mx][1] < dp[j+1][1] {
            mx = j + 1
        }
    }
    for mx > 0 {
        res = append(res, nums[mx-1])
        mx = dp[mx][0]
    }
    if len(res) == 0 {
        res = append(res, nums[0])
    }
    return res
}

func largestDivisibleSubset2(nums []int) []int {
    n := len(nums)
    if n == 0 { return []int{} }
    sort.Ints(nums)
    dp := make([]int, n) // dp[i] = 以nums[i]结尾的最大整除子集的大小
    for i := range dp { // 初始化所有元素为1（单个元素自成一组）
        dp[i] = 1
    }
    maxSize, maxIndex := 1, 0 // 记录最大子集的大小和结束索引
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ { // 计算每个位置的最大子集大小
        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        if dp[i] > maxSize {
            maxSize, maxIndex = dp[i], i
        }
    }
    res,num, size := []int{}, nums[maxIndex], maxSize
    for i := maxIndex; i >= 0; i-- { // 根据dp数组构建结果 从后向前重建子集
        if num % nums[i] == 0 && dp[i] == size {
            res = append(res, nums[i])
            num = nums[i]
            size--
        }
    }
    return res // 无需反转结果，因为我们是从大到小构建的
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: [1,2]
    // Explanation: [1,3] is also accepted.
    fmt.Println(largestDivisibleSubset([]int{1,2,3})) // [1,2]
    // Example 2:
    // Input: nums = [1,2,4,8]
    // Output: [1,2,4,8]
    fmt.Println(largestDivisibleSubset([]int{1,2,4,8})) // [1,2,4,8]

    fmt.Println(largestDivisibleSubset([]int{1,2,3,4,5,6,7,8,9})) // [1,2,4,8]
    fmt.Println(largestDivisibleSubset([]int{9,8,7,6,5,4,3,2,1})) // [1,2,4,8]

    fmt.Println(largestDivisibleSubset1([]int{1,2,3})) // [1,2]
    fmt.Println(largestDivisibleSubset1([]int{1,2,4,8})) // [1,2,4,8]
    fmt.Println(largestDivisibleSubset1([]int{1,2,3,4,5,6,7,8,9})) // [1,2,4,8]
    fmt.Println(largestDivisibleSubset1([]int{9,8,7,6,5,4,3,2,1})) // [1,2,4,8]

    fmt.Println(largestDivisibleSubset2([]int{1,2,3})) // [1,2]
    fmt.Println(largestDivisibleSubset2([]int{1,2,4,8})) // [1,2,4,8]
    fmt.Println(largestDivisibleSubset2([]int{1,2,3,4,5,6,7,8,9})) // [1,2,4,8]
    fmt.Println(largestDivisibleSubset2([]int{9,8,7,6,5,4,3,2,1})) // [1,2,4,8]
}