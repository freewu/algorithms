package main

// 2271. Maximum White Tiles Covered by a Carpet
// You are given a 2D integer array tiles where tiles[i] = [li, ri] represents 
// that every tile j in the range li <= j <= ri is colored white.

// You are also given an integer carpetLen, the length of a single carpet that can be placed anywhere.

// Return the maximum number of white tiles that can be covered by the carpet.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/25/example1drawio3.png" />
// Input: tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
// Output: 9
// Explanation: Place the carpet starting on tile 10. 
// It covers 9 white tiles, so we return 9.
// Note that there may be other places where the carpet covers 9 white tiles.
// It can be shown that the carpet cannot cover more than 9 white tiles.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/24/example2drawio.png" />
// Input: tiles = [[10,11],[1,1]], carpetLen = 2
// Output: 2
// Explanation: Place the carpet starting on tile 10. 
// It covers 2 white tiles, so we return 2.

// Constraints:
//     1 <= tiles.length <= 5 * 10^4
//     tiles[i].length == 2
//     1 <= li <= ri <= 10^9
//     1 <= carpetLen <= 10^9
//     The tiles are non-overlapping.

import "fmt"
import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
    sort.Slice(tiles, func(i, j int) bool {
        if tiles[i][0] == tiles[j][0] { return tiles[i][1] < tiles[j][1] }
        return tiles[i][0] < tiles[j][0]
    })
    n, val, res := len(tiles), 0, 0
    psum := make([]int, n)
    for i := 0; i < n; i++ {
        val += (tiles[i][1] - tiles[i][0] + 1)
        psum[i] = val
    }
    binarySearch := func(tiles [][]int, find int) int{
        start, end := 0, len(tiles)
        for start < end {
            mid := start + (end - start) / 2
            if tiles[mid][0] <= find {
                start = mid + 1
            } else {
                end = mid
            }
        }
        return start
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i:= 0; i < n; i++ {
        start := tiles[i][0]
        find := start + carpetLen - 1
        index := binarySearch(tiles, find)
        sum := 0
        if index - 2 >= 0 && i - 1 >= 0 { sum = psum[index - 2] - psum[i - 1] }
        if index - 2 >= 0  && i == 0 { sum = psum[index-2] }
        if find >= tiles[index - 1][1] {
            sum += (tiles[index - 1][1] - tiles[index - 1][0] + 1)
        } else {
            sum += (find - tiles[index - 1][0] + 1)
        }  
        res = max(res,sum)
    }
    return res
}

func maximumWhiteTiles1(tiles [][]int, carpetLen int) int {
    sort.Slice(tiles, func(i, j int) bool {
        return tiles[i][0] < tiles[j][0]
    })
    res, n := 0, len(tiles)
    s, j := make([]int, n + 1), n - 1
    for i, v := range tiles {
        s[i+1] = s[i] + v[1] - v[0] + 1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        cur := tiles[i]
        for j > 0 && cur[1]-tiles[j][0]+1 < carpetLen { j-- }
        // 左边界
        mostCoverLeft := max(cur[1]-carpetLen+1, 0)
        sum := s[i+1] - s[j+1] + max(tiles[j][1] - max(mostCoverLeft, tiles[j][0]) + 1, 0)
        res = max(res, sum)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/25/example1drawio3.png" />
    // Input: tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
    // Output: 9
    // Explanation: Place the carpet starting on tile 10. 
    // It covers 9 white tiles, so we return 9.
    // Note that there may be other places where the carpet covers 9 white tiles.
    // It can be shown that the carpet cannot cover more than 9 white tiles.
    fmt.Println(maximumWhiteTiles([][]int{{1,5},{10,11},{12,18},{20,25},{30,32}}, 10)) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/24/example2drawio.png" />
    // Input: tiles = [[10,11],[1,1]], carpetLen = 2
    // Output: 2
    // Explanation: Place the carpet starting on tile 10. 
    // It covers 2 white tiles, so we return 2.
    fmt.Println(maximumWhiteTiles([][]int{{10,11},{1,1}}, 2)) // 2

    fmt.Println(maximumWhiteTiles1([][]int{{1,5},{10,11},{12,18},{20,25},{30,32}}, 10)) // 9
    fmt.Println(maximumWhiteTiles1([][]int{{10,11},{1,1}}, 2)) // 2
}