package main

// 1346. Check If N and Its Double Exist
// Given an array arr of integers, check if there exist two indices i and j such that :
//     i != j
//     0 <= i, j < arr.length
//     arr[i] == 2 * arr[j]
    
// Example 1:
// Input: arr = [10,2,5,3]
// Output: true
// Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

// Example 2:
// Input: arr = [3,1,7,11]
// Output: false
// Explanation: There is no i and j that satisfy the conditions.

// Constraints:
//     2 <= arr.length <= 500
//     -10^3 <= arr[i] <= 10^3

import "fmt"

func checkIfExist(arr []int) bool {
    mp, zero := make(map[int]bool), 0
    for _, v := range arr {
        if v == 0 { zero++ }
        mp[v] = true
    }
    if zero >= 2 { return true } // 处理有多个 0 的情况
    for _, v := range arr {
        if v == 0 { continue }
        if mp[2 * v] { return true }
    }
    return false
}

func checkIfExist1(arr []int) bool {
    mp := make(map[int]bool)
    for _, v := range arr {
        if mp[v*2] || v % 2 == 0 && mp[v / 2] {
            return true
        }
        mp[v] = true
    }
    return false
}

func main() {
    // Example 1:
    // Input: arr = [10,2,5,3]
    // Output: true
    // Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
    fmt.Println(checkIfExist([]int{10,2,5,3})) // true
    // Example 2:
    // Input: arr = [3,1,7,11]
    // Output: false
    // Explanation: There is no i and j that satisfy the conditions.
    fmt.Println(checkIfExist([]int{3,1,7,11})) // false

    fmt.Println(checkIfExist([]int{-2,0,10,-19,4,6,-8})) // false
    fmt.Println(checkIfExist([]int{0, 0})) // true

    fmt.Println(checkIfExist1([]int{10,2,5,3})) // true
    fmt.Println(checkIfExist1([]int{3,1,7,11})) // false
    fmt.Println(checkIfExist1([]int{-2,0,10,-19,4,6,-8})) // false
    fmt.Println(checkIfExist1([]int{0, 0})) // true
}