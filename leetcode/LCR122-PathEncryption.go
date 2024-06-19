package main

// LCR 122. 路径加密
// 假定一段路径记作字符串 path，其中以 "." 作为分隔符。现需将路径加密，加密方法为将 path 中的分隔符替换为空格 " "，请返回加密后的字符串。

// 示例 1：
// 输入：path = "a.aef.qerf.bb"
// 输出："a aef qerf bb"

// 限制：
//     0 <= path.length <= 10000

import "fmt"
import "strings"

func pathEncryption(path string) string {
    return strings.Join(strings.Split(path,".")," ")
}

// use replace
func pathEncryption1(s string) string {
    return strings.Replace(s,"."," ", -1)
}

func main() {
    // 示例 1：
    // 输入：path = "a.aef.qerf.bb"
    // 输出："a aef qerf bb"
    fmt.Println(pathEncryption("a.aef.qerf.bb")) // "a aef qerf bb"

    fmt.Println(pathEncryption1("a.aef.qerf.bb")) // "a aef qerf bb"
}