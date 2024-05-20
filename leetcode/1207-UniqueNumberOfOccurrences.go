package main

// 1207. Unique Number of Occurrences
// Given an array of integers arr, 
// return true if the number of occurrences of each value in the array is unique or false otherwise.

// Example 1:
// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: 
// The value 1 has 3 occurrences, 2 has 2 and 3 has 1. 
// No two values have the same number of occurrences.

// Example 2:
// Input: arr = [1,2]
// Output: false

// Example 3:
// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
// Output: true
 
// Constraints:
//     1 <= arr.length <= 1000
//     -1000 <= arr[i] <= 1000

import "fmt"

func uniqueOccurrences(arr []int) bool {
    m := make(map[int]int)
    for _,v := range arr { // 累加每个数值的出现次数
        m[v]++ 
    }
    c := make(map[int]bool)
    for _,v := range m { 
        if _, ok := c[v]; ok { // 出现相同次数
            return false
        }
        c[v] = true
    }
    return true
}

func main() {
    // Example 1:
    // Input: arr = [1,2,2,1,1,3]
    // Output: true
    // Explanation: 
    // The value 1 has 3 occurrences, 2 has 2 and 3 has 1. 
    // No two values have the same number of occurrences.
    fmt.Println(uniqueOccurrences([]int{1,2,2,1,1,3})) // true
    // Example 2:
    // Input: arr = [1,2]
    // Output: false
    fmt.Println(uniqueOccurrences([]int{1,2})) // false
    // Example 3:
    // Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
    // Output: true
    fmt.Println(uniqueOccurrences([]int{-3,0,1,-3,1,1,1,-3,10,0})) // true
}