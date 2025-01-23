package main

// 2571. Minimum Operations to Reduce an Integer to 0
// You are given a positive integer n, you can do the following operation any number of times:
//     Add or subtract a power of 2 from n.

// Return the minimum number of operations to make n equal to 0.

// A number x is power of 2 if x == 2i where i >= 0.

// Example 1:
// Input: n = 39
// Output: 3
// Explanation: We can do the following operations:
// - Add 20 = 1 to n, so now n = 40.
// - Subtract 23 = 8 from n, so now n = 32.
// - Subtract 25 = 32 from n, so now n = 0.
// It can be shown that 3 is the minimum number of operations we need to make n equal to 0.

// Example 2:
// Input: n = 54
// Output: 3
// Explanation: We can do the following operations:
// - Add 21 = 2 to n, so now n = 56.
// - Add 23 = 8 to n, so now n = 64.
// - Subtract 26 = 64 from n, so now n = 0.
// So the minimum number of operations is 3.

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func minOperations(n int) int {
    res := 0
    log2 := func(n int) int {
        res := 0
        for n > 1 {
            n >>= 1
            res++
        }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for n > 0 {
        res++
        pow2 := log2(n)
        n = min(n - (1 << pow2), (1 << (pow2 + 1)) - n)
    }
    return res
}

func minOperations1(n int) int {
    //lower_bit 
    res := 1
    for n & (n - 1) > 0 {
        bit := n & -n
        if n & (bit << 1) > 0 {
            n += bit
        } else {
            n -= bit
        }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 39
    // Output: 3
    // Explanation: We can do the following operations:
    // - Add 20 = 1 to n, so now n = 40.
    // - Subtract 23 = 8 from n, so now n = 32.
    // - Subtract 25 = 32 from n, so now n = 0.
    // It can be shown that 3 is the minimum number of operations we need to make n equal to 0.
    fmt.Println(minOperations(39)) // 3
    // Example 2:
    // Input: n = 54
    // Output: 3
    // Explanation: We can do the following operations:
    // - Add 21 = 2 to n, so now n = 56.
    // - Add 23 = 8 to n, so now n = 64.
    // - Subtract 26 = 64 from n, so now n = 0.
    // So the minimum number of operations is 3.
    fmt.Println(minOperations(54)) // 3

    fmt.Println(minOperations(1)) // 1
    fmt.Println(minOperations(2)) // 1
    fmt.Println(minOperations(3)) // 2
    fmt.Println(minOperations(8)) // 1
    fmt.Println(minOperations(64)) // 1
    fmt.Println(minOperations(100)) // 3
    fmt.Println(minOperations(999)) // 4
    fmt.Println(minOperations(1000)) // 3
    fmt.Println(minOperations(1024)) // 1
    fmt.Println(minOperations(9999)) // 5
    fmt.Println(minOperations(10000)) // 4

    fmt.Println(minOperations1(39)) // 3
    fmt.Println(minOperations1(54)) // 3
    fmt.Println(minOperations1(1)) // 1
    fmt.Println(minOperations1(2)) // 1
    fmt.Println(minOperations1(3)) // 2
    fmt.Println(minOperations1(8)) // 1
    fmt.Println(minOperations1(64)) // 1
    fmt.Println(minOperations1(100)) // 3
    fmt.Println(minOperations1(999)) // 4
    fmt.Println(minOperations1(1000)) // 3
    fmt.Println(minOperations1(1024)) // 1
    fmt.Println(minOperations1(9999)) // 5
    fmt.Println(minOperations1(10000)) // 4
}