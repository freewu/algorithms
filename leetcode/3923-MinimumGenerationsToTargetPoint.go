package main

// 3923. Minimum Generations to Target Point
// You are given a 2D integer array points where points[i] = [xi, yi, zi] represents a point in 3D space, and an integer array target representing a target point.

// Define generation 0 as the initial list of points. For each integer k >= 1, form generation k as follows:
//     1. Consider every pair of two distinct points a = [x1, y1, z1] and b = [x2, y2, z2] taken from all points produced in generations 0 through k - 1.
//     2. For each such pair, compute c = [floor((x1 + x2) / 2), floor((y1 + y2) / 2), floor((z1 + z2) / 2)] and collect every such c into a generation k.
//     3. All points in the generation k are produced simultaneously from points in generations 0 through​​​​​​​ k - 1.
//     4. After generation k is formed, the points in the generation k are considered available for forming later generations.

// Return the smallest integer k such that the target appears in one of the generations 0 through k. 
// If the target is already in the initial points, return 0. 
// If it is impossible to obtain the target, return -1.

// Notes:
//     1. floor denotes rounding down to the nearest integer.
//     2. "Two distinct points" means the two chosen points must have different (x, y, z) coordinates. 
//        A point cannot be paired with itself, and pairing two points with identical coordinates is not possible.

// Example 1:
// Input: points = [[0,0,0],[6,6,6]], target = [3,3,3]
// Output: 1
// Explanation:
// Generation 0: The initial points = [[0, 0, 0], [6, 6, 6]].
// The target = [3, 3, 3] does not exist in generation 0.
// Generation 1: For each pair of points in generation 0, we create new points.
// Using [0, 0, 0] and [6, 6, 6], we generate [3, 3, 3].
// After generation 1, points = [[0, 0, 0], [6, 6, 6], [3, 3, 3]].
// The target = [3, 3, 3] is found in generation 1, so the smallest k is 1.

// Example 2:
// Input: points = [[0,0,0],[5,5,5]], target = [1,1,1]
// Output: 2
// Explanation:
// Generation 0: The initial points = [[0, 0, 0], [5, 5, 5]].
// The target = [1, 1, 1] does not exist in generation 0.
// Generation 1: For each pair of points in generation 0, we create new points.
// Using [0, 0, 0] and [5, 5, 5], we generate [2, 2, 2].
// After generation 1, points = [[0, 0, 0], [5, 5, 5], [2, 2, 2]].
// Generation 2: For each pair of points available after generation 1, we create new points.
// Using [0, 0, 0] and [5, 5, 5], we generate [2, 2, 2].
// Using [0, 0, 0] and [2, 2, 2], we generate [1, 1, 1].
// Using [5, 5, 5] and [2, 2, 2], we generate [3, 3, 3].
// After generation 2, points = [[0, 0, 0], [5, 5, 5], [2, 2, 2], [1, 1, 1], [3, 3, 3]].
// The target = [1, 1, 1] is found in generation 2, so the smallest k is 2.

// Example 3:
// Input: points = [[0,0,0],[2,2,2],[3,3,3]], target = [2,2,2]
// Output: 0
// Explanation:
// Generation 0: The initial points = [[0, 0, 0], [2, 2, 2], [3, 3, 3]].
// The target = [2, 2, 2] already exists in generation 0, so the smallest k is 0.

// Example 4:
// Input: points = [[1,2,3]], target = [5,5,5]
// Output: -1
// Explanation:
// Only one initial point is available, so no new points can be generated.
// Therefore, the target cannot be obtained, and the answer is -1.

// Constraints:
//     1 <= points.length <= 20
//     points[i] = [xi, yi, zi​​​​​​​]
//     0 <= xi, yi, zi <= 6
//     target.length == 3
//     ​​​​​​​0 <= target[i] <= 6
//     The initial set of points contains no duplicates.

import "fmt"
import "maps"

func minGenerations(points [][]int, target []int) int {
    type Point struct{ x, y, z int }
    tar := Point{target[0], target[1], target[2]}
    curr := make(map[Point]struct{}, len(points))
    for _, p := range points {
        curr[Point{p[0], p[1], p[2]}] = struct{}{}
    }
    for i := 0; ; i++ {
        if _, ok := curr[tar]; ok {
            return i
        }
        next := maps.Clone(curr)
        for p := range curr {
            for q := range curr { // 枚举 curr 中的所有点对 (p, q)
                next[Point{(p.x + q.x) / 2, (p.y + q.y) / 2, (p.z + q.z) / 2}] = struct{}{}
            }
        }
        if len(next) == len(curr) { // 没有产生新的点
            return -1
        }
        curr = next
    }
}

func minGenerations1(points [][]int, target []int) int {
    var visited [7][7][7]bool // 三维访问标记数组：坐标范围 [0,6]，标记点是否已生成
    type Point struct{ x, y, z int } // Point 表示三维空间中的一个点 (x,y,z)
    var queue []Point // 初始化 BFS 队列，存储所有已生成的点
    targetPoint := Point{target[0], target[1], target[2]}
    // 遍历初始点，去重并加入队列
    for _, p := range points {
        current := Point{p[0], p[1], p[2]}
        if !visited[current.x][current.y][current.z] {
            visited[current.x][current.y][current.z] = true
            queue = append(queue, current)
        }
    }
    // 目标点已存在，直接返回 0 代
    if visited[targetPoint.x][targetPoint.y][targetPoint.z] {
        return 0
    }
    // BFS 搜索：start 为当前层起始索引，generation 为当前代数
    start, generation := 0, 1
    for {
        // 当前层结束索引
        end := len(queue)
        // 无新节点可遍历，退出循环
        if start == end {
            break
        }
        // 不足两个点，无法生成中点，直接返回失败
        if len(queue) < 2 {
            return -1
        }
        // 存储当前代新生成的点
        var nextPoints []Point
        // 遍历当前层所有点，两两组合生成中点
        for i := start; i < end; i++ {
            for j := 0; j < end; j++ {
                p1, p2 := queue[i], queue[j]
                // 跳过相同点，无法生成新中点
                if p1.x == p2.x && p1.y == p2.y && p1.z == p2.z {
                    continue
                }
                // 计算两点中点（整数除法）
                midX := (p1.x + p2.x) / 2
                midY := (p1.y + p2.y) / 2
                midZ := (p1.z + p2.z) / 2
                midPoint := Point{midX, midY, midZ}
                // 中点未访问过：标记并加入下一层
                if !visited[midX][midY][midZ] {
                    visited[midX][midY][midZ] = true
                    nextPoints = append(nextPoints, midPoint)
                    // 找到目标点，返回当前代数
                    if midPoint == targetPoint {
                        return generation
                    }
                }
            }
        }
        // 无新点生成，退出循环
        if len(nextPoints) == 0 {
            break
        }
        // 将新点加入总队列
        queue = append(queue, nextPoints...)
        // 移动到下一层起始位置
        start = end
        // 代数 +1
        generation++
    }
    // 无法生成目标点
    return -1
}

func main() {
    // Example 1:
    // Input: points = [[0,0,0],[6,6,6]], target = [3,3,3]
    // Output: 1
    // Explanation:
    // Generation 0: The initial points = [[0, 0, 0], [6, 6, 6]].
    // The target = [3, 3, 3] does not exist in generation 0.
    // Generation 1: For each pair of points in generation 0, we create new points.
    // Using [0, 0, 0] and [6, 6, 6], we generate [3, 3, 3].
    // After generation 1, points = [[0, 0, 0], [6, 6, 6], [3, 3, 3]].
    // The target = [3, 3, 3] is found in generation 1, so the smallest k is 1.
    fmt.Println(minGenerations([][]int{{0,0,0},{6,6,6}}, []int{3,3,3})) // 1  
    // Example 2:
    // Input: points = [[0,0,0],[5,5,5]], target = [1,1,1]
    // Output: 2
    // Explanation:
    // Generation 0: The initial points = [[0, 0, 0], [5, 5, 5]].
    // The target = [1, 1, 1] does not exist in generation 0.
    // Generation 1: For each pair of points in generation 0, we create new points.
    // Using [0, 0, 0] and [5, 5, 5], we generate [2, 2, 2].
    // After generation 1, points = [[0, 0, 0], [5, 5, 5], [2, 2, 2]].
    // Generation 2: For each pair of points available after generation 1, we create new points.
    // Using [0, 0, 0] and [5, 5, 5], we generate [2, 2, 2].
    // Using [0, 0, 0] and [2, 2, 2], we generate [1, 1, 1].
    // Using [5, 5, 5] and [2, 2, 2], we generate [3, 3, 3].
    // After generation 2, points = [[0, 0, 0], [5, 5, 5], [2, 2, 2], [1, 1, 1], [3, 3, 3]].
    // The target = [1, 1, 1] is found in generation 2, so the smallest k is 2.
    fmt.Println(minGenerations([][]int{{0,0,0},{5,5,5}}, []int{1,1,1})) // 2
    // Example 3:
    // Input: points = [[0,0,0],[2,2,2],[3,3,3]], target = [2,2,2]
    // Output: 0
    // Explanation:
    // Generation 0: The initial points = [[0, 0, 0], [2, 2, 2], [3, 3, 3]].
    // The target = [2, 2, 2] already exists in generation 0, so the smallest k is 0.
    fmt.Println(minGenerations([][]int{{0,0,0},{2,2,2},{3,3,3}}, []int{2,2,2})) // 0
    // Example 4:
    // Input: points = [[1,2,3]], target = [5,5,5]
    // Output: -1
    // Explanation:
    // Only one initial point is available, so no new points can be generated.
    // Therefore, the target cannot be obtained, and the answer is -1.
    fmt.Println(minGenerations([][]int{{1,2,3}}, []int{5,5,5})) // -1

    fmt.Println(minGenerations1([][]int{{0,0,0},{6,6,6}}, []int{3,3,3})) // 1  
    fmt.Println(minGenerations1([][]int{{0,0,0},{5,5,5}}, []int{1,1,1})) // 2
    fmt.Println(minGenerations1([][]int{{0,0,0},{2,2,2},{3,3,3}}, []int{2,2,2})) // 0
    fmt.Println(minGenerations1([][]int{{1,2,3}}, []int{5,5,5})) // -1
}