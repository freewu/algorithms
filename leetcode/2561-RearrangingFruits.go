package main

// 2561. Rearranging Fruits
// You have two fruit baskets containing n fruits each. 
// You are given two 0-indexed integer arrays basket1 and basket2 representing the cost of fruit in each basket. 
// You want to make both baskets equal. 
// To do so, you can use the following operation as many times as you want:
//     1. Chose two indices i and j, and swap the ith fruit of basket1 with the jth fruit of basket2.
//     2. The cost of the swap is min(basket1[i],basket2[j]).

// Two baskets are considered equal if sorting them according to the fruit cost makes them exactly the same baskets.

// Return the minimum cost to make both the baskets equal or -1 if impossible.

// Example 1:
// Input: basket1 = [4,2,2,2], basket2 = [1,4,1,2]
// Output: 1
// Explanation: Swap index 1 of basket1 with index 0 of basket2, which has cost 1. Now basket1 = [4,1,2,2] and basket2 = [2,4,1,2]. Rearranging both the arrays makes them equal.

// Example 2:
// Input: basket1 = [2,3,4,1], basket2 = [3,2,5,1]
// Output: -1
// Explanation: It can be shown that it is impossible to make both the baskets equal.

// Constraints:
//     basket1.length == basket2.length
//     1 <= basket1.length <= 10^5
//     1 <= basket1[i],basket2[i] <= 10^9

import "fmt"
import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
    count := make(map[int]int)
    for i, v := range basket1 {
        count[v]++
        count[basket2[i]]--
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res, mn, arr := int64(0), 1 << 31, []int{}
    for k, v := range count {
        if v % 2 != 0 { // 存在奇数个的情况
            return -1
        }
        for i := abs(v) / 2; i > 0; i-- {
            arr = append(arr, k)
        }
        mn = min(mn, k)
    }
    sort.Ints(arr)
    n := len(arr)
    for i := 0; i < n / 2; i++ {
        res += int64(min(arr[i], mn * 2))
    }
    return res
}

func minCost1(basket1 []int, basket2 []int) int64 {
    count := make(map[int]int)
    if len(basket1) != len(basket2) { return -1 }
    n := len(basket1)
    for i := 0; i < n; i++ {
        count[basket1[i]]++
        count[basket2[i]]--
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res, mn, arr := 0, 1 << 31, []int{}
    for k, v := range count {
        mn = min(mn, k)
        if abs(v) % 2 != 0 { return -1 }
        for i := abs(v) / 2; i > 0; i-- {
            arr = append(arr, k)
        }
    }
    sort.Ints(arr)
    for i := 0; i < len(arr) / 2; i++ {
        res += min(arr[i], mn * 2)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: basket1 = [4,2,2,2], basket2 = [1,4,1,2]
    // Output: 1
    // Explanation: Swap index 1 of basket1 with index 0 of basket2, which has cost 1. Now basket1 = [4,1,2,2] and basket2 = [2,4,1,2]. Rearranging both the arrays makes them equal.
    fmt.Println(minCost([]int{4,2,2,2}, []int{1,4,1,2})) // 1
    // Example 2:
    // Input: basket1 = [2,3,4,1], basket2 = [3,2,5,1]
    // Output: -1
    // Explanation: It can be shown that it is impossible to make both the baskets equal.
    fmt.Println(minCost([]int{2,3,4,1}, []int{3,2,5,1})) // -1

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost([]int{9,9,8,8,7,7,6,6,5}, []int{1,1,2,2,3,3,4,4,5})) // 7
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0

    fmt.Println(minCost1([]int{4,2,2,2}, []int{1,4,1,2})) // 1
    fmt.Println(minCost1([]int{2,3,4,1}, []int{3,2,5,1})) // -1
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost1([]int{9,9,8,8,7,7,6,6,5}, []int{1,1,2,2,3,3,4,4,5})) // 7
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
}