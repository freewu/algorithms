package main

// LCP 53. 守护太空城
// 各位勇者请注意，力扣太空城发布陨石雨红色预警。

// 太空城中的一些舱室将要受到陨石雨的冲击，这些舱室按照编号0 ~ N的顺序依次排列。
// 为了阻挡陨石损毁舱室，太空城可以使用能量展开防护屏障，具体消耗如下：
//     1. 选择一个舱室开启屏障，能量消耗为2
//     2. 选择相邻两个舱室开启联合屏障，能量消耗为3
//     3. 对于已开启的一个屏障，多维持一时刻，能量消耗为1

// 已知陨石雨的影响范围和到达时刻，time[i]和position[i]分别表示该陨石的到达时刻和冲击位置。
// 请返回太空舱能够守护所有舱室所需要的最少能量。

// 注意：
//     同一时间，一个舱室不能被多个屏障覆盖
//     陨石雨仅在到达时刻对冲击位置处的舱室有影响

// 示例 1：
// 输入：time = [1,2,1], position = [6,3,3]
// 输出：5
// 解释： 时刻 1，分别开启编号 3、6 舱室的屏障，能量消耗 2*2 = 4 
//       时刻 2，维持编号 3 舱室的屏障，能量消耗 1 
//       因此，最少需要能量 5

// 示例 2：
// 输入：time = [1,1,1,2,2,3,5], position = [1,2,3,1,2,1,3]
// 输出：9
// 解释： 时刻 1，开启编号 1、2 舱室的联合屏障，能量消耗 3 
//       时刻 1，开启编号 3 舱室的屏障，能量消耗 2 
//       时刻 2，维持编号 1、2 舱室的联合屏障，能量消耗 1 
//       时刻 3，维持编号 1、2 舱室的联合屏障，能量消耗 1 
//       时刻 5，重新开启编号 3 舱室的联合屏障，能量消耗 2 
//       因此，最少需要能量 9

// 提示：
//     1 <= time.length == position.length <= 500
//     1 <= time[i] <= 5
//     0 <= position[i] <= 100

import "fmt"

func defendSpaceCity(time, position []int) int {
    n, m := 100, 1 << 5
    rain := make([]int, n + 1)
    for i, t := range time {
        rain[position[i]] |= 1 << (t - 1)
    }
    union, single := make([]int, m), make([]int, m)
    for i := 1; i < m; i++ {
        lb := i & -i
        j := i ^ lb
        lb2 := j & -j
        if lb == lb2 >> 1 { // 两个时间点相邻
            union[i], single[i]  = union[j] + 1, single[j] + 1 // 递推
        } else {
            // 若 i 的二进制包含 101，对于联合屏障选择继续维持是更优的，
            // 不过下面的 DP 已经枚举了所有的情况，自然会选择更优的方案
            union[i], single[i] = union[j] + 3, single[j] + 2
        }
    }
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, m)
        if i == 0 {
            for j := range dp[0] {
                dp[0][j] = union[j] + single[(m - 1 ^ j) & rain[0]]
            }
        }
    }
    for i := 1; i <= n; i++ {
        for j := range dp[i] {
            dp[i][j] = 1 << 31
            mask := m - 1 ^ j
            for pre := mask; ; pre = (pre - 1) & mask { // 枚举 j 的补集 mask 中的子集 pre
                cost := dp[i-1][pre] + union[j] + single[(mask^pre)&rain[i]]
                dp[i][j] = min(dp[i][j], cost)
                if pre == 0 { break }
            }
        }
    }
    return dp[n][0]
}

func main() {
    // 示例 1：
    // 输入：time = [1,2,1], position = [6,3,3]
    // 输出：5
    // 解释： 时刻 1，分别开启编号 3、6 舱室的屏障，能量消耗 2*2 = 4 
    //       时刻 2，维持编号 3 舱室的屏障，能量消耗 1 
    //       因此，最少需要能量 5
    fmt.Println(defendSpaceCity([]int{1,2,1}, []int{6,3,3})) // 5
    // 示例 2：
    // 输入：time = [1,1,1,2,2,3,5], position = [1,2,3,1,2,1,3]
    // 输出：9
    // 解释： 时刻 1，开启编号 1、2 舱室的联合屏障，能量消耗 3 
    //       时刻 1，开启编号 3 舱室的屏障，能量消耗 2 
    //       时刻 2，维持编号 1、2 舱室的联合屏障，能量消耗 1 
    //       时刻 3，维持编号 1、2 舱室的联合屏障，能量消耗 1 
    //       时刻 5，重新开启编号 3 舱室的联合屏障，能量消耗 2 
    //       因此，最少需要能量 9
    fmt.Println(defendSpaceCity([]int{1,1,1,2,2,3,5}, []int{1,2,3,1,2,1,3})) // 9

    fmt.Println(defendSpaceCity([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 10
    fmt.Println(defendSpaceCity([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(defendSpaceCity([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 10
    fmt.Println(defendSpaceCity([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 10
}