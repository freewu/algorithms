package main

// 面试题 08.10. Color Fill LCCI
// Implement the "paint fill" function that one might see on many image editing programs. 
// That is, given a screen (represented by a two-dimensional array of colors), a point, and a new color, fill in the surrounding area until the color changes from the original color.

// Example1:
// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: 
// From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected 
// by a path of the same color as the starting pixel are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected
// to the starting pixel.

// Note:
//     The length of image and image[0] will be in the range [1, 50].
//     The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
//     The value of each color in image[i][j] and newColor will be an integer in [0, 65535].

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    n, m := len(image), len(image[0])
    queue, same := [][]int{{ sr, sc }}, image[sr][sc]
    image[sr][sc] = newColor
    visited := make([][]bool, n)
    for i := range image {
        visited[i] = make([]bool, m)
    }
    visited[sr][sc] = true
    for len(queue) > 0 {
        tmp := queue[0]
        queue = queue[1:]
        for _, v := range [][]int{ {1,0},{0,1},{-1, 0},{0,-1}} {
            x, y := tmp[0] + v[0], tmp[1] + v[1]
            if x >= 0 && x <= n- 1 && y >= 0 && y <= m - 1 && !visited[x][y] {
                visited[x][y] = true
                if image[x][y] == same {
                    image[x][y] = newColor
                    queue = append(queue, []int{x, y})
                } 
            }
        }
    }
    return image
}

func main() {
    // Example1:
    // Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
    // Output: [[2,2,2],[2,2,0],[2,0,1]]
    // Explanation: 
    // From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected 
    // by a path of the same color as the starting pixel are colored with the new color.
    // Note the bottom corner is not colored 2, because it is not 4-directionally connected
    // to the starting pixel.
    fmt.Println(floodFill([][]int{{1,1,1},{1,1,0},{1,0,1}}, 1, 1, 2)) // [[2,2,2],[2,2,0],[2,0,1]]
}