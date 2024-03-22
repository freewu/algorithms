package main

// 8. String to Integer (atoi)
// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).
// The algorithm for myAtoi(string s) is as follows:
//     1 Read in and ignore any leading whitespace.
//     2 Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
//     3 Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
//     4 Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
//     5 If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
//     6 Return the integer as the final result.

// Note:
//     Only the space character ' ' is considered a whitespace character.
//     Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.

// Constraints:
//     0 <= s.length <= 200
//     s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'

// Example 1:
// Input: s = "42"
// Output: 42
// Explanation: The underlined characters are what is read in, the caret is the current reader position.
//     Step 1: "42" (no characters read because there is no leading whitespace)^
//     Step 2: "42" (no characters read because there is neither a '-' nor '+')^
//     Step 3: "42" ("42" is read in)^
// The parsed integer is 42.
// Since 42 is in the range [-231, 231 - 1], the final result is 42.

// Example 2:
// Input: s = "   -42"
// Output: -42
// Explanation:
//     Step 1: "   -42" (leading whitespace is read and ignored)
//     Step 2: "   -42" ('-' is read, so the result should be negative)^
//     Step 3: "   -42" ("42" is read in) ^
//     The parsed integer is -42.
//     Since -42 is in the range [-231, 231 - 1], the final result is -42.

// Example 3:
// Input: s = "4193 with words"
// Output: 4193
// Explanation:
//     Step 1: "4193 with words" (no characters read because there is no leading whitespace)
//     Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
//     Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
// The parsed integer is 4193.
// Since 4193 is in the range [-231, 231 - 1], the final result is 4193.

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

func myAtoi1(s string) int {
    maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}
    for _, c := range s {
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
        digits = append(digits, int(c-48))
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
func myAtoiBest(s string) int {
    ans := int64(0)
    negative := int64(1)
    high := int64(2147483647)
    low := int64(-2147483648)
    digitOnly := false
    for _,char := range s {
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
            ans = ans * 10 + (negative * int64(char - '0'))
            if ans > high {
                return int(high)
            } else if ans < low {
                return int(low)
            }
        }
    }
    return int(ans)
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

    fmt.Printf("myAtoiBest(\"42\") = %v\n",myAtoiBest("42")) // 42
    fmt.Printf("myAtoiBest(\"    -42\") = %v\n",myAtoiBest("    -42")) // -42
    fmt.Printf("myAtoiBest(\"4193 with words\") = %v\n",myAtoiBest("4193 with words")) // 4193

    fmt.Printf("myAtoi3(\"42\") = %v\n",myAtoi3("42")) // 42
    fmt.Printf("myAtoi3(\"    -42\") = %v\n",myAtoi3("    -42")) // -42
    fmt.Printf("myAtoi3(\"4193 with words\") = %v\n",myAtoi3("4193 with words")) // 4193
}
