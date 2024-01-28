package main

import "fmt"

// 1074. Number of Submatrices That Sum to Target
// Given a matrix and a target, return the number of non-empty submatrices that sum to target.
// A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
// Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.

// Example 1:
// Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
// Output: 4
// Explanation: The four 1x1 submatrices that only contain 0.

// Example 2:
// Input: matrix = [[1,-1],[-1,1]], target = 0
// Output: 5
// Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.

// Example 3:
// Input: matrix = [[904]], target = 0
// Output: 0
 
// Constraints:

// 		1 <= matrix.length <= 100
// 		1 <= matrix[0].length <= 100
// 		-1000 <= matrix[i] <= 1000
// 		-10^8 <= target <= 10^8

// 给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n, res := len(matrix), len(matrix[0]), 0
	for row := range matrix {
		for col := 1; col < len(matrix[row]); col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			counterMap, sum := make(map[int]int, m), 0
			counterMap[0] = 1 // 题目保证一定有解，所以这里初始化是 1
			for row := 0; row < m; row++ {
				if i > 0 {
					sum += matrix[row][j] - matrix[row][i-1]
				} else {
					sum += matrix[row][j]
				}
				res += counterMap[sum-target]
				counterMap[sum]++
			}
		}
	}
	return res
}

// 暴力解法 O(n^4)
func numSubmatrixSumTarget1(matrix [][]int, target int) int {
	m, n, res, sum := len(matrix), len(matrix[0]), 0, 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			counterMap := map[int]int{}
			counterMap[0] = 1 // 题目保证一定有解，所以这里初始化是 1
			sum = 0
			for row := 0; row < m; row++ {
				for k := i; k <= j; k++ {
					sum += matrix[row][k]
				}
				res += counterMap[sum-target]
				counterMap[sum]++
			}
		}
	}
	return res
}

func main() {
	matrix1 := [][]int{[]int{0,1,0},[]int{1,1,1},[]int{0,1,0}}
	matrix2 := [][]int{[]int{1,-1},[]int{-1,1}}
	fmt.Println(numSubmatrixSumTarget(matrix1,0)) // 4
	fmt.Println(numSubmatrixSumTarget(matrix2,0)) // 5

	fmt.Println(numSubmatrixSumTarget1(matrix1,0)) // 4
	fmt.Println(numSubmatrixSumTarget1(matrix2,0)) // 5
}