package main

// 3756. Concatenate Non-Zero Digits and Multiply by Sum II
// You are given a string s of length m consisting of digits. 
// You are also given a 2D integer array queries, where queries[i] = [li, ri].

// For each queries[i], extract the substring s[li..ri]. Then, perform the following:
//     1. Form a new integer x by concatenating all the non-zero digits from the substring in their original order. 
//        If there are no non-zero digits, x = 0.
//     2. Let sum be the sum of digits in x. The answer is x * sum.

// Return an array of integers answer where answer[i] is the answer to the ith query.

// Since the answers may be very large, return them modulo 10^9 + 7.

// Example 1:
// Input: s = "10203004", queries = [[0,7],[1,3],[4,6]]
// Output: [12340, 4, 9]
// Explanation:
// s[0..7] = "10203004"
// x = 1234
// sum = 1 + 2 + 3 + 4 = 10
// Therefore, answer is 1234 * 10 = 12340.
// s[1..3] = "020"
// x = 2
// sum = 2
// Therefore, the answer is 2 * 2 = 4.
// s[4..6] = "300"
// x = 3
// sum = 3
// Therefore, the answer is 3 * 3 = 9.

// Example 2:
// Input: s = "1000", queries = [[0,3],[1,1]]
// Output: [1, 0]
// Explanation:
// s[0..3] = "1000"
// x = 1
// sum = 1
// Therefore, the answer is 1 * 1 = 1.
// s[1..1] = "0"
// x = 0
// sum = 0
// Therefore, the answer is 0 * 0 = 0.

// Example 3:
// Input: s = "9876543210", queries = [[0,9]]
// Output: [444444137]
// Explanation:
// s[0..9] = "9876543210"
// x = 987654321
// sum = 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 = 45
// Therefore, the answer is 987654321 * 45 = 44444444445.
// We return 44444444445 modulo (109 + 7) = 444444137.

// Constraints:
//     1 <= m == s.length <= 10^5
//     s consists of digits only.
//     1 <= queries.length <= 10^5
//     queries[i] = [li, ri]
//     0 <= li <= ri < m

import "fmt"

func sumAndMultiply(s string, queries [][]int) []int {
    const MOD = 1_000_000_007
    n, q := len(s), len(queries)
    // 1. 预处理前缀数组
    // pow10[i] 存储 10^i mod MOD 的结果
    pow10 := make([]int64, n+1)
    pow10[0] = 1
    for i := 1; i <= n; i++ {
        pow10[i] = (pow10[i-1] * 10) % MOD
    }
    // nonZeroCount[i] 存储前 i 个字符中非零数字的个数
    nonZeroCount := make([]int, n+1)
    // prefixValue[i] 存储由前 i 个非零数字组成的数 mod MOD 的结果
    prefixValue := make([]int64, n+1)
    // prefixDigitSum[i] 存储前 i 个非零数字的总和
    prefixDigitSum := make([]int64, n+1)

    for i := 0; i < n; i++ {
        digit := s[i] - '0'
        // 默认继承前一个位置的状态
        nonZeroCount[i+1] = nonZeroCount[i]
        prefixValue[i+1] = prefixValue[i]
        prefixDigitSum[i+1] = prefixDigitSum[i]

        // 只有当当前字符是 non-zero 时，才更新所有前缀数组
        if digit != 0 {
            nonZeroCount[i+1]++
            // 更新数值：当前数值 = 之前的数值 * 10 + 新数字
            prefixValue[i+1] = (prefixValue[i]*10 + int64(digit)) % MOD
            // 更新数字和
            prefixDigitSum[i+1] += int64(digit)
        }
    }
    // 2. 处理每个查询
    res := make([]int, q)
    for i, query := range queries {
        L, R := query[0], query[1]
        // 计算 [L, R] 区间内非零数字的个数
        count := nonZeroCount[R+1] - nonZeroCount[L]
        // 计算 [L, R] 区间内非零数字的总和
        digitSum := prefixDigitSum[R+1] - prefixDigitSum[L]
        var x int64
        if count == 0 {
            // 如果没有非零数字，数值为 0
            x = 0
        } else {
            // 计算 [L, R] 区间内非零数字组成的数值 x
            // x = (prefixValue[R+1] - prefixValue[L] * 10^count) mod MOD
            leftContribution := (prefixValue[L] * pow10[count]) % MOD
            // (a - b) mod MOD 的正确计算方式，确保结果非负
            x = (prefixValue[R+1] - leftContribution + MOD) % MOD
        }
        // 计算最终结果 (x * digitSum) mod MOD
        res[i] = int((x * digitSum) % MOD)
    }
    return res
}

func sumAndMultiply1(s string, queries [][]int) []int {
    const MOD = 1_000_000_007
    n := len(s)
    sum, count, value, p10 := make([]int64, n+1), make([]int64, n+1), make([]int64, n+1), make([]int64, n+1)
    p10[0] = 1
    for i := 1; i <= n; i++ {
        tmp := int64(s[i-1] - '0')
        sum[i], count[i], value[i] = sum[i-1], count[i-1], value[i-1]
        p10[i] = (p10[i-1] * 10) % MOD
        if tmp != 0 {
            sum[i] += tmp
            count[i]++
            value[i] = (value[i]*10 + tmp) % MOD
        }
    }
    res := make([]int, len(queries))
    for i := 0; i < len(queries); i++ {
        l, r := queries[i][0], queries[i][1]
        v := sum[r+1] - sum[l]
        cn := count[r+1] - count[l]
        if cn == 0 {
            res[i] = 0
        } else {
            full := value[r+1] 
            pre := value[l]
            x := (full - (pre * p10[cn]) % MOD + MOD) % MOD
            res[i] = int((x * v) % MOD)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "10203004", queries = [[0,7],[1,3],[4,6]]
    // Output: [12340, 4, 9]
    // Explanation:
    // s[0..7] = "10203004"
    // x = 1234
    // sum = 1 + 2 + 3 + 4 = 10
    // Therefore, answer is 1234 * 10 = 12340.
    // s[1..3] = "020"
    // x = 2
    // sum = 2
    // Therefore, the answer is 2 * 2 = 4.
    // s[4..6] = "300"
    // x = 3
    // sum = 3
    // Therefore, the answer is 3 * 3 = 9.
    fmt.Println(sumAndMultiply("10203004", [][]int{{0,7},{1,3},{4,6}})) // [12340, 4, 9]
    // Example 2:
    // Input: s = "1000", queries = [[0,3],[1,1]]
    // Output: [1, 0]
    // Explanation:
    // s[0..3] = "1000"
    // x = 1
    // sum = 1
    // Therefore, the answer is 1 * 1 = 1.
    // s[1..1] = "0"
    // x = 0
    // sum = 0
    // Therefore, the answer is 0 * 0 = 0.
    fmt.Println(sumAndMultiply("1000", [][]int{{0,3},{1,1}})) // [1, 0]
    // Example 3:
    // Input: s = "9876543210", queries = [[0,9]]
    // Output: [444444137]
    // Explanation:
    // s[0..9] = "9876543210"
    // x = 987654321
    // sum = 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 = 45
    // Therefore, the answer is 987654321 * 45 = 44444444445.
    // We return 44444444445 modulo (109 + 7) = 444444137.
    fmt.Println(sumAndMultiply("9876543210", [][]int{{0,9}})) // [444444137]

    fmt.Println(sumAndMultiply("123456789", [][]int{{0,7},{1,3},{4,6}})) // [444444408 2106 10206]
    fmt.Println(sumAndMultiply("987654321", [][]int{{0,7},{1,3},{4,6}})) // [345678980 18396 6516]

    fmt.Println(sumAndMultiply1("10203004", [][]int{{0,7},{1,3},{4,6}})) // [12340, 4, 9]
    fmt.Println(sumAndMultiply1("1000", [][]int{{0,3},{1,1}})) // [1, 0]
    fmt.Println(sumAndMultiply1("9876543210", [][]int{{0,9}})) // [444444137]
    fmt.Println(sumAndMultiply1("123456789", [][]int{{0,7},{1,3},{4,6}})) // [444444408 2106 10206]
    fmt.Println(sumAndMultiply1("987654321", [][]int{{0,7},{1,3},{4,6}})) // [345678980 18396 6516]
}

