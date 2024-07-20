package main

// LCR 167. 招式拆解 I
// 某套连招动作记作序列 arr，其中 arr[i] 为第 i 个招式的名字。请返回 arr 中最多可以出连续不重复的多少个招式。

// 示例 1:
// 输入: arr = "dbascDdad"
// 输出: 6
// 解释: 因为连续且最长的招式序列是 "dbascD" 或 "bascDd"，所以其长度为 6。

// 示例 2:
// 输入: arr = "KKK"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "K"，所以其长度为 1。

// 示例 3:
// 输入: arr = "pwwkew"
// 输出: 3
// 解释: 因为连续且最长的招式序列是 "wke"，所以其长度为 3。     
// 请注意区分 子串 与 子序列 的概念：你的答案必须是 连续招式 的长度，也就是 子串。而 "pwke" 是一个非连续的 子序列，不是 子串。

// 提示：
//     0 <= arr.length <= 40000
//     arr 由英文字母、数字、符号和空格组成。

import "fmt"
import "strings"

func dismantlingAction(s string) int {
    if 0 == len(s) {  return 0 }
    l, sl, m := 0, 0, "" // 子串
    for i := 0; i < len(s); i++ {
        // 判断当前子符是否存在子串里
        if strings.Index(m, string(s[i])) == -1 {
            l++
            m += string(s[i])
        } else {
            if sl < l {
                sl = l
            }
            // 返回到s[i]之后开始的字符串
            t := s[i]
            for {
                i--
                if t == s[i] {
                    m = string(s[i+1])
                    i++
                    break
                }
            }
            l = 1
        }
    }
    // 如果最后一段字符是最长的
    if sl < l {
        sl = l
    }
    return sl
}

// best speed solution
func dismantlingAction1(s string) int {
    // index 初始一个list
    index, res, start, tmp := [128]int{}, 0, 0, 0
    for i, j := range s {
        //fmt.Println(i, j)
        if start < index[j] { // 如果
            start = index[j] //
        }
        tmp = i - start + 1 //
        if res < tmp {
            res = tmp
        }
        index[j] = i + 1
    }
    return res
}

// 位图
func dismantlingAction2(s string) int {
    if len(s) == 0 {
        return 0
    }
    bitSet := [256]bool{}
    res, left, right := 0, 0, 0
    for left < len(s) {
        // 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
        if bitSet[s[right]] {
            bitSet[s[left]] = false
            left++
        } else {
            bitSet[s[right]] = true
            right++
        }
        if res < right-left {
            res = right - left
        }
        if left + res >= len(s) || right >= len(s) {
            break
        }
    }
    return res
}

// 滑动窗口
func dismantlingAction3(s string) int {
    if len(s) == 0 {
        return 0
    }
    freq := [127]int{}
    res, left, right := 0, 0, -1
    for left < len(s) {
        if right+1 < len(s) && freq[s[right+1]] == 0 {
            freq[s[right+1]]++
            right++

        } else {
            freq[s[left]]--
            left++
        }
        res = max(res, right-left+1)
    }
    return res
}

// 滑动窗口-哈希桶
func dismantlingAction4(s string) int {
    right, left, res := 0, 0, 0
    m := make(map[byte]int, len(s))
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for left < len(s) {
        if index, ok := m[s[left]]; ok && index >= right {
            right = index + 1
        }
        m[s[left]] = left
        left++
        res = max(res, left - right)
    }
    return res
}

func main() {
    // Explanation: The answer is "abc", with the length of 3.
    fmt.Println(dismantlingAction("abcabcbb")) // 3
    // Explanation: The answer is "b", with the length of 1.
    fmt.Println(dismantlingAction("bbbbb")) // 1
    // Explanation: The answer is "wke", with the length of 3.
    // Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
    fmt.Println(dismantlingAction("pwwkew")) // 3

    fmt.Println(dismantlingAction1("abcabcbb")) // 3
    fmt.Println(dismantlingAction1("bbbbb")) // 1
    fmt.Println(dismantlingAction1("pwwkew")) // 3

    fmt.Println(dismantlingAction2("abcabcbb")) // 3
    fmt.Println(dismantlingAction2("bbbbb")) // 1
    fmt.Println(dismantlingAction2("pwwkew")) // 3

    fmt.Println(dismantlingAction3("abcabcbb")) // 3
    fmt.Println(dismantlingAction3("bbbbb")) // 1
    fmt.Println(dismantlingAction3("pwwkew")) // 3

    fmt.Println(dismantlingAction4("abcabcbb")) // 3
    fmt.Println(dismantlingAction4("bbbbb")) // 1
    fmt.Println(dismantlingAction4("pwwkew")) // 3
 
    fmt.Println(dismantlingAction("dvdf"))     // 3
    fmt.Println(dismantlingAction("aac"))      // 2
    fmt.Println(dismantlingAction("abc"))      // 3
    fmt.Println(dismantlingAction(""))         // 0
    fmt.Println(dismantlingAction("a"))        // 1
    fmt.Println(dismantlingAction("abcabcbb")) // 3
    fmt.Println(dismantlingAction("bbbbb"))    // 1
    fmt.Println(dismantlingAction("pwwkew"))   // 3

    fmt.Println()

    fmt.Println(dismantlingAction1("dvdf"))     // 3
    fmt.Println(dismantlingAction1("aac"))      // 2
    fmt.Println(dismantlingAction1("abc"))      // 3
    fmt.Println(dismantlingAction1(""))         // 0
    fmt.Println(dismantlingAction1("a"))        // 1
    fmt.Println(dismantlingAction1("abcabcbb")) // 3
    fmt.Println(dismantlingAction1("bbbbb"))    // 1
    fmt.Println(dismantlingAction1("pwwkew"))   // 3

    fmt.Printf("dismantlingAction2(\"abcabcbb\") = %v\n", dismantlingAction2("abcabcbb")) // 3
    fmt.Printf("dismantlingAction2(\"bbbbb\") = %v\n", dismantlingAction2("bbbbb")) // 1
    fmt.Printf("dismantlingAction2(\"abcabcbb\") = %v\n", dismantlingAction2("pwwkew")) // 3

    fmt.Printf("dismantlingAction3(\"abcabcbb\") = %v\n", dismantlingAction3("abcabcbb")) // 3
    fmt.Printf("dismantlingAction3(\"bbbbb\") = %v\n", dismantlingAction3("bbbbb")) // 1
    fmt.Printf("dismantlingAction3(\"abcabcbb\") = %v\n", dismantlingAction3("pwwkew")) // 3

    fmt.Printf("dismantlingAction4(\"abcabcbb\") = %v\n", dismantlingAction4("abcabcbb")) // 3
    fmt.Printf("dismantlingAction4(\"bbbbb\") = %v\n", dismantlingAction4("bbbbb")) // 1
    fmt.Printf("dismantlingAction4(\"abcabcbb\") = %v\n", dismantlingAction4("pwwkew")) // 3
}