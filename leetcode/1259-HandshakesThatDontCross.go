package main

// 1259. Handshakes That Don't Cross
// You are given an even number of people numPeople that stand around a circle and each person shakes hands with someone else so that there are numPeople / 2 handshakes total.
// Return the number of ways these handshakes could occur such that none of the handshakes cross.
// Since the answer could be very large, return it modulo 10^9 + 7.

// Example 1:
// <img src ="https://assets.leetcode.com/uploads/2019/07/11/5125_example_2.png" />
// Input: numPeople = 4
// Output: 2
// Explanation: There are two ways to do it, the first way is [(1,2),(3,4)] and the second one is [(2,3),(4,1)].

// Example 2:
// <img src ="https://assets.leetcode.com/uploads/2019/07/11/5125_example_3.png" />
// Input: numPeople = 6
// Output: 5
 
// Constraints:
//     2 <= numPeople <= 1000
//     numPeople is even.

import "fmt"

func numberOfWays(numPeople int) int {
    mod := 1_000_000_007
    dp := map[int]int{0: 1, 2: 1}
    // i 是从2到n的偶数
    // 自下向上，计算不同人数时的握手方案
    var helper func(x int) int
    helper = func (x int) int {
        if _, ok := dp[x]; ok {
            return dp[x]
        }
        dp[x] = 0
        // 假设选最后一个编号为x的人作为起点，和第1个人开始握手（1,3,5...x-1)
        // 每次和第i个人握手，就将图分为两部分，左边剩下x-1个人，共有dp[x-1]种握手排列组合方案，
        // 右边有x-i-1个人，共有dp[x-i-1]种握手方案，所以当x个人的时候，跟第i个人这一次的握手的方案数为dp[i-1]*dp[x-i-1].
        // 遍历i，求和即可得到x人时，总的握手方案
        for i := 1; i < x; i += 2 { // 自增2是因为不能留下奇数个，保证两边剩下的也都能握手
            dp[x] = dp[x] + (helper(i-1) * helper(x-i-1)) % mod
        }
        return dp[x]
    }
    for i := 2; i <= numPeople; i += 2 {
        dp[i] = helper(i) % mod
    }
    return dp[numPeople]
}

func numberOfWays1(numPeople int) int {
    mod := 1_000_000_007
    dp := make([]int, numPeople / 2 + 1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= numPeople / 2; i++ {
        for j := 0; j < i; j++ {
            dp[i] = (dp[i] + (dp[j] * dp[i - j - 1]) % mod) % mod
        }
    }
    return dp[numPeople / 2]
}

func numberOfWays2(numPeople int) int {
    if numPeople < 2 { return 0 }
    pairCount, mod := numPeople / 2, 1_000_000_007
    dp := make([]int, pairCount + 1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= pairCount; i ++ {
        border := (i - 1) / 2
        for j := 0; j <= border; j ++ {
            if j == i - 1 - j {
                dp[i] += (dp[j] * dp[i - 1 - j]) % mod
            } else {
                dp[i] += (2 * dp[j] * dp[i - 1 - j]) % mod
            }
        }
        dp[i] = dp[i] % mod
    }
    return dp[pairCount]
}

func main() {
    // Example 1:
    // <img src ="https://assets.leetcode.com/uploads/2019/07/11/5125_example_2.png" />
    // Input: numPeople = 4
    // Output: 2
    // Explanation: There are two ways to do it, the first way is [(1,2),(3,4)] and the second one is [(2,3),(4,1)].
    fmt.Println(numberOfWays(4)) // 2
    // Example 2:
    // <img src ="https://assets.leetcode.com/uploads/2019/07/11/5125_example_3.png" />
    // Input: numPeople = 6
    // Output: 5
    fmt.Println(numberOfWays(6)) // 5

    fmt.Println(numberOfWays(2)) // 1
    fmt.Println(numberOfWays(8)) // 14
    fmt.Println(numberOfWays(999)) // 0
    fmt.Println(numberOfWays(1000)) // 591137401

    fmt.Println(numberOfWays1(4)) // 2
    fmt.Println(numberOfWays1(6)) // 5
    fmt.Println(numberOfWays1(2)) // 1
    fmt.Println(numberOfWays1(8)) // 14
    fmt.Println(numberOfWays1(999)) // 948528453
    fmt.Println(numberOfWays1(1000)) // 591137401

    fmt.Println(numberOfWays2(4)) // 2
    fmt.Println(numberOfWays2(6)) // 5
    fmt.Println(numberOfWays2(2)) // 1
    fmt.Println(numberOfWays2(8)) // 14
    fmt.Println(numberOfWays2(999)) // 948528453
    fmt.Println(numberOfWays2(1000)) // 591137401
}