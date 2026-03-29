package main 

// 3881. Direction Assignments with Exactly K Visible People
// You are given three integers n, pos, and k.

// There are n people standing in a line indexed from 0 to n - 1. Each person independently chooses a direction:
//     1. 'L': visible only to people on their right
//     2. 'R': visible only to people on their left

// A person at index pos sees others as follows:
//     1. A person i < pos is visible if and only if they choose 'L'.
//     2. A person i > pos is visible if and only if they choose 'R'.

// Return the number of possible direction assignments such that the person at index pos sees exactly k people.

// Since the answer may be large, return it modulo 10^9 + 7.
 
// Example 1:
// Input: n = 3, pos = 1, k = 0
// Output: 2
// Explanation:​​​​​​​
// Index 0 is to the left of pos = 1, and index 2 is to the right of pos = 1.
// To see k = 0 people, index 0 must choose 'R' and index 2 must choose 'L', keeping both invisible.
// The person at index 1 can choose 'L' or 'R' since it does not affect the count. Thus, the answer is 2.

// Example 2:
// Input: n = 3, pos = 2, k = 1
// Output: 4
// Explanation:
// Index 0 and index 1 are left of pos = 2, and there is no index to the right.
// To see k = 1 person, exactly one of index 0 or index 1 must choose 'L', and the other must choose 'R'.
// There are 2 ways to choose which index is visible from the left.
// The person at index 2 can choose 'L' or 'R' since it does not affect the count. Thus, the answer is 2 + 2 = 4.

// Example 3:
// Input: n = 1, pos = 0, k = 0
// Output: 2
// Explanation:
// There are no indices to the left or right of pos = 0.
// To see k = 0 people, no additional condition is required.
// The person at index 0 can choose 'L' or 'R'. Thus, the answer is 2.
 
// Constraints:
//     1 <= n <= 10^5
//     0 <= pos, k <= n - 1

import "fmt"

const MOD = 1_000_000_007
const MX = 100_001

var fac [MX]int  // fac[i] = i!
var invF [MX]int // invF[i] = i!^-1 = pow(i!, mod-2)

func init() {
    fac[0] = 1
    for i := 1; i < MX; i++ {
        fac[i] = fac[i-1] * i % MOD
    }
    invF[MX-1] = pow(fac[MX-1], MOD-2)
    for i := MX - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % MOD
    }
}

func pow(x, n int) int {
    res := 1
    for ; n > 0; n /= 2 {
        if n % 2 > 0 {
            res = res * x % MOD
        }
        x = x * x % MOD
    }
    return res
}

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
    if m < 0 || m > n {
        return 0
    }
    return fac[n] * invF[m] % MOD * invF[n-m] % MOD
}

func countVisiblePeople(n, _, k int) int {
    return comb(n-1, k) * 2 % MOD
}

func countVisiblePeople1(n, _, k int) int {
    mx, mod := 100_001, 1_000_000_007
    fac, invF := make([]int, mx), make([]int, mx)  // fac[i] = i!,  invF[i] = i!^-1 = pow(i!, mod-2)
    pow := func (x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    init :=func () {
        fac[0] = 1
        for i := 1; i < mx; i++ {
            fac[i] = fac[i-1] * i % mod
        }
        invF[mx-1] = pow(fac[mx - 1], mod - 2)
        for i := mx - 1; i > 0; i-- {
            invF[i-1] = invF[i] * i % mod
        }
    }
    comb :=func (n, m int) int { // 从 n 个数中选 m 个数的方案数
        if m < 0 || m > n {
            return 0
        }
        return fac[n] * invF[m] % mod * invF[n-m] % mod
    }
    init()
    return comb(n-1, k) * 2 % mod
}

func main() {
    // Example 1:
    // Input: n = 3, pos = 1, k = 0
    // Output: 2
    // Explanation:​​​​​​​
    // Index 0 is to the left of pos = 1, and index 2 is to the right of pos = 1.
    // To see k = 0 people, index 0 must choose 'R' and index 2 must choose 'L', keeping both invisible.
    // The person at index 1 can choose 'L' or 'R' since it does not affect the count. Thus, the answer is 2.
    fmt.Println(countVisiblePeople(3, 1, 0)) // 2
    // Example 2:
    // Input: n = 3, pos = 2, k = 1
    // Output: 4
    // Explanation:
    // Index 0 and index 1 are left of pos = 2, and there is no index to the right.
    // To see k = 1 person, exactly one of index 0 or index 1 must choose 'L', and the other must choose 'R'.
    // There are 2 ways to choose which index is visible from the left.
    // The person at index 2 can choose 'L' or 'R' since it does not affect the count. Thus, the answer is 2 + 2 = 4.
    fmt.Println(countVisiblePeople(3, 2, 1)) // 4
    // Example 3:
    // Input: n = 1, pos = 0, k = 0
    // Output: 2
    // Explanation:
    // There are no indices to the left or right of pos = 0.
    // To see k = 0 people, no additional condition is required.
    // The person at index 0 can choose 'L' or 'R'. Thus, the answer is 2. 
    fmt.Println(countVisiblePeople(1, 0, 0)) // 2

    fmt.Println(countVisiblePeople(100_000, 0, 0)) // 2
    fmt.Println(countVisiblePeople(100_000, 99_999, 0)) // 2
    fmt.Println(countVisiblePeople(100_000, 0, 99_999)) // 2
    fmt.Println(countVisiblePeople(100_000, 99_999, 99_999)) // 2

    fmt.Println(countVisiblePeople1(3, 1, 0)) // 2
    fmt.Println(countVisiblePeople1(3, 2, 1)) // 4
    fmt.Println(countVisiblePeople1(1, 0, 0)) // 2
    fmt.Println(countVisiblePeople1(100_000, 0, 0)) // 2
    fmt.Println(countVisiblePeople1(100_000, 99_999, 0)) // 2
    fmt.Println(countVisiblePeople1(100_000, 0, 99_999)) // 2
    fmt.Println(countVisiblePeople1(100_000, 99_999, 99_999)) // 2
}