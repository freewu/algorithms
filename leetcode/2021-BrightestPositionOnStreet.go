package main

// 2021. Brightest Position on Street
// A perfectly straight street is represented by a number line. 
// The street has street lamp(s) on it and is represented by a 2D integer array lights. 
// Each lights[i] = [positioni, rangei] indicates that there is a street lamp at position positioni 
// that lights up the area from [positioni - rangei, positioni + rangei] (inclusive).

// The brightness of a position p is defined as the number of street lamp that light up the position p.

// Given lights, return the brightest position on the street. 
// If there are multiple brightest positions, return the smallest one.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/28/image-20210928155140-1.png" />
// Input: lights = [[-3,2],[1,2],[3,3]]
// Output: -1
// Explanation:
// The first street lamp lights up the area from [(-3) - 2, (-3) + 2] = [-5, -1].
// The second street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].
// The third street lamp lights up the area from [3 - 3, 3 + 3] = [0, 6].
// Position -1 has a brightness of 2, illuminated by the first and second street light.
// Positions 0, 1, 2, and 3 have a brightness of 2, illuminated by the second and third street light.
// Out of all these positions, -1 is the smallest, so return it.

// Example 2:
// Input: lights = [[1,0],[0,1]]
// Output: 1
// Explanation:
// The first street lamp lights up the area from [1 - 0, 1 + 0] = [1, 1].
// The second street lamp lights up the area from [0 - 1, 0 + 1] = [-1, 1].
// Position 1 has a brightness of 2, illuminated by the first and second street light.
// Return 1 because it is the brightest position on the street.

// Example 3:
// Input: lights = [[1,2]]
// Output: -1
// Explanation:
// The first street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].
// Positions -1, 0, 1, 2, and 3 have a brightness of 1, illuminated by the first street light.
// Out of all these positions, -1 is the smallest, so return it.

// Constraints:
//     1 <= lights.length <= 10^5
//     lights[i].length == 2
//     -10^8 <= positioni <= 10^8
//     0 <= rangei <= 10^8

import "fmt"
import "sort"

func brightestPosition(lights [][]int) int {
    arr := make([]int, 0, len(lights) * 2)
    for _, v := range lights {
        position, scope := v[0], v[1]
        arr = append(arr, (position - scope) << 1 | 1, (position + scope + 1) << 1) // 最低位存储是区间左端点还是区间右端点
    }
    sort.Ints(arr)
    res, sum, mx := 0, 0, 0
    for i, v := range arr {
        sum += v & 1 << 1 - 1 // 根据最低位来 +1 或 -1
        if (i == len(arr) - 1 || arr[i+1] >> 1 != v >> 1) && sum > mx {
            mx = sum
            res = v >> 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/28/image-20210928155140-1.png" />
    // Input: lights = [[-3,2],[1,2],[3,3]]
    // Output: -1
    // Explanation:
    // The first street lamp lights up the area from [(-3) - 2, (-3) + 2] = [-5, -1].
    // The second street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].
    // The third street lamp lights up the area from [3 - 3, 3 + 3] = [0, 6].
    // Position -1 has a brightness of 2, illuminated by the first and second street light.
    // Positions 0, 1, 2, and 3 have a brightness of 2, illuminated by the second and third street light.
    // Out of all these positions, -1 is the smallest, so return it.
    fmt.Println(brightestPosition([][]int{{-3,2},{1,2},{3,3}})) // -1
    // Example 2:
    // Input: lights = [[1,0],[0,1]]
    // Output: 1
    // Explanation:
    // The first street lamp lights up the area from [1 - 0, 1 + 0] = [1, 1].
    // The second street lamp lights up the area from [0 - 1, 0 + 1] = [-1, 1].
    // Position 1 has a brightness of 2, illuminated by the first and second street light.
    // Return 1 because it is the brightest position on the street.
    fmt.Println(brightestPosition([][]int{{1,0},{0,1}})) // 1
    // Example 3:
    // Input: lights = [[1,2]]
    // Output: -1
    // Explanation:
    // The first street lamp lights up the area from [1 - 2, 1 + 2] = [-1, 3].
    // Positions -1, 0, 1, 2, and 3 have a brightness of 1, illuminated by the first street light.
    // Out of all these positions, -1 is the smallest, so return it.
    fmt.Println(brightestPosition([][]int{{1,2}})) // -1
}