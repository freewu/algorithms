package main

// 1036. Escape a Large Maze
// There is a 1 million by 1 million grid on an XY-plane, and the coordinates of each grid square are (x, y).

// We start at the source = [sx, sy] square and want to reach the target = [tx, ty] square. 
// There is also an array of blocked squares, where each blocked[i] = [xi, yi] represents a blocked square with coordinates (xi, yi).

// Each move, we can walk one square north, east, south, or west if the square is not in the array of blocked squares.
// We are also not allowed to walk outside of the grid.

// Return true if and only if it is possible to reach the target square from the source square through a sequence of valid moves.

// Example 1:
// Input: blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
// Output: false
// Explanation: The target square is inaccessible starting from the source square because we cannot move.
// We cannot move north or east because those squares are blocked.
// We cannot move south or west because we cannot go outside of the grid.

// Example 2:
// Input: blocked = [], source = [0,0], target = [999999,999999]
// Output: true
// Explanation: Because there are no blocked cells, it is possible to reach the target square.

// Constraints:
//     0 <= blocked.length <= 200
//     blocked[i].length == 2
//     0 <= xi, yi < 10^6
//     source.length == target.length == 2
//     0 <= sx, sy, tx, ty < 10^6
//     source != target
//     It is guaranteed that source and target are not blocked.

import "fmt"
import "sort"
import "slices"
// import "math"

// // bfs 
// func isEscapePossible(blocked [][]int, source []int, target []int) bool {
//     moves := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
//     formatSlice := func(a []int) [2]int { return [2]int{a[0], a[1]} }
//     blockedSet := make(map[[2]int]bool)
//     for i := range blocked {
//         key := formatSlice(blocked[i])
//         blockedSet[key] = true
//     }
//     var bfs func([2]int, [2]int) bool
//     bfs = func(source [2]int, target [2]int) bool {
//         queue, visited, level := [][2]int{}, make(map[[2]int]bool), 0
//         queue = append(queue, source)
//         for len(queue) > 0 {
//             for range queue { 
//                 first := queue[0] // pop
//                 queue = queue[1:]
//                 if first[0] == target[0] && first[1] == target[1] {
//                     return true
//                 }
//                 for i := range moves {
//                     xs, ys := first[0] + moves[i][0], first[1] + moves[i][1]
//                     point := [2]int{ xs, ys }
//                     _, ok1 := visited[point] // 是否访问过
//                     _, ok2 := blockedSet[point] // 是否禁止
//                     if xs >= 0 && xs <= int(math.Pow10(6)) && ys >= 0 && ys <= int(math.Pow10(6)) && !ok1 && !ok2 {
//                         visited[point] = true
//                         queue = append(queue, point)
//                     }
//                 }
//             }
//             level += 1
//             if level >= len(blockedSet) {
//                 return true
//             }
//         }
//         return false
//     }
//     sk, tk := formatSlice(source), formatSlice(target)
//     return bfs(sk, tk) && bfs(tk, sk)
// }

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
    type Point struct { X, Y int }
    visited1, visited2 := make(map[Point]bool), make(map[Point]bool) // true 则该位置已经访问 | false 则该位置未被访问
    for _, v := range blocked {
        visited1[Point{ X: v[0], Y: v[1], }] = true
        visited2[Point{ X: v[0], Y: v[1], }] = true
    }
    newSource, newTarget := Point{ X: source[0], Y: source[1], }, Point{ X: target[0], Y: target[1], }
    maxCount, t1, t2 := len(blocked) * (len(blocked) - 1) / 2, 0, 0
    var bfs func (x, y, maxCount int, count *int, visited map[Point]bool, target Point) bool
    bfs = func (x, y, maxCount int, count *int, visited map[Point]bool, target Point) bool {
        if x < 0 || y < 0 || x >= 1000000 || y >= 1000000 || visited[Point{  X: x, Y: y }] == true {
            return false
        } else {
            *count++
            visited1[Point{ X: x, Y: y, }] = true
            if *count > maxCount || (x == target.X && y == target.Y) ||
                bfs(x+1, y, maxCount, count, visited, target) || // 向右
                bfs(x-1, y, maxCount, count, visited, target) || // 向左
                bfs(x, y+1, maxCount, count, visited, target) || // 向下
                bfs(x, y-1, maxCount, count, visited, target) {  // 向上
                return true
            }
        }
        return false
    }
    return bfs(source[0], source[1], maxCount, &t1, visited1, newTarget) && bfs(target[0], target[1], maxCount, &t2, visited2, newSource)
}

func isEscapePossible1(blocked [][]int, source []int, target []int) bool {
    sort.Slice(blocked, func(i, j int) bool {
        return blocked[i][0] < blocked[j][0]
    })
    mr := int(1e6) - 1
    var rows []int
    for i, m := 0, len(blocked); i < m; {
        c := blocked[i][0]
        rows = append(rows, c)
        for ; i < m && blocked[i][0] == c; i++ {
        }
    }
    rows = append(rows, source[0], target[0])
    sort.Ints(rows)
    rows = slices.Compact(rows)
    sort.Slice(blocked, func(i, j int) bool {
        return blocked[i][1] < blocked[j][1]
    })
    var cols []int
    for i, m := 0, len(blocked); i < m; {
        c := blocked[i][1]
        cols = append(cols, c)
        for ; i < m && blocked[i][1] == c; i++ {
        }
    }
    cols = append(cols, source[1], target[1])
    sort.Ints(cols)
    cols = slices.Compact(cols)
    m, rowFlat := discretize(rows, 0, mr)
    n, colFlat := discretize(cols, 0, mr)
    blockMap := twoDimensionArrayShape(m, n, false)
    for _, b := range blocked {
        blockMap[rowFlat[b[0]]][colFlat[b[1]]] = true
    }
    //fmt.Println(rows, cols, m, n, rowFlat, colFlat)
    sx, sy := rowFlat[source[0]], colFlat[source[1]]
    tx, ty := rowFlat[target[0]], colFlat[target[1]]
    visited := twoDimensionArrayShape(m, n, false)
    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if i == tx && j == ty {
            return true
        }
        visited[i][j] = true
        for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            x, y := i+d[0], j+d[1]
            if x >= m || x < 0 || y >= n || y < 0 {
                continue
            }
            if visited[x][y] || blockMap[x][y] {
                continue
            }
            if dfs(x, y) {
                return true
            }
        }
        return false
    }
    return dfs(sx, sy)
}

func discretize(blocked []int, minVal, maxVal int) (int, map[int]int) {
    preRow := minVal - 1
    j := 0
    hash := map[int]int{}
    for _, b := range blocked {
        if b > preRow+1 {
            j++
        }
        hash[b] = j
        j++
        preRow = b
    }
    if maxVal > preRow {
        j++
    }
    return j, hash
}

func arrayShape[T int | uint | int64 | int32 | byte | bool | *int](n int, defaultVal T) (ans []T) {
    ans = make([]T, n)
    for i := range ans {
        ans[i] = defaultVal
    }
    return
}

func twoDimensionArrayShape[T int | uint | int64 | int32 | byte | bool | *int](m, n int, defaultVal T) (ans [][]T) {
    ans = make([][]T, m)
    for i := range ans {
        ans[i] = make([]T, n)
        for j := range ans[i] {
            ans[i][j] = defaultVal
        }
    }
    return
}

func main() {
    // Example 1:
    // Input: blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
    // Output: false
    // Explanation: The target square is inaccessible starting from the source square because we cannot move.
    // We cannot move north or east because those squares are blocked.
    // We cannot move south or west because we cannot go outside of the grid.
    fmt.Println(isEscapePossible([][]int{{0,1},{1,0}}, []int{0,0}, []int{0,2})) // false
    // Example 2:
    // Input: blocked = [], source = [0,0], target = [999999,999999]
    // Output: true
    // Explanation: Because there are no blocked cells, it is possible to reach the target square.
    fmt.Println(isEscapePossible([][]int{}, []int{0,0}, []int{999999,999999})) // true

    fmt.Println(isEscapePossible([][]int{{0,999991},{0,999993},{0,999996},{1,999996},{1,999997},{1,999998},{1,999999}}, []int{0,999997}, []int{0,2})) // false

    fmt.Println(isEscapePossible1([][]int{{0,1},{1,0}}, []int{0,0}, []int{0,2})) // false
    fmt.Println(isEscapePossible1([][]int{}, []int{0,0}, []int{999999,999999})) // true
    fmt.Println(isEscapePossible1([][]int{{0,999991},{0,999993},{0,999996},{1,999996},{1,999997},{1,999998},{1,999999}}, []int{0,999997}, []int{0,2})) // false
}