package main

// 4021. Minimum Operations to Make a Rotated Palindrome I
// You are given a string s consisting of lowercase English letters.

// You can perform the following operations any number of times (including zero) and in any order:
//     1. Increment: Choose any index i and replace s[i] with the next lowercase English letter. 
//        The letter after 'z' is 'a'.
//     2. Left rotate: Move the first character of the string to the end.

// Return the minimum number of operations required to make s a palindrome.

// A palindrome is a string that reads the same forward and backward.

// Example 1:
// Input: s = "abc"
// Output: 2
// Explanation:
// One optimal solution:
// Left rotate the string: "abc" -> "bca".
// Increment 'a' to 'b': "bca" -> "bcb".
// "bcb" is a palindrome. Thus, the answer is 2.

// Example 2:
// Input: s = "yb"
// Output: 3
// Explanation:
// Increment the first character three times: "yb" -> "zb" -> "ab" -> "bb".
// "bb" is a palindrome. Thus, the answer is 3.

// Constraints:
//     2 <= s.length <= 2000
//     s consists only of lowercase English letters.

import "fmt"

func minOperations(s string) int {
    res, n := 1 << 61, len(s)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = int(s[i] - 'a')
    }
    for k := 0; k < n; k++ {
        cost := k
        for i := 0; i < n/2; i++ {
            left, right := arr[(k + i) % n], arr[(k + n - 1 - i) % n]
            diff := left - right
            if diff < 0 {
                diff = -diff
            }
            if 26 - diff < diff {
                cost += 26 - diff
            } else {
                cost += diff
            }
        }
        if cost < res {
            res = cost
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: 2
    // Explanation:
    // One optimal solution:
    // Left rotate the string: "abc" -> "bca".
    // Increment 'a' to 'b': "bca" -> "bcb".
    // "bcb" is a palindrome. Thus, the answer is 2.
    fmt.Println(minOperations("abc")) // 2
    // Example 2:
    // Input: s = "yb"
    // Output: 3
    // Explanation:
    // Increment the first character three times: "yb" -> "zb" -> "ab" -> "bb".
    // "bb" is a palindrome. Thus, the answer is 3.
    fmt.Println(minOperations("yb")) // 3

    fmt.Println(minOperations("bluefrog")) // 12
    fmt.Println(minOperations("leetcode")) // 20
    fmt.Println(minOperations("freewu")) // 16
}