package main

// 1243. Array Transformation
// Given an initial array arr, every day you produce a new array using the array of the previous day.

// On the i-th day, you do the following operations on the array of day i-1 to produce the array of day i:
//     1. If an element is smaller than both its left neighbor and its right neighbor, then this element is incremented.
//     2. If an element is bigger than both its left neighbor and its right neighbor, then this element is decremented.
//     3. The first and last elements never change.

// After some days, the array does not change. Return that final array.

// Example 1:
// Input: arr = [6,2,3,4]
// Output: [6,3,3,4]
// Explanation: 
// On the first day, the array is changed from [6,2,3,4] to [6,3,3,4].
// No more operations can be done to this array.

// Example 2:
// Input: arr = [1,6,3,4,3,5]
// Output: [1,4,4,4,4,5]
// Explanation: 
// On the first day, the array is changed from [1,6,3,4,3,5] to [1,5,4,3,4,5].
// On the second day, the array is changed from [1,5,4,3,4,5] to [1,4,4,4,4,5].
// No more operations can be done to this array.

// Constraints:
//     3 <= arr.length <= 100
//     1 <= arr[i] <= 100

import "fmt"
import "reflect"

// func transformArray(arr []int) []int {
//     n, count := len(arr), 0 // 判断是否有变动
//     for true {
//         for i := 1; i < n - 1; i++ {
//             if arr[i] > arr[i - 1] && arr[i] > arr[i + 1] { // 假如一个元素大于它的左右邻居，那么该元素自减 1
//                 arr[i]--
//                 count++
//             } else if arr[i] < arr[i - 1] && arr[i] < arr[i + 1] { // 假如一个元素小于它的左右邻居，那么该元素自增 1
//                 arr[i]++
//                 count++
//             }
//         } 
//         if count == 0 {
//             break
//         }
//         count = 0
//     }
//     return arr
// }

func transformArray(arr []int) []int {
    if len(arr) < 3 {
        return arr
    }
    res := make([]int, len(arr))
    for !reflect.DeepEqual(res, arr) {
        copy(res,arr)
        for i := 1; i < len(arr)-1; i++ {
            if res[i] > res[i-1] && res[i] > res[i+1] { // 假如一个元素大于它的左右邻居，那么该元素自减 1
                arr[i]--
            }
            if res[i] < res[i-1] && res[i] < res[i+1] { // 假如一个元素小于它的左右邻居，那么该元素自增 1
                arr[i]++
            }
        }
    }
    return res
}

func transformArray1(arr []int) []int {
    ops, count := make([]int, len(arr)), 0
    for {
        count = 0
        for i := 1; i < len(arr) - 1; i++ {
            if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
                ops[i] = -1
                count++
            } else if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
                ops[i] = 1
                count++
            } else {
                ops[i] = 0
            }
        }
        for i := 1; i < len(arr) - 1; i++ {
            arr[i] += ops[i]
        }
        if count == 0 { // 数组没有变化了
            break
        }
    }
    return arr
}

func main() {
    // Example 1:
    // Input: arr = [6,2,3,4]
    // Output: [6,3,3,4]
    // Explanation: 
    // On the first day, the array is changed from [6,2,3,4] to [6,3,3,4].
    // No more operations can be done to this array.
    fmt.Println(transformArray([]int{6,2,3,4})) // [6,3,3,4]
    // Example 2:
    // Input: arr = [1,6,3,4,3,5]
    // Output: [1,4,4,4,4,5]
    // Explanation: 
    // On the first day, the array is changed from [1,6,3,4,3,5] to [1,5,4,3,4,5].
    // On the second day, the array is changed from [1,5,4,3,4,5] to [1,4,4,4,4,5].
    // No more operations can be done to this array.
    fmt.Println(transformArray([]int{1,6,3,4,3,5})) // [1,4,4,4,4,5]

    fmt.Println(transformArray([]int{2,1,2,1,1,2,2,1})) // [2,2,1,1,1,2,2,1]

    fmt.Println(transformArray1([]int{6,2,3,4})) // [6,3,3,4]
    fmt.Println(transformArray1([]int{1,6,3,4,3,5})) // [1,4,4,4,4,5]
    fmt.Println(transformArray1([]int{2,1,2,1,1,2,2,1})) // [2,2,1,1,1,2,2,1]
}