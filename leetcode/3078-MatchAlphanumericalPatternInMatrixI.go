package main

// 3078. Match Alphanumerical Pattern in Matrix I
// You are given a 2D integer matrix board and a 2D character matrix pattern. 
// Where 0 <= board[r][c] <= 9 and each element of pattern is either a digit or a lowercase English letter.

// Your task is to find a submatrix of board that matches pattern.

// An integer matrix part matches pattern if we can replace cells containing letters in pattern 
// with some digits (each distinct letter with a unique digit) in such a way 
// that the resulting matrix becomes identical to the integer matrix part. 
// In other words,
//     1. The matrices have identical dimensions.
//     2. If pattern[r][c] is a digit, then part[r][c] must be the same digit.
//     3. If pattern[r][c] is a letter x:
//             For every pattern[i][j] == x, part[i][j] must be the same as part[r][c].
//             For every pattern[i][j] != x, part[i][j] must be different than part[r][c].

// Return an array of length 2 containing the row number and column number of the upper-left corner of a submatrix of board which matches pattern. 
// If there is more than one such submatrix, return the coordinates of the submatrix with the lowest row index, and in case there is still a tie, return the coordinates of the submatrix with the lowest column index. 
// If there are no suitable answers, return [-1, -1].

// Example 1:
// [1	2]	2   a	b
// [2	2]	3   b	b
// 2	3	3
// Input: board = [[1,2,2],[2,2,3],[2,3,3]], pattern = ["ab","bb"]
// Output: [0,0]
// Explanation: If we consider this mapping: "a" -> 1 and "b" -> 2; the submatrix with the upper-left corner (0,0) is a match as outlined in the matrix above.
// Note that the submatrix with the upper-left corner (1,1) is also a match but since it comes after the other one, we return [0,0].

// Example 2:
// 1	1	2   a	b
// 3	[3	4]   6	6
// 6	[6	6]
// Input: board = [[1,1,2],[3,3,4],[6,6,6]], pattern = ["ab","66"]
// Output: [1,1]
// Explanation: If we consider this mapping: "a" -> 3 and "b" -> 4; the submatrix with the upper-left corner (1,1) is a match as outlined in the matrix above.
// Note that since the corresponding values of "a" and "b" must differ, the submatrix with the upper-left corner (1,0) is not a match. Hence, we return [1,1].

// Example 3:
// 1	2   x	x
// 2	1
// Input: board = [[1,2],[2,1]], pattern = ["xx"]
// Output: [-1,-1]
// Explanation: Since the values of the matched submatrix must be the same, there is no match. Hence, we return [-1,-1].

// Constraints:
//     1 <= board.length <= 50
//     1 <= board[i].length <= 50
//     0 <= board[i][j] <= 9
//     1 <= pattern.length <= 50
//     1 <= pattern[i].length <= 50
//     pattern[i][j] is either a digit represented as a string or a lowercase English letter.

import "fmt"

func findPattern(board [][]int, pattern []string) []int {
    m, n, r, c := len(board), len(board[0]), len(pattern), len(pattern[0])
    check := func(i, j int) bool {
        d1, d2 := [26]int{}, [10]int{}
        for a := 0; a < r; a++ {
            for b := 0; b < c; b++ {
                x, y := i+a, j+b
                if pattern[a][b] >= '0' && pattern[a][b] <= '9' { // 0-9
                    if int(pattern[a][b] - '0') != board[x][y] { return false }
                } else { // a - z
                    v := int(pattern[a][b] - 'a')
                    if d1[v] > 0 && d1[v]-1 != board[x][y] { return false }
                    if d2[board[x][y]] > 0 && d2[board[x][y]]-1 != v { return false }
                    d1[v], d2[board[x][y]] = board[x][y] + 1, v + 1
                }
            }
        }
        return true
    }
    for i := 0; i < m - r + 1; i++ {
        for j := 0; j < n - c + 1; j++ {
            if check(i, j) {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}

func main() {
    // Example 1:
    // [1	2]	2   a	b
    // [2	2]	3   b	b
    // 2	3	3
    // Input: board = [[1,2,2],[2,2,3],[2,3,3]], pattern = ["ab","bb"]
    // Output: [0,0]
    // Explanation: If we consider this mapping: "a" -> 1 and "b" -> 2; the submatrix with the upper-left corner (0,0) is a match as outlined in the matrix above.
    // Note that the submatrix with the upper-left corner (1,1) is also a match but since it comes after the other one, we return [0,0].
    fmt.Println(findPattern([][]int{{1,2,2},{2,2,3},{2,3,3}}, []string{"ab","bb"})) // [0,0]
    // Example 2:
    // 1	1	2   a	b
    // 3	[3	4]   6	6
    // 6	[6	6]
    // Input: board = [[1,1,2],[3,3,4],[6,6,6]], pattern = ["ab","66"]
    // Output: [1,1]
    // Explanation: If we consider this mapping: "a" -> 3 and "b" -> 4; the submatrix with the upper-left corner (1,1) is a match as outlined in the matrix above.
    // Note that since the corresponding values of "a" and "b" must differ, the submatrix with the upper-left corner (1,0) is not a match. Hence, we return [1,1].
    fmt.Println(findPattern([][]int{{1,1,2},{3,3,4},{6,6,6}}, []string{"ab","66"})) // [1,1]
    // Example 3:
    // 1	2   x	x
    // 2	1
    // Input: board = [[1,2],[2,1]], pattern = ["xx"]
    // Output: [-1,-1]
    // Explanation: Since the values of the matched submatrix must be the same, there is no match. Hence, we return [-1,-1].
    fmt.Println(findPattern([][]int{{1,2},{2,1}}, []string{"xx"})) // [-1,-1]
}