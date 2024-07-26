package main

// LCR 187. 破冰游戏
// 社团共有 num 位成员参与破冰游戏，编号为 0 ~ num-1。
// 成员们按照编号顺序围绕圆桌而坐。社长抽取一个数字 target，从 0 号成员起开始计数，排在第 target 位的成员离开圆桌，且成员离开后从下一个成员开始计数。
// 请返回游戏结束时最后一位成员的编号。

// 示例 1：
// 输入：num = 7, target = 4
// 输出：1

// 示例 2：
// 输入：num = 12, target = 5
// 输出：0

// 提示：
//     1 <= num <= 10^5
//     1 <= target <= 10^6

import "fmt"

// 递归
func iceBreakingGame(num int, target int) int {
    if (num == 1) {
        return 0
    }
    prevRemaining := iceBreakingGame(num - 1, target)
    return (prevRemaining + target) % num
}

// 模拟
func iceBreakingGame1(num int, target int) int {
    remaining := 0
    for i := 2; i <= num; i++ {
        remaining = (remaining + target) % i
    }
    return remaining
}

func main() {
    // 示例 1：
    // 输入：num = 7, target = 4
    // 输出：1
    fmt.Println(iceBreakingGame(7, 4)) // 1
    // 示例 2：
    // 输入：num = 12, target = 5
    // 输出：0
    fmt.Println(iceBreakingGame(12, 5)) // 0

    fmt.Println(iceBreakingGame1(7, 4)) // 1
    fmt.Println(iceBreakingGame1(12, 5)) // 0
}