package main

// 59. Spiral Matrix II
// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

// Example 1:
// Input: n = 3
// Output: [[1,2,3],[8,9,4],[7,6,5]]
// <img scr="https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg" />

// Example 2:
// Input: n = 1
// Output: [[1]]

// Constraints:
//     1 <= n <= 20

// 解题思路:
//     给出一个数 n，要求输出一个 n * n 的二维数组，里面元素是 1 - n*n，
//     数组排列顺序是螺旋排列的

import "fmt"

func generateMatrix(n int) [][]int {
    if n == 0 { return [][]int{} }
    if n == 1 { return [][]int{[]int{1}} }
    res, visited := make([][]int, n), make([][]bool, n)
    round, x, y := 0, 0, 0
    directions := [][]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }  // 朝右, 朝下, 朝左, 朝上
    for i := 0; i < n; i++ {
        visited[i], res[i] = make([]bool, n), make([]int, n)
    }
    res[x][y], visited[x][y] = 1, true
    for i := 0; i < n*n; i++ {
        x += directions[round % 4][0]
        y += directions[round % 4][1]
        if (x == 0 && y == n-1) || (x == n-1 && y == n-1) || (y == 0 && x == n-1) {
            round++
        }
        if x > n-1 || y > n-1 || x < 0 || y < 0 {
            return res
        }
        if !visited[x][y] {
            visited[x][y], res[x][y] = true , i + 2
        }
        switch round % 4 {
        case 0:
            if y + 1 <= n - 1 && visited[x][y + 1] {
                round++
                continue
            }
        case 1:
            if x + 1 <= n - 1 && visited[x + 1][y] {
                round++
                continue
            }
        case 2:
            if y - 1 >= 0 && visited[x][y - 1] {
                round++
                continue
            }
        case 3:
            if x - 1 >= 0 && visited[x - 1][y] {
                round++
                continue
            }
        }
    }
    return res
}

func generateMatrix1(n int) [][]int {
    directions := [4][2]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0},}
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    d, y, x := 0, 0, 0
    for i, e := 1, n * n; i <= e; i++ {
        res[y][x] = i
        y1, x1 := y + directions[d % 4][0], x + directions[d % 4][1]
        if y1 < 0 || y1 >= n || x1 < 0 || x1 >= n || res[y1][x1] > 0 {
            d++
            y1, x1 = y + directions[d % 4][0], x + directions[d % 4][1]
        }
        y, x = y1, x1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: [[1,2,3],[8,9,4],[7,6,5]]
    fmt.Printf("generateMatrix(3) = %v\n", generateMatrix(3)) // [[1,2,3],[8,9,4],[7,6,5]]
    // Example 2:
    // Input: n = 1
    // Output: [[1]]
    fmt.Printf("generateMatrix(1) = %v\n", generateMatrix(1)) // [[1]]

    fmt.Printf("generateMatrix(2) = %v\n", generateMatrix(2)) //  [[1 2] [4 3]]
    fmt.Printf("generateMatrix(4) = %v\n", generateMatrix(4)) // [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
    fmt.Printf("generateMatrix(5) = %v\n", generateMatrix(5)) // [[1 2 3 4 5] [16 17 18 19 6] [15 24 25 20 7] [14 23 22 21 8] [13 12 11 10 9]]

    fmt.Printf("generateMatrix1(3) = %v\n", generateMatrix1(3)) // [[1,2,3],[8,9,4],[7,6,5]]
    fmt.Printf("generateMatrix1(1) = %v\n", generateMatrix1(1)) // [[1]]
    fmt.Printf("generateMatrix1(2) = %v\n", generateMatrix1(2)) //  [[1 2] [4 3]]
    fmt.Printf("generateMatrix1(4) = %v\n", generateMatrix1(4)) // [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
    fmt.Printf("generateMatrix1(5) = %v\n", generateMatrix1(5)) // [[1 2 3 4 5] [16 17 18 19 6] [15 24 25 20 7] [14 23 22 21 8] [13 12 11 10 9]]
}
