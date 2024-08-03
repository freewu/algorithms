package main

// 1460. Make Two Arrays Equal by Reversing Subarrays
// You are given two integer arrays of equal length target and arr. 
// In one step, you can select any non-empty subarray of arr and reverse it. 
// You are allowed to make any number of steps.

// Return true if you can make arr equal to target or false otherwise.

// Example 1:
// Input: target = [1,2,3,4], arr = [2,4,1,3]
// Output: true
// Explanation: You can follow the next steps to convert arr to target:
// 1- Reverse subarray [2,4,1], arr becomes [1,4,2,3]
// 2- Reverse subarray [4,2], arr becomes [1,2,4,3]
// 3- Reverse subarray [4,3], arr becomes [1,2,3,4]
// There are multiple ways to convert arr to target, this is not the only way to do so.

// Example 2:
// Input: target = [7], arr = [7]
// Output: true
// Explanation: arr is equal to target without any reverses.

// Example 3:
// Input: target = [3,7,9], arr = [3,7,11]
// Output: false
// Explanation: arr does not have value 9 and it can never be converted to target.
 
// Constraints:
//     target.length == arr.length
//     1 <= target.length <= 1000
//     1 <= target[i] <= 1000
//     1 <= arr[i] <= 1000

import "fmt"
import "sort"

func canBeEqual(target []int, arr []int) bool {
    if len(target) != len(arr) {
        return false
    }
    sort.Ints(target)
    sort.Ints(arr)
    for i, v := range target { // 判断两个数组是否相等
        if v != arr[i] {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: target = [1,2,3,4], arr = [2,4,1,3]
    // Output: true
    // Explanation: You can follow the next steps to convert arr to target:
    // 1- Reverse subarray [2,4,1], arr becomes [1,4,2,3]
    // 2- Reverse subarray [4,2], arr becomes [1,2,4,3]
    // 3- Reverse subarray [4,3], arr becomes [1,2,3,4]
    // There are multiple ways to convert arr to target, this is not the only way to do so.
    fmt.Println(canBeEqual([]int{1,2,3,4},[]int{2,4,1,3})) // true
    // Example 2:
    // Input: target = [7], arr = [7]
    // Output: true
    // Explanation: arr is equal to target without any reverses.
    fmt.Println(canBeEqual([]int{7},[]int{7})) // true
    // Example 3:
    // Input: target = [3,7,9], arr = [3,7,11]
    // Output: false
    // Explanation: arr does not have value 9 and it can never be converted to target.
    fmt.Println(canBeEqual([]int{3,7,9},[]int{3,7,11})) // false
}