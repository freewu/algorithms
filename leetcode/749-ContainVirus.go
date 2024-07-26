package main

// 749. Contain Virus
// A virus is spreading rapidly, and your task is to quarantine the infected area by installing walls.

// The world is modeled as an m x n binary grid isInfected, where isInfected[i][j] == 0 represents uninfected cells, 
// and isInfected[i][j] == 1 represents cells contaminated with the virus. 
// A wall (and only one wall) can be installed between any two 4-directionally adjacent cells, on the shared boundary.

// Every night, the virus spreads to all neighboring cells in all four directions unless blocked by a wall. 
// Resources are limited. Each day, you can install walls around only one region (i.e., the affected area (continuous block of infected cells) 
// that threatens the most uninfected cells the following night). There will never be a tie.

// Return the number of walls used to quarantine all the infected regions. 
// If the world will become fully infected, return the number of walls used.


// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/01/virus11-grid.jpg" />
// Input: isInfected = [[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]
// Output: 10
// Explanation: There are 2 contaminated regions.
// On the first day, add 5 walls to quarantine the viral region on the left. The board after the virus spreads is:
// <img src="https://assets.leetcode.com/uploads/2021/06/01/virus12edited-grid.jpg" />
// On the second day, add 5 walls to quarantine the viral region on the right. The virus is fully contained.
// <img src="https://assets.leetcode.com/uploads/2021/06/01/virus13edited-grid.jpg" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/01/virus2-grid.jpg" />
// Input: isInfected = [[1,1,1],[1,0,1],[1,1,1]]
// Output: 4
// Explanation: Even though there is only one cell saved, there are 4 walls built.
// Notice that walls are only built on the shared boundary of two different cells.

// Example 3:
// Input: isInfected = [[1,1,1,0,0,0,0,0,0],[1,0,1,0,1,1,1,1,1],[1,1,1,0,0,0,0,0,0]]
// Output: 13
// Explanation: The region on the left only builds two new walls.

// Constraints:
//     m == isInfected.length
//     n == isInfected[i].length
//     1 <= m, n <= 50
//     isInfected[i][j] is either 0 or 1.
//     There is always a contiguous viral region throughout the described process that will infect strictly more uncontaminated squares in the next round.

import "fmt"
import "sort"

const (
    Infected   = 1
    Uninfected = 0
    Blocked    = -1
)

type pii struct {
    X, Y int
}

// cluster 病毒集群
type cluster struct {
    // walls 需要建的墙数
    walls int
    // infectedCells 该集群中感染区域的坐标
    infectedCells map[pii]bool
    // uninfectedNeighbor 该集群附近未感染区域的坐标，用于病毒扩散
    uninfectedNeighbor map[pii]bool
}

func (c *cluster) dfs(isInfected [][]int, visited [][]bool, x, y int) {
    if x < 0 || x >= len(isInfected) || y < 0 || y >= len(isInfected[0]) || visited[x][y] || isInfected[x][y] == Blocked {
        // 超出范围的、已访问的、已经被封锁的
        return
    }
    if isInfected[x][y] == Uninfected {
        // 有多少个未感染的邻居，就有多少个墙
        c.walls++
        c.uninfectedNeighbor[pii{x, y}] = true
        return
    }
    c.infectedCells[pii{x, y}] = true
    visited[x][y] = true
    c.dfs(isInfected, visited, x+1, y)
    c.dfs(isInfected, visited, x-1, y)
    c.dfs(isInfected, visited, x, y+1)
    c.dfs(isInfected, visited, x, y-1)
}

func containVirus(isInfected [][]int) (res int) {
    // 模拟每一轮变化
    for {
        visited := make([][]bool, len(isInfected))
        for i := range visited {
            visited[i] = make([]bool, len(isInfected[0]))
        }
        var clusters []*cluster
        // dfs 找到所有未隔离的病毒集群
        for i := range isInfected {
            for j := range isInfected[i] {
                if !visited[i][j] && isInfected[i][j] == Infected {
                    c := &cluster{
                        infectedCells:      map[pii]bool{},
                        uninfectedNeighbor: map[pii]bool{},
                    }
                    c.dfs(isInfected, visited, i, j)
                    clusters = append(clusters, c)
                }
            }
        }
        // 已经没有未隔离的病毒集群了
        if len(clusters) == 0 {
            break
        }
        // 按照题意，找到威胁最大的
        sort.Slice(clusters, func(i, j int) bool {
            return len(clusters[i].uninfectedNeighbor) > len(clusters[j].uninfectedNeighbor)
        })
        // 给威胁最大的加上防火墙隔离
        for pii := range clusters[0].infectedCells {
            isInfected[pii.X][pii.Y] = Blocked
        }
        // 记录增加的防火墙个数
        res += clusters[0].walls
        // 病毒扩散
        for i := 1; i < len(clusters); i++ {
            for pii := range clusters[i].uninfectedNeighbor {
                isInfected[pii.X][pii.Y] = Infected
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/01/virus11-grid.jpg" />
    // Input: isInfected = [[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]
    // Output: 10
    // Explanation: There are 2 contaminated regions.
    // On the first day, add 5 walls to quarantine the viral region on the left. The board after the virus spreads is:
    // <img src="https://assets.leetcode.com/uploads/2021/06/01/virus12edited-grid.jpg" />
    // On the second day, add 5 walls to quarantine the viral region on the right. The virus is fully contained.
    // <img src="https://assets.leetcode.com/uploads/2021/06/01/virus13edited-grid.jpg" />
    isInfected1 := [][]int{
        {0,1,0,0,0,0,0,1},
        {0,1,0,0,0,0,0,1},
        {0,0,0,0,0,0,0,1},
        {0,0,0,0,0,0,0,0},
    }
    fmt.Println(containVirus(isInfected1)) // 10
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/01/virus2-grid.jpg" />
    // Input: isInfected = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: 4
    // Explanation: Even though there is only one cell saved, there are 4 walls built.
    // Notice that walls are only built on the shared boundary of two different cells.
    isInfected2 := [][]int{
        {1,1,1},
        {1,0,1},
        {1,1,1},
    }
    fmt.Println(containVirus(isInfected2)) // 4
    // Example 3:
    // Input: isInfected = [[1,1,1,0,0,0,0,0,0],[1,0,1,0,1,1,1,1,1],[1,1,1,0,0,0,0,0,0]]
    // Output: 13
    // Explanation: The region on the left only builds two new walls.
    isInfected3 := [][]int{
        {1,1,1,0,0,0,0,0,0},
        {1,0,1,0,1,1,1,1,1},
        {1,1,1,0,0,0,0,0,0},
    }
    fmt.Println(containVirus(isInfected3)) // 13
}