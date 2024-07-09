package main

// LCR 177. 撞色搭配
// 整数数组 sockets 记录了一个袜子礼盒的颜色分布情况，其中 sockets[i] 表示该袜子的颜色编号。
// 礼盒中除了一款撞色搭配的袜子，每种颜色的袜子均有两只。
// 请设计一个程序，在时间复杂度 O(n)，空间复杂度O(1) 内找到这双撞色搭配袜子的两个颜色编号。

// 示例 1：
// 输入：sockets = [4, 5, 2, 4, 6, 6]
// 输出：[2,5] 或 [5,2]

// 示例 2：
// 输入：sockets = [1, 2, 4, 1, 4, 3, 12, 3]
// 输出：[2,12] 或 [12,2]

// 提示：
//     2 <= sockets.length <= 10000

import "fmt"

// 位运算
func sockCollocation1(sockets []int) []int {
    tmp := 0
    for _, v := range sockets {
        tmp ^= v
    }
    flag := 1
    for flag & tmp == 0 {
        flag <<= 1
    }
    x, y := 0, 0
    for _, v := range sockets {
        if v & flag > 0 {
            x ^= v
        } else {
            y ^= v
        }
    }
    return []int{x, y}
}

// map
func sockCollocation(sockets []int) []int {
    res, mp := []int{}, make(map[int]int, len(sockets) / 2 + 1)
    for _, v := range sockets {
        mp[v]++
    }
    for i, c := range mp {
        if c == 1 { // 只出现了一次的
            res = append(res, i)
        } 
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：sockets = [4, 5, 2, 4, 6, 6]
    // 输出：[2,5] 或 [5,2]
    fmt.Println(sockCollocation([]int{4, 5, 2, 4, 6, 6})) // [2,5]
    // 示例 2：
    // 输入：sockets = [1, 2, 4, 1, 4, 3, 12, 3]
    // 输出：[2,12] 或 [12,2]
    fmt.Println(sockCollocation([]int{1, 2, 4, 1, 4, 3, 12, 3})) // [2,12]

    fmt.Println(sockCollocation1([]int{4, 5, 2, 4, 6, 6})) // [2,5]
    fmt.Println(sockCollocation1([]int{1, 2, 4, 1, 4, 3, 12, 3})) // [2,12]
}