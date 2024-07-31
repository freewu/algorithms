package main

// LCP 40. 心算挑战
// 「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，
// 若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。 
// 给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 
// 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。

// 示例 1：
// 输入：cards = [1,2,8,9], cnt = 3
// 输出：18
// 解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。

// 示例 2：
// 输入：cards = [3,3,1], cnt = 1
// 输出：0
// 解释：不存在获取有效得分的卡牌方案。

// 提示：
//     1 <= cnt <= cards.length <= 10^5
//     1 <= cards[i] <= 1000

import "fmt"
import "sort"

// func maxmiumScore(cards []int, cnt int) int {
//     res := 0
//     sort.Ints(cards)
//     for i := len(cards) - 1; i >= 0; i-- {
//         res += cards[i]
//         cnt--
//         if cnt == 0 {
//             if res % 2 == 0 { // 偶数直接返回
//                 return res
//             }
//             res -= cards[i]
//             cnt++
//         }
//     }
//     return 0
// }

func maxmiumScore(cards []int, cnt int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(cards)))
    res, tmp, odd, even := 0, 0, -1, -1
    for i := 0; i < cnt; i++ {
        tmp += cards[i]
        if cards[i] % 2 == 1 {
            odd = cards[i]
        } else {
            even = cards[i]
        }
    }
    if tmp % 2 == 0 {
        return tmp
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := cnt; i < len(cards); i++ {
        if cards[i] % 2 == 1 {
            if even != -1 { // 在数组前面找到一个最大的偶数与后 cnt 个数中最小的奇数进行替换
                res = max(res, tmp - even + cards[i])
            }
        } else {
            if odd != -1 { // 在数组前面找到一个最大的奇数与后 cnt 个数中最小的偶数进行替换
                res = max(res, tmp - odd + cards[i])
            }
        }
    }
    return res
}

// hash
func maxmiumScore1(cards []int, cnt int) int {
    val := make([]int, 1005)
    for _, card := range cards {
        val[card]++
    }
    res, tmp, odd, even, ed := 0, 0, -1, -1, -1
    for i := 1000; i >= 1; i-- {
        if val[i] == 0 {
            continue
        }
        if val[i] > cnt {
            tmp += cnt * i
            cnt = 0
        } else {
            tmp += val[i] * i
            cnt -= val[i]
            val[i] = 0
        }
        if i%2 == 1 {
            odd = i
        } else {
            even = i
        }
        if cnt == 0 {
            if val[i] > 0 {
                ed = i
            } else {
                ed = i - 1
            }
            break
        }
    }
    if tmp % 2 == 0 {
        return tmp
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := ed; i >= 1; i-- {
        if val[i] == 0 {
            continue
        }
        if i % 2 == 1 {
            if even != -1 {
                res = max(res, tmp - even + i)
            }
        } else {
            if odd != -1 {
                res = max(res, tmp - odd + i)
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：cards = [1,2,8,9], cnt = 3
    // 输出：18
    // 解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。
    fmt.Println(maxmiumScore([]int{1,2,8,9}, 3)) // 18
    // 示例 2：
    // 输入：cards = [3,3,1], cnt = 1
    // 输出：0
    // 解释：不存在获取有效得分的卡牌方案。
    fmt.Println(maxmiumScore([]int{3,3,1}, 1)) // 0

    fmt.Println(maxmiumScore([]int{1,3,4,5}, 4)) // 0
    fmt.Println(maxmiumScore([]int{9,5,9,1,6,10,3,4,5,1}, 2)) // 18

    fmt.Println(maxmiumScore1([]int{1,2,8,9}, 3)) // 18
    fmt.Println(maxmiumScore1([]int{3,3,1}, 1)) // 0
    fmt.Println(maxmiumScore1([]int{1,3,4,5}, 4)) // 0
    fmt.Println(maxmiumScore1([]int{9,5,9,1,6,10,3,4,5,1}, 2)) // 18
}
