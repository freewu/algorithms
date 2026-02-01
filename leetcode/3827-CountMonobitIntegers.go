package main

// 3827. Count Monobit Integers
// You are given an integer n.

// An integer is called Monobit if all bits in its binary representation are the same.

// Return the count of Monobit integers in the range [0, n] (inclusive).

// Example 1:
// Input: n = 1
// Output: 2
// Explanation:
// The integers in the range [0, 1] have binary representations "0" and "1".
// Each representation consists of identical bits. Thus, the answer is 2.

// Example 2:
// Input: n = 4
// Output: 3
// Explanation:
// The integers in the range [0, 4] include binaries "0", "1", "10", "11", and "100".
// Only 0, 1 and 3 satisfy the Monobit condition. Thus, the answer is 3.
 
// Constraints:
//     0 <= n <= 1000

import "fmt"
import "math/bits"

func countMonobit(n int) int {
    return bits.Len(uint(n + 1))
}

func countMonobit1(n int) int {
    count := 0
    if n >= 0 {
        count = 1 
    }
    for k := 1; ; k++ {
        v := (1 << k) - 1
        if v > n { break }
        count++
    }
    return count
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 2
    // Explanation:
    // The integers in the range [0, 1] have binary representations "0" and "1".
    // Each representation consists of identical bits. Thus, the answer is 2.
    fmt.Println(countMonobit(1)) // 2
    // Example 2:
    // Input: n = 4
    // Output: 3
    // Explanation:
    // The integers in the range [0, 4] include binaries "0", "1", "10", "11", and "100".
    // Only 0, 1 and 3 satisfy the Monobit condition. Thus, the answer is 3.   
    fmt.Println(countMonobit(4)) // 3

    fmt.Println(countMonobit(0)) // 1
    fmt.Println(countMonobit(8)) // 4
    fmt.Println(countMonobit(64)) // 7
    fmt.Println(countMonobit(99)) // 7
    fmt.Println(countMonobit(100)) // 7
    fmt.Println(countMonobit(128)) // 8
    fmt.Println(countMonobit(999)) // 10
    fmt.Println(countMonobit(1000)) // 10
    fmt.Println(countMonobit(1024)) // 11

    fmt.Println(countMonobit1(1)) // 2
    fmt.Println(countMonobit1(4)) // 3
    fmt.Println(countMonobit1(0)) // 1
    fmt.Println(countMonobit1(8)) // 4
    fmt.Println(countMonobit1(64)) // 7
    fmt.Println(countMonobit1(99)) // 7
    fmt.Println(countMonobit1(100)) // 7
    fmt.Println(countMonobit1(128)) // 8
    fmt.Println(countMonobit1(999)) // 10
    fmt.Println(countMonobit1(1000)) // 10
    fmt.Println(countMonobit1(1024)) // 11
}