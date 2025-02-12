package main

// 2954. Count the Number of Infection Sequences
// You are given an integer n and an array sick sorted in increasing order, 
// representing positions of infected people in a line of n people.

// At each step, one uninfected person adjacent to an infected person gets infected. 
// This process continues until everyone is infected.

// An infection sequence is the order in which uninfected people become infected, 
// excluding those initially infected.

// Return the number of different infection sequences possible, modulo 10^9+7.

// Example 1:
// Input: n = 5, sick = [0,4]
// Output: 4
// Explanation:
// There is a total of 6 different sequences overall.
// Valid infection sequences are [1,2,3], [1,3,2], [3,2,1] and [3,1,2].
// [2,3,1] and [2,1,3] are not valid infection sequences because the person at index 2 cannot be infected at the first step.

// Example 2:
// Input: n = 4, sick = [1]
// Output: 3
// Explanation:
// There is a total of 6 different sequences overall.
// Valid infection sequences are [0,2,3], [2,0,3] and [2,3,0].
// [3,2,0], [3,0,2], and [0,3,2] are not valid infection sequences because the infection starts at the person at index 1, then the order of infection is 2, then 3, and hence 3 cannot be infected earlier than 2.

// Constraints:
//     2 <= n <= 10^5
//     1 <= sick.length <= n - 1
//     0 <= sick[i] <= n - 1
//     sick is sorted in increasing order.

import "fmt"

// const MX = 100_000
// const MOD = 1_000_000_007

// var facts [MX + 1]int

// func init() {
//     facts[0] = 1
//     for i := 1; i <= MX; i++ {
//         facts[i] = facts[i-1] * i % MOD
//     }
// }

// func numberOfSequence(n int, sick []int) int {
//     m := len(sick)
//     nums := make([]int, m + 1)
//     nums[0], nums[m] = sick[0], n - sick[m-1] - 1
//     for i := 1; i < m; i++ {
//         nums[i] = sick[i] - sick[i-1] - 1
//     }
//     sum := 0
//     for _, v := range nums {
//         sum += v
//     }
//     pow := func(x, n int) int {
//         res := 1
//         for n > 0 {
//             if n & 1 == 1 {
//                 res = (res * x) % MOD
//             }
//             x = (x * x) % MOD
//             n >>= 1
//         }
//         return res
//     }
//     res := facts[sum]
//     for _, v := range nums {
//         if v > 0 {
//             res = res * pow(facts[v], MOD - 2) % MOD
//         }
//     }
//     for i := 1; i < len(nums) - 1; i++ {
//         if nums[i] > 1 {
//             res = res * pow(2, nums[i] - 1) % MOD
//         }
//     }
//     return res
// }

const mod = 1_000_000_007
const mx = 100_000

var fac, invFac [mx]int

func pow(x, n int) int {
    res := 1
    for ; n > 0; n /= 2 {
        if n%2 > 0 {
            res = res * x % mod
        }
        x = x * x % mod
    }
    return res
}

func init() {
    fac[0] = 1
    for i := 1; i < mx; i++ {
        fac[i] = fac[i-1] * i % mod
    }
    invFac[mx-1] = pow(fac[mx-1], mod-2)
    for i := mx - 1; i > 0; i-- {
        invFac[i-1] = invFac[i] * i % mod
    }
}

func numberOfSequence(n int, a []int) int {
    m := len(a)
    total := n - m
    comb := func(n, k int) int { return fac[n] * invFac[k] % mod * invFac[n-k] % mod }
    res := comb(total, a[0]) * comb(total-a[0], n-a[m-1]-1) % mod
    total -= a[0] + n - a[m-1] - 1
    e := 0
    for i := 1; i < m; i++ {
        k := a[i] - a[i-1] - 1
        if k > 0 {
            e += k - 1
            res = res * comb(total, k) % mod
            total -= k
        }
    }
    return res * pow(2, e) % mod
}

func main() {
    // Example 1:
    // Input: n = 5, sick = [0,4]
    // Output: 4
    // Explanation:
    // There is a total of 6 different sequences overall.
    // Valid infection sequences are [1,2,3], [1,3,2], [3,2,1] and [3,1,2].
    // [2,3,1] and [2,1,3] are not valid infection sequences because the person at index 2 cannot be infected at the first step.
    fmt.Println(numberOfSequence(5, []int{0,4})) // 4
    // Example 2:
    // Input: n = 4, sick = [1]
    // Output: 3
    // Explanation:
    // There is a total of 6 different sequences overall.
    // Valid infection sequences are [0,2,3], [2,0,3] and [2,3,0].
    // [3,2,0], [3,0,2], and [0,3,2] are not valid infection sequences because the infection starts at the person at index 1, then the order of infection is 2, then 3, and hence 3 cannot be infected earlier than 2.
    fmt.Println(numberOfSequence(4, []int{1})) // 3

    fmt.Println(numberOfSequence(10, []int{1,2,3,4,5,6,7,8,9})) // 1
}