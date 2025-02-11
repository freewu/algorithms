package main

// 面试题 10.03. Search Rotate Array LCCI
// Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. 
// You may assume that the array was originally sorted in increasing order. 
// If there are more than one target elements in the array, return the smallest index.

// Example1:
// Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
// Output: 8 (the index of 5 in the array)

// Example2:
// Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
// Output: -1 (not found)

// Note:
//     1 <= arr.length <= 1000000

import "fmt"

func search(arr []int, target int) int {
    // 如果arr[0] == target,使用二分法的时候会导致左右两边都有可能出现目标值，需要特殊处理才能解决，
    if arr[0] == target {
        return 0
    }
    left, right := 1, len(arr) - 1
    for left <= right {
        mid := left + (right-left)>>1
        if arr[mid] == target {
            right = mid - 1
        } else if arr[left] <= arr[mid] {
            // 左边为有序序列
            if target >= arr[left] && target < arr[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
                // 右边为有序序列
            if target > arr[mid] && target <= arr[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    if left < len(arr) && arr[left] == target {
        return left
    }
    return -1
}

func main() {
    // Example1:
    // Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
    // Output: 8 (the index of 5 in the array)
    fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5)) // 8
    // Example2:
    // Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
    // Output: -1 (not found)
    fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11)) // -1

    fmt.Println(search([]int{1,2,3,4,5,6,7,8,9}, 5)) // 4
    fmt.Println(search([]int{9,8,7,6,5,4,3,2,1}, 5)) // -1
}