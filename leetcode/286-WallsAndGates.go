package main

// 286. Walls and Gates
// You are given an m x n grid rooms initialized with these three possible values.
//     -1 A wall or an obstacle.
//     0 A gate.
//     INF Infinity means an empty room. We use the value 2^31 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.

// Fill each empty room with the distance to its nearest gate. 
// If it is impossible to reach a gate, it should be filled with INF.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/03/grid.jpg" />
// Input: rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
// Output: [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

// Example 2:
// Input: rooms = [[-1]]
// Output: [[-1]]

// Constraints:
//     m == rooms.length
//     n == rooms[i].length
//     1 <= m, n <= 250
//     rooms[i][j] is -1, 0, or 2^31 - 1.

import "fmt"

// dfs
func wallsAndGates(rooms [][]int) {
    var dfs func(rooms [][]int, i, j, distance int)
    dfs = func(rooms [][]int, i, j, distance int) {
        if i == -1 || j == -1 || i == len(rooms) || j == len(rooms[0]) {
            return
        }
        val := rooms[i][j]
        // 用 val <= distance 来剪枝，避免重复走本次遍历已经遍历过的节点，如果另一个门到这里的距离更短，那么也没必要继续遍历了
        if val == 0 || val == -1 || val <= distance {
            return
        }
        rooms[i][j] = distance
        distance++
        dfs(rooms, i, j+1, distance)
        dfs(rooms, i, j-1, distance)
        dfs(rooms, i+1, j, distance)
        dfs(rooms, i-1, j, distance)
    }
    for i := 0; i < len(rooms); i++ {
        for j := 0; j < len(rooms[0]); j++ {
            if rooms[i][j] == 0 {
                dfs(rooms, i, j+1, 1)
                dfs(rooms, i, j-1, 1)
                dfs(rooms, i+1, j, 1)
                dfs(rooms, i-1, j, 1)
            }
        }
    }
}

// bfs
func wallsAndGates1(rooms [][]int) {
    q := [][]int{}
    for i := 0; i < len(rooms); i++ {
        for j := 0; j < len(rooms[0]); j++ {
            if rooms[i][j] == 0 {
                q = append(q, []int{i, j})
            }
        }
    }
    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for i := 0; i < len(q); i++ {
        p := q[i]
        for _, direct := range directions {
            j, k := p[0]+direct[0], p[1]+direct[1]
            if j < 0 || j == len(rooms) || k < 0 || k == len(rooms[0]) {
                continue
            }

            val := rooms[j][k]
            if val == -1 || val == 0 || val <= rooms[p[0]][p[1]]+1 {
                continue
            }
            rooms[j][k] = rooms[p[0]][p[1]] + 1
            q = append(q, []int{j, k})
        }
    }
}

func wallsAndGates2(rooms [][]int) {
    mx, m, n := 2147483647, len(rooms), len(rooms[0])
    queue := [][2]int{} // 队列
    var bfs func() // 返回当前到门的距离
    bfs = func() {
        for len(queue) > 0 {
            cur := queue[0]
            a := cur[0]
            b := cur[1]
            if rooms[a][b] < 0 {
                queue = queue[1:]
                continue
            }
            if a > 0 && mx == rooms[a-1][b] { // 向上
                queue = append(queue, [2]int{a - 1, b})
                rooms[a-1][b] = rooms[a][b] + 1
            }
            if b < n-1 && mx == rooms[a][b+1] { // 向右
                queue = append(queue, [2]int{a, b + 1})
                rooms[a][b+1] = rooms[a][b] + 1
            }
            if a < m-1 && mx == rooms[a+1][b] { // 向下
                queue = append(queue, [2]int{a + 1, b})
                rooms[a+1][b] = rooms[a][b] + 1
            }
            if b > 0 && mx == rooms[a][b-1] { // 向左
                queue = append(queue, [2]int{a, b - 1})
                rooms[a][b-1] = rooms[a][b] + 1
            }
            queue = queue[1:]
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                queue = append(queue, [2]int{i, j})
            }
        }
    }
    bfs()
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/03/grid.jpg" />
    // Input: rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
    // Output: [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]
    room1 := [][]int{
        {2147483647,-1,0,2147483647},
        {2147483647,2147483647,2147483647,-1},
        {2147483647,-1,2147483647,-1},
        {0,-1,2147483647,2147483647},
    }
    fmt.Println("before: ", room1) 
    wallsAndGates(room1)
    fmt.Println("after: ", room1) // [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]
    // Example 2:
    // Input: rooms = [[-1]]
    // Output: [[-1]]
    room2 := [][]int{{-1}}
    fmt.Println("before: ", room2) 
    wallsAndGates(room2)
    fmt.Println("after: ", room2) // [[-1]]

    room11 := [][]int{
        {2147483647,-1,0,2147483647},
        {2147483647,2147483647,2147483647,-1},
        {2147483647,-1,2147483647,-1},
        {0,-1,2147483647,2147483647},
    }
    fmt.Println("before: ", room11) 
    wallsAndGates1(room11)
    fmt.Println("after: ", room11) // [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

    room12 := [][]int{{-1}}
    fmt.Println("before: ", room12) 
    wallsAndGates1(room12)
    fmt.Println("after: ", room12) // [[-1]]

    room21 := [][]int{
        {2147483647,-1,0,2147483647},
        {2147483647,2147483647,2147483647,-1},
        {2147483647,-1,2147483647,-1},
        {0,-1,2147483647,2147483647},
    }
    fmt.Println("before: ", room21) 
    wallsAndGates2(room21)
    fmt.Println("after: ", room21) // [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

    room22 := [][]int{{-1}}
    fmt.Println("before: ", room22) 
    wallsAndGates2(room22)
    fmt.Println("after: ", room22) // [[-1]]
}