package main

import "fmt"

/**
54. Spiral Matrix
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

	Input:
	[
	 [ 1, 2, 3 ],
	 [ 4, 5, 6 ],
	 [ 7, 8, 9 ]
	]
	Output: [1,2,3,6,9,8,7,4,5]

Example 2:

	Input:
	[
	  [1, 2, 3, 4],
	  [5, 6, 7, 8],
	  [9,10,11,12]
	]
	Output: [1,2,3,4,8,12,11,10,9,5,6,7]

解题思路:
	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	if len(matrix) == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			res = append(res, matrix[0][i])
		}
		return res
	}
	if len(matrix[0]) == 1 {
		for i := 0; i < len(matrix); i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}
	visit, m, n, round, x, y, spDir := make([][]int, len(matrix)), len(matrix), len(matrix[0]), 0, 0, 0, [][]int{
		[]int{0, 1},  // 朝右
		[]int{1, 0},  // 朝下
		[]int{0, -1}, // 朝左
		[]int{-1, 0}, // 朝上
	}
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}
	visit[x][y] = 1
	res = append(res, matrix[x][y])
	for i := 0; i < m*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
		if (x == 0 && y == n-1) || (x == m-1 && y == n-1) || (y == 0 && x == m-1) {
			round++
		}
		if x > m-1 || y > n-1 || x < 0 || y < 0 {
			return res
		}
		if visit[x][y] == 0 {
			visit[x][y] = 1
			res = append(res, matrix[x][y])
		}
		switch round % 4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1] == 1 {
				round++
				continue
			}
		case 1:
			if x+1 <= m-1 && visit[x+1][y] == 1 {
				round++
				continue
			}
		case 2:
			if y-1 >= 0 && visit[x][y-1] == 1 {
				round++
				continue
			}
		case 3:
			if x-1 >= 0 && visit[x-1][y] == 1 {
				round++
				continue
			}
		}
	}
	return res
}

func spiralOrder1(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	// top、left、right、bottom 分别是剩余区域的上、左、右、下的下标
	top, left, bottom, right := 0, 0, m-1, n-1
	count, sum := 0, m*n
	res := []int{}

	// 外层循环每次遍历一圈
	for count < sum {
		i, j := top, left
		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top + 1, right
		for i <= bottom && count < sum {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i, j = bottom, right - 1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i, j = bottom - 1, left
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		// 进入到下一层
		top, left, bottom, right = top+1, left+1, bottom-1, right-1
	}
	return res
}

// best solution
func spiralOrderBest(matrix [][]int) []int {
	var res []int
	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)

	for left < right && top < bottom {
		//->
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top += 1
		//↓
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right -= 1
		if !(left < right && top < bottom) {
			break
		}
		//←
		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom -= 1
		//↑
		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left += 1
	}
	return res
}

func main() {
	arr1 := [][]int{[]int{ 1, 2, 3},[]int{ 4, 5, 6  },[]int{ 7, 8, 9  } }
	arr2 := [][]int{[]int{ 1, 2, 3, 4},[]int{ 5, 6, 7, 8  },[]int{ 9,10,11,12 } }
	fmt.Printf(" spiralOrder(%v) = %v\n",arr1,spiralOrder(arr1))
	fmt.Printf(" spiralOrder(%v) = %v\n",arr2,spiralOrder(arr2))
	fmt.Printf(" spiralOrder1(%v) = %v\n",arr1,spiralOrder1(arr1))
	fmt.Printf(" spiralOrder1(%v) = %v\n",arr2,spiralOrder1(arr2))
	fmt.Printf(" spiralOrderBest(%v) = %v\n",arr1,spiralOrderBest(arr1))
	fmt.Printf(" spiralOrderBest(%v) = %v\n",arr2,spiralOrderBest(arr2))
}