package main

// LCR 016. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。

// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3 
// 解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。

// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。

// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 示例 4:
// 输入: s = ""
// 输出: 0

// 提示：
//     0 <= s.length <= 5 * 10^4
//     s 由英文字母、数字、符号和空格组成

import "fmt"
import "strings"

func lengthOfLongestSubstring(s string) int {
    if 0 == len(s) {
        return 0
    }
    var l = 0
    var sl = 0
    var m = "" // 子串
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
            var t = s[i]
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
func lengthOfLongestSubstring1(s string) int {
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
func lengthOfLongestSubstring2(s string) int {
    if len(s) == 0 {
        return 0
    }
    var bitSet [256]bool
    result, left, right := 0, 0, 0
    for left < len(s) {
        // 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
        if bitSet[s[right]] {
            bitSet[s[left]] = false
            left++
        } else {
            bitSet[s[right]] = true
            right++
        }
        if result < right-left {
            result = right - left
        }
        if left+result >= len(s) || right >= len(s) {
            break
        }
    }
    return result
}

// 滑动窗口
func lengthOfLongestSubstring3(s string) int {
    if len(s) == 0 {
        return 0
    }
    var freq [127]int
    result, left, right := 0, 0, -1

    for left < len(s) {
        if right+1 < len(s) && freq[s[right+1]] == 0 {
            freq[s[right+1]]++
            right++

        } else {
            freq[s[left]]--
            left++
        }
        result = max(result, right-left+1)
    }
    return result
}

// 滑动窗口-哈希桶
func lengthOfLongestSubstring4(s string) int {
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
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
    // Explanation: The answer is "b", with the length of 1.
    fmt.Println(lengthOfLongestSubstring("bbbbb")) // 1
    // Explanation: The answer is "wke", with the length of 3.
    // Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
    fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3

    fmt.Println(lengthOfLongestSubstring1("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring1("bbbbb")) // 1
    fmt.Println(lengthOfLongestSubstring1("pwwkew")) // 3

    fmt.Println(lengthOfLongestSubstring2("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring2("bbbbb")) // 1
    fmt.Println(lengthOfLongestSubstring2("pwwkew")) // 3

    fmt.Println(lengthOfLongestSubstring3("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring3("bbbbb")) // 1
    fmt.Println(lengthOfLongestSubstring3("pwwkew")) // 3

    fmt.Println(lengthOfLongestSubstring4("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring4("bbbbb")) // 1
    fmt.Println(lengthOfLongestSubstring4("pwwkew")) // 3
 
    fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
    fmt.Println(lengthOfLongestSubstring("aac"))      // 2
    fmt.Println(lengthOfLongestSubstring("abc"))      // 3
    fmt.Println(lengthOfLongestSubstring(""))         // 0
    fmt.Println(lengthOfLongestSubstring("a"))        // 1
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
    fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3

    fmt.Println()

    fmt.Println(lengthOfLongestSubstring1("dvdf"))     // 3
    fmt.Println(lengthOfLongestSubstring1("aac"))      // 2
    fmt.Println(lengthOfLongestSubstring1("abc"))      // 3
    fmt.Println(lengthOfLongestSubstring1(""))         // 0
    fmt.Println(lengthOfLongestSubstring1("a"))        // 1
    fmt.Println(lengthOfLongestSubstring1("abcabcbb")) // 3
    fmt.Println(lengthOfLongestSubstring1("bbbbb"))    // 1
    fmt.Println(lengthOfLongestSubstring1("pwwkew"))   // 3

    fmt.Printf("lengthOfLongestSubstring2(\"abcabcbb\") = %v\n",lengthOfLongestSubstring2("abcabcbb")) // 3
    fmt.Printf("lengthOfLongestSubstring2(\"bbbbb\") = %v\n",lengthOfLongestSubstring2("bbbbb")) // 1
    fmt.Printf("lengthOfLongestSubstring2(\"abcabcbb\") = %v\n",lengthOfLongestSubstring2("pwwkew")) // 3

    fmt.Printf("lengthOfLongestSubstring3(\"abcabcbb\") = %v\n",lengthOfLongestSubstring3("abcabcbb")) // 3
    fmt.Printf("lengthOfLongestSubstring3(\"bbbbb\") = %v\n",lengthOfLongestSubstring3("bbbbb")) // 1
    fmt.Printf("lengthOfLongestSubstring3(\"abcabcbb\") = %v\n",lengthOfLongestSubstring3("pwwkew")) // 3

    fmt.Printf("lengthOfLongestSubstring4(\"abcabcbb\") = %v\n",lengthOfLongestSubstring4("abcabcbb")) // 3
    fmt.Printf("lengthOfLongestSubstring4(\"bbbbb\") = %v\n",lengthOfLongestSubstring4("bbbbb")) // 1
    fmt.Printf("lengthOfLongestSubstring4(\"abcabcbb\") = %v\n",lengthOfLongestSubstring4("pwwkew")) // 3
}