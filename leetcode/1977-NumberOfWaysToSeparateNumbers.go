package main

// 1977. Number of Ways to Separate Numbers
// You wrote down many positive integers in a string called num. 
// However, you realized that you forgot to add commas to seperate the different numbers. 
// You remember that the list of integers was non-decreasing and that no integer had leading zeros.

// Return the number of possible lists of integers that you could have written down to get the string num. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: num = "327"
// Output: 2
// Explanation: You could have written down the numbers:
// 3, 27
// 327

// Example 2:
// Input: num = "094"
// Output: 0
// Explanation: No numbers can have leading zeros and all numbers must be positive.

// Example 3:
// Input: num = "0"
// Output: 0
// Explanation: No numbers can have leading zeros and all numbers must be positive.

// Constraints:
//     1 <= num.length <= 3500
//     num consists of digits '0' through '9'.

import "fmt"
import "sort"

// DP + Binary Search + Prefix Sum
func numberOfCombinations(num string) int {
    n, mod := len(num), 1_000_000_007
    last, prefix := make([][]int, n + 1), make([][]int, n + 1)
    last[0], prefix[0] = []int{ 0 }, []int{0, 1}
    compare := func (a, b string) bool {
        if len(a) != len(b) { return len(a) > len(b) }
        for i := 0; i < len(a); i++ {
            if a[i] == b[i] { continue }
            return a[i] > b[i]
        }
        return false
    }
    for end := 1; end <= len(num); end++ {
        prefix[end] = append(prefix[end], 0)
        index := 0
        for curStart := end - 1; curStart >= 0; curStart-- {
            if num[curStart] == '0' { continue }
            val := num[curStart:end]
            i := sort.Search(len(last[curStart]), func(j int) bool {
                prevStart := last[curStart][j]
                return compare(num[prevStart:curStart], val)
            })
            if count := prefix[curStart][i]; count > 0 {
                last[end] = append(last[end], curStart)
                prefix[end] = append(prefix[end], prefix[end][index] + count)
                prefix[end][index] %= mod
                index++
            }
        }
    }
    return prefix[n][len(prefix[n]) - 1] % mod
}

// 超出时间限制 184 / 257 
func numberOfCombinations1(num string) int {
    n, mod := len(num), 1_000_000_007
    var dp, lcp [3501][3501]int
    for i := 0; i <= n; i++ {
        lcp[i][n] = 0
    }
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= i; j-- {
            lcp[i][j] = 0
            if num[i] == num[j] {
                lcp[i][j] = 1 + lcp[i+1][j+1]
            }
        }
    }
    checkSmall := func(l1, l2, l int) bool {
        c := lcp[l1][l2]
        if c >= l { return true }
        return num[l1 + c] < num[l2 + c]
    }
    for i := 0; i < n; i++ {
        for l := 1; l <= i+1; l++ {
            dp[i][l] = 0
            j := i - l + 1
            if num[j] != '0' {
                if j == 0 {
                    dp[i][l] = 1
                } else {
                    maxL2 := 0
                    if j < l {
                        maxL2 = j
                    } else if checkSmall(j-l, j, l) {
                        maxL2 = l
                    } else {
                        maxL2 = l - 1
                    }
                    dp[i][l] = dp[j-1][maxL2]
                }
            }
            dp[i][l] = (dp[i][l-1] + dp[i][l]) % mod
        }
    }
    return dp[n-1][n]
}

func numberOfCombinations2(num string) int {
    n, mod := len(num), 1_000_000_007
    lcp, dp := make([][]int, n+1), make([][]int, n+1)
    for i := range lcp {
        lcp[i], dp[i] = make([]int, n+1), make([]int, n+1)
    }
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if num[i] == num[j] {
                lcp[i][j] = 1 + lcp[i+1][j+1]
            }
        }
    }
    compare := func(i, j, k int) bool {
        x := lcp[i][j]
        return x >= k || num[i+x] >= num[j+x]
    }
    dp[0][0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            v := 0
            if num[i-j] != '0' {
                if i-j-j >= 0 && compare(i-j, i-j-j, j) {
                    v = dp[i-j][j]
                } else {
                    v = dp[i-j][min(j-1, i-j)]
                }
            }
            dp[i][j] = (dp[i][j-1] + v) % mod
        }
    }
    return dp[n][n]
}

func main() {
    // Example 1:
    // Input: num = "327"
    // Output: 2
    // Explanation: You could have written down the numbers:
    // 3, 27
    // 327
    fmt.Println(numberOfCombinations("327")) // 2
    // Example 2:
    // Input: num = "094"
    // Output: 0
    // Explanation: No numbers can have leading zeros and all numbers must be positive.
    fmt.Println(numberOfCombinations("094")) // 0
    // Example 3:
    // Input: num = "0"
    // Output: 0
    // Explanation: No numbers can have leading zeros and all numbers must be positive.
    fmt.Println(numberOfCombinations("0")) // 0
    fmt.Println(numberOfCombinations("466183982622444721264647154379")) // 1090

    fmt.Println(numberOfCombinations1("327")) // 2
    fmt.Println(numberOfCombinations1("094")) // 0
    fmt.Println(numberOfCombinations1("0")) // 0
    fmt.Println(numberOfCombinations1("466183982622444721264647154379")) // 1090

    fmt.Println(numberOfCombinations2("327")) // 2
    fmt.Println(numberOfCombinations2("094")) // 0
    fmt.Println(numberOfCombinations2("0")) // 0
    fmt.Println(numberOfCombinations2("466183982622444721264647154379")) // 1090
}