package main

// LCP 62. 交通枢纽
// 为了缓解「力扣嘉年华」期间的人流压力，组委会在活动期间开设了一些交通专线。
// path[i] = [a, b]表示有一条从地点a通往地点b的单向交通专线。 
// 若存在一个地点，满足以下要求，我们则称之为交通枢纽：
//     1. 所有地点（除自身外）均有一条单向专线直接通往该地点；
//     2. 该地点不存在任何通往其他地点的单向专线。

// 请返回交通专线的交通枢纽。若不存在，则返回-1。

// 注意：
//     1. 对于任意一个地点，至少被一条专线连通。

// 示例 1：
// 输入：path = [[0,1],[0,3],[1,3],[2,0],[2,3]]
// 输出：3
// 解释：如下图所示： 地点0,1,2各有一条通往地点3的交通专线， 且地点3不存在任何通往其他地点的交通专线。
// <img src="https://pic.leetcode-cn.com/1663902572-yOlUCr-image.png" />

// 示例 2：
// 输入：path = [[0,3],[1,0],[1,3],[2,0],[3,0],[3,2]]
// 输出：-1
// 解释：如下图所示：不存在满足交通枢纽的地点。
// <img src="https://pic.leetcode-cn.com/1663902595-McsEkY-image.png" />

// 提示：
//     1 <= path.length <= 1000
//     0 <= path[i][0], path[i][1] <= 1000
//     path[i][0]与path[i][1]不相等

import "fmt"

func transportationHub(path [][]int) int {
    out, in, location := make(map[int]bool), make(map[int]int), make(map[int]bool) // 出, 入, 所有地点
    for i := 0; i < len(path); i++ {
        if !out[path[i][0]] { out[path[i][0]] = true }
        if !location[path[i][0]] { location[path[i][0]] = true }
        if !location[path[i][1]] { location[path[i][1]] = true }
        in[path[i][1]]++
    }
    for k, v := range in {
        if v == len(location) - 1 && !out[k] {
            return k
        }
    }
    return -1
}

func main() {
    // 示例 1：
    // 输入：path = [[0,1],[0,3],[1,3],[2,0],[2,3]]
    // 输出：3
    // 解释：如下图所示： 地点0,1,2各有一条通往地点3的交通专线， 且地点3不存在任何通往其他地点的交通专线。
    // <img src="https://pic.leetcode-cn.com/1663902572-yOlUCr-image.png" />
    fmt.Println(transportationHub([][]int{{0,1},{0,3},{1,3},{2,0},{2,3}})) // 3
    // 示例 2：
    // 输入：path = [[0,3],[1,0],[1,3],[2,0],[3,0],[3,2]]
    // 输出：-1
    // 解释：如下图所示：不存在满足交通枢纽的地点。
    // <img src="https://pic.leetcode-cn.com/1663902595-McsEkY-image.png" />
    fmt.Println(transportationHub([][]int{{0,3},{1,0},{1,3},{2,0},{3,0},{3,2}})) // -1
}