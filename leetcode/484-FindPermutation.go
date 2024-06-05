package main

// 484. Find Permutation
// A permutation perm of n integers of all the integers in the range [1, n] can be represented as a string s of length n - 1 where:
//     s[i] == 'I' if perm[i] < perm[i + 1], and
//     s[i] == 'D' if perm[i] > perm[i + 1].

// Given a string s, reconstruct the lexicographically smallest permutation perm and return it.

// Example 1:
// Input: s = "I"
// Output: [1,2]
// Explanation: [1,2] is the only legal permutation that can represented by s, where the number 1 and 2 construct an increasing relationship.

// Example 2:
// Input: s = "DI"
// Output: [2,1,3]
// Explanation: Both [2,1,3] and [3,1,2] can be represented as "DI", but since we want to find the smallest lexicographical permutation, you should return [2,1,3]
 
// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either 'I' or 'D'.

import "fmt"

func findPermutation(s string) []int {
    res := make([]int, len(s) + 1)
    for i := range res { // 先按照 123... 装填 res
        res[i] = i + 1
    }
    // 翻转区间数组
    reverse := func (a []int, i, j int) {
        for i < j {
            a[i], a[j] = a[j], a[i]
            i++
            j--
        }
    }
    index := 0  // 标记 res 下标的索引
    for i := 0; i < len(s); i++ {
        index = i // 标记第一个为D的下标
        for i < len(s) && s[i] == 'D' {
            i++
        }
        // 此时i是某连续D的最后一个D的下一个元素I
        reverse(res, index, i)  // 同时对应翻转res中下标index到i的元素
    }
    return res
}

// 回溯
func findPermutation1(s string) []int {
    n := len(s) + 1 // 目标的排列长度
    sel := make([]int, n) // 选择列表
    for i := range sel {
         sel[i] = i + 1
    }
    res, path, isVisited := make([]int, 0), make([]int, 0), make([]bool, n + 1) // 1...n 的元素，所以这里长度为 n + 1，标记是否访问过
    // index参数 标记目前排列到的path的下标
    var traverse func(path []int, index int, sel []int, pattern string)
    traverse = func(path []int, index int, sel []int, pattern string) {
        if index > 0 && index < len(sel) && pattern[index - 1] == 'I' && path[index - 1] > path[index] { // 要求上升，但实际是下降
            return 
        }
        if index > 0 && index < len(sel) && pattern[index - 1] == 'D' && path[index - 1] < path[index] { // 要求下降，但实际是上升
            return
        } 
        // 下面的if 判断要放在上面两个if的下面。
        if index == len(sel) - 1 && len(res) == 0 { // 都选择完毕了 && 目前还没有res
            tmp := make([]int, len(path))
            copy(tmp,path)
            res = tmp
            return
        }             
        for i := 0; i < len(sel); i++ {
            if isVisited[sel[i]] {continue} // 已选择过了，排列问题的去重核心
            // 做选择
            isVisited[sel[i]] = true //回溯算法，这里是添加的选择的内容，而不是当前节点的内容和DFS的区别
            path = append(path, sel[i])
            traverse(path, index + 1, sel, pattern) // 向下回溯
            isVisited[sel[i]] = false  // 撤销选择
            path = path[:len(path) - 1]
        } 
    }
    traverse(path, -1, sel, s)
    return res
}

func main() {
    // Example 1:
    // Input: s = "I"
    // Output: [1,2]
    // Explanation: [1,2] is the only legal permutation that can represented by s, where the number 1 and 2 construct an increasing relationship.
    fmt.Println(findPermutation("I")) // [1,2]
    // Example 2:
    // Input: s = "DI"
    // Output: [2,1,3]
    // Explanation: Both [2,1,3] and [3,1,2] can be represented as "DI", but since we want to find the smallest lexicographical permutation, you should return [2,1,3]
    fmt.Println(findPermutation("DI")) //  [2,1,3]

    fmt.Println(findPermutation1("I")) // [1,2]
    fmt.Println(findPermutation1("DI")) //  [2,1,3]
}