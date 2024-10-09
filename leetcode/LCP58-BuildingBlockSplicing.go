package main

// LCP 58. 积木拼接
// 欢迎各位勇者来到力扣城，本次试炼主题为「积木拼接」。 
// 勇者面前有 6 片积木（厚度均为 1），每片积木的形状记录于二维字符串数组 shapes 中，shapes[i] 表示第 i 片积木，其中 1 表示积木对应位置无空缺，0 表示积木对应位置有空缺。 
// 例如 ["010","111","010"] 对应积木形状为
// <img src="https://pic.leetcode-cn.com/1616125620-nXMCxX-image.png" />

// 拼接积木的规则如下：
//     1. 积木片可以旋转、翻面
//     2. 积木片边缘必须完全吻合才能拼接在一起
//     3. 每片积木片 shapes[i] 的中心点在拼接时必须处于正方体对应面的中心点

// 例如 3*3、4*4 的积木片的中心点如图所示（红色点）：
// <img src="https://pic.leetcode-cn.com/1650509082-wObiEp-middle_img_v2_c2d91eb5-9beb-4c06-9726-f7dae149d86g.png" />

// 请返回这 6 片积木能否拼接成一个严丝合缝的正方体且每片积木正好对应正方体的一个面。

// 注意：
//     输入确保每片积木均无空心情况（即输入数据保证对于大小 N*N 的 shapes[i]，内部的 (N-2)*(N-2) 的区域必然均为 1）
//     输入确保每片积木的所有 1 位置均连通

// 示例 1：
// 输入：shapes = [["000","110","000"],["110","011","000"],["110","011","110"],["000","010","111"],["011","111","011"],["011","010","000"]]
// 输出：true
// 解释：
// <img src="https://pic.leetcode-cn.com/1616125823-hkXAeN-cube.gif" />

// 示例 2：
// 输入：shapes = [["101","111","000"],["000","010","111"],["010","011","000"],["010","111","010"],["101","111","010"],["000","010","011"]]
// 输出：false
// 解释： 由于每片积木片的中心点在拼接时必须处于正方体对应面的中心点，积木片 ["010","011","000"] 不能作为 ["100","110","000"] 使用，因此无法构成正方体

// 提示：
//     shapes.length == 6
//     shapes[i].length == shapes[j].length
//     shapes[i].length == shapes[i][j].length
//     3 <= shapes[i].length <= 10

import "fmt"

func composeCube(shapes [][]string) bool {
    encode := func(a [][]byte) (res [4][2]int) { // 每条边压缩成二进制
        n := len(a)
        for i, b := range a[0] {
            res[0][0] |= int(b&1) << i // 正向
            res[0][1] |= int(b&1) << (n - 1 - i) // 反向
            res[2][0] |= int(a[n-1][i]&1) << i
            res[2][1] |= int(a[n-1][i]&1) << (n - 1 - i)
        }
        for i, r := range a {
            res[1][0] |= int(r[n-1]&1) << i
            res[1][1] |= int(r[n-1]&1) << (n - 1 - i)
            res[3][0] |= int(r[0]&1) << i
            res[3][1] |= int(r[0]&1) << (n - 1 - i)
        }
        return
    }
    rotate := func(a [][]byte) [][]byte { // 顺时针旋转矩阵 90°
        n, m := len(a), len(a[0])
        b := make([][]byte, m)
        for i := range b {
            b[i] = make([]byte, n)
        }
        for i, r := range a {
            for j, v := range r {
                b[j][n-1-i] = v
            }
        }
        return b
    }
    n := len(shapes[0])
    a := [6][8][4][2]int{} // [积木][旋转+翻转][边][0-正向/1-反向]
    for i, shape := range shapes {
        t := make([][]byte, n)
        for j, s := range shape {
            t[j] = []byte(s)
        }
        for j := 0; j < 4; j++ {
            a[i][j] = encode(t)
            t = rotate(t)
        }
        for _, r := range t {
            for j := 0; j < n/2; j++ {
                r[j], r[n-1-j] = r[n-1-j], r[j]
            }
        }
        for j := 4; j < 8; j++ {
            a[i][j] = encode(t)
            t = rotate(t)
        }
    }
    // 判断两条边是否恰好重叠（除了顶角）
    MASK := 1<<(n-1) - 2
    check := func(v, w int) bool { return v&w == 0 && (v|w)&MASK == MASK }
    type pair struct{ who, rot int }
    fill := [6]pair{} // 枚举每个积木以什么旋转/翻转姿势放在哪个面（0-顶面，1234-侧面，5-底面）
    vis := 0
    var dfs func(int) bool
    dfs = func(p int) bool { // 当前考虑的面
        if p == 6 { return true }
        for cur := 1; cur < 6; cur++ { // 枚举 6 个积木（固定第一个积木放在顶面）
            if vis>>cur&1 > 0 { continue }
            vis ^= 1 << cur
            for rot := 0; rot < 8; rot++ { // 枚举 8 种旋转+翻转的情况
                switch p {
                case 1:
                    // 1 和 0 是否有冲突
                    if !check(a[cur][rot][0][0], a[0][0][2][0]) {
                        continue
                    }
                case 2:
                    // 2 和 0 1 是否有冲突
                    w, r := fill[p-1].who, fill[p-1].rot
                    if !check(a[cur][rot][0][0], a[0][0][1][1]) || // 边是否冲突
                        !check(a[cur][rot][3][0], a[w][r][1][0]) ||
                        a[0][0][2][1]&1 == 0 && a[cur][rot][0][0]&1 == 0 && a[w][r][0][1]&1 == 0 { // 角是否冲突
                        continue
                    }
                case 3:
                    // 3 和 0 2 是否有冲突
                    w, r := fill[p-1].who, fill[p-1].rot
                    if !check(a[cur][rot][0][0], a[0][0][0][1]) ||
                        !check(a[cur][rot][3][0], a[w][r][1][0]) ||
                        a[0][0][1][0]&1 == 0 && a[cur][rot][0][0]&1 == 0 && a[w][r][0][1]&1 == 0 {
                        continue
                    }
                case 4:
                    // 4 和 0 1 3 是否有冲突
                    w, r := fill[p-1].who, fill[p-1].rot
                    w1, r1 := fill[1].who, fill[1].rot
                    if !check(a[cur][rot][0][0], a[0][0][3][0]) ||
                        !check(a[cur][rot][3][0], a[w][r][1][0]) ||
                        !check(a[cur][rot][1][0], a[w1][r1][3][0]) ||
                        a[0][0][3][0]&1 == 0 && a[cur][rot][0][0]&1 == 0 && a[w][r][0][1]&1 == 0 ||
                        a[0][0][2][0]&1 == 0 && a[cur][rot][0][1]&1 == 0 && a[w1][r1][0][0]&1 == 0 {
                        continue
                    }
                default:
                    // 5 和 1 2 3 4 是否有冲突
                    w1, r1 := fill[1].who, fill[1].rot
                    w2, r2 := fill[2].who, fill[2].rot
                    w3, r3 := fill[3].who, fill[3].rot
                    w4, r4 := fill[4].who, fill[4].rot
                    if !check(a[cur][rot][0][0], a[w1][r1][2][0]) ||
                        !check(a[cur][rot][1][0], a[w2][r2][2][0]) ||
                        !check(a[cur][rot][2][1], a[w3][r3][2][0]) ||
                        !check(a[cur][rot][3][1], a[w4][r4][2][0]) ||
                        a[cur][rot][0][1]&1 == 0 && a[w1][r1][2][1]&1 == 0 && a[w2][r2][2][0]&1 == 0 ||
                        a[cur][rot][1][1]&1 == 0 && a[w2][r2][2][1]&1 == 0 && a[w3][r3][2][0]&1 == 0 ||
                        a[cur][rot][2][0]&1 == 0 && a[w3][r3][2][1]&1 == 0 && a[w4][r4][2][0]&1 == 0 ||
                        a[cur][rot][0][0]&1 == 0 && a[w4][r4][2][1]&1 == 0 && a[w1][r1][2][0]&1 == 0 {
                        continue
                    }
                }
                fill[p] = pair{cur, rot}
                if dfs(p + 1) { return true }
            }
            vis ^= 1 << cur
        }
        return false
    }
    return dfs(1)
}

func main() {
    // 示例 1：
    // 输入：shapes = [["000","110","000"],["110","011","000"],["110","011","110"],["000","010","111"],["011","111","011"],["011","010","000"]]
    // 输出：true
    // 解释：
    // <img src="https://pic.leetcode-cn.com/1616125823-hkXAeN-cube.gif" />
    shapes1 := [][]string{
        {"000","110","000"},
        {"110","011","000"},
        {"110","011","110"},
        {"000","010","111"},
        {"011","111","011"},
        {"011","010","000"},
    }
    fmt.Println(composeCube(shapes1)) // true
    // 示例 2：
    // 输入：shapes = [["101","111","000"],["000","010","111"],["010","011","000"],["010","111","010"],["101","111","010"],["000","010","011"]]
    // 输出：false
    // 解释： 由于每片积木片的中心点在拼接时必须处于正方体对应面的中心点，积木片 ["010","011","000"] 不能作为 ["100","110","000"] 使用，因此无法构成正方体
    shapes2 := [][]string{
        {"101","111","000"},
        {"000","010","111"},
        {"010","011","000"},
        {"010","111","010"},
        {"101","111","010"},
        {"000","010","011"},
    }
    fmt.Println(composeCube(shapes2)) // false
}