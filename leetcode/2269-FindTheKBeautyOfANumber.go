package main

// 2269. Find the K-Beauty of a Number
// The k-beauty of an integer num is defined as the number of substrings of num 
// when it is read as a string that meet the following conditions:
//     It has a length of k.
//     It is a divisor of num.

// Given integers num and k, return the k-beauty of num.

// Note:
//     Leading zeros are allowed.
//     0 is not a divisor of any value.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: num = 240, k = 2
// Output: 2
// Explanation: The following are the substrings of num of length k:
// - "24" from "240": 24 is a divisor of 240.
// - "40" from "240": 40 is a divisor of 240.
// Therefore, the k-beauty is 2.

// Example 2:
// Input: num = 430043, k = 2
// Output: 2
// Explanation: The following are the substrings of num of length k:
// - "43" from "430043": 43 is a divisor of 430043.
// - "30" from "430043": 30 is not a divisor of 430043.
// - "00" from "430043": 0 is not a divisor of 430043.
// - "04" from "430043": 4 is not a divisor of 430043.
// - "43" from "430043": 43 is a divisor of 430043.
// Therefore, the k-beauty is 2.

// Constraints:
//     1 <= num <= 10^9
//     1 <= k <= num.length (taking num as a string)

import "fmt"
import "strconv"

func divisorSubstrings(num int, k int) int {
    str := strconv.Itoa(num)
    res, n := 0, len(str)
    for i := 0; i <= n - k; i++ {
        sub := str[i:i+k]
        v, _ := strconv.Atoi(sub)
        if v != 0 && num % v == 0 {
            res += 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 240, k = 2
    // Output: 2
    // Explanation: The following are the substrings of num of length k:
    // - "24" from "240": 24 is a divisor of 240.
    // - "40" from "240": 40 is a divisor of 240.
    // Therefore, the k-beauty is 2.
    fmt.Println(divisorSubstrings(240, 2)) // 2
    // Example 2:
    // Input: num = 430043, k = 2
    // Output: 2
    // Explanation: The following are the substrings of num of length k:
    // - "43" from "430043": 43 is a divisor of 430043.
    // - "30" from "430043": 30 is not a divisor of 430043.
    // - "00" from "430043": 0 is not a divisor of 430043.
    // - "04" from "430043": 4 is not a divisor of 430043.
    // - "43" from "430043": 43 is a divisor of 430043.
    // Therefore, the k-beauty is 2.
    fmt.Println(divisorSubstrings(430043, 2)) // 2

    fmt.Println(divisorSubstrings(1_000_000_000, 2)) // 1
    fmt.Println(divisorSubstrings(999_999_999, 2)) // 0
}