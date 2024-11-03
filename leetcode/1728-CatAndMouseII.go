package main

// 1728. Cat and Mouse II
// A game is played by a cat and a mouse named Cat and Mouse.

// The environment is represented by a grid of size rows x cols, 
// where each element is a wall, floor, player (Cat, Mouse), or food.
//     1. Players are represented by the characters 'C'(Cat),'M'(Mouse).
//     2. Floors are represented by the character '.' and can be walked on.
//     3. Walls are represented by the character '#' and cannot be walked on.
//     4. Food is represented by the character 'F' and can be walked on.
//     5. There is only one of each character 'C', 'M', and 'F' in grid.

// Mouse and Cat play according to the following rules:
//     1. Mouse moves first, then they take turns to move.
//     2. During each turn, Cat and Mouse can jump in one of the four directions (left, right, up, down). 
//        They cannot jump over the wall nor outside of the grid.
//     3. catJump, mouseJump are the maximum lengths Cat and Mouse can jump at a time, respectively. 
//        Cat and Mouse can jump less than the maximum length.
//     4. Staying in the same position is allowed.
//     5. Mouse can jump over Cat.

// The game can end in 4 ways:
//     1. If Cat occupies the same position as Mouse, Cat wins.
//     2. If Cat reaches the food first, Cat wins.
//     3. If Mouse reaches the food first, Mouse wins.
//     4. If Mouse cannot get to the food within 1000 turns, Cat wins.

// Given a rows x cols matrix grid and two integers catJump and mouseJump, 
// return true if Mouse can win the game if both Cat and Mouse play optimally, otherwise return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/12/sample_111_1955.png" />
// Input: grid = ["####F","#C...","M...."], catJump = 1, mouseJump = 2
// Output: true
// Explanation: Cat cannot catch Mouse on its turn nor can it get the food before Mouse.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/12/sample_2_1955.png" />
// Input: grid = ["M.C...F"], catJump = 1, mouseJump = 4
// Output: true

// Example 3:
// Input: grid = ["M.C...F"], catJump = 1, mouseJump = 3
// Output: false

// Constraints:
//     rows == grid.length
//     cols = grid[i].length
//     1 <= rows, cols <= 8
//     grid[i][j] consist only of characters 'C', 'M', 'F', '.', and '#'.
//     There is only one of each character 'C', 'M', and 'F' in grid.
//     1 <= catJump, mouseJump <= 8

import "fmt"

// func canMouseWin(grid []string, catJump int, mouseJump int) bool {
//     direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
//     n, m := len(grid), len(grid[0])
//     ocx, ocy, omx, omy, fx, fy := 0, 0, 0, 0, 0, 0
//     maxStep := m*n + 2
//     var dp [8][8][8][8][100]int
//     // 边界检查
//     borderCheck := func(x, y int) bool { return x < 0 || y < 0 || x >= n || y >= m }
//     // 找到 猫，老鼠，食物的位置
//     for i := 0; i < n; i++ {
//         for j := 0; j < m; j++ {
//             switch grid[i][j] {
//                 case 'C': ocx, ocy = i, j // cat
//                 case 'M': omx, omy = i, j // mouse
//                 case 'F': fx, fy = i, j // food
//             }
//         }
//     }
//     var dfs func(cx, cy, mx, my, step int) int
//     dfs = func(cx, cy, mx, my, step int) int { //  猫的当前位置，老鼠的当前位置，第几步
//         if v := dp[cx][cy][mx][my][step]; v != 0 { return v }
//         res := -1
//         if step & 1 == 1 { // step 从0 开始，偶数 为 老鼠，奇数为猫
//             // is Cat
//             // 可以尝试走的位置
//             tempPosition := [][]int{[]int{ cx,  }}
//             for _, dire := range direction { //方位
//                 dx, dy := dire[0], dire[1]
//                 for i := 1; i <= catJump; i++ { //可以跳的步数
//                     nx, ny := cx+dx*i, cy+dy*i
//                     if borderCheck(nx, ny) { break }
//                     if grid[nx][ny] == '#' { break }
//                     if nx == mx && ny == my { // win 判断
//                         res = 1
//                         break
//                     }
//                     if nx == fx && ny == fy { // win 判断
//                         res = 1
//                         break
//                     }
//                     tempPosition = append(tempPosition, []int{nx, ny})
//                 }
//                 if res == 1 { break } // 找必赢的方案，只要有赢的方案，就可以跳出循环了
//             }
//             // 如果没有赢，尝试以下步骤，找对方输
//             for res != 1 && len(tempPosition) > 0 {
//                 position := tempPosition[0]
//                 tempPosition = tempPosition[1:]
//                 if dfs(position[0], position[1], mx, my, step + 1) == -1 {
//                     res = 1
//                     break
//                 }
//             }
//         } else if step <= maxStep {
//             //it is mouse's turn and step is less than maxStep
//             tempPosition := [][]int{[]int{mx, my}} // 原地不动
//             for _, dire := range direction {
//                 dx, dy := dire[0], dire[1]
//                 for i := 1; i <= mouseJump; i++ {
//                     nx, ny := mx+dx*i, my+dy*i
//                     if borderCheck(nx, ny) { break }
//                     if grid[nx][ny] == '#' { break }
//                     if nx == fx && ny == fy { res = 1; break; }
//                     if nx == cx && ny == cy { continue }
//                     tempPosition = append(tempPosition, []int{ nx, ny})
//                 }
//                 if res == 1 { break }
//             }
//             // 逻辑同猫
//             for res != 1 && len(tempPosition) > 0 {
//                 position := tempPosition[0]
//                 tempPosition = tempPosition[1:]
//                 if dfs(cx, cy, position[0], position[1], step + 1) == -1 {
//                     res = 1
//                     break
//                 }
//             }
//         }
//         dp[cx][cy][mx][my][step] = res
//         return res
//     }
//     return dfs(ocx, ocy, omx, omy, 0) == 1
// }

const (
    MouseTurn = 0
    CatTurn   = 1
    UNKNOWN   = 0
    MouseWin  = 1
    CatWin    = 2
    MaxMoves  = 1000
)

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
    rows, cols := len(grid), len(grid[0])
    getPos := func(row, col int) int { return row*cols + col }
    var startMouse, startCat, food int
    for i, row := range grid {
        for j, ch := range row {
            if ch == 'M' {
                startMouse = getPos(i, j)
            } else if ch == 'C' {
                startCat = getPos(i, j)
            } else if ch == 'F' {
                food = getPos(i, j)
            }
        }
    }

    // 计算每个状态的度
    total := rows * cols
    degrees := [64][64][2]int{}
    for mouse := 0; mouse < total; mouse++ {
        mouseRow := mouse / cols
        mouseCol := mouse % cols
        if grid[mouseRow][mouseCol] == '#' {
            continue
        }
        for cat := 0; cat < total; cat++ {
            catRow := cat / cols
            catCol := cat % cols
            if grid[catRow][catCol] == '#' {
                continue
            }
            degrees[mouse][cat][MouseTurn]++
            degrees[mouse][cat][CatTurn]++
            for _, dir := range dirs {
                for row, col, jump := mouseRow+dir.x, mouseCol+dir.y, 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != '#' && jump <= mouseJump; jump++ {
                    nextMouse := getPos(row, col)
                    nextCat := getPos(catRow, catCol)
                    degrees[nextMouse][nextCat][MouseTurn]++
                    row += dir.x
                    col += dir.y
                }
                for row, col, jump := catRow+dir.x, catCol+dir.y, 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != '#' && jump <= catJump; jump++ {
                    nextMouse := getPos(mouseRow, mouseCol)
                    nextCat := getPos(row, col)
                    degrees[nextMouse][nextCat][CatTurn]++
                    row += dir.x
                    col += dir.y
                }
            }
        }
    }

    results := [64][64][2][2]int{}
    type state struct{ mouse, cat, turn int }
    q := []state{}

    // 猫和老鼠在同一个单元格，猫获胜
    for pos := 0; pos < total; pos++ {
        row := pos / cols
        col := pos % cols
        if grid[row][col] == '#' {
            continue
        }
        results[pos][pos][MouseTurn][0] = CatWin
        results[pos][pos][MouseTurn][1] = 0
        results[pos][pos][CatTurn][0] = CatWin
        results[pos][pos][CatTurn][1] = 0
        q = append(q, state{pos, pos, MouseTurn}, state{pos, pos, CatTurn})
    }

    // 猫和食物在同一个单元格，猫获胜
    for mouse := 0; mouse < total; mouse++ {
        mouseRow := mouse / cols
        mouseCol := mouse % cols
        if grid[mouseRow][mouseCol] == '#' || mouse == food {
            continue
        }
        results[mouse][food][MouseTurn][0] = CatWin
        results[mouse][food][MouseTurn][1] = 0
        results[mouse][food][CatTurn][0] = CatWin
        results[mouse][food][CatTurn][1] = 0
        q = append(q, state{mouse, food, MouseTurn}, state{mouse, food, CatTurn})
    }

    // 老鼠和食物在同一个单元格且猫和食物不在同一个单元格，老鼠获胜
    for cat := 0; cat < total; cat++ {
        catRow := cat / cols
        catCol := cat % cols
        if grid[catRow][catCol] == '#' || cat == food {
            continue
        }
        results[food][cat][MouseTurn][0] = MouseWin
        results[food][cat][MouseTurn][1] = 0
        results[food][cat][CatTurn][0] = MouseWin
        results[food][cat][CatTurn][1] = 0
        q = append(q, state{food, cat, MouseTurn}, state{food, cat, CatTurn})
    }

    getPrevStates := func(mouse, cat, turn int) []state {
        mouseRow := mouse / cols
        mouseCol := mouse % cols
        catRow := cat / cols
        catCol := cat % cols
        prevTurn := MouseTurn
        if turn == MouseTurn {
            prevTurn = CatTurn
        }
        maxJump, startRow, startCol := catJump, catRow, catCol
        if prevTurn == MouseTurn {
            maxJump, startRow, startCol = mouseJump, mouseRow, mouseCol
        }
        prevStates := []state{{mouse, cat, prevTurn}}
        for _, dir := range dirs {
            for i, j, jump := startRow+dir.x, startCol+dir.y, 1; i >= 0 && i < rows && j >= 0 && j < cols && grid[i][j] != '#' && jump <= maxJump; jump++ {
                prevMouseRow := mouseRow
                prevMouseCol := mouseCol
                prevCatRow := i
                prevCatCol := j
                if prevTurn == MouseTurn {
                    prevMouseRow = i
                    prevMouseCol = j
                    prevCatRow = catRow
                    prevCatCol = catCol
                }
                prevMouse := getPos(prevMouseRow, prevMouseCol)
                prevCat := getPos(prevCatRow, prevCatCol)
                prevStates = append(prevStates, state{prevMouse, prevCat, prevTurn})
                i += dir.x
                j += dir.y
            }
        }
        return prevStates
    }

    // 拓扑排序
    for len(q) > 0 {
        s := q[0]
        q = q[1:]
        mouse, cat, turn := s.mouse, s.cat, s.turn
        result := results[mouse][cat][turn][0]
        moves := results[mouse][cat][turn][1]
        for _, s := range getPrevStates(mouse, cat, turn) {
            prevMouse, prevCat, prevTurn := s.mouse, s.cat, s.turn
            if results[prevMouse][prevCat][prevTurn][0] == UNKNOWN {
                canWin := result == MouseWin && prevTurn == MouseTurn || result == CatWin && prevTurn == CatTurn
                if canWin {
                    results[prevMouse][prevCat][prevTurn][0] = result
                    results[prevMouse][prevCat][prevTurn][1] = moves + 1
                    q = append(q, state{prevMouse, prevCat, prevTurn})
                } else {
                    degrees[prevMouse][prevCat][prevTurn]--
                    if degrees[prevMouse][prevCat][prevTurn] == 0 {
                        loseResult := MouseWin
                        if prevTurn == MouseTurn {
                            loseResult = CatWin
                        }
                        results[prevMouse][prevCat][prevTurn][0] = loseResult
                        results[prevMouse][prevCat][prevTurn][1] = moves + 1
                        q = append(q, state{prevMouse, prevCat, prevTurn})
                    }
                }
            }
        }
    }
    return results[startMouse][startCat][MouseTurn][0] == MouseWin && results[startMouse][startCat][MouseTurn][1] <= MaxMoves
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/12/sample_111_1955.png" />
    // Input: grid = ["####F","#C...","M...."], catJump = 1, mouseJump = 2
    // Output: true
    // Explanation: Cat cannot catch Mouse on its turn nor can it get the food before Mouse.
    fmt.Println(canMouseWin([]string{"####F","#C...","M...."}, 1, 2)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/12/sample_2_1955.png" />
    // Input: grid = ["M.C...F"], catJump = 1, mouseJump = 4
    // Output: true
    fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 4)) // true
    // Example 3:
    // Input: grid = ["M.C...F"], catJump = 1, mouseJump = 3
    // Output: false
    fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 3)) // false

    fmt.Println(canMouseWin([]string{"C#......","M..####.","###.....","....####",".####...","......#.","#######.","F......."}, 1, 1)) // true
}