package main

// 2485. Find the Pivot Integer
// Given a positive integer n, find the pivot integer x such that:
//     The sum of all elements between 1 and x inclusively equals the sum of all elements between x and n inclusively.

// Return the pivot integer x. 
// If no such integer exists, return -1. 
// It is guaranteed that there will be at most one pivot index for the given input.

// Example 1:
// Input: n = 8
// Output: 6
// Explanation: 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.

// Example 2:
// Input: n = 1
// Output: 1
// Explanation: 1 is the pivot integer since: 1 = 1.

// Example 3:
// Input: n = 4
// Output: -1
// Explanation: It can be proved that no such integer exist.

// Constraints:
//     1 <= n <= 1000

import "fmt"
import "math"

// 高斯
func pivotInteger(n int) int {
    for  i := 1; i <= n; i += 1 {
        if (1+i)*(i-1+1) == (i+n)*(n-i+1) {
            return i
        }
    }
    return -1
}

// sqrt
func pivotInteger1(n int) int {
    m := n * (n + 1) / 2
    x := int(math.Sqrt(float64(m)))
    // return x * x == m ? x : -1;
    if x * x == m {
        return x
    }
    return -1
}

func pivotInteger2(n int) int {
    s := int(n * (n+1) / 2)
    for x := 1 ;x <= n; x++ {
        if x * (x+1) - x == s {
            return x
        }
    }
    return -1
}

func main() {
    // 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.
    fmt.Println(pivotInteger(8)) // 6
    // 1 is the pivot integer since: 1 = 1
    fmt.Println(pivotInteger(1)) // 1
    fmt.Println(pivotInteger(4)) // -1

    // 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.
    fmt.Println(pivotInteger1(8)) // 6
    // 1 is the pivot integer since: 1 = 1
    fmt.Println(pivotInteger1(1)) // 1
    fmt.Println(pivotInteger1(4)) // -1
    
    // 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.
    fmt.Println(pivotInteger2(8)) // 6
    // 1 is the pivot integer since: 1 = 1
    fmt.Println(pivotInteger2(1)) // 1
    fmt.Println(pivotInteger2(4)) // -1
}