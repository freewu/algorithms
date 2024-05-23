package main

// 1198. Find Smallest Common Element in All Rows
// Given an m x n matrix mat where every row is sorted in strictly increasing order, 
// return the smallest common element in all rows.

// If there is no common element, return -1.

// Example 1:
// Input: mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
// Output: 5

// Example 2:
// Input: mat = [[1,2,3],[2,3,4],[2,3,5]]
// Output: 2

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 500
//     1 <= mat[i][j] <= 10^4
//     mat[i] is sorted in strictly increasing order.

import "fmt"
import "sort"

func smallestCommonElement(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    arr := []int{} // 可提前分配内存, 也可考虑 双端队列
    for _, x := range mat[0] { // 将第一行的元素做为基准, 每个都有
        arr = append(arr, x)
    }
    for i := 1; i < m; i++ { // 逐行遍历, 基准数组中每个元素, 没找到就去掉元素
        for j := len(arr) - 1; j >= 0; j-- {  // 倒序遍历, 因为删除元素后不用更新下标
            idx := sort.SearchInts(mat[i], arr[j])
            if idx == n || mat[i][idx] != arr[j] {
                arr = append(arr[:j], arr[j+1:]...)
            }
        }
    }
    if len(arr) == 0 {
        return -1
    }
    return arr[0]
}

func smallestCommonElement1(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    arr := make([]int, m)
    for {
        flag := false
        for i := 1; i < m; i++ {
            j, k := arr[i-1], arr[i]
            for j < n && k < n && mat[i-1][j] != mat[i][k] {
                flag = true
                if mat[i-1][j] < mat[i][k] {
                    j++
                } else {
                    k++
                }
            }
            if j == n || k == n {
                return -1
            }
            if flag {
                arr[i-1] = j
                arr[i] = k
            }
        }
        if !flag {
            return mat[0][arr[0]]
        }
    }
}

func main() {
    // Example 1:
    // Input: mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
    // Output: 5
    fmt.Println(smallestCommonElement([][]int{{1,2,3,4,5},{2,4,5,8,10},{3,5,7,9,11},{1,3,5,7,9}})) // 5
    // Example 2:
    // Input: mat = [[1,2,3],[2,3,4],[2,3,5]]
    // Output: 2f
    fmt.Println(smallestCommonElement([][]int{{1,2,3},{2,3,4},{2,3,5}})) // 2

    fmt.Println(smallestCommonElement1([][]int{{1,2,3,4,5},{2,4,5,8,10},{3,5,7,9,11},{1,3,5,7,9}})) // 5
    fmt.Println(smallestCommonElement1([][]int{{1,2,3},{2,3,4},{2,3,5}})) // 2
}