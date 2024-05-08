package main

// 465. Optimal Account Balancing
// You are given an array of transactions transactions where transactions[i] = [fromi, toi, amounti] indicates 
// that the person with ID = fromi gave amounti $ to the person with ID = toi.
// Return the minimum number of transactions required to settle the debt.

// Example 1:
// Input: transactions = [[0,1,10],[2,0,5]]
// Output: 2
// Explanation:
// Person #0 gave person #1 $10.
// Person #2 gave person #0 $5.
// Two transactions are needed. One way to settle the debt is person #1 pays person #0 and #2 $5 each.

// Example 2:
// Input: transactions = [[0,1,10],[1,0,1],[1,2,5],[2,0,5]]
// Output: 1
// Explanation:
// Person #0 gave person #1 $10.
// Person #1 gave person #0 $1.
// Person #1 gave person #2 $5.
// Person #2 gave person #0 $5.
// Therefore, person #1 only need to give person #0 $4, and all debt is settled.
 
// Constraints:
//     1 <= transactions.length <= 8
//     transactions[i].length == 3
//     0 <= fromi, toi < 12
//     fromi != toi
//     1 <= amounti <= 100

import "fmt"
import "math/bits"

func minTransfers(transactions [][]int) int {
    cnt := make([]int, 12) // 0 <= fromi, toi < 12 // 用一个数组 cnt 记录每个人的钱是多了还是少了
    for _, p := range transactions {
        cnt[p[0]] -= p[2]
        cnt[p[1]] += p[2]
    }
    m, inf := 1 << len(cnt), 1 << 32 - 1
    dp := make([]int, m) // 定义 dp[i] 表示把集合 i 的所有元素值调整为 0，所需要的最少的还钱次数
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        sum := 0
        for j, v := range cnt {
            sum += i >> j & 1 * v
        }
        if sum != 0 { // 如果集合 i 的元素值之和不为 0，则 dp[i] = ∞
            dp[i] = inf / 2 // 防止溢出
        } else { // 如果集合 i 的元素值之和为 0，dp[i] 至多为 ∣i∣−1。然后枚举 i 的所有子集 j 和对应的补集 ∁ij
            dp[i] = bits.OnesCount(uint(i)) - 1 // 用二进制表示集合，二进制的第 i 位为 1 表示 i 在集合中
            for j := (i - 1) & i; j > 0; j = (j - 1) & i {
                dp[i] = min(dp[i], dp[j] + dp[i^j])
            }
        }
    }
    return dp[m-1]
}

func minTransfers1(transactions [][]int) int {
    money, inf := make(map[int]int), 1 << 32 - 1
    for _, record := range transactions {
        money[record[0]] += record[2]
        money[record[1]] -= record[2]
    }
    n, debit := len(money), make([]int, 0)
    for _, v := range money {
        debit = append(debit, v)
    }
    dp := make([]int, 1 << n)
    for i := 1; i < 1 << n; i++ {
        sum, ones := 0, 0
        for j := 0; j < n; j++ {
            if (i >> j) & 1 == 1 {
                sum += debit[j]
                ones++
            }
        }
        if sum != 0 {
            dp[i] = inf / 2
        } else {
            dp[i] = ones - 1
            for j := (i - 1) & i; j > 0; j = (j - 1) & i {
                if dp[i] > dp[j] + dp[i ^ j] {
                    dp[i] = dp[j] + dp[i ^ j]
                }
            }
        }
    }
    return dp[(1 << n) - 1]
}

func main() {
    // Example 1:
    // Input: transactions = [[0,1,10],[2,0,5]]
    // Output: 2
    // Explanation:
    // Person #0 gave person #1 $10.
    // Person #2 gave person #0 $5.
    // Two transactions are needed. One way to settle the debt is person #1 pays person #0 and #2 $5 each.
    fmt.Println(minTransfers([][]int{{0,1,10},{2,0,5}})) // 2
    // Example 2:
    // Input: transactions = [[0,1,10],[1,0,1],[1,2,5],[2,0,5]]
    // Output: 1
    // Explanation:
    // Person #0 gave person #1 $10.
    // Person #1 gave person #0 $1.
    // Person #1 gave person #2 $5.
    // Person #2 gave person #0 $5.
    // Therefore, person #1 only need to give person #0 $4, and all debt is settled.
    fmt.Println(minTransfers([][]int{{0,1,10},{1,0,1},{1,2,5},{2,0,5}})) // 1

    fmt.Println(minTransfers1([][]int{{0,1,10},{2,0,5}})) // 2
    fmt.Println(minTransfers1([][]int{{0,1,10},{1,0,1},{1,2,5},{2,0,5}})) // 1
}