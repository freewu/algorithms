package main

// LCR 192. 把字符串转换成整数 (atoi)
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
// 函数 myAtoi(string s) 的算法如下：
//     读入字符串并丢弃无用的前导空格
//     检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//     读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//     将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//     如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。

// 返回整数作为最终结果。

// 注意：
//     本题中的空白字符只包括空格字符 ' ' 。
//     除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

// 示例 1：
// 输入：s = "42"
// 输出：42
// 解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
// 第 1 步："42"（当前没有读入字符，因为没有前导空格）
//          ^
// 第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//          ^
// 第 3 步："42"（读入 "42"）
//            ^
// 解析得到整数 42 。
// 由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。

// 示例 2：
// 输入：s = "   -42"
// 输出：-42
// 解释：
// 第 1 步："   -42"（读入前导空格，但忽视掉）
//             ^
// 第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
//              ^
// 第 3 步："   -42"（读入 "42"）
//                ^
// 解析得到整数 -42 。
// 由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。

// 示例 3：
// 输入：s = "4193 with words"
// 输出：4193
// 解释：
// 第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
//          ^
// 第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//          ^
// 第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
//              ^
// 解析得到整数 4193 。
// 由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。

// 提示：
//     0 <= s.length <= 200
//     s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成

import "fmt"
import "math"
import "strings"

func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    var l = len(str)
    if l == 0 {
        return 0
    }
    var flag = 1 // 正负符号
    var s = 0

    switch true {
    case '-' == str[0]:
        flag = -1
    case '+' == str[0]:
        s = 0
    case str[0] >= 48 && str[0] <= 57:
        s = int(str[0]) - 48
    default:
        return 0
    }

    for i := 1; i < l; i++ {
        if str[i] > 57 || str[i] < 48 {
            break
        }
        s = s*10 + (int(str[i]) - 48)

        if s > math.MaxInt32 {
            if -1 == flag {
                s = math.MinInt32
                flag = 1
            } else {
                s = math.MaxInt32
            }
            break
        }
    }
    return flag * s
}

func myAtoi1(str string) int {
    maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2 << 30), true, true, 1, []int{}
    for _, c := range str {
        if c == ' ' && whitespaceAllowed {
            continue
        }
        if signAllowed {
            if c == '+' {
                signAllowed = false
                whitespaceAllowed = false
                continue
            } else if c == '-' {
                sign = -1
                signAllowed = false
                whitespaceAllowed = false
                continue
            }
        }
        if c < '0' || c > '9' {
            break
        }
        whitespaceAllowed, signAllowed = false, false
        digits = append(digits, int(c - 48))
    }
    var num, place int64
    place, num = 1, 0
    lastLeading0Index := -1
    for i, d := range digits {
        if d == 0 {
            lastLeading0Index = i
        } else {
            break
        }
    }
    if lastLeading0Index > -1 {
        digits = digits[lastLeading0Index+1:]
    }
    var rtnMax int64
    if sign > 0 {
        rtnMax = maxInt - 1
    } else {
        rtnMax = maxInt
    }
    digitsCount := len(digits)
    for i := digitsCount - 1; i >= 0; i-- {
        num += int64(digits[i]) * place
        place *= 10
        if digitsCount-i > 10 || num > rtnMax {
            return int(int64(sign) * rtnMax)
        }
    }
    num *= int64(sign)
    return int(num)
}

// best solution
func myAtoi2(str string) int {
    res, negative, high, low, digitOnly := int64(0), int64(1), int64(1 << 32 -1), int64(-1 << 32 - 1), false
    for _,char := range str {
        if char == ' ' && digitOnly == false {
            continue
        } else if char == '-' && digitOnly == false {
            negative = -1
            digitOnly = true
        } else if char == '+' && digitOnly == false {
            negative = 1
            digitOnly = true
        } else if char < '0' || char > '9' {
            break
        } else {
            digitOnly = true
            res = res * 10 + (negative * int64(char - '0'))
            if res > high {
                return int(high)
            } else if res < low {
                return int(low)
            }
        }
    }
    return int(res)
}

func myAtoi3(s string) int {
    res, i, sign := 0, 0, 1
    for i < len(s) && s[i] == ' ' {
        i++
    }
    if i < len(s) {
        if s[i] == '-' {
            sign = -1
            i++
        } else if s[i] == '+' {
            i++
        }
    }
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        num := int(s[i] - '0')
        if res*10 + num > math.MaxInt32 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        res = res*10 + num
        i++
    }
    return res*sign
}

func main() {
    fmt.Println(myAtoi("aaaa"))
    fmt.Println(myAtoi("1a1aa"))
    fmt.Println(myAtoi("12345"))

    fmt.Println(myAtoi("-12345"))
    fmt.Println(myAtoi("+12345"))

    fmt.Println(myAtoi("   +12345"))

    fmt.Println(myAtoi("2147483648"))          // 2147483647
    fmt.Println(myAtoi("9223372036854775809")) // 2147483647

    fmt.Printf("myAtoi(\"42\") = %v\n",myAtoi("42")) // 42
    fmt.Printf("myAtoi(\"    -42\") = %v\n",myAtoi("    -42")) // -42
    fmt.Printf("myAtoi(\"4193 with words\") = %v\n",myAtoi("4193 with words")) // 4193

    fmt.Printf("myAtoi1(\"42\") = %v\n",myAtoi1("42")) // 42
    fmt.Printf("myAtoi1(\"    -42\") = %v\n",myAtoi1("    -42")) // -42
    fmt.Printf("myAtoi1(\"4193 with words\") = %v\n",myAtoi1("4193 with words")) // 4193

    fmt.Printf("myAtoi2(\"42\") = %v\n",myAtoi2("42")) // 42
    fmt.Printf("myAtoi2(\"    -42\") = %v\n",myAtoi2("    -42")) // -42
    fmt.Printf("myAtoi2(\"4193 with words\") = %v\n",myAtoi2("4193 with words")) // 4193

    fmt.Printf("myAtoi3(\"42\") = %v\n",myAtoi3("42")) // 42
    fmt.Printf("myAtoi3(\"    -42\") = %v\n",myAtoi3("    -42")) // -42
    fmt.Printf("myAtoi3(\"4193 with words\") = %v\n",myAtoi3("4193 with words")) // 4193
}
