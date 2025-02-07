package main

// 2939. Maximum Xor Product
// Given three integers a, b, and n, return the maximum value of (a XOR x) * (b XOR x) where 0 <= x < 2n.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Note that XOR is the bitwise XOR operation.

// Example 1:
// Input: a = 12, b = 5, n = 4
// Output: 98
// Explanation: For x = 2, (a XOR x) = 14 and (b XOR x) = 7. Hence, (a XOR x) * (b XOR x) = 98. 
// It can be shown that 98 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.

// Example 2:
// Input: a = 6, b = 7 , n = 5
// Output: 930
// Explanation: For x = 25, (a XOR x) = 31 and (b XOR x) = 30. Hence, (a XOR x) * (b XOR x) = 930.
// It can be shown that 930 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.

// Example 3:
// Input: a = 1, b = 6, n = 3
// Output: 12
// Explanation: For x = 5, (a XOR x) = 4 and (b XOR x) = 3. Hence, (a XOR x) * (b XOR x) = 12.
// It can be shown that 12 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.

// Constraints:
//     0 <= a, b < 250
//     0 <= n <= 50

import "fmt"
import "math/bits"

func maximumXorProduct(a int64, b int64, n int) int {
    mask, diff, mod := (int64(1) << n) - 1, a ^ b, int64(1_000_000_007)
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    x := (^min(a, b) ^ int64(uint64(-diff) >> 63) << bits.Len64(uint64(diff >> 1))) & mask
    return int((a ^ x) % mod * ((b ^ x) % mod) % mod)
}

func maximumXorProduct1(a int64, b int64, n int) int {
    const mod = 1_000_000_007
    // max((a^x)*(x^b))
    // 思路：按照每一个bit,从大到小计算出x的值
    x, newa,newb := 0, 0,0
    for i := 50 ; i >= 0 ; i-- {
        if i <= n- 1 {
            if (a >> i) & 1 == 1 && (b >> i) & 1 == 1 {
                newa,newb = newa + 1 << i,newb + 1 << i
            } else if (a >> i) & 1 == 0 && (b >> i) & 1 == 0 {
                newa, newb, x = newa + 1 << i,newb + 1 << i, x + 1 << i
            } else if (a >> i) & 1 == 1 && (b >> i) & 1 == 0 {
                if newa > newb {
                    newa, newb, x = newa ,newb + 1 << i, x + 1 << i
                } else {
                    newa, newb = newa + 1 << i ,newb 
                }
            } else if (a >> i) & 1 == 0 && (b >> i) & 1 == 1 {
                if newa > newb {
                    newa, newb = newa ,newb + 1 << i
                } else {
                    newa, newb, x = newa + 1 << i, newb, x + 1 << i
                }
            }
        } else {
            if (a >> i) & 1 == 1 {
                newa = newa +  1 << i
            }
            if (b >> i) & 1 == 1 {
                newb = newb + 1 << i
            }
        }
    }
    // 为了避免溢出，按照位数依次计算
    a = (a ^ int64(x)) % mod
    b = (b ^ int64(x)) % mod
    return int ((a * b) % mod)
}

func maximumXorProduct2(a int64, b int64, n int) int {
    const mod = 1_000_000_007
    for i := n - 1; i >= 0; i-- {
        mask := int64(1) << i
        if a & mask == b & mask {
            a |= mask; b |= mask    // bit set
        } else {
            if a > b {
                a &= ^mask; b |= mask
            } else {
                a |= mask; b &= ^mask
            }
        }
    }
    return int((a % mod) * (b % mod) % mod)
}

func main() {
    // Example 1:
    // Input: a = 12, b = 5, n = 4
    // Output: 98
    // Explanation: For x = 2, (a XOR x) = 14 and (b XOR x) = 7. Hence, (a XOR x) * (b XOR x) = 98. 
    // It can be shown that 98 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.
    fmt.Println(maximumXorProduct(12, 5, 4)) // 98
    // Example 2:
    // Input: a = 6, b = 7 , n = 5
    // Output: 930
    // Explanation: For x = 25, (a XOR x) = 31 and (b XOR x) = 30. Hence, (a XOR x) * (b XOR x) = 930.
    // It can be shown that 930 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.
    fmt.Println(maximumXorProduct(6, 7, 5)) // 930
    // Example 3:
    // Input: a = 1, b = 6, n = 3
    // Output: 12
    // Explanation: For x = 5, (a XOR x) = 4 and (b XOR x) = 3. Hence, (a XOR x) * (b XOR x) = 12.
    // It can be shown that 12 is the maximum value of (a XOR x) * (b XOR x) for all 0 <= x < 2n.
    fmt.Println(maximumXorProduct(1, 6, 3)) // 12

    fmt.Println(maximumXorProduct(0, 0, 0)) // 0
    fmt.Println(maximumXorProduct(250, 250, 50)) // 178448631
    fmt.Println(maximumXorProduct(0, 250, 50)) // 438133322
    fmt.Println(maximumXorProduct(250, 0, 50)) // 438133322
    fmt.Println(maximumXorProduct(250, 250, 0)) // 62500
    fmt.Println(maximumXorProduct(0, 0, 50)) // 178448631
    fmt.Println(maximumXorProduct(0, 250, 0)) // 0
    fmt.Println(maximumXorProduct(250, 0, 0)) // 0

    fmt.Println(maximumXorProduct1(12, 5, 4)) // 98
    fmt.Println(maximumXorProduct1(6, 7, 5)) // 930
    fmt.Println(maximumXorProduct1(1, 6, 3)) // 12
    fmt.Println(maximumXorProduct1(0, 0, 0)) // 0
    fmt.Println(maximumXorProduct1(250, 250, 50)) // 178448631
    fmt.Println(maximumXorProduct1(0, 250, 50)) // 438133322
    fmt.Println(maximumXorProduct1(250, 0, 50)) // 438133322
    fmt.Println(maximumXorProduct1(250, 250, 0)) // 62500
    fmt.Println(maximumXorProduct1(0, 0, 50)) // 178448631
    fmt.Println(maximumXorProduct1(0, 250, 0)) // 0
    fmt.Println(maximumXorProduct1(250, 0, 0)) // 0

    fmt.Println(maximumXorProduct2(12, 5, 4)) // 98
    fmt.Println(maximumXorProduct2(6, 7, 5)) // 930
    fmt.Println(maximumXorProduct2(1, 6, 3)) // 12
    fmt.Println(maximumXorProduct2(0, 0, 0)) // 0
    fmt.Println(maximumXorProduct2(250, 250, 50)) // 178448631
    fmt.Println(maximumXorProduct2(0, 250, 50)) // 438133322
    fmt.Println(maximumXorProduct2(250, 0, 50)) // 438133322
    fmt.Println(maximumXorProduct2(250, 250, 0)) // 62500
    fmt.Println(maximumXorProduct2(0, 0, 50)) // 178448631
    fmt.Println(maximumXorProduct2(0, 250, 0)) // 0
    fmt.Println(maximumXorProduct2(250, 0, 0)) // 0
}