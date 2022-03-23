package main

import "fmt"

/**
59. Spiral Matrix II
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

Constraints:
	1 <= n <= 20

Example 1:

	Input: n = 3
	Output: [[1,2,3],[8,9,4],[7,6,5]]

Example 2:

	Input: n = 1
	Output: [[1]]

解题思路:
	给出一个数 n，要求输出一个 n * n 的二维数组，里面元素是 1 - n*n，
	数组排列顺序是螺旋排列的

 */

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{[]int{1}}
	}
	res, visit, round, x, y, spDir := make([][]int, n), make([][]int, n), 0, 0, 0, [][]int{
		[]int{0, 1},  // 朝右
		[]int{1, 0},  // 朝下
		[]int{0, -1}, // 朝左
		[]int{-1, 0}, // 朝上
	}
	for i := 0; i < n; i++ {
		visit[i] = make([]int, n)
		res[i] = make([]int, n)
	}
	visit[x][y] = 1
	res[x][y] = 1
	for i := 0; i < n*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
		if (x == 0 && y == n-1) || (x == n-1 && y == n-1) || (y == 0 && x == n-1) {
			round++
		}
		if x > n-1 || y > n-1 || x < 0 || y < 0 {
			return res
		}
		if visit[x][y] == 0 {
			visit[x][y] = 1
			res[x][y] = i + 2
		}
		switch round % 4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1] == 1 {
				round++
				continue
			}
		case 1:
			if x+1 <= n-1 && visit[x+1][y] == 1 {
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

func main() {
	fmt.Printf("generateMatrix(1) = %v\n",generateMatrix(1)) // [[1]]
	fmt.Printf("generateMatrix(2) = %v\n",generateMatrix(2)) //  [[1 2] [4 3]]
	fmt.Printf("generateMatrix(3) = %v\n",generateMatrix(3)) // [[1,2,3],[8,9,4],[7,6,5]]
	fmt.Printf("generateMatrix(4) = %v\n",generateMatrix(4)) // [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
}
