package main

// LCR 159. 库存管理 III
// 仓库管理员以数组 stock 形式记录商品库存表，其中 stock[i] 表示对应商品库存余量。
// 请返回库存余量最少的 cnt 个商品余量，返回 顺序不限。

// 示例 1：
// 输入：stock = [2,5,7,4], cnt = 1
// 输出：[2]

// 示例 2：
// 输入：stock = [0,2,3,6], cnt = 2
// 输出：[0,2] 或 [2,0]

// 提示：
//     0 <= cnt <= stock.length <= 10000
// 	   0 <= stock[i] <= 10000

import "fmt"
import "sort"

func inventoryManagement(stock []int, cnt int) []int {
    if len(stock) <= cnt {
        return stock
    }
    sort.Ints(stock)
    return stock[:cnt]
}

func inventoryManagement1(stock []int, cnt int) []int {
    if cnt <= 0 { return nil }
    if cnt >= len(stock) { return stock }
    var quickPartition func(arr []int, left, right, k int) int
    quickPartition = func(arr []int, left, right, k int) int {
        if left > right { return 0 }
        p, leftIndex := right, left
        for i:=left;i<right;i++{
            if arr[i] < arr[p] {
                arr[i], arr[leftIndex] = arr[leftIndex], arr[i]
                leftIndex++
            }
        }
        arr[leftIndex], arr[p] = arr[p], arr[leftIndex]
        if leftIndex == k {
            return k
        } else if leftIndex > k {
            return quickPartition(arr, left, leftIndex-1, k)
        }
        return quickPartition(arr, leftIndex+1, right, k)
    }
    if quickPartition(stock, 0, len(stock)-1, cnt) != cnt {
        return nil
    }
    return stock[:cnt]
}

func main() {
    // 示例 1：
    // 输入：stock = [2,5,7,4], cnt = 1
    // 输出：[2] 
    fmt.Println(inventoryManagement([]int{2,5,7,4},1)) // [2]
    // 示例 2：
    // 输入：stock = [0,2,3,6], cnt = 2
    // 输出：[0,2] 或 [2,0]
    fmt.Println(inventoryManagement([]int{0,2,3,6},2)) // [0,2] 或 [2,0]
    
    fmt.Println(inventoryManagement1([]int{2,5,7,4},1)) // [2]
    fmt.Println(inventoryManagement1([]int{0,2,3,6},2)) // [0,2] 或 [2,0]
}
