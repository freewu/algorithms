package main

// LCP 45. 自行车炫技赛场
// 「力扣挑战赛」中N*M大小的自行车炫技赛场的场地由一片连绵起伏的上下坡组成，
// 场地的高度值记录于二维数组terrain中，场地的减速值记录于二维数组obstacle中。

//     1. 若选手骑着自行车从高度为h1 且减速值为o1 的位置到高度为h2 且减速值为o2 的相邻位置（上下左右四个方向），
//        速度变化值为h1 - h2 - o2（负值减速，正值增速）。

// 选手初始位于坐标position处且初始速度为 1，请问选手可以刚好到其他哪些位置时速度依旧为 1。
// 请以二维数组形式返回这些位置。
// 若有多个位置则按行坐标升序排列，若有多个位置行坐标相同则按列坐标升序排列。

// 注意：骑行过程中速度不能为零或负值

// 示例 1：
// 输入：position = [0,0], terrain = [[0,0],[0,0]], obstacle = [[0,0],[0,0]]
// 输出：[[0,1],[1,0],[1,1]]
// 解释： 由于当前场地属于平地，根据上面的规则，选手从[0,0]的位置出发都能刚好在其他处的位置速度为 1。

// 示例 2：
// 输入：position = [1,1], terrain = [[5,0],[0,6]], obstacle = [[0,6],[7,0]]
// 输出：[[0,1]]
// 解释： 选手从[1,1]处的位置出发，到[0,1]处的位置时恰好速度为 1。

// 提示：
//     n == terrain.length == obstacle.length
//     m == terrain[i].length == obstacle[i].length
//     1 <= n <= 100
//     1 <= m <= 100
//     0 <= terrain[i][j], obstacle[i][j] <= 100
//     position.length == 2
//     0 <= position[0] < n
//     0 <= position[1] < m

import "fmt"
import "sort"

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    n, m := len(terrain), len(terrain[0])
    res, visited := [][]int{}, make(map[[3]int]bool) // 使用三维数组记录访问状态：坐标(x, y) + 速度
    var dfs func(x, y, speed int)
    dfs = func(x, y, speed int) {
        if x < 0 || y < 0 || x >= n || y >= m || speed <= 0 { return } // 越界或速度为0时返回
        if visited[[3]int{x, y, speed}] { return } // 已访问过此状态时返回
        visited[[3]int{x, y, speed}] = true // 标记为已访问
        if speed == 1 && !(x == position[0] && y == position[1]) { // 速度为1且不是起始位置时加入结果
            res = append(res, []int{ x, y })
        }
        for _, d := range directions { // 遍历四个方向
            nx, ny := x + d[0], y + d[1]
            if nx >= 0 && ny >= 0 && nx < n && ny < m {
                dfs(nx, ny, (speed + terrain[x][y] - terrain[nx][ny] - obstacle[nx][ny])) // new speed
            }
        }
    }
    dfs(position[0], position[1], 1) // 从起始位置开始深度优先搜索
    sort.Slice(res, func(i, j int) bool { // 若有多个位置则按行坐标升序排列，
        if res[i][0] == res[j][0] { return res[i][1] < res[j][1] } // 若有多个位置行坐标相同则按列坐标升序排列
        return res[i][0] < res[j][0]
    })
    return res
}

func main() {
    // 示例 1：
    // 输入：position = [0,0], terrain = [[0,0],[0,0]], obstacle = [[0,0],[0,0]]
    // 输出：[[0,1],[1,0],[1,1]]
    // 解释： 由于当前场地属于平地，根据上面的规则，选手从[0,0]的位置出发都能刚好在其他处的位置速度为 1。
    fmt.Println(bicycleYard([]int{0,0}, [][]int{{0,0},{0,0}}, [][]int{{0,0},{0,0}})) // [[0,1],[1,0],[1,1]]
    // 示例 2：
    // 输入：position = [1,1], terrain = [[5,0],[0,6]], obstacle = [[0,6],[7,0]]
    // 输出：[[0,1]]
    // 解释： 选手从[1,1]处的位置出发，到[0,1]处的位置时恰好速度为 1。
    fmt.Println(bicycleYard([]int{1,1}, [][]int{{5,0},{0,6}}, [][]int{{0,6},{7,0}})) // [[0,1]]
}