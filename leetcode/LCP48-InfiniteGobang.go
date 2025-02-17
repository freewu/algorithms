package main

// LCP 48. 无限棋局
// 小力正在通过残局练习来备战「力扣挑战赛」中的「五子棋」项目，他想请你能帮他预测当前残局的输赢情况。
// 棋盘中的棋子分布信息记录于二维数组 pieces 中，其中 pieces[i] = [x,y,color] 表示第 i 枚棋子的横坐标为 x，纵坐标为 y，棋子颜色为 color(0 表示黑棋，1 表示白棋)。
// 假如黑棋先行，并且黑棋和白棋都按最优策略落子，请你求出当前棋局在三步（按 黑、白、黑 的落子顺序）之内的输赢情况（三步之内先构成同行、列或对角线连续同颜色的至少 5 颗即为获胜）：
//     1. 黑棋胜, 请返回 "Black"
//     2. 白棋胜, 请返回 "White"
//     3. 仍无胜者, 请返回 "None"

// 注意：
//     1. 和传统的五子棋项目不同，「力扣挑战赛」中的「五子棋」项目 不存在边界限制，即可在 任意位置 落子；
//     2. 黑棋和白棋均按 3 步内的输赢情况进行最优策略的选择
//     3. 测试数据保证所给棋局目前无胜者；
//     4. 测试数据保证不会存在坐标一样的棋子。

// 示例 1：
// 输入： pieces = [[0,0,1],[1,1,1],[2,2,0]]
// 输出："None"
// 解释：无论黑、白棋以何种方式落子，三步以内都不会产生胜者。

// 示例 2：
// 输入： pieces = [[1,2,1],[1,4,1],[1,5,1],[2,1,0],[2,3,0],[2,4,0],[3,2,1],[3,4,0],[4,2,1],[5,2,1]]
// 输出："Black"
// 解释：三步之内黑棋必胜，以下是一种可能的落子情况：
// <img src="https://pic.leetcode-cn.com/1629800639-KabOfY-902b87df29998b1c181146c8fdb3a4b6.gif" />

// 提示：
//     0 <= pieces.length <= 1000
//     pieces[i].length = 3
//     -10^9 <= pieces[i][0], pieces[i][1] <=10^9
//     0 <= pieces[i][2] <=1

import "fmt"

func gobang(pieces [][]int) string {
    const (
        black = "Black"
        white = "White"
        none  = "None"
    )
    type Pair struct{ x, y int }
    direactions := []Pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
    color := make(map[Pair]int, len(pieces) + 1)
    for _, p := range pieces {
        p[2]++ // 为方便利用零值，将黑改为 1，白改为 2
        color[Pair{p[0], p[1]}] = p[2]
    }
    // 在 (i,j) 落子，颜色为 c，判断落子的一方是否获胜
    checkWin := func(i, j, c int) bool {
        for k, d := range direactions[:4] {
            cnt := 1
            // 检查一个方向
            for x, y := i+d.x, j+d.y; color[Pair{x, y}] == c; x, y = x+d.x, y+d.y {
                cnt++
            }
            // 检查相反的另一方向
            d = direactions[k^4]
            for x, y := i+d.x, j+d.y; color[Pair{x, y}] == c; x, y = x+d.x, y+d.y {
                cnt++
            }
            if cnt >= 5 {
                return true
            }
        }
        return false
    }
    // 1. 黑第一手就可以获胜
    for _, p := range pieces {
        if p[2] == 2 {
            continue
        }
        i, j := p[0], p[1]
        for _, d := range direactions { // 黑要想一步获胜只能下在黑子周围
            if x, y := i+d.x, j+d.y; color[Pair{x, y}] == 0 && checkWin(x, y, 1) {
                return black
            }
        }
    }
    // 2. 黑第一手无法获胜
    whites := map[Pair]bool{}
    posW := Pair{}
    for _, p := range pieces {
        if p[2] == 1 {
            continue
        }
        i, j := p[0], p[1]
        for _, d := range direactions { // 白要想一步获胜只能下在白子周围
            x, y := i+d.x, j+d.y
            q := Pair{x, y}
            if color[q] == 0 && checkWin(x, y, 2) {
                // 2.1 白可以一步胜，且获胜位置不止一处
                if whites[q] = true; len(whites) > 1 {
                    return white
                }
                posW = q
            }
        }
    }
    // 2.2 白可以一步胜，但获胜位置只有一处
    if len(whites) == 1 {
        color[posW] = 1 // 黑第一手下在该处，阻止白获胜
        blacks := map[Pair]bool{}
        // 检查第三步的黑能否获胜
        checkBlackWin := func(i, j int) bool {
            for _, d := range direactions { // 黑要获胜只能下在黑子周围
                x, y := i+d.x, j+d.y
                p := Pair{x, y}
                if color[p] == 0 && checkWin(x, y, 1) {
                    if blacks[p] = true; len(blacks) > 1 {
                        return true
                    }
                }
            }
            return false
        }
        checkBlackWin(posW.x, posW.y)
        for _, p := range pieces {
            if p[2] == 1 && checkBlackWin(p[0], p[1]) {
                return black
            }
        }
        return none
    }
    // 3. 白无法获胜，于是白的策略是防止黑获胜
    // 根据黑第一步的落子位置，在该位置周围枚举黑第三步的落子位置，检查黑能否获胜
    checkBlackWin := func(i0, j0 int) bool {
        blacks := map[Pair]bool{}
        for k, d := range direactions {
            for l, i, j := 0, i0, j0; l < 5; l++ { // 如果黑可以获胜，这两枚黑子的距离不会超过 5
                i += d.x
                j += d.y
                p := Pair{i, j}
                if color[p] > 0 {
                    continue
                }
                cnt := 1
                // 检查一个方向
                for x, y := i+d.x, j+d.y; color[Pair{x, y}] == 1; x, y = x+d.x, y+d.y {
                    cnt++
                }
                // 检查相反的另一方向
                d2 := direactions[k^4]
                for x, y := i+d2.x, j+d2.y; color[Pair{x, y}] == 1; x, y = x+d2.x, y+d2.y {
                    cnt++
                }
                if cnt >= 5 {
                    if blacks[p] = true; len(blacks) > 1 {
                        return true
                    }
                }
            }
        }
        return false
    }
    vis := map[Pair]bool{} // 常数优化：避免重复枚举
    for _, p := range pieces {
        if p[2] == 2 {
            continue
        }
        i, j := p[0], p[1]
        // 枚举黑第一步的落子。由于黑要下两手棋，需要枚举黑子周围两圈
        for dx := -2; dx <= 2; dx++ {
            for dy := -2; dy <= 2; dy++ {
                if dx == 0 && dy == 0 {
                    continue
                }
                x, y := i+dx, j+dy
                q := Pair{x, y}
                if vis[q] || color[q] > 0 {
                    continue
                }
                color[q] = 1 // 黑落子
                vis[q] = true
                if checkBlackWin(x, y) {
                    return black
                }
                delete(color, q)
            }
        }
    }
    return none
}

func main() {
    // 示例 1：
    // 输入： pieces = [[0,0,1],[1,1,1],[2,2,0]]
    // 输出："None"
    // 解释：无论黑、白棋以何种方式落子，三步以内都不会产生胜者。
    fmt.Println(gobang([][]int{{0,0,1},{1,1,1},{2,2,0}})) // "None"
    // 示例 2：
    // 输入： pieces = [[1,2,1],[1,4,1],[1,5,1],[2,1,0],[2,3,0],[2,4,0],[3,2,1],[3,4,0],[4,2,1],[5,2,1]]
    // 输出："Black"
    // 解释：三步之内黑棋必胜，以下是一种可能的落子情况：
    // <img src="https://pic.leetcode-cn.com/1629800639-KabOfY-902b87df29998b1c181146c8fdb3a4b6.gif" />
    fmt.Println(gobang([][]int{{1,2,1},{1,4,1},{1,5,1},{2,1,0},{2,3,0},{2,4,0},{3,2,1},{3,4,0},{4,2,1},{5,2,1}})) // "Black"
}