package main

// 868. Binary Gap
// Given a positive integer n, 
// find and return the longest distance between any two adjacent 1's in the binary representation of n. 
// If there are no two adjacent 1's, return 0.

// Two 1's are adjacent if there are only 0's separating them (possibly no 0's). 
// The distance between two 1's is the absolute difference between their bit positions. 
// For example, the two 1's in "1001" have a distance of 3.

// Example 1:
// Input: n = 22
// Output: 2
// Explanation: 22 in binary is "10110".
// The first adjacent pair of 1's is "10110" with a distance of 2.
// The second adjacent pair of 1's is "10110" with a distance of 1.
// The answer is the largest of these two distances, which is 2.
// Note that "10110" is not a valid pair since there is a 1 separating the two 1's underlined.

// Example 2:
// Input: n = 8
// Output: 0
// Explanation: 8 in binary is "1000".
// There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.

// Example 3:
// Input: n = 5
// Output: 2
// Explanation: 5 in binary is "101".

// Constraints:
//     1 <= n <= 10^9

import "fmt"
import "math/bits"
import "strconv"

func binaryGap(n int) int {
    res, s := 0, ""
    for n != 0 { // 拼成2进制字符串
        if n % 2 == 1 {
            s += "1"
            n -= 1
        } else {
            s += "0"
        }
        n /= 2
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            if s[i] != '1' || s[j] != '1' {
                continue
            }
            fine := true
            for k := i + 1; k < j; k++ {
                if s[k] == '1' {
                    fine = false
                    break
                }
            }
            if fine {
                res = max(res, j - i)
            }
        }
    }
    return res
}

func binaryGap1(n int) int {
    res, last := 0, -1
    for tmp := n; tmp > 0; {
        lb := tmp & (-tmp)
        tmp ^= lb
        lbIdx := bits.TrailingZeros(uint(lb))
        if last != -1 {
            res = max(res, lbIdx - last) // 
        }
        last = lbIdx 
    }
    return res
}

func binaryGap2(n int) int {
    prev, mx := -1, 0
    str := strconv.FormatInt(int64(n), 2)
    for k, v := range str {
        if v == '1' {
            if prev != -1 {
                if k - prev > mx { // 间隔比之前的大
                    mx = k - prev
                }
            }
            prev = k
        }
    }
    return mx
}

func binaryGap3(n int) int {
    res, count, curr  := 0, 0, 1
    for n > 0 {
        if n % 2 == 0 && count != 0 {
            curr++
        } else if n % 2 == 1 {
            if curr > res { res = curr }
            curr = 1
            count++
        }
        n /= 2
    }
    if count < 2 {
        return 0
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 22
    // Output: 2
    // Explanation: 22 in binary is "10110".
    // The first adjacent pair of 1's is "10110" with a distance of 2.
    // The second adjacent pair of 1's is "10110" with a distance of 1.
    // The answer is the largest of these two distances, which is 2.
    // Note that "10110" is not a valid pair since there is a 1 separating the two 1's underlined.
    fmt.Println(binaryGap(22)) // 2  10110
    // Example 2:
    // Input: n = 8
    // Output: 0
    // Explanation: 8 in binary is "1000".
    // There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.
    fmt.Println(binaryGap(8)) // 0  1000
    // Example 3:
    // Input: n = 5
    // Output: 2
    // Explanation: 5 in binary is "101".
    fmt.Println(binaryGap(5)) // 2  101

    fmt.Println(binaryGap(1)) // 0
    fmt.Println(binaryGap(64)) // 0
    fmt.Println(binaryGap(99)) // 4
    fmt.Println(binaryGap(100)) // 3
    fmt.Println(binaryGap(1024)) // 0
    fmt.Println(binaryGap(1025)) // 10
    fmt.Println(binaryGap(999_999_999)) // 3
    fmt.Println(binaryGap(1_000_000_000)) // 3

    fmt.Println(binaryGap1(22)) // 2  10110
    fmt.Println(binaryGap1(8)) // 0  1000
    fmt.Println(binaryGap1(5)) // 2  101
    fmt.Println(binaryGap1(1)) // 0
    fmt.Println(binaryGap1(64)) // 0
    fmt.Println(binaryGap1(99)) // 4
    fmt.Println(binaryGap1(100)) // 3
    fmt.Println(binaryGap1(1024)) // 0
    fmt.Println(binaryGap1(1025)) // 10
    fmt.Println(binaryGap1(999_999_999)) // 3
    fmt.Println(binaryGap1(1_000_000_000)) // 3

    fmt.Println(binaryGap2(22)) // 2  10110
    fmt.Println(binaryGap2(8)) // 0  1000
    fmt.Println(binaryGap2(5)) // 2  101
    fmt.Println(binaryGap2(1)) // 0
    fmt.Println(binaryGap2(64)) // 0
    fmt.Println(binaryGap2(99)) // 4
    fmt.Println(binaryGap2(100)) // 3
    fmt.Println(binaryGap2(1024)) // 0
    fmt.Println(binaryGap2(1025)) // 10
    fmt.Println(binaryGap2(999_999_999)) // 3
    fmt.Println(binaryGap2(1_000_000_000)) // 3

    fmt.Println(binaryGap3(22)) // 2  10110
    fmt.Println(binaryGap3(8)) // 0  1000
    fmt.Println(binaryGap3(5)) // 2  101
    fmt.Println(binaryGap3(1)) // 0
    fmt.Println(binaryGap3(64)) // 0
    fmt.Println(binaryGap3(99)) // 4
    fmt.Println(binaryGap3(100)) // 3
    fmt.Println(binaryGap3(1024)) // 0
    fmt.Println(binaryGap3(1025)) // 10
    fmt.Println(binaryGap3(999_999_999)) // 3
    fmt.Println(binaryGap3(1_000_000_000)) // 3
}