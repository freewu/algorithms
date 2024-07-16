package main

// LCR 178. 训练计划 VI
// 教学过程中，教练示范一次，学员跟做三次。
// 该过程被混乱剪辑后，记录于数组 actions，其中 actions[i] 表示做出该动作的人员编号。
// 请返回教练的编号。

// 示例 1：
// 输入：actions = [5, 7, 5, 5]
// 输出：7

// 示例 2：
// 输入：actions = [12, 1, 6, 12, 6, 12, 6]
// 输出：1

// 提示：
//     1 <= actions.length <= 10000
//     1 <= actions[i] < 2^31

import "fmt"

func trainingPlan(actions []int) int {
    mp := make(map[int]int)
    for _, v := range actions {
        mp[v]++
    }
    for i, c := range mp {
        if c == 1 {
            return i
        }
    }
    return -1
}

// 位运算
func trainingPlan1(actions []int) int {
    res, counts := 0, make([]int,32)
    for _,action := range actions {
        for i := 0; i < 32 ; i++ {
            counts[i] += action & 1
            action = action >> 1
        }
    }
    for i := 31; i >= 0;i-- {
        res = res << 1
        res = (res | (counts[i] % 3))
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：actions = [5, 7, 5, 5]
    // 输出：7
    fmt.Println(trainingPlan([]int{5, 7, 5, 5})) // 7
    // 示例 2：
    // 输入：actions = [12, 1, 6, 12, 6, 12, 6]
    // 输出：1
    fmt.Println(trainingPlan([]int{12, 1, 6, 12, 6, 12, 6})) // 1

    fmt.Println(trainingPlan1([]int{5, 7, 5, 5})) // 7
    fmt.Println(trainingPlan1([]int{12, 1, 6, 12, 6, 12, 6})) // 1
}