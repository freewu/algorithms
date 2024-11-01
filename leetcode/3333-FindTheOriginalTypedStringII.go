package main

// 3333. Find the Original Typed String II
// Alice is attempting to type a specific string on her computer. 
// However, she tends to be clumsy and may press a key for too long, resulting in a character being typed multiple times.

// You are given a string word, which represents the final output displayed on Alice's screen.
// You are also given a positive integer k.

// Return the total number of possible original strings that Alice might have intended to type, if she was trying to type a string of size at least k.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: word = "aabbccdd", k = 7
// Output: 5
// Explanation:
// The possible strings are: "aabbccdd", "aabbccd", "aabbcdd", "aabccdd", and "abbccdd".

// Example 2:
// Input: word = "aabbccdd", k = 8
// Output: 1
// Explanation:
// The only possible string is "aabbccdd".

// Example 3:
// Input: word = "aaabbb", k = 3
// Output: 8

// Constraints:
//     1 <= word.length <= 5 * 10^5
//     word consists only of lowercase English letters.
//     1 <= k <= 2000

import "fmt"

func possibleStringCount(word string, k int) int {
    i, n, arr, mod := 0, len(word), []int{}, 1_000_000_007
    for i < n {
        j := i + 1
        for j < n && word[j] == word[j - 1] { j++ }
        arr = append(arr, j - i)
        i = j
    }
    m := len(arr)
    power := make([]int, m)
    power[m - 1] = arr[m - 1]
    for i := m - 2; i >= 0; i-- {
        power[i] = (power[i + 1] * arr[i]) % mod
    }
    if m >= k { return power[0] }
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, k - m + 1)
    }
    for i := 0; i < k - m + 1; i++ {
        if arr[m-1] + i + m > k {
            dp[m-1][i] = arr[m-1] - (k-m-i)
        }
    }
    for i := m - 2; i >= 0; i-- {
        sum := (dp[i+1][k-m] * arr[i]) % mod
        for j := k - m; j >= 0; j-- {
            sum += dp[i+1][j]
            if j + arr[i] > k - m {
                sum = (sum - dp[i + 1][k - m] + mod) % mod
            } else {
                sum = (sum - dp[i + 1][j + arr[i]] + mod) % mod
            }
            dp[i][j] = sum
        }
    }
    return dp[0][0]
}

func possibleStringCount1(word string, k int) int {
    if len(word) < k {  return 0 } // 无法满足要求
    res, count, mod := 1, 0, 1_000_000_007
    arr := []int{}
    for i := range word {
        count++
        if i == len(word)-1 || word[i] != word[i+1] {
            if count > 1 { // 如果 cnt = 1，这组字符串必选，无需参与计算
                if k > 0 {
                    arr = append(arr, count - 1)
                }
                res = res * count % mod
            }
            k-- // 注意这里把 k 减小了
            count = 0
        }
    }
    if k <= 0 { return res }
    dp := make([]int, k)
    dp[0] = 1
    for _, v := range arr {
        for j := 1; j < k; j++ { // 原地计算 dp 的前缀和
            dp[j] = (dp[j] + dp[j-1]) % mod
        }
        for j := k - 1; j > v; j-- { // 计算子数组和
            dp[j] -= dp[j-v-1]
        }
    }
    for _, v := range dp {
        res -= v
    }
    return (res % mod + mod) % mod // 保证结果非负
}

func main() {
    // Example 1:
    // Input: word = "aabbccdd", k = 7
    // Output: 5
    // Explanation:
    // The possible strings are: "aabbccdd", "aabbccd", "aabbcdd", "aabccdd", and "abbccdd".
    fmt.Println(possibleStringCount("aabbccdd", 7)) // 5
    // Example 2:
    // Input: word = "aabbccdd", k = 8
    // Output: 1
    // Explanation:
    // The only possible string is "aabbccdd".
    fmt.Println(possibleStringCount("aabbccdd", 8)) // 1
    // Example 3:
    // Input: word = "aaabbb", k = 3
    // Output: 8
    fmt.Println(possibleStringCount("aaabbb", 3)) // 8
    
    fmt.Println(possibleStringCount1("aabbccdd", 7)) // 5
    fmt.Println(possibleStringCount1("aabbccdd", 8)) // 1
    fmt.Println(possibleStringCount1("aaabbb", 3)) // 8
}