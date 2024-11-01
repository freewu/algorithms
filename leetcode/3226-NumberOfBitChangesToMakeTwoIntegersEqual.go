package main

// 3226. Number of Bit Changes to Make Two Integers Equal
// You are given two positive integers n and k.

// You can choose any bit in the binary representation of n that is equal to 1 and change it to 0.

// Return the number of changes needed to make n equal to k. If it is impossible, return -1.

// Example 1:
// Input: n = 13, k = 4
// Output: 2
// Explanation:
// Initially, the binary representations of n and k are n = (1101)2 and k = (0100)2.
// We can change the first and fourth bits of n. The resulting integer is n = (0100)2 = k.

// Example 2:
// Input: n = 21, k = 21
// Output: 0
// Explanation:
// n and k are already equal, so no changes are needed.

// Example 3:
// Input: n = 14, k = 13
// Output: -1
// Explanation:
// It is not possible to make n equal to k.

// Constraints:
//     1 <= n, k <= 10^6

import "fmt"
import "math/bits"

func minChanges(n int, k int) int {
    res := 0
    for n > 0 || k > 0 {
        nmod, kmod := n % 2, k % 2
        if nmod == kmod {
            n, k = n / 2, k / 2
            continue
        }
        if nmod == 0 && kmod == 1 { return -1 }
        n, k = n / 2, k / 2
        res++
    }
    return res
}

func minChanges1(n int, k int) int {
    if n & k == k {
        return bits.OnesCount(uint(n ^ k))
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 13, k = 4
    // Output: 2
    // Explanation:
    // Initially, the binary representations of n and k are n = (1101)2 and k = (0100)2.
    // We can change the first and fourth bits of n. The resulting integer is n = (0100)2 = k.
    fmt.Println(minChanges(13, 4)) // 2
    // Example 2:
    // Input: n = 21, k = 21
    // Output: 0
    // Explanation:
    // n and k are already equal, so no changes are needed.
    fmt.Println(minChanges(21, 21)) // 0
    // Example 3:
    // Input: n = 14, k = 13
    // Output: -1
    // Explanation:
    // It is not possible to make n equal to k.
    fmt.Println(minChanges(14, 13)) // -1

    fmt.Println(minChanges(1, 1)) // 0
    fmt.Println(minChanges(100000, 100000)) // 0
    fmt.Println(minChanges(1, 100000)) // -1
    fmt.Println(minChanges(100000, 1)) // -1
    fmt.Println(minChanges(1, 99999)) // -1
    fmt.Println(minChanges(99999, 1)) // 9

    fmt.Println(minChanges1(13, 4)) // 2
    fmt.Println(minChanges1(21, 21)) // 0
    fmt.Println(minChanges1(14, 13)) // -1
    fmt.Println(minChanges1(1, 1)) // 0
    fmt.Println(minChanges1(100000, 100000)) // 0
    fmt.Println(minChanges1(1, 100000)) // -1
    fmt.Println(minChanges1(100000, 1)) // -1
    fmt.Println(minChanges1(1, 99999)) // -1
    fmt.Println(minChanges1(99999, 1)) // 9
}