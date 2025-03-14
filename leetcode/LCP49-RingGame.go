package main

// LCP 49. 环形闯关游戏
// 「力扣挑战赛」中有一个由 N 个关卡组成的环形闯关游戏，关卡编号为 0~N-1，编号 0 的关卡和编号 N-1 的关卡相邻。
// 每个关卡均有积分要求，challenge[i] 表示挑战编号 i 的关卡最少需要拥有的积分。
// <img src="https://pic.leetcode-cn.com/1630392170-ucncVS-%E5%9B%BE%E7%89%87.png" />

// 小扣想要挑战关卡，闯关具体规则如下：
//     1. 初始小扣可以指定其中一个关卡为「开启」状态，其余关卡将处于「未开启」状态。
//     2. 小扣可以挑战处于「开启」状态且满足最少积分要求的关卡，若小扣挑战该关卡前积分为 score，挑战结束后，积分将增长为 score|challenge[i]（即位运算中的 "OR" 运算）
//     3. 在挑战某个关卡后，该关卡两侧相邻的关卡将会开启（若之前未开启）

// 请帮助小扣进行计算，初始最少需要多少积分，可以挑战 环形闯关游戏 的所有关卡。

// 示例1：
// 输入：challenge = [5,4,6,2,7]
// 输出：4
// 解释： 初始选择编号 3 的关卡开启，积分为 4 挑战编号 3 的关卡，积分变为 4∣2=6，开启 2、4 处的关卡 挑战编号 2 的关卡，积分变为 6∣6=6，开启 1 处的关卡 挑战编号 1 的关卡，积分变为 6∣4=6，开启 0 处的关卡 挑战编号 0 的关卡，积分变为 6∣5=7 挑战编号 4 的关卡，顺利完成全部的关卡

// 示例2：
// 输入：challenge = [12,7,11,3,9]
// 输出：8
// 解释： 初始选择编号 3 的关卡开启，积分为 8 挑战编号 3 的关卡，积分变为 8∣3=11，开启 2、4 处的关卡 挑战编号 2 的关卡，积分变为 11∣11=11，开启 1 处的关卡 挑战编号 4 的关卡，积分变为 11∣9=11，开启 0 处的关卡 挑战编号 1 的关卡，积分变为 11∣7=15 挑战编号 0 的关卡，顺利完成全部的关卡

// 示例3：
// 输入：challenge = [1,1,1]
// 输出：1

// 提示：
//     1 <= challenge.length <= 5*10^4
//     1 <= challenge[i] <= 10^14

import "fmt"

func ringGame(challenge []int64) int64 {
    if len(challenge) == 0 { return 0 }
    if len(challenge) == 1 { return challenge[0] }
    res, b, mx := int64(0), 0, challenge[0]
    for _, c := range challenge {
        if c > mx {
            mx = c
        }
    }
    for b = 62; b >= 0; b-- { // 找最高位
        t := int64((1 << b))
        if  (t & mx) > 0 {
            res = int64(t) // 设置最高位
            break
        }
    }
    b--
    check := func(challenge []int64, m int64) bool {
        n := len(challenge)
        for s := 0; s < n; s ++ {
            if challenge[s] > m { continue }
            score := challenge[s] | m
            // 往右挑战
            finished := true
            for j := (s + 1) % n; j != s; j = (j + 1) % n {
                if score < challenge[j] {
                    finished = false
                    break
                }
                score = score | challenge[j]
            }
            if finished { return finished }
            // 往左
            finished = true
            for j := (s - 1 + n) % n; j != s; j=(j - 1 + n) % n{
                if score < challenge[j] {
                    finished = false
                    break
                }
                score = score | challenge[j]
            }
            if finished { return finished }
        }
        return false
    }
    for b >= 0 {
        if !check(challenge, res | ((1 << b) - 1)) {
            res |= (1 << b)
        }
        b--
    }
    return res
}

func main() {
    // 示例1：
    // 输入：challenge = [5,4,6,2,7]
    // 输出：4
    // 解释： 初始选择编号 3 的关卡开启，积分为 4 挑战编号 3 的关卡，积分变为 4∣2=6，开启 2、4 处的关卡 挑战编号 2 的关卡，积分变为 6∣6=6，开启 1 处的关卡 挑战编号 1 的关卡，积分变为 6∣4=6，开启 0 处的关卡 挑战编号 0 的关卡，积分变为 6∣5=7 挑战编号 4 的关卡，顺利完成全部的关卡
    fmt.Println(ringGame([]int64{5,4,6,2,7})) // 4
    // 示例2：
    // 输入：challenge = [12,7,11,3,9]
    // 输出：8
    // 解释： 初始选择编号 3 的关卡开启，积分为 8 挑战编号 3 的关卡，积分变为 8∣3=11，开启 2、4 处的关卡 挑战编号 2 的关卡，积分变为 11∣11=11，开启 1 处的关卡 挑战编号 4 的关卡，积分变为 11∣9=11，开启 0 处的关卡 挑战编号 1 的关卡，积分变为 11∣7=15 挑战编号 0 的关卡，顺利完成全部的关卡
    fmt.Println(ringGame([]int64{12,7,11,3,9})) // 8
    // 示例3：
    // 输入：challenge = [1,1,1]
    // 输出：1
    fmt.Println(ringGame([]int64{1,1,1})) // 1

    fmt.Println(ringGame([]int64{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(ringGame([]int64{9,8,7,6,5,4,3,2,1})) // 8
}