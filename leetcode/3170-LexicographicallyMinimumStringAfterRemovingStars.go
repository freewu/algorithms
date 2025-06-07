package main

// 3170. Lexicographically Minimum String After Removing Stars
// You are given a string s. It may contain any number of '*' characters. 
// Your task is to remove all '*' characters.

// While there is a '*', do the following operation:
//     1. Delete the leftmost '*' and the smallest non-'*' character to its left. 
//        If there are several smallest characters, you can delete any of them.

// Return the lexicographically smallest resulting string after removing all '*' characters.

// Example 1:
// Input: s = "aaba*"
// Output: "aab"
// Explanation:
// We should delete one of the 'a' characters with '*'. If we choose s[3], s becomes the lexicographically smallest.

// Example 2:
// Input: s = "abc"
// Output: "abc"
// Explanation:
// There is no '*' in the string.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters and '*'.
//     The input is generated such that it is possible to delete all '*' characters.

import "fmt"
import "sort"

func clearStars(s string) string {
    indexs, deleted := [26][]int{}, make(map[int]bool)
    for i := range s {
        if s[i] != '*' {
            indexs[int(s[i]-'a')] = append(indexs[int(s[i]-'a')], i)
            continue
        }
        for j := 0; j < 26 ; j++ {
            if len(indexs[j]) > 0 {
                pos :=  indexs[j][len(indexs[j]) - 1]
                indexs[j] = indexs[j][:len(indexs[j]) - 1]
                deleted[pos] = true
                break
            }
        }
    }
    res := []byte{}
    for i := range s {
        if !deleted[i] && s[i] != '*' {
            res = append(res, s[i])
        }
    }
    return string(res)
}

func clearStars1(s string) string {
    arr := make([][]int, 26)
    for i, v := range s {
        if v != '*' {
            index := v - 'a'
            arr[index] = append(arr[index], i)
            continue
        }
        for j, _ := range arr { // *
            if len(arr[j]) > 0 {
                arr[j] = arr[j][:len(arr[j]) - 1]
                break
            }
        }
    }
    order := []int{}
    for _, v := range arr {
        order = append(order, v...)
    }
    sort.Ints(order)
    res := make([]byte, len(order))
    for i, v := range order {
        res[i] = s[v]
    }
    return string(res)
}

func clearStars2(s string) string {
    arr, n, pos := []byte(s), len(s), [26][]int{}
    for i := 0; i < n; i++ {
        if arr[i] != '*' {
            c := arr[i] - 'a'
            pos[c] = append(pos[c], i)
            continue
        }
        j := 0
        for ; j < 26; j++ {
            if len(pos[j]) > 0 { break }
        }
        p := pos[j][len(pos[j]) - 1]
        pos[j] = pos[j][:len(pos[j])-1]
        arr[p] = '*'
    }
    res := make([]byte, 0, n)
    for i := 0; i < n; i++ {
        if arr[i] == '*' { continue }
        res = append(res, arr[i])
    }
    return string(res)
}


func clearStars3(s string) string {
    arr := []byte(s) // 将输入字符串转换为字节切片以便修改
    stack := make([][]int, 26) // 创建一个包含26个切片的数组，每个切片存储对应字母的位置索引 stack[0]对应'a'stack[1]对应'b'，依此类推
    for i, c := range s { // 遍历字符串中的每个字符及其索引
        if c != '*' { // 如果当前字符不是星号        
            stack[c - 'a'] = append(stack[c - 'a'], i) // 将该字符的位置追加到对应字母的切片中 c - 'a' 计算字符c在字母表中的位置（0-25）
        } else { // 如果当前字符是星号 
            for j, row := range stack { // 遍历所有字母的切片（从a开始）
                if m := len(row); m > 0 {  // 检查当前字母的切片是否非空
                    arr[row[m - 1]] = '*' // 将该字母最近出现的位置（切片最后一个元素）标记为星号
                    // 从切片中移除最后一个元素（相当于pop操作）
                    stack[j] = row[:m-1]
                    break // 处理完一个星号后立即跳出循环
                }
            }
        }
    }
    // 构建结果字符串：过滤掉所有星号字符
    // 使用原地修改技术来节省空间
    res := arr[:0] // 创建一个长度为0但容量不变的切片
    for _, c := range arr {
        if c != '*' { // 只保留非星号字符
            res = append(res, c)
        }
    }
    return string(res)
}


func main() {
    // Example 1:
    // Input: s = "aaba*"
    // Output: "aab"
    // Explanation:
    // We should delete one of the 'a' characters with '*'. If we choose s[3], s becomes the lexicographically smallest.
    fmt.Println(clearStars("aaba*")) // aab
    // Example 2:
    // Input: s = "abc"
    // Output: "abc"
    // Explanation:
    // There is no '*' in the string.
    fmt.Println(clearStars("abc*")) // abc

    fmt.Println(clearStars("bleufrog*")) // leufrog
    fmt.Println(clearStars("leet*code")) // letcode

    fmt.Println(clearStars1("aaba*")) // aab
    fmt.Println(clearStars1("abc*")) // abc
    fmt.Println(clearStars1("bleufrog*")) // leufrog
    fmt.Println(clearStars1("leet*code")) // letcode

    fmt.Println(clearStars2("aaba*")) // aab
    fmt.Println(clearStars2("abc*")) // abc
    fmt.Println(clearStars2("bleufrog*")) // leufrog
    fmt.Println(clearStars2("leet*code")) // letcode

    fmt.Println(clearStars3("aaba*")) // aab
    fmt.Println(clearStars3("abc*")) // abc
    fmt.Println(clearStars3("bleufrog*")) // leufrog
    fmt.Println(clearStars3("leet*code")) // letcode
}