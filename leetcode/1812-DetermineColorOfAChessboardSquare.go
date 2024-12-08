package main

// 1812. Determine Color of a Chessboard Square
// You are given coordinates, a string that represents the coordinates of a square of the chessboard. 
// Below is a chessboard for your reference.
// <img src="https://assets.leetcode.com/uploads/2021/02/19/screenshot-2021-02-20-at-22159-pm.png" />

// Return true if the square is white, and false if the square is black.

// The coordinate will always represent a valid chessboard square. 
// The coordinate will always have the letter first, and the number second.

// Example 1:
// Input: coordinates = "a1"
// Output: false
// Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.

// Example 2:
// Input: coordinates = "h3"
// Output: true
// Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.

// Example 3:
// Input: coordinates = "c7"
// Output: false

// Constraints:
//     coordinates.length == 2
//     'a' <= coordinates[0] <= 'h'
//     '1' <= coordinates[1] <= '8'

import "fmt"

func squareIsWhite(coordinates string) bool {
    if int(coordinates[0] - 'a') % 2 == 0 { // 偶数列 奇数为 白 ( 0 开头)
        return int(coordinates[1] - '0') % 2 == 0
    }
    return int(coordinates[1] - '0') % 2 == 1
}

func squareIsWhite1(coordinates string) bool {
    return (coordinates[0] ^ coordinates[1]) & 1 == 1
}

func main() {
    // Example 1:
    // Input: coordinates = "a1"
    // Output: false
    // Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.
    fmt.Println(squareIsWhite("a1")) // false
    // Example 2:
    // Input: coordinates = "h3"
    // Output: true
    // Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.
    fmt.Println(squareIsWhite("h3")) // true
    // Example 3:
    // Input: coordinates = "c7"
    // Output: false
    fmt.Println(squareIsWhite("c7")) // false

    fmt.Println(squareIsWhite1("a1")) // false
    fmt.Println(squareIsWhite1("h3")) // true
    fmt.Println(squareIsWhite1("c7")) // false
}