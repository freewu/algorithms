package main

// LCR 186. 文物朝代判断
// 展览馆展出来自 13 个朝代的文物，每排展柜展出 5 个文物。
// 某排文物的摆放情况记录于数组 places，其中 places[i] 表示处于第 i 位文物的所属朝代编号。
// 其中，编号为 0 的朝代表示未知朝代。
// 请判断并返回这排文物的所属朝代编号是否连续（如遇未知朝代可算作连续情况）。

// 示例 1：
// 输入: places = [0, 6, 9, 0, 7]
// 输出: True

// 示例 2：
// 输入: places = [7, 8, 9, 10, 11]
// 输出: True

// 提示：
//     places.length = 5
//     0 <= places[i] <= 13

import "fmt"

func checkDynasty(places []int) bool {
    mn, mx := 1 << 32 - 1, -1 << 32 + 1
    hash := map[int]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range places {
        if v == 0 {
            continue
        } 
        hash[v]++
        if hash[v] > 1 { // 存在重复直接false
            return false
        }
        mn = min(mn, v)
        mx = max(mx, v)
    }
    return mx- mn < 5 // 如果 max - min < 5 ,说明就能顺子
}

func main() {
    // 示例 1：
    // 输入: places = [0, 6, 9, 0, 7]
    // 输出: True
    fmt.Println(checkDynasty([]int{0, 6, 9, 0, 7})) // true
    // 示例 2：
    // 输入: places = [7, 8, 9, 10, 11]
    // 输出: True
    fmt.Println(checkDynasty([]int{7, 8, 9, 10, 11})) // true

    fmt.Println(checkDynasty([]int{6, 8, 9, 10, 11})) // false
}