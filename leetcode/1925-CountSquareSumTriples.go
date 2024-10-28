package main

// 1925. Count Square Sum Triples
// A square triple (a,b,c) is a triple where a, b, and c are integers and a2 + b2 = c2.

// Given an integer n, return the number of square triples such that 1 <= a, b, c <= n.

// Example 1:
// Input: n = 5
// Output: 2
// Explanation: The square triples are (3,4,5) and (4,3,5).

// Example 2:
// Input: n = 10
// Output: 4
// Explanation: The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).

// Constraints:
//     1 <= n <= 250

import "fmt"
import "math"

func countTriples(n int) int {
    res := 0
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            v := i * i + j * j
            k := int(math.Sqrt(float64(v)))
            if k <= n && k * k == v {
                res++
            }
        }
    }
    return res
}

func countTriples1(n int) int {
    res := 0
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            for k := j + 1; k <= n;k++ {
                if  i*i + j*j == k*k {
                    res += 2
                } 
            }
        }
    }
    return res
}

func countTriples2(n int) int {
    res := 0
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            v := i * i + j * j
            k := int(math.Sqrt(float64(v)))
            if k <= n && k * k == v {
                res++
            }
        }
    }
    return res * 2
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 2
    // Explanation: The square triples are (3,4,5) and (4,3,5).
    fmt.Println(countTriples(5)) // 2
    // Example 2:
    // Input: n = 10
    // Output: 4
    // Explanation: The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).
    fmt.Println(countTriples(10)) // 4

    fmt.Println(countTriples(1)) // 0
    fmt.Println(countTriples(16)) // 8
    fmt.Println(countTriples(32)) // 22
    fmt.Println(countTriples(64)) // 54
    fmt.Println(countTriples(100)) // 104
    fmt.Println(countTriples(249)) // 324
    fmt.Println(countTriples(250)) // 330

    fmt.Println(countTriples1(5)) // 2
    fmt.Println(countTriples1(10)) // 4
    fmt.Println(countTriples1(1)) // 0
    fmt.Println(countTriples1(16)) // 8
    fmt.Println(countTriples1(32)) // 22
    fmt.Println(countTriples1(64)) // 54
    fmt.Println(countTriples1(100)) // 104
    fmt.Println(countTriples1(249)) // 324
    fmt.Println(countTriples1(250)) // 330

    fmt.Println(countTriples2(5)) // 2
    fmt.Println(countTriples2(10)) // 4
    fmt.Println(countTriples2(1)) // 0
    fmt.Println(countTriples2(16)) // 8
    fmt.Println(countTriples2(32)) // 22
    fmt.Println(countTriples2(64)) // 54
    fmt.Println(countTriples2(100)) // 104
    fmt.Println(countTriples2(249)) // 324
    fmt.Println(countTriples2(250)) // 330
}