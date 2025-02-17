package main

// LCP 43. 十字路口的交通
// 前往「力扣挑战赛」场馆的道路上，有一个拥堵的十字路口，该十字路口由两条双向两车道的路交叉构成。
// 由于信号灯故障，交警需要手动指挥拥堵车辆。
// 假定路口没有新的来车且一辆车从一个车道驶入另一个车道所需的时间恰好为一秒钟，长度为 4 的一维字符串数组 directions 中按照 东、南、西、北 顺序记录了四个方向从最靠近路口到最远离路口的车辆计划开往的方向。
// 其中：
//     "E" 表示向东行驶；
//     "S" 表示向南行驶；
//     "W" 表示向西行驶；
//     "N" 表示向北行驶。

// 交警每秒钟只能指挥各个车道距离路口最近的一辆车，且每次指挥需要满足如下规则：
//     1. 同一秒钟内，一个方向的车道只允许驶出一辆车；
//     2. 同一秒钟内，一个方向的车道只允许驶入一辆车；
//     3. 同一秒钟内，车辆的行驶路线不可相交。

// 请返回最少需要几秒钟，该十字路口等候的车辆才能全部走完。

// 各个车道驶出的车辆可能的行驶路线如图所示：

// <img src="https://pic.leetcode-cn.com/1630393755-gyPeMM-%E5%9B%BE%E7%89%87.png">

// 注意：
//     测试数据保证不会出现掉头行驶指令，即某一方向的行驶车辆计划开往的方向不会是当前车辆所在的车道的方向;
//     表示堵塞车辆行驶方向的字符串仅用大写字母 "E"，"N"，"W"，"S" 表示。

// 示例 1：
// 输入：directions = ["W","N","ES","W"]
// 输出：2
// 解释：  第 1 秒：东西方向排在最前的车先行，剩余车辆状态 ["","N","S","W"]；  
//        第 2 秒：南、西、北方向的车行驶，路口无等待车辆； 
//        因此最少需要 2 秒，返回 2。

// 示例 2：
// 输入：directions = ["NS","WE","SE","EW"]
// 输出：3
// 解释： 第 1 秒：四个方向排在最前的车均可驶出； 
//        第 2 秒：东南方向的车驶出，剩余车辆状态 ["","","E","W"]； 
//        第 3 秒：西北方向的车驶出。

// 提示：
//     directions.length = 4
//     0 <= directions[i].length <= 20

import "fmt"
import "math/bits"

var mp = []int{'E': 0, 'S': 1, 'W': 2, 'N': 3} // 将方向字符映射到 0-3 上
var valid = [4][4][4]int{ // 当前车去方向，另一辆车来方向，另一辆车去方向
    {},
    {
        {},
        {1},
        {0, 0, 0, 1},
        {0, 0, 1},
    },
    {
        {},
        {1},
        {1, 1},
    },
    {
        {},
        {1, 0, 1},
        {1, 1},
        {1, 1, 1},
    },
}

// 两辆车 X 和 Y，将 X 旋转至方向 0 上，Y 跟着 X 一起旋转
// 这样上面就不用写一个很长的 valid 数组
func ok(fromX, toX, fromY, toY int) bool {
    toX = (toX - fromX + 4) % 4
    fromY = (fromY - fromX + 4) % 4
    toY = (toY - fromX + 4) % 4
    return valid[toX][fromY][toY] == 1
}

func trafficCommand(ds []string) int {
    n0, n1, n2, n3 := len(ds[0]), len(ds[1]), len(ds[2]), len(ds[3])
    dp := [21][21][21][21]int{}
    vis := [21][21][21][21]bool{}

    var f func(int, int, int, int) int
    f = func(p0, p1, p2, p3 int) (res int) {
        if p0 == n0 && p1 == n1 && p2 == n2 && p3 == n3 {
            return
        }
        dv := &dp[p0][p1][p2][p3]
        if vis[p0][p1][p2][p3] {
            return *dv
        }
        vis[p0][p1][p2][p3] = true
        defer func() { *dv = res }()
        ps := [4]int{p0, p1, p2, p3}
        res = 1e9
    outer:
        for sub := uint(1); sub < 16; sub++ { // 枚举选择哪几个方向的车通行
            a := [][2]int{}
            for s := sub; s > 0; s &= s - 1 {
                from := bits.TrailingZeros(s)
                if ps[from] == len(ds[from]) {
                    continue outer
                }
                to := mp[ds[from][ps[from]]]
                for _, q := range a {
                    if !ok(from, to, q[0], q[1]) {
                        continue outer
                    }
                }
                a = append(a, [2]int{from, to})
            }
            for _, p := range a {
                ps[p[0]]++
            }
            res = min(res, f(ps[0], ps[1], ps[2], ps[3])+1)
            for _, p := range a {
                ps[p[0]]--
            }
        }
        return
    }
    return f(0, 0, 0, 0)
}

func main() {
    // 示例 1：
    // 输入：directions = ["W","N","ES","W"]
    // 输出：2
    // 解释：  第 1 秒：东西方向排在最前的车先行，剩余车辆状态 ["","N","S","W"]；  
    //        第 2 秒：南、西、北方向的车行驶，路口无等待车辆； 
    //        因此最少需要 2 秒，返回 2。
    fmt.Println(trafficCommand([]string{"W","N","ES","W"})) // 2
    // 示例 2：
    // 输入：directions = ["NS","WE","SE","EW"]
    // 输出：3
    // 解释： 第 1 秒：四个方向排在最前的车均可驶出； 
    //        第 2 秒：东南方向的车驶出，剩余车辆状态 ["","","E","W"]； 
    //        第 3 秒：西北方向的车驶出。
    fmt.Println(trafficCommand([]string{"NS","WE","SE","EW"})) // 3
}