package main

// LCP 38. 守卫城堡
// 城堡守卫游戏的胜利条件为使恶魔无法从出生点到达城堡。
// 游戏地图可视作 2*N 的方格图，记作字符串数组 grid，其中：
//     "." 表示恶魔可随意通行的平地；
//     "#" 表示恶魔不可通过的障碍物，玩家可通过在 平地 上设置障碍物，即将 "." 变为 "#" 以阻挡恶魔前进；
//     "S" 表示恶魔出生点，将有大量的恶魔该点生成，恶魔可向上/向下/向左/向右移动，且无法移动至地图外；
//     "P" 表示瞬移点，移动到 "P" 点的恶魔可被传送至任意一个 "P" 点，也可选择不传送；
//     "C" 表示城堡。

// 然而在游戏中用于建造障碍物的金钱是有限的，请返回玩家最少需要放置几个障碍物才能获得胜利。
// 若无论怎样放置障碍物均无法获胜，请返回 -1。

// 注意：
//     地图上可能有一个或多个出生点
//     地图上有且只有一个城堡

// 示例 1
// 输入：grid = ["S.C.P#P.", ".....#.S"]
// 输出：3
// 解释：至少需要放置三个障碍物
// <img src="https://pic.leetcode-cn.com/1614828255-uuNdNJ-image.png" />

// 示例 2：
// 输入：grid = ["SP#P..P#PC#.S", "..#P..P####.#"]
// 输出：-1
// 解释：无论怎样修筑障碍物，均无法阻挡最左侧出生的恶魔到达城堡位置
// <img src="https://pic.leetcode-cn.com/1614828208-oFlpVs-image.png" />

// 示例 3：
// 输入：grid = ["SP#.C.#PS", "P.#...#.P"]
// 输出：0
// 解释：无需放置障碍物即可获得胜利
// <img src="https://pic.leetcode-cn.com/1614828242-oveClu-image.png" />

// 示例 4：
// 输入：grid = ["CP.#.P.", "...S..S"]
// 输出：4
// 解释：至少需要放置 4 个障碍物，示意图为放置方法之一
// <img src="https://pic.leetcode-cn.com/1614828218-sIAYkb-image.png" />

// 提示：
//     grid.length == 2
//     2 <= grid[0].length == grid[1].length <= 10^4
//     grid[i][j] 仅包含字符 "."、"#"、"C"、"P"、"S"

import "fmt"

func guardCastle(grid []string) int {
    n := len(grid[0]) + 2
    g0, g1, g0t, g1t := make([]byte, n), make([]byte, n), make([]byte, n), make([]byte, n)
    copy(g0[1:], grid[0])
    copy(g1[1:], grid[1])
    g0[0], g0[n - 1], g1[0], g1[n - 1] = '.', '.', '.', '.'
    copy(g0t, g0)
    copy(g1t, g1)
    cal1 := func(i, p int) int {
        for j := p + 1; j <= i; j++ {
            if (g0[j] == '#' && (g1[j] == '#' || g1[j-1] == '#')) ||
               (g1[j] == '#' && (g0[j] == '#' || g0[j-1] == '#')) {
                return 0
            }
        }
        for j := i; j >= p; j-- {
            if g0[j] == '#' {
                if j < i && g1[j+1] == '.' {
                    g1[j+1] = '#'
                }
                return 1
            }
            if g1[j] == '#' {
                if j < i && g0[j+1] == '.' {
                    g0[j+1] = '#'
                }
                return 1
            }
        }
        if g0[i] == '.' {
            g0[i] = '#'
        }
        if g1[i] == '.' {
            g1[i] = '#'
        }
        return 2
    }
    cal := func() int {
        for i := 1; i < n; i++ {
            if g0[i]+g1[i] == 3 || g0[i-1]+g0[i] == 3 || g1[i-1]+g1[i] == 3 {
                return 1e5
            }
        }
        o, p, t := 0, 0, -1
        for i := 1; i < n; i++ {
            if g0[i] == 1 || g1[i] == 1 {
                if t == 2 {
                    o += cal1(i, p)
                }
                p, t = i, 1
            } else if g0[i] == 2 || g1[i] == 2 {
                if t == 1 {
                    o += cal1(i, p)
                }
                p, t = i, 2
            }
        }
        return o
    }
    for i := 1; i < n - 1; i++ {
        switch g0[i] {
        case 'S', 'P':
            g0[i] = 1
        case 'C':
            g0[i] = 2
        }
        switch g1[i] {
        case 'S', 'P':
            g1[i] = 1
        case 'C':
            g1[i] = 2
        }
    }
    res := cal()
    g0, g1 = g0t, g1t
    for i := 1; i < n - 1; i++ {
        switch g0[i] {
        case 'S':
            g0[i] = 1
        case 'C', 'P':
            g0[i] = 2
        }
        switch g1[i] {
        case 'S':
            g1[i] = 1
        case 'C', 'P':
            g1[i] = 2
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res = min(res, cal())
    if res == 1e5 {
        return -1
    }
    return res
}

func main() {
    // 示例 1
    // 输入：grid = ["S.C.P#P.", ".....#.S"]
    // 输出：3
    // 解释：至少需要放置三个障碍物
    // <img src="https://pic.leetcode-cn.com/1614828255-uuNdNJ-image.png" />
    fmt.Println(guardCastle([]string{"S.C.P#P.", ".....#.S"})) // 3
    // 示例 2：
    // 输入：grid = ["SP#P..P#PC#.S", "..#P..P####.#"]
    // 输出：-1
    // 解释：无论怎样修筑障碍物，均无法阻挡最左侧出生的恶魔到达城堡位置
    // <img src="https://pic.leetcode-cn.com/1614828208-oFlpVs-image.png" />
    fmt.Println(guardCastle([]string{"SP#P..P#PC#.S", "..#P..P####.#"})) // -1
    // 示例 3：
    // 输入：grid = ["SP#.C.#PS", "P.#...#.P"]
    // 输出：0
    // 解释：无需放置障碍物即可获得胜利
    // <img src="https://pic.leetcode-cn.com/1614828242-oveClu-image.png" />
    fmt.Println(guardCastle([]string{"SP#.C.#PS", "P.#...#.P"})) // 0
    // 示例 4：
    // 输入：grid = ["CP.#.P.", "...S..S"]
    // 输出：4
    // 解释：至少需要放置 4 个障碍物，示意图为放置方法之一
    // <img src="https://pic.leetcode-cn.com/1614828218-sIAYkb-image.png" />
    fmt.Println(guardCastle([]string{"CP.#.P.", "...S..S"})) // 4
}