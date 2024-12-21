package main 

// 2266. Count Number of Texts
// Alice is texting Bob using her phone. The mapping of digits to letters is shown in the figure below.
// <img src="https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png" />

// In order to add a letter, Alice has to press the key of the corresponding digit i times, 
// where i is the position of the letter in the key.
//     1. For example, to add the letter 's', Alice has to press '7' four times. 
//        Similarly, to add the letter 'k', Alice has to press '5' twice.
//     2. Note that the digits '0' and '1' do not map to any letters, so Alice does not use them.

// However, due to an error in transmission, 
// Bob did not receive Alice's text message but received a string of pressed keys instead.
//     For example, when Alice sent the message "bob", Bob received the string "2266622".

// Given a string pressedKeys representing the string received by Bob, 
// return the total number of possible text messages Alice could have sent.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: pressedKeys = "22233"
// Output: 8
// Explanation:
// The possible text messages Alice could have sent are:
// "aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae", and "ce".
// Since there are 8 possible messages, we return 8.

// Example 2:
// Input: pressedKeys = "222222222222222222222222222222222222"
// Output: 82876089
// Explanation:
// There are 2082876103 possible text messages Alice could have sent.
// Since we need to return the answer modulo 109 + 7, we return 2082876103 % (109 + 7) = 82876089.

// Constraints:
//     1 <= pressedKeys.length <= 10^5
//     pressedKeys only consists of digits from '2' - '9'.

import "fmt"

func countTexts(pressedKeys string) int {
    n, mod := len(pressedKeys), 1_000_000_007
    dp := make([]int, n + 1)
    dp[0], dp[1] = 1, 1
    for i := 1; i < n; i++ {
        dp[i + 1] = dp[i]
        if i >= 1 && pressedKeys[i] == pressedKeys[i - 1] {
            dp[i + 1] = dp[i] + dp[i - 1]
            if i >= 2 && pressedKeys[i] == pressedKeys[i - 2] {
                dp[i + 1] = dp[i] + dp[i - 1] + dp[i - 2]
                if i >= 3 && (string(pressedKeys[i]) == "7" || string(pressedKeys[i]) == "9") {
                    if pressedKeys[i] == pressedKeys[i - 3] {
                        dp[i + 1] = dp[i] + dp[i - 1] + dp[i - 2] + dp[i - 3]
                    }
                }
            }
        }
        dp[i + 1] %= mod
    }
    return dp[n] 
}

func countTexts1(pressedKeys string) int {
    mod := 1_000_000_007
    num1 := [100_001]int{1, 1, 2, 4}
    num2 := num1
    init := func() {
        for i := 4; i < 100_001; i++ {
            num1[i] = (num1[i-1] + num1[i-2] + num1[i-3]) % mod
            num2[i] = (num2[i-1] + num2[i-2] + num2[i-3] + num2[i-4]) % mod
        }
    }
    init()
    res, count, n := 1, 0, len(pressedKeys)
    for i, c := range pressedKeys {
        count++
        if i == n - 1 || pressedKeys[i+1] != byte(c) { // 找到一个完整的组
            if c != '7' && c != '9' {
                res = res * num1[count] % mod
            } else {
                res = res * num2[count] % mod
            }
            count = 0
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: pressedKeys = "22233"
    // Output: 8
    // Explanation:
    // The possible text messages Alice could have sent are:
    // "aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae", and "ce".
    // Since there are 8 possible messages, we return 8.
    fmt.Println(countTexts("22233")) // 8
    // Example 2:
    // Input: pressedKeys = "222222222222222222222222222222222222"
    // Output: 82876089
    // Explanation:
    // There are 2082876103 possible text messages Alice could have sent.
    // Since we need to return the answer modulo 109 + 7, we return 2082876103 % (109 + 7) = 82876089.
    fmt.Println(countTexts("222222222222222222222222222222222222")) // 882876089
    fmt.Println(countTexts("344644885")) // 8

    fmt.Println(countTexts1("22233")) // 8
    fmt.Println(countTexts1("222222222222222222222222222222222222")) // 882876089
    fmt.Println(countTexts1("344644885")) // 8
}