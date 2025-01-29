package main

// 2749. Minimum Operations to Make the Integer Zero
// You are given two integers num1 and num2.

// In one operation, you can choose integer i in the range [0, 60] and subtract 2i + num2 from num1.

// Return the integer denoting the minimum number of operations needed to make num1 equal to 0.

// If it is impossible to make num1 equal to 0, return -1.

// Example 1:
// Input: num1 = 3, num2 = -2
// Output: 3
// Explanation: We can make 3 equal to 0 with the following operations:
// - We choose i = 2 and subtract 22 + (-2) from 3, 3 - (4 + (-2)) = 1.
// - We choose i = 2 and subtract 22 + (-2) from 1, 1 - (4 + (-2)) = -1.
// - We choose i = 0 and subtract 20 + (-2) from -1, (-1) - (1 + (-2)) = 0.
// It can be proven, that 3 is the minimum number of operations that we need to perform.

// Example 2:
// Input: num1 = 5, num2 = 7
// Output: -1
// Explanation: It can be proven, that it is impossible to make 5 equal to 0 with the given operation.

// Constraints:
//     1 <= num1 <= 10^9
//     -10^9 <= num2 <= 10^9

import "fmt"
import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
    for k := 1; ; k++ {
        x := num1 - k * num2
        if x < 0 {
            break
        }
        if bits.OnesCount(uint(x)) <= k && k <= x {
            return k
        }
    }
    return -1
}

func makeTheIntegerZero1(num1 int, num2 int) int {
    k := int64(1) // 0 operation is a valid output
    mx, power2 := int64(1 << 61), int64(num1) - int64(num2)
    for ; power2 < mx && power2 > 0; k++ {
        count, p2 := int64(0), power2
        for p2 > 0 {
            if p2 % 2 == 1 {
                count++
            }
            p2 = p2 / 2
        }
        if count <= k && power2 >= k {
            return int(k)
        }
        power2 -= int64(num2)
    }
    return -1
}

func main() {
    // Example 1:
    // Input: num1 = 3, num2 = -2
    // Output: 3
    // Explanation: We can make 3 equal to 0 with the following operations:
    // - We choose i = 2 and subtract 22 + (-2) from 3, 3 - (4 + (-2)) = 1.
    // - We choose i = 2 and subtract 22 + (-2) from 1, 1 - (4 + (-2)) = -1.
    // - We choose i = 0 and subtract 20 + (-2) from -1, (-1) - (1 + (-2)) = 0.
    // It can be proven, that 3 is the minimum number of operations that we need to perform.
    fmt.Println(makeTheIntegerZero(3, -2)) // 3
    // Example 2:
    // Input: num1 = 5, num2 = 7
    // Output: -1
    // Explanation: It can be proven, that it is impossible to make 5 equal to 0 with the given operation.
    fmt.Println(makeTheIntegerZero(5, 7)) // -1

    fmt.Println(makeTheIntegerZero(1, 1)) // -1
    fmt.Println(makeTheIntegerZero(1, -1_000_000_000)) // 13
    fmt.Println(makeTheIntegerZero(1_000_000_000, 1_000_000_000)) // -1
    fmt.Println(makeTheIntegerZero(1, 1_000_000_000)) // -1
    fmt.Println(makeTheIntegerZero(1_000_000_000, 1)) // 19
    fmt.Println(makeTheIntegerZero(1_000_000_000, -1_000_000_000)) // 12

    fmt.Println(makeTheIntegerZero1(3, -2)) // 3
    fmt.Println(makeTheIntegerZero1(5, 7)) // -1
    fmt.Println(makeTheIntegerZero1(1, 1)) // -1
    fmt.Println(makeTheIntegerZero1(1, -1_000_000_000)) // 13
    fmt.Println(makeTheIntegerZero1(1_000_000_000, 1_000_000_000)) // -1
    fmt.Println(makeTheIntegerZero1(1, 1_000_000_000)) // -1
    fmt.Println(makeTheIntegerZero1(1_000_000_000, 1)) // 19
    fmt.Println(makeTheIntegerZero1(1_000_000_000, -1_000_000_000)) // 12
}