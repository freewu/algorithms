package main

// 1394. Find Lucky Integer in an Array
// Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

// Return the largest lucky integer in the array. If there is no lucky integer return -1.

// Example 1:
// Input: arr = [2,2,3,4]
// Output: 2
// Explanation: The only lucky number in the array is 2 because frequency[2] == 2.

// Example 2:
// Input: arr = [1,2,2,3,3,3]
// Output: 3
// Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.

// Example 3:
// Input: arr = [2,2,2,3,3]
// Output: -1
// Explanation: There are no lucky numbers in the array.

// Constraints:
//     1 <= arr.length <= 500
//     1 <= arr[i] <= 500

import "fmt"

func findLucky(arr []int) int {
    res, mp := -1, make(map[int]int)
    for _, v := range arr {
        mp[v]++
    }
    for k, v := range mp {
        if k == v && k > res { // 出现频次和它的数值大小相等
            res = k
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [2,2,3,4]
    // Output: 2
    // Explanation: The only lucky number in the array is 2 because frequency[2] == 2.
    fmt.Println(findLucky([]int{2,2,3,4})) // 2
    // Example 2:
    // Input: arr = [1,2,2,3,3,3]
    // Output: 3
    // Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
    fmt.Println(findLucky([]int{1,2,2,3,3,3})) // 3
    // Example 3:
    // Input: arr = [2,2,2,3,3]
    // Output: -1
    // Explanation: There are no lucky numbers in the array.
    fmt.Println(findLucky([]int{2,2,2,3,3})) // -1
}