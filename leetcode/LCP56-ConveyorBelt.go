package main

// LCP 56. 信物传送
// 欢迎各位勇者来到力扣城，本次试炼主题为「信物传送」。

// 本次试炼场地设有若干传送带，matrix[i][j]表示第i行j列的传送带运作方向，"^","v","<",">"这四种符号分别表示上、下、左、右四个方向。
// 信物会随传送带的方向移动。勇者每一次施法操作，可临时变更一处传送带的方向，在物品经过后传送带恢复原方向。
// <img src="https://pic.leetcode-cn.com/1649835246-vfupSL-lcp%20(2).gif" />

// 通关信物初始位于坐标start处，勇者需要将其移动到坐标end处，请返回勇者施法操作的最少次数。

// 注意：
//     start和end的格式均为[i,j]

// 示例 1：
// 输入：matrix = [">>v","v^<","<><"], start = [0,1], end = [2,0]
// 输出：1
// 解释： 如上图所示 当信物移动到[1,1]时，勇者施法一次将[1,1]的传送方向^从变更为<从而信物移动到[1,0]，后续到达end位置 因此勇者最少需要施法操作 1 次

// 示例 2：
// 输入：matrix = [">>v",">>v","^<<"], start = [0,0], end = [1,1]
// 输出：0
// 解释：勇者无需施法，信物将自动传送至end位置

// 示例 3：
// 输入：matrix = [">^^>","<^v>","^v^<"], start = [0,0], end = [1,3]
// 输出：3

// 提示：
//     matrix中仅包含'^'、'v'、'<'、'>'
//     0 < matrix.length <= 100
//     0 < matrix[i].length <= 100
//     0 <= start[0],end[0] < matrix.length
//     0 <= start[1],end[1] < matrix[i].length

import "fmt"
import "container/list"

func conveyorBelt1(matrix []string, start []int, end []int) int {
    type Node struct { i, j, d int }
    directions := map[byte][]int{
        '>': {0, 1},
        '<': {0, -1},
        '^': {-1, 0},
        'v': {1, 0},
    }
    m, n := len(matrix), len(matrix[0])
    dst := make([][]int, m)
    for i := range dst {
        dst[i] = make([]int, n)
        for j := range dst[i] {
            dst[i][j] = 1 << 31
        }
    }
    q := list.New()
    q.PushFront(Node{start[0], start[1], 0})
    for q.Len() != 0 {
        cur := q.Front().Value.(Node)
        q.Remove(q.Front())
        i, j, d := cur.i, cur.j, cur.d
        if i == end[0] && j == end[1] {
            return d
        }
        for k, v := range directions {
            x, y := i + v[0], j + v[1]
            if x < 0 || x >= m || y < 0 || y >= n { continue }
            if k == matrix[i][j] {
                if dst[x][y] > d {
                    dst[x][y] = d 
                    q.PushFront(Node{ x, y, d })
                }
            } else {
                if dst[x][y] > d + 1 {
                    dst[x][y] = d + 1 
                    q.PushBack(Node{ x, y, d + 1 })
                }
            }
        }
    }
    return -1
}

func conveyorBelt(matrix []string, start []int, end []int) int {
    n, m := len(matrix), len(matrix[0])
    total := n * m
    dst := make([][]int, n)
    for i := range dst {
        dst[i] = make([]int, m)
        for j := range dst[i] {
            dst[i][j] = total
        }
    }
    type Point struct{ x, y int }
    var directions = []Point{ {-1, 0}, {0, 1}, {1, 0}, {0, -1} }
    arrowCorrect := "v<^>"
    queue := []Point{{ end[0], end[1] }}
    dst[end[0]][end[1]] = 0
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        ux, uy := p.x, p.y
        for i := 0; i < 4; i++ {
            nx, ny := ux + directions[i].x, uy + directions[i].y
            if nx >= 0 && ny >= 0 && nx < n && ny < m {
                if matrix[nx][ny] == arrowCorrect[i] {
                    if dst[nx][ny] > dst[ux][uy] {
                        dst[nx][ny] = dst[ux][uy]
                        queue = append(queue, Point{ nx, ny })
                    }
                } else {
                    if dst[nx][ny] > dst[ux][uy]+1 {
                        dst[nx][ny] = dst[ux][uy] + 1
                        queue = append(queue, Point{ nx, ny })
                    }
                }
            }
        }
    }
    return dst[start[0]][start[1]]
}

func main() {
    // 示例 1：
    // 输入：matrix = [">>v","v^<","<><"], start = [0,1], end = [2,0]
    // 输出：1
    // 解释： 如上图所示 当信物移动到[1,1]时，勇者施法一次将[1,1]的传送方向^从变更为<从而信物移动到[1,0]，后续到达end位置 因此勇者最少需要施法操作 1 次
    fmt.Println(conveyorBelt([]string{">>v","v^<","<><"}, []int{0,1}, []int{2,0})) // 1
    // 示例 2：
    // 输入：matrix = [">>v",">>v","^<<"], start = [0,0], end = [1,1]
    // 输出：0
    // 解释：勇者无需施法，信物将自动传送至end位置
    fmt.Println(conveyorBelt([]string{">>v",">>v","^<<"}, []int{0,0}, []int{1,1})) // 0
    // 示例 3：
    // 输入：matrix = [">^^>","<^v>","^v^<"], start = [0,0], end = [1,3]
    // 输出：3
    fmt.Println(conveyorBelt([]string{">^^>","<^v>","^v^<"}, []int{0,0}, []int{1,3})) // 3

    fmt.Println(conveyorBelt1([]string{">>v","v^<","<><"}, []int{0,1}, []int{2,0})) // 1
    fmt.Println(conveyorBelt1([]string{">>v",">>v","^<<"}, []int{0,0}, []int{1,1})) // 0
    fmt.Println(conveyorBelt1([]string{">^^>","<^v>","^v^<"}, []int{0,0}, []int{1,3})) // 3
}