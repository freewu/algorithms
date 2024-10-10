package main

// 1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
// Given an integer k, return the minimum number of Fibonacci numbers whose sum is equal to k. 
// The same Fibonacci number can be used multiple times.

// The Fibonacci numbers are defined as:
//     F1 = 1
//     F2 = 1
//     Fn = Fn-1 + Fn-2 for n > 2.

// It is guaranteed that for the given constraints we can always find such Fibonacci numbers that sum up to k.

// Example 1:
// Input: k = 7
// Output: 2 
// Explanation: The Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, ... 
// For k = 7 we can use 2 + 5 = 7.

// Example 2:
// Input: k = 10
// Output: 2 
// Explanation: For k = 10 we can use 2 + 8 = 10.

// Example 3:
// Input: k = 19
// Output: 3 

// Explanation: For k = 19 we can use 1 + 5 + 13 = 19.

// Constraints:
//     1 <= k <= 10^9

import "fmt"

func findMinFibonacciNumbers(k int) int {
    res, f1, f2 := 0, 1, 1
    arr := []int{f1, f2}
    for {
        if f1 + f2 > k { break }
        f1, f2 = f2, f1 + f2
        arr = append(arr, f2)
    }
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] <= k {
            res++
            k = k - arr[i]
        }
        if k == 0 { 
            return res 
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: k = 7
    // Output: 2 
    // Explanation: The Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, ... 
    // For k = 7 we can use 2 + 5 = 7.
    fmt.Println(findMinFibonacciNumbers(7)) // 2
    // Example 2:
    // Input: k = 10
    // Output: 2 
    // Explanation: For k = 10 we can use 2 + 8 = 10.
    fmt.Println(findMinFibonacciNumbers(10)) // 2
    // Example 3:
    // Input: k = 19
    // Output: 3 
    fmt.Println(findMinFibonacciNumbers(19)) // 3

    fmt.Println(findMinFibonacciNumbers(1)) // 1
    fmt.Println(findMinFibonacciNumbers(1024)) // 3
    fmt.Println(findMinFibonacciNumbers(100000000)) // 13
    fmt.Println(findMinFibonacciNumbers(99999999)) // 13
}