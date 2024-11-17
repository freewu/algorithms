package main

// 661. Image Smoother
// An image smoother is a filter of the size 3 x 3 
// that can be applied to each cell of an image by rounding down the average of the cell 
// and the eight surrounding cells (i.e., the average of the nine cells in the blue smoother). 
// If one or more of the surrounding cells of a cell is not present, 
// we do not consider it in the average (i.e., the average of the four cells in the red smoother).
// <img src="https://assets.leetcode.com/uploads/2021/05/03/smoother-grid.jpg" />

// Given an m x n integer matrix img representing the grayscale of an image, 
// return the image after applying the smoother on each cell of it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/smooth-grid.jpg" />
// Input: img = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[0,0,0],[0,0,0],[0,0,0]]
// Explanation:
// For the points (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
// For the points (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
// For the point (1,1): floor(8/9) = floor(0.88888889) = 0

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/smooth2-grid.jpg" />
// Input: img = [[100,200,100],[200,50,200],[100,200,100]]
// Output: [[137,141,137],[141,138,141],[137,141,137]]
// Explanation:
// For the points (0,0), (0,2), (2,0), (2,2): floor((100+200+200+50)/4) = floor(137.5) = 137
// For the points (0,1), (1,0), (1,2), (2,1): floor((200+200+50+200+100+100)/6) = floor(141.666667) = 141
// For the point (1,1): floor((50+200+200+200+200+100+100+100+100)/9) = floor(138.888889) = 138
 
// Constraints:
//     m == img.length
//     n == img[i].length
//     1 <= m, n <= 200
//     0 <= img[i][j] <= 255

import "fmt"

func imageSmoother(img [][]int) [][]int {
    processedImage := make([][]int, len(img))
    for i := range processedImage {
        processedImage[i] = make([]int, len(img[0]))
    }
    gray := func(img [][]int, i, j int) (int, bool) {
        if i < 0 || i >= len(img) || j < 0 || j >= len(img[0]) {
            return 0, false
        }
        return img[i][j], true
    }
    average := func (img [][]int, i, j int) int {
        count, summary := 0, 0
        for m := i-1; m <= i+1; m++ {
            for n := j - 1; n <= j+1; n++ {
                if v, exist := gray(img, m, n); exist {
                    count++
                    summary += v
                }
            }
        }
        return summary / count
    }
    for i := range processedImage {
        for j := range processedImage[i] {
            processedImage[i][j] = average(img, i, j)
        }
    }
    return processedImage
}

func imageSmoother1(img [][]int) [][]int {
    m, n := len(img), len(img[0])
    res := make([][]int, m)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range res {
        res[i] = make([]int, n)
        for j := range res[i] {
            sum, num := 0, 0
            for _, row := range img[max(i-1, 0):min(i+2, m)] {
                for _, v := range row[max(j-1, 0):min(j+2, n)] {
                    sum += v
                    num++
                }
            }
            res[i][j] = sum / num
        }
    }
    return res
}

func imageSmoother2(img [][]int) [][]int {
    n, m := len(img), len(img[0])
    res := make([][]int, n)
    directions := [][]int{{-1,-1}, {-1,0}, {-1,1}, {0,-1}, {0,1}, {1,-1}, {1,0}, {1,1}}
    for i := range img {
        res[i] = make([]int, m)
        for j := range img[i] {
            sum, count := img[i][j], 1
            for _, v := range directions {
                x, y := i + v[0], j + v[1]
                if x >= 0 && y >= 0 && x < n && y < m {
                    count++
                    sum += img[x][y]
                }
            }
            res[i][j] = sum / count
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/smooth-grid.jpg" />
    // Input: img = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: [[0,0,0],[0,0,0],[0,0,0]]
    // Explanation:
    // For the points (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
    // For the points (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
    // For the point (1,1): floor(8/9) = floor(0.88888889) = 0
    fmt.Println(imageSmoother([][]int{{1,1,1},{1,0,1},{1,1,1}})) // [[0,0,0],[0,0,0],[0,0,0]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/smooth2-grid.jpg" />
    // Input: img = [[100,200,100],[200,50,200],[100,200,100]]
    // Output: [[137,141,137],[141,138,141],[137,141,137]]
    // Explanation:
    // For the points (0,0), (0,2), (2,0), (2,2): floor((100+200+200+50)/4) = floor(137.5) = 137
    // For the points (0,1), (1,0), (1,2), (2,1): floor((200+200+50+200+100+100)/6) = floor(141.666667) = 141
    // For the point (1,1): floor((50+200+200+200+200+100+100+100+100)/9) = floor(138.888889) = 138
    fmt.Println(imageSmoother([][]int{{100,200,100},{200,50,200},{100,200,100}})) // [[137,141,137],[141,138,141],[137,141,137]]
    
    fmt.Println(imageSmoother1([][]int{{1,1,1},{1,0,1},{1,1,1}})) // [[0,0,0],[0,0,0],[0,0,0]]
    fmt.Println(imageSmoother1([][]int{{100,200,100},{200,50,200},{100,200,100}})) // [[137,141,137],[141,138,141],[137,141,137]]

    fmt.Println(imageSmoother2([][]int{{1,1,1},{1,0,1},{1,1,1}})) // [[0,0,0],[0,0,0],[0,0,0]]
    fmt.Println(imageSmoother2([][]int{{100,200,100},{200,50,200},{100,200,100}})) // [[137,141,137],[141,138,141],[137,141,137]]

}