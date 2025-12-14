package main

// 3776. Minimum Moves to Balance Circular Array
// You are given a circular array balance of length n, where balance[i] is the net balance of person i.

// In one move, a person can transfer exactly 1 unit of balance to either their left or right neighbor.

// Return the minimum number of moves required so that every person has a non-negative balance. If it is impossible, return -1.

// Note: You are guaranteed that at most 1 index has a negative balance initially.

// Example 1:
// Input: balance = [5,1,-4]
// Output: 4
// Explanation:
// One optimal sequence of moves is:
// Move 1 unit from i = 1 to i = 2, resulting in balance = [5, 0, -3]
// Move 1 unit from i = 0 to i = 2, resulting in balance = [4, 0, -2]
// Move 1 unit from i = 0 to i = 2, resulting in balance = [3, 0, -1]
// Move 1 unit from i = 0 to i = 2, resulting in balance = [2, 0, 0]
// Thus, the minimum number of moves required is 4.

// Example 2:
// Input: balance = [1,2,-5,2]
// Output: 6
// Explanation:
// One optimal sequence of moves is:
// Move 1 unit from i = 1 to i = 2, resulting in balance = [1, 1, -4, 2]
// Move 1 unit from i = 1 to i = 2, resulting in balance = [1, 0, -3, 2]
// Move 1 unit from i = 3 to i = 2, resulting in balance = [1, 0, -2, 1]
// Move 1 unit from i = 3 to i = 2, resulting in balance = [1, 0, -1, 0]
// Move 1 unit from i = 0 to i = 1, resulting in balance = [0, 1, -1, 0]
// Move 1 unit from i = 1 to i = 2, resulting in balance = [0, 0, 0, 0]
// Thus, the minimum number of moves required is 6.​​​

// Example 3:
// Input: balance = [-3,2]
// Output: -1
// Explanation:
// ​​​​​​​It is impossible to make all balances non-negative for balance = [-3, 2], so the answer is -1.

// Constraints:
//     1 <= n == balance.length <= 10^5
//     -10^9 <= balance[i] <= 10^9
//     There is at most one negative value in balance initially.

import "fmt"

func minMoves(balance []int) int64 {
    sum, negi := 0, -1
    for i, v := range balance {
        sum += v // 累加和
        if v < 0 { // 找到负数位置
            negi = i
        }
    }
    if sum < 0 { return -1 } // 总和必须非负
    if negi < 0 { return 0 } // 没有负数，无需操作

    res, n := 0, len(balance)
    need := -balance[negi]
    for dis := 1; ; dis++ { // 把与 negi 相距 dis 的数移到 negi
        s := balance[(negi - dis + n)%n] + balance[(negi + dis) % n]
        if s >= need {
            res += need * dis // need 个 1 移动 dis 次
            return int64(res)
        }
        res += s * dis // s 个 1 移动 dis 次
        need -= s
    }
}

func main() {
    // Example 1:
    // Input: balance = [5,1,-4]
    // Output: 4
    // Explanation:
    // One optimal sequence of moves is:
    // Move 1 unit from i = 1 to i = 2, resulting in balance = [5, 0, -3]
    // Move 1 unit from i = 0 to i = 2, resulting in balance = [4, 0, -2]
    // Move 1 unit from i = 0 to i = 2, resulting in balance = [3, 0, -1]
    // Move 1 unit from i = 0 to i = 2, resulting in balance = [2, 0, 0]
    // Thus, the minimum number of moves required is 4.
    fmt.Println(minMoves([]int{5,1,-4})) // 4
    // Example 2:
    // Input: balance = [1,2,-5,2]
    // Output: 6
    // Explanation:
    // One optimal sequence of moves is:
    // Move 1 unit from i = 1 to i = 2, resulting in balance = [1, 1, -4, 2]
    // Move 1 unit from i = 1 to i = 2, resulting in balance = [1, 0, -3, 2]
    // Move 1 unit from i = 3 to i = 2, resulting in balance = [1, 0, -2, 1]
    // Move 1 unit from i = 3 to i = 2, resulting in balance = [1, 0, -1, 0]
    // Move 1 unit from i = 0 to i = 1, resulting in balance = [0, 1, -1, 0]
    // Move 1 unit from i = 1 to i = 2, resulting in balance = [0, 0, 0, 0]
    // Thus, the minimum number of moves required is 6.​​​
    fmt.Println(minMoves([]int{1,2,-5,2})) // 6
    // Example 3:
    // Input: balance = [-3,2]
    // Output: -1
    // Explanation:
    // ​​​​​​​It is impossible to make all balances non-negative for balance = [-3, 2], so the answer is -1. 
    fmt.Println(minMoves([]int{-3,2})) // -1

    fmt.Println(minMoves([]int{-1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(minMoves([]int{9,8,7,6,5,4,3,2,-1})) // 1
}