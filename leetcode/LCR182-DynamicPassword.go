package main

// LCR 182. 动态口令
// 某公司门禁密码使用动态口令技术。初始密码为字符串 password，密码更新均遵循以下步骤：
//     设定一个正整数目标值 target
//     将 password 前 target 个字符按原顺序移动至字符串末尾

// 请返回更新后的密码字符串。

// 示例 1：
// 输入: password = "s3cur1tyC0d3", target = 4
// 输出: "r1tyC0d3s3cu"

// 示例 2：
// 输入: password = "lrloseumgh", target = 6
// 输出: "umghlrlose"

// 提示：
//     1 <= target < password.length <= 10000

import "fmt"

func dynamicPassword(password string, target int) string {
    return string(password[target:] + password[:target]) 
}

func dynamicPassword1(password string, target int) string {
    n := len(password)
    res, k := make([]byte, n), 0
    for i := target; i < target + n; i++ {
        res[k] = password[i % n]
        k++
    }
    return string(res)
}

func main() {
    // 示例 1：
    // 输入: password = "s3cur1tyC0d3", target = 4
    // 输出: "r1tyC0d3s3cu"
    fmt.Println(dynamicPassword("s3cur1tyC0d3", 4)) // "r1tyC0d3s3cu"
    // 示例 2：
    // 输入: password = "lrloseumgh", target = 6
    // 输出: "umghlrlose"
    fmt.Println(dynamicPassword("lrloseumgh", 6)) // "umghlrlose"

    fmt.Println(dynamicPassword1("s3cur1tyC0d3", 4)) // "r1tyC0d3s3cu"
    fmt.Println(dynamicPassword1("lrloseumgh", 6)) // "umghlrlose"
}