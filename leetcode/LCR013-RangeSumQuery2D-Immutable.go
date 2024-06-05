package main

// LCR 013. 二维区域和检索 - 矩阵不可变
// 给定一个二维矩阵 matrix，以下类型的多个请求：
//     计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。

// 实现 NumMatrix 类：
//     NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//     int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。

// 示例 1：
// <img src="https://pic.leetcode-cn.com/1626332422-wUpUHT-image.png" />
// 输入: 
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
// [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
// 输出: 
// [null, 8, 11, 12]
// 解释:
// NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
 
// 提示：
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 200
//     -10^5 <= matrix[i][j] <= 10^5
//     0 <= row1 <= row2 < m
//     0 <= col1 <= col2 < n
//     最多调用 10^4 次 sumRegion 方法

import "fmt"

type NumMatrix struct {
    cumsum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 {
        return NumMatrix{nil}
    }
    // 生成一个二维数组保存结果
    cumsum := make([][]int, len(matrix)+1)
    cumsum[0] = make([]int, len(matrix[0])+1)
    for i := range matrix {
        cumsum[i+1] = make([]int, len(matrix[i])+1)
        for j := range matrix[i] {
            cumsum[i+1][j+1] = matrix[i][j] + cumsum[i][j+1] + cumsum[i+1][j] - cumsum[i][j]
        }
    }
    fmt.Println(cumsum)
    return NumMatrix{cumsum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    cumsum := this.cumsum
    return cumsum[row2+1][col2+1] - cumsum[row1][col2+1] - cumsum[row2+1][col1] + cumsum[row1][col1]
}

func main() {
    numMatrix := Constructor([][]int{
        {3, 0, 1, 4, 2}, 
        {5, 6, 3, 2, 1}, 
        {1, 2, 0, 1, 5}, 
        {4, 1, 0, 1, 7}, 
        {1, 0, 3, 0, 5},
    })
    fmt.Println(numMatrix.SumRegion(2, 1, 4, 3)); // return 8 (i.e sum of the red rectangle)
    fmt.Println(numMatrix.SumRegion(1, 1, 2, 2)); // return 11 (i.e sum of the green rectangle)
    fmt.Println(numMatrix.SumRegion(1, 2, 2, 4)); // return 12 (i.e sum of the blue rectangle)
}