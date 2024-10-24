package main

// 1924. Erect the Fence II
// You are given a 2D integer array trees where trees[i] = [xi, yi] represents the location of the ith tree in the garden.

// You are asked to fence the entire garden using the minimum length of rope possible. 
// The garden is well-fenced only if all the trees are enclosed and the rope used forms a perfect circle. 
// A tree is considered enclosed if it is inside or on the border of the circle.

// More formally, you must form a circle using the rope with a center (x, y) 
// and radius r where all trees lie inside or on the circle and r is minimum.

// Return the center and radius of the circle as a length 3 array [x, y, r]. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/06/trees1.png" />
// Input: trees = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
// Output: [2.00000,2.00000,2.00000]
// Explanation: The fence will have center = (2, 2) and radius = 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/06/trees2.png" />
// Input: trees = [[1,2],[2,2],[4,2]]
// Output: [2.50000,2.00000,1.50000]
// Explanation: The fence will have center = (2.5, 2) and radius = 1.5

// Constraints:
//     1 <= trees.length <= 3000
//     trees[i].length == 2
//     0 <= xi, yi <= 3000

import "fmt"
import "math"
import "time"
import "math/rand"

// 最小圆覆盖 Welzl 算法
func outerTrees(trees [][]int) []float64 {
    const eps = 1e-8
    type Point struct{ X, Y float64 }
    arr := make([]Point, len(trees))
    for i, tree := range trees {
        arr[i] = Point{ float64(tree[0]), float64(tree[1]) }
    }
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }) // 随机打乱
    o := arr[0] // 圆心
    r2 := 0.0  // 半径的平方
    // 两点距离的平方
    dis2 := func(p, q Point) float64 { return (p.X - q.X) * (p.X - q.X) + (p.Y - q.Y) * (p.Y - q.Y ) }
    // 三角形外心
    circumcenter := func(a, b, c Point) Point {
        a1, b1, a2, b2 := b.X - a.X, b.Y - a.Y, c.X - a.X, c.Y - a.Y
        c1, c2, d := a1 * a1 + b1 * b1, a2 * a2 + b2 * b2, 2 * (a1 * b2 - a2 * b1)
        return Point{a.X + (c1 * b2 - c2 * b1) / d, a.Y + (a1 * c2 - a2 * c1) / d }
    }
    for i, p := range arr {
        if dis2(o, p) < r2 + eps { continue } // p 在最小圆内部或边上
        o, r2 = p, 0
        for j, q := range arr[:i] {
            if dis2(o, q) < r2+eps { continue } // q 在最小圆内部或边上
            o = Point{(p.X + q.X) / 2, (p.Y + q.Y) / 2}
            r2 = dis2(o, p)
            for _, x := range arr[:j] {
                if dis2(o, x) > r2 + eps { // 保证三点不会共线
                    o = circumcenter(p, q, x)
                    r2 = dis2(o, p)
                }
            }
        }
    }
    return []float64{ o.X, o.Y, math.Sqrt(r2) }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/06/trees1.png" />
    // Input: trees = [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
    // Output: [2.00000,2.00000,2.00000]
    // Explanation: The fence will have center = (2, 2) and radius = 2
    fmt.Println(outerTrees([][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}})) // [2.00000,2.00000,2.00000]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/06/trees2.png" />
    // Input: trees = [[1,2],[2,2],[4,2]]
    // Output: [2.50000,2.00000,1.50000]
    // Explanation: The fence will have center = (2.5, 2) and radius = 1.5
    fmt.Println(outerTrees([][]int{{1,2},{2,2},{4,2}})) // [2.50000,2.00000,1.50000]
}