package main

// LCR 127. 跳跃训练
// 今天的有氧运动训练内容是在一个长条形的平台上跳跃。
// 平台有 num 个小格子，每次可以选择跳 一个格子 或者 两个格子。请返回在训练过程中，学员们共有多少种不同的跳跃方式。

// 结果可能过大，因此结果需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// 示例 1：
// 输入：n = 2
// 输出：2

// 示例 2：
// 输入：n = 5
// 输出：8

// 提示：
//     0 <= n <= 100

import "fmt"
import "time"

// 递归
func trainWays(n int) int {
    if 0 == n { return 1 }
    if n == 1 || n == 2 {
        return n
    }
    return trainWays(n - 1) + trainWays(n - 2)
}

// 利用缓存 备忘录
func trainWays1(n int) int  {
    if 0 == n { return 1 }
    var dfs func(n int,arr []int) int 
    dfs = func(n int,arr []int) int {
        if arr[n] != 0 {
            return arr[n]
        } else {
            arr[n] = (dfs(n -1, arr) + dfs(n - 2, arr)) % 1_000_000_007
            return arr[n]
        }
    }
    if n == 1 || n == 2 {
        return n
    } else {
        arr := make([]int, n + 1)
        arr[1] = 1
        arr[2] = 2
        return dfs(n,arr)
    }
}

// 动态规划法 (利用数组来存储)
func trainWays2(n int) int  {
    if 0 == n { return 1 }
    if n == 1 || n == 2 { return n }
    arr := make([]int, n + 1)
    arr[1] = 1
    arr[2] = 2
    for i := 3 ; i <= n; i++ {
        arr[i] = (arr[i - 1] + arr[i - 2]) % 1_000_000_007
    }
    return arr[n]
}

func main() {
    var start,end int64
    start = time.Now().UnixNano()
    fmt.Println(trainWays(2)) // 2
    fmt.Println(trainWays(3)) // 3
    fmt.Println(trainWays(5)) // 8
    end = time.Now().UnixNano()
    fmt.Printf("trainWays used: %d ns \n",end - start)

    start = time.Now().UnixNano()
    fmt.Println(trainWays1(2)) // 2
    fmt.Println(trainWays1(3)) // 3
    fmt.Println(trainWays1(5)) // 8
    end = time.Now().UnixNano()
    fmt.Printf("trainWays1 used: %d ns \n",end - start)

    start = time.Now().UnixNano()
    fmt.Println(trainWays2(2)) // 2
    fmt.Println(trainWays2(3)) // 3
    fmt.Println(trainWays2(5)) // 8
    end = time.Now().UnixNano()
    fmt.Printf("trainWays2 used: %d ns \n",end - start)
}
