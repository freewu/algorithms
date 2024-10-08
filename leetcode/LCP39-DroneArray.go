package main

// LCP 39. 无人机方阵
// 在 「力扣挑战赛」 开幕式的压轴节目 「无人机方阵」中，每一架无人机展示一种灯光颜色。 无人机方阵通过两种操作进行颜色图案变换：
//     调整无人机的位置布局
//     切换无人机展示的灯光颜色

// 给定两个大小均为 N*M 的二维数组 source 和 target 表示无人机方阵表演的两种颜色图案，由于无人机切换灯光颜色的耗能很大，
// 请返回从 source 到 target 最少需要多少架无人机切换灯光颜色。

// 注意： 调整无人机的位置布局时无人机的位置可以随意变动。

// 示例 1：
// 输入：source = [[1,3],[5,4]], target = [[3,1],[6,5]]
// 输出：1
// 解释： 
//     最佳方案为 将 [0,1] 处的无人机移动至 [0,0] 处； 将 [0,0] 处的无人机移动至 [0,1] 处； 将 [1,0] 处的无人机移动至 [1,1] 处； 
//     将 [1,1] 处的无人机移动至 [1,0] 处，其灯光颜色切换为颜色编号为 6 的灯光； 
//     因此从source 到 target 所需要的最少灯光切换次数为 1。
//     <img src="https://pic.leetcode-cn.com/1628823765-uCDaux-8819ccdd664e91c78cde3bba3c701986.gif" />

// 示例 2：
// 输入：source = [[1,2,3],[3,4,5]], target = [[1,3,5],[2,3,4]]
// 输出：0 解释： 仅需调整无人机的位置布局，便可完成图案切换。因此不需要无人机切换颜色

// 提示： 
//     n == source.length == target.length 
//     m == source[i].length == target[i].length 
//     1 <= n, m <=100 
//     1 <= source[i][j], target[i][j] <=10^4

import "fmt"

func minimumSwitchingTimes(source [][]int, target [][]int) int {
    res, mp := 0, make([]int, 10001)
    for i := range source {
        for j := range source[i] {
            mp[source[i][j]]++
        }
    }
    for i := range target {
        for j := range target[i] {
            if mp[target[i][j]] > 0 {
                mp[target[i][j]]--
            }
        }
    }
    for _, v := range mp {
        res += v
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：source = [[1,3],[5,4]], target = [[3,1],[6,5]]
    // 输出：1
    // 解释： 
    //     最佳方案为 将 [0,1] 处的无人机移动至 [0,0] 处； 将 [0,0] 处的无人机移动至 [0,1] 处； 将 [1,0] 处的无人机移动至 [1,1] 处； 
    //     将 [1,1] 处的无人机移动至 [1,0] 处，其灯光颜色切换为颜色编号为 6 的灯光； 
    //     因此从source 到 target 所需要的最少灯光切换次数为 1。
    //     <img src="https://pic.leetcode-cn.com/1628823765-uCDaux-8819ccdd664e91c78cde3bba3c701986.gif" />
    fmt.Println(minimumSwitchingTimes([][]int{{1, 3},{5, 4}}, [][]int{{3,1},{6,5}})) // 1
    // 示例 2：
    // 输入：source = [[1,2,3],[3,4,5]], target = [[1,3,5],[2,3,4]]
    // 输出：0 解释： 仅需调整无人机的位置布局，便可完成图案切换。因此不需要无人机切换颜色
    fmt.Println(minimumSwitchingTimes([][]int{{1,2,3}, {3,4,5}}, [][]int{{1,3,5},{2,3,4}})) // 1
}