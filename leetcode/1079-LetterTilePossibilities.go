package main

// 1079. Letter Tile Possibilities
// You have n  tiles, where each tile has one letter tiles[i] printed on it.

// Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.

// Example 1:
// Input: tiles = "AAB"
// Output: 8
// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

// Example 2:
// Input: tiles = "AAABBC"
// Output: 188

// Example 3:
// Input: tiles = "V"
// Output: 1

// Constraints:
//     1 <= tiles.length <= 7
//     tiles consists of uppercase English letters.

import "fmt"
import "sort"

// back track
func numTilePossibilities(tiles string) int {
    res, n := 0, len(tiles)
    if n <= 1 { return n }
    mp := make(map[rune]int)
    for _, v := range tiles {
        mp[v]++
    }
    var backtrack func(pos int)
    backtrack = func(pos int) {
        if pos > 0 {
            res++
            if pos == n { return }
        }	
        for k, v := range mp {
            if v <= 0 { continue }
            mp[k]--
            backtrack(pos + 1)
            mp[k]++
        }
    }
    backtrack(0)
    return res
}

func numTilePossibilities1(tiles string) int {
    res, path, visited, tile := 0, []byte{}, make([]bool,len(tiles)), []byte(tiles)
    sort.Slice(tile,func(i,j int) bool {
        return tile[i] < tile[j]
    })
    var backtrack func()
    backtrack = func() {
        if len(path) > 0 {
            res++
        }
        for i:=0;i<len(tile);i++{
            if i > 0 && tile[i] == tile[i-1] && !visited[i-1] {
                continue
            }
            if !visited[i] {
                path = append(path,tile[i])
                visited[i] = true
                backtrack()
                path = path[:len(path)-1]
                visited[i] = false
            }
        }
    }
    backtrack()
    return res
}

func main() {
    // Example 1:
    // Input: tiles = "AAB"
    // Output: 8
    // Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
    fmt.Println(numTilePossibilities("AAB")) // 8 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"
    // Example 2:
    // Input: tiles = "AAABBC"
    // Output: 188
    fmt.Println(numTilePossibilities("AAABBC")) // 188 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"
    // Example 3:
    // Input: tiles = "V"
    // Output: 1
    fmt.Println(numTilePossibilities("V")) // 1 "V"

    fmt.Println(numTilePossibilities1("AAB")) // 8 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"
    fmt.Println(numTilePossibilities1("AAABBC")) // 188 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"
    fmt.Println(numTilePossibilities1("V")) // 1 "V"
}