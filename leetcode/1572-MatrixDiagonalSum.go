package main

// 1572. Matrix Diagonal Sum
// Given a square matrix mat, return the sum of the matrix diagonals.
// Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/14/sample_1911.png" />
// Input: mat = [[1,2,3],
//               [4,5,6],
//               [7,8,9]]
// Output: 25
// Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
// Notice that element mat[1][1] = 5 is counted only once.

// Example 2:
// Input: mat = [[1,1,1,1],
//               [1,1,1,1],
//               [1,1,1,1],
//               [1,1,1,1]]
// Output: 8

// Example 3:
// Input: mat = [[5]]
// Output: 5
 
// Constraints:
//     n == mat.length == mat[i].length
//     1 <= n <= 100
//     1 <= mat[i][j] <= 100

import "fmt"

// ([0,0] + [1,1] + [2,2] + ... [n,n]) + ([0, n] + [1, n - 1] + [2, n - 2] + [3, n - 3] + .. + [n, 0])
func diagonalSum(mat [][]int) int {
    n, res := len(mat), 0
    for i := 0; i < n; i++ {
        res += mat[i][i]
        // 排除交叉项
        if i != n - i - 1 {
            res += mat[i][n - i - 1]
        }
    }
    return res
}

func diagonalSum1(mat [][]int) int {
    l, res := len(mat) - 1, 0
    for i := 0; i <= l; i++ {
        res += mat[i][i]
        // 排除交叉项
        if i != l - i {
            res += mat[i][l - i]
        }
    }
    return res
}

func diagonalSum2(mat [][]int) int {
    l := len(mat)
    res := 0
    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            if i == j || i + j == l-1 {
                res += mat[i][j]
            }
        }
    }
    return res
}

func main() {
    // Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
    // Notice that element mat[1][1] = 5 is counted only once.
    fmt.Println(
        diagonalSum(
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
        ),
    ) // 25

    fmt.Println(
        diagonalSum(
            [][]int{
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
            },
        ),
    ) // 8

    fmt.Println(
        diagonalSum(
            [][]int{
                []int{5},
            },
        ),
    ) // 5


    fmt.Println(
        diagonalSum1(
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
        ),
    ) // 25

    fmt.Println(
        diagonalSum1(
            [][]int{
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
            },
        ),
    ) // 8

    fmt.Println(
        diagonalSum1(
            [][]int{
                []int{5},
            },
        ),
    ) // 5

    fmt.Println(
        diagonalSum2(
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
                []int{7,8,9},
            },
        ),
    ) // 25

    fmt.Println(
        diagonalSum2(
            [][]int{
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
                []int{1,1,1,1},
            },
        ),
    ) // 8

    fmt.Println(
        diagonalSum2(
            [][]int{
                []int{5},
            },
        ),
    ) // 5
}