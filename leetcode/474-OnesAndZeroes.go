package main

// 474. Ones and Zeroes
// You are given an array of binary strings strs and two integers m and n.
// Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
// A set x is a subset of a set y if all elements of x are also elements of y.

// Example 1:
// Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
// Output: 4
// Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
// Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
// {"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.

// Example 2:
// Input: strs = ["10","0","1"], m = 1, n = 1
// Output: 2
// Explanation: The largest subset is {"0", "1"}, so the answer is 2.
 
// Constraints:
//     1 <= strs.length <= 600
//     1 <= strs[i].length <= 100
//     strs[i] consists only of digits '0' and '1'.
//     1 <= m, n <= 100

import "fmt"
import "strings"

func findMaxForm(strs []string, m int, n int) int {
    var dp = [101][101]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, str := range strs {
        cnt0 := strings.Count(str, "0")
        cnt1 := (len(str) - cnt0)
        for zero := m; zero >= cnt0; zero-- {
            for one := n; one >= cnt1; one-- {
                dp[zero][one] = max(dp[zero][one], dp[zero - cnt0][one - cnt1] + 1)
            }
        }
    }
    return dp[m][n]
}

func main() {
    // Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
    // Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
    // {"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.
    fmt.Println(findMaxForm([]string{"10","0001","111001","1","0"}, 5, 3)) // 4
    // Explanation: The largest subset is {"0", "1"}, so the answer is 2.
    fmt.Println(findMaxForm([]string{"10","0","1"}, 1, 1)) // 2
}