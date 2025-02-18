package main

// LCP 75. 传送卷轴
// 随着不断的深入，小扣来到了守护者之森寻找的魔法水晶。首先，他必须先通过守护者的考验。

// 考验的区域是一个正方形的迷宫，maze[i][j]表示在迷宫i行j列的地形：
//     1. 若为.，表示可以到达的空地；
//     2. 若为#，表示不可到达的墙壁；
//     3. 若为S，表示小扣的初始位置；
//     4. 若为T，表示魔法水晶的位置。

// 小扣每次可以向 上、下、左、右 相邻的位置移动一格。而守护者拥有一份「传送魔法卷轴」，使用规则如下：
//     1. 魔法需要在小扣位于空地时才能释放，发动后卷轴消失；；
//     2. 发动后，小扣会被传送到水平或者竖直的镜像位置，且目标位置不得为墙壁(如下图所示)；
//        <img src="https://pic.leetcode.cn/1681789509-wTekFu-image.png" />

// 在使用卷轴后，小扣将被「附加负面效果」，因此小扣需要尽可能缩短传送后到达魔法水晶的距离。
// 而守护者的目标是阻止小扣到达魔法水晶的位置；如果无法阻止，则尽可能增加小扣传送后到达魔法水晶的距离。 
// 假设小扣和守护者都按最优策略行事，返回小扣需要在 「附加负面效果」的情况下最少移动多少次才能到达魔法水晶。
// 如果无法到达，返回-1。

// 注意：
//     守护者可以不使用卷轴；
//     传送后的镜像位置可能与原位置相同。

// 示例 1：
// 输入：maze = [".....","##S..","...#.","T.#..","###.."]
// 输出：7
// 解释：如下图所示：守护者释放魔法的两个最佳的位置为 [2,0] 或 [3,1]： 
//     若小扣经过 [2,0]，守护者在该位置释放魔法， 小扣被传送至 [2,4] 处且加上负面效果，此时小扣还需要移动 7 次才能到达魔法水晶； 
//     若小扣经过 [3,1]，守护者在该位置释放魔法， 小扣被传送至 [3,3] 处且加上负面效果，此时小扣还需要移动 9 次才能到达魔法水晶； 
//     因此小扣负面效果下最少需要移动 7 次才能到达魔法水晶。
//     <img src="https://pic.leetcode.cn/1681714676-gksEMT-image.png" />

// 示例 2：
// 输入：maze = [".#..","..##",".#S.",".#.T"]
// 输出：-1
// 解释：如下图所示。 
//     若小扣向下移动至 [3,2]，守护者使其传送至 [0,2]，小扣将无法到达魔法水晶； 
//     若小扣向右移动至 [2,3]，守护者使其传送至 [2,0]，小扣将无法到达魔法水晶；
//     <img src="https://pic.leetcode.cn/1681714693-LsxKAh-image.png" />

// 示例 3：
// 输入：maze = ["S###.","..###","#..##","##..#","###.T"]
// 输出：5
// 解释：如下图所示： 守护者需要小扣在空地才能释放，因此初始无法将其从 [0,0] 传送至 [0,4]; 
//       当小扣移动至 [2,1] 时，释放卷轴将其传送至水平方向的镜像位置 [2,1]（为原位置） 而后小扣需要移动 5 次才能到达魔法水晶。
//       <img src="https://pic.leetcode.cn/1681800985-KrSdru-image.png" />

// 提示：
//     4 <= maze.length == maze[i].length <= 200
//     maze[i][j]仅包含"."、"#"、"S"、"T"

import "fmt"
import "sort"

func challengeOfTheKeeper(maze []string) int {
    type Point struct{ x, y int }
    var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m, n := len(maze), len(maze[0])
    // 1. 找到起点终点坐标
    var sx, sy, tx, ty int
    for i, row := range maze {
        for j, c := range row {
            if c == 'S' {
                sx, sy = i, j
            } else if c == 'T' {
                tx, ty = i, j
            }
        }
    }
    // 2. BFS 计算终点到其余点的最短距离
    disFromT := make([][]int, m)
    for i := range disFromT {
        disFromT[i] = make([]int, n)
        for j := range disFromT[i] {
            disFromT[i][j] = 1 << 31
        }
    }
    disFromT[tx][ty] = 0
    q := []Point{{tx, ty}}
    for step := 1; len(q) > 0; step++ {
        tmp := q
        q = nil
        for _, p := range tmp {
            for _, d := range directions {
                x, y := p.x + d.x, p.y + d.y
                if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] != '#' && disFromT[x][y] == 1 << 31 {
                    disFromT[x][y] = step
                    q = append(q, Point{x, y})
                }
            }
        }
    }
    // 3. 剪枝：如果 S 无法到达 T，直接返回 -1
    if disFromT[sx][sy] == 1 << 31 {
        return -1
    }
    // 4. 二分
    visited := make([][]int, m)
    for i := range visited {
        visited[i] = make([]int, n)
    }
    res := sort.Search(m*n+1, func(maxDis int) bool {
        // DFS，看能否在「附加负面效果」的情况下，移动不超过 maxDis 步到达终点
        var dfs func(int, int) bool
        dfs = func(i, j int) bool {
            if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == maxDis + 1 || maze[i][j] == '#' {
                return false
            }
            if maze[i][j] == 'T' { // 到达终点
                return true
            }
            visited[i][j] = maxDis + 1 // 为避免反复创建 visited，用一个每次二分都不一样的数来标记
            if maze[i][j] == '.' {
                // 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
                if x, y := i, n-1-j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
                    return false
                }
                if x, y := m-1-i, j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
                    return false
                }
            }
            // 枚举四个方向
            for _, d := range directions {
                if dfs(i + d.x, j + d.y) { // 到达终点
                    return true
                }
            }
            return false // 无法到达终点
        }
        return dfs(sx, sy)
    })
    if res > m * n { // 守护者使用卷轴传送小扣，可以把小扣传送到一个无法到达终点的位置
        return -1
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：maze = [".....","##S..","...#.","T.#..","###.."]
    // 输出：7
    // 解释：如下图所示：守护者释放魔法的两个最佳的位置为 [2,0] 或 [3,1]： 
    //     若小扣经过 [2,0]，守护者在该位置释放魔法， 小扣被传送至 [2,4] 处且加上负面效果，此时小扣还需要移动 7 次才能到达魔法水晶； 
    //     若小扣经过 [3,1]，守护者在该位置释放魔法， 小扣被传送至 [3,3] 处且加上负面效果，此时小扣还需要移动 9 次才能到达魔法水晶； 
    //     因此小扣负面效果下最少需要移动 7 次才能到达魔法水晶。
    //     <img src="https://pic.leetcode.cn/1681714676-gksEMT-image.png" />
    fmt.Println(challengeOfTheKeeper([]string{".....","##S..","...#.","T.#..","###.."})) // 7
    // 示例 2：
    // 输入：maze = [".#..","..##",".#S.",".#.T"]
    // 输出：-1
    // 解释：如下图所示。 
    //     若小扣向下移动至 [3,2]，守护者使其传送至 [0,2]，小扣将无法到达魔法水晶； 
    //     若小扣向右移动至 [2,3]，守护者使其传送至 [2,0]，小扣将无法到达魔法水晶；
    //     <img src="https://pic.leetcode.cn/1681714693-LsxKAh-image.png" />
    fmt.Println(challengeOfTheKeeper([]string{".#..","..##",".#S.",".#.T"})) // -1
    // 示例 3：
    // 输入：maze = ["S###.","..###","#..##","##..#","###.T"]
    // 输出：5
    // 解释：如下图所示： 守护者需要小扣在空地才能释放，因此初始无法将其从 [0,0] 传送至 [0,4]; 
    //       当小扣移动至 [2,1] 时，释放卷轴将其传送至水平方向的镜像位置 [2,1]（为原位置） 而后小扣需要移动 5 次才能到达魔法水晶。
    //       <img src="https://pic.leetcode.cn/1681800985-KrSdru-image.png" />
    fmt.Println(challengeOfTheKeeper([]string{"S###.","..###","#..##","##..#","###.T"})) // 5
}