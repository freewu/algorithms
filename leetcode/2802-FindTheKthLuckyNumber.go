package main

// 2802. Find The K-th Lucky Number
// We know that 4 and 7 are lucky digits. Also, a number is called lucky if it contains only lucky digits.

// You are given an integer k, return the kth lucky number represented as a string.

// Example 1:
// Input: k = 4
// Output: "47"
// Explanation: The first lucky number is 4, the second one is 7, the third one is 44 and the fourth one is 47.

// Example 2:
// Input: k = 10
// Output: "477"
// Explanation: Here are lucky numbers sorted in increasing order:
// 4, 7, 44, 47, 74, 77, 444, 447, 474, 477. So the 10th lucky number is 477.

// Example 3:
// Input: k = 1000
// Output: "777747447"
// Explanation: It can be shown that the 1000th lucky number is 777747447.

// Constraints:
//     1 <= k <= 10^9

import "fmt"

func kthLuckyNumber(k int) string {
    res, n := []byte{}, 1
    for k > 1 << n {
        k -= 1 << n
        n++
    }
    for n > 0 {
        n--
        if k <= 1 << n {
            res = append(res, '4')
        } else {
            res = append(res, '7')
            k -= 1 << n
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: k = 4
    // Output: "47"
    // Explanation: The first lucky number is 4, the second one is 7, the third one is 44 and the fourth one is 47.
    fmt.Println(kthLuckyNumber(4)) // 47
    // Example 2:
    // Input: k = 10
    // Output: "477"
    // Explanation: Here are lucky numbers sorted in increasing order:
    // 4, 7, 44, 47, 74, 77, 444, 447, 474, 477. So the 10th lucky number is 477.
    fmt.Println(kthLuckyNumber(10)) // 477
    // Example 3:
    // Input: k = 1000
    // Output: "777747447"
    // Explanation: It can be shown that the 1000th lucky number is 777747447.
    fmt.Println(kthLuckyNumber(1000)) // 777747447

    fmt.Println(kthLuckyNumber(1)) // 4
    fmt.Println(kthLuckyNumber(2)) // 7
    fmt.Println(kthLuckyNumber(1024)) // 4444444447
    fmt.Println(kthLuckyNumber(999_999_999)) // 77477744774747744747444444444
    fmt.Println(kthLuckyNumber(1_000_000_000)) // 77477744774747744747444444447
}