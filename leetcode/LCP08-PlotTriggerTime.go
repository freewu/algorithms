package main

// LCP 08. 剧情触发时间
// 在战略游戏中，玩家往往需要发展自己的势力来触发各种新的剧情。
// 一个势力的主要属性有三种，分别是文明等级（C），资源储备（R）以及人口数量（H）。
// 在游戏开始时（第 0 天），三种属性的值均为 0。

// 随着游戏进程的进行，每一天玩家的三种属性都会对应增加，我们用一个二维数组 increase 来表示每天的增加情况。
// 这个二维数组的每个元素是一个长度为 3 的一维数组，
// 例如 [[1,2,1],[3,4,2]] 表示第一天三种属性分别增加 1,2,1 而第二天分别增加 3,4,2。

// 所有剧情的触发条件也用一个二维数组 requirements 表示。
// 这个二维数组的每个元素是一个长度为 3 的一维数组，对于某个剧情的触发条件 c[i], r[i], h[i]，如果当前 C >= c[i] 且 R >= r[i] 且 H >= h[i] ，则剧情会被触发。

// 根据所给信息，请计算每个剧情的触发时间，并以一个数组返回。如果某个剧情不会被触发，则该剧情对应的触发时间为 -1 。

// 示例 1：
// 输入： increase = [[2,8,4],[2,5,0],[10,9,8]] requirements = [[2,11,3],[15,10,7],[9,17,12],[8,1,14]]
// 输出: [2,-1,3,-1]
// 解释：
// 初始时，C = 0，R = 0，H = 0
// 第 1 天，C = 2，R = 8，H = 4
// 第 2 天，C = 4，R = 13，H = 4，此时触发剧情 0
// 第 3 天，C = 14，R = 22，H = 12，此时触发剧情 2
// 剧情 1 和 3 无法触发。

// 示例 2：
// 输入： increase = [[0,4,5],[4,8,8],[8,6,1],[10,10,0]] requirements = [[12,11,16],[20,2,6],[9,2,6],[10,18,3],[8,14,9]]
// 输出: [-1,4,3,3,3]

// 示例 3：
// 输入： increase = [[1,1,1]] requirements = [[0,0,0]]
// 输出: [0]

// 限制：
//     1 <= increase.length <= 10000
//     1 <= requirements.length <= 100000
//     0 <= increase[i] <= 10
//     0 <= requirements[i] <= 100000

import "fmt"
import "sort"

func getTriggerTime(increase [][]int, requirements [][]int) []int {
    crh := make([][3]int, len(increase) + 1) // 存储每一天结束时的总 C，R，H
    crh[0] = [3]int{0, 0, 0} // 第 0 天，所有属性都是 0
    for i, inc := range increase { // 计算每一天的总 C，R，H
        crh[i+1] = [3]int{ crh[i][0] + inc[0], crh[i][1] + inc[1], crh[i][2] + inc[2], }
    }
    res := make([]int, len(requirements))
    for i, req := range requirements { // 遍历所有剧情触发条件
        res[i] = -1 // 默认不触发
        for j, r := range crh { // 遍历每一天的总 C，R，H 来检查是否满足条件
            if r[0] >= req[0] && r[1] >= req[1] && r[2] >= req[2] {
                res[i] = j // 记录触发时间
                break
            }
        }
    }
    return res
}

func getTriggerTime1(increase [][]int, requirements [][]int) []int {
    prefix := make([][3]int,len(increase)+1)
    for i := range increase {
        prefix[i+1] = [3]int{ prefix[i][0] + increase[i][0], prefix[i][1] + increase[i][1], prefix[i][2] + increase[i][2] }
    }
    res := make([]int,len(requirements))
    trigger := func(c, r, h int) (int,bool) {
        index := sort.Search(len(prefix), func(i int) bool {
            return prefix[i][0] >= c && prefix[i][1] >= r && prefix[i][2] >= h
        })
        if index == len(prefix) {
            return -1, false 
        }
        return index, prefix[index][0] >= c && prefix[index][1] >= r && prefix[index][2] >= h
    }
    for i, req := range requirements{
        index, flag := trigger(req[0], req[1], req[2])
        if flag {
            res[i] = index
        } else {
            res[i] = -1
        }
    }
    return res 
}

func main() {
    // 示例 1：
    // 输入： increase = [[2,8,4],[2,5,0],[10,9,8]] requirements = [[2,11,3],[15,10,7],[9,17,12],[8,1,14]]
    // 输出: [2,-1,3,-1]
    // 解释：
    // 初始时，C = 0，R = 0，H = 0
    // 第 1 天，C = 2，R = 8，H = 4
    // 第 2 天，C = 4，R = 13，H = 4，此时触发剧情 0
    // 第 3 天，C = 14，R = 22，H = 12，此时触发剧情 2
    // 剧情 1 和 3 无法触发。
    fmt.Println(getTriggerTime([][]int{{2,8,4},{2,5,0},{10,9,8}}, [][]int{{2,11,3},{15,10,7},{9,17,12},{8,1,14}})) // [2,-1,3,-1]
    // 示例 2：
    // 输入： increase = [[0,4,5],[4,8,8],[8,6,1],[10,10,0]] requirements = [[12,11,16],[20,2,6],[9,2,6],[10,18,3],[8,14,9]]
    // 输出: [-1,4,3,3,3]
    fmt.Println(getTriggerTime([][]int{{0,4,5},{4,8,8},{8,6,1},{10,10,0}}, [][]int{{12,11,16},{20,2,6},{9,2,6},{10,18,3},{8,14,9}})) // [-1,4,3,3,3]
    // 示例 3：
    // 输入： increase = [[1,1,1]] requirements = [[0,0,0]]
    // 输出: [0]
    fmt.Println(getTriggerTime([][]int{{1,1,1}}, [][]int{{0,0,0}})) // [0]

    fmt.Println(getTriggerTime1([][]int{{2,8,4},{2,5,0},{10,9,8}}, [][]int{{2,11,3},{15,10,7},{9,17,12},{8,1,14}})) // [2,-1,3,-1]
    fmt.Println(getTriggerTime1([][]int{{0,4,5},{4,8,8},{8,6,1},{10,10,0}}, [][]int{{12,11,16},{20,2,6},{9,2,6},{10,18,3},{8,14,9}})) // [-1,4,3,3,3]
    fmt.Println(getTriggerTime1([][]int{{1,1,1}}, [][]int{{0,0,0}})) // [0]
}