package main

// 279. Perfect Squares
// Given an integer n, return the least number of perfect square numbers that sum to n.
// A perfect square is an integer that is the square of an integer; 
// in other words, it is the product of some integer with itself. 
// For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

// Example 1:
// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4.

// Example 2:
// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9.
 
// Constraints:
//     1 <= n <= 10^4

// 由拉格朗日的四平方定理可得，每个自然数都可以表示为四个整数平方之和。 其中四个数字是整数。
// 四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和。这给出了本题的答案的上界

import "fmt"
import "math"

func numSquares(n int) int {
    // 判断是否为完全平方数
    isPerfectSquare := func (n int) bool {
        sq := int(math.Floor(math.Sqrt(float64(n))))
        return sq * sq == n
    }
    // 四平方和定理可以推出三平方和推论：
    // 		当且仅当 n != 4^k*(8m+7) 时，n 可以被表示为至多三个正整数的平方和。
    // 		所以当 n = 4^k*(8m+7) 时，n 只能被表示为四个正整数的平方和。此时我们可以直接返回 4
    // 判断是否能表示为 4^k*(8m+7)
    checkAnswer4 := func (x int) bool {
        for x % 4 == 0 {
            x /= 4
        }
        return x % 8 == 7
    }
    if checkAnswer4(n) {
        return 4
    }
    // 当 n != 4^k*(8m+7) 时，需要判断 n 到底可以分解成几个完全平方数之和。答案肯定是 1，2，3 中的一个
    // 从 1 开始一个个判断是否满足。如果答案为 1，代表 n 为完全平方数
    if isPerfectSquare(n) {
        return 1
    }
    // 如果答案为 2，代表 n = a^2 + b^2 ，枚举 1 <= a <= sqrt(n) ，判断 n - a^2  是否为完全平方数
    for i := 1; i * i <= n; i++ {
        j := n - i * i
        if isPerfectSquare(j) {
            return 2
        }
    }
    return 3
}

// dp
func numSquares1(n int) int {
    dp := make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        val := math.MaxInt32
        for j := 1; j * j <= i; j++ {
            val = min(val, dp[i - j * j])
        }
        dp[i] = val + 1
    }
    return dp[n]
}

func main() {
    // Explanation: 12 = 4 + 4 + 4.
    fmt.Println(numSquares(12)) // 3
    // Explanation: 13 = 4 + 9.
    fmt.Println(numSquares(13)) // 2
    fmt.Println(numSquares(15)) // 4  (1 + 1 + 4 + 9)
    fmt.Println(numSquares(3)) // 3  (1 + 1 + 1)
    fmt.Println(numSquares(1)) // 1 （4）

    // Explanation: 12 = 4 + 4 + 4.
    fmt.Println(numSquares1(12)) // 3
    // Explanation: 13 = 4 + 9.
    fmt.Println(numSquares1(13)) // 2
    fmt.Println(numSquares1(15)) // 4 (1 + 1 + 4 + 9)
    fmt.Println(numSquares1(3)) // 3  (1 + 1 + 1)
    fmt.Println(numSquares1(1)) // 1 （4）
}