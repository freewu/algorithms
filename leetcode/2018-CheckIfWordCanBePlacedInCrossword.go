package main

// 2018. Check if Word Can Be Placed In Crossword
// You are given an m x n matrix board, representing the current state of a crossword puzzle. 
// The crossword contains lowercase English letters (from solved words), ' ' to represent any empty cells, 
// and '#' to represent any blocked cells.

// A word can be placed horizontally (left to right or right to left) 
// or vertically (top to bottom or bottom to top) in the board if:
//     1. It does not occupy a cell containing the character '#'.
//     2. The cell each letter is placed in must either be ' ' (empty) or match the letter already on the board.
//     3. There must not be any empty cells ' ' or other lowercase letters directly left 
//        or right of the word if the word was placed horizontally.
//     4. There must not be any empty cells ' ' or other lowercase letters directly above 
//        or below the word if the word was placed vertically.

// Given a string word, return true if word can be placed in board, or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex1-1.png" />
// Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", "c", " "]], word = "abc"
// Output: true
// Explanation: The word "abc" can be placed as shown above (top to bottom).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex2-1.png" />
// Input: board = [[" ", "#", "a"], [" ", "#", "c"], [" ", "#", "a"]], word = "ac"
// Output: false
// Explanation: It is impossible to place the word because there will always be a space/letter above or below it.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex3-1.png" />
// Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", " ", "c"]], word = "ca"
// Output: true
// Explanation: The word "ca" can be placed as shown above (right to left). 

// Constraints:
//     m == board.length
//     n == board[i].length
//     1 <= m * n <= 2 * 10^5
//     board[i][j] will be ' ', '#', or a lowercase English letter.
//     1 <= word.length <= max(m, n)
//     word will contain only lowercase English letters.

import "fmt"

func placeWordInCrossword(board [][]byte, word string) bool {
    n, m := len(board), len(board[0])
    inRange := func(row, col, rows, cols int) bool { return row >= 0 && row < rows && col >= 0 && col < cols }
    isValid := func(row, col, rows, cols int) bool {
        // Is to check the left and right element is valid, when placing horizontally. And vise-versa
        if inRange(row, col, rows, cols) { return board[row][col] == '#' }
        return true
    }
    check := func(row, col, rows, cols, d1, d2 int) bool {
        i, l := 0, len(word)
        for i < l && inRange(row, col, rows, cols) && (board[row][col] == ' ' || board[row][col] == word[i]) {
            row += d1
            col += d2
            i++
        }
        if i < l { return false } // If all characters are not travelled in the word
        return isValid(row, col, rows, cols) // The ending char should be valid, either out of index or '#'
    }
    isPossible := func(row, col, rows, cols int) bool {
        if board[row][col] == '#' { return false } // The start character should not '#'
        // Check if word can be placed in left direction. For that you have to check if the right element of starting point should be either out of index, or should be '#'
        if isValid(row, col + 1, rows, cols) && check(row, col, rows, cols, 0, -1) { return true }
        // Check if word can be placed in right direction. For that you have to check if the right element of starting point should be either out of index, or should be '#'
        if isValid(row, col - 1, rows, cols) && check(row, col, rows, cols, 0, 1) { return true }
        // Check if word can be placed in top direction. For that you have to check if the bottom element of starting point should be either out of index, or should be '#'
        if isValid(row + 1, col, rows, cols) && check(row, col, rows, cols, -1, 0) { return true }
        // Check if word can be placed in bottom direction. For that you have to check if the top element of starting point should be either out of index, or should be '#'
        return isValid(row - 1, col, rows, cols) && check(row, col, rows, cols, 1, 0)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if isPossible(i, j, n, m) {
                return true
            }
        }
    }
    return false
}

func placeWordInCrossword1(board [][]byte, word string) bool {
    m, n, k := len(board), len(board[0]), len(word)
    // 遍历行
    for _, row := range board {
        for j := 0; j < n; j++ {
            if row[j] == '#' { continue }
            // 遍历并匹配两个 # 之间的字符
            j0, ok1, ok2 := j, true, true
            for ; j < n && row[j] != '#'; j++ { // 注意这里的 j 就是外层循环的 j，因此整体复杂度是线性的
                if j-j0 >= k || row[j] != ' ' && row[j] != word[j-j0] { // 正序匹配 word
                    ok1 = false
                }
                if j-j0 >= k || row[j] != ' ' && row[j] != word[k-1-j+j0] { // 倒序匹配 word
                    ok2 = false
                }
            }
            if (ok1 || ok2) && j - j0 == k { // 只要正序和倒序中有一个匹配成功，且两个 # 之间的字符长度恰好为 word 的长度，就返回 true
                return true
            }
        }
    }
    // 遍历列（同上）
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            if board[i][j] == '#' { continue }
            i0, ok1, ok2 := i, true, true
            for ; i < m && board[i][j] != '#'; i++ {
                if i-i0 >= k || board[i][j] != ' ' && board[i][j] != word[i-i0] {
                    ok1 = false
                }
                if i-i0 >= k || board[i][j] != ' ' && board[i][j] != word[k-1-i+i0] {
                    ok2 = false
                }
            }
            if (ok1 || ok2) && i - i0 == k {
                return true
            }
        }
    }
    return false
}

func placeWordInCrossword2(board [][]byte, word string) bool {
    m, n, k := len(board), len(board[0]), len(word)
    check := func(i, j, a, b int) bool {
        x, y := i + a * k, j + b * k
        if x >= 0 && x < m && y >= 0 && y < n && board[x][y] != '#' { return false }
        for _, c := range word {
            if i < 0 || i >= m || j < 0 || j >= n || (board[i][j] != ' ' && board[i][j] != byte(c)) { return false }
            i, j = i + a, j + b
        }
        return true
    }
    for i := range board {
        for j := range board[i] {
            leftToRight := (j == 0 || board[i][j-1] == '#') && check(i, j, 0, 1)
            rightToLeft := (j == n-1 || board[i][j+1] == '#') && check(i, j, 0, -1)
            upToDown := (i == 0 || board[i-1][j] == '#') && check(i, j, 1, 0)
            downToUp := (i == m-1 || board[i+1][j] == '#') && check(i, j, -1, 0)
            if leftToRight || rightToLeft || upToDown || downToUp {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex1-1.png" />
    // Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", "c", " "]], word = "abc"
    // Output: true
    // Explanation: The word "abc" can be placed as shown above (top to bottom).
    fmt.Println(placeWordInCrossword([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}, "abc")) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex2-1.png" />
    // Input: board = [[" ", "#", "a"], [" ", "#", "c"], [" ", "#", "a"]], word = "ac"
    // Output: false
    // Explanation: It is impossible to place the word because there will always be a space/letter above or below it.
    fmt.Println(placeWordInCrossword([][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}, "ac")) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/10/04/crossword-ex3-1.png" />
    // Input: board = [["#", " ", "#"], [" ", " ", "#"], ["#", " ", "c"]], word = "ca"
    // Output: true
    // Explanation: The word "ca" can be placed as shown above (right to left). 
    fmt.Println(placeWordInCrossword([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', ' ', 'c'}}, "ca")) // true

    fmt.Println(placeWordInCrossword1([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}, "abc")) // true
    fmt.Println(placeWordInCrossword1([][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}, "ac")) // false
    fmt.Println(placeWordInCrossword1([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', ' ', 'c'}}, "ca")) // true

    fmt.Println(placeWordInCrossword2([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}, "abc")) // true
    fmt.Println(placeWordInCrossword2([][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}, "ac")) // false
    fmt.Println(placeWordInCrossword2([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', ' ', 'c'}}, "ca")) // true
}