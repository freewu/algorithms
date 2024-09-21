package main

// 1284. Minimum Number of Flips to Convert Binary Matrix to Zero Matrix
// Given a m x n binary matrix mat. 
// In one step, you can choose one cell and flip it and all the four neighbors of it if they exist (Flip is changing 1 to 0 and 0 to 1). 
// A pair of cells are called neighbors if they share one edge.

// Return the minimum number of steps required to convert mat to a zero matrix or -1 if you cannot.

// A binary matrix is a matrix with all cells equal to 0 or 1 only.

// A zero matrix is a matrix with all cells equal to 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/28/matrix.png" />
// Input: mat = [[0,0],[0,1]]
// Output: 3
// Explanation: One possible solution is to flip (1, 0) then (0, 1) and finally (1, 1) as shown.

// Example 2:
// Input: mat = [[0]]
// Output: 0
// Explanation: Given matrix is a zero matrix. We do not need to change it.

// Example 3:
// Input: mat = [[1,0,0],[1,0,0]]
// Output: -1
// Explanation: Given matrix cannot be a zero matrix.

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 3
//     mat[i][j] is either 0 or 1.

import "fmt"
import "strconv"
import "strings"

func minFlips(mat [][]int) int {
    //brute force BFS for every status
    //each status has m*n options to flip
    m, n := len(mat),len(mat[0])
    queue, set:= make([]string, 0, 0), make(map[string]bool)
    target, temp := "", ""
    flip := func(b byte)byte{
        if b == '1' { return '0' } 
        return '1'
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            temp = temp + strconv.Itoa(mat[i][j])
            target = target + "0"
        }
    }
    queue = append(queue, string(temp))
    set[string(temp)] = true
    step := 0
    for len(queue) > 0 {
        size := len(queue)
        for t := 0; t < size; t++ {
            cur := queue[t]
            if cur==target{
                return step
            }
            for i := 0; i < len(cur); i++ {
                next:=[]byte(cur)
                next[i] = flip(next[i])
                x := i / n
                y := i % n
                if x>0{
                    next[i-n] = flip(next[i-n])
                }
                if x<m-1{
                    next[i+n] = flip(next[i+n])
                }
                if y>0 {
                    next[i-1] = flip(next[i-1])
                }
                if y<n-1{
                    next[i+1] = flip(next[i+1])
                }
                nextstr := string(next)
                if !set[nextstr]{
                    set[nextstr] = true
                    queue = append(queue, nextstr)
                }
            }
        }
        step++
        queue = queue[size:]
    }
    return -1
}

func minFlips1(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    target := strings.Repeat("0", m * n)
    neighbours := func(pos int) []int {
        x, y, res := pos / n, pos % n, []int{}
        ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
        for _, d := range ds {
            nx, ny := x + d[0], y + d[1]
            if 0 <= nx && nx < m && 0 <= ny && ny < n {
                res = append(res, nx*n+ny)
            }
        }
        return res
    }
    getNextStatus := func(status string) []string {
        bs, res := []byte(status), []string{}
        for i := 0; i < m*n; i++ {
            indices2Flip := append(neighbours(i), i)
            for _, idx := range indices2Flip {
                if bs[idx] == '0' {
                    bs[idx] = '1'
                } else {
                    bs[idx] = '0'
                }
            }
            res = append(res, string(bs))
            for _, idx := range indices2Flip { // flip back
                if bs[idx] == '0' {
                    bs[idx] = '1'
                } else {
                    bs[idx] = '0'
                }
            }
        }
        return res
    }
    sb := make([]byte, 0, m*n)
    for _, rows := range mat {
        for _, c := range rows {
            sb = append(sb, byte('0'+c))
        }
    }
    start := string(sb)
    if start == target {
        return 0
    }
    seen := map[string]bool{}
    seen[start] = true
    type pair struct {
        status string
        step int
    }
    queue := []*pair{{start, 0}}
    for len(queue) > 0 {
        cur := queue[0] // pop
        queue = queue[1:]
        for _, nextStatus := range getNextStatus(cur.status) {
            if !seen[nextStatus] {
                if nextStatus == target {
                    return cur.step + 1
                }
                seen[nextStatus] = true
                queue = append(queue, &pair{nextStatus, cur.step+1})
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/11/28/matrix.png" />
    // Input: mat = [[0,0],[0,1]]
    // Output: 3
    // Explanation: One possible solution is to flip (1, 0) then (0, 1) and finally (1, 1) as shown.
    fmt.Println(minFlips([][]int{{0,0},{0,1}})) // 3
    // Example 2:
    // Input: mat = [[0]]
    // Output: 0
    // Explanation: Given matrix is a zero matrix. We do not need to change it.
    fmt.Println(minFlips([][]int{{0}})) // 0
    // Example 3:
    // Input: mat = [[1,0,0],[1,0,0]]
    // Output: -1
    // Explanation: Given matrix cannot be a zero matrix.
    fmt.Println(minFlips([][]int{{1,0,0},{1,0,0}})) // -1

    fmt.Println(minFlips1([][]int{{0,0},{0,1}})) // 3
    fmt.Println(minFlips1([][]int{{0}})) // 0
    fmt.Println(minFlips1([][]int{{1,0,0},{1,0,0}})) // -1
}