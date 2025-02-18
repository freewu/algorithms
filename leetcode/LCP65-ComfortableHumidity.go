package main

// LCP 65. 舒适的湿度
// 力扣嘉年华为了确保更舒适的游览环境条件，在会场的各处设置了湿度调节装置，这些调节装置受控于总控室中的一台控制器。 
// 控制器中已经预设了一些调节指令，整数数组operate[i]表示第i条指令增加空气湿度的大小。
// 现在你可以将任意数量的指令修改为降低湿度（变化的数值不变），以确保湿度尽可能的适宜：
//     1. 控制器会选择一段连续的指令，从而进行湿度调节的操作；
//     2. 这段指令最终对湿度影响的绝对值，即为当前操作的「不适宜度」
//     3. 在控制器所有可能的操作中，最大的「不适宜度」即为「整体不适宜度」

// 请返回在所有修改指令的方案中，可以得到的最小「整体不适宜度」。

// 示例 1：
// 输入：operate = [5,3,7]
// 输出：8
// 解释：对于方案2的[5,3,-7]操作指令[5],[3],[-7]的「不适宜度」分别为5,3,7操作指令[5,3],[3,-7]的「不适宜度」分别为8,4操作指令[5,3,-7]的「不适宜度」为1， 因此对于方案[5,3,-7]的「整体不适宜度」为8，其余方案的「整体不适宜度」均不小于8，如下表所示：
// <img src="https://pic.leetcode-cn.com/1663902759-dgDCxn-image.png" />

// 示例 2：
// 输入：operate = [20,10]
// 输出：20

// 提示：
//     1 <= operate.length <= 1000
//     1 <= operate[i] <= 1000

import "fmt"

func unSuitability(operate []int) int {
    mx, inf := 0, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range operate {
        mx = max(mx, v)
    }
    mx *= 2
    pre := make([]int, mx + 1)
    for i := range pre {
        pre[i] = inf
    }
    pre[0] = 0
    f := make([]int, mx + 1)
    for _, v := range operate {
        for i := range f {
            f[i] = inf
        }
        for j, dis := range pre {
            if pre[j] == inf { // 无效的长度（无法组成）
                continue
            }
            if j + v <= mx {
                f[j + v] = min(f[j + v], max(dis, j + v))
            }
            if j >= v {
                f[j - v] = min(f[j - v], dis)
            } else {
                f[0] = min(f[0], dis - j + v)
            }
        }
        pre, f = f, pre
    }
    res := inf
    for _, v := range pre {
        res = min(res, v)
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：operate = [5,3,7]
    // 输出：8
    // 解释：对于方案2的[5,3,-7]操作指令[5],[3],[-7]的「不适宜度」分别为5,3,7操作指令[5,3],[3,-7]的「不适宜度」分别为8,4操作指令[5,3,-7]的「不适宜度」为1， 因此对于方案[5,3,-7]的「整体不适宜度」为8，其余方案的「整体不适宜度」均不小于8，如下表所示：
    // <img src="https://pic.leetcode-cn.com/1663902759-dgDCxn-image.png" />
    fmt.Println(unSuitability([]int{5,3,7})) // 8
    // 示例 2：
    // 输入：operate = [20,10]
    // 输出：20
    fmt.Println(unSuitability([]int{20,10})) // 20

    fmt.Println(unSuitability([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(unSuitability([]int{9,8,7,6,5,4,3,2,1})) // 9
}