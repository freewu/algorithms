package main

// 2647. Color the Triangle Red
// You are given an integer n. Consider an equilateral triangle of side length n, broken up into n2 unit equilateral triangles. 
// The triangle has n 1-indexed rows where the ith row has 2i - 1 unit equilateral triangles.

// The triangles in the ith row are also 1-indexed with coordinates from (i, 1) to (i, 2i - 1). 
// The following image shows a triangle of side length 4 with the indexing of its triangle.
// <img src="https://assets.leetcode.com/uploads/2022/09/01/triangle4.jpg" />

// Two triangles are neighbors if they share a side. For example:
//     Triangles (1,1) and (2,2) are neighbors
//     Triangles (3,2) and (3,3) are neighbors.
//     Triangles (2,2) and (3,3) are not neighbors because they do not share any side.

// Initially, all the unit triangles are white. You want to choose k triangles and color them red. 
// We will then run the following algorithm:
//     1. Choose a white triangle that has at least two red neighbors.
//         If there is no such triangle, stop the algorithm.
//     2. Color that triangle red.
//     3. Go to step 1.

// Choose the minimum k possible and set k triangles red before running this algorithm such that after the algorithm stops, all unit triangles are colored red.

// Return a 2D list of the coordinates of the triangles that you will color red initially. 
// The answer has to be of the smallest size possible. 
// If there are multiple valid solutions, return any.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/01/example1.jpg" />
// Input: n = 3
// Output: [[1,1],[2,1],[2,3],[3,1],[3,5]]
// Explanation: Initially, we choose the shown 5 triangles to be red. Then, we run the algorithm:
// - Choose (2,2) that has three red neighbors and color it red.
// - Choose (3,2) that has two red neighbors and color it red.
// - Choose (3,4) that has three red neighbors and color it red.
// - Choose (3,3) that has three red neighbors and color it red.
// It can be shown that choosing any 4 triangles and running the algorithm will not make all triangles red.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/01/example2.jpg" />
// Input: n = 2
// Output: [[1,1],[2,1],[2,3]]
// Explanation: Initially, we choose the shown 3 triangles to be red. Then, we run the algorithm:
// - Choose (2,2) that has three red neighbors and color it red.
// It can be shown that choosing any 2 triangles and running the algorithm will not make all triangles red.

// Constraints:
//     1 <= n <= 1000

import "fmt"

// // 解答错误 4 / 18
// func colorRed(n int) [][]int {
//     res, size := [][]int{}, 2 * n 
//     for i := n; i > 1; i -= 4 {
//         for j := 1; j < size; j += 2 {
//             res = append(res, []int{i, j})
//         }
//         if i >= 3 {
//             res = append(res, []int{i - 1, 2})
//         }
//         i -= 2
//         if i >= 2 {
//             for j := 3; j < size - 4; j += 2 {
//                 res = append(res, []int{i, j})
//             }
//         }
//         if i >= 3 {
//             res = append(res, []int{i - 1, 1})
//         }
//         i -= 2
//         size -= 8
//     }
//     res = append(res, []int{1, 1})
//     return res
// }

func colorRed(n int) [][]int {
    res := [][]int{{1, 1}}
    for i, k := n, 0; i > 1; i, k = i - 1, (k + 1) % 4 {
        if k == 0 {
            for j := 1; j < i << 1; j += 2 {
                res = append(res, []int{i, j})
            }
        } else if k == 1 {
            res = append(res, []int{i, 2})
        } else if k == 2 {
            for j := 3; j < i << 1; j += 2 {
                res = append(res, []int{i, j})
            }
        } else {
            res = append(res, []int{i, 1})
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/01/example1.jpg" />
    // Input: n = 3
    // Output: [[1,1],[2,1],[2,3],[3,1],[3,5]]
    // Explanation: Initially, we choose the shown 5 triangles to be red. Then, we run the algorithm:
    // - Choose (2,2) that has three red neighbors and color it red.
    // - Choose (3,2) that has two red neighbors and color it red.
    // - Choose (3,4) that has three red neighbors and color it red.
    // - Choose (3,3) that has three red neighbors and color it red.
    // It can be shown that choosing any 4 triangles and running the algorithm will not make all triangles red.
    fmt.Println(colorRed(3)) // [[1,1],[2,1],[2,3],[3,1],[3,5]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/01/example2.jpg" />
    // Input: n = 2
    // Output: [[1,1],[2,1],[2,3]]
    // Explanation: Initially, we choose the shown 3 triangles to be red. Then, we run the algorithm:
    // - Choose (2,2) that has three red neighbors and color it red.
    // It can be shown that choosing any 2 triangles and running the algorithm will not make all triangles red.
    fmt.Println(colorRed(2)) // [[1,1],[2,1],[2,3]]

    fmt.Println(colorRed(1)) // [[1,1]]
    fmt.Println(colorRed(8)) // [[8 1] [8 3] [8 5] [8 7] [8 9] [8 11] [8 13] [8 15] [7 2] [6 3] [6 5] [6 7] [6 9] [6 11] [5 1] [1 1]]
    //fmt.Println(colorRed(64)) // [[1,1],[2,1],[2,3]]
    //fmt.Println(colorRed(1000)) // [[1,1],[2,1],[2,3]]

    fmt.Println(colorRed(10)) // [[1,1],[2,1],[2,3],[3,1],[4,3],[4,5],[4,7],[5,2],[6,1],[6,3],[6,5],[6,7],[6,9],[6,11],[7,1],[8,3],[8,5],[8,7],[8,9],[8,11],[8,13],[8,15],[9,2],[10,1],[10,3],[10,5],[10,7],[10,9],[10,11],[10,13],[10,15],[10,17],[10,19]]
}