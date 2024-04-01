package main

// 240. Search a 2D Matrix II
// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. 
// This matrix has the following properties:
//     Integers in each row are sorted in ascending from left to right.
//     Integers in each column are sorted in ascending from top to bottom.
	
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/24/searchgrid2.jpg" />
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/24/searchgrid.jpg" />
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// Output: false

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= n, m <= 300
//     -10^9 <= matrix[i][j] <= 10^9
//     All the integers in each row are sorted in ascending order.
//     All the integers in each column are sorted in ascending order.
//     -10^9 <= target <= 10^9

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//     每行的元素从左到右升序排列。
//     每列的元素从上到下升序排列。


// 模拟，时间复杂度 O(m+n)
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    row, col := 0, len(matrix[0])-1
    for col >= 0 && row <= len(matrix)-1 {
        if target == matrix[row][col] {
            return true
        } else if target > matrix[row][col] { // 如果当前值小于目标值向下一层走，否则向里走一步
            row++
        } else {
            col--
        }
    }
    return false
}

// 二分搜索，时间复杂度 O(n log n)
func searchMatrix1(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    for _, row := range matrix { // 循环每行
        // 每行做二分法查找
        low, high := 0, len(matrix[0])-1
        for low <= high {
            mid := low + (high-low) >> 1
            if row[mid] > target {
                high = mid - 1
            } else if row[mid] < target {
                low = mid + 1
            } else { // row[mid] == target 找到了
                return true
            }
        }
    }
    return false
}

func main() {
	matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	fmt.Printf("searchMatrix1(matrix,5) = %v\n",searchMatrix1(matrix,5)) // true
	fmt.Printf("searchMatrix1(matrix,20) = %v\n",searchMatrix1(matrix,20)) // false

	fmt.Printf("searchMatrix(matrix,5) = %v\n",searchMatrix(matrix,5)) // true
	fmt.Printf("searchMatrix(matrix,20) = %v\n",searchMatrix(matrix,20)) // false
}