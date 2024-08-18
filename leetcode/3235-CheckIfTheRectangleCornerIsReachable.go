package main

// 3235. Check if the Rectangle Corner Is Reachable
// You are given two positive integers xCorner and yCorner, 
// and a 2D array circles, where circles[i] = [xi, yi, ri] denotes a circle with center at (xi, yi) and radius ri.

// There is a rectangle in the coordinate plane with its bottom left corner at the origin 
// and top right corner at the coordinate (xCorner, yCorner). 
// You need to check whether there is a path from the bottom left corner to the top right corner such 
// that the entire path lies inside the rectangle, does not touch or lie inside any circle, 
// and touches the rectangle only at the two corners.

// Return true if such a path exists, and false otherwise.

// Example 1:
// Input: xCorner = 3, yCorner = 4, circles = [[2,1,1]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/18/example2circle1.png" />
// The black curve shows a possible path between (0, 0) and (3, 4).

// Example 2:
// Input: xCorner = 3, yCorner = 3, circles = [[1,1,2]]
// Output: false
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/18/example1circle.png" />
// No path exists from (0, 0) to (3, 3).

// Example 3:
// Input: xCorner = 3, yCorner = 3, circles = [[2,1,1],[1,2,1]]
// Output: false
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/18/example0circle.png" />
// No path exists from (0, 0) to (3, 3).

// Example 4:
// Input: xCorner = 4, yCorner = 4, circles = [[5,5,1]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/04/rectangles.png" />

// Constraints:
//     3 <= xCorner, yCorner <= 10^9
//     1 <= circles.length <= 1000
//     circles[i].length == 3
//     1 <= xi, yi, ri <= 10^9

import "fmt"

func canReachCorner(X, Y int, circles [][]int) bool {
    vis := make([]bool, len(circles))
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // inCircle returns whether Point (x, y) is located inside Circle (ox, oy, r) (including the border)
    inCircle := func (ox, oy, r, x, y int) bool { return (ox-x)*(ox-x)+(oy-y)*(oy-y) <= r*r }
    var dfs func(i int) bool
    dfs = func(i int) bool {
        x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
        if y1 <= Y && abs(x1-X) <= r1 || x1 <= X && y1 <= r1 || x1 > X && inCircle(x1, y1, r1, X, 0) { // Circle i intersects with the bottom or right border of the rectangle
            return true
        }
        vis[i] = true
        for j, circle := range circles {
            x2, y2, r2 := circle[0], circle[1], circle[2]
            // Let Point A which |O1A| / |O1O2| = r1 / (r1+r2). If two circles are connected to each other then A must be inside the intersection
            // And its coordinate is: (x1·r2+x2·r1)/(r1+r2), (y1·r2+y2·r1)/(r1+r2)
            if !vis[j] && (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= (r1+r2)*(r1+r2) && x1*r2+x2*r1 < (r1+r2)*X && y1*r2+y2*r1 < (r1+r2)*Y && dfs(j) {
                return true
            }
        }
        return false
    }
    for i, circle := range circles {
        x, y, r := circle[0], circle[1], circle[2]
        if inCircle(x, y, r, 0, 0) || inCircle(x, y, r, X, Y) || !vis[i] && (x <= X && abs(y-Y) <= r || y <= Y && x <= r || y > Y && inCircle(x, y, r, 0, Y)) && dfs(i) { // DFS starts from circles which intersects the left/top border of the rectangle
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: xCorner = 3, yCorner = 4, circles = [[2,1,1]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/18/example2circle1.png" />
    // The black curve shows a possible path between (0, 0) and (3, 4).
    fmt.Println(canReachCorner(3,4,[][]int{{2,1,1}})) // true
    // Example 2:
    // Input: xCorner = 3, yCorner = 3, circles = [[1,1,2]]
    // Output: false
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/18/example1circle.png" />
    // No path exists from (0, 0) to (3, 3).
    fmt.Println(canReachCorner(3,3,[][]int{{1,1,2}})) // false
    // Example 3:
    // Input: xCorner = 3, yCorner = 3, circles = [[2,1,1],[1,2,1]]
    // Output: false
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/18/example0circle.png" />
    // No path exists from (0, 0) to (3, 3).
    fmt.Println(canReachCorner(3,3,[][]int{{2,1,1},{1,2,1}})) // false
    // Example 4:
    // Input: xCorner = 4, yCorner = 4, circles = [[5,5,1]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/04/rectangles.png" />
    fmt.Println(canReachCorner(4,4,[][]int{{5,5,1}})) // true
}