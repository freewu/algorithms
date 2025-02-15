package main

// LCP 12. 小张刷题计划
// 为了提高自己的代码能力，小张制定了 LeetCode 刷题计划，他选中了 LeetCode 题库中的 n 道题，
// 编号从 0 到 n-1，并计划在 m 天内按照题目编号顺序刷完所有的题目（注意，小张不能用多天完成同一题）。

// 在小张刷题计划中，小张需要用 time[i] 的时间完成编号 i 的题目。
// 此外，小张还可以使用场外求助功能，通过询问他的好朋友小杨题目的解法，可以省去该题的做题时间。
// 为了防止“小张刷题计划”变成“小杨刷题计划”，小张每天最多使用一次求助。

// 我们定义 m 天中做题时间最多的一天耗时为 T（小杨完成的题目不计入做题总时间）。
// 请你帮小张求出最小的 T是多少。

// 示例 1：
// 输入：time = [1,2,3,3], m = 2
// 输出：3
// 解释：第一天小张完成前三题，其中第三题找小杨帮忙；第二天完成第四题，并且找小杨帮忙。这样做题时间最多的一天花费了 3 的时间，并且这个值是最小的。

// 示例 2：
// 输入：time = [999,999,999], m = 4
// 输出：0
// 解释：在前三天中，小张每天求助小杨一次，这样他可以在三天内完成所有的题目并不花任何时间。

// 限制：
//     1 <= time.length <= 10^5
//     1 <= time[i] <= 10000
//     1 <= m <= 1000

import "fmt"
import "sort"

func minTime(time []int, m int) int {
    if len(time) <= m { return 0 } // 如果天数大于题目 可以0做完 
    right, left := 0, time[0]
    for _, v := range time {
        right += v
        if left > v {
            left = v
        }
    }
    // 计算当前的值是否能够做完题
    var calc func(key int, day int) bool
    calc = func(key int, day int) bool {
        flag, temp, mx := -1, 0, 0 // flag 标志位 用来判断是否已经使用过一次求助, temp 用来记录当前回答问题耗费的时间, mx 记录今天的所回答问题的最大值
        for i := 0; i < len(time); {
            temp += time[i] // 求今天回答问题耗费的时间
            if time[i] > mx { // 获取今天的最大值
                mx = time[i]
            }
            if temp > key { // 如果今天耗费的时间大于给定的时间
                if flag == -1 { // 是否能够使用一次求助 如果可以 就使用一次求助 并把最大值减掉 这样能够回答问题最大化  贪心
                    temp, flag = temp - mx, -flag // 更改标志位
                } else {
                    day-- // 如果已经使用过了 就将day-1
                    if day <= 0 {  return false } // 如果day使用完成 但是还没有遍历完成 返回false
                    // temp mx flag 重置
                    temp, mx, flag = 0, 0, -flag
                    continue // 跳过这次循环
                }
            }
            i++
        }
        return true	
    }
    for left < right { // 二分查找 left和right分别为 数组的最小值和总和 这样才能够保证把题做完
        mid := (left + right) / 2
        if !calc(mid, m) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func minTime1(time []int, m int) int {
    return sort.Search(1_000_000_000, func(p int) bool {
        sum, mx, day := 0, 0, 1
        for _, v := range time {
            if mx < v {
                mx = v
            }
            if sum + v - mx <= p {
                sum += v
                continue
            } else {
                day++
                sum, mx = v, v
            }
        }
        return day <= m
    })
}

func main() {
    // 示例 1：
    // 输入：time = [1,2,3,3], m = 2
    // 输出：3
    // 解释：第一天小张完成前三题，其中第三题找小杨帮忙；第二天完成第四题，并且找小杨帮忙。这样做题时间最多的一天花费了 3 的时间，并且这个值是最小的。
    fmt.Println(minTime([]int{1,2,3,3}, 2)) // 3
    // 示例 2：
    // 输入：time = [999,999,999], m = 4
    // 输出：0
    // 解释：在前三天中，小张每天求助小杨一次，这样他可以在三天内完成所有的题目并不花任何时间。
    fmt.Println(minTime([]int{999,999,999}, 4)) // 0

    fmt.Println(minTime([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(minTime([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15

    fmt.Println(minTime1([]int{1,2,3,3}, 2)) // 3
    fmt.Println(minTime1([]int{999,999,999}, 4)) // 0
    fmt.Println(minTime1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(minTime1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15
}