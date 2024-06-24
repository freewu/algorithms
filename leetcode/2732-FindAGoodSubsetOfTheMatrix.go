package main

// 2732. Find a Good Subset of the Matrix
// You are given a 0-indexed m x n binary matrix grid.
// Let us call a non-empty subset of rows good if the sum of each column of the subset is at most half of the length of the subset.
// More formally, if the length of the chosen subset of rows is k, then the sum of each column should be at most floor(k / 2).
// Return an integer array that contains row indices of a good subset sorted in ascending order.
// If there are multiple good subsets, you can return any of them. If there are no good subsets, return an empty array.
// A subset of rows of the matrix grid is any matrix that can be obtained by deleting some (possibly none or all) rows from grid.

// Example 1:
// Input: grid = [[0,1,1,0],[0,0,0,1],[1,1,1,1]]
// Output: [0,1]
// Explanation: We can choose the 0th and 1st rows to create a good subset of rows.
// The length of the chosen subset is 2.
// - The sum of the 0th column is 0 + 0 = 0, which is at most half of the length of the subset.
// - The sum of the 1st column is 1 + 0 = 1, which is at most half of the length of the subset.
// - The sum of the 2nd column is 1 + 0 = 1, which is at most half of the length of the subset.
// - The sum of the 3rd column is 0 + 1 = 1, which is at most half of the length of the subset.

// Example 2:
// Input: grid = [[0]]
// Output: [0]
// Explanation: We can choose the 0th row to create a good subset of rows.
// The length of the chosen subset is 1.
// - The sum of the 0th column is 0, which is at most half of the length of the subset.

// Example 3:
// Input: grid = [[1,1,1],[1,1,1]]
// Output: []
// Explanation: It is impossible to choose any subset of rows to create a good subset.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m <= 10^4
//     1 <= n <= 5
//     grid[i][j] is either 0 or 1.

import "fmt"

func goodSubsetofBinaryMatrix(grid [][]int) []int {
    n := len(grid)
    nums := make([]int, n)
    for i, block := range(grid) {
        num, power := 0, 1
        for _,item := range(block) {
            num += item * power
            power *= 2
        }
        nums[i] = num
    }
    // fmt.Println(nums)
    if n == 1 && nums[0] == 0{
        return []int{0}
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if nums[i] & nums[j] == 0 {
                return []int{i,j}
            }
        }
    }       
    return []int{}
}

func goodSubsetofBinaryMatrix1(grid [][]int) []int {
    match := make([]int, 1 << len(grid[0])) // 列值
    for i := range match {
        match[i] = -1
    }
    for i := range grid {
        cur := 0
        for j, n := range grid[i] { // 求当前行的列值
            cur += n * (1 << j)
        } 
        if cur == 0 { // 列值为0，即该行全部是0，直接返回该行
            return []int{i}
        }
        for j, m := range match {
            if m != -1 && cur&j == 0 { // 若该列值与前面某一行的列值与运算为0，则返回
                return []int{m, i}
            }
        } 
        if match[cur] == -1 { // 若该列值没有出现过，记录该列值第一次出现的行
            match[cur] = i
        } 
    }
    return []int{} // 无解，没有任意2列的与值为0
}

func goodSubsetofBinaryMatrix2(grid [][]int) []int {
    res, mp := []int{}, make(map[int]int)
    m, n := len(grid), len(grid[0])
    for j := 0; j < m; j++ {
        st := 0
        for i := 0; i < n; i++ {
            st |= (grid[j][i] << i)
        }
        mp[st] = j
    }
    if _, ok := mp[0]; ok {
        res = append(res, mp[0])
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for x, i := range mp {
        for y, j := range mp {
            if x & y == 0 {
                return []int{min(i, j), max(i, j)}
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[0,1,1,0],[0,0,0,1],[1,1,1,1]]
    // Output: [0,1]
    // Explanation: We can choose the 0th and 1st rows to create a good subset of rows.
    // The length of the chosen subset is 2.
    // - The sum of the 0th column is 0 + 0 = 0, which is at most half of the length of the subset.
    // - The sum of the 1st column is 1 + 0 = 1, which is at most half of the length of the subset.
    // - The sum of the 2nd column is 1 + 0 = 1, which is at most half of the length of the subset.
    // - The sum of the 3rd column is 0 + 1 = 1, which is at most half of the length of the subset.
    fmt.Println(goodSubsetofBinaryMatrix([][]int{{0,1,1,0},{0,0,0,1},{1,1,1,1}})) // [0,1]
    // Example 2:
    // Input: grid = [[0]]
    // Output: [0]
    // Explanation: We can choose the 0th row to create a good subset of rows.
    // The length of the chosen subset is 1.
    // - The sum of the 0th column is 0, which is at most half of the length of the subset.
    fmt.Println(goodSubsetofBinaryMatrix([][]int{{0}})) // [0]
    // Example 3:
    // Input: grid = [[1,1,1],[1,1,1]]
    // Output: []
    // Explanation: It is impossible to choose any subset of rows to create a good subset.
    fmt.Println(goodSubsetofBinaryMatrix([][]int{{1,1,1}, {1,1,1}})) // []

    fmt.Println(goodSubsetofBinaryMatrix1([][]int{{0,1,1,0},{0,0,0,1},{1,1,1,1}})) // [0,1]
    fmt.Println(goodSubsetofBinaryMatrix1([][]int{{0}})) // [0]
    fmt.Println(goodSubsetofBinaryMatrix1([][]int{{1,1,1}, {1,1,1}})) // []

    fmt.Println(goodSubsetofBinaryMatrix2([][]int{{0,1,1,0},{0,0,0,1},{1,1,1,1}})) // [0,1]
    fmt.Println(goodSubsetofBinaryMatrix2([][]int{{0}})) // [0]
    fmt.Println(goodSubsetofBinaryMatrix2([][]int{{1,1,1}, {1,1,1}})) // []
}