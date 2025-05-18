package main

// 1931. Painting a Grid With Three Different Colors
// You are given two integers m and n. 
// Consider an m x n grid where each cell is initially white. 
// You can paint each cell red, green, or blue. 
// All cells must be painted.

// Return the number of ways to color the grid with no two adjacent cells having the same color. 
// Since the answer can be very large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/22/colorthegrid.png" />
// Input: m = 1, n = 1
// Output: 3
// Explanation: The three possible colorings are shown in the image above.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/22/copy-of-colorthegrid.png" />
// Input: m = 1, n = 2
// Output: 6
// Explanation: The six possible colorings are shown in the image above.

// Example 3:
// Input: m = 5, n = 5
// Output: 580986

// Constraints:
//     1 <= m <= 5
//     1 <= n <= 1000

import "fmt"
import "math"

// func colorTheGrid(m int, n int) int {
//     mod := 1_000_000_007
//     if m == 1 { return 3 * (int(math.Pow(float64(2), float64(n-1))) % mod) % mod }
//     if m == 2 { return 2 * (int(math.Pow(float64(3), float64(n))) % mod) % mod }
//     if m == 3 {
//         x0, x1 := 0, 3
//         for i := 0; i < n; i++ {
//             x0, x1 = (3*x0 + 2*x1) %mod, (2*x0 + 2*x1) %mod
//         }
//         return (x0 + x1) % mod
//     }
//     if m == 4 {
//         x0, x1, x2 := 0, 2, 2
//         for i := 0; i < n; i++ {
//             x0, x1, x2 = (3*x0 + 2*x1 + x2) %mod, (4*x0 + 4*x1 + 2*x2) %mod, ( x0 + x1 + 2*x2) % mod
//         }
//         return (x0+ x1 + x2) % mod
//     }
//     if m == 5 {
//         x0, x1, x2, x3, x4, x5, x6, x7 := 0, 0, 0, 0, 3, 0, 3, 0
//         for i := 0; i < n; i++ {
//             x0 = (3*x0 + 2*x1 + 2*x2 +   x3 +          x5 + 2*x6 + 2*x7) %mod
//             x1 = (2*x0 + 2*x1 + 2*x2 +   x3 +   x4 +   x5 +   x6 +   x7) %mod
//             x2 = (2*x0 + 2*x1 + 2*x2 +   x3 +          x5 + 2*x6 + 2*x7) %mod
//             x3 = (1*x0 + 1*x1 + 1*x2 + 2*x3 +   x4 +   x5 +   x6 +   x7) %mod
//             x4 = (     + 1*x1 +      +   x3 + 2*x4 +   x5 +          x7) %mod
//             x5 = (1*x0 + 1*x1 + 1*x2 +   x3 +   x4 + 2*x5 +   x6 +   x7) %mod
//             x6 = (2*x0 + 1*x1 + 2*x2 +   x3 +          x5 + 2*x6 +   x7) %mod
//             x7 = (2*x0 + 1*x1 + 2*x2 +   x3 +   x4 +   x5 +   x6 + 2*x7) %mod
//         }
//         return (x0 + x1 + x2 + x3 + x4 + x5 + x6 + x7) % mod
//     }
//     return -1
// }

// class Solution:
//     def colorTheGrid(self, m: int, n: int, mod = 1_000_000_007) -> int:

//         if m == 1:
//             return 3*pow(2, n-1, mod) %mod

//         if m == 2:
//             return 2*pow(3, n, mod) %mod

//         if m == 3:
//             x0, x1 = 0, 3

//             for _ in range(n):
//                 x0, x1 = ((3*x0 + 2*x1) %mod, 
//                           (2*x0 + 2*x1) %mod)

//             return (x0+x1) %mod   

//         if m == 4:
//             x0, x1, x2 = 0, 2, 2

//             for _ in range(n):
//                 x0, x1, x2 = ((3*x0 + 2*x1 +   x2) %mod,
//                               (4*x0 + 4*x1 + 2*x2) %mod,
//                               (  x0 +   x1 + 2*x2) %mod)

//             return (x0+ x1 + x2) %mod

//         if m == 5:
//             (x0, x1, x2, x3, x4, x5, x6, x7) = (0, 0, 0, 0, 3, 0, 3, 0)

//             for _ in range(n):

//                 (x0, x1, x2, x3, x4, x5, x6, x7) = (

//                     (3*x0 + 2*x1 + 2*x2 +   x3 +          x5 + 2*x6 + 2*x7) %mod,
//                     (2*x0 + 2*x1 + 2*x2 +   x3 +   x4 +   x5 +   x6 +   x7) %mod,
//                     (2*x0 + 2*x1 + 2*x2 +   x3 +          x5 + 2*x6 + 2*x7) %mod,
//                     (1*x0 + 1*x1 + 1*x2 + 2*x3 +   x4 +   x5 +   x6 +   x7) %mod,
//                     (     + 1*x1 +      +   x3 + 2*x4 +   x5 +          x7) %mod,
//                     (1*x0 + 1*x1 + 1*x2 +   x3 +   x4 + 2*x5 +   x6 +   x7) %mod,
//                     (2*x0 + 1*x1 + 2*x2 +   x3 +          x5 + 2*x6 +   x7) %mod,
//                     (2*x0 + 1*x1 + 2*x2 +   x3 +   x4 +   x5 +   x6 + 2*x7) %mod )

//             return (x0 + x1 + x2 + x3 + x4 + x5 + x6 + x7) %mod

func colorTheGrid(m int, n int) int {
    if n == 1 {
        return len(allType(m))
    }
    next := nextType(m)
    l := len(next)
    // 第一列所有长度为m的序列均为可行解
    pre := make([]int, l, l)
    for i := 0; i < l; i++ { 
        pre[i] = 1
    }
    // 后续列安装状态转移关系进行转移。
    cur := make([]int, l, l)
    for i := 1; i < n; i++ {
        for j := 0; j < l; j++ {
            cur[j] = 0
        }
        for j := 0; j < l; j++ {
            for k := 0; k < l; k++ {
                cur[j] += pre[k] * next[k][j] % (1e9+7)
            }
        }
        pre, cur = cur, pre
    }
    res := 0
    for i := 0; i < l; i++ {
        res += pre[i]
        res %= 1_000_000_007
    }
    return res
}

// 确定长度为m（每一列）的三色染色序列总共有多少种，保证同一列相邻网格不同色
func allType(m int) [][]int{
    res, t := [][]int{}, []int{}
    var bt func(i int)
    bt = func(d int) {
        if d == m {
            tt := make([]int, 0, m)
            tt = append(tt, t...)
            res = append(res, tt) 
        } else if d == 0 {
            for i := 0; i < 3; i++ {
                t = append(t, i)
                bt(d+1)
                t = t[:len(t)-1]
            }
        } else {
            for i := 0; i < 3; i++ {
                if i != t[len(t)-1] {
                    t = append(t, i)
                    bt(d+1)
                    t = t[:len(t)-1]
                }
            }
        }
    }
    bt(0)
    return res
}

// 确定相邻列之间的转移关系，转移后确保同行网格不同色
func nextType(m int) [][]int {
    at := allType(m)
    n := len(at)
    res := make([][]int, n, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n, n)
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            flag := 0
            for k := 0; k < m; k++ {
                if at[i][k] == at[j][k] {
                    flag = 1
                    break
                }
            }
            if flag == 0 {
                res[i][j], res[j][i] = 1, 1
            }
        }
    }
    return res
}

func colorTheGrid1(m int, n int) int{
    f1 := func(x int) bool {
        last := -1
        for i := 0; i < m; i++ {
            if x % 3 == last { return false }
            last = x % 3
            x /= 3
        }
        return true
    }
    f2 := func(x, y int) bool {
        for i := 0; i < m; i++ {
            if x % 3 == y % 3 { return false }
            x /= 3
            y /= 3
        }
        return true
    }
    mx := int(math.Pow(3, float64(m)))
    valid := map[int]bool{}
    f := make([]int, mx)
    for i := 0; i < mx; i++ {
        if f1(i) {
            valid[i] = true
            f[i] = 1
        }
    }
    d := map[int][]int{}
    for i := range valid {
        for j := range valid {
            if f2(i, j) {
                d[i] = append(d[i], j)
            }
        }
    }
    res, mod := 0, 1_000_000_007
    for k := 1; k < n; k++ {
        g := make([]int, mx)
        for i := range valid {
            for _, j := range d[i] {
                g[i] = (g[i] + f[j]) % mod
            }
        }
        f = g
    }
    for _, x := range f {
        res = (res + x) % mod
    }
    return res
}

const mod = 1e9 + 7
var ways [1024][1024]int
var rowmasks [6][]int
var pow3 = []int{1, 3, 9, 27, 81, 243}

func init() {
    for m := 1; m <= 5; m++ {
    nextmask: for mask := 0; mask < pow3[m]; mask++ {
        bm, prev := 0, -1
        for i, t := 0, mask; i < m; i, t = i+1, t/3 {
            if t%3 == prev { continue nextmask }
            bm, prev = bm*4 + t % 3 + 1, t % 3
        }
        rowmasks[m] = append(rowmasks[m], bm)
        ways[1][bm] = 1
    }
    for n := 2; n <= 1000; n++ {
        for _, mask := range rowmasks[m] {
            for _, prevmask := range rowmasks[m] {
                ok := true
                for i := 0; ok && i < m; i++ {
                    ok = mask >> (2*i) & 3 != prevmask >> (2*i) & 3
                }
                if ok { ways[n][mask] += ways[n-1][prevmask] }
            }
                ways[n][mask] %= mod
            }
        }
    }
}

func colorTheGrid2(m int, n int) (r int) {
    for _, mask := range rowmasks[m] { r += ways[n][mask] }
    return r % mod
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/22/colorthegrid.png" />
    // Input: m = 1, n = 1
    // Output: 3
    // Explanation: The three possible colorings are shown in the image above.
    fmt.Println(colorTheGrid(1,1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/22/copy-of-colorthegrid.png" />
    // Input: m = 1, n = 2
    // Output: 6
    // Explanation: The six possible colorings are shown in the image above.
    fmt.Println(colorTheGrid(1,2)) // 6
    // Example 3:
    // Input: m = 5, n = 5
    // Output: 580986
    fmt.Println(colorTheGrid(5,5)) // 580986

    fmt.Println(colorTheGrid(5,1000)) // 408208448
    fmt.Println(colorTheGrid(1,1000)) // 32634808
    fmt.Println(colorTheGrid(5,1)) // 48
    fmt.Println(colorTheGrid(1,1)) // 3

    fmt.Println(colorTheGrid1(1,1)) // 3
    fmt.Println(colorTheGrid1(1,2)) // 6
    fmt.Println(colorTheGrid1(5,5)) // 580986
    fmt.Println(colorTheGrid1(5,1000)) // 408208448
    fmt.Println(colorTheGrid1(1,1000)) // 32634808
    fmt.Println(colorTheGrid1(5,1)) // 48
    fmt.Println(colorTheGrid1(1,1)) // 3

    fmt.Println(colorTheGrid2(1,1)) // 3
    fmt.Println(colorTheGrid2(1,2)) // 6
    fmt.Println(colorTheGrid2(5,5)) // 580986
    fmt.Println(colorTheGrid2(5,1000)) // 408208448
    fmt.Println(colorTheGrid2(1,1000)) // 32634808
    fmt.Println(colorTheGrid2(5,1)) // 48
    fmt.Println(colorTheGrid2(1,1)) // 3
}