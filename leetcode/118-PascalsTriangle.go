package main

// 118. Pascal’s Triangle
// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
// <img src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif" />

// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

// Example 2:
// Input: numRows = 1
// Output: [[1]]

// Constraints:
//     1 <= numRows <= 30

// # 解题思路
//     给定一个 n，要求打印杨辉三角的前 n 行

import "fmt"

func generate(numRows int) [][]int {
    res := [][]int{}
    for i := 0; i < numRows; i++ { // 循环 numRows
        row := []int{}
        for j := 0; j < i+1; j++ { // 每次 循环 i + 1
            if j == 0 || j == i { // 如果是开头或者最后 都是 1
                row = append(row, 1)
            } else if i > 1 { // 0,1 都是 1,2以后开始计算
                // result[i-1][j-1]  上前
                // result[i-1][j] 上
                row = append(row, res[i-1][j-1] + res[i-1][j])
            }
        }
        res = append(res, row)
    }
    return res
}

// best solution
func generate1(numRows int) [][]int {
    res := make([][]int,numRows)
    res[0] = make([]int, 1)
    res[0] = append([]int{1}) // 直接加上0的 少一次循环

    for i := 1; i < numRows; i++ {
        res[i] = append(res[i], 1) // 开头是 1
        for j := 1; j < len(res[i-1]); j++ { // 只循环 1 - 上一层长度
            res[i] = append(res[i], res[i-1][j-1] + res[i-1][j])
        }
        res[i] = append(res[i], 1) // 结尾也 1
    }
    return res
}

func main() {
	fmt.Printf("generate(1) = %v\n",generate(1)) // [[1]]
	fmt.Printf("generate(5) = %v\n",generate(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Printf("generate(8) = %v\n",generate(8)) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1] [1 5 10 10 5 1] [1 6 15 20 15 6 1] [1 7 21 35 35 21 7 1]]

	fmt.Printf("generate1(1) = %v\n",generate1(1)) // [[1]]
	fmt.Printf("generate1(5) = %v\n",generate1(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Printf("generate1(8) = %v\n",generate1(8)) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1] [1 5 10 10 5 1] [1 6 15 20 15 6 1] [1 7 21 35 35 21 7 1]]
}
