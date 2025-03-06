package main

// 面试题 16.11. Diving Board LCCI
// You are building a diving board by placing a bunch of planks of wood end-to-end. 
// There are two types of planks, one of length shorter and one of length longer. 
// You must use exactly K planks of wood. 
// Write a method to generate all possible lengths for the diving board.

// return all lengths in non-decreasing order.

// Example:
// Input: 
// shorter = 1
// longer = 2
// k = 3
// Output:  {3,4,5,6}

// Note:
//     0 < shorter <= longer
//     0 <= k <= 100000

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {
    if k == 0 { return nil }
    if shorter == longer { // 重复，不管怎么选长度始终一样
        return []int{ shorter * k }
    }
    start := shorter * k // 数列起点
    end := longer * k // 数列终点
    diff := longer - shorter // 公差
    res := make([]int, (end - start) / diff + 1)
    if start > end { // shorter > longer 的情况，不清楚是否有这样的测试用例，但是测试的时候是可运行的。
        for i := start; i >= end; i += diff {
            res[(i - start) / end] = i
        }
    } else {
        for i := start; i <= end; i += diff {
            res[(i - start) / diff] = i
        }
    }
    return res
}

func main() {
    // Example:
    // Input: 
    // shorter = 1
    // longer = 2
    // k = 3
    // Output:  {3,4,5,6}
    fmt.Println(divingBoard(1, 2, 3)) // {3,4,5,6}

    fmt.Println(divingBoard(0, 0, 0)) // []
    fmt.Println(divingBoard(100000, 100000, 100000)) // [10000000000]
    // fmt.Println(divingBoard(0, 100000, 100000))
    fmt.Println(divingBoard(0, 0, 100000)) // [0]
    fmt.Println(divingBoard(0, 100000, 0)) // []
}