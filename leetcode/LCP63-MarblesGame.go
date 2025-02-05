package main

// LCP 63. 弹珠游戏
// 欢迎各位来到「力扣嘉年华」，接下来将为各位介绍在活动中广受好评的弹珠游戏。

// N*M 大小的弹珠盘的初始状态信息记录于一维字符串型数组 plate 中，
// 数组中的每个元素为仅由 "O"、"W"、"E"、"." 组成的字符串。其中：
//     1. "O" 表示弹珠洞（弹珠到达后会落入洞中，并停止前进）；
//     2. "W" 表示逆时针转向器（弹珠经过时方向将逆时针旋转 90 度）；
//     3. "E" 表示顺时针转向器（弹珠经过时方向将顺时针旋转 90 度）；
//     4. "." 表示空白区域（弹珠可通行）。

// 游戏规则要求仅能在边缘位置的 空白区域 处（弹珠盘的四角除外）沿 与边缘垂直 的方向打入弹珠，并且打入后的每颗弹珠最多能 前进 num 步。
// 请返回符合上述要求且可以使弹珠最终入洞的所有打入位置。你可以 按任意顺序 返回答案。

// 注意：
//     若弹珠已到达弹珠盘边缘并且仍沿着出界方向继续前进，则将直接出界。

// 示例 1：
// 输入： num = 4 plate = ["..E.",".EOW","..W."]
// 输出：[[2,1]]
// 解释： 在 [2,1] 处打入弹珠，弹珠前进 1 步后遇到转向器，前进方向顺时针旋转 90 度，再前进 1 步进入洞中。
// <img src="https://pic.leetcode-cn.com/1630392649-BoQncz-b054955158a99167b8d51da0e22a54da.gif" />

// 示例 2：
// 输入： num = 5 plate = [".....","..E..",".WO..","....."]
// 输出：[[0,1],[1,0],[2,4],[3,2]]
// 解释： 在 [0,1] 处打入弹珠，弹珠前进 2 步，遇到转向器后前进方向逆时针旋转 90 度，再前进 1 步进入洞中。 
// 在 [1,0] 处打入弹珠，弹珠前进 2 步，遇到转向器后前进方向顺时针旋转 90 度，再前进 1 步进入洞中。
// 在 [2,4] 处打入弹珠，弹珠前进 2 步后进入洞中。 在 [3,2] 处打入弹珠，弹珠前进 1 步后进入洞中。
// <img src="https://pic.leetcode-cn.com/1630392625-rckbdy-b44e9963239ae368badf3d00b7563087.gif" />

// 示例 3：
// 输入： num = 3 plate = [".....","....O","....O","....."]
// 输出：[]
// 解释： 由于弹珠被击中后只能前进 3 步，且不能在弹珠洞和弹珠盘四角打入弹珠，故不存在能让弹珠入洞的打入位置。

// 提示：
//     1 <= num <= 10^6
//     1 <= plate.length, plate[i].length <= 1000
//     plate[i][j] 仅包含 "O"、"W"、"E"、"."

import "fmt"

// brute force
func ballGame(num int, plate []string) [][]int {
    directions := []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上（顺时针）
    m, n := len(plate), len(plate[0])
    check := func(x, y, d int) bool {
        for left := num; plate[x][y] != 'O'; left-- {
            if left == 0 { return false } // 无剩余步数
            if plate[x][y] == 'W' { // 逆时针
                d = (d + 3) % 4
            } else if plate[x][y] == 'E' { // 顺时针
                d = (d + 1) % 4
            }
            x += directions[d].x
            y += directions[d].y
            if x < 0 || x >= m || y < 0 || y >= n { // 从另一边出去了
                return false
            }
        }
        return true
    }
    res := [][]int{}
    for j := 1; j < n-1; j++ {
        if plate[0][j] == '.'     && check(0, j, 1)     { res = append(res, []int{0, j}) }
        if plate[m - 1][j] == '.' && check(m - 1, j, 3) { res = append(res, []int{m - 1, j}) }
    }
    for i := 1; i < m - 1; i++ {
        if plate[i][0] == '.'     && check(i, 0, 0)     { res = append(res, []int{i, 0}) }
        if plate[i][n - 1] == '.' && check(i, n - 1, 2) { res = append(res, []int{i, n - 1}) }
    }
    return res
}

func ballGame1(num int, plate []string) [][]int {
    // 仅能在边缘位置的空白区域处，沿垂直方向打入弹珠
    directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // 上 左 下 右
    var dfs func(r int, c int, step int, d int) bool
    dfs = func(r int, c int, step int, d int) bool {
        x, y := r + directions[d][0], c + directions[d][1]
        if x > len(plate)-1 || x < 0 || y > len(plate[0])-1 || y < 0 { return false }// 边界检测
        step++
        if step > num { return false }
        if plate[x][y] == 'O' {
            return true
        } else if plate[x][y] == 'W' { // 逆时针
            d = (d + 1) % 4
        } else if plate[x][y] == 'E' { // 顺时针
            d = (d + 3) % 4
        }
        if dfs(x, y, step, d) { return true }
        return false
    }
    res := make([][]int, 0)
    for c := 1; c < len(plate[0])-1; c++ { // 上下
        if plate[0][c] == '.' && dfs(0, c, 0, 2) {
            res = append(res, []int{0, c})
        }
        if plate[len(plate)-1][c] == '.' && dfs(len(plate)-1, c, 0, 0) {
            res = append(res, []int{len(plate) - 1, c})
        }
    }
    for r := 1; r < len(plate)-1; r++ { // 左右
        if plate[r][0] == '.' && dfs(r, 0, 0, 3) {
            res = append(res, []int{r, 0})
        }
        if plate[r][len(plate[0])-1] == '.' && dfs(r, len(plate[0])-1, 0, 1) {
            res = append(res, []int{r, len(plate[0]) - 1})
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入： num = 4 plate = ["..E.",".EOW","..W."]
    // 输出：[[2,1]]
    // 解释： 在 [2,1] 处打入弹珠，弹珠前进 1 步后遇到转向器，前进方向顺时针旋转 90 度，再前进 1 步进入洞中。
    // <img src="https://pic.leetcode-cn.com/1630392649-BoQncz-b054955158a99167b8d51da0e22a54da.gif" />
    fmt.Println(ballGame(4, []string{"..E.",".EOW","..W."})) // [[2,1]]
    // 示例 2：
    // 输入： num = 5 plate = [".....","..E..",".WO..","....."]
    // 输出：[[0,1],[1,0],[2,4],[3,2]]
    // 解释： 在 [0,1] 处打入弹珠，弹珠前进 2 步，遇到转向器后前进方向逆时针旋转 90 度，再前进 1 步进入洞中。 
    // 在 [1,0] 处打入弹珠，弹珠前进 2 步，遇到转向器后前进方向顺时针旋转 90 度，再前进 1 步进入洞中。
    // 在 [2,4] 处打入弹珠，弹珠前进 2 步后进入洞中。 在 [3,2] 处打入弹珠，弹珠前进 1 步后进入洞中。
    // <img src="https://pic.leetcode-cn.com/1630392625-rckbdy-b44e9963239ae368badf3d00b7563087.gif" />
    fmt.Println(ballGame(5, []string{".....","..E..",".WO..","....."})) // [[0,1],[1,0],[2,4],[3,2]]
    // 示例 3：
    // 输入： num = 3 plate = [".....","....O","....O","....."]
    // 输出：[]
    // 解释： 由于弹珠被击中后只能前进 3 步，且不能在弹珠洞和弹珠盘四角打入弹珠，故不存在能让弹珠入洞的打入位置。
    fmt.Println(ballGame(3, []string{".....","....O","....O","....."})) // []

    fmt.Println(ballGame1(4, []string{"..E.",".EOW","..W."})) // [[2,1]]
    fmt.Println(ballGame1(5, []string{".....","..E..",".WO..","....."})) // [[0,1],[1,0],[2,4],[3,2]]
    fmt.Println(ballGame1(3, []string{".....","....O","....O","....."})) // []
}