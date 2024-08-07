package main

// LCS 02. 完成一半题目
// 有 N 位扣友参加了微软与力扣举办了「以扣会友」线下活动。
// 主办方提供了 2*N 道题目，整型数组 questions 中每个数字对应了每道题目所涉及的知识点类型。 
// 若每位扣友选择不同的一题，请返回被选的 N 道题目至少包含多少种知识点类型。

// 示例 1：
// 输入：questions = [2,1,6,2]
// 输出：1
// 解释：有 2 位扣友在 4 道题目中选择 2 题。 可选择完成知识点类型为 2 的题目时，此时仅一种知识点类型 因此至少包含 1 种知识点类型。

// 示例 2：
// 输入：questions = [1,5,1,3,4,5,2,5,3,3,8,6]
// 输出：2
// 解释：有 6 位扣友在 12 道题目中选择题目，需要选择 6 题。 选择完成知识点类型为 3、5 的题目，因此至少包含 2 种知识点类型。

// 提示：
//     questions.length == 2*n
//     2 <= questions.length <= 10^5
//     1 <= questions[i] <= 1000

import "fmt"
import "sort"

func halfQuestions(questions []int) int {
    n := len(questions)
    mp := make(map[int]int)
    for _, v := range questions { // 统计出所有的类型数量
        mp[v]++
    }
    arr := []int{}
    for _, v := range mp {
        arr = append(arr, v)
    }
    sort.Ints(arr)
    res, count := 1, 0
    for i := len(arr) - 1; i > 0; i-- {
        count += arr[i]
        if count >= n / 2 {
            return res
        } else {
            res++
        }
    }
    return res
}

func halfQuestions1(questions []int) int {
    res, cnts := 0, [1001]int{}
    for _,q := range questions{ // 统计出所有的类型数量
        cnts[q]++
    }
    sort.Ints(cnts[:])
    for i,n := 1000, len(questions) / 2; n > 0; i-- {
        res++
        n -= cnts[i]
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：questions = [2,1,6,2]
    // 输出：1
    // 解释：有 2 位扣友在 4 道题目中选择 2 题。 可选择完成知识点类型为 2 的题目时，此时仅一种知识点类型 因此至少包含 1 种知识点类型。
    fmt.Println(halfQuestions([]int{2,1,6,2})) // 1
    // 示例 2：
    // 输入：questions = [1,5,1,3,4,5,2,5,3,3,8,6]
    // 输出：2
    // 解释：有 6 位扣友在 12 道题目中选择题目，需要选择 6 题。 选择完成知识点类型为 3、5 的题目，因此至少包含 2 种知识点类型。
    fmt.Println(halfQuestions([]int{1,5,1,3,4,5,2,5,3,3,8,6})) // 2

    fmt.Println(halfQuestions1([]int{2,1,6,2})) // 1
    fmt.Println(halfQuestions1([]int{1,5,1,3,4,5,2,5,3,3,8,6})) // 2
}

