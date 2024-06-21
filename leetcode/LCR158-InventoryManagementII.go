package main

// LCR 158. 库存管理 II
// 仓库管理员以数组 stock 形式记录商品库存表。
// stock[i] 表示商品 id，可能存在重复。请返回库存表中数量大于 stock.length / 2 的商品 id。

// 示例 1:
// 输入: stock = [6, 1, 3, 1, 1, 1]
// 输出: 1

// 限制：
//     1 <= stock.length <= 50000
//     给定数组为非空数组，且存在结果数字

// # 解题思路
//     给定一个大小为 n 的数组，找到其中的众数。
//     众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。

import "fmt"

// 时间复杂度 O(n) 空间复杂度 O(1)
func inventoryManagement(stock []int) int {
    res, count := 0, 0 // 默认第一的值就是返回值 只有一个的话 直接返回了
    for i := 0; i < len(stock); i++ {
        if count == 0 { // 如果累加到 0 重新赋值
            res, count = stock[i], 1
        } else {
            if stock[i] == res {
                count++ // 如果还是自己 普累加
            } else {
                count-- // 如果不是自己就减去 1
            }
        }
    }
    return res
}

// 时间复杂度 O(n) 空间复杂度 O(n)
func inventoryManagement1(stock []int) int {
    m := make(map[int]int) // 声明一个map来保存数量
    l := len(stock) / 2
    for _, v := range stock {
        m[v]++
        if m[v] > l { // 如果统计到的数值多于一半了直接返回
            return v
        }
    }
    return 0
}

// 思路和解法1一样
func inventoryManagement2(stock []int) int {
    major, count := stock[0], 1
    for i := 1; i < len(stock); i++ {
        if count == 0 {
            major = stock[i]
            count++
        } else if major == stock[i] {
            count++
        } else {
            count--
        }
    }
    return major
}

func inventoryManagement3(stock []int) int {
    res, num := stock[0], 1
    for _, v := range stock[1:] {
        if v == res {
            num++
        } else if num == 0{
            res = v
            num = 1
        }else {
            num--
        }
    }
    return res
}

func main() {
    fmt.Printf("inventoryManagement([]int{ 3,2,3 }) = %v\n",inventoryManagement([]int{ 3,2,3 })) // 3
    fmt.Printf("inventoryManagement([]int{ 2,2,1,1,1,2,2 }) = %v\n",inventoryManagement([]int{ 2,2,1,1,1,2,2 })) // 2
    fmt.Printf("inventoryManagement([]int{ 6, 1, 3, 1, 1, 1 }) = %v\n",inventoryManagement([]int{ 6, 1, 3, 1, 1, 1})) // 1

    fmt.Printf("inventoryManagement1([]int{ 3,2,3 }) = %v\n",inventoryManagement1([]int{ 3,2,3 })) // 3
    fmt.Printf("inventoryManagement1([]int{ 2,2,1,1,1,2,2 }) = %v\n",inventoryManagement1([]int{ 2,2,1,1,1,2,2 })) // 2
    fmt.Printf("inventoryManagement1([]int{ 6, 1, 3, 1, 1, 1 }) = %v\n",inventoryManagement1([]int{ 6, 1, 3, 1, 1, 1})) // 1

    fmt.Printf("inventoryManagement2([]int{ 3,2,3 }) = %v\n",inventoryManagement2([]int{ 3,2,3 })) // 3
    fmt.Printf("inventoryManagement2([]int{ 2,2,1,1,1,2,2 }) = %v\n",inventoryManagement2([]int{ 2,2,1,1,1,2,2 })) // 2
    fmt.Printf("inventoryManagement2([]int{ 6, 1, 3, 1, 1, 1 }) = %v\n",inventoryManagement2([]int{ 6, 1, 3, 1, 1, 1})) // 1
    
    fmt.Printf("inventoryManagement3([]int{ 3,2,3 }) = %v\n",inventoryManagement3([]int{ 3,2,3 })) // 3
    fmt.Printf("inventoryManagement3([]int{ 2,2,1,1,1,2,2 }) = %v\n",inventoryManagement3([]int{ 2,2,1,1,1,2,2 })) // 2
    fmt.Printf("inventoryManagement3([]int{ 6, 1, 3, 1, 1, 1 }) = %v\n",inventoryManagement3([]int{ 6, 1, 3, 1, 1, 1})) // 1
}