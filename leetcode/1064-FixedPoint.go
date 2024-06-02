package main

// 1064. Fixed Point
// Given an array of distinct integers arr, where arr is sorted in ascending order, 
// return the smallest index i that satisfies arr[i] == i. If there is no such index, return -1.

// Example 1:
// Input: arr = [-10,-5,0,3,7]
// Output: 3
// Explanation: For the given array, arr[0] = -10, arr[1] = -5, arr[2] = 0, arr[3] = 3, thus the output is 3.

// Example 2:
// Input: arr = [0,2,5,8,17]
// Output: 0
// Explanation: arr[0] = 0, thus the output is 0.

// Example 3:
// Input: arr = [-10,-5,3,4,7,9]
// Output: -1
// Explanation: There is no such i that arr[i] == i, thus the output is -1.

// Constraints:
//     1 <= arr.length < 10^4
//     -10^9 <= arr[i] <= 10^9

// Follow up: The O(n) solution is very straightforward. Can we do better?

import "fmt"
import "sort"

func fixedPoint(arr []int) int {
    for i,v := range arr {
        if i == v { // 遇见第一个不动点(arr[i] == i) 则返回
            return i
        }
    }
    return -1 // 如果不存在这样的 i，返回 -1
}

func fixedPoint1(arr []int) int {
    index := sort.Search(len(arr), func(i int) bool {return arr[i] >= i})
    if index < len(arr) && arr[index] == index {
        return index
    }
    return -1
}

func main() {
    // Example 1:
    // Input: arr = [-10,-5,0,3,7]
    // Output: 3
    // Explanation: For the given array, arr[0] = -10, arr[1] = -5, arr[2] = 0, arr[3] = 3, thus the output is 3.
    fmt.Println(fixedPoint([]int{-10,-5,0,3,7})) // 3
    // Example 2:
    // Input: arr = [0,2,5,8,17]
    // Output: 0
    // Explanation: arr[0] = 0, thus the output is 0.
    fmt.Println(fixedPoint([]int{0,2,5,8,17})) // 0
    // Example 3:
    // Input: arr = [-10,-5,3,4,7,9]
    // Output: -1
    // Explanation: There is no such i that arr[i] == i, thus the output is -1.
    fmt.Println(fixedPoint([]int{-10,-5,3,4,7,9})) // -1

    fmt.Println(fixedPoint1([]int{-10,-5,0,3,7})) // 3
    fmt.Println(fixedPoint1([]int{0,2,5,8,17})) // 0
    fmt.Println(fixedPoint1([]int{-10,-5,3,4,7,9})) // -1
}