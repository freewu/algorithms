package main

// 302. Smallest Rectangle Enclosing Black Pixels
// You are given an m x n binary matrix image where 0 represents a white pixel and 1 represents a black pixel.

// The black pixels are connected (i.e., there is only one black region). 
// Pixels are connected horizontally and vertically.

// Given two integers x and y that represents the location of one of the black pixels,
// return the area of the smallest (axis-aligned) rectangle that encloses all black pixels.

// You must write an algorithm with less than O(mn) runtime complexity

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/pixel-grid.jpg" />
// Input: image = [["0","0","1","0"],["0","1","1","0"],["0","1","0","0"]], x = 0, y = 2
// Output: 6

// Example 2:
// Input: image = [["1"]], x = 0, y = 0
// Output: 1
 
// Constraints:
//     m == image.length
//     n == image[i].length
//     1 <= m, n <= 100
//     image[i][j] is either '0' or '1'.
//     0 <= x < m
//     0 <= y < n
//     image[x][y] == '1'.
//     The black pixels in the image only form one component.

import "fmt"

// bfs 
func minArea(image [][]byte, x int, y int) int {
    directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    m, n := len(image), len(image[0])
    xMin, xMax, yMin, yMax := m, 0, n, 0
    q := [][]int{{x, y}}
    for i := 0; i < len(q); i++ {
        node := q[i]
        x, y := node[0], node[1]
        if x < xMin { xMin = x; }
        if x > xMax { xMax = x; }
        if y < yMin { yMin = y; }
        if y > yMax { yMax = y; }
        image[x][y] = '3'
        for _, direct := range directions {
            x1, y1 := x+direct[0], y+direct[1]
            if x1 < 0 || y1 < 0 || x1 == m || y1 == n {
                continue
            }
            if image[x1][y1] == '1' {
                q = append(q, []int{x1, y1})
            }
        }
    }
    for i := xMin; i <= xMax; i++ {
        for j := yMin; j <= yMax; j++ {
            if image[i][j] == '3' {
                image[i][j] = '1'
            }
        }
    }
    return (xMax - xMin + 1) * (yMax - yMin + 1)
}

// dfs
func minArea1(image [][]byte, x int, y int) int {
    m, n := len(image), len(image[0])
    xMin, xMax, yMin, yMax := m, 0, n, 0
    var dfs func(i, j int)
    dfs = func(x, y int) {
        if x < 0 || y < 0 || x == m || y == n || image[x][y] != '1' { return }
        if x < xMin { xMin = x; }
        if x > xMax { xMax = x; }
        if y < yMin { yMin = y; }
        if y > yMax { yMax = y; }
        image[x][y] = '0'
        dfs(x+1, y)
        dfs(x-1, y)
        dfs(x, y+1)
        dfs(x, y-1)
    }
    dfs(x, y)
    return (xMax - xMin + 1) * (yMax - yMin + 1)
}

// 二分查找
func minArea2(image [][]byte, x int, y int) int {
    findXExtremum := func (image [][]byte, l, r int, min bool) int {
        for l <= r {
            mid, found := (r-l)/2 + l, false
            for _, b := range image[mid] {
                if b == '1' {
                    found = true
                    break
                }
            }
            if found == min {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        if min {
            return l
        }
        return r
    }
    findYExtremum := func (image [][]byte, l, r int, min bool) int {
        for l <= r {
            mid, found := (r-l)/2 + l, false
            for _, row := range image {
                if row[mid] == '1' {
                    found = true
                    break
                }
            }
            if found == min {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        if min {
            return l
        }
        return r
    }
    m, n := len(image), len(image[0])
    xMin := findXExtremum(image, 0, x, true)
    xMax := findXExtremum(image, x, m-1, false)
    yMin := findYExtremum(image, 0, y, true)
    yMax := findYExtremum(image, y, n-1, false)
    return (xMax - xMin + 1) * (yMax - yMin + 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/pixel-grid.jpg" />
    // Input: image = [["0","0","1","0"],["0","1","1","0"],["0","1","0","0"]], x = 0, y = 2
    // Output: 6
    fmt.Println(minArea([][]byte{{'0','0','1','0'},{'0','1','1','0'},{'0','1','0','0'}}, 0, 2))
    // Example 2:
    // Input: image = [["1"]], x = 0, y = 0
    // Output: 1
    fmt.Println(minArea([][]byte{{1}}, 0, 0)) // 1

    fmt.Println(minArea1([][]byte{{'0','0','1','0'},{'0','1','1','0'},{'0','1','0','0'}}, 0, 2))
    fmt.Println(minArea1([][]byte{{1}}, 0, 0)) // 1

    fmt.Println(minArea2([][]byte{{'0','0','1','0'},{'0','1','1','0'},{'0','1','0','0'}}, 0, 2))
    fmt.Println(minArea2([][]byte{{1}}, 0, 0)) // 1
}