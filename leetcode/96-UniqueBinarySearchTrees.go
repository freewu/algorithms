package main 

// 96. Unique Binary Search Trees
// Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

// Example 1:
//     1       1           2               3       3
//      \       \         /  \            /       /
//       3       2       1    3          2       2
//       /        \                     /         \ 
//     2           3                   1            1
// <img src="https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg" />
// Input: n = 3
// Output: 5

// Example 2:
// Input: n = 1
// Output: 1

// Constraints:
//     1 <= n <= 19

// # 这题的解题思路是 DP。
// dp[n] 代表 1-n 个数能组成多少个不同的二叉排序树，F(i,n) 代表以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数。
// 由于题意，我们可以得到这个等式：
//      dp[n] = F(1,n) + F(2,n) + F(3,n) + …… + F(n,n) 
// 初始值 dp[0] = 1，dp[1] = 1。分析 dp 和 F(i,n) 的关系又可以得到下面这个等式:
//      F(i,n) = dp[i-1] * dp[n-i] 
// 举例，[1,2,3,4,…, i ,…,n-1,n]，以 i 为 根节点，
//      那么左半边 [1,2,3,……,i-1] 
//      右半边 [i+1,i+2,……,n-1,n] 
//      分别能组成二叉排序树的不同个数相乘，即为以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数，也即 F(i,n)。

// 由于二叉排序树本身的性质，右边的子树一定比左边的子树，值都要大。
// 所以这里只需要根节点把树分成左右，不需要再关心左右两边数字的大小，只需要关心数字的个数。

// 状态转移方程是 dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]，最终要求的结果是 dp[n]
import "fmt"

func numTrees(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        // dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]
        for j := 1; j <= i; j++ {
            dp[i] += dp[j - 1] * dp[i - j]
        }
    }
    return dp[n]
}

func main() {
    fmt.Println(numTrees(1)) // 1
    fmt.Println(numTrees(2)) // 2
    fmt.Println(numTrees(3)) // 5
    fmt.Println(numTrees(4)) // 14
    fmt.Println(numTrees(5)) // 42
}