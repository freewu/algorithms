package main

// 3448. Count Substrings Divisible By Last Digit
// You are given a string s consisting of digits.

// Return the number of substrings of s divisible by their non-zero last digit.

// Note: A substring may contain leading zeros.

// Example 1:
// Input: s = "12936"
// Output: 11
// Explanation:
// Substrings "29", "129", "293" and "2936" are not divisible by their last digit. There are 15 substrings in total, so the answer is 15 - 4 = 11.

// Example 2:
// Input: s = "5701283"
// Output: 18
// Explanation:
// Substrings "01", "12", "701", "012", "128", "5701", "7012", "0128", "57012", "70128", "570128", and "701283" are all divisible by their last digit. Additionally, all substrings that are just 1 non-zero digit are divisible by themselves. Since there are 6 such digits, the answer is 12 + 6 = 18.

// Example 3:
// Input: s = "1010101010"
// Output: 25
// Explanation:
// Only substrings that end with digit '1' are divisible by their last digit. There are 25 such substrings.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of digits only.

import "fmt"

func countSubstrings(s string) int64 {
    mem := make([][][]int, 2)
    for i := range mem {
        mem[i] = make([][]int, 10)
        for j := range mem[i] {
            mem[i][j] = make([]int, 9)
        }
    }
    res, cur := 0, 1
    for i := 0; i < len(s); i++ {
        x := int(s[i] - '0')
        for j := 1; j < 10; j++ {
            for k := 0; k < 9; k++ {
                mem[cur][j][k] = 0
            }
        }
        for j := 1; j < 10; j++ {
            for k := 0; k < 9; k++ {
                mem[cur][j][(k*10+x)%j] += mem[1-cur][j][k]
            }
            mem[cur][j][x%j]++ // don't forget to add s[i]%j
        }
        res += mem[cur][x][0]
        cur = 1 - cur
    }
    return int64(res)
}

func countSubstrings1(s string) int64 {
    res, n := s , len(s)
    var ans int64 = 0
    for j := 0; j < n; j++ {
        ch := res[j]
        if ch == '1' || ch == '2' || ch == '5' {ans += int64(j + 1)}
    }
    mod3 := 0
    var freq3 [3]int64
    freq3[0] = 1 
    for j := 0; j < n; j++ {
        digit , old := int(res[j] - '0') , mod3
        mod3 = (mod3 + digit) % 3
        if res[j] == '3' {ans += freq3[old]}
        freq3[mod3]++
    }
    mod6 := 0
    var freq6 [3]int64
    freq6[0] = 1 
    for j := 0; j < n; j++ {
        digit , old := int(res[j] - '0') , mod6
        mod6 = (mod6*4 + digit) % 6
        if res[j] == '6' {ans += freq6[old%3]}
        freq6[mod6%3]++
    }
    mod9 := 0
    freq9 := make([]int64, 9)
    freq9[0] = 1
    for j := 0; j < n; j++ {
        digit , old := int(res[j] - '0') , mod9
        mod9 = (mod9 + digit) % 9
        if res[j] == '9' {ans += freq9[old]}
        freq9[mod9]++
    }
    mod4 := 0
    for j := 0; j < n; j++ {
        digit := int(res[j] - '0')
        mod4 = (mod4*10 + digit) % 4
        if res[j] == '4' {
            if j == 0 {
                ans += 1
            } else {
                if mod4 == 0 {
                    ans += int64(j) + 1
                } else {
                    ans += 1
                }
            }
        }
    }
    pre7 := 0
    var freq7 [6][7]int64
    freq7[0][0] = 1 
    invFactor := [6]int{1, 5, 4, 6, 2, 3}
    for m := 1; m <= n; m++ {
        j := m - 1
        digit := int(res[j] - '0')
        pre7 = (pre7*10 + digit) % 7
        if res[j] == '7' {
            for r := 0; r < 6; r++ {
                targetResidue := ((m % 6) - r + 6) % 6
                desired := (pre7 * invFactor[r]) % 7
                ans += freq7[targetResidue][desired]
            }
        }
        freq7[m%6][pre7]++
    }
    pre8Arr := make([]int, n+1)
    pre8Arr[0] = 0
    for j := 0; j < n; j++ {
        digit := int(res[j] - '0')
        pre8Arr[j+1] = (pre8Arr[j]*10 + digit) % 8
    }
    for j := 0; j < n; j++ {
        if res[j] == '8' {
            add := 1 
            if j >= 1 {
                if pre8Arr[j+1] == (4*pre8Arr[j-1])%8 {
                    add += 1
                }
            }
            if j >= 2 {
                if pre8Arr[j+1] == 0 {
                    add += j - 1
                }
            }
            ans += int64(add)
        }
    }
    return ans
}

func main() {
    // Example 1:
    // Input: s = "12936"
    // Output: 11
    // Explanation:
    // Substrings "29", "129", "293" and "2936" are not divisible by their last digit. There are 15 substrings in total, so the answer is 15 - 4 = 11.
    fmt.Println(countSubstrings("12936")) // 11
    // Example 2:
    // Input: s = "5701283"
    // Output: 18
    // Explanation:
    // Substrings "01", "12", "701", "012", "128", "5701", "7012", "0128", "57012", "70128", "570128", and "701283" are all divisible by their last digit. Additionally, all substrings that are just 1 non-zero digit are divisible by themselves. Since there are 6 such digits, the answer is 12 + 6 = 18.
    fmt.Println(countSubstrings("5701283")) // 18
    // Example 3:
    // Input: s = "1010101010"
    // Output: 25
    // Explanation:
    // Only substrings that end with digit '1' are divisible by their last digit. There are 25 such substrings.
    fmt.Println(countSubstrings("1010101010")) // 25

    fmt.Println(countSubstrings("999999999")) // 45
    fmt.Println(countSubstrings("123456789")) // 20
    fmt.Println(countSubstrings("987654321")) // 35

    fmt.Println(countSubstrings1("12936")) // 11
    fmt.Println(countSubstrings1("5701283")) // 18
    fmt.Println(countSubstrings1("1010101010")) // 25
    fmt.Println(countSubstrings1("999999999")) // 45
    fmt.Println(countSubstrings1("123456789")) // 20
    fmt.Println(countSubstrings1("987654321")) // 35
}