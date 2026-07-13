package main

// 3991. Sort Array Using Prefix Reversals
// You are given an integer array nums of length n, where nums is a permutation of the integers in the range [0, n - 1].

// You are also given an integer array pre, where each pre[i] is a valid prefix length.

// In one operation, you may choose any length x from pre and reverse the first x elements of nums.

// For example, applying a prefix reversal of length 3 on [4, 1, 2, 3] results in [2, 1, 4, 3].

// Return the minimum number of operations required to sort nums in ascending order. 
// If it is impossible to sort nums, return -1.

// Example 1:
// Input: nums = [2,0,1], pre = [2,3]
// Output: 2
// Explanation:
// Reverse pre[1] = 3 elements to get nums = [1, 0, 2].
// Then reverse pre[0] = 2 elements to get nums = [0, 1, 2].
// Thus, the minimum number of prefix reversal required is 2.

// Example 2:
// Input: nums = [1,0,2], pre = [1,3]
// Output: -1
// Explanation:
// It is impossible to sort the array using the given prefix lengths, so the answer is -1.

// Example 3:
// Input: nums = [0,1], pre = [2]
// Output: 0
// Explanation:
// Since nums is already sorted, no prefix reversals are needed. Thus, the answer is 0.

// Constraints:
//     1 <= n == nums.length <= 8
//     0 <= nums[i] <= n - 1
//     1 <= pre.length <= n
//     1 <= pre[i] <= n
//     ​​​​​​​nums is a permutation of integers from 0 to n - 1.
//     pre consists of unique integers.

import "fmt"
import "strings"

func sortArray(nums []int, pre []int) int {
    n := len(nums)
    // 构造目标有序数组
    target := make([]int, n)
    for i := 0; i < n; i++ {
        target[i] = i
    }
    // 判断数组是否相等
    equal := func(a, b []int) bool {
        for i := range a {
            if a[i] != b[i] {
                return false
            }
        }
        return true
    }
    if equal(nums, target) {
        return 0
    }
    // 数组转字符串用于去重
    arrToString := func(arr []int) string {
        var s []string
        for _, v := range arr {
            s = append(s, fmt.Sprintf("%d", v))
        }
        return strings.Join(s, ",")
    }
    // 翻转前x个元素，返回新数组
    reversePrefix := func(arr []int, x int) []int {
        newArr := make([]int, n)
        copy(newArr, arr)
        l, r := 0, x-1
        for l < r {
            newArr[l], newArr[r] = newArr[r], newArr[l]
            l++
            r--
        }
        return newArr
    }
    type State struct {
        arr []int
        step int
    }
    queue := []State{{arr: nums, step: 0}}
    visited := make(map[string]bool)
    visited[arrToString(nums)] = true
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        // 遍历所有可用翻转长度
        for _, x := range pre {
            nextArr := reversePrefix(cur.arr, x)
            if equal(nextArr, target) {
                return cur.step + 1
            }
            key := arrToString(nextArr)
            if !visited[key] {
                visited[key] = true
                queue = append(queue, State{arr: nextArr, step: cur.step + 1})
            }
        }
    }
    // 队列遍历完无法得到有序数组
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,0,1], pre = [2,3]
    // Output: 2
    // Explanation:
    // Reverse pre[1] = 3 elements to get nums = [1, 0, 2].
    // Then reverse pre[0] = 2 elements to get nums = [0, 1, 2].
    // Thus, the minimum number of prefix reversal required is 2.
    fmt.Println(sortArray([]int{2,0,1}, []int{2,3})) // 2
    // Example 2:
    // Input: nums = [1,0,2], pre = [1,3]
    // Output: -1
    // Explanation:
    // It is impossible to sort the array using the given prefix lengths, so the answer is -1.
    fmt.Println(sortArray([]int{1,0,2}, []int{1,3})) // -1
    // Example 3:
    // Input: nums = [0,1], pre = [2]
    // Output: 0
    // Explanation:
    // Since nums is already sorted, no prefix reversals are needed. Thus, the answer is 0.
    fmt.Println(sortArray([]int{0,1}, []int{2})) // 0

    fmt.Println(sortArray([]int{0,1,2,3,4,5,6,7}, []int{1,2,3,4,5,6,7,8})) // 0
    fmt.Println(sortArray([]int{0,1,2,3,4,5,6,7}, []int{8,7,6,5,4,3,2,1})) // 0
    fmt.Println(sortArray([]int{7,6,5,4,3,2,1,0}, []int{1,2,3,4,5,6,7,8})) // 1
    fmt.Println(sortArray([]int{7,6,5,4,3,2,1,0}, []int{8,7,6,5,4,3,2,1})) // 1
}