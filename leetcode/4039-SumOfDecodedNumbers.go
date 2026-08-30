package main

// 4039. Sum of Decoded Numbers
// You are given an integer array nums.

// Each nums[i] is an encoded integer representing two positive integers xi and yi. To decode nums[i], define:
//     1. widthi = nums[i] % 10.
//     2. di = floor(nums[i] / 10).
//     3. xi as the integer formed by the first widthi digits of the decimal representation of di.
//     4. yi as the integer formed by all remaining digits of the decimal representation of di.

// It is guaranteed that the decimal representation of di contains more than widthi digits. 
// Therefore, both xi and yi contain at least one digit.

// The decoded value of nums[i] is xiyi.

// Return the sum of the decoded values of all elements in nums, modulo 109 + 7.

// The floor() function returns the integer part of the division.

// Example 1:
// Input: nums = [231]
// Output: 8
// Explanation:
// For 231, we have width = 1, d = 23, x = 2, and y = 3.
// The decoded value of 231 is 23 = 8.
// Since there is only one element in nums, the sum of the decoded values is 8.

// Example 2:
// Input: nums = [2522,2101]
// Output: 1649
// Explanation:
// For 2522, we have width = 2, d = 252, x = 25, and y = 2.
// The decoded value of 2522 is 252 = 625.
// For 2101, we have width = 1, d = 210, x = 2, and y = 10.
// The decoded value of 2101 is 210 = 1024.
// The sum of the decoded values is 625 + 1024 = 1649.

// Example 3:
// Input: nums = [2301]
// Output: 73741817
// Explanation:
// For 2301, we have width = 1, d = 230, x = 2, and y = 30.
// The decoded value is 230 = 1073741824.
// Therefore, the answer is 1073741824 modulo (109 + 7) = 73741817.

// Constraints:
//     1 <= nums.length <= 10^5
//     100 < nums[i] < 10^15
//     1 <= widthi <= 9
//     1 <= xi, yi < 10^9
//     The digit sequences used to form xi and yi do not have leading zeros.
//     It is guaranteed that every element in nums is a valid encoded integer.

import "fmt"
import "math"

func sumDecoded(nums []int64) int {
    res, mod:= 0, 1_000_000_007
    modPow := func(base, exp, mod int) int {
        res := 1
        base %= mod
        if base == 0 { return 0 }
        for ; exp > 0; exp >>= 1 {
            if (exp & 1) == 1 { 
                res = (res * base) % mod 
            }
            base = (base * base) % mod
        }
        return res
    }
    for _, n := range nums {
        x := int(n)
        d := x / 10
        // 计算 d 的十进制长度
        l := 0
        for v := d; v > 0; v /= 10 {
            l++
        }
        pow10 := int(math.Pow10(l - x % 10))
        // 根据 pow10 求出 x = d/pow10 和 y = d%pow10
        res += modPow(d / pow10, d % pow10, mod)
    }
    return res % mod
}

func sumDecoded1(nums []int64) int {
    const MOD = 1e9 + 7
    res := int64(0)
    fastPow := func(x, n int64) int64 {
        res := int64(1)
        for n > 0 {
            if n % 2 == 1 {
                res = (res * x) % MOD
            }
            x = (x * x) % MOD
            n /= 2
        }
        return res
    }
    for _, v := range nums {
        width, d := v % 10, v / 10
        count, copyD := int64(0), d
        for copyD > 0 {
            count++
            copyD /= 10
        }
        separateDigit := count - width
        x := d / fastPow(10, separateDigit)
        y := d % fastPow(10, separateDigit)
        res = (res + fastPow(x, y)) % MOD
    }
    return int(res)
}

func main() {
    // Example 1:
    // Input: nums = [231]
    // Output: 8
    // Explanation:
    // For 231, we have width = 1, d = 23, x = 2, and y = 3.
    // The decoded value of 231 is 23 = 8.
    // Since there is only one element in nums, the sum of the decoded values is 8.
    fmt.Println(sumDecoded([]int64{ 231 })) // 8 
    // Example 2:
    // Input: nums = [2522,2101]
    // Output: 1649
    // Explanation:
    // For 2522, we have width = 2, d = 252, x = 25, and y = 2.
    // The decoded value of 2522 is 252 = 625.
    // For 2101, we have width = 1, d = 210, x = 2, and y = 10.
    // The decoded value of 2101 is 210 = 1024.
    // The sum of the decoded values is 625 + 1024 = 1649.
    fmt.Println(sumDecoded([]int64{ 2522, 2101 })) // 1649
    // Example 3:
    // Input: nums = [2301]
    // Output: 73741817
    // Explanation:
    // For 2301, we have width = 1, d = 230, x = 2, and y = 30.
    // The decoded value is 230 = 1073741824.
    // Therefore, the answer is 1073741824 modulo (109 + 7) = 73741817.
    fmt.Println(sumDecoded([]int64{ 2301 })) // 73741817

    //fmt.Println(sumDecoded([]int64{ 123456789 })) // 
    fmt.Println(sumDecoded([]int64{ 987654321 })) // 448907077

    // fmt.Println(sumDecoded([]int64{1,2,3,4,5,6,7,8,9})) // 
    //fmt.Println(sumDecoded([]int64{9,8,7,6,5,4,3,2,1})) // 

    fmt.Println(sumDecoded1([]int64{ 231 })) // 8 
    fmt.Println(sumDecoded1([]int64{ 2522, 2101 })) // 1649
    fmt.Println(sumDecoded1([]int64{ 2301 })) // 73741817
    //fmt.Println(sumDecoded([]int64{ 123456789 })) // 
    fmt.Println(sumDecoded1([]int64{ 987654321 })) // 448907077
}