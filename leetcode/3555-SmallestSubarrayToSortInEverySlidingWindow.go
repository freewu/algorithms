package main

// 3555. Smallest Subarray to Sort in Every Sliding Window
// You are given an integer array nums and an integer k.

// For each contiguous subarray of length k, determine the minimum length of a continuous segment that must be sorted so that the entire window becomes non‑decreasing; if the window is already sorted, its required length is zero.

// Return an array of length n − k + 1 where each element corresponds to the answer for its window.

// Example 1:
// Input: nums = [1,3,2,4,5], k = 3
// Output: [2,2,0]
// Explanation:
// nums[0...2] = [1, 3, 2]. Sort [3, 2] to get [1, 2, 3], the answer is 2.
// nums[1...3] = [3, 2, 4]. Sort [3, 2] to get [2, 3, 4], the answer is 2.
// nums[2...4] = [2, 4, 5] is already sorted, so the answer is 0.

// Example 2:
// Input: nums = [5,4,3,2,1], k = 4
// Output: [4,4]
// Explanation:
// nums[0...3] = [5, 4, 3, 2]. The whole subarray must be sorted, so the answer is 4.
// nums[1...4] = [4, 3, 2, 1]. The whole subarray must be sorted, so the answer is 4.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= k <= nums.length
//     1 <= nums[i] <= 10^6

import "fmt"
import "sort"
import "reflect"

func minSubarraySort(nums []int, k int) []int {
    n := len(nums)
    res := make([]int, n - k + 1)
    for i := 0; i <= n - k; i++ {
        // 使用滑动窗口方法，遍历数组中所有长度为 k 的子数组
        window := make([]int, k)
        copy(window, nums[i:i+k])
        // 对每个窗口，创建其副本并排序，以便与原始窗口进行比较
        sortedWindow := make([]int, k)
        copy(sortedWindow, window)
        sort.Ints(sortedWindow)
        // 检查是否已排序：使用reflect.DeepEqual比较原始窗口和排序后的窗口，如果相同则直接记录结果为 0
        if reflect.DeepEqual(window, sortedWindow) { 
            res[i] = 0
            continue
        }
        left, right := 0, k-1
        // 左边界：从左向右找到第一个与排序后位置不同的元素
        for left < k && window[left] == sortedWindow[left] {
            left++
        }
        // 右边界：从右向左找到第一个与排序后位置不同的元素。
        for right >= 0 && window[right] == sortedWindow[right] {
            right--
        }
        // 长度计算：右边界减去左边界加 1 即为需要排序的最小长度
        res[i] = right - left + 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,4,5], k = 3
    // Output: [2,2,0]
    // Explanation:
    // nums[0...2] = [1, 3, 2]. Sort [3, 2] to get [1, 2, 3], the answer is 2.
    // nums[1...3] = [3, 2, 4]. Sort [3, 2] to get [2, 3, 4], the answer is 2.
    // nums[2...4] = [2, 4, 5] is already sorted, so the answer is 0.
    fmt.Println(minSubarraySort([]int{1,3,2,4,5}, 3)) // [2,2,0]
    // Example 2:
    // Input: nums = [5,4,3,2,1], k = 4
    // Output: [4,4]
    // Explanation:
    // nums[0...3] = [5, 4, 3, 2]. The whole subarray must be sorted, so the answer is 4.
    // nums[1...4] = [4, 3, 2, 1]. The whole subarray must be sorted, so the answer is 4.
    fmt.Println(minSubarraySort([]int{5,4,3,2,1}, 4)) // [4,4]

    fmt.Println(minSubarraySort([]int{1,2,3,4,5,6,7,8,9}, 4)) // [4,4]
    fmt.Println(minSubarraySort([]int{9,8,7,6,5,4,3,2,1}, 4)) // [4,4]
}