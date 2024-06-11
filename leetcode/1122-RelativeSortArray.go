package main

// 1122. Relative Sort Array
// Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
// Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2. 
// Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.

// Example 1:
// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// Output: [2,2,2,1,4,3,3,9,6,7,19]

// Example 2:
// Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
// Output: [22,28,8,6,17,44]
 
// Constraints:
//     1 <= arr1.length, arr2.length <= 1000
//     0 <= arr1[i], arr2[i] <= 1000
//     All the elements of arr2 are distinct.
//     Each arr2[i] is in arr1.

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

func main() {
    // Example 1:
    // Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
    // Output: [2,2,2,1,4,3,3,9,6,7,19]
    fmt.Println(relativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6})) // [2,2,2,1,4,3,3,9,6,7,19]
    // Example 2:
    // Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
    // Output: [22,28,8,6,17,44]
    fmt.Println(relativeSortArray([]int{28,6,22,8,44,17}, []int{22,28,8,6})) // [22,28,8,6,17,44]
}