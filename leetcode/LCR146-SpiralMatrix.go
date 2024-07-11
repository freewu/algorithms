package main

// LCR 146. 螺旋遍历二维数组
// 给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。
// 螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。

// 示例 1：
// 输入：array = [[1,2,3],[8,9,4],[7,6,5]]
// 输出：[1,2,3,4,5,6,7,8,9]

// 示例 2：
// 输入：array  = [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]
// 输出：[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]

// 限制：
//     0 <= array.length <= 100
//     0 <= array[i].length <= 100

// 解题思路:
// 	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

import "fmt"

func spiralArray(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { // 处理为空的情况
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

func spiralArray1(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { // 处理为空的情况
        return []int{}
    }
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

func spiralArray2(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { // 处理为空的情况
        return []int{}
    }
    res := []int{}
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
    fmt.Printf(" spiralArray(%v) = %v\n",arr1,spiralArray(arr1))
    fmt.Printf(" spiralArray(%v) = %v\n",arr2,spiralArray(arr2))
    fmt.Printf(" spiralArray1(%v) = %v\n",arr1,spiralArray1(arr1))
    fmt.Printf(" spiralArray1(%v) = %v\n",arr2,spiralArray1(arr2))
    fmt.Printf(" spiralArray2(%v) = %v\n",arr1,spiralArray2(arr1))
    fmt.Printf(" spiralArray2(%v) = %v\n",arr2,spiralArray2(arr2))
}