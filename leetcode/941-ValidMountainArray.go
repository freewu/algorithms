package main

// 941. Valid Mountain Array
// Given an array of integers arr, return true if and only if it is a valid mountain array.

// Recall that arr is a mountain array if and only if:
//     arr.length >= 3
//     There exists some i with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// <img src="https://assets.leetcode.com/uploads/2019/10/20/hint_valid_mountain_array.png" />

// Example 1:
// Input: arr = [2,1]
// Output: false

// Example 2:
// Input: arr = [3,5,5]
// Output: false

// Example 3:
// Input: arr = [0,3,2,1]
// Output: true

// Constraints:
//     1 <= arr.length <= 10^4
//     0 <= arr[i] <= 10^4

import "fmt"

// func validMountainArray(arr []int) bool {
//     if len(arr) < 3 { 
//         return false 
//     }
//     flag := true // true 上升 false 下降
//     for i := 1; i < len(arr); i++ {
//         if flag == true { // 判断是否上升
//             if arr[i - 1] > arr[i] { // 开始下降
//                 if i == 0 { // 一开就下降
//                     return false
//                 }
//                 flag = false
//             }
//         } else {
//             if arr[i - 1] <= arr[i] { // 出现再次上升的情况
//                 return false
//             }
//         }
//     }
//     return true
// }

// 双指针
func validMountainArray(arr []int) bool {
    n := len(arr)
    l, r := 0, n - 1
    for l + 1 < n && arr[l] < arr[l + 1] { 
        l++ 
    }
    for r - 1 > 0 && arr[r] < arr[r - 1] { 
        r-- 
    }
    return l > 0 && r < n - 1 && l == r
}

func validMountainArray1(arr []int) bool {
    i, n := 1, len(arr)
    for i < n && arr[i] > arr[i-1] { // 向上走
        i++
    }
    if i == n || i == 1 { // 一直向上走完了
        return false
    }
    for i < n && arr[i] < arr[i-1] { // 向下
        i++
    }
    return i == n
}

func main() {
    // Example 1:
    // Input: arr = [2,1]
    // Output: false
    fmt.Println(validMountainArray([]int{2,1})) // false
    // Example 2:
    // Input: arr = [3,5,5]
    // Output: false
    fmt.Println(validMountainArray([]int{3,5,5})) // false
    // Example 3:
    // Input: arr = [0,3,2,1]
    // Output: true
    fmt.Println(validMountainArray([]int{0,3,2,1})) // true

    fmt.Println(validMountainArray1([]int{2,1})) // false
    fmt.Println(validMountainArray1([]int{3,5,5})) // false
    fmt.Println(validMountainArray1([]int{0,3,2,1})) // true
}