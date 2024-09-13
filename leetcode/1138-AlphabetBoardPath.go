package main

// 1138. Alphabet Board Path
// On an alphabet board, we start at position (0, 0), corresponding to character board[0][0].

// Here, board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"], as shown in the diagram below.
// <img src="https://assets.leetcode.com/uploads/2019/07/28/azboard.png" />

// We may make the following moves:
//     'U' moves our position up one row, if the position exists on the board;
//     'D' moves our position down one row, if the position exists on the board;
//     'L' moves our position left one column, if the position exists on the board;
//     'R' moves our position right one column, if the position exists on the board;
//     '!' adds the character board[r][c] at our current position (r, c) to the answer.

// (Here, the only positions that exist on the board are positions with letters on them.)

// Return a sequence of moves that makes our answer equal to target in the minimum number of moves.  
// You may return any path that does so.

// Example 1:
// Input: target = "leet"
// Output: "DDR!UURRR!!DDD!"

// Example 2:
// Input: target = "code"
// Output: "RR!DDRR!UUL!R!"

// Constraints:
//     1 <= target.length <= 100
//     target consists only of English lowercase letters.

import "fmt"

func alphabetBoardPath(target string) string {
    res, chars := []byte{}, []byte(target)
    x, y := 0, 0
    move := func(dir byte, count int) {
        for count > 0 {
            res = append(res, dir)
            count--
        }
    }
    for i := 0; i < len(chars); i++ {
        position := int(chars[i] - 'a')
        nx, ny := position / 5, position % 5
        move('U', x - nx)
        move('L', y - ny)
        move('D', nx - x)
        move('R', ny - y)
        res = append(res, '!')
        x, y = nx, ny
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: target = "leet"
    // Output: "DDR!UURRR!!DDD!"
    fmt.Println(alphabetBoardPath("leet")) // "DDR!UURRR!!DDD!"
    // Example 2:
    // Input: target = "code"
    // Output: "RR!DDRR!UUL!R!"
    fmt.Println(alphabetBoardPath("code")) // "RR!DDRR!UUL!R!"
}