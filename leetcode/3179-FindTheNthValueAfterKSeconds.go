package main 

// 3179. Find the N-th Value After K Seconds
// You are given two integers n and k.

// Initially, you start with an array a of n integers where a[i] = 1 for all 0 <= i <= n - 1. 
// After each second, you simultaneously update each element to be the sum of all its preceding elements plus the element itself. 
// For example, after one second, a[0] remains the same, a[1] becomes a[0] + a[1], a[2] becomes a[0] + a[1] + a[2], and so on.

// Return the value of a[n - 1] after k seconds.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 4, k = 5
// Output: 56
// Explanation:
// Second	State After
// 0	[1,1,1,1]
// 1	[1,2,3,4]
// 2	[1,3,6,10]
// 3	[1,4,10,20]
// 4	[1,5,15,35]
// 5	[1,6,21,56]

// Example 2:
// Input: n = 5, k = 3
// Output: 35
// Explanation:
// Second	State After
// 0	[1,1,1,1,1]
// 1	[1,2,3,4,5]
// 2	[1,3,6,10,15]
// 3	[1,4,10,20,35]

// Constraints:
//     1 <= n, k <= 1000

import "fmt"

func valueAfterKSeconds(n int, k int) int {
    prefix := make([]int, n)
    for i := 0; i < n; i++ { // fill 1
        prefix[i] = 1
    }
    for i := 1; i <= k; i++ {
        for j := 1; j < n; j++ {
            prefix[j] = (prefix[j - 1] + prefix[j]) % 1_000_000_007
        }
    }
    return prefix[n - 1]
}

const mod = 1_000_000_007
const mx = 2000
var fac, invFac [mx + 1]int

func init() {
    fac[0] = 1
    for i := 1; i <= mx; i++ {
        fac[i] = fac[i-1] * i % mod
    }
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    invFac[mx] = pow(fac[mx], mod - 2)
    for i := mx; i > 0; i-- {
        invFac[i-1] = invFac[i] * i % mod
    }
}

func valueAfterKSeconds1(n, k int) int {
    return fac[n + k - 1] * invFac[n - 1] % mod * invFac[k] % mod
}

func main() {
    // Example 1:
    // Input: n = 4, k = 5
    // Output: 56
    // Explanation:
    // Second	State After
    // 0	[1,1,1,1]
    // 1	[1,2,3,4]
    // 2	[1,3,6,10]
    // 3	[1,4,10,20]
    // 4	[1,5,15,35]
    // 5	[1,6,21,56]
    fmt.Println(valueAfterKSeconds(4, 5)) // 56
    // Example 2:
    // Input: n = 5, k = 3
    // Output: 35
    // Explanation:
    // Second	State After
    // 0	[1,1,1,1,1]
    // 1	[1,2,3,4,5]
    // 2	[1,3,6,10,15]
    // 3	[1,4,10,20,35]
    fmt.Println(valueAfterKSeconds(5, 3)) // 35

    fmt.Println(valueAfterKSeconds(1, 1)) // 1
    fmt.Println(valueAfterKSeconds(1000, 1000)) // 36237869
    fmt.Println(valueAfterKSeconds(1, 1000)) // 1
    fmt.Println(valueAfterKSeconds(1000, 1)) // 1000

    fmt.Println(valueAfterKSeconds1(4, 5)) // 56
    fmt.Println(valueAfterKSeconds1(5, 3)) // 35
    fmt.Println(valueAfterKSeconds1(1, 1)) // 1
    fmt.Println(valueAfterKSeconds1(1000, 1000)) // 36237869
    fmt.Println(valueAfterKSeconds1(1, 1000)) // 1
    fmt.Println(valueAfterKSeconds1(1000, 1)) // 1000
}