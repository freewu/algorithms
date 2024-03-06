package main

// LCS 01. 下载插件
// 小扣打算给自己的 VS code 安装使用插件，初始状态下带宽每分钟可以完成 1 个插件的下载。假定每分钟选择以下两种策略之一:
//         使用当前带宽下载插件
//         将带宽加倍（下载插件数量随之加倍）

// 请返回小扣完成下载 n 个插件最少需要多少分钟。

// 注意：实际的下载的插件数量可以超过 n 个

// 示例 1：
// 输入：n = 2
// 输出：2
// 解释： 以下两个方案，都能实现 2 分钟内下载 2 个插件
//         方案一：第一分钟带宽加倍，带宽可每分钟下载 2 个插件；第二分钟下载 2 个插件
//         方案二：第一分钟下载 1 个插件，第二分钟下载 1 个插件

// 示例 2：
// 输入：n = 4
// 输出：3
// 解释： 
//     最少需要 3 分钟可完成 4 个插件的下载，以下是其中一种方案:   
//         第一分钟带宽加倍，带宽可每分钟下载 2 个插件; 
//         第二分钟下载 2 个插件; 
//         第三分钟下载 2 个插件。

// 提示：
//         1 <= n <= 10^5

import "fmt"
// import "bits"

// func leastMinutes(n int) int {
//     return bits.Len(uint(n-1)) + 1
// }

// 动态规划
func leastMinutes(n int) int {
    dp := make([]int,n+1)
    dp[1] = 1
    for i := 2; i <= n ; i++ {
        dp[i] = dp[(i+1)/2] + 1
    }
    return dp[n]
}

// 贪心
func leastMinutes1(n int) int {
    res := 0
	for i := 1; i < n; i *= 2 {
		res++
	}
	return res + 1
}

func main() {
    fmt.Println(leastMinutes(2)) // 2
    fmt.Println(leastMinutes(4)) // 3
    fmt.Println(leastMinutes(7)) // 4
    fmt.Println(leastMinutes(15)) // 5
    fmt.Println(leastMinutes(63)) // 7
    fmt.Println(leastMinutes(127)) // 8
    fmt.Println(leastMinutes(128)) // 8
    fmt.Println(leastMinutes(255)) // 9
    fmt.Println(leastMinutes(511)) // 10

    fmt.Println(leastMinutes1(2)) // 2
    fmt.Println(leastMinutes1(4)) // 3
    fmt.Println(leastMinutes1(7)) // 4
    fmt.Println(leastMinutes1(15)) // 5
    fmt.Println(leastMinutes1(63)) // 7
    fmt.Println(leastMinutes1(127)) // 8
    fmt.Println(leastMinutes1(128)) // 8
    fmt.Println(leastMinutes1(255)) // 9
    fmt.Println(leastMinutes1(511)) // 10
}