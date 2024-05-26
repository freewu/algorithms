package main

// 531. Lonely Pixel I
// Given an m x n picture consisting of black 'B' and white 'W' pixels, return the number of black lonely pixels.
// A black lonely pixel is a character 'B' 
// that located at a specific position where the same row and same column don't have any other black pixels.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/pixel1.jpg" />
// Input: picture = [["W","W","B"],["W","B","W"],["B","W","W"]]
// Output: 3
// Explanation: All the three 'B's are black lonely pixels.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/pixel2.jpg" />
// Input: picture = [["B","B","B"],["B","B","W"],["B","B","B"]]
// Output: 0
 
// Constraints:
//     m == picture.length
//     n == picture[i].length
//     1 <= m, n <= 500
//     picture[i][j] is 'W' or 'B'.

import "fmt"

func findLonelyPixel(picture [][]byte) int {
    res, m, n := 0, len(picture), len(picture[0])
    byR, byC := make([][]int, m), make([]int, n)
    for i := 0; i < m; i++ { // 先扫描一遍，得到以行号和列号为索引的B位置信息
        for j := 0; j < n; j++ {
            if picture[i][j] == 'B' {
                // 这里的if判断是一个内存优化，因为如果某行/列超过2个黑块，
                // 后面的其实不用再存了，因为后面二次扫描的时候只关心长度是否为1
                if len(byR[i]) < 2 {
                    byR[i] = append(byR[i], j)
                }
                byC[j]++
            }
        }
    }
    // 黑色孤独像素 的定义为：如果黑色像素 'B' 所在的同一行和同一列不存在其他黑色像素，那么这个黑色像素就是黑色孤独像素
    for i := 0; i < m; i++ { // 扫描预处理的结果
        if len(byR[i]) != 1 { continue;  } // 行不为 1 直接继续下一个
        if byC[byR[i][0]] == 1 { res++ } // 如果某行只有一个黑块，则看这个黑块对应列是不是也只有一个黑块
    }
    return res
}

func findLonelyPixel1(picture [][]byte) int {
    if len(picture) == 0 { return 0 }
    res, m, n := 0, len(picture), len(picture[0])
    column, row := make(map[int][]int, m), make(map[int][]int, n)
    for i,arr := range picture {
        for j,v := range arr{
            if v == 'B' {
                column[i] = append(column[i], j)
                row[j] = append(row[j],i)
            }
        }
    }
    for _,v := range column {
        if len(v) != 1 { continue } // 本列已超过一个像素的 b 黑色
        if len(row[v[0]]) !=1 { continue } // 本行已超过一个像素的 b 黑色
        res ++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/pixel1.jpg" />
    // Input: picture = [["W","W","B"],["W","B","W"],["B","W","W"]]
    // Output: 3
    // Explanation: All the three 'B's are black lonely pixels.
    fmt.Println(findLonelyPixel([][]byte{{'W','W','B'},{'W','B','W'},{'B','W','W'}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/pixel2.jpg" />
    // Input: picture = [["B","B","B"],["B","B","W"],["B","B","B"]]
    // Output: 0
    fmt.Println(findLonelyPixel([][]byte{{'B','B','B'},{'B','B','W'},{'B','B','B'}})) // 0

    fmt.Println(findLonelyPixel1([][]byte{{'W','W','B'},{'W','B','W'},{'B','W','W'}})) // 3
    fmt.Println(findLonelyPixel1([][]byte{{'B','B','B'},{'B','B','W'},{'B','B','B'}})) // 0
}