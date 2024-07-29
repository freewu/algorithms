package main

// 1067. Digit Count in Range
// Given a single-digit integer d and two integers low and high, 
// return the number of times that d occurs as a digit in all integers in the inclusive range [low, high].

// Example 1:
// Input: d = 1, low = 1, high = 13
// Output: 6
// Explanation: The digit d = 1 occurs 6 times in 1, 10, 11, 12, 13.
// Note that the digit d = 1 occurs twice in the number 11.

// Example 2:
// Input: d = 3, low = 100, high = 250
// Output: 35
// Explanation: The digit d = 3 occurs 35 times in 103,113,123,130,131,...,238,239,243.

// Constraints:
//     0 <= d <= 9
//     1 <= low <= high <= 2 * 10^8

import "fmt"
import "strconv"

func digitsCount(d int, low int, high int) int {
    var dfs func(int, int, bool, bool, [][]int, string, int) int
    // mask记录之前有多少个d，isLimitL、isLimitH记录上下界是否受到限制
    dfs = func(idx, mask int, isLimit, isNum bool, dp [][]int, str string, length int) int {
        if idx == length {
            return mask
        }
        if !isLimit && isNum && dp[idx][mask] != -1 {
            return dp[idx][mask]
        }
        down, up := 0, 9
        res := 0
        if !isNum {
            down = 1 //前面没填过数，当前位下界至少为1
            res += dfs(idx+1, mask, false, false, dp, str, length)
        }
        if isLimit {
            up = int(str[idx] - '0')
        }
        for i := down; i <= up; i++ {
            if i == d {
                res += dfs(idx+1, mask+1, isLimit && i == up, true, dp, str, length)
            } else {
                res += dfs(idx+1, mask, isLimit && i == up, true, dp, str, length)
            }
        }
        if !isLimit && isNum && dp[idx][mask] == -1 {
            dp[idx][mask] = res
        }
        return res
    }
    numCount := func(d int, num int) int {
        str := strconv.Itoa(num)
        length := len(str)
        dp := make([][]int, length)
        for i := 0; i < length; i++ {
            dp[i] = make([]int, 10)
            for j := 0; j < 10; j++ {
                dp[i][j] = -1
            }
        }
        return dfs(0, 0, true, false, dp, str, length)
    }
    return numCount(d, high) - numCount(d, low-1)
}

func main() {
    // Example 1:
    // Input: d = 1, low = 1, high = 13
    // Output: 6
    // Explanation: The digit d = 1 occurs 6 times in 1, 10, 11, 12, 13.
    // Note that the digit d = 1 occurs twice in the number 11.
    fmt.Println(digitsCount(1, 1, 13)) // 6
    // Example 2:
    // Input: d = 3, low = 100, high = 250
    // Output: 35
    // Explanation: The digit d = 3 occurs 35 times in 103,113,123,130,131,...,238,239,243.
    fmt.Println(digitsCount(3, 100, 250)) // 35
}