package main

// 1362. Closest Divisors
// Given an integer num, find the closest two integers in absolute difference whose product equals num + 1 or num + 2.
// Return the two integers in any order.

// Example 1:
// Input: num = 8
// Output: [3,3]
// Explanation: For num + 1 = 9, the closest divisors are 3 & 3, for num + 2 = 10, the closest divisors are 2 & 5, hence 3 & 3 is chosen.

// Example 2:
// Input: num = 123
// Output: [5,25]

// Example 3:
// Input: num = 999
// Output: [40,25]

// Constraints:
//     1 <= num <= 10^9

import "fmt"
import "math"

func closestDivisors(num int) []int {
    helper := func(num int) []int {
        root := int(math.Sqrt(float64(num)))
        for i := root; i >= 1; i-- {
            if num % i == 0 {
                return []int{ i, num / i }
            }
        }
        return []int{}
    }
    v1, v2 := helper(num + 1), helper(num + 2)
    if v1[1] - v1[0] < v2[1] - v2[0] {
        return v1
    }
    return v2
}

func closestDivisors1(num int) []int {
    if num == 1 { return []int{1, 2} }
    v := num + 2
    i := int(math.Sqrt(float64(v)))
    for v % i > 1 {
        i-- // 当 i == 2 时必退出循环
    } // 这也要求 v >= 4, num >= 2
    return []int{ i, v / i }
}

func main() {
    // Example 1:
    // Input: num = 8
    // Output: [3,3]
    // Explanation: For num + 1 = 9, the closest divisors are 3 & 3, for num + 2 = 10, the closest divisors are 2 & 5, hence 3 & 3 is chosen.
    fmt.Println(closestDivisors(8)) // [3,3]
    // Example 2:
    // Input: num = 123
    // Output: [5,25]
    fmt.Println(closestDivisors(123)) // [5,25]
    // Example 3:
    // Input: num = 999
    // Output: [40,25]
    fmt.Println(closestDivisors(999)) // [40,25]

    fmt.Println(closestDivisors1(8)) // [3,3]
    fmt.Println(closestDivisors1(123)) // [5,25]
    fmt.Println(closestDivisors1(999)) // [40,25]
}