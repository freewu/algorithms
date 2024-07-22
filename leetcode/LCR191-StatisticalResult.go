package main

// LCR 191. 按规则计算统计结果
// 为了深入了解这些生物群体的生态特征，你们进行了大量的实地观察和数据采集。
// 数组 arrayA 记录了各个生物群体数量数据，其中 arrayA[i] 表示第 i 个生物群体的数量。
// 请返回一个数组 arrayB，该数组为基于数组 arrayA 中的数据计算得出的结果，其中 arrayB[i] 表示将第 i 个生物群体的数量从总体中排除后的其他数量的乘积。

// 示例 1：
// 输入：arrayA = [2, 4, 6, 8, 10]
// 输出：[1920, 960, 640, 480, 384]
 
// 提示：
//     所有元素乘积之和不会溢出 32 位整数
//     arrayA.length <= 100000

import "fmt"

func statisticalResult(arrayA []int) []int {
    // 分为左右计算，又或者说上三角下三角
    n := len(arrayA)
    if n == 0 {
        return []int{}
    }
    arrayB := make([]int,n)
    arrayB[0] = 1
    for i := 1; i < n; i++ {
        arrayB[i] = arrayB[i-1] * arrayA[i-1]
    }
    right :=1
    for i := n -1; i >= 0; i-- {
        arrayB[i] = arrayB[i] * right // arrayB[i] 表示将第 i 个生物群体的数量从总体中排除后的其他数量的乘积
        right = right * arrayA[i]
    }
    return arrayB
}

func main() {
    // 示例 1：
    // 输入：arrayA = [2, 4, 6, 8, 10]
    // 输出：[1920, 960, 640, 480, 384]
    fmt.Println(statisticalResult([]int{2, 4, 6, 8, 10})) // [1920, 960, 640, 480, 384]
}