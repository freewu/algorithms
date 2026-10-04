package main

// 4071. Minimum Rotations to Dial a Number II
// You are given an integer n and a string s of length n consisting of digits.

// The dial contains the digits 0 through 9 in order and is circular, so 0 and 9 are adjacent. 
// The pointer initially points to 0.

// To dial each digit of s in order, rotate the pointer until it points to that digit. 
// Each rotation moves the pointer to an adjacent digit, and you may rotate in either direction. 
// Dialing a digit that the pointer already points to requires no rotations.

// Before dialing, you may perform the following operation at most once:
//     1. Choose an index k such that 0 <= k < n and reverse the suffix s[k..n - 1].

// Return the minimum total number of rotations needed to dial the string after optimally choosing whether to perform the operation and which suffix to reverse.

// Example 1:
// Input: n = 4, s = "1502"
// Output: 9
// Explanation:
// Reverse the suffix starting at k = 1 to obtain "1205", then dial it.
// Step	From	To	Rotations
// 1	0	1	1
// 2	1	2	1
// 3	2	0	2
// 4	0	5	5
// The total is 1 + 1 + 2 + 5 = 9, which is the minimum total number of rotations.

// Exa
// Input: n = 4, s = "2916"
// Output: 12
// Explanation:
// Choose not to reverse a suffix and dial "2916".
// Step	From	To	Rotations
// 1	0	2	2
// 2	2	9	3
// 3	9	1	2
// 4	1	6	5
// The total is 2 + 3 + 2 + 5 = 12, which is the minimum total number of rotations.

// Example 3:
// Input: n = 4, s = "4219"
// Output: 6
// Explanation:
// Reverse the suffix starting at k = 0, which reverses the entire string, to obtain "9124", then dial it.
// Step	From	To	Rotations
// 1	0	9	1
// 2	9	1	2
// 3	1	2	1
// 4	2	4	2
// The total is 1 + 2 + 1 + 2 = 6, which is the minimum total number of rotations.

// Constraints:
//     1 <= n == s.length <= 10^5​​​​​​​
//     s consists only of digits '0' to '9'

import "fmt"

func minRotations(n int, s string) int {
    pre, base, mn := byte('0'), 0, 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    dis := func(x, y byte) int { // 指针从数字 x 旋转到数字 y 的最少旋转次数
        d := abs(int(x) - int(y))
        return min(d, 10-d)
    }
    for _, v := range s {
        op := dis(pre, byte(v))
        base += op
        mn = min(mn, dis(pre, byte(s[n-1]))-op)
        pre = byte(v)
    }
    return base + mn
}

func minRotations1(n int, s string) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; } 
    dist := func(a, b int) int {
        d := abs(a-b)
        return min(d, 10-d)
    }
    sum := dist(0, int(s[0]-'0'))  
    for i := 1; i<n; i++ {
        sum += dist(int(s[i-1]-'0'), int(s[i]-'0'))
    }
    last := int(s[n-1] - '0')
    candidate := sum - dist(0, int(s[0]-'0')) + dist(0, last)
    res := min(sum, candidate)
    for k := 1; k < n; k++ {
        prev, curr := int(s[k-1] - '0'), int(s[k] - '0')
        candidate = sum - dist(prev, curr) + dist(prev, last)
        res = min(res, candidate)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, s = "1502"
    // Output: 9
    // Explanation:
    // Reverse the suffix starting at k = 1 to obtain "1205", then dial it.
    // Step	From	To	Rotations
    // 1	0	1	1
    // 2	1	2	1
    // 3	2	0	2
    // 4	0	5	5
    // The total is 1 + 1 + 2 + 5 = 9, which is the minimum total number of rotations.
    fmt.Println(minRotations(4, "1502")) // 9
    // Exa
    // Input: n = 4, s = "2916"
    // Output: 12
    // Explanation:
    // Choose not to reverse a suffix and dial "2916".
    // Step	From	To	Rotations
    // 1	0	2	2
    // 2	2	9	3
    // 3	9	1	2
    // 4	1	6	5
    // The total is 2 + 3 + 2 + 5 = 12, which is the minimum total number of rotations.
    fmt.Println(minRotations(4, "2916")) // 12
    // Example 3:
    // Input: n = 4, s = "4219"
    // Output: 6
    // Explanation:
    // Reverse the suffix starting at k = 0, which reverses the entire string, to obtain "9124", then dial it.
    // Step	From	To	Rotations
    // 1	0	9	1
    // 2	9	1	2
    // 3	1	2	1
    // 4	2	4	2
    // The total is 1 + 2 + 1 + 2 = 6, which is the minimum total number of rotations.
    fmt.Println(minRotations(4, "4219")) // 6

    fmt.Println(minRotations(10, "0123456789")) // 9
    fmt.Println(minRotations(10, "9876543210")) // 9

    fmt.Println(minRotations1(4, "1502")) // 9
    fmt.Println(minRotations1(4, "2916")) // 12
    fmt.Println(minRotations1(4, "4219")) // 6
    fmt.Println(minRotations1(10, "0123456789")) // 9
    fmt.Println(minRotations1(10, "9876543210")) // 9
}