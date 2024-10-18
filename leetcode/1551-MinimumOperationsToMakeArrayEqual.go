package main

// 1551. Minimum Operations to Make Array Equal
// You have an array arr of length n where arr[i] = (2 * i) + 1 for all valid values of i (i.e., 0 <= i < n).

// In one operation, you can select two indices x and y where 0 <= x, y < n and subtract 1 from arr[x] and add 1 to arr[y] (i.e., perform arr[x] -=1 and arr[y] += 1). 
// The goal is to make all the elements of the array equal. 
// It is guaranteed that all the elements of the array can be made equal using some operations.

// Given an integer n, the length of the array, 
// return the minimum number of operations needed to make all the elements of arr equal.

// Example 1:
// Input: n = 3
// Output: 2
// Explanation: arr = [1, 3, 5]
// First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
// In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].

// Example 2:
// Input: n = 6
// Output: 9

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func minOperations(n int) int {
    if n == 1 { return 0 }
    if n & 1 == 0 { return n * ((n - 2) / 2 + 1) / 2  } 
    return (n + 1) * ((n - 3) / 2 + 1) / 2
}

func minOperations1(n int) int {
    res, target := 0, 0
    if n % 2 != 0 {
        target = 2 * (n/2) + 1
    } else {
        target = 2*(n/2)
    }
    for i := 0; i < n / 2; i++ {
        res += (target - (2 * i + 1))
    }
    return res
}

func minOperations2(n int) int {
    // middle index = (n/2)
    // target = 2 * middle index
    // if n is odd, target = middle index + 1
    // target = n
    res, start := 0, 1
    for start < n {
        res += n - start
        start += 2
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 2
    // Explanation: arr = [1, 3, 5]
    // First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
    // In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].
    fmt.Println(minOperations(3)) // 2
    // Example 2:
    // Input: n = 6
    // Output: 9
    fmt.Println(minOperations(6)) // 9

    fmt.Println(minOperations(1)) // 0
    fmt.Println(minOperations(99)) // 2450
    fmt.Println(minOperations(100)) // 2500
    fmt.Println(minOperations(1024)) // 262144
    fmt.Println(minOperations(9999)) // 24995000
    fmt.Println(minOperations(10000)) // 25000000

    fmt.Println(minOperations1(3)) // 2
    fmt.Println(minOperations1(6)) // 9
    fmt.Println(minOperations1(1)) // 0
    fmt.Println(minOperations1(99)) // 2450
    fmt.Println(minOperations1(100)) // 2500
    fmt.Println(minOperations1(1024)) // 262144
    fmt.Println(minOperations1(9999)) // 24995000
    fmt.Println(minOperations1(10000)) // 25000000

    fmt.Println(minOperations2(3)) // 2
    fmt.Println(minOperations2(6)) // 9
    fmt.Println(minOperations2(1)) // 0
    fmt.Println(minOperations2(99)) // 2450
    fmt.Println(minOperations2(100)) // 2500
    fmt.Println(minOperations2(1024)) // 262144
    fmt.Println(minOperations2(9999)) // 24995000
    fmt.Println(minOperations2(10000)) // 25000000
}