package main

import "fmt"

/*
498	Diagonal Traverse

Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

 1 2 3
 4 5 6
 7 8 9

对角线 1 / 2 4 / 7 5 3 / 6 8 / 9

Example 1:

	Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
	Output: [1,2,4,7,5,3,6,8,9]

    [[0,0] [0,1],[1,0],[2,0],[1,1],[0,2],[1,2],[2,1],[2,2]]

Example 2:

	Input: mat = [[1,2],[3,4]]
	Output: [1,2,3,4]

Note:
	The total number of elements of the given matrix will not exceed 10,000.

对角线的顺序
 */

// 解法一
func findDiagonalOrder1(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col, dir, i, x, y, d := len(matrix), len(matrix[0]), [2][2]int{
		{-1, 1},
		{1, -1},
	}, 0, 0, 0, 0
	total := row * col
	res := make([]int, total)
	for i < total {
		for x >= 0 && x < row && y >= 0 && y < col {
			res[i] = matrix[x][y]
			i++
			x += dir[d][0]
			y += dir[d][1]
		}
		d = (d + 1) % 2
		if x == row {
			x--
			y += 2
		}
		if y == col {
			y--
			x += 2
		}
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
	}
	return res
}

// 解法二
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	// dir = 0 代表从右上到左下的方向， dir = 1 代表从左下到右上的方向 dir = -1 代表上一次转变了方向
	m, n, i, j, dir, res := len(matrix), len(matrix[0]), 0, 0, 0, []int{}
	for index := 0; index < m*n; index++ {
		if dir == -1 {
			if (i == 0 && j < n-1) || (j == n-1) { // 上边界和右边界
				i++
				if j > 0 {
					j--
				}
				dir = 0
				addTraverse(matrix, i, j, &res)
				continue
			}
			if (j == 0 && i < m-1) || (i == m-1) { // 左边界和下边界
				if i > 0 {
					i--
				}
				j++
				dir = 1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if i == 0 && j == 0 {
			res = append(res, matrix[i][j])
			if j < n-1 {
				j++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			} else {
				if i < m-1 {
					i++
					dir = -1
					addTraverse(matrix, i, j, &res)
					continue
				}
			}
		}
		if i == 0 && j < n-1 { // 上边界
			if j < n-1 {
				j++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if j == 0 && i < m-1 { // 左边界
			if i < m-1 {
				i++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if j == n-1 { // 右边界
			if i < m-1 {
				i++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if i == m-1 { // 下边界
			j++
			dir = -1
			addTraverse(matrix, i, j, &res)
			continue
		}
		if dir == 1 {
			i--
			j++
			addTraverse(matrix, i, j, &res)
			continue
		}
		if dir == 0 {
			i++
			j--
			addTraverse(matrix, i, j, &res)
			continue
		}
	}
	return res
}

func addTraverse(matrix [][]int, i, j int, res *[]int) {
	if i >= 0 && i <= len(matrix)-1 && j >= 0 && j <= len(matrix[0])-1 {
		*res = append(*res, matrix[i][j])
	}
}

// best solution
func findDiagonalOrderBest(matrix [][]int) []int {
	i, j, s:= 0, 0, 1
	width := len(matrix[0])
	height := len(matrix)
	length := width*height
	o := make([]int, length)
	for k := 0; k < length; k++ {
		o[k] = matrix[j][i]
		i += s
		j -= s
		if s == 1 {
			if i == width {
				j += 2
				i--
				s = -1
			} else if j < 0 {
				j = 0
				s = -1
			}
		} else {
			if j == height {
				i += 2
				j--
				s = 1
			} else if i < 0 {
				i = 0
				s = 1
			}
		}
	}
	return o
}

func main() {
	matrix := [][]int{ []int{1,2,3},  []int{4,5,6}, []int{7,8,9}}
	fmt.Printf("matrix: %v \n", matrix)
	fmt.Printf("findDiagonalOrder(matrix) =  %v\n",findDiagonalOrder(matrix))
	fmt.Printf("findDiagonalOrder1(matrix) =  %v\n",findDiagonalOrder1(matrix))
	fmt.Printf("findDiagonalOrderBest(matrix) =  %v\n",findDiagonalOrderBest(matrix))
}