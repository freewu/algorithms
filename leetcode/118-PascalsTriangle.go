package main

import "fmt"

/**
118. Pascal’s Triangle
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Constraints:

	1 <= numRows <= 30

Example 1:

	Input: 5
	Output:
	[
		 [1],
		[1,1],
	   [1,2,1],
	  [1,3,3,1],
	 [1,4,6,4,1]
	]

Example 2:

	Input: numRows = 1
	Output: [[1]]

# 解题思路
	给定一个 n，要求打印杨辉三角的前 n 行

*/

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ { // 循环 numRows
		var row []int
		for j := 0; j < i+1; j++ { // 每次 循环 i + 1
			if j == 0 || j == i { // 如果是开头或者最后 都是 1
				row = append(row, 1)
			} else if i > 1 { // 0,1 都是 1,2以后开始计算
				// result[i-1][j-1]  上前
				// result[i-1][j] 上
				row = append(row, result[i-1][j-1] + result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}

// best solution
func generateBest(numRows int) [][]int {
	array := make([][]int,numRows)
	array[0] = make([]int, 1)
	array[0] = append([]int{1}) // 直接加上0的 少一次循环

	for k := 1; k < numRows; k++ {
		array[k] = append(array[k], 1) // 开头是 1
		for i := 1; i < len(array[k-1]); i++ { // 只循环 1 - 上一层长度
			array[k] = append(array[k], array[k-1][i-1] + array[k-1][i])
		}
		array[k] = append(array[k], 1) // 结尾也 1
	}
	return array
}

func main() {
	fmt.Printf("generate(1) = %v\n",generate(1)) // [[1]]
	fmt.Printf("generate(5) = %v\n",generate(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Printf("generate(8) = %v\n",generate(8)) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1] [1 5 10 10 5 1] [1 6 15 20 15 6 1] [1 7 21 35 35 21 7 1]]

	fmt.Printf("generateBest(1) = %v\n",generateBest(1)) // [[1]]
	fmt.Printf("generateBest(5) = %v\n",generateBest(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Printf("generateBest(8) = %v\n",generateBest(8)) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1] [1 5 10 10 5 1] [1 6 15 20 15 6 1] [1 7 21 35 35 21 7 1]]
}
