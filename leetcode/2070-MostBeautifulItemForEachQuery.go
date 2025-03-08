package main

// 2070. Most Beautiful Item for Each Query
// You are given a 2D integer array items where items[i] = [pricei, beautyi] denotes the price and beauty of an item respectively.

// You are also given a 0-indexed integer array queries. 
// For each queries[j], you want to determine the maximum beauty of an item whose price is less than or equal to queries[j]. 
// If no such item exists, then the answer to this query is 0.

// Return an array answer of the same length as queries where answer[j] is the answer to the jth query.

// Example 1:
// Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
// Output: [2,4,5,5,6,6]
// Explanation:
// - For queries[0]=1, [1,2] is the only item which has price <= 1. Hence, the answer for this query is 2.
// - For queries[1]=2, the items which can be considered are [1,2] and [2,4]. 
//   The maximum beauty among them is 4.
// - For queries[2]=3 and queries[3]=4, the items which can be considered are [1,2], [3,2], [2,4], and [3,5].
//   The maximum beauty among them is 5.
// - For queries[4]=5 and queries[5]=6, all items can be considered.
//   Hence, the answer for them is the maximum beauty of all items, i.e., 6.

// Example 2:
// Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
// Output: [4]
// Explanation: 
// The price of every item is equal to 1, so we choose the item with the maximum beauty 4. 
// Note that multiple items can have the same price and/or beauty.  

// Example 3:
// Input: items = [[10,1000]], queries = [5]
// Output: [0]
// Explanation:
// No item has a price less than or equal to 5, so no item can be chosen.
// Hence, the answer to the query is 0.

// Constraints:
//     1 <= items.length, queries.length <= 10^5
//     items[i].length == 2
//     1 <= pricei, beautyi, queries[j] <= 10^9

import "fmt"
import "sort"
import "slices"

func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool { // 排序
        return items[i][0] < items[j][0]
    })
    for i := 1; i < len(items); i++ {
        if items[i][1] < items[i-1][1] { 
            items[i][1] = items[i-1][1] 
        }
    }
    res := []int{}
    for _, query := range queries {
        left, right := 0, len(items) // 二分
        for left < right {
            mid := (right - left) / 2 + left
            if items[mid][0] > query {
                right = mid
            } else {
                left = mid + 1
            }
        }
        if left == 0 { // 所有物品的价格都大于查询价格
            res = append(res, []int{0}...)
        } else { // 返回小于等于查询价格的物品的最大美丽值
            res = append(res, []int{items[left - 1][1]}...)
        }
    }
    return res
}

// 超出时间限制 33 / 35 
func maximumBeauty1(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool { // 按价格排序
        return items[i][0] < items[j][0] 
    }) 
    for i, q := range queries {
        queries[i] = q << 32 | i // 这样排序时可以保留查询的下标
    }
    sort.Ints(queries)
    res, mx := make([]int, len(queries)), 0
    for _, q := range queries {
        for i := 0; i < len(items) && items[i][0] <= q >> 32; i++ {
            if items[i][1] > mx { mx = items[i][1] }
        }
        res[q & (1 << 32 - 1)] = mx
    }
    return res
}

func maximumBeauty2(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool { // 按价格排序
        return items[i][0] < items[j][0]
    })
    prefMaxBeauty := make([]int, len(items))
    for i, item := range items {
        prefMaxBeauty[i] = item[1]
        if i > 0 && prefMaxBeauty[i - 1] > prefMaxBeauty[i] {
            prefMaxBeauty[i] = prefMaxBeauty[i - 1]
        }
    }
    res := make([]int, 0, len(queries))
    for _, query := range queries {
        pos, _ := slices.BinarySearchFunc(items, query + 1, func(item []int, target int) int {
            return item[0] - target
        })
        v := 0
        if pos > 0 {
            v = prefMaxBeauty[pos - 1]
        }
        res = append(res, v)
    }
    return res
}

func maximumBeauty3(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool { // 按价格排序
        return items[i][0] < items[j][0]
    })
    k := 0
    for _, item := range items[1:] {
        if item[1] > items[k][1] {
            k++
            items[k] = item
        }
    }
    for i, q := range queries {
        j := sort.Search(k + 1, func(i int) bool { 
            return items[i][0] > q 
        })
        if j > 0 {
            queries[i] = items[j - 1][1]
        } else {
            queries[i] = 0
        }
    }
    return queries
}

func main() {
    // Example 1:
    // Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
    // Output: [2,4,5,5,6,6]
    // Explanation:
    // - For queries[0]=1, [1,2] is the only item which has price <= 1. Hence, the answer for this query is 2.
    // - For queries[1]=2, the items which can be considered are [1,2] and [2,4]. 
    //   The maximum beauty among them is 4.
    // - For queries[2]=3 and queries[3]=4, the items which can be considered are [1,2], [3,2], [2,4], and [3,5].
    //   The maximum beauty among them is 5.
    // - For queries[4]=5 and queries[5]=6, all items can be considered.
    //   Hence, the answer for them is the maximum beauty of all items, i.e., 6.
    fmt.Println(maximumBeauty([][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}, []int{1,2,3,4,5,6})) // [2,4,5,5,6,6]
    // Example 2:
    // Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
    // Output: [4]
    // Explanation: 
    // The price of every item is equal to 1, so we choose the item with the maximum beauty 4. 
    // Note that multiple items can have the same price and/or beauty.  
    fmt.Println(maximumBeauty([][]int{{1,2},{1,2},{1,3},{1,4}}, []int{1})) // [4]
    // Example 3:
    // Input: items = [[10,1000]], queries = [5]
    // Output: [0]
    // Explanation:
    // No item has a price less than or equal to 5, so no item can be chosen.
    // Hence, the answer to the query is 0.
    fmt.Println(maximumBeauty([][]int{{10,1000}}, []int{5})) // [0]

    fmt.Println(maximumBeauty1([][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}, []int{1,2,3,4,5,6})) // [2,4,5,5,6,6] 
    fmt.Println(maximumBeauty1([][]int{{1,2},{1,2},{1,3},{1,4}}, []int{1})) // [4]
    fmt.Println(maximumBeauty1([][]int{{10,1000}}, []int{5})) // [0]

    fmt.Println(maximumBeauty2([][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}, []int{1,2,3,4,5,6})) // [2,4,5,5,6,6] 
    fmt.Println(maximumBeauty2([][]int{{1,2},{1,2},{1,3},{1,4}}, []int{1})) // [4]
    fmt.Println(maximumBeauty2([][]int{{10,1000}}, []int{5})) // [0]

    fmt.Println(maximumBeauty3([][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}, []int{1,2,3,4,5,6})) // [2,4,5,5,6,6] 
    fmt.Println(maximumBeauty3([][]int{{1,2},{1,2},{1,3},{1,4}}, []int{1})) // [4]
    fmt.Println(maximumBeauty3([][]int{{10,1000}}, []int{5})) // [0]
}