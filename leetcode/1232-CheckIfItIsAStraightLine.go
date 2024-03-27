package main

// 1232. Check If It Is a Straight Line
// You are given an array coordinates, coordinates[i] = [x, y], 
// where [x, y] represents the coordinate of a point. 
// Check if these points make a straight line in the XY plane.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/15/untitled-diagram-2.jpg" />
// Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/09/untitled-diagram-1.jpg" />
// Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
// Output: false
 
// Constraints:
//     2 <= coordinates.length <= 1000
//     coordinates[i].length == 2
//     -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
//     coordinates contains no duplicate point.

import "fmt"

// func checkStraightLine(coordinates [][]int) bool {
//     x, y := coordinates[1][0] - coordinates[0][0], coordinates[1][1] - coordinates[0][1]
//     for i := 2 ; i < len(coordinates); i++ {
//         // 差值不一样就不是在同一条直线上
//         if coordinates[i][0] - coordinates[i-1][0] != x || coordinates[i][1] - coordinates[i-1][1] != y {
//             return false
//         }
//     }
//     return true
// }

func checkStraightLine(coordinates [][]int) bool {
    // 平移 每个坐标都减去 coordinates[0] 
    // [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
    //      => [[1-1,2-2],[2-1,3-2], ..., [6-1,7-5]]
    //      => [[0 0] [1 1] [2 2] [3 3] [4 4] [5 5]]
    x, y := coordinates[0][0], coordinates[0][1]
    for _, v := range coordinates {
        v[0] -= x
        v[1] -= y
    }
    a, b := coordinates[1][1], -coordinates[1][0]
    for _, v := range coordinates {
        if a * v[0] + b * v[1] != 0 {
            return false
        }
    }
    return true
}

func main() {
    // [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
    fmt.Println(
        checkStraightLine(
            [][]int{
                []int{1,2},
                []int{2,3},
                []int{3,4},
                []int{4,5},
                []int{5,6},
                []int{6,7},
            },
        ),
    ) // true

    //  [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
    fmt.Println(
        checkStraightLine(
            [][]int{
                []int{1,1},
                []int{2,2},
                []int{3,4},
                []int{4,5},
                []int{5,6},
                []int{7,7},
            },
        ),
    ) // false
}