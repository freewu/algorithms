package main

// 3725. Count Ways to Choose Coprime Integers from Rows
// You are given a m x n matrix mat of positive integers.

// Return an integer denoting the number of ways to choose exactly one integer from each row of mat such that the greatest common divisor of all chosen integers is 1.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: mat = [[1,2],[3,4]]
// Output: 3
// Explanation:
// Chosen integer in the first row	Chosen integer in the second row	Greatest common divisor of chosen integers
// 1	3	1
// 1	4	1
// 2	3	1
// 2	4	2
// 3 of these combinations have a greatest common divisor of 1. Therefore, the answer is 3.

// Example 2:
// Input: mat = [[2,2],[2,2]]
// Output: 0
// Explanation:
// Every combination has a greatest common divisor of 2. Therefore, the answer is 0.

// Constraints:
//     1 <= m == mat.length <= 150
//     1 <= n == mat[i].length <= 150
//     1 <= mat[i][j] <= 150

import "fmt"
import "slices"

func countCoprime(mat [][]int) int {
    m, mod,mx := len(mat), 1_000_000_007, 0
    for _, row := range mat {
        mx = max(mx, slices.Max(row))
    }
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, mx+1)
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    var dfs func(i, g int) int
    dfs = func(i, g int) (res int) {
        if i < 0 {
            if g == 1 {
                return 1
            }
            return
        }
        p := &memo[i][g]
        if *p != -1 { // 之前计算过
            return *p
        }
        for _, x := range mat[i] {
            res += dfs(i-1, gcd(g, x))
        }
        res %= mod
        *p = res // 记忆化
        return
    }
    return dfs(m - 1, 0)
}

const MX = 151
var divisors [MX][]int

// 预处理每个数的因子的出现次数
func init() {
    for i := 1; i < MX; i++ {
        for j := i; j < MX; j += i {
            divisors[j] = append(divisors[j], i)
        }
    }
}

// 倍数容斥
func countCoprime1(mat [][]int) int {
    const mod = 1_000_000_007
    // 预处理每行的因子个数
    divisorCnt := make([][]int, len(mat))
    mx := 0
    for i, row := range mat {
        rowMax := slices.Max(row)
        mx = max(mx, rowMax)
        divisorCnt[i] = make([]int, rowMax + 1)
        for _, x := range row {
            for _, d := range divisors[x] {
                divisorCnt[i][d]++
            }
        }
    }
    cntGcd := make([]int, mx + 1)
    for i := mx; i > 0; i-- {
        // 每行选一个 i 的倍数的方案数
        res := 1
        for _, cnt := range divisorCnt {
            if i >= len(cnt) || cnt[i] == 0 {
                res = 0
                break
            }
            res = res * cnt[i] % mod // 乘法原理
        }
        for j := i; j <= mx; j += i {
            res -= cntGcd[j] // 注意这里有减法，可能导致 res 是负数
        }
        cntGcd[i] = res % mod
    }
    return (cntGcd[1] + mod) % mod // 保证结果非负
}

func main() {
    // Example 1:
    // Input: mat = [[1,2],[3,4]]
    // Output: 3
    // Explanation:
    // Chosen integer in the first row	Chosen integer in the second row	Greatest common divisor of chosen integers
    // 1	3	1
    // 1	4	1
    // 2	3	1
    // 2	4	2
    // 3 of these combinations have a greatest common divisor of 1. Therefore, the answer is 3.
    fmt.Println(countCoprime([][]int{{1,2},{3,4}})) // 3    
    // Example 2:
    // Input: mat = [[2,2],[2,2]]
    // Output: 0
    // Explanation:
    // Every combination has a greatest common divisor of 2. Therefore, the answer is 0.
    fmt.Println(countCoprime([][]int{{2,2},{2,2}})) // 0

    fmt.Println(countCoprime1([][]int{{1,2},{3,4}})) // 3    
    fmt.Println(countCoprime1([][]int{{2,2},{2,2}})) // 0
}