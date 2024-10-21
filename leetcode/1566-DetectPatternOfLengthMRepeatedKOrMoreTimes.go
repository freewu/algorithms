package main

// 1566. Detect Pattern of Length M Repeated K or More Times
// Given an array of positive integers arr, find a pattern of length m that is repeated k or more times.

// A pattern is a subarray (consecutive sub-sequence) that consists of one or more values, 
// repeated multiple times consecutively without overlapping. 
// A pattern is defined by its length and the number of repetitions.

// Return true if there exists a pattern of length m that is repeated k or more times, otherwise return false.

// Example 1:
// Input: arr = [1,2,4,4,4,4], m = 1, k = 3
// Output: true
// Explanation: The pattern (4) of length 1 is repeated 4 consecutive times. Notice that pattern can be repeated k or more times but not less.

// Example 2:
// Input: arr = [1,2,1,2,1,1,1,3], m = 2, k = 2
// Output: true
// Explanation: The pattern (1,2) of length 2 is repeated 2 consecutive times. Another valid pattern (2,1) is also repeated 2 times.

// Example 3:
// Input: arr = [1,2,1,2,1,3], m = 2, k = 3
// Output: false
// Explanation: The pattern (1,2) is of length 2 but is repeated only 2 times. There is no pattern of length 2 that is repeated 3 or more times.

// Constraints:
//     2 <= arr.length <= 100
//     1 <= arr[i] <= 100
//     1 <= m <= 100
//     2 <= k <= 100

import "fmt"

// brute force
func containsPattern(arr []int, m int, k int) bool {
    arrAreEqual := func (a, b []int) bool {
        for k, v := range a {
            if b[k] != v {
                return false
            }
        }
        return true
    }
    for i := 0; i <= len(arr) - (m * k); i++ {
        pattern, isValid := arr[i:i + m], true
        for j := i + m; j < i + (m * k); j += m {
            if !arrAreEqual(arr[j:j + m], pattern) {
                isValid = false
                break
            }
        }
        if isValid { 
            return true 
        }
    }
    return false
}

func containsPattern1(arr []int, m int, k int) bool {
    for start := 0; start < len(arr); start++ { // 模式起点
        i := 1
        for ; i < k; i++ {
            flag := true
            for j := 0; j < m; j++ {
                if start+j >= len(arr) || start+i*m+j >= len(arr) ||
                    arr[start+j] != arr[start+i*m+j] {
                    flag = false
                    break
                }
            }
            if !flag {
                break
            }
        }
        if i == k {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: arr = [1,2,4,4,4,4], m = 1, k = 3
    // Output: true
    // Explanation: The pattern (4) of length 1 is repeated 4 consecutive times. Notice that pattern can be repeated k or more times but not less.
    fmt.Println(containsPattern([]int{1,2,4,4,4,4}, 1, 3)) // true
    // Example 2:
    // Input: arr = [1,2,1,2,1,1,1,3], m = 2, k = 2
    // Output: true
    // Explanation: The pattern (1,2) of length 2 is repeated 2 consecutive times. Another valid pattern (2,1) is also repeated 2 times.
    fmt.Println(containsPattern([]int{1,2,1,2,1,1,1,3}, 2, 2)) // true
    // Example 3:
    // Input: arr = [1,2,1,2,1,3], m = 2, k = 3
    // Output: false
    // Explanation: The pattern (1,2) is of length 2 but is repeated only 2 times. There is no pattern of length 2 that is repeated 3 or more times.
    fmt.Println(containsPattern([]int{1,2,1,2,1,3}, 2, 3)) // false

    fmt.Println(containsPattern1([]int{1,2,4,4,4,4}, 1, 3)) // true
    fmt.Println(containsPattern1([]int{1,2,1,2,1,1,1,3}, 2, 2)) // true
    fmt.Println(containsPattern1([]int{1,2,1,2,1,3}, 2, 3)) // false
}