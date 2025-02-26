package main

// LCP 33. 蓄水
// 给定 N 个无限容量且初始均空的水缸，每个水缸配有一个水桶用来打水，第 i 个水缸配备的水桶容量记作 bucket[i]。
// 小扣有以下两种操作：
//     1. 升级水桶：选择任意一个水桶，使其容量增加为 bucket[i] + 1
//     2. 蓄水：将全部水桶接满水，倒入各自对应的水缸

// 每个水缸对应最低蓄水量记作 vat[i]，返回小扣至少需要多少次操作可以完成所有水缸蓄水要求。

// 注意：实际蓄水量 达到或超过 最低蓄水量，即完成蓄水要求。

// 示例 1：
// 输入：bucket = [1,3], vat = [6,8]
// 输出：4
// 解释： 第 1 次操作升级 bucket[0]； 第 2 ~ 4 次操作均选择蓄水，即可完成蓄水要求。vat1.gif

// 示例 2：
// 输入：bucket = [9,0,1], vat = [0,2,2]
// 输出：3
// 解释： 第 1 次操作均选择升级 bucket[1] 第 2~3 次操作选择蓄水，即可完成蓄水要求。

// 提示：
//     1 <= bucket.length == vat.length <= 100
//     0 <= bucket[i], vat[i] <= 10^4

import "fmt"

func storeWater(bucket []int, vat []int) int {
    res, mx := 1 << 31, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range vat {
        mx = max(mx, v)
    }
    if mx == 0 { return 0 } // 如果最低蓄水量 vat 中所有元素都为 0，说明不需要蓄水
    for x := 1; x <= mx; x++ {
        y := 0
        for i, v := range vat {
            y += max(0, (v + x - 1) / x - bucket[i])
        }
        res = min(res, x + y)
    }
    return res
}

func storeWater1(bucket []int, vat []int) int {
    res, n, mx := 1 << 31, len(bucket), 0
    for _, v := range vat {
        if v > mx {
            mx = v
        }
    }
    if mx == 0 { return 0 } // 特殊情况：如果所有 vat 都是0，不需要蓄水
    for k := 1; k <= 10000; k++ { // 枚举总蓄水次数k，k从1开始
        if k >= res { break }
        upgrades := 0
        // 对于每个水缸，计算需要的升级次数
        for i := 0; i < n; i++ {
            // 计算每个水缸需要的水桶容量
            need := (vat[i] + k - 1) / k  // 向上取整
            if need > bucket[i] {
                upgrades += need - bucket[i]
            }
        }
        // 更新答案：总操作次数 = 升级次数 + 蓄水次数
        total := upgrades + k
        if total < res {
            res = total
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：bucket = [1,3], vat = [6,8]
    // 输出：4
    // 解释： 第 1 次操作升级 bucket[0]； 第 2 ~ 4 次操作均选择蓄水，即可完成蓄水要求。vat1.gif
    fmt.Println(storeWater([]int{1,3}, []int{6,8})) // 4
    // 示例 2：
    // 输入：bucket = [9,0,1], vat = [0,2,2]
    // 输出：3
    // 解释： 第 1 次操作均选择升级 bucket[1] 第 2~3 次操作选择蓄水，即可完成蓄水要求。
    fmt.Println(storeWater([]int{9,0,1}, []int{0,2,2})) // 3

    fmt.Println(storeWater1([]int{1,3}, []int{6,8})) // 4
    fmt.Println(storeWater1([]int{9,0,1}, []int{0,2,2})) // 3
}