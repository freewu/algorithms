package main

// 13. Roman to Integer
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//     Symbol       Value
//     I             1
//     V             5
//     X             10
//     L             50
//     C             100
//     D             500
//     M             1000
// For example, two is written as II in Roman numeral, just two one’s added together.
// Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
// Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

//     I can be placed before V (5) and X (10) to make 4 and 9.
//     X can be placed before L (50) and C (100) to make 40 and 90.
//     C can be placed before D (500) and M (1000) to make 400 and 900.

// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

// Example 1:
//     Input: "III"
//     Output: 3

// Example 2:
//     Input: "IV"
//     Output: 4

// Example 3:
//     Input: "IX"
//     Output: 9

// Example 4:
//     Input: "LVIII"
//     Output: 58
//     Explanation: L = 50, V= 5, III = 3.

// Example 5:
//     Input: "MCMXCIV"
//     Output: 1994
//     Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

// Constraints:
//     1 <= s.length <= 15
//     s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
//     It is guaranteed that s is a valid roman numeral in the range [1, 3999].

import "fmt"

func romanToInt(s string) int {
    m := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
    res := m[string(s[0])]
    for i := 1; i < len(s); i++ {
        if m[string(s[i])] > m[string(s[i - 1])] {
            // 1 + 5 - 2 * 1 IV
            res = res + m[string(s[i])] - 2 * m[string(s[i - 1])]
        } else {
            res += m[string(s[i])]
        }
    }
    return res
}

func romanToInt1(s string) int {
    if s == "" {
        return 0
    }
    m := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
    num, lastint, res := 0, 0, 0
    for i := 0; i < len(s); i++ {
        char := s[len(s)-(i+1) : len(s)-i]
        num = m[char]
        if num < lastint {
            res = res - num
        } else {
            res = res + num
        }
        lastint = num
    }
    return res
}

// best solution
func romanToInt2(s string)  int {
    k, res := []rune(s), 0
    m := map[rune]int{ 'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000 }
    for i:=0; i<len(k);i++{
        if i < len(k)-1 && m[k[i]] < m[k[i+1]] {
            res += m[k[i+1]]- m[k[i]]
            i = i+1
        } else {
            res += m[k[i]]
        }
    }
    return res
}

func main() {
    fmt.Printf("romanToInt(\"VII\") = %v\n",romanToInt("VII")) // 7
    fmt.Printf("romanToInt(\"IV\") = %v\n",romanToInt("IV")) // 4
    fmt.Printf("romanToInt(\"IX\") = %v\n",romanToInt("IX")) // 9
    fmt.Printf("romanToInt(\"VIII\") = %v\n",romanToInt("VIII")) // 8
    fmt.Printf("romanToInt(\"V\") = %v\n",romanToInt("V")) // 5
    fmt.Printf("romanToInt(\"XV\") = %v\n",romanToInt("XV")) // 15
    fmt.Printf("romanToInt(\"DCXXI\") = %v\n",romanToInt("DCXXI")) // 621

    fmt.Printf("romanToInt1(\"VII\") = %v\n",romanToInt1("VII"))
    fmt.Printf("romanToInt1(\"IV\") = %v\n",romanToInt1("IV"))
    fmt.Printf("romanToInt1(\"IX\") = %v\n",romanToInt1("IX"))
    fmt.Printf("romanToInt1(\"VIII\") = %v\n",romanToInt1("VIII"))
    fmt.Printf("romanToInt1(\"V\") = %v\n",romanToInt1("V"))
    fmt.Printf("romanToInt1(\"XV\") = %v\n",romanToInt1("XV"))
    fmt.Printf("romanToInt1(\"DCXXI\") = %v\n",romanToInt1("DCXXI")) // 621

    fmt.Printf("romanToInt2(\"VII\") = %v\n",romanToInt2("VII"))
    fmt.Printf("romanToInt2(\"IV\") = %v\n",romanToInt2("IV"))
    fmt.Printf("romanToInt2(\"IX\") = %v\n",romanToInt2("IX"))
    fmt.Printf("romanToInt2(\"VIII\") = %v\n",romanToInt2("VIII"))
    fmt.Printf("romanToInt2(\"V\") = %v\n",romanToInt2("V"))
    fmt.Printf("romanToInt2(\"XV\") = %v\n",romanToInt2("XV"))
    fmt.Printf("romanToInt2(\"DCXXI\") = %v\n",romanToInt2("DCXXI")) // 621
}