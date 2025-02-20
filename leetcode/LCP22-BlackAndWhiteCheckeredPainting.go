package main

// LCP 22. 黑白方格画
// 小扣注意到秋日市集上有一个创作黑白方格画的摊位。
// 摊主给每个顾客提供一个固定在墙上的白色画板，画板不能转动。
// 画板上有n * n的网格。
// 绘画规则为，小扣可以选择任意多行以及任意多列的格子涂成黑色（选择的整行、整列均需涂成黑色），所选行数、列数均可为 0。

// 小扣希望最终的成品上需要有k个黑色格子，请返回小扣共有多少种涂色方案。

// 注意：两个方案中任意一个相同位置的格子颜色不同，就视为不同的方案。

// 示例 1：
// 输入：n = 2, k = 2
// 输出：4
// 解释：一共有四种不同的方案： 第一种方案：涂第一列； 第二种方案：涂第二列； 第三种方案：涂第一行； 第四种方案：涂第二行。

// 示例 2：
// 输入：n = 2, k = 1
// 输出：0
// 解释：不可行，因为第一次涂色至少会涂两个黑格。

// 示例 3：
// 输入：n = 2, k = 4
// 输出：1
// 解释：共有 2*2=4 个格子，仅有一种涂色方案。

// 限制：
//     1 <= n <= 6
//     0 <= k <= n * n

import "fmt"

func paintingPlan(n int, k int) int {
    if k == n * n { return 1 }
    res := 0
    calc := func(base, up int) int {
        a, b := 1, 1
        for i := 0; i < up; i++ {
            a *= base - i
        }
        for i := 1; i <= up; i++ {
            b *= i
        }
        return a / b
    }
    for i := 0; i <= n; i++ {
        for j := 0; j <= n ; j++ {
            if i * n + j * n - i * j == k {
                res += calc(n, i) * calc(n, j)
            }
        } 
    }
    return res
}


func paintingPlan1(n int, k int) int {
    res, sum := 0, n * n
    if sum == k { return 1 }
    calc := func(a int,b int) int {
        res := 1
        for i := 0; i < b; i++ {
            res = res * a / (i + 1)
            a--
        }
        return res
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if sum - (n - i) * (n - j) == k {
                res += calc(n, i) * calc(n, j)
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：n = 2, k = 2
    // 输出：4
    // 解释：一共有四种不同的方案： 第一种方案：涂第一列； 第二种方案：涂第二列； 第三种方案：涂第一行； 第四种方案：涂第二行。
    fmt.Println(paintingPlan(2, 2)) // 4
    // 示例 2：
    // 输入：n = 2, k = 1
    // 输出：0
    // 解释：不可行，因为第一次涂色至少会涂两个黑格。
    fmt.Println(paintingPlan(2, 1)) // 0
    // 示例 3：
    // 输入：n = 2, k = 4
    // 输出：1
    // 解释：共有 2*2=4 个格子，仅有一种涂色方案。
    fmt.Println(paintingPlan(2, 4)) // 1

    fmt.Println(paintingPlan(1, 0)) // 1
    fmt.Println(paintingPlan(1, 1)) // 1
    fmt.Println(paintingPlan(6, 0)) // 1
    fmt.Println(paintingPlan(6, 36)) // 1

    fmt.Println(paintingPlan1(2, 2)) // 4
    fmt.Println(paintingPlan1(2, 1)) // 0
    fmt.Println(paintingPlan1(2, 4)) // 1
    fmt.Println(paintingPlan1(1, 0)) // 1
    fmt.Println(paintingPlan1(1, 1)) // 1
    fmt.Println(paintingPlan1(6, 0)) // 1
    fmt.Println(paintingPlan1(6, 36)) // 1
}