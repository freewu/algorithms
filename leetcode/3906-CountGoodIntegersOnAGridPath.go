package main 

// 3906. Count Good Integers on a Grid Path
// You are given two integers l and r, and a string directions consisting of exactly three 'D' characters and three 'R' characters.

// For each integer x in the range [l, r] (inclusive), perform the following steps:
//     1. If x has fewer than 16 digits, pad it on the left with leading zeros to obtain a 16-digit string.
//     2. Place the 16 digits into a 4 × 4 grid in row-major order (the first 4 digits form the first row from left to right, 
//        the next 4 digits form the second row, and so on).
//     3. Starting at the top-left cell (row = 0, column = 0), apply the 6 characters of directions in order:
//         3.1 'D' character increments the row by 1.
//         3.2 'R' character increments the column by 1.
//     4. Record the sequence of digits visited along the path (including the starting cell), producing a sequence of length 7.

// The integer x is considered good if the recorded sequence is non-decreasing.

// Return an integer representing the number of good integers in the range [l, r].

// Example 1:
// Input: l = 8, r = 10, directions = "DDDRRR"
// Output: 2
// Explanation:
// The grid for x = 8:
// 0	0	0	0
// 0	0	0	0
// 0	0	0	0
// 0	0	0	8
// Path: (0,0) → (1,0) → (2,0) → (3,0) → (3,1) → (3,2) → (3,3)
// The sequence of digits visited is [0, 0, 0, 0, 0, 0, 8].
// As the sequence of digits visited is non-decreasing, 8 is a good integer.
// The grid for x = 9:
// 0	0	0	0
// 0	0	0	0
// 0	0	0	0
// 0	0	0	9
// The sequence of digits visited is [0, 0, 0, 0, 0, 0, 9].
// As the sequence of digits visited is non-decreasing, 9 is a good integer.
// The grid for x = 10:
// 0	0	0	0
// 0	0	0	0
// 0	0	0	0
// 0	0	1	0
// The sequence of digits visited is [0, 0, 0, 0, 0, 1, 0].
// As the sequence of digits visited is not non-decreasing, 10 is not a good integer.
// Hence, only 8 and 9 are good, giving a total of 2 good integers in the range.

// Example 2:
// Input: l = 123456789, r = 123456790, directions = "DDRRDR"
// Output: 1
// Explanation:
// The grid for x = 123456789:
// 0	0	0	0
// 0	0	0	1
// 2	3	4	5
// 6	7	8	9
// Path: (0,0) → (1,0) → (2,0) → (2,1) → (2,2) → (3,2) → (3,3)
// The sequence of digits visited is [0, 0, 2, 3, 4, 8, 9].
// As the sequence of digits visited is non-decreasing, 123456789 is a good integer.
// The grid for x = 123456790:
// 0	0	0	0
// 0	0	0	1
// 2	3	4	5
// 6	7	9	0
// The sequence of digits visited is [0, 0, 2, 3, 4, 9, 0].
// As the sequence of digits visited is not non-decreasing, 123456790 is not a good integer.
// Hence, only 123456789 is good, giving a total of 1 good integer in the range.

// Example 3:
// Input: l = 1288561398769758, r = 1288561398769758, directions = "RRRDDD"
// Output: 0
// Explanation:
// The grid for x = 1288561398769758:
// 1	2	8	8
// 5	6	1	3
// 9	8	7	6
// 9	7	5	8
// Path: (0,0) → (0,1) → (0,2) → (0,3) → (1,3) → (2,3) → (3,3)
// The sequence of digits visited is [1, 2, 8, 8, 3, 6, 8].
// ​​​​​​​As the sequence of digits visited is not non-decreasing, 1288561398769758 is not a good integer.
// No numbers are good, giving a total of 0 good integers in the range.

// Constraints:
//     1 <= l <= r <= 9 × 10^15
//     directions.length == 6
//     directions consists of exactly three 'D' characters and three 'R' characters.

import "fmt"
import "strconv"

func countGoodIntegersOnPath(l, r int64, directions string) int64 {
    lowS := strconv.FormatInt(l, 10)
    highS := strconv.FormatInt(r, 10)
    n := len(highS)
    inPath := make([]bool, n)
    pos := n - 16 // 右下角是下标 n-1，那么左上角是下标 n-16
    for _, d := range directions {
        if pos >= 0 { // 只需要对网格图的后 n 个格子做标记
            inPath[pos] = true // 标记在路径中的格子
        }
        if d == 'R' { // 往右
            pos++
        } else { // 往下
            pos += 4 // 相当于往右数 4 个位置
        }
    }
    inPath[n-1] = true // 终点一定在路径中
    diffLH := n - len(lowS)
    memo := make([][10]int64, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int, bool, bool) int64
    dfs = func(i, pre int, limitLow, limitHigh bool) int64 {
        if i == n { // 成功到达终点
            return 1 // 找到了一个好整数
        }
        res, low, high := int64(0), 0, 9
        if !limitLow && !limitHigh {
            p := &memo[i][pre]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        if limitLow && i >= diffLH {
            low = int(lowS[i-diffLH] - '0')
        }
        if limitHigh {
            high = int(highS[i] - '0')
        }
        d := low
        if inPath[i] { // 当前位置在路径中
            d = max(d, pre) // 当前位置填的数必须 >= pre
        }
        for ; d <= high; d++ {
            p := pre
            if inPath[i] {
                p = d
            }
            res += dfs(i+1, p, limitLow && d == low, limitHigh && d == high)
        }
        return res
    }
    return dfs(0, 0, true, true)
}

func countGoodIntegersOnPath1(l int64, r int64, directions string) int64 {
    var isPath [16]bool
    row, col := 0, 0
    isPath[0] = true
    for _, d := range directions {
        if d == 'D' {
            row++
        } else if d == 'R' {
            col++
        }
        isPath[row*4+col] = true
    }
    calc := func(n int64) int64 {
        if n < 0 { 
            return 0 
        }
        tar := fmt.Sprintf("%016d", n)
        var mem[16][2][10]int64
        for i := 0; i < 16; i++ {
            for j := 0; j < 2; j++ {
                for k := 0; k < 10; k++ {
                    mem[i][j][k] = -1
                }
            }
        }
        var tab func(pos, isLess, lastPath int) int64
        tab = func(pos, isLess, lastPath int) int64 {
            if pos == 16 {
                return 1
            }
            if mem[pos][isLess][lastPath] != -1 {
                return mem[pos][isLess][lastPath]
            }
            lim := 9
            if isLess == 0 {
                lim = int(tar[pos] - '0')
            }
            res := int64(0)
            for d := 0; d <= lim; d++ {
                if isPath[pos] && d < lastPath {
                    continue
                }
                nextLess := isLess
                if d < lim {
                    nextLess = 1
                }
                nextPath := lastPath
                if isPath[pos] {
                    nextPath = d
                }
                res += tab(pos+1, nextLess, nextPath)
            }
            mem[pos][isLess][lastPath] = res
            return res
        }
        return tab(0, 0, 0)
    }
    return calc(r) - calc(l-1)
}

func main() {
    // Example 1:
    // Input: l = 8, r = 10, directions = "DDDRRR"
    // Output: 2
    // Explanation:
    // The grid for x = 8:
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	0	8
    // Path: (0,0) → (1,0) → (2,0) → (3,0) → (3,1) → (3,2) → (3,3)
    // The sequence of digits visited is [0, 0, 0, 0, 0, 0, 8].
    // As the sequence of digits visited is non-decreasing, 8 is a good integer.
    // The grid for x = 9:
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	0	9
    // The sequence of digits visited is [0, 0, 0, 0, 0, 0, 9].
    // As the sequence of digits visited is non-decreasing, 9 is a good integer.
    // The grid for x = 10:
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	0	0
    // 0	0	1	0
    // The sequence of digits visited is [0, 0, 0, 0, 0, 1, 0].
    // As the sequence of digits visited is not non-decreasing, 10 is not a good integer.
    // Hence, only 8 and 9 are good, giving a total of 2 good integers in the range.
    fmt.Println(countGoodIntegersOnPath(8, 10, "DDDRRR")) // 2
    // Example 2:
    // Input: l = 123456789, r = 123456790, directions = "DDRRDR"
    // Output: 1
    // Explanation:
    // The grid for x = 123456789:
    // 0	0	0	0
    // 0	0	0	1
    // 2	3	4	5
    // 6	7	8	9
    // Path: (0,0) → (1,0) → (2,0) → (2,1) → (2,2) → (3,2) → (3,3)
    // The sequence of digits visited is [0, 0, 2, 3, 4, 8, 9].
    // As the sequence of digits visited is non-decreasing, 123456789 is a good integer.
    // The grid for x = 123456790:
    // 0	0	0	0
    // 0	0	0	1
    // 2	3	4	5
    // 6	7	9	0
    // The sequence of digits visited is [0, 0, 2, 3, 4, 9, 0].
    // As the sequence of digits visited is not non-decreasing, 123456790 is not a good integer.
    // Hence, only 123456789 is good, giving a total of 1 good integer in the range.
    fmt.Println(countGoodIntegersOnPath(123456789, 123456790, "DDRRDR")) // 1
    // Example 3:
    // Input: l = 1288561398769758, r = 1288561398769758, directions = "RRRDDD"
    // Output: 0
    // Explanation:
    // The grid for x = 1288561398769758:
    // 1	2	8	8
    // 5	6	1	3
    // 9	8	7	6
    // 9	7	5	8
    // Path: (0,0) → (0,1) → (0,2) → (0,3) → (1,3) → (2,3) → (3,3)
    // The sequence of digits visited is [1, 2, 8, 8, 3, 6, 8].
    // ​​​​​​​As the sequence of digits visited is not non-decreasing, 1288561398769758 is not a good integer.
    // No numbers are good, giving a total of 0 good integers in the range.
    fmt.Println(countGoodIntegersOnPath(1288561398769758, 1288561398769758, "RRRDDD")) // 0

    fmt.Println(countGoodIntegersOnPath(8, 10, "DDDRRR")) // 2
    fmt.Println(countGoodIntegersOnPath(123456789, 123456790, "DDRRDR")) // 1
    fmt.Println(countGoodIntegersOnPath(1288561398769758, 1288561398769758, "RRRDDD")) // 0
}

