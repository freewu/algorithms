package main

// 835. Image Overlap
// You are given two images, img1 and img2, represented as binary, square matrices of size n x n. 
// A binary matrix has only 0s and 1s as values.

// We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. 
// We then place it on top of the other image.
// We can then calculate the overlap by counting the number of positions that have a 1 in both images.

// Note also that a translation does not include any kind of rotation. 
// Any 1 bits that are translated outside of the matrix borders are erased.

// Return the largest possible overlap.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap1.jpg" />
// Input: img1 = [[1,1,0],[0,1,0],[0,1,0]], img2 = [[0,0,0],[0,1,1],[0,0,1]]
// Output: 3
// Explanation: We translate img1 to right by 1 unit and down by 1 unit.
//     <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap_step1.jpg" />
// The number of positions that have a 1 in both images is 3 (shown in red).
//     <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap_step2.jpg" />

// Example 2:
// Input: img1 = [[1]], img2 = [[1]]
// Output: 1

// Example 3:
// Input: img1 = [[0]], img2 = [[0]]
// Output: 0

// Constraints:
//     n == img1.length == img1[i].length
//     n == img2.length == img2[i].length
//     1 <= n <= 30
//     img1[i][j] is either 0 or 1.
//     img2[i][j] is either 0 or 1.

import "fmt"

// 解答错误 49 / 59 个通过的测试用例
// func largestOverlap(img1 [][]int, img2 [][]int) int {
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     helper := func (img1, img2 [][]int) int {
//         n, c := len(img1), 0
//         for x := 0; x < n; x++ {
//             for y := 0; y < n; y++ {
//                 t := 0
//                 for i := x; i < n; i++ {
//                     for j := y; j < n; j++ {
//                         if img1[i][j] == 1 && img2[i - x][j - y] == 1 { t++ } // 两个图像 都 具有 1 
//                     }
//                 }
//                 c = max(c, t)
//             }
//         }
//         return c
//     }
//     return max(helper(img1, img2), helper(img2, img1))
// }

// func largestOverlap(img1, img2 [][]int) int {
//     n, largest := len(img1), 0
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     overlap := func(img1, img2  [][]int, h, w int) int {
//         n, count := len(img1), 0
//         for i := 0; i < h; i++ {
//             for j := 0; j < w; j++ {
//                 if img2[i][j] == 1 && img2[i][j] == img1[n - h + i][n - w + j] {
//                     count++
//                 }
//             }
//         }
//         return count
//     }
//     for h := 1; h <= n; h++ {
//         for w := 1; w <= n; w++ {
//             largest = max(largest, overlap(img1, img2, h, w))
//             largest = max(largest, overlap(img2, img1, h, w))
//         }
//     }
//     return largest
// }

func largestOverlap(img1 [][]int, img2 [][]int) int {
    maxOverlap, n := 0, len(img1)
    maxMove := n - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := -maxMove; i <= maxMove; i++ {
        for j := -maxMove; j <= maxMove; j++ {
            overlap := 0
            rowStart, rowEnd := max(0, i), min(n, n + i)
            columnStart, columnEnd := max(0, j), min(n, n + j)
            for row := rowStart; row < rowEnd; row++ {
                for column := columnStart; column < columnEnd; column++ {
                    if img1[row - i][column - j] == 1 && img2[row][column] == 1 {
                        overlap++
                    }
                }
            }
            maxOverlap = max(maxOverlap, overlap)
        }
    }
    return maxOverlap
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap1.jpg" />
    // Input: img1 = [[1,1,0],[0,1,0],[0,1,0]], img2 = [[0,0,0],[0,1,1],[0,0,1]]
    // Output: 3
    // Explanation: We translate img1 to right by 1 unit and down by 1 unit.
    //     <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap_step1.jpg" />
    // The number of positions that have a 1 in both images is 3 (shown in red).
    //     <img src="https://assets.leetcode.com/uploads/2020/09/09/overlap_step2.jpg" />
    fmt.Println(largestOverlap([][]int{{1,1,0},{0,1,0},{0,1,0}}, [][]int{{0,0,0},{0,1,1},{0,0,1}})) // 3
    // Example 2:
    // Input: img1 = [[1]], img2 = [[1]]
    // Output: 1
    fmt.Println(largestOverlap([][]int{{1}},[][]int{{1}})) // 1
    // Example 3:
    // Input: img1 = [[0]], img2 = [[0]]
    // Output: 0
    fmt.Println(largestOverlap([][]int{{0}},[][]int{{0}})) // 0

    img11 := [][]int{{0,0,0,0,1},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},}
    img12 := [][]int{{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{1,0,0,0,0},}
    fmt.Println(largestOverlap(img11, img12)) // 1
}