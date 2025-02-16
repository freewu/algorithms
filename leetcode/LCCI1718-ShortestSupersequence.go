package main

// 面试题 17.18. Shortest Supersequence LCCI
// You are given two arrays, one shorter (with all distinct elements) and one longer. 
// Find the shortest subarray in the longer array that contains all the elements in the shorter array. 
// The items can appear in any order.

// Return the indexes of the leftmost and the rightmost elements of the array. 
// If there are more than one answer, return the one that has the smallest left index. 
// If there is no answer, return an empty array.

// Example 1:
// Input: big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7], small = [1,5,9]
// Output: [7,10]

// Example 2:
// Input: big = [1,2,3], small = [4]
// Output: []

// Note:
//     big.length <= 100000
//     1 <= small.length <= 100000

import "fmt"

func shortestSeq(big []int, small []int) []int {
    left, right, count, n := 0, 0, 0, len(big)
    win, need :=  make(map[int]int), make(map[int]int)
    for _, v := range small { // 需要的数字
        need[v]++
    }
    res := []int{ 0, 1 << 31 } // 初始化一个最大值
    for right < n {
        win[big[right]]++ // 进窗口
        if win[big[right]] == need[big[right]] { // 如果数字符合要求, cnt++
            count++
        }
        for count == len(need) { // 找到了所有符合要求的数字
            if right - left < res[1] - res[0] { // 更新最小长度（因为从左往右, 所以默认返回的就是左端点最小的一个） 
                res[0],  res[1] = left, right
            }
            if win[big[left]] == need[big[left]] { // 如果出窗口的是符合要求的数字, count--
                count--
            }
            win[big[left]]-- // 出窗口
            left++
        }
        right++
    }
    if res[1] == 1 << 31 { // 如果 res 还是初始值, 证明找不到符合的子数组, 返回空
        return []int{}
    }
    return res
}

func shortestSeq1(big []int, small []int) []int {
    n, m := len(big), len(small)
    if n < m { return nil }
    freq := make(map[int]int)
    for _, v := range small {
        freq[v]++
    }
    res, count, left := []int{}, m, 0
    for right, v := range big {
        if c, ok := freq[v]; ok {
            if c >= 1 {
                count--
            }
            freq[v]--
        }
        for ; count == 0; left++ {
            if len(res) == 0 || right - left < res[1] - res[0] {
                res = []int{ left, right }
            }
            if c, ok := freq[big[left]]; ok {
                if c >= 0 {
                    count++
                }
                freq[big[left]]++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7], small = [1,5,9]
    // Output: [7,10]
    fmt.Println(shortestSeq([]int{7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7}, []int{1,5,9})) // [7,10]
    // Example 2:
    // Input: big = [1,2,3], small = [4]
    // Output: []
    fmt.Println(shortestSeq([]int{1,2,3}, []int{4})) // []

    fmt.Println(shortestSeq([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3})) // [0,2]
    fmt.Println(shortestSeq([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3})) // [6,8]

    fmt.Println(shortestSeq1([]int{7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7}, []int{1,5,9})) // [7,10]
    fmt.Println(shortestSeq1([]int{1,2,3}, []int{4})) // []
    fmt.Println(shortestSeq1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3})) // [0,2]
    fmt.Println(shortestSeq1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3})) // [6,8]
}