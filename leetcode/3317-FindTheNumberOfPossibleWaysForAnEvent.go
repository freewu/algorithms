package main

// 3317. Find the Number of Possible Ways for an Event
// You are given three integers n, x, and y.

// An event is being held for n performers. 
// When a performer arrives, they are assigned to one of the x stages. 
// All performers assigned to the same stage will perform together as a band, though some stages might remain empty.

// After all performances are completed, the jury will award each band a score in the range [1, y].

// Return the total number of possible ways the event can take place.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that two events are considered to have been held differently if either of the following conditions is satisfied:
//     1. Any performer is assigned a different stage.
//     2. Any band is awarded a different score.

// Example 1:
// Input: n = 1, x = 2, y = 3
// Output: 6
// Explanation:
// There are 2 ways to assign a stage to the performer.
// The jury can award a score of either 1, 2, or 3 to the only band.

// Example 2:
// Input: n = 5, x = 2, y = 1
// Output: 32
// Explanation:
// Each performer will be assigned either stage 1 or stage 2.
// All bands will be awarded a score of 1.

// Example 3:
// Input: n = 3, x = 3, y = 4
// Output: 684

// Constraints:
//     1 <= n, x, y <= 1000

import "fmt"

func numberOfWays(n int, x int, y int) int {
    res, p, mod := 0, 1, 1_000_000_007
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, x + 1)
    }
    dp[0][0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= x; j++ {
            dp[i][j] = (dp[i-1][j] * j % mod + dp[i-1][j-1] * (x - (j - 1) ) % mod) % mod
        }
    }
    for j := 1; j <= x; j++ {
        p = p * y % mod
        res = (res + dp[n][j] * p % mod) % mod
    }
    return res
}

const mx = 1_000
const mod = 1_000_000_007

var fact [mx + 1]int
var invFact [mx + 1]int
var stirlingII [mx + 1][mx + 1]int

func init() {
    fact[0] = 1
    for i := 1; i < mx+1; i++ {
        fact[i] = fact[i-1] * i % mod
    }
    invFact[mx] = pow(fact[mx], mod-2)
    for i := mx; i > 0; i-- {
        invFact[i-1] = invFact[i] * i % mod
    }
    stirlingII[0][0] = 1
    for i := 1; i < mx+1; i++ {
        for j := 1; j < mx+1; j++ {
            stirlingII[i][j] = j * (stirlingII[i-1][j-1] + stirlingII[i-1][j]) % mod
        }
    }
}

func pow(x, n int) int {
    res := 1
    for ; n > 0; n >>= 1 {
        if n&1 == 1 {
            res = res * x % mod
        }
        x = x * x % mod
    }
    return res
}

func comb(n, m int) int {
    return fact[n] * invFact[m] % mod * invFact[n-m] % mod
}

func numberOfWays1(n int, x int, y int) int {
    res := 0
    for i := 0; i < x; i++ {
        res += comb(x, i+1) * stirlingII[n][i+1] % mod * pow(y, i+1) % mod
    }
    return res % mod
}

func main() {
    // Example 1:
    // Input: n = 1, x = 2, y = 3
    // Output: 6
    // Explanation:
    // There are 2 ways to assign a stage to the performer.
    // The jury can award a score of either 1, 2, or 3 to the only band.
    fmt.Println(numberOfWays(1, 2, 3)) // 6
    // Example 2:
    // Input: n = 5, x = 2, y = 1
    // Output: 32
    // Explanation:
    // Each performer will be assigned either stage 1 or stage 2.
    // All bands will be awarded a score of 1.
    fmt.Println(numberOfWays(5, 2, 1)) // 32
    // Example 3:
    // Input: n = 3, x = 3, y = 4
    // Output: 684
    fmt.Println(numberOfWays(3, 3, 4)) // 684

    fmt.Println(numberOfWays(1, 1, 1)) // 1
    fmt.Println(numberOfWays(1000, 1000, 1000)) // 295964505
    fmt.Println(numberOfWays(1, 1000, 1000)) // 1000000
    fmt.Println(numberOfWays(1000, 1, 1000)) // 1000
    fmt.Println(numberOfWays(1000, 1000, 1)) // 524700271
    fmt.Println(numberOfWays(1, 1, 1000)) // 1000
    fmt.Println(numberOfWays(1000, 1, 1)) // 1
    fmt.Println(numberOfWays(1, 1000, 1)) // 1000

    fmt.Println(numberOfWays1(1, 2, 3)) // 6
    fmt.Println(numberOfWays1(5, 2, 1)) // 32
    fmt.Println(numberOfWays1(3, 3, 4)) // 684
    fmt.Println(numberOfWays1(1, 1, 1)) // 1
    fmt.Println(numberOfWays1(1000, 1000, 1000)) // 295964505
    fmt.Println(numberOfWays1(1, 1000, 1000)) // 1000000
    fmt.Println(numberOfWays1(1000, 1, 1000)) // 1000
    fmt.Println(numberOfWays1(1000, 1000, 1)) // 524700271
    fmt.Println(numberOfWays1(1, 1, 1000)) // 1000
    fmt.Println(numberOfWays1(1000, 1, 1)) // 1
    fmt.Println(numberOfWays1(1, 1000, 1)) // 1000
}