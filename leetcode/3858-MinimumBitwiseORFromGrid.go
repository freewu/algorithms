package main

// 3858. Minimum Bitwise OR From Grid
// You are given a 2D integer array grid of size m x n.

// You must select exactly one integer from each row of the grid.

// Return an integer denoting the minimum possible bitwise OR of the selected integers from each row.

// Example 1:
// Input: grid = [[1,5],[2,4]]
// Output: 3
// Explanation:
// Choose 1 from the first row and 2 from the second row.
// The bitwise OR of 1 | 2 = 3​​​​​​​, which is the minimum possible.

// Example 2:
// Input: grid = [[3,5],[6,4]]
// Output: 5
// Explanation:
// Choose 5 from the first row and 4 from the second row.
// The bitwise OR of 5 | 4 = 5​​​​​​​, which is the minimum possible.

// Example 3:
// Input: grid = [[7,9,8]]
// Output: 7
// Explanation:
// Choosing 7 gives the minimum bitwise OR.
 
// Constraints:
//     1 <= m == grid.length <= 10^5
//     1 <= n == grid[i].length <= 10^5
//     m * n <= 10^5
//     1 <= grid[i][j] <= 10^5

import "fmt"
import "slices"
import "math/bits"

func minimumOR(grid [][]int) int {
    res, mx := 0, 0
    for _, row := range grid { // 找到最大值
        mx = max(mx, slices.Max(row))
    }
    // 试填法：res 的第 i 位能不能是 0？
    // 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 res 的第 i 位可以是 0，否则必须是 1
    for i := bits.Len(uint(mx)) - 1; i >= 0; i-- {
        mask := res | (1<<i - 1) // mask 低于 i 的比特位全是 1，表示 grid[i][j] 的低位是 0 还是 1 无所谓
    next:
        for _, row := range grid {
            for _, x := range row {
                // x 的高于 i 的比特位，如果 res 是 0，那么 x 的这一位必须也是 0    
                // x 的低于 i 的比特位，随意
                // x 的第 i 个比特位，我们期望它是 0
                if x|mask == mask { // x 可以选，且第 i 位是 0
                    continue next
                }
            }
            // 这一行的可选数字中，第 i 位全是 1
            res |= 1 << i // res 第 i 位必须是 1
            break // 填下一位
        }
    }
    return res
}

func minimumOR1(grid [][]int) int {
    res := 0
    for i := 16; i >= 0; i-- {
        mask, flag := res | ((1 << i) - 1), true
        for _, row := range grid {
            match := false
            for _, v := range row {
                if v & ^mask == 0 {
                    match = true
                    break
                }
            }
            if !match {
                flag = false
                break
            }
        }
        if !flag {
            res = res | (1 << i)
        }
    }
    return res  
}

func main() {
    // Example 1:
    // Input: grid = [[1,5],[2,4]]
    // Output: 3
    // Explanation:
    // Choose 1 from the first row and 2 from the second row.
    // The bitwise OR of 1 | 2 = 3​​​​​​​, which is the minimum possible.
    fmt.Println(minimumOR([][]int{{1,5},{2,4}})) // 3
    // Example 2:
    // Input: grid = [[3,5],[6,4]]
    // Output: 5
    // Explanation:
    // Choose 5 from the first row and 4 from the second row.
    // The bitwise OR of 5 | 4 = 5​​​​​​​, which is the minimum possible.
    fmt.Println(minimumOR([][]int{{3,5},{6,4}})) // 5
    // Example 3:
    // Input: grid = [[7,9,8]]
    // Output: 7
    // Explanation:
    // Choosing 7 gives the minimum bitwise OR.   
    fmt.Println(minimumOR([][]int{{7,9,8}})) // 7

    fmt.Println(minimumOR1([][]int{{1,5},{2,4}})) // 3
    fmt.Println(minimumOR1([][]int{{3,5},{6,4}})) // 5
    fmt.Println(minimumOR1([][]int{{7,9,8}})) // 7
}