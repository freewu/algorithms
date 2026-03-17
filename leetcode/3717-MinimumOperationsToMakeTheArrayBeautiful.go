package main

// 3717. Minimum Operations to Make the Array Beautiful
// You are given an integer array nums.

// An array is called beautiful if for every index i > 0, the value at nums[i] is divisible by nums[i - 1].

// In one operation, you may increment any element of nums by 1.

// Return the minimum number of operations required to make the array beautiful.

// Example 1:
// Input: nums = [3,7,9]
// Output: 2
// Explanation:
// Applying the operation twice on nums[1] makes the array beautiful: [3,9,9]

// Example 2:
// Input: nums = [1,1,1]
// Output: 0
// Explanation:
// The given array is already beautiful.

// Example 3:
// Input: nums = [4]
// Output: 0
// Explanation:
// The array has only one element, so it's already beautiful.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "slices"

// Wrong Answer 916 / 1047 testcases passed
// 美丽数组的定义：对于数组中的每个元素（除了第一个元素），该元素必须能被前一个元素整除
// 解题思路：
//     1. 从数组的第二个元素开始遍历，计算当前元素与前一个元素的余数
//     2. 如果余数不为 0，计算需要增加的值（使得当前元素能被前一个元素整除），并将该值累加到总操作次数中
//     3. 更新当前元素的值（增加后的值），并将其设为下一次迭代的前一个元素
//     4. 遍历完成后，返回总操作次数
func minOperations(nums []int) int {
    res, n, prev := 0, len(nums), nums[0]
    if n <= 1 { return 0 } // 如果数组长度小于等于 1，数组已经是美丽数组，直接返回 0
    for i := 1; i < n; i++ {
        curr:= nums[i]
        rem := curr % prev // 计算当前元素与前一个元素的余数
        if rem != 0 { // 如果余数不为 0，计算需要增加的值（使得当前元素能被前一个元素整除），并将该值累加到总操作次数中。
            add := prev - rem
            res += add
            curr += add
        }
        prev = curr // 更新当前元素的值（增加后的值），并将其设为下一次迭代的前一个元素
    }
    return res
}

// Wrong Answer 916 / 1047 testcases passed
func minOperations1(nums []int) int {
    res, n, prev := 0, len(nums), nums[0]
    if n <= 1 {
        return 0
    }
    for i := 1; i < n; i++ {
        curr := nums[i]
        k := (curr + prev - 1) / prev
        res += (k * prev - curr)
        prev= k * prev
    }
    return res
}

func minOperations2(nums []int) int {
    n, inf := len(nums), 1 << 61
    if n == 0 { return 0 }
    mx := nums[0]
    for _, v := range nums { // 找到数组中的最大值
        if v > mx {
            mx = v
        }
    }
    threshold := 2 * mx
    dp := make([]int, threshold + 1)
    for i := range dp { // 初始化 dp 数组，填充无穷大
        dp[i] = inf
    }
    dp[nums[0]] = 0
    for i := 1; i < n; i++ { // 动态规划核心逻辑
        // 临时数组存储当前轮次的 dp 值（避免覆盖）
        newDp := make([]int, threshold + 1)
        copy(newDp, dp)
        for j := threshold; j >= nums[i]; j-- { // 倒序遍历 j，从 threshold 到 nums[i]
            curr := inf
            // 遍历 k 从 nums[i-1] 到 j，寻找能整除 j 的 k
            for k := nums[i-1]; k <= j; k++ {
                if j % k == 0 && dp[k] != inf {
                    // 计算操作数：之前的操作数 + 当前需要增加的数值
                    op := dp[k] + (j - nums[i])
                    if op < curr {
                        curr = op
                    }
                }
            }
            newDp[j] = curr
        }
        for j := 0; j < nums[i]; j++ { // 将小于 nums[i] 的位置重置为无穷大（无法通过增加得到）
            newDp[j] = inf
        }
        dp = newDp
    }
    res := inf
    for _, v := range dp { // 找到 dp 数组中的最小值
        if v < res {
            res = v
        }
    }
    return res
}

func minOperations3(nums []int) int {
    const inf = 1 << 61
    top := slices.Max(nums) * 2
    f, g := make([]int, top), make([]int, top)
    for i := range f {
        f[i] = inf
    }
    f[nums[0]] = 0
    for _, v := range nums[1:] {
        for i := range g {
            g[i] = inf
        }
        for preVal, preCnt := range f {
            if preCnt >= inf {
                continue
            }
            for newVal := (v + preVal - 1) / preVal * preVal; newVal < top; newVal += preVal {
                g[newVal] = min(g[newVal], preCnt + newVal - v)
            }
        }
        f, g = g, f
    }
    return slices.Min(f)
}

func main() {
    // Example 1:
    // Input: nums = [3,7,9]
    // Output: 2
    // Explanation:
    // Applying the operation twice on nums[1] makes the array beautiful: [3,9,9]
    fmt.Println(minOperations([]int{3,7,9})) // 2
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 0
    // Explanation:
    // The given array is already beautiful.
    fmt.Println(minOperations([]int{1,1,1})) // 0
    // Example 3:
    // Input: nums = [4]
    // Output: 0
    // Explanation:
    // The array has only one element, so it's already beautiful.
    fmt.Println(minOperations([]int{4})) // 0

    fmt.Println(minOperations([]int{5,13,18})) // 9
    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 36

    fmt.Println(minOperations1([]int{3,7,9})) // 2
    fmt.Println(minOperations1([]int{1,1,1})) // 0
    fmt.Println(minOperations1([]int{4})) // 0
    fmt.Println(minOperations1([]int{5,13,18})) // 9
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1})) // 36

    fmt.Println(minOperations2([]int{3,7,9})) // 2
    fmt.Println(minOperations2([]int{1,1,1})) // 0
    fmt.Println(minOperations2([]int{4})) // 0
    fmt.Println(minOperations2([]int{5,13,18})) // 9
    fmt.Println(minOperations2([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(minOperations2([]int{9,8,7,6,5,4,3,2,1})) // 36
    
    fmt.Println(minOperations3([]int{3,7,9})) // 2
    fmt.Println(minOperations3([]int{1,1,1})) // 0
    fmt.Println(minOperations3([]int{4})) // 0
    fmt.Println(minOperations3([]int{5,13,18})) // 9
    fmt.Println(minOperations3([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(minOperations3([]int{9,8,7,6,5,4,3,2,1})) // 36
}

