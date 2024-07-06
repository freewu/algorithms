package main

// LCR 069. 山脉数组的峰顶索引
// 符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
//     arr.length >= 3
//     存在 i（0 < i < arr.length - 1）使得：
//         arr[0] < arr[1] < ... arr[i-1] < arr[i] 
//         arr[i] > arr[i+1] > ... > arr[arr.length - 1]

// 给定由整数组成的山峰数组 arr ，
// 返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i ，即山峰顶部。


// 示例 1：
// 输入：arr = [0,1,0]
// 输出：1

// 示例 2：
// 输入：arr = [1,3,5,4,2]
// 输出：2

// 示例 3：
// 输入：arr = [0,10,5,2]
// 输出：1

// 示例 4：
// 输入：arr = [3,4,5,1]
// 输出：2

// 示例 5：
// 输入：arr = [24,69,100,99,79,78,67,36,26,19]
// 输出：2

// 提示：
//     3 <= arr.length <= 10^4
//     0 <= arr[i] <= 10^6
//     题目数据保证 arr 是一个山脉数组

// 进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？

import "fmt"

func peakIndexInMountainArray(arr []int) int {
    res := len(arr) / 2 // 先取中间
    left,right := 0, len(arr) - 1
    for left < right {
        if arr[res + 1] < arr[res] && arr[res - 1 ] < arr[res] {// 找到了尖峰
            return res
        }
        if arr[res + 1] > arr[res] { // 还在左山坡上
            left = res
            right = right
        } else { // 还在右山坡上
            left = left
            right = res
        }
        res = (left + right) / 2
    }
    return res
}

func main() {
    fmt.Println(peakIndexInMountainArray([]int{ 0,1,0 })) // 1
    fmt.Println(peakIndexInMountainArray([]int{ 0,2,1,0 })) // 1
    fmt.Println(peakIndexInMountainArray([]int{ 0,10,5,2 })) // 1

    // 示例 1：
    // 输入：arr = [0,1,0]
    // 输出：1
    fmt.Println(peakIndexInMountainArray([]int{ 0,1,0 })) // 1
    // 示例 2：
    // 输入：arr = [1,3,5,4,2]
    // 输出：2
    fmt.Println(peakIndexInMountainArray([]int{ 1,3,5,4,2 })) // 2
    // 示例 3：
    // 输入：arr = [0,10,5,2]
    // 输出：1
    fmt.Println(peakIndexInMountainArray([]int{ 0,10,5,2 })) // 1
    // 示例 4：
    // 输入：arr = [3,4,5,1]
    // 输出：2
    fmt.Println(peakIndexInMountainArray([]int{ 3,4,5,1 })) // 2
    // 示例 5：
    // 输入：arr = [24,69,100,99,79,78,67,36,26,19]
    // 输出：2
    fmt.Println(peakIndexInMountainArray([]int{ 24,69,100,99,79,78,67,36,26,19 })) // 2
}