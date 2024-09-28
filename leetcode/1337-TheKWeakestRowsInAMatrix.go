package main

// 1337. The K Weakest Rows in a Matrix
// You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). 
// The soldiers are positioned in front of the civilians. 
// That is, all the 1's will appear to the left of all the 0's in each row.

// A row i is weaker than a row j if one of the following is true:
//     The number of soldiers in row i is less than the number of soldiers in row j.
//     Both rows have the same number of soldiers and i < j.

// Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.

// Example 1:
// Input: mat = 
// [[1,1,0,0,0],
//  [1,1,1,1,0],
//  [1,0,0,0,0],
//  [1,1,0,0,0],
//  [1,1,1,1,1]], 
// k = 3
// Output: [2,0,3]
// Explanation: 
// The number of soldiers in each row is: 
// - Row 0: 2 
// - Row 1: 4 
// - Row 2: 1 
// - Row 3: 2 
// - Row 4: 5 
// The rows ordered from weakest to strongest are [2,0,3,1,4].

// Example 2:
// Input: mat = 
// [[1,0,0,0],
//  [1,1,1,1],
//  [1,0,0,0],
//  [1,0,0,0]], 
// k = 2
// Output: [0,2]
// Explanation: 
// The number of soldiers in each row is: 
// - Row 0: 1 
// - Row 1: 4 
// - Row 2: 1 
// - Row 3: 1 
// The rows ordered from weakest to strongest are [0,2,3,1].

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     2 <= n, m <= 100
//     1 <= k <= m
//     matrix[i][j] is either 0 or 1.

import "fmt"
import "sort"

func kWeakestRows(mat [][]int, k int) []int {
    type entry struct{
        index int // 行数
        soldiers int // 战士数量
    }
    count := make([]entry, len(mat), len(mat))
    for r, row := range mat {
        soldiers := 0
        for _, v := range row {
            soldiers += v
        }
        count[r] = entry{ r, soldiers }
    }
    sort.Slice(count, func(i, j int) bool {
        if count[i].soldiers == count[j].soldiers {
            return count[i].index < count[j].index
        }
        return count[i].soldiers < count[j].soldiers
    })
    res := make([]int, 0, k)
    for i := 0; i < k; i++ { // 取 n 个
        res = append(res, count[i].index)
    }
    return res
}

func kWeakestRows1(mat [][]int, k int) []int {
    m, n := len(mat), len(mat[0])
    type Row struct {
        Count int // 士兵个数
        Index int // 行数
    }
    rows := make([]Row, m)
    for i := 0; i < m; i++ { // 计算每行的士兵数量
        count := 0
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                count++
            } else {
                break // 因为士兵总是在前面，遇到0可以直接停止计数
            }
        }
        rows[i] = Row{ Count: count, Index: i}
    }
    insertionSort := func(rows []Row) {
        n := len(rows)
        for i := 1; i < n; i++ {
            key := rows[i]
            j := i - 1
            for j >= 0 && (rows[j].Count > key.Count || (rows[j].Count == key.Count && rows[j].Index > key.Index)) {
                rows[j+1] = rows[j]
                j--
            }
            rows[j+1] = key
        }
    }
    insertionSort(rows) // 使用插入排序对行进行排序
    // 提取前K个最弱行的索引
    res := make([]int, k)
    for i := 0; i < k; i++ {
        res[i] = rows[i].Index
    }
    return res
}

func main() {
    // Example 1:
    // Input: mat = 
    // [[1,1,0,0,0],
    //  [1,1,1,1,0],
    //  [1,0,0,0,0],
    //  [1,1,0,0,0],
    //  [1,1,1,1,1]], 
    // k = 3
    // Output: [2,0,3]
    // Explanation: 
    // The number of soldiers in each row is: 
    // - Row 0: 2 
    // - Row 1: 4 
    // - Row 2: 1 
    // - Row 3: 2 
    // - Row 4: 5 
    // The rows ordered from weakest to strongest are [2,0,3,1,4].
    mat1 := [][]int{
        {1,1,0,0,0},
        {1,1,1,1,0},
        {1,0,0,0,0},
        {1,1,0,0,0},
        {1,1,1,1,1},
    }
    fmt.Println(kWeakestRows(mat1, 3)) // [2,0,3]
    // Example 2:
    // Input: mat = 
    // [[1,0,0,0],
    //  [1,1,1,1],
    //  [],
    //  [1,0,0,0]], 
    // k = 2
    // Output: [0,2]
    // Explanation: 
    // The number of soldiers in each row is: 
    // - Row 0: 1 
    // - Row 1: 4 
    // - Row 2: 1 
    // - Row 3: 1 
    // The rows ordered from weakest to strongest are [0,2,3,1].
    mat2 := [][]int{
        {1,0,0,0},
        {1,1,1,1},
        {1,0,0,0},
        {1,0,0,0},
    }
    fmt.Println(kWeakestRows(mat2, 2)) // [0,2]

    fmt.Println(kWeakestRows1(mat1, 3)) // [2,0,3]
    fmt.Println(kWeakestRows1(mat2, 2)) // [0,2]
}