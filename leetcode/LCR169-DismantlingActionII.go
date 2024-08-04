package main 

// LCR 169. 招式拆解 II
// 某套连招动作记作仅由小写字母组成的序列 arr，其中 arr[i] 第 i 个招式的名字。
// 请返回第一个只出现一次的招式名称，如不存在请返回空格。

// 示例 1：
// 输入：arr = "abbccdeff"
// 输出：'a'

// 示例 2：
// 输入：arr = "ccdd"
// 输出：' '

// 限制：
//     0 <= arr.length <= 50000

import "fmt"

func dismantlingAction(arr string) byte {
    seen := [26]int{}
    for i := range arr {
        seen[arr[i]-'a']++
    }
    for i := range arr {
        if seen[arr[i]-'a'] == 1 {
            return arr[i]
        }
    }
    return ' '
}

func main() {
    // 示例 1：
    // 输入：arr = "abbccdeff"
    // 输出：'a'
    fmt.Printf("%c\n", dismantlingAction("abbccdeff")) // a
    // 示例 2：
    // 输入：arr = "ccdd"
    // 输出：' '
    fmt.Printf("%c\n", dismantlingAction("ccdd")) // ' '
}