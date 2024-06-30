package main

// LCR 128. 库存管理 I
// 仓库管理员以数组 stock 形式记录商品库存表。
// stock[i] 表示商品 id，可能存在重复。
// 原库存表按商品 id 升序排列。现因突发情况需要进行商品紧急调拨，管理员将这批商品 id 提前依次整理至库存表最后。
// 请你找到并返回库存表中编号的 最小的元素 以便及时记录本次调拨。

// 示例 1：
// 输入：stock = [4,5,8,3,4]
// 输出：3

// 示例 2：
// 输入：stock = [5,7,9,1,2]
// 输出：1

// 提示：
//     1 <= stock.length <= 5000
//     -5000 <= stock[i] <= 5000

import "fmt"

func stockManagement(stock []int) int {
    low, high := 0, len(stock)-1
    for low < high {
        if stock[low] < stock[high] {
            return stock[low]
        }
        mid := low + (high-low) >> 1
        if stock[mid] > stock[low] {
            low = mid + 1
        } else if stock[mid] == stock[low] { // 判断是否是相等元素
            low++
        } else {
            high = mid
        }
    }
    return stock[low]
}

// best solution
func stockManagement1(stock []int) int {
    low, high := 0, len(stock) - 1
    for low < high {
        mid := low + (high - low) / 2
        if stock[high] < stock[mid] {
            low = mid + 1
        } else if stock[high] == stock[mid] { // 判断是否是相等元素
            high--
        } else {
            high = mid
        }
    }
    return stock[low]
}

func stockManagement2(stock []int) int {
    mn := stock[0]
    for i := 1; i < len(stock); i++ {
        if stock[i] < mn {
            mn = stock[i]
        }
    }
    return mn
}

func main() {
    // 示例 1：
    // 输入：stock = [4,5,8,3,4]
    // 输出：3
    fmt.Printf("stockManagement([]int{ 4,5,8,3,4 }) = %v\n", stockManagement([]int{ 4,5,8,3,4 })) // 3
    // 示例 2：
    // 输入：stock = [5,7,9,1,2]
    // 输出：1
    fmt.Printf("stockManagement([]int{ 5,7,9,1,2 }) = %v\n", stockManagement([]int{ 5,7,9,1,2 })) // 1

    fmt.Printf("stockManagement([]int{ 1,3,5 }) = %v\n", stockManagement([]int{ 1,3,5 })) // 1
    fmt.Printf("stockManagement([]int{ 2,2,2,0,1 }) = %v\n", stockManagement([]int{ 2,2,2,0,1 })) // 0

    fmt.Printf("stockManagement1([]int{ 4,5,8,3,4 }) = %v\n", stockManagement1([]int{ 4,5,8,3,4 })) // 3
    fmt.Printf("stockManagement1([]int{ 5,7,9,1,2 }) = %v\n", stockManagement1([]int{ 5,7,9,1,2 })) // 1
    fmt.Printf("stockManagement1([]int{ 1,3,5 }) = %v\n", stockManagement1([]int{ 1,3,5 })) // 1
    fmt.Printf("stockManagement1([]int{ 2,2,2,0,1 }) = %v\n", stockManagement1([]int{ 2,2,2,0,1 })) // 0

    fmt.Printf("stockManagement2([]int{ 4,5,8,3,4 }) = %v\n", stockManagement2([]int{ 4,5,8,3,4 })) // 3
    fmt.Printf("stockManagement2([]int{ 5,7,9,1,2 }) = %v\n", stockManagement2([]int{ 5,7,9,1,2 })) // 1
    fmt.Printf("stockManagement2([]int{ 1,3,5 }) = %v\n", stockManagement2([]int{ 1,3,5 })) // 1
    fmt.Printf("stockManagement2([]int{ 2,2,2,0,1 }) = %v\n", stockManagement2([]int{ 2,2,2,0,1 })) // 0
}
