package main

// 1304. Find N Unique Integers Sum up to Zero
// Given an integer n, return any array containing n unique integers such that they add up to 0.

// Example 1:
// Input: n = 5
// Output: [-7,-1,1,3,4]
// Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].

// Example 2:
// Input: n = 3
// Output: [-1,0,1]

// Example 3:
// Input: n = 1
// Output: [0]

// Constraints:
//     1 <= n <= 1000

import "fmt"

func sumZero(n int) []int {
    res := make([]int, n)
    for num, i := n / 2, 0; i < n/2; i, num = i + 1, num - 1 {
        res[i] = -num
        res[n / 2 + i] = num
    }
    return res
}

func sumZero1(n int) []int {
    res := []int{}
    for i := 0; i < n / 2; i++ {
        res = append(res, +(i + 1))
        res = append(res, -(i + 1))
    }
    if n % 2 == 1 { // 奇数补个0
        res = append(res, 0)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: [-7,-1,1,3,4]
    // Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
    fmt.Println(sumZero(5)) // [-2 -1 2 1 0]
    // Example 2:
    // Input: n = 3
    // Output: [-1,0,1]
    fmt.Println(sumZero(3)) // [-1 1 0]
    // Example 3:
    // Input: n = 1
    // Output: [0]
    fmt.Println(sumZero(1)) // [0]

    fmt.Println(sumZero(10)) // [-5 -4 -3 -2 -1 5 4 3 2 1]
    fmt.Println(sumZero(16)) // [-8 -7 -6 -5 -4 -3 -2 -1 8 7 6 5 4 3 2 1]
    fmt.Println(sumZero(32)) // [-16 -15 -14 -13 -12 -11 -10 -9 -8 -7 -6 -5 -4 -3 -2 -1 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1]
    fmt.Println(sumZero(1000))

    fmt.Println(sumZero1(5)) // [-2 -1 2 1 0]
    fmt.Println(sumZero1(3)) // [-1 1 0]
    fmt.Println(sumZero1(1)) // [0]
    fmt.Println(sumZero1(10)) // [-5 -4 -3 -2 -1 5 4 3 2 1]
    fmt.Println(sumZero1(16)) // [-8 -7 -6 -5 -4 -3 -2 -1 8 7 6 5 4 3 2 1]
    fmt.Println(sumZero1(32)) // [-16 -15 -14 -13 -12 -11 -10 -9 -8 -7 -6 -5 -4 -3 -2 -1 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1]
    fmt.Println(sumZero1(1000))
}