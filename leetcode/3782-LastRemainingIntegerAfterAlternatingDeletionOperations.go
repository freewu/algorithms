package main

// 3782. Last Remaining Integer After Alternating Deletion Operations
// You are given an integer n.

// We write the integers from 1 to n in a sequence from left to right. 
// Then, alternately apply the following two operations until only one integer remains, starting with operation 1:
//     1. Operation 1: Starting from the left, delete every second number.
//     2. Operation 2: Starting from the right, delete every second number.

// Return the last remaining integer.

// Example 1:
// Input: n = 8
// Output: 3
// Explanation:
// Write [1, 2, 3, 4, 5, 6, 7, 8] in a sequence.
// Starting from the left, we delete every second number: [1, 2, 3, 4, 5, 6, 7, 8]. The remaining integers are [1, 3, 5, 7].
// Starting from the right, we delete every second number: [1, 3, 5, 7]. The remaining integers are [3, 7].
// Starting from the left, we delete every second number: [3, 7]. The remaining integer is [3].

// Example 2:
// Input: n = 5
// Output: 1
// Explanation:
// Write [1, 2, 3, 4, 5] in a sequence.
// Starting from the left, we delete every second number: [1, 2, 3, 4, 5]. The remaining integers are [1, 3, 5].
// Starting from the right, we delete every second number: [1, 3, 5]. The remaining integers are [1, 5].
// Starting from the left, we delete every second number: [1, 5]. The remaining integer is [1].

// Example 3:
// Input: n = 1
// Output: 1
// Explanation:
// Write [1] in a sequence.
// The last remaining integer is 1.

// Constraints:
//     1 <= n <= 10^15

import "fmt"

func lastInteger(n int64) int64 {
    const mask = 0xAAAAAAAAAAAAAAA // ...1010
    return (n - 1) & mask + 1
}

func lastInteger1(n int64) int64 {
    if n <= 2 { return 1 }
    switch n & 3 {
    case 0:
        return 4 * lastInteger1(n / 4) - 1
    case 1:
        return 4 * lastInteger1((n + 3) / 4) - 3
    case 2:
        return 4 * lastInteger1((n + 2) / 4) - 3
    case 3:
        return 4 * lastInteger1((n + 1) / 4) - 1
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 8
    // Output: 3
    // Explanation:
    // Write [1, 2, 3, 4, 5, 6, 7, 8] in a sequence.
    // Starting from the left, we delete every second number: [1, 2, 3, 4, 5, 6, 7, 8]. The remaining integers are [1, 3, 5, 7].
    // Starting from the right, we delete every second number: [1, 3, 5, 7]. The remaining integers are [3, 7].
    // Starting from the left, we delete every second number: [3, 7]. The remaining integer is [3].
    fmt.Println(lastInteger(8)) // 3
    // Example 2:
    // Input: n = 5
    // Output: 1
    // Explanation:
    // Write [1, 2, 3, 4, 5] in a sequence.
    // Starting from the left, we delete every second number: [1, 2, 3, 4, 5]. The remaining integers are [1, 3, 5].
    // Starting from the right, we delete every second number: [1, 3, 5]. The remaining integers are [1, 5].
    // Starting from the left, we delete every second number: [1, 5]. The remaining integer is [1].
    fmt.Println(lastInteger(5)) // 1
    // Example 3:
    // Input: n = 1
    // Output: 1
    // Explanation:
    // Write [1] in a sequence.
    // The last remaining integer is 1.
    fmt.Println(lastInteger(1)) // 1

    fmt.Println(lastInteger(64)) // 43
    fmt.Println(lastInteger(99)) // 35
    fmt.Println(lastInteger(100)) // 35
    fmt.Println(lastInteger(999)) // 675
    fmt.Println(lastInteger(1000)) // 675
    fmt.Println(lastInteger(1001)) // 681
    fmt.Println(lastInteger(1024)) // 683
    fmt.Println(lastInteger(999_999_999_999_999)) // 712666616310443
    fmt.Println(lastInteger(1_000_000_000_000_000)) // 712666616310443

    fmt.Println(lastInteger1(8)) // 3
    fmt.Println(lastInteger1(5)) // 1
    fmt.Println(lastInteger1(1)) // 1
    fmt.Println(lastInteger1(64)) // 43
    fmt.Println(lastInteger1(99)) // 35
    fmt.Println(lastInteger1(100)) // 35
    fmt.Println(lastInteger1(999)) // 675
    fmt.Println(lastInteger1(1000)) // 675
    fmt.Println(lastInteger1(1001)) // 681
    fmt.Println(lastInteger1(1024)) // 683
    fmt.Println(lastInteger1(999_999_999_999_999)) // 712666616310443
    fmt.Println(lastInteger1(1_000_000_000_000_000)) // 712666616310443
}