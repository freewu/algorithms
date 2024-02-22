package main

// 2861. Maximum Number of Alloys 
// You are the owner of a company that creates alloys using various types of metals. 
// There are n different types of metals available, and you have access to k machines that can be used to create alloys. 
// Each machine requires a specific amount of each metal type to create an alloy.

// For the ith machine to create an alloy, it needs composition[i][j] units of metal of type j. 
// Initially, you have stock[i] units of metal type i, and purchasing one unit of metal type i costs cost[i] coins.

// Given integers n, k, budget, a 1-indexed 2D array composition, and 1-indexed arrays stock and cost, 
// your goal is to maximize the number of alloys the company can create while staying within the budget of budget coins.
// All alloys must be created with the same machine.

// Return the maximum number of alloys that the company can create.

// Example 1:
// Input: n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,0], cost = [1,2,3]
// Output: 2
// Explanation: It is optimal to use the 1st machine to create alloys.
// To create 2 alloys we need to buy the:
// - 2 units of metal of the 1st type.
// - 2 units of metal of the 2nd type.
// - 2 units of metal of the 3rd type.
// In total, we need 2 * 1 + 2 * 2 + 2 * 3 = 12 coins, which is smaller than or equal to budget = 15.
// Notice that we have 0 units of metal of each type and we have to buy all the required units of metal.
// It can be proven that we can create at most 2 alloys.

// Example 2:
// Input: n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,100], cost = [1,2,3]
// Output: 5
// Explanation: It is optimal to use the 2nd machine to create alloys.
// To create 5 alloys we need to buy:
// - 5 units of metal of the 1st type.
// - 5 units of metal of the 2nd type.
// - 0 units of metal of the 3rd type.
// In total, we need 5 * 1 + 5 * 2 + 0 * 3 = 15 coins, which is smaller than or equal to budget = 15.
// It can be proven that we can create at most 5 alloys.

// Example 3:
// Input: n = 2, k = 3, budget = 10, composition = [[2,1],[1,2],[1,1]], stock = [1,1], cost = [5,5]
// Output: 2
// Explanation: It is optimal to use the 3rd machine to create alloys.
// To create 2 alloys we need to buy the:
// - 1 unit of metal of the 1st type.
// - 1 unit of metal of the 2nd type.
// In total, we need 1 * 5 + 1 * 5 = 10 coins, which is smaller than or equal to budget = 10.
// It can be proven that we can create at most 2 alloys.

// Constraints:
//         1 <= n, k <= 100
//         0 <= budget <= 10^8
//         composition.length == k
//         composition[i].length == n
//         1 <= composition[i][j] <= 100
//         stock.length == cost.length == n
//         0 <= stock[i] <= 10^8
//         1 <= cost[i] <= 100

import "fmt"

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    // 最多每一种金属最多只能补到 2×10^8 个
    left, right, ans := 1, int(2e8), 0
    for left <= right {
        mid := (left + right) / 2
        var valid bool
        // 所有合金都需要由同一台机器制造
        for i := 0; i < k; i++ {
            var spend int64
            for j := 0; j < n; j++ {
                // 对于第 i 台机器以及第 j 种金属，它需要的数量为 composition[i][j]，当前已拥有的数量为 stock[j]
                spend += max(int64(composition[i][j]) * int64(mid) - int64(stock[j]), int64(0)) * int64(cost[j])
            }
            if spend <= int64(budget) {
                valid = true
                break
            }
        }
        if valid {
            ans, left = mid, mid + 1
        } else {
            right = mid - 1
        }
    }
    return ans
}

func maxNumberOfAlloys1(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    isValid := func(target int) bool {
        for _, currMachine := range composition {
            remain := budget
            for i, x := range currMachine {
                need := max(0, x*target-stock[i])
                remain -= need * cost[i]
            }
            if remain >= 0 {
                return true
            }
        }
        return false
    }

    l, r := 0, budget + stock[0]
    for l < r {
        mid := (l + r + 1) >> 1
        if isValid(mid) {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l
}

func main() {
// Example 1:
// Input: n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,0], cost = [1,2,3]
// Output: 2
// Explanation: It is optimal to use the 1st machine to create alloys.
// To create 2 alloys we need to buy the:
// - 2 units of metal of the 1st type.
// - 2 units of metal of the 2nd type.
// - 2 units of metal of the 3rd type.
// In total, we need 2 * 1 + 2 * 2 + 2 * 3 = 12 coins, which is smaller than or equal to budget = 15.
// Notice that we have 0 units of metal of each type and we have to buy all the required units of metal.
// It can be proven that we can create at most 2 alloys.
    fmt.Println(maxNumberOfAlloys(
        3,2,15,
        [][]int{[]int{1,1,1},[]int{1,1,10}},
        []int{0,0,0},
        []int{1,2,3},
    )) // 2

// Example 2:
// Input: n = 3, k = 2, budget = 15, composition = [[1,1,1],[1,1,10]], stock = [0,0,100], cost = [1,2,3]
// Output: 5
// Explanation: It is optimal to use the 2nd machine to create alloys.
// To create 5 alloys we need to buy:
// - 5 units of metal of the 1st type.
// - 5 units of metal of the 2nd type.
// - 0 units of metal of the 3rd type.
// In total, we need 5 * 1 + 5 * 2 + 0 * 3 = 15 coins, which is smaller than or equal to budget = 15.
// It can be proven that we can create at most 5 alloys.
    fmt.Println(maxNumberOfAlloys(
        3,2,15,
        [][]int{[]int{1,1,1},[]int{1,1,10}},
        []int{0,0,100},
        []int{1,2,3},
    )) // 5

// Example 3:
// Input: n = 2, k = 3, budget = 10, composition = [[2,1],[1,2],[1,1]], stock = [1,1], cost = [5,5]
// Output: 2

    fmt.Println(maxNumberOfAlloys(
        2,3,10,
        [][]int{[]int{2,1},[]int{1,2},[]int{1,1}},
        []int{1,1},
        []int{5,5},
    )) // 2

    fmt.Println(maxNumberOfAlloys1(
        3,2,15,
        [][]int{[]int{1,1,1},[]int{1,1,10}},
        []int{0,0,0},
        []int{1,2,3},
    )) // 2
    fmt.Println(maxNumberOfAlloys1(
        3,2,15,
        [][]int{[]int{1,1,1},[]int{1,1,10}},
        []int{0,0,100},
        []int{1,2,3},
    )) // 5
    fmt.Println(maxNumberOfAlloys1(
        2,3,10,
        [][]int{[]int{2,1},[]int{1,2},[]int{1,1}},
        []int{1,1},
        []int{5,5},
    )) // 2
}