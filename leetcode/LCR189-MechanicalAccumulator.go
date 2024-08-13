package main

// LCR 189. 设计机械累加器
// 请设计一个机械累加器，计算从 1、2... 一直累加到目标数值 target 的总和。
// 注意这是一个只能进行加法操作的程序，不具备乘除、if-else、switch-case、for 循环、while 循环，及条件判断语句等高级功能。

// 示例 1：
// 输入: target = 5
// 输出: 15

// 示例 2：
// 输入: target = 7
// 输出: 28

// 提示：
//     1 <= target <= 10000

import "fmt"

// 使用 for
func mechanicalAccumulator(target int) int {
    res := 0
    for i := 1; i <= target; i++ {
        res += i
    }
    return res
}

// 奇数 5 的累加：
//     1 + 5 | 2 + 4 | 3
//     这是 (5/2)(5+1) + (5+1)/2
// 偶数 6 的累加：
//     1 + 6 | 2 + 5 | 3 + 4
//     这是 (6/2)(6+1)
//     可以得出表达式 (target/2)*(target+1)+((target%2)*((target+1)/2))
func mechanicalAccumulator1(target int) int {
    return (target / 2) * (target + 1) + ((target % 2)*((target + 1)/2))
}

// 递归
func mechanicalAccumulator2(target int) int {
    if target == 0 { return 0 }
    return target + mechanicalAccumulator2(target - 1)
}

func mechanicalAccumulator3(target int) int {
    res := 0
    var sum func(target int) bool 
    sum = func(target int) bool {
        res += target
        return target > 0 && sum(target-1)
    }
    sum(target)
    return res
}

func main() {
    // 示例 1：
    // 输入: target = 5
    // 输出: 15
    fmt.Println(mechanicalAccumulator(5)) // 15  1 + 2 + 3 + 4 + 5
    // 示例 2：
    // 输入: target = 7
    // 输出: 28
    fmt.Println(mechanicalAccumulator(7)) // 28  1 + 2 + 3 + 4 + 5 + 6 + 7

    fmt.Println(mechanicalAccumulator1(5)) // 15  1 + 2 + 3 + 4 + 5
    fmt.Println(mechanicalAccumulator1(7)) // 28  1 + 2 + 3 + 4 + 5 + 6 + 7

    fmt.Println(mechanicalAccumulator2(5)) // 15  1 + 2 + 3 + 4 + 5
    fmt.Println(mechanicalAccumulator2(7)) // 28  1 + 2 + 3 + 4 + 5 + 6 + 7

    fmt.Println(mechanicalAccumulator3(5)) // 15  1 + 2 + 3 + 4 + 5
    fmt.Println(mechanicalAccumulator3(7)) // 28  1 + 2 + 3 + 4 + 5 + 6 + 7
}