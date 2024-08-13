package main

// 3030. Find the Grid of Region Average
// You are given m x n grid image which represents a grayscale image, 
// where image[i][j] represents a pixel with intensity in the range [0..255]. 
// You are also given a non-negative integer threshold.

// Two pixels are adjacent if they share an edge.

// A region is a 3 x 3 subgrid where the absolute difference in intensity between any two adjacent pixels is less than or equal to threshold.

// All pixels in a region belong to that region, note that a pixel can belong to multiple regions.

// You need to calculate a m x n grid result, where result[i][j] is the average intensity of the regions to which image[i][j] belongs, rounded down to the nearest integer. 
// If image[i][j] belongs to multiple regions, result[i][j] is the average of the rounded-down average intensities of these regions, rounded down to the nearest integer. 
// If image[i][j] does not belong to any region, result[i][j] is equal to image[i][j].

// Return the grid result.

// Example 1:
// Input: image = [[5,6,7,10],[8,9,10,10],[11,12,13,10]], threshold = 3
// Output: [[9,9,9,9],[9,9,9,9],[9,9,9,9]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2023/12/21/example0corrected.png" />
// There are two regions as illustrated above. 
// The average intensity of the first region is 9, while the average intensity of the second region is 9.67 which is rounded down to 9. The average intensity of both of the regions is (9 + 9) / 2 = 9. 
// As all the pixels belong to either region 1, region 2, or both of them, the intensity of every pixel in the result is 9.
// Please note that the rounded-down values are used when calculating the average of multiple regions, hence the calculation is done using 9 as the average intensity of region 2, not 9.67.

// Example 2:
// Input: image = [[10,20,30],[15,25,35],[20,30,40],[25,35,45]], threshold = 12
// Output: [[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2023/12/21/example1corrected.png" />
// There are two regions as illustrated above. 
// The average intensity of the first region is 25, while the average intensity of the second region is 30. 
// The average intensity of both of the regions is (25 + 30) / 2 = 27.5 which is rounded down to 27.
// All the pixels in row 0 of the image belong to region 1, hence all the pixels in row 0 in the result are 25. 
// Similarly, all the pixels in row 3 in the result are 30. 
// The pixels in rows 1 and 2 of the image belong to region 1 and region 2, hence their assigned value is 27 in the result.

// Example 3:
// Input: image = [[5,6,7],[8,9,10],[11,12,13]], threshold = 1
// Output: [[5,6,7],[8,9,10],[11,12,13]]
// Explanation:
// There is only one 3 x 3 subgrid, while it does not have the condition on difference of adjacent pixels, 
// for example, the difference between image[0][0] and image[1][0] is |5 - 8| = 3 > threshold = 1. 
// None of them belong to any valid regions, so the result should be the same as image.

// Constraints:
//     3 <= n, m <= 500
//     0 <= image[i][j] <= 255
//     0 <= threshold <= 255

import "fmt"

func resultGrid(image [][]int, threshold int) [][]int {
    n, m := len(image), len(image[0])
    sum := make([][]int, n)
    num := make([][]int, n)
    for i := range sum {
        sum[i] = make([]int, m)
        num[i] = make([]int, m)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    traverseRegion := func (row, col, t int) {
        xDir, yDir := []int{0, 1}, []int{-1, 0}
        avg := 0
        for i := row; row+3 <= n && i < row+3; i++ {
            for j := col; col+3 <= m && j < col+3; j++ {
                for d := 0; d < 2; d++ {
                    adji, adjj := i+xDir[d], j+yDir[d]                
                    if adji >= row && adji < row + 3 && adjj >= col && adjj < col+3 && abs(image[i][j] - image[adji][adjj]) > t {
                        return
                    }
                }
                avg += image[i][j]
            }
        }
        for i := row; row+3 <= n && i < row+3; i++ {
            for j := col; col+3 <= m && j < col+3; j++ {            
                sum[i][j] = sum[i][j] + (avg / 9)
                num[i][j]++
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            traverseRegion(i, j, threshold)
        }
    }
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, m)
        for j := 0; j < m; j++ {
            v := image[i][j]
            if num[i][j] > 0 {
                v = sum[i][j] / num[i][j]
            }
            res[i][j] = v
        }
    }
    return res
}

func resultGrid1(image [][]int, threshold int) [][]int {
    m, n := len(image), len(image[0])
    res, cnt := make([][]int, m), make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
        cnt[i] = make([]int, n)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    check := func(i, j int) bool {
        // 检查左右相邻格子
        for _, row := range image[i-2 : i+1] {
            if abs(row[j-2]-row[j-1]) > threshold || abs(row[j-1]-row[j]) > threshold {
                return false // 不合法，下一个
            }
        }
        // 检查上下相邻格子
        for y := j - 2; y <= j; y++ {
            if abs(image[i-2][y] - image[i-1][y]) > threshold || abs(image[i-1][y]-image[i][y]) > threshold {
                return false // 不合法，下一个
            }
        }
        // 合法，计算 3x3 子网格的平均值
        avg := 0
        for x := i - 2; x <= i; x++ {
            for y := j - 2; y <= j; y++ {
                avg += image[x][y]
            }
        }
        avg /= 9
        // 更新 3x3 子网格内的 result
        for x := i - 2; x <= i; x++ {
            for y := j - 2; y <= j; y++ {
                res[x][y] += avg // 先累加，最后再求平均值
                cnt[x][y]++
            }
        }
        return true
    }
    for i := 2; i < m; i++ {
        for j := 2; j < n; j++ {
            if !check(i,j) {
                break
            }
        }
    }
    for i, row := range cnt {
        for j, c := range row {
            if c == 0 { // (i,j) 不属于任何子网格
                res[i][j] = image[i][j]
            } else {
                res[i][j] /= c // 求平均值
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: image = [[5,6,7,10],[8,9,10,10],[11,12,13,10]], threshold = 3
    // Output: [[9,9,9,9],[9,9,9,9],[9,9,9,9]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2023/12/21/example0corrected.png" />
    // There are two regions as illustrated above. 
    // The average intensity of the first region is 9, while the average intensity of the second region is 9.67 which is rounded down to 9. The average intensity of both of the regions is (9 + 9) / 2 = 9. 
    // As all the pixels belong to either region 1, region 2, or both of them, the intensity of every pixel in the result is 9.
    // Please note that the rounded-down values are used when calculating the average of multiple regions, hence the calculation is done using 9 as the average intensity of region 2, not 9.67.
    image1 := [][]int{
        {5,6,7,10},
        {8,9,10,10},
        {11,12,13,10},
    }
    fmt.Println(resultGrid(image1, 3)) // [[9,9,9,9],[9,9,9,9],[9,9,9,9]]
    // Example 2:
    // Input: image = [[10,20,30],[15,25,35],[20,30,40],[25,35,45]], threshold = 12
    // Output: [[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2023/12/21/example1corrected.png" />
    // There are two regions as illustrated above. 
    // The average intensity of the first region is 25, while the average intensity of the second region is 30. 
    // The average intensity of both of the regions is (25 + 30) / 2 = 27.5 which is rounded down to 27.
    // All the pixels in row 0 of the image belong to region 1, hence all the pixels in row 0 in the result are 25. 
    // Similarly, all the pixels in row 3 in the result are 30. 
    // The pixels in rows 1 and 2 of the image belong to region 1 and region 2, hence their assigned value is 27 in the result.
    image2 := [][]int{
        {10,20,30},
        {15,25,35},
        {20,30,40},
        {25,35,45},
    }
    fmt.Println(resultGrid(image2, 12)) // [[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
    // Example 3:
    // Input: image = [[5,6,7],[8,9,10],[11,12,13]], threshold = 1
    // Output: [[5,6,7],[8,9,10],[11,12,13]]
    // Explanation:
    // There is only one 3 x 3 subgrid, while it does not have the condition on difference of adjacent pixels, 
    // for example, the difference between image[0][0] and image[1][0] is |5 - 8| = 3 > threshold = 1. 
    // None of them belong to any valid regions, so the result should be the same as image.
    image3 := [][]int{
        {5,6,7},
        {8,9,10},
        {11,12,13},
    }
    fmt.Println(resultGrid(image3, 1)) // [[5,6,7],[8,9,10],[11,12,13]]

    fmt.Println(resultGrid1(image1, 3)) // [[9,9,9,9],[9,9,9,9],[9,9,9,9]]
    fmt.Println(resultGrid1(image2, 12)) // [[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
    fmt.Println(resultGrid1(image3, 1)) // [[5,6,7],[8,9,10],[11,12,13]]
}