package main

// LCP 03. 机器人大冒险
// 力扣团队买了一个可编程机器人，机器人初始位置在原点(0, 0)。
// 小伙伴事先给机器人输入一串指令command，机器人就会无限循环这条指令的步骤进行移动。
// 指令有两种：
//     U: 向y轴正方向移动一格
//     R: 向x轴正方向移动一格。

// 不幸的是，在 xy 平面上还有一些障碍物，他们的坐标用obstacles表示。
// 机器人一旦碰到障碍物就会被损毁。

// 给定终点坐标(x, y)，返回机器人能否完好地到达终点。如果能，返回true；否则返回false。

// 示例 1：
// 输入：command = "URR", obstacles = [], x = 3, y = 2
// 输出：true
// 解释：U(0, 1) -> R(1, 1) -> R(2, 1) -> U(2, 2) -> R(3, 2)。

// 示例 2：
// 输入：command = "URR", obstacles = [[2, 2]], x = 3, y = 2
// 输出：false
// 解释：机器人在到达终点前会碰到(2, 2)的障碍物。

// 示例 3：
// 输入：command = "URR", obstacles = [[4, 2]], x = 3, y = 2
// 输出：true
// 解释：到达终点后，再碰到障碍物也不影响返回结果。

// 限制：
//     2 <= command的长度 <= 1000
//     command由U，R构成，且至少有一个U，至少有一个R
//     0 <= x <= 1e9, 0 <= y <= 1e9
//     0 <= obstacles的长度 <= 1000
//     obstacles[i]不为原点或者终点

import "fmt"
import "strings"

func robot(command string, obstacles [][]int, x int, y int) bool {
    isOnThePath := func(command string, x int, y int) bool {
        u := strings.Count(command, "U") * ((x + y) / len(command)) + strings.Count(command[0:(x + y) % len(command)], "U")
        r := strings.Count(command, "R") * ((x + y) / len(command)) + strings.Count(command[0:(x + y) % len(command)], "R")
        if u == y && r == x { return true }
        return false
    }
    // 如果目标点不在路径上，返回失败
    if !isOnThePath(command, x, y) {
        return false
    }
    for _, o := range obstacles {
        // 判断有效的故障点是否在路径上（故障的步数大于等于目标的点，视为无效故障）
        if (x + y > o[0] + o[1]) && isOnThePath(command, o[0], o[1]) {
            return false
        }
    }
    return true
}

func main() {
    // 示例 1：
    // 输入：command = "URR", obstacles = [], x = 3, y = 2
    // 输出：true
    // 解释：U(0, 1) -> R(1, 1) -> R(2, 1) -> U(2, 2) -> R(3, 2)。
    fmt.Println(robot("URR", [][]int{}, 3, 2)) // true
    // 示例 2：
    // 输入：command = "URR", obstacles = [[2, 2]], x = 3, y = 2
    // 输出：false
    // 解释：机器人在到达终点前会碰到(2, 2)的障碍物。
    fmt.Println(robot("URR", [][]int{{2, 2}}, 3, 2)) // false
    // 示例 3：
    // 输入：command = "URR", obstacles = [[4, 2]], x = 3, y = 2
    // 输出：true
    // 解释：到达终点后，再碰到障碍物也不影响返回结果。
    fmt.Println(robot("URR", [][]int{{4, 2}}, 3, 2)) // true
}