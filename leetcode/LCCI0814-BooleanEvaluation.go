package main

// 面试题 08.14. Boolean Evaluation LCCI
// Given a boolean expression consisting of the symbols 0 (false), 1 (true), & (AND), | (OR), and ^ (XOR), 
// and a desired boolean result value result, implement a function to count the number of ways of parenthesizing the expression such that it evaluates to result.

// Example 1:
// Input: s = "1^0|0|1", result = 0
// Output: 2
// Explanation: Two possible parenthesizing ways are:
// 1^(0|(0|1))
// 1^((0|0)|1)

// Example 2:
// Input: s = "0&0&0&1^1|0", result = 1
// Output: 10

// Note:
//     There are no more than 19 operators in s.

import "fmt"

// dp[i][j][0/1]，表示s[i:j+1]结果为 0/1 的括号方法种数
func countEval(s string, result int) int {
    dp := make([][][2]int, len(s))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([][2]int, len(s))
    }
    calc := func(arg1, arg2 int, operator byte) int {
        if operator == '&' { return arg1 & arg2 } 
        if operator == '|' { return arg1 | arg2 }
        return arg1 ^ arg2
    }
    // dp[0][len-1][result]即为答案，i需要倒序遍历，j正序遍历
    for i := len(s) - 1; i >= 0; i -= 2 {
        for j := i; j < len(s); j += 2 {
            if i == j {
                if s[i] == '0' {
                    dp[i][j][0]++
                } else {
                    dp[i][j][1]++
                }
                continue
            }
            // 操作符s[k]将s[i:j+1]分为s[i:k]和s[k+1:j+1]分为前后两块
            // 遍历各自dp[i][k-1][0/1]和dp[k+1][j][0/1]共四种情况，累计dp[i][j][0/1]的值
            for k := i + 1; k < j; k += 2 {
                for arg1 := 0; arg1 <= 1; arg1++ {
                    for arg2 := 0; arg2 <= 1; arg2++ {
                        if calc(arg1, arg2, s[k]) == 0 {
                            dp[i][j][0] += dp[i][k-1][arg1] * dp[k+1][j][arg2]
                        } else {
                            dp[i][j][1] += dp[i][k-1][arg1] * dp[k+1][j][arg2]
                        }
                    }
                }
            }
        }
    }
    return dp[0][len(s)-1][result]
}

func countEval1(s string, result int) int {
    n := len(s)
    dp := make([][][2]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][2]int, n)
        if i % 2 == 1 { continue }
        if s[i] == '0' {
            dp[i][i] = [2]int{ 1, 0 }
        } else {
            dp[i][i] = [2]int{ 0, 1 }
        }
    }
    for d := 2; d < n; d += 2 {
        for i := 0; i < n; i += 2 {
            j := i + d
            if j >= n { continue }
            count0, count1 := 0, 0
            for t := i + 1; t < j; t += 2 {
                left0, left1, right0, right1 := dp[i][t - 1][0], dp[i][t - 1][1], dp[t + 1][j][0], dp[t + 1][j][1]
                total := (left0 + left1) * (right0 + right1)
                switch s[t] {
                case '&':
                    count0 += total - left1 * right1
                    count1 += left1 * right1
                case '|':
                    count0 += left0 * right0
                    count1 += total - left0 * right0
                case '^':
                    count0 += left0 * right0 + left1 * right1
                    count1 += left0 * right1 + left1 * right0
                }
            }
            dp[i][j] = [2]int{ count0, count1 }
        }
    }
    if result == 0 {
        return dp[0][n-1][0]
    }
    return dp[0][n-1][1]
}

func main() {
    // Example 1:
    // Input: s = "1^0|0|1", result = 0
    // Output: 2
    // Explanation: Two possible parenthesizing ways are:
    // 1^(0|(0|1))
    // 1^((0|0)|1)
    fmt.Println(countEval("1^0|0|1", 0)) // 2
    // Example 2:
    // Input: s = "0&0&0&1^1|0", result = 1
    // Output: 10
    fmt.Println(countEval("0&0&0&1^1|0", 1)) // 10

    fmt.Println(countEval1("1^0|0|1", 0)) // 2
    fmt.Println(countEval1("0&0&0&1^1|0", 1)) // 10
}