package main

// 1632. Rank Transform of a Matrix
// Given an m x n matrix, return a new matrix answer where answer[row][col] is the rank of matrix[row][col].

// The rank is an integer that represents how large an element is compared to other elements. 
// It is calculated using the following rules:
//     The rank is an integer starting from 1.
//     If two elements p and q are in the same row or column, then:
//         If p < q then rank(p) < rank(q)
//         If p == q then rank(p) == rank(q)
//         If p > q then rank(p) > rank(q)
//     The rank should be as small as possible.

// The test cases are generated so that answer is unique under the given rules.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/18/rank1.jpg" />
// Input: matrix = [[1,2],[3,4]]
// Output: [[1,2],[2,3]]
// Explanation:
// The rank of matrix[0][0] is 1 because it is the smallest integer in its row and column.
// The rank of matrix[0][1] is 2 because matrix[0][1] > matrix[0][0] and matrix[0][0] is rank 1.
// The rank of matrix[1][0] is 2 because matrix[1][0] > matrix[0][0] and matrix[0][0] is rank 1.
// The rank of matrix[1][1] is 3 because matrix[1][1] > matrix[0][1], matrix[1][1] > matrix[1][0], and both matrix[0][1] and matrix[1][0] are rank 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/18/rank2.jpg" />
// Input: matrix = [[7,7],[7,7]]
// Output: [[1,1],[1,1]]

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/18/rank3.jpg" />
// Input: matrix = [[20,-21,14],[-19,4,19],[22,-47,24],[-19,4,19]]
// Output: [[4,2,3],[1,3,4],[5,1,6],[1,3,4]]

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 500
//     -10^9 <= matrix[row][col] <= 10^9

import "fmt"
import "slices"

func matrixRankTransform(matrix [][]int) [][]int {
    n, m := len(matrix), len(matrix[0])
    trackers1, trackers2 :=  make([]int, n), make([]int, m) // rank trackers
    flags1, flags2 := make([]bool, n), make([]bool, m) // rank increment flags
    elems := make([][3]int, n * m) // elements queue
    for i := range matrix { // fill queue with elem values and coords
        for j, v := range matrix[i] {
            elems[i * m + j] = [3]int{ v, i, j }
        }
    }
    slices.SortFunc(elems, func(a, b [3]int) int { return a[0] - b[0] })
    for l, r := 0, 0; l < len(elems); l = r { // iterate over increasing elements
        for { // increment rank trackers if they weren't incremented already
            if !flags1[elems[r][1]] {
                trackers1[elems[r][1]]++
                flags1[elems[r][1]] = true
            }
            if !flags2[elems[r][2]] {
                trackers2[elems[r][2]]++
                flags2[elems[r][2]] = true
            }
            r++ 
            if r == len(elems) || elems[r-1][0] != elems[r][0] { break } // next elem isn't the same
        }
        // merge ranks:
        // rank trackers are related via elems coordinates,
        // so we iterate over each relation (elem) and
        // make ranks equal to their max
        // until no changes is needed.
        for update := true; update; {
            update = false
            for i := l; i < r; i++ {
                if trackers1[elems[i][1]] != trackers2[elems[i][2]] {
                    mx := max(trackers1[elems[i][1]], trackers2[elems[i][2]])
                    trackers1[elems[i][1]], trackers2[elems[i][2]] = mx, mx
                    update = true
                }
            }
        }
        for i := l; i < r; i++ { // set rank for each elem and clear tracker flags
            flags1[elems[i][1]], flags2[elems[i][2]] = false, false
            matrix[elems[i][1]][elems[i][2]] = trackers1[elems[i][1]]
        }
    }
    return matrix
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/18/rank1.jpg" />
    // Input: matrix = [[1,2],[3,4]]
    // Output: [[1,2],[2,3]]
    // Explanation:
    // The rank of matrix[0][0] is 1 because it is the smallest integer in its row and column.
    // The rank of matrix[0][1] is 2 because matrix[0][1] > matrix[0][0] and matrix[0][0] is rank 1.
    // The rank of matrix[1][0] is 2 because matrix[1][0] > matrix[0][0] and matrix[0][0] is rank 1.
    // The rank of matrix[1][1] is 3 because matrix[1][1] > matrix[0][1], matrix[1][1] > matrix[1][0], and both matrix[0][1] and matrix[1][0] are rank 2.
    fmt.Println(matrixRankTransform([][]int{{1,2},{3,4}})) // [[1,2],[2,3]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/18/rank2.jpg" />
    // Input: matrix = [[7,7],[7,7]]
    // Output: [[1,1],[1,1]]
    fmt.Println(matrixRankTransform([][]int{{7,7},{7,7}})) // [[1,1],[1,1]]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/10/18/rank3.jpg" />
    // Input: matrix = [[20,-21,14],[-19,4,19],[22,-47,24],[-19,4,19]]
    // Output: [[4,2,3],[1,3,4],[5,1,6],[1,3,4]]
    fmt.Println(matrixRankTransform([][]int{{20,-21,14},{-19,4,19},{22,-47,24},{-19,4,19}})) // [[4,2,3],[1,3,4],[5,1,6],[1,3,4]]
}