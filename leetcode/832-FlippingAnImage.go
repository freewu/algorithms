package main

// 832. Flipping an Image
// Given an n x n binary matrix image, flip the image horizontally, then invert it, and return the resulting image.

// To flip an image horizontally means that each row of the image is reversed.
//     For example, flipping [1,1,0] horizontally results in [0,1,1].

// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
//     For example, inverting [0,1,1] results in [1,0,0].

// Example 1:
// Input: image = [[1,1,0],[1,0,1],[0,0,0]]
// Output: [[1,0,0],[0,1,0],[1,1,1]]
// Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
// Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

// Example 2:
// Input: image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
// Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
// Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
// Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

// Constraints:
//     n == image.length
//     n == image[i].length
//     1 <= n <= 20
//     images[i][j] is either 0 or 1.

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
    for _, row := range image {
        for i, j := 0, len(row) - 1; i < j; i, j = i + 1, j - 1 {
            row[i], row[j] = row[j], row[i] // 水平 翻转图像
        }
        for k, _ := range row {
            row[k] ^= 1 //  反转
        }
    }
    return image
}

func main() {
    // Example 1:
    // Input: image = [[1,1,0],[1,0,1],[0,0,0]]
    // Output: [[1,0,0],[0,1,0],[1,1,1]]
    // Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
    // Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
    fmt.Println(flipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}})) // [[1,0,0],[0,1,0],[1,1,1]]
    // Example 2: 
    // Input: image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
    // Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
    // Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
    // Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
    fmt.Println(flipAndInvertImage([][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}})) // [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
}