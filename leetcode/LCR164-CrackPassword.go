package main

// LCR 164. 破解闯关密码
// 闯关游戏需要破解一组密码，闯关组给出的有关密码的线索是：
//     一个拥有密码所有元素的非负整数数组 password
//     密码是 password 中所有元素拼接后得到的最小的一个数

// 请编写一个程序返回这个密码。

// 示例 1:
// 输入: password = [15, 8, 7]
// 输出: "1578"

// 示例 2:
// 输入: password = [0, 3, 30, 34, 5, 9]
// 输出: "03033459"

// 提示:
//     0 < password.length <= 100

// 说明:
// 输出结果可能非常大，所以你需要返回一个字符串而不是整数
// 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

import "fmt"
import "sort"
import "strconv"

func crackPassword(password []int) string {
    sort.Slice(password, func(m, n int) bool {
        mstr, nstr := strconv.Itoa(password[m]), strconv.Itoa(password[n])
        return mstr + nstr < nstr + mstr
    })
    res := ""
    for _, v := range password {
        res += strconv.Itoa(v)
    }
    return res
}

func main() {
    // 示例 1:
    // 输入: password = [15, 8, 7]
    // 输出: "1578"
    fmt.Println(crackPassword([]int{15, 8, 7})) // "1578"
    // 示例 2:
    // 输入: password = [0, 3, 30, 34, 5, 9]
    // 输出: "03033459"
    fmt.Println(crackPassword([]int{0, 3, 30, 34, 5, 9})) // "03033459"
}