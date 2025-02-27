package main

// 面试题 17.17. Multi Search LCCI
// Given a string band an array of smaller strings T, design a method to search b for each small string in T. 
// Output positions of all strings in smalls that appear in big, where positions[i] is all positions of smalls[i].

// Example:
// Input: 
// big = "mississippi"
// smalls = ["is","ppi","hi","sis","i","ssippi"]
// Output:  [[1,4],[8],[],[3],[1,4,7,10],[5]]

// Note:
//     0 <= len(big) <= 1000
//     0 <= len(smalls[i]) <= 1000
//     The total number of characters in smalls will not exceed 100000.
//     No duplicated strings in smalls.
//     All characters are lowercase letters.

import "fmt"

func multiSearch(big string, smalls []string) [][]int {
    res := [][]int{}
    m := len(big)
    for _, v := range smalls {
        arr := []int{}
        i, j, n := 0, 0, len(v)
        for i < m && j < n {
            // 如果big和small当前字符相等,则继续尝试下一个
            if big[i] == v[j] && j < n-1 {
                i++
                j++
                continue
            }
            // 如果一直到small末尾都相等,则找到一个完整匹配,记录下来.
            if big[i] == v[j] && j == n-1 {
                arr = append(arr, i - n + 1)
            }
            // 到了这里有两种可能:要么刚刚一次匹配成功,要么匹配失败;但是都要确定下一次匹配的开始位置.
            // 需要回溯:big回溯到上一次匹配起始位置的右边第一个与small首字符相等的位置,small回溯到0.
            k := i - j + 1
            for k < i && big[k] != v[0] {
                k++
            }
            i, j = k, 0
        }
        res = append(res, arr)
    }
    return res
}

func main() {
    // Example:
    // Input: 
    // big = "mississippi"
    // smalls = ["is","ppi","hi","sis","i","ssippi"]
    // Output:  [[1,4],[8],[],[3],[1,4,7,10],[5]]
    fmt.Println(multiSearch("mississippi", []string{"is","ppi","hi","sis","i","ssippi"})) // [[1,4],[8],[],[3],[1,4,7,10],[5]]
}