package main

import "fmt"

/**
74. Search a 2D Matrix

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.

Example 1:

	Input:
	matrix = [
	  [1,  3,  5,  7],
	  [10, 11, 16, 20],
	  [23, 30, 34, 50]
	]
	target = 3
	Output: true

Example 2:

	Input:
	matrix = [
	  [1,  3,  5,  7],
	  [10, 11, 16, 20],
	  [23, 30, 34, 50]
	]
	target = 13
	Output: false

解题思路:
	有序的
	使用二分法 先按列，再找行

 */

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for low <= high {
		mid := low + (high-low)>>1
		fmt.Printf("mid = %v\n",mid)
		fmt.Printf("matrix[mid/m][midm] = matrix[%v][%v] \n",mid/m,mid%m)
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{[]int{1, 3, 5, 7},[]int{10, 11, 16, 20},[]int{23, 30, 34, 50}}
	fmt.Printf("matrix = %v\n",matrix)
	fmt.Printf("searchMatrix(matrix,3) = %v\n",searchMatrix(matrix,3))
	fmt.Printf("searchMatrix(matrix,2) = %v\n",searchMatrix(matrix,2))
}
