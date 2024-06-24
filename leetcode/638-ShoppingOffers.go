package main

// 638. Shopping Offers
// In LeetCode Store, there are n items to sell. 
// Each item has a price. However, there are some special offers, 
// and a special offer consists of one or more different kinds of items with a sale price.

// You are given an integer array price where price[i] is the price of the ith item, 
// and an integer array needs where needs[i] is the number of pieces of the ith item you want to buy.

// You are also given an array special where special[i] is of size n + 1
// where special[i][j] is the number of pieces of the jth item in the ith offer and special[i][n] (i.e., the last integer in the array) is the price of the ith offer.

// Return the lowest price you have to pay for exactly certain items as given, where you could make optimal use of the special offers. 
// You are not allowed to buy more items than you want, even if that would lower the overall price. 
// You could use any of the special offers as many times as you want.

// Example 1:
// Input: price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
// Output: 14
// Explanation: There are two kinds of items, A and B. Their prices are $2 and $5 respectively. 
// In special offer 1, you can pay $5 for 3A and 0B
// In special offer 2, you can pay $10 for 1A and 2B. 
// You need to buy 3A and 2B, so you may pay $10 for 1A and 2B (special offer #2), and $4 for 2A.

// Example 2:
// Input: price = [2,3,4], special = [[1,1,0,4],[2,2,1,9]], needs = [1,2,1]
// Output: 11
// Explanation: The price of A is $2, and $3 for B, $4 for C. 
// You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C. 
// You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer #1), and $3 for 1B, $4 for 1C. 
// You cannot add more items, though only $9 for 2A ,2B and 1C.

// Constraints:
//     n == price.length == needs.length
//     1 <= n <= 6
//     0 <= price[i], needs[i] <= 10
//     1 <= special.length <= 100
//     special[i].length == n + 1
//     0 <= special[i][j] <= 50

import "fmt"

func shoppingOffers(price []int, special [][]int, needs []int) int {
    res := 1 << 32 -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i, curPrice int)
    dfs = func(sp, curPrice int) { // sp:当前的大礼包号
        if curPrice > res { return }// 如果当前总花费已经超过了目前的最小花费 直接返回
        for _, need := range needs { // 不能购买超出需求的物品
            if need < 0 { return }
        }
        if sp == len(special) { // 大礼包购买完毕(可能一个都没有购买或者某几个礼包重复购买)
            for i := 0; i < len(needs); i++ { // 加上购买大礼包剩下的物品总价
                curPrice += needs[i] * price[i]
            }
            res = min(res, curPrice) // 更便宜的
            return
        }
        for i := 0; i < len(price); i++ { // 购买大礼包
            needs[i] -= special[sp][i]
        }
        dfs(sp, curPrice+special[sp][len(price)])
        for i := 0; i < len(price); i++ { // 恢复
            needs[i] += special[sp][i]
        }
        dfs(sp+1, curPrice) // 不购买大礼包
    }
    dfs(0, 0)
    return res
}

func shoppingOffers1(price []int, special [][]int, needs []int) int {
    bits, n, mask, f  := 4, len(needs), 0, make(map[int]int)
    for i, need := range needs {
        mask |= need << (i * bits)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(int) int
    dfs = func(cur int) int {
        if v, ok := f[cur]; ok {
            return v
        }
        res := 0
        for i := 0; i < n; i++ {
            res += price[i] * ((cur >> (i * bits)) & 0xf)
        }
        for _, offer := range special {
            next := cur
            ok := true
            for j := 0; j < n; j++ {
                if ((cur >> (j * bits)) & 0xf) < offer[j] {
                    ok = false
                    break
                }
                next -= offer[j] << (j * bits)
            }
            if ok {
                res = min(res, offer[n] + dfs(next))
            }
        }
        f[cur] = res
        return res
    }
    return dfs(mask)
}

func main() {
    // Example 1:
    // Input: price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
    // Output: 14
    // Explanation: There are two kinds of items, A and B. Their prices are $2 and $5 respectively. 
    // In special offer 1, you can pay $5 for 3A and 0B
    // In special offer 2, you can pay $10 for 1A and 2B. 
    // You need to buy 3A and 2B, so you may pay $10 for 1A and 2B (special offer #2), and $4 for 2A.
    fmt.Println(shoppingOffers([]int{2,5},[][]int{{3,0,5},{1,2,10}},[]int{3,2})) // 14
    // Example 2:
    // Input: price = [2,3,4], special = [[1,1,0,4],[2,2,1,9]], needs = [1,2,1]
    // Output: 11
    // Explanation: The price of A is $2, and $3 for B, $4 for C. 
    // You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C. 
    // You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer #1), and $3 for 1B, $4 for 1C. 
    // You cannot add more items, though only $9 for 2A ,2B and 1C.
    fmt.Println(shoppingOffers([]int{2,3,4},[][]int{{1,1,0,4},{2,2,1,9}},[]int{1,2,1})) // 11

    fmt.Println(shoppingOffers1([]int{2,5},[][]int{{3,0,5},{1,2,10}},[]int{3,2})) // 14
    fmt.Println(shoppingOffers1([]int{2,3,4},[][]int{{1,1,0,4},{2,2,1,9}},[]int{1,2,1})) // 11
}