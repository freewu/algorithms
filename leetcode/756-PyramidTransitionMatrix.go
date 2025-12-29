package main

// 756. Pyramid Transition Matrix
// You are stacking blocks to form a pyramid. 
// Each block has a color, which is represented by a single letter. 
// Each row of blocks contains one less block than the row beneath it and is centered on top.

// To make the pyramid aesthetically pleasing, there are only specific triangular patterns that are allowed. 
// A triangular pattern consists of a single block stacked on top of two blocks. 
// The patterns are given as a list of three-letter strings allowed, where the first two characters of a pattern represent the left and right bottom blocks respectively, and the third character is the top block.
//     For example, "ABC" represents a triangular pattern with a 'C' block stacked on top of an 'A' (left) and 'B' (right) block. 
//     Note that this is different from "BAC" where 'B' is on the left bottom and 'A' is on the right bottom.

// You start with a bottom row of blocks bottom, given as a single string, 
// that you must use as the base of the pyramid.

// Given bottom and allowed, return true if you can build the pyramid all the way to the top such that every triangular pattern in the pyramid is in allowed, or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/26/pyramid1-grid.jpg" />
// Input: bottom = "BCD", allowed = ["BCC","CDE","CEA","FFF"]
// Output: true
// Explanation: The allowed triangular patterns are shown on the right.
// Starting from the bottom (level 3), we can build "CE" on level 2 and then build "A" on level 1.
// There are three triangular patterns in the pyramid, which are "BCC", "CDE", and "CEA". All are allowed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/26/pyramid2-grid.jpg" />
// Input: bottom = "AAAA", allowed = ["AAB","AAC","BCD","BBE","DEF"]
// Output: false
// Explanation: The allowed triangular patterns are shown on the right.
// Starting from the bottom (level 4), there are multiple ways to build level 3, but trying all the possibilites, you will get always stuck before building level 1.

// Constraints:
//     2 <= bottom.length <= 6
//     0 <= allowed.length <= 2^16
//     allowed[i].length == 3
//     The letters in all input strings are from the set {'A', 'B', 'C', 'D', 'E', 'F'}.
//     All the values of allowed are unique.

import "fmt"

// // dfs
// func pyramidTransition(bottom string, allowed []string) bool {
//     nexts := map[string][]byte{}
//     for _, v := range allowed {
//         nexts[v[:2]] = append(nexts[v[:2]], v[2])
//     }
//     var dfs func(cur []byte, i int) bool
//     dfs = func(cur []byte, i int) bool {
//         if len(cur) == 1 {
//             return true
//         }
//         if len(cur) == i+1 {
//             return dfs(cur[:len(cur)-1], 0)
//         }
//         s := string(cur[i : i+2])
//         for _, c := range nexts[s] {
//             cur[i] = c // no need to backtrack
//             if dfs(cur, i+1) {
//                 return true
//             }
//         }
//         return false
//     }
//     return dfs([]byte(bottom), 0)
// }

func pyramidTransition(bottom string, allowed []string) bool {
    type base struct {
        left, right byte
    }
    var isPossibleToBuild func(bottom, nextLevel []byte, index map[base][]byte) bool
    isPossibleToBuild = func(bottom, nextLevel []byte, index map[base][]byte) bool {
        if len(bottom) == 1 {
            if len(nextLevel) == 0 {
                return true
            }
            return isPossibleToBuild(nextLevel, []byte{}, index)
        }
        b := base{left:bottom[0], right:bottom[1]}
        tt, _ := index[b]
        for _, t := range tt {
            if isPossibleToBuild(bottom[1:], append(nextLevel, t), index) {
                return true
            }
        }
        return false
    }
    index := make(map[base][]byte)
    for _, triple := range allowed {
        b := base{left:triple[0], right:triple[1]}
        t := triple[2]
        index[b] = append(index[b], t)
    }
    return isPossibleToBuild(([]byte)(bottom), []byte{}, index)
}

func pyramidTransition1(bottom string, allowed []string) bool {
    groups := [7][7][]byte{}
    for _, s := range allowed {
        a, b := s[0]&31, s[1]&31 // A~F -> 1~6
        groups[a][b] = append(groups[a][b], s[2]&31)
    }
    n := len(bottom)
    pyramid := make([]int, n)
    for i, ch := range bottom {
        pyramid[n-1] |= int(ch&31) << (i * 3) // 等价于 pyramid[n-1][i] = ch&31
    }
    visted := make([]bool, 1<<((n-1)*3))
    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i < 0 { return true }
        if visted[pyramid[i]] { return false }
        if j == i+1 {
            visted[pyramid[i]] = true
            return dfs(i-1, 0)
        }
        for _, top := range groups[pyramid[i+1]>>(j*3)&7][pyramid[i+1]>>((j+1)*3)&7] {
            pyramid[i] &^= 7 << (j * 3) // 清除之前填的字母，等价于 pyramid[i][j] = 0
            pyramid[i] |= int(top) << (j * 3) // 等价于 pyramid[i][j] = top
            if dfs(i, j + 1) {
                return true
            }
        }
        return false
    }
    return dfs(n - 2, 0)
}


func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/26/pyramid1-grid.jpg" />
    // Input: bottom = "BCD", allowed = ["BCC","CDE","CEA","FFF"]
    // Output: true
    // Explanation: The allowed triangular patterns are shown on the right.
    // Starting from the bottom (level 3), we can build "CE" on level 2 and then build "A" on level 1.
    // There are three triangular patterns in the pyramid, which are "BCC", "CDE", and "CEA". All are allowed.
    fmt.Println(pyramidTransition("BCD",[]string{"BCC","CDE","CEA","FFF"})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/26/pyramid2-grid.jpg" />
    // Input: bottom = "AAAA", allowed = ["AAB","AAC","BCD","BBE","DEF"]
    // Output: false
    // Explanation: The allowed triangular patterns are shown on the right.
    // Starting from the bottom (level 4), there are multiple ways to build level 3, but trying all the possibilites, you will get always stuck before building level 1.
    fmt.Println(pyramidTransition("AAAA",[]string{"AAB","AAC","BCD","BBE","DEF"})) // false

    fmt.Println(pyramidTransition("DBCDA",[]string{"DBD","BCC","CDD","DAD","DDA","AAC","CCA","BCD"})) // true

    fmt.Println(pyramidTransition1("BCD",[]string{"BCC","CDE","CEA","FFF"})) // true
    fmt.Println(pyramidTransition1("AAAA",[]string{"AAB","AAC","BCD","BBE","DEF"})) // false
    fmt.Println(pyramidTransition1("DBCDA",[]string{"DBD","BCC","CDD","DAD","DDA","AAC","CCA","BCD"})) // true
}