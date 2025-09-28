package main

// 3697. Compute Decimal Representation
// You are given a positive integer n.

// A positive integer is a base-10 component if it is the product of a single digit from 1 to 9 and a non-negative power of 10. 
// For example, 500, 30, and 7 are base-10 components, while 537, 102, and 11 are not.

// Express n as a sum of only base-10 components, using the fewest base-10 components possible.

// Return an array containing these base-10 components in descending order.

// Example 1:
// Input: n = 537
// Output: [500,30,7]
// Explanation:
// We can express 537 as 500 + 30 + 7. It is impossible to express 537 as a sum using fewer than 3 base-10 components.

// Example 2:
// Input: n = 102
// Output: [100,2]
// Explanation:
// We can express 102 as 100 + 2. 102 is not a base-10 component, which means 2 base-10 components are needed.

// Example 3:
// Input: n = 6
// Output: [6]
// Explanation:
// 6 is a base-10 component.

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func decimalRepresentation(n int) []int {
    res := []int{}
    base := 1
    for n > 0 {
        d := n % 10
        n /= 10
        if d > 0 {
            res = append(res, d * base)
        }
        base *= 10
    }
    // 翻转数组
    for i, j := 0, len(res) - 1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 537
    // Output: [500,30,7]
    // Explanation:
    // We can express 537 as 500 + 30 + 7. It is impossible to express 537 as a sum using fewer than 3 base-10 components.
    fmt.Println(decimalRepresentation(537)) // [500,30,7]
    // Example 2:
    // Input: n = 102
    // Output: [100,2]
    // Explanation:
    // We can express 102 as 100 + 2. 102 is not a base-10 component, which means 2 base-10 components are needed.
    fmt.Println(decimalRepresentation(102)) // [100,2]
    // Example 3:
    // Input: n = 6
    // Output: [6]
    // Explanation:
    // 6 is a base-10 component.
    fmt.Println(decimalRepresentation(6)) // [6]

    fmt.Println(decimalRepresentation(1)) // [1]
    fmt.Println(decimalRepresentation(999)) // [900,90,9]
    fmt.Println(decimalRepresentation(1024)) // [1000,20,4]
    fmt.Println(decimalRepresentation(999_999_999)) // [900000000 90000000 9000000 900000 90000 9000 900 90 9]
    fmt.Println(decimalRepresentation(1_000_000_000)) // [1000000000]
}