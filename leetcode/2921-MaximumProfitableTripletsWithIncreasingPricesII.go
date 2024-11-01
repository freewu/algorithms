package main

// 2921. Maximum Profitable Triplets With Increasing Prices II
// Given the 0-indexed arrays prices and profits of length n. 
// There are n items in an store where the ith item has a price of prices[i] and a profit of profits[i].

// We have to pick three items with the following condition:
//     prices[i] < prices[j] < prices[k] where i < j < k.

// If we pick items with indices i, j and k satisfying the above condition, the profit would be profits[i] + profits[j] + profits[k].

// Return the maximum profit we can get, and -1 if it's not possible to pick three items with the given condition.

// Example 1:
// Input: prices = [10,2,3,4], profits = [100,2,7,10]
// Output: 19
// Explanation: We can't pick the item with index i=0 since there are no indices j and k such that the condition holds.
// So the only triplet we can pick, are the items with indices 1, 2 and 3 and it's a valid pick since prices[1] < prices[2] < prices[3].
// The answer would be sum of their profits which is 2 + 7 + 10 = 19.

// Example 2:
// Input: prices = [1,2,3,4,5], profits = [1,5,3,4,6]
// Output: 15
// Explanation: We can select any triplet of items since for each triplet of indices i, j and k such that i < j < k, the condition holds.
// Therefore the maximum profit we can get would be the 3 most profitable items which are indices 1, 3 and 4.
// The answer would be sum of their profits which is 5 + 4 + 6 = 15.

// Example 3:
// Input: prices = [4,3,2,1], profits = [33,20,19,87]
// Output: -1
// Explanation: We can't select any triplet of indices such that the condition holds, so we return -1.

// Constraints:
//     3 <= prices.length == profits.length <= 50000
//     1 <= prices[i] <= 5000
//     1 <= profits[i] <= 10^6

import "fmt"
import "slices"

type BinaryIndexedTree struct {
    Size int
    Tree []int
}

func Constructor(n int) BinaryIndexedTree {
    return BinaryIndexedTree{ Size: n, Tree: make([]int, n + 1), }
}

func (this BinaryIndexedTree) lowbit(x int) int {
    return x & (-x)
}

func (this BinaryIndexedTree) GetMax(index int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for index > 0 {
        res = max(res, this.Tree[index])
        index -= this.lowbit(index)
    }
    return res
}

func (this BinaryIndexedTree) Update(index, val int) {
    for index <= this.Size && this.Tree[index] < val {
        this.Tree[index] = val
        index += this.lowbit(index)
    }
}

func maxProfit(prices []int, profits []int) int {
    res, n, mx := -1, len(prices), 0
    leftMaxProfits, rightMaxProfits := make([]int, n), make([]int, n)
    for _, v := range prices { // 获取最大的价格
        if v > mx { mx = v }
    }
    bit1, bit2 := Constructor(mx + 1), Constructor(mx + 1)
    for i := 0; i < n; i++ {
        index := prices[i];
        leftMaxProfits[i] = bit1.GetMax(index - 1)
        bit1.Update(index, profits[i])
    }
    for i := n - 1; i >= 0; i-- {
        index := mx - prices[i] + 1
        rightMaxProfits[i] = bit2.GetMax(index - 1)
        bit2.Update(index, profits[i])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n - 1; i++ {
        if leftMaxProfits[i] > 0 && rightMaxProfits[i] > 0 {
            res = max(res, leftMaxProfits[i] + profits[i] + rightMaxProfits[i])
        }
    }
    return res
}

func maxProfit1(prices []int, profits []int) int {
    // 树状数组  + 枚举中点 O(n* logm) m为 max(price)
    // 树状数组 - 维护前缀最大值
    // 而对于i位置的右侧, 需要求出的是 [x+1,sz]的后缀最大值,  需要将后缀转换为前缀(即在fenwick中的坐标继续, idx变为 sz-idx)
    n,inf := len(prices), 1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // fenwick
    update := func(tree []int, i int, val int) {
        for ; i < len(tree); i += i & -i {
            tree[i] = max(tree[i], val)
        }
    }
    query := func(tree []int, i int) int {
        res := -inf
        for ; i > 0; i &= i - 1 {
            res = max(res, tree[i])
        }
        return res
    }
    // 计算i左右侧 满足 <[i] 的最大profit和 >[i] 的最大profit
    mx := slices.Max(prices)
    sz := mx + 1// 请求的范围 是 值+1,所以必须申请的范围应该是 值域的最大值+1
    leftBit, rightBit := make([]int, mx + 2), make([]int, sz + 1) // bit: BinaryIndexedTree
    for i := range leftBit {
        leftBit[i], rightBit[i] = -inf, -inf // 取值是可能取不到值(没有一个满足条件的,此时不能返回0,应该返回一个极值代表无效情况
    }
    rightMx := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        rightMx[i] = query(rightBit, sz - (prices[i]+1)) // 树状数组只支持 前缀最大值, 求[x,sz]的后缀最大值应该转换前缀最大值
        update(rightBit, sz-prices[i], profits[i])
    }
    res := -1
    for i, price := range prices {
        v := query(leftBit, price - 1)
        res = max(res, v + profits[i] + rightMx[i])
        update(leftBit, price, profits[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: prices = [10,2,3,4], profits = [100,2,7,10]
    // Output: 19
    // Explanation: We can't pick the item with index i=0 since there are no indices j and k such that the condition holds.
    // So the only triplet we can pick, are the items with indices 1, 2 and 3 and it's a valid pick since prices[1] < prices[2] < prices[3].
    // The answer would be sum of their profits which is 2 + 7 + 10 = 19.
    fmt.Println(maxProfit([]int{10,2,3,4}, []int{100,2,7,10})) // 19
    // Example 2:
    // Input: prices = [1,2,3,4,5], profits = [1,5,3,4,6]
    // Output: 15
    // Explanation: We can select any triplet of items since for each triplet of indices i, j and k such that i < j < k, the condition holds.
    // Therefore the maximum profit we can get would be the 3 most profitable items which are indices 1, 3 and 4.
    // The answer would be sum of their profits which is 5 + 4 + 6 = 15.
    fmt.Println(maxProfit([]int{1,2,3,4,5}, []int{1,5,3,4,6})) // 15
    // Example 3:
    // Input: prices = [4,3,2,1], profits = [33,20,19,87]
    // Output: -1
    // Explanation: We can't select any triplet of indices such that the condition holds, so we return -1.
    fmt.Println(maxProfit([]int{4,3,2,1}, []int{33,20,19,87})) // -1

    fmt.Println(maxProfit1([]int{10,2,3,4}, []int{100,2,7,10})) // 19
    fmt.Println(maxProfit1([]int{1,2,3,4,5}, []int{1,5,3,4,6})) // 15
    fmt.Println(maxProfit1([]int{4,3,2,1}, []int{33,20,19,87})) // -1
}