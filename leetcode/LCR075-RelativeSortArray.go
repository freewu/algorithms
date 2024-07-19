package main

// LCR 075. 数组的相对排序
// 给定两个数组，arr1 和 arr2，
//     arr2 中的元素各不相同
//     arr2 中的每个元素都出现在 arr1 中

// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。
// 未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

// 示例：
// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// 输出：[2,2,2,1,4,3,3,9,6,7,19]
 
// 提示：
//     1 <= arr1.length, arr2.length <= 1000
//     0 <= arr1[i], arr2[i] <= 1000
//     arr2 中的元素 arr2[i] 各不相同
//     arr2 中的每个元素 arr2[i] 都出现在 arr1 中

import "fmt"
import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
    res,mp := []int{}, make(map[int]int) 
    for _, v := range arr1 { // 记录 arr1 出现的次数
        mp[v]++
    }
    for _, v := range arr2 {
        c := mp[v]
        for c > 0 { // 重放 c 次
            res = append(res, v)
            c--
        }
        delete(mp,v)
    }
    remind := []int{}
    for k, c := range mp { // 未在 arr2 中出现过的元素
        for c > 0 { // 重放 c 次
            remind = append(remind, k)
            c--
        }
    }
    sort.Ints(remind) // 未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾
    res = append(res,remind...)
    return res
}

func relativeSortArray1(arr1 []int, arr2 []int) []int {
    m := make(map[int]int)
    for i := 0;i < len(arr1);i++{
        m[arr1[i]]++
    }
    res, pos := make([]int,len(arr1)), 0
    for i := 0;i < len(arr2);i++{
        for j := 0;j < m[arr2[i]];j++{
            res[pos] = arr2[i]
            pos++
        }
        delete(m,arr2[i])
    }
    keys := make([]int,0)
    for k ,_ := range m {
        keys = append(keys,k)
    }
    sort.Ints(keys)
    for i := 0;i < len(keys);i++{
        for j := 0;j < m[keys[i]];j++ {
            res[pos] = keys[i]
            pos++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
    // Output: [2,2,2,1,4,3,3,9,6,7,19]
    fmt.Println(relativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6})) // [2,2,2,1,4,3,3,9,6,7,19]
    // Example 2:
    // Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
    // Output: [22,28,8,6,17,44]
    fmt.Println(relativeSortArray([]int{28,6,22,8,44,17}, []int{22,28,8,6})) // [22,28,8,6,17,44]

    fmt.Println(relativeSortArray1([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6})) // [2,2,2,1,4,3,3,9,6,7,19]
    fmt.Println(relativeSortArray1([]int{28,6,22,8,44,17}, []int{22,28,8,6})) // [22,28,8,6,17,44]
}