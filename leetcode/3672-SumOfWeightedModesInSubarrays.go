package main

// 3672. Sum of Weighted Modes in Subarrays
// You are given an integer array nums and an integer k.

// For every subarray of length k:
//     1. The mode is defined as the element with the highest frequency. 
//        If there are multiple choices for a mode, the smallest such element is taken.
//     2. The weight is defined as mode * frequency(mode).

// Return the sum of the weights of all subarrays of length k.

// Note:
//     1. A subarray is a contiguous non-empty sequence of elements within an array.
//     2. The frequency of an element x is the number of times it occurs in the array.

// Example 1:
// Input: nums = [1,2,2,3], k = 3
// Output: 8
// Explanation:
// Subarrays of length k = 3 are:
// Subarray	Frequencies	Mode	Mode
// ​​​​​​​Frequency	Weight
// [1, 2, 2]	1: 1, 2: 2	2	2	2 × 2 = 4
// [2, 2, 3]	2: 2, 3: 1	2	2	2 × 2 = 4
// Thus, the sum of weights is 4 + 4 = 8.

// Example 2:
// Input: nums = [1,2,1,2], k = 2
// Output: 3
// Explanation:
// Subarrays of length k = 2 are:
// Subarray	Frequencies	Mode	Mode
// Frequency	Weight
// [1, 2]	1: 1, 2: 1	1	1	1 × 1 = 1
// [2, 1]	2: 1, 1: 1	1	1	1 × 1 = 1
// [1, 2]	1: 1, 2: 1	1	1	1 × 1 = 1
// Thus, the sum of weights is 1 + 1 + 1 = 3.

// Example 3:
// Input: nums = [4,3,4,3], k = 3
// Output: 14
// Explanation:
// Subarrays of length k = 3 are:
// Subarray	Frequencies	Mode	Mode
// Frequency	Weight
// [4, 3, 4]	4: 2, 3: 1	4	2	2 × 4 = 8
// [3, 4, 3]	3: 2, 4: 1	3	2	2 × 3 = 6
// Thus, the sum of weights is 8 + 6 = 14.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= nums.length

import "fmt"

// 超出时间限制 993 / 999 
func modeWeight(nums []int, k int) int64 {
    n := len(nums)
    if n == 0 || k <= 0 || k > n {
        return 0
    }
    // 频率映射：元素 -> 频率, 频率到元素集合的映射：频率 -> 元素集合
    freq, freqToNums := make(map[int]int), make(map[int]map[int]bool)
    res, mx := 0, 0
    // 辅助函数：更新频率映射
    updateFrequency := func(num int, oldFreq, newFreq int) {
        // 从旧频率集合中移除
        if oldFreq > 0 {
            delete(freqToNums[oldFreq], num)
            if len(freqToNums[oldFreq]) == 0 {
                delete(freqToNums, oldFreq)
            }
        }
        // 如果新频率大于0，添加到新频率集合中
        if newFreq > 0 {
            if _, exists := freqToNums[newFreq]; !exists {
                freqToNums[newFreq] = make(map[int]bool)
            }
            freqToNums[newFreq][num] = true
        }
    }
    // 找到集合中的最小元素
    findMinInSet := func(s map[int]bool) int {
        mn := -1
        for v := range s {
            if mn == -1 || v < mn {
                mn = v
            }
        }
        return mn
    }
    // 检查 map 中是否包含指定的键
    containsKey := func(mp map[int]map[int]bool, key int) bool {
        _, exists := mp[key]
        return exists
    } 
    // 初始化第一个窗口
    for i := 0; i < k; i++ {
        num := nums[i]
        oldFreq := freq[num]
        freq[num]++
        newFreq := oldFreq + 1
        updateFrequency(num, oldFreq, newFreq)
        if newFreq > mx {
            mx = newFreq
        }
    }
    // 找到第一个窗口的众数
    mode := findMinInSet(freqToNums[mx])
    res += (mode * mx)
    
    // 滑动窗口处理剩余子数组
    for i := k; i < n; i++ {
        // 移除窗口最左边的元素
        leftNum := nums[i-k]
        oldFreq := freq[leftNum]
        freq[leftNum]--
        newFreq := oldFreq - 1
        updateFrequency(leftNum, oldFreq, newFreq)
        if newFreq == 0 {
            delete(freq, leftNum)
        }
        // 如果移除的是最大频率，需要重新寻找最大频率
        if oldFreq == mx && !containsKey(freqToNums, oldFreq) {
            mx--
        }
        // 添加新的元素到窗口右边
        rightNum := nums[i]
        oldFreq = freq[rightNum]
        freq[rightNum]++
        newFreq = oldFreq + 1
        updateFrequency(rightNum, oldFreq, newFreq)
        // 更新最大频率
        if newFreq > mx {
            mx = newFreq
        }
        // 找到当前窗口的众数并计算权重
        mode := findMinInSet(freqToNums[mx])
        res += mode * mx
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3], k = 3
    // Output: 8
    // Explanation:
    // Subarrays of length k = 3 are:
    // Subarray	Frequencies	Mode	Mode
    // ​​​​​​​Frequency	Weight
    // [1, 2, 2]	1: 1, 2: 2	2	2	2 × 2 = 4
    // [2, 2, 3]	2: 2, 3: 1	2	2	2 × 2 = 4
    // Thus, the sum of weights is 4 + 4 = 8.
    fmt.Println(modeWeight([]int{1,2,2,3}, 3)) // 8
    // Example 2:
    // Input: nums = [1,2,1,2], k = 2
    // Output: 3
    // Explanation:
    // Subarrays of length k = 2 are:
    // Subarray	Frequencies	Mode	Mode
    // Frequency	Weight
    // [1, 2]	1: 1, 2: 1	1	1	1 × 1 = 1
    // [2, 1]	2: 1, 1: 1	1	1	1 × 1 = 1
    // [1, 2]	1: 1, 2: 1	1	1	1 × 1 = 1
    // Thus, the sum of weights is 1 + 1 + 1 = 3.
    fmt.Println(modeWeight([]int{1,2,1,2}, 2)) // 3
    // Example 3:
    // Input: nums = [4,3,4,3], k = 3
    // Output: 14
    // Explanation:
    // Subarrays of length k = 3 are:
    // Subarray	Frequencies	Mode	Mode
    // Frequency	Weight
    // [4, 3, 4]	4: 2, 3: 1	4	2	2 × 4 = 8
    // [3, 4, 3]	3: 2, 4: 1	3	2	2 × 3 = 6
    // Thus, the sum of weights is 8 + 6 = 14.
    fmt.Println(modeWeight([]int{4,3,4,3}, 3)) // 14

    fmt.Println(modeWeight([]int{1,2,3,4,5,6,7,8,9}, 2)) // 36
    fmt.Println(modeWeight([]int{9,8,7,6,5,4,3,2,1}, 2)) // 36
}