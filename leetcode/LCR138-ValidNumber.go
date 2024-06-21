package main

// LCR 138. 有效数字
// 有效数字（按顺序）可以分成以下几个部分：
//     若干空格
//     一个 小数 或者 整数
//     （可选）一个 'e' 或 'E' ，后面跟着一个 整数
//     若干空格

// 小数（按顺序）可以分成以下几个部分：
//     （可选）一个符号字符（'+' 或 '-'）
//     下述格式之一：
//     至少一位数字，后面跟着一个点 '.'
//     至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//     一个点 '.' ，后面跟着至少一位数字

// 整数（按顺序）可以分成以下几个部分：
//     （可选）一个符号字符（'+' 或 '-'）
//     至少一位数字

// 部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
// 部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
// 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。

// 示例 1：
// 输入：s = "0"
// 输出：true

// 示例 2：
// 输入：s = "e"
// 输出：false

// 示例 3：
// 输入：s = "."
// 输出：false

// 提示：
//     1 <= s.length <= 20
//     s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。

import "fmt"
import "strings"
import "unicode"

func validNumber1(s string) bool {
    s = strings.TrimSpace(s) // 移除前导和尾随空格
    index, length, hasDigits := 0, len(s), false
    if index < length && (s[index] == '+' || s[index] == '-') {
        index++
    }
    // 更精确地处理数字和小数点
    if index < length && unicode.IsDigit(rune(s[index])) {
        hasDigits = true
        for ; index < length && unicode.IsDigit(rune(s[index])); index++ {
        }
        if index < length && s[index] == '.' {
            index++
            for ; index < length && unicode.IsDigit(rune(s[index])); index++ {
            }
        }
    } else if index < length && s[index] == '.' {
        index++
        if index == length || !unicode.IsDigit(rune(s[index])) {
            return false
        }
        for ; index < length && unicode.IsDigit(rune(s[index])); index++ {
        }
        hasDigits = true
    } else {
        return false
    }
    // 更严格的科学记数法处理
    if index < length && (s[index] == 'e' || s[index] == 'E') {
        index++
        if index == length || (!unicode.IsDigit(rune(s[index])) && s[index] != '+' && s[index] != '-') {
            return false // 直接跟'e/E'后面没有数字或正负号，则返回false
        }
        if index < length && (s[index] == '+' || s[index] == '-') {
            index++
        }
        if index == length || !unicode.IsDigit(rune(s[index])) {
            return false // 'e/E'后面没有数字，则返回false
        }
        for ; index < length && unicode.IsDigit(rune(s[index])); index++ {
        }
        hasDigits = true
    }
    return hasDigits && index == length
}

func validNumber(s string) bool {
    index, length := 0, len(s)
    for index < length && s[index] == ' ' { // 去掉前面的空格
        index++
    }
    if index == length { // 都是空格
        return false
    }
    if s[index] == '+' || s[index] == '-' { // 符号
        index++
    }
    if index == length { // 只有 + / - 的情况
        return false
    }
    isDigit := func (b byte) bool {
        return b >= '0' && b <= '9'
    }
    isLetter := func (b byte) bool {
        return (b >= 'A' && b < 'Z') || (b >= 'a' && b <= 'z')
    }
    beforePoint, afterPoint, point := 0, 0, false
    for index < length && s[index] != ' ' && !isLetter(s[index]) {
        c := s[index]
        if c == '.' {
            if !point {
                point = true
            } else {
                return false
            }
        } else if isDigit(c) {
            if !point {
                beforePoint++
            } else {
                afterPoint++
            }
        } else {
            return false
        }
        index++
    }
    if beforePoint == 0 && afterPoint == 0 {
        return false
    }
    if index < length && s[index] == ' ' {
        for index < length {
            if s[index] != ' ' {
                return false
            }
            index++
        }
    }
    if index == length {
        return true
    }
    letter := s[index]
    if letter != 'E' && letter != 'e' {
        return false
    }
    index++
    if index < length && (s[index] == '+' || s[index] == '-') {
        index++
    }
    if index == length || s[index] == ' ' {
        return false
    }
    for index < length && s[index] != ' ' {
        if !isDigit(s[index]) {
            return false
        }
        index++
    }
    for index < length {
        if s[index] != ' ' {
            return false
        }
        index++
    }
    return true
}

func main() {
    // 示例 1：
    // 输入：s = "0"
    // 输出：true
    fmt.Println(validNumber("0")) // true
    // 示例 2：
    // 输入：s = "e"
    // 输出：false
    fmt.Println(validNumber("e")) // false
    // 示例 3：
    // 输入：s = "."
    // 输出：false
    fmt.Println(validNumber(".")) // false

    fmt.Println(validNumber1("0")) // true
    fmt.Println(validNumber1("e")) // false
    fmt.Println(validNumber1(".")) // falses
}