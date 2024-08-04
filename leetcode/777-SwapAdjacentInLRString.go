package main

// 777. Swap Adjacent in LR String
// In a string composed of 'L', 'R', and 'X' characters, like "RXXLRXRXL", 
// a move consists of either replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR". 
// Given the starting string start and the ending string end, 
// return True if and only if there exists a sequence of moves to transform one string to the other.

// Example 1:
// Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
// Output: true
// Explanation: We can transform start to end following these steps:
// RXXLRXRXL ->
// XRXLRXRXL ->
// XRLXRXRXL ->
// XRLXXRRXL ->
// XRLXXRRLX

// Example 2:
// Input: start = "X", end = "L"
// Output: false

// Constraints:
//     1 <= start.length <= 10^4
//     start.length == end.length
//     Both start and end will only consist of characters in 'L', 'R', and 'X'.

import "fmt"

func canTransform(start string, end string) bool {
    currentL, currentR := 0, 0
    for i := 0; i < len(start); i++ {
        switch start[i] {
        case 'L': // meet L, R must be 0
            if currentR != 0 { return false }
            currentL--
        case 'R': // meet R, L must be 0
            if currentL != 0 { return false }
            currentR++
        }
        switch end[i] {
        case 'L': currentL++
        case 'R': currentR--
        }
        if currentL < 0 || currentR < 0 { return false }
        if currentL > 0 && currentR > 0 { return false }
    }
    return currentL == 0 && currentR == 0
}

func main() {
    // Example 1:
    // Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
    // Output: true
    // Explanation: We can transform start to end following these steps:
    // RXXLRXRXL ->
    // XRXLRXRXL ->
    // XRLXRXRXL ->
    // XRLXXRRXL ->
    // XRLXXRRLX
    fmt.Println(canTransform("RXXLRXRXL","XRLXXRRLX")) // true
    // Example 2:
    // Input: start = "X", end = "L"
    // Output: false
    fmt.Println(canTransform("X","L")) // true
}