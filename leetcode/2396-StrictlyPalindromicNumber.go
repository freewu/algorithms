package main

// 2396. Strictly Palindromic Number
// An integer n is strictly palindromic if, for every base b between 2 and n - 2 (inclusive), the string representation of the integer n in base b is palindromic.

// Given an integer n, return true if n is strictly palindromic and false otherwise.

// A string is palindromic if it reads the same forward and backward.

// Example 1:
// Input: n = 9
// Output: false
// Explanation: In base 2: 9 = 1001 (base 2), which is palindromic.
// In base 3: 9 = 100 (base 3), which is not palindromic.
// Therefore, 9 is not strictly palindromic so we return false.
// Note that in bases 4, 5, 6, and 7, n = 9 is also not palindromic.

// Example 2:
// Input: n = 4
// Output: false
// Explanation: We only consider base 2: 4 = 100 (base 2), which is not palindromic.
// Therefore, we return false.

// Constraints:
//     4 <= n <= 10^5

import "fmt"
import "strconv"

func isStrictlyPalindromic(n int) bool {
    reverse := func(s string) string {
        arr := []byte(s) // convert to byte
        for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
        return string(arr)
    }
    for i := 2; i <= n - 2; i++ {
        num := strconv.FormatInt(int64(n), i)
        if num != reverse(num) {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: n = 9
    // Output: false
    // Explanation: In base 2: 9 = 1001 (base 2), which is palindromic.
    // In base 3: 9 = 100 (base 3), which is not palindromic.
    // Therefore, 9 is not strictly palindromic so we return false.
    // Note that in bases 4, 5, 6, and 7, n = 9 is also not palindromic.
    fmt.Println(isStrictlyPalindromic(9)) // false
    // Example 2:
    // Input: n = 4
    // Output: false
    // Explanation: We only consider base 2: 4 = 100 (base 2), which is not palindromic.
    // Therefore, we return false.
    fmt.Println(isStrictlyPalindromic(4)) // false

    fmt.Println(isStrictlyPalindromic(11)) // false
    fmt.Println(isStrictlyPalindromic(15)) // false
    fmt.Println(isStrictlyPalindromic(16)) // false
    fmt.Println(isStrictlyPalindromic(121)) // false
    fmt.Println(isStrictlyPalindromic(99999)) // false
    fmt.Println(isStrictlyPalindromic(1000000)) // false
}