package main

// 852. Peak Index in a Mountain Array
// An array arr is a mountain if the following properties hold:    
//     arr.length >= 3
//     There exists some i with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// Given a mountain array arr, 
// return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].
// You must solve it in O(log(arr.length)) time complexity.

// Example 1:
// Input: arr = [0,1,0]
// Output: 1

// Example 2:
// Input: arr = [0,2,1,0]
// Output: 1

// Example 3:
// Input: arr = [0,10,5,2]
// Output: 1
 
// Constraints:
//     3 <= arr.length <= 10^5
//     0 <= arr[i] <= 10^6
//     arr is guaranteed to be a mountain array.

import "fmt"

func peakIndexInMountainArray(arr []int) int {
    res := len(arr) / 2 // 先取中间
    left,right := 0, len(arr) - 1
    for left < right {
        // 找到了尖峰
        if arr[res + 1] < arr[res] && arr[res - 1 ] < arr[res] {
            return res
        }
        // 还在左山坡上
        if arr[res + 1] > arr[res] {
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
}