package main

// LCR 165. 解密数字
// 现有一串神秘的密文 ciphertext，经调查，密文的特点和规则如下：
//     密文由非负整数组成
//     数字 0-25 分别对应字母 a-z

// 请根据上述规则将密文 ciphertext 解密为字母，并返回共有多少种解密结果。

// 示例 1:
// 输入: ciphertext = 216612
// 输出: 6
// 解释: 216612 解密后有 6 种不同的形式，分别是 "cbggbc"，"vggbc"，"vggm"，"cbggm"，"cqggbc" 和 "cqggm" 

// 提示：
//     0 <= ciphertext < 2^31

import "fmt"
import "strconv"

func crackNumber(ciphertext int) int {
    nums := strconv.Itoa(ciphertext)
    n := len(nums)

    dp := make([]int, n + 1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        if nums[i-2] == '1' || (nums[i-2] == '2' && nums[i-1] < '6') { // 10 - 26 的处理
            dp[i] = dp[i-1] + dp[i-2]
        } else {
            dp[i] = dp[i-1]
        }
    }
    return dp[n]
}

func crackNumber1(ciphertext int) int {
    text := strconv.Itoa(ciphertext)
    if len(text) == 1 { return 1 }
    n := len(text)
    dp := make([]int, n)
    dp[0] = 1
    if n >= 2 && text[0:2] >= "10" && text[0:2] <= "25" {
        dp[1] = 2
    } else if n >= 2 {
        dp[1] = 1
    }
    for i := 2; i < n; i++ {
        pre := text[i-1:i+1]
        if pre >= "10" && pre <= "25" {
            dp[i] = dp[i-2]+dp[i-1]
        } else {
            dp[i] = dp[i-1]
        }
    }
    return dp[n-1]
}

func main() {
    // 示例 1:
    // 输入: ciphertext = 216612
    // 输出: 6
    // 解释: 216612 解密后有 6 种不同的形式，分别是 "cbggbc"，"vggbc"，"vggm"，"cbggm"，"cqggbc" 和 "cqggm" 
    fmt.Println(crackNumber(216612)) // 6
    fmt.Println(crackNumber(1024)) // 4
    fmt.Println(crackNumber(2048)) // 2

    fmt.Println(crackNumber1(216612)) // 6
    fmt.Println(crackNumber(1024)) // 4
    fmt.Println(crackNumber(2048)) // 2
}