package main

// LCR 139. 训练计划 I
// 教练使用整数数组 actions 记录一系列核心肌群训练项目编号。
// 为增强训练趣味性，需要将所有奇数编号训练项目调整至偶数编号训练项目之前。请将调整后的训练项目编号以 数组 形式返回。


// 示例 1：
// 输入：actions = [1,2,3,4,5]
// 输出：[1,3,5,2,4] 
// 解释：为正确答案之一

// 提示：
//     0 <= actions.length <= 50000
//     0 <= actions[i] <= 10000

import "fmt"

func trainingPlan(actions []int) []int {
    odds, evens := []int{}, []int{}
    for _, v := range actions {
        if v % 2 == 0 {
            evens = append(evens, v)
        } else {
            odds = append(odds, v)
        }
    }
    for _, v := range evens {
        odds = append(odds, v)
    }
    return odds
}

// 双指针
func trainingPlan1(actions []int) []int {
    left, right := 0, len(actions) - 1
    for left <= right {
        for left <= right && actions[left] % 2 == 1 {
            left++
        }
        for left <= right && actions[right] % 2 == 0 {
            right--
        }
        if left <= right {
            actions[left], actions[right] = actions[right], actions[left]
            left++
            right--
        }
    }
    return actions
}

func main() {
// 示例 1：
// 输入：actions = [1,2,3,4,5]
// 输出：[1,3,5,2,4] 
// 解释：为正确答案之一
fmt.Println(trainingPlan([]int{1,2,3,4,5})) // [1,3,5,2,4] 

fmt.Println(trainingPlan1([]int{1,2,3,4,5})) // [1,3,5,2,4] 
}