package main

// 面试题 16.06. Smallest Difference LCCI
// Given two arrays of integers, compute the pair of values (one value in each array) with the smallest (non-negative) difference. 
// Return the difference.

// Example:
// Input: {1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
// Output:  3, the pair (11, 8)

// Note:
//     1 <= a.length, b.length <= 100000
//     -2147483648 <= a[i], b[i] <= 2147483647
//     The result is in the range [0, 2147483647]

import "fmt"
import "sort"

func smallestDifference(a []int, b []int) int {
    sort.Ints(a)
    sort.Ints(b)
    res, i, j, diff := 1 << 31, 0, 0, 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            diff = b[j] - a[i]
            i++
        } else {
            diff = a[i] - b[j]
            j++
        }
        if diff < res {
            res = diff
        }
    }
    return res
}

func smallestDifference1(a []int, b []int) int {
    sort.Ints(a)
    sort.Ints(b)
    // 排序好之后 两个指针从头扫到尾 找到最小的差
    res, i, j, m, n := 1 << 31, 0, 0, len(a), len(b)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i < m && j < n {
        if a[i] == b[j] { // 相等的情况肯定是最小的差了
            return 0
        } else {
            res = min(res, abs(b[j] - a[i]))
            if a[i] > b[j] {
                j++
            } else {
                i++
            }
        }
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    // Example:
    // Input: {1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
    // Output:  3, the pair (11, 8)
    fmt.Println(smallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8})) // 3

    fmt.Println(smallestDifference([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(smallestDifference1([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8})) // 3
    fmt.Println(smallestDifference1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
}