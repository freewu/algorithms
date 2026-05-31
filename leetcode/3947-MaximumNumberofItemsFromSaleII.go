package main

// 3947. Maximum Number of Items From Sale II
// You are given a 2D integer array items, where items[i] = [factori, pricei] represents the ith item. 
// You are also given an integer budget.

// There are unlimited copies of each item available for purchase. 
// You may buy any number of copies of any items such that the total cost of the purchased copies is at most budget.

// After buying items, you may receive free copies according to the following rules:
//     1. Each purchased copy of item i can give you at most one free copy of another item j.
//     2. The free item must satisfy i != j and factori divides factorj.
//     3. For each ordered pair (i, j), you can receive a free copy of item j from purchases of item i at most once, regardless of how many copies of item i you buy.  
//     4. The same item j can be received multiple times for free if it is received from purchases of different item types.

// Return the maximum total number of item copies you can obtain, including both purchased copies and free copies, while spending at most budget on purchased items.

// Example 1:
// Input: items = [[1,6],[2,4],[3,5]], budget = 19
// Output: 5
// Explanation:
// You can buy 2 copies of item 0 and 1 copy of item 1 for a total cost of 2 * 6 + 4 = 16, which is not greater than budget = 19.
// One purchased copy of item 0 gives 1 free copy of item 1, because factor0 = 1 divides factor1 = 2.
// The other purchased copy of item 0 gives 1 free copy of item 2, because factor0 = 1 divides factor2 = 3.
// You leave with 3 purchased copies and 2 free copies, for a total of 5 item copies.

// Example 2:
// Input: items = [[2,8],[1,10],[6,6],[4,12],[5,20],[5,17]], budget = 35
// Output: 7
// Explanation:
// You can buy 2 copies of item 0, 1 copy of item 1, and 1 copy of item 2 for a total cost of 2 * 8 + 10 + 6 = 32, which is not greater than budget = 35.
// One purchased copy of item 0 gives 1 free copy of item 2, because factor0 = 2 divides factor2 = 6.
// The other purchased copy of item 0 gives 1 free copy of item 3, because factor0 = 2 divides factor3 = 4.
// The purchased copy of item 1 gives 1 free copy of item 2, because factor1 = 1 divides factor2 = 6.
// Buying item 2 gives no free copy, because factor2 = 6 does not divide the factor of any other item.
// You leave with 4 purchased copies and 3 free copies, for a total of 7 item copies.

// Constraints:
//     1 <= items.length <= 10^5
//     items[i] = [factori, pricei]
//     1 <= factori <= items.length
//     1 <= pricei <= 10^9
//     1 <= budget <= 10^9

import "fmt"
import "slices"

// 贪心
func maximumSaleItems(items [][]int, budget int) int {
    res, n, minPrice := 0, len(items), 1 << 31
    count := make([]int, n+1)
    for _, p := range items {
        count[p[0]]++
        minPrice = min(minPrice, p[1])
    }
    multi := make([]int, n + 1)
    type Pair struct{ price, cnt int }
    arr := []Pair{}
    for _, p := range items {
        factor, price := p[0], p[1]
        if price >= minPrice * 2 { // 
            continue
        }
        if multi[factor] == 0 { // 之前没有计算过
            for j := factor; j <= n; j += factor {
                multi[factor] += count[j]
            }
        }
        if multi[factor] > 1 {
            arr = append(arr, Pair{price, multi[factor] - 1}) // factor 的倍数不包括该物品
        }
    }
    slices.SortFunc(arr, func(a, b Pair) int { return a.price - b.price })
    for _, p := range arr {
        if budget < p.price { // 没钱了
            break
        }
        c := min(p.cnt, budget/p.price) // 该物品最多买 c 个
        budget -= p.price * c
        res += c * 2
    }
    return res + budget / minPrice // 剩余的钱买最便宜的物品
}

func main() {
    // Example 1:
    // Input: items = [[1,6],[2,4],[3,5]], budget = 19
    // Output: 5
    // Explanation:
    // You can buy 2 copies of item 0 and 1 copy of item 1 for a total cost of 2 * 6 + 4 = 16, which is not greater than budget = 19.
    // One purchased copy of item 0 gives 1 free copy of item 1, because factor0 = 1 divides factor1 = 2.
    // The other purchased copy of item 0 gives 1 free copy of item 2, because factor0 = 1 divides factor2 = 3.
    // You leave with 3 purchased copies and 2 free copies, for a total of 5 item copies.
    fmt.Println(maximumSaleItems([][]int{{1,6},{2,4},{3,5}}, 19)) // 5
    // Example 2:
    // Input: items = [[2,8],[1,10],[6,6],[4,12],[5,20],[5,17]], budget = 35
    // Output: 7
    // Explanation:
    // You can buy 2 copies of item 0, 1 copy of item 1, and 1 copy of item 2 for a total cost of 2 * 8 + 10 + 6 = 32, which is not greater than budget = 35.
    // One purchased copy of item 0 gives 1 free copy of item 2, because factor0 = 2 divides factor2 = 6.
    // The other purchased copy of item 0 gives 1 free copy of item 3, because factor0 = 2 divides factor3 = 4.
    // The purchased copy of item 1 gives 1 free copy of item 2, because factor1 = 1 divides factor2 = 6.
    // Buying item 2 gives no free copy, because factor2 = 6 does not divide the factor of any other item.
    // You leave with 4 purchased copies and 3 free copies, for a total of 7 item copies.
    fmt.Println(maximumSaleItems([][]int{{2,8},{1,10},{6,6},{4,12},{5,20},{5,17}}, 35)) // 7
}