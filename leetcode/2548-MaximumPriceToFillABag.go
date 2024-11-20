package main

// 2548. Maximum Price to Fill a Bag
// You are given a 2D integer array items where items[i] = [pricei, weighti] denotes the price and weight of the ith item, respectively.

// You are also given a positive integer capacity.

// Each item can be divided into two items with ratios part1 and part2, where part1 + part2 == 1.
//     1. The weight of the first item is weighti * part1 and the price of the first item is pricei * part1.
//     2. Similarly, the weight of the second item is weighti * part2 and the price of the second item is pricei * part2.

// Return the maximum total price to fill a bag of capacity capacity with given items. 
// If it is impossible to fill a bag return -1. 
// Answers within 10^-5 of the actual answer will be considered accepted.

// Example 1:
// Input: items = [[50,1],[10,8]], capacity = 5
// Output: 55.00000
// Explanation: 
// We divide the 2nd item into two parts with part1 = 0.5 and part2 = 0.5.
// The price and weight of the 1st item are 5, 4. And similarly, the price and the weight of the 2nd item are 5, 4.
// The array items after operation becomes [[50,1],[5,4],[5,4]]. 
// To fill a bag with capacity 5 we take the 1st element with a price of 50 and the 2nd element with a price of 5.
// It can be proved that 55.0 is the maximum total price that we can achieve.

// Example 2:
// Input: items = [[100,30]], capacity = 50
// Output: -1.00000
// Explanation: It is impossible to fill a bag with the given item.

// Constraints:
//     1 <= items.length <= 10^5
//     items[i].length == 2
//     1 <= pricei, weighti <= 10^4
//     1 <= capacity <= 10^9

import "fmt"
import "sort"

func maxPrice(items [][]int, capacity int) float64 {
    res := float64(0)
    sort.Slice(items, func(i, j int) bool { // 按总价从小到大排序
        return items[i][1] * items[j][0] < items[i][0] * items[j][1] 
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, item := range items {
        price, weight := item[0], item[1]
        v :=  min(weight, capacity)
        res += float64(v) / float64(weight) * float64(price)
        capacity -= v
    }
    if capacity > 0 { return -1 } // 无法填满背包
    return res
}

func main() {
    // Example 1:
    // Input: items = [[50,1],[10,8]], capacity = 5
    // Output: 55.00000
    // Explanation: 
    // We divide the 2nd item into two parts with part1 = 0.5 and part2 = 0.5.
    // The price and weight of the 1st item are 5, 4. And similarly, the price and the weight of the 2nd item are 5, 4.
    // The array items after operation becomes [[50,1],[5,4],[5,4]]. 
    // To fill a bag with capacity 5 we take the 1st element with a price of 50 and the 2nd element with a price of 5.
    // It can be proved that 55.0 is the maximum total price that we can achieve.
    fmt.Println(maxPrice([][]int{{50,1},{10,8}}, 5)) // 55.00000
    // Example 2:
    // Input: items = [[100,30]], capacity = 50
    // Output: -1.00000
    // Explanation: It is impossible to fill a bag with the given item.
    fmt.Println(maxPrice([][]int{{100,30}}, 50)) // -1.00000
}