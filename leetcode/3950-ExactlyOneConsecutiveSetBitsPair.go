package main

// 3950. Exactly One Consecutive Set Bits Pair
// You are given an integer n.

// Return true if its binary representation contains exactly one pair of consecutive set bits, and false otherwise.

// Example 1:
// Input: nums = 6
// Output: true
// Explanation:
// Binary representation of 6 is 110.
// There is exactly one pair of consecutive set bits ("11"). Thus, the answer is true​​​​​​​.

// Example 2:
// Input: nums = 5
// Output: false
// Explanation:
// Binary representation of 5 is 101.
// There are no consecutive set bits. Thus, the answer is false​​​​​​​.

// Constraints:
//     0 <= n <= 10^5

import "fmt"
import "strconv"

func consecutiveSetBits(n int) bool {
    binary := strconv.FormatInt(int64(n), 2)
    i, j, count := 0, 1, 0
    for i < j && j < len(binary) {
        if binary[i] == 49 && binary[j] == 49 {
            count++
        }
        i++
        j++
    }
    return count == 1
}

func consecutiveSetBits1(n int) bool {
    flag := 0
    for n != 0 {
        if (n & 3) == 3 {
            flag++
        }
        n /= 2
    }
    return flag == 1
}

func main() {
    // Example 1:
    // Input: nums = 6
    // Output: true
    // Explanation:
    // Binary representation of 6 is 110.
    // There is exactly one pair of consecutive set bits ("11"). Thus, the answer is true​​​​​​​.
    fmt.Println(consecutiveSetBits(6)) // true
    // Example 2:
    // Input: nums = 5
    // Output: false
    // Explanation:
    // Binary representation of 5 is 101.
    // There are no consecutive set bits. Thus, the answer is false​​​​​​​.
    fmt.Println(consecutiveSetBits(5)) // false

    fmt.Println(consecutiveSetBits(1)) // false
    fmt.Println(consecutiveSetBits(2)) // false
    fmt.Println(consecutiveSetBits(3)) // true
    fmt.Println(consecutiveSetBits(8)) // false
    fmt.Println(consecutiveSetBits(64)) // false
    fmt.Println(consecutiveSetBits(99)) // false
    fmt.Println(consecutiveSetBits(100)) // true
    fmt.Println(consecutiveSetBits(101)) // true
    fmt.Println(consecutiveSetBits(999)) // false
    fmt.Println(consecutiveSetBits(1000)) // false
    fmt.Println(consecutiveSetBits(1024)) // false
    fmt.Println(consecutiveSetBits(99_999)) // false
    fmt.Println(consecutiveSetBits(100_000)) // false

    fmt.Println(consecutiveSetBits1(6)) // true
    fmt.Println(consecutiveSetBits1(5)) // false
    fmt.Println(consecutiveSetBits1(1)) // false
    fmt.Println(consecutiveSetBits1(2)) // false
    fmt.Println(consecutiveSetBits1(3)) // true
    fmt.Println(consecutiveSetBits1(8)) // false
    fmt.Println(consecutiveSetBits1(64)) // false
    fmt.Println(consecutiveSetBits1(99)) // false
    fmt.Println(consecutiveSetBits1(100)) // true
    fmt.Println(consecutiveSetBits1(101)) // true
    fmt.Println(consecutiveSetBits1(999)) // false
    fmt.Println(consecutiveSetBits1(1000)) // false
    fmt.Println(consecutiveSetBits1(1024)) // false
    fmt.Println(consecutiveSetBits1(99_999)) // false
    fmt.Println(consecutiveSetBits1(100_000)) // false
}