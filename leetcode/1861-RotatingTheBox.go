package main

// 1861. Rotating the Box
// You are given an m x n matrix of characters box representing a side-view of a box. 
// Each cell of the box is one of the following:
//     A stone '#'
//     A stationary obstacle '*'
//     Empty '.'

// The box is rotated 90 degrees clockwise, causing some of the stones to fall due to gravity. 
// Each stone falls down until it lands on an obstacle, another stone, or the bottom of the box. 
// Gravity does not affect the obstacles' positions, and the inertia from the box's rotation does not affect the stones' horizontal positions.

// It is guaranteed that each stone in box rests on an obstacle, another stone, or the bottom of the box.

// Return an n x m matrix representing the box after the rotation described above.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcodewithstones.png" />
// Input: box = [["#",".","#"]]
// Output: [["."],
//          ["#"],
//          ["#"]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode2withstones.png" />
// Input: box = [["#",".","*","."],
//               ["#","#","*","."]]
// Output: [["#","."],
//          ["#","#"],
//          ["*","*"],
//          [".","."]]

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode3withstone.png" />
// Input: box = [["#","#","*",".","*","."],
//               ["#","#","#","*",".","."],
//               ["#","#","#",".","#","."]]
// Output: [[".","#","#"],
//          [".","#","#"],
//          ["#","#","*"],
//          ["#","*","."],
//          ["#",".","*"],
//          ["#",".","."]]

// Constraints:
//     m == box.length
//     n == box[i].length
//     1 <= m, n <= 500
//     box[i][j] is either '#', '*', or '.'.

import "fmt"

func rotateTheBox(box [][]byte) [][]byte {
    m, n := len(box[0]), len(box)
    res := make([][]byte, m)
    for i := range res {
        res[i] = make([]byte, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for r := n - 1; r >= 0; r-- {
        t, j := n - 1 - r, m - 1
        for i := m - 1; i >= 0; i-- {
            res[i][r] = box[t][i]
            if res[i][r] == '.' {
                j = min(j, i - 1)
                for j >= 0 && box[t][j] != '*' {
                    if box[t][j] == '#' {
                        box[t][j] = '.'
                        res[j][r] = '.'
                        res[i][r] = '#'
                        break
                    }
                    j--
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcodewithstones.png" />
    // Input: box = [["#",".","#"]]
    // Output: [["."],
    //          ["#"],
    //          ["#"]]
    box1 := [][]byte{{'#','.','#'}}
    fmt.Println(rotateTheBox(box1)) // 
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode2withstones.png" />
    // Input: box = [["#",".","*","."],
    //               ["#","#","*","."]]
    // Output: [["#","."],
    //          ["#","#"],
    //          ["*","*"],
    //          [".","."]]
    box2 := [][]byte{{'#','.','*','.'}, {'#','#','*','.'}}
    fmt.Println(rotateTheBox(box2)) // 
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/rotatingtheboxleetcode3withstone.png" />
    // Input: box = [["#","#","*",".","*","."],
    //               ["#","#","#","*",".","."],
    //               ["#","#","#",".","#","."]]
    // Output: [[".","#","#"],
    //          [".","#","#"],
    //          ["#","#","*"],
    //          ["#","*","."],
    //          ["#",".","*"],
    //          ["#",".","."]]
    box3 := [][]byte{{'#','#','*','.','*','.'}, {'#','#','#','*','.','.'}, {'#','#','#','.','#','.'}}
    fmt.Println(rotateTheBox(box3)) // 
}