package main

// 733. Flood Fill
// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

// You are also given three integers sr, sc, and color. 
// You should perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, 
// plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, 
// plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. 
// Replace the color of all of the aforementioned pixels with color.

// Return the modified image after performing the flood fill.

// Example 1:
// <img src="" />
// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

// Example 2:
// Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
// Output: [[0,0,0],[0,0,0]]
// Explanation: The starting pixel is already colored 0, so no changes are made to the image.

// Constraints:
//     m == image.length
//     n == image[i].length
//     1 <= m, n <= 50
//     0 <= image[i][j], color < 2^16
//     0 <= sr < m
//     0 <= sc < n

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    oldColor := image[sr][sc]
    var changeColor func(image [][]int, x, y int, oldColor, newColor *int) 
    changeColor = func(image [][]int, x, y int, oldColor, newColor *int) {
        if image[x][y] == *oldColor {
            image[x][y] = *newColor
            if x >= 1 {
                changeColor(image, x-1, y, oldColor, newColor)
            }
            if y >= 1 {
                changeColor(image, x, y-1, oldColor, newColor)
            }
            if x+1 < len(image) {
                changeColor(image, x+1, y, oldColor, newColor)
            }
            if y+1 < len(image[0]) {
                changeColor(image, x, y+1, oldColor, newColor)
            }
        }
    }
    if oldColor != color {
        changeColor(image, sr, sc, &oldColor, &color)
    }
    return image
}

func floodFill1(image [][]int, sr int, sc int, color int) [][]int {
    var dfs func(sr int, sc int, oldColor int, newColor int)
    dfs = func(sr int, sc int, oldColor int, newColor int) {
        if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != oldColor {
            return
        }
        image[sr][sc] = newColor
        dfs(sr+1, sc, oldColor, newColor)
        dfs(sr-1, sc, oldColor, newColor)
        dfs(sr, sc+1, oldColor, newColor)
        dfs(sr, sc-1, oldColor, newColor)
    }
    if image[sr][sc] != color {
        dfs(sr, sc, image[sr][sc], color)
    }
    return image
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
    // Output: [[2,2,2],[2,2,0],[2,0,1]]
    // Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
    // Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
    image1 := [][]int{
        {1,1,1},
        {1,1,0},
        {1,0,1},
    }
    fmt.Println(floodFill(image1,1,1,2)) // [[2,2,2],[2,2,0],[2,0,1]]
    // Example 2:
    // Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
    // Output: [[0,0,0],[0,0,0]]
    // Explanation: The starting pixel is already colored 0, so no changes are made to the image.
    image2 := [][]int{
        {0,0,0},
        {0,0,0},
    }
    fmt.Println(floodFill(image2,0,0,0)) // [[0,0,0],[0,0,0]]

    image11 := [][]int{
        {1,1,1},
        {1,1,0},
        {1,0,1},
    }
    fmt.Println(floodFill1(image11,1,1,2)) // [[2,2,2],[2,2,0],[2,0,1]]
    image12 := [][]int{
        {0,0,0},
        {0,0,0},
    }
    fmt.Println(floodFill1(image12,0,0,0)) // [[0,0,0],[0,0,0]]
}