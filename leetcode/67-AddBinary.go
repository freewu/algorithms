package main

// 67. Add Binary
// Given two binary strings, return their sum (also a binary string).
// The input strings are both non-empty and contains only characters 1 or 0.

// Example 1:
// Input: a = "11", b = "1"
// Output: "100"

// Example 2:
// Input: a = "1010", b = "1011"
// Output: "10101"

// Constraints:
//     1 <= a.length, b.length <= 10^4
//     a and b consist only of '0' or '1' characters.
//     Each string does not contain leading zeros except for the zero itself.

import "fmt"
import "strings"
import "strconv"

func addBinary(a string, b string) string {
    var al = len(a) - 1
    var bl = len(b) - 1

    var s = ""
    var flag = 0 // 进位标识

    for {
        // 短的结束了
        if al < 0 || bl < 0 {
            break
        }
        if '1' == a[al] && '1' == b[bl] { // 1/1
            if 0 == flag {
                s = "0" + s
            } else {
                s = "1" + s
            }
            flag = 1
        } else if '0' == a[al] && '0' == b[bl] { // 0/0
            if 0 == flag {
                s = "0" + s
            } else {
                s = "1" + s
            }
            flag = 0
        } else { // 0/1 1/0的情况
            if 0 == flag {
                s = "1" + s
                flag = 0
            } else {
                s = "0" + s
                flag = 1
            }
        }
        al--
        bl--
    }
    // 连接其它的部分
    for {
        if al < 0 {
            break
        }
        if a[al] == '1' && 1 == flag {
            flag = 1
            s = "0" + s
        } else if a[al] == '0' && 0 == flag {
            flag = 0
            s = string(a[al]) + s
        } else {
            flag = 0
            s = "1" + s
        }
        al--
    }
    for {
        if bl < 0 {
            break
        }
        if b[bl] == '1' && 1 == flag {
            flag = 1
            s = "0" + s
        } else if b[bl] == '0' && 0 == flag {
            flag = 0
            s = string(b[bl]) + s
        } else {
            flag = 0
            s = "1" + s
        }
        bl--
    }

    if flag == 1 {
        s = "1" + s
    }

    return s
}

// best solution
func addBinary1(a string, b string) string {
    if len(b) > len(a) {
        a, b = b, a
    }
    res := make([]string, len(a)+1) // 声明一个数组来保存结果 长度是 最长的 + 1
    i, j, k, c := len(a)-1, len(b)-1, len(a), 0
    // 先把对齐的长度进行相加处理
    for i >= 0 && j >= 0 {
        ai, _ := strconv.Atoi(string(a[i]))
        bj, _ := strconv.Atoi(string(b[j]))
        res[k] = strconv.Itoa((ai + bj + c) % 2)
        c = (ai + bj + c) / 2 // 判断是否进位
        i--
        j--
        k--
    }
    // 处理超过长度的部分
    for i >= 0 {
        ai, _ := strconv.Atoi(string(a[i]))
        res[k] = strconv.Itoa((ai + c) % 2)
        c = (ai + c) / 2
        i--
        k--
    }
    // 有进位处理
    if c > 0 {
        res[k] = strconv.Itoa(c)
    }
    return strings.Join(res, "")
}

func main() {
    fmt.Printf("addBinary(\"11\", \"1\") = %v\n",addBinary("11", "1"))  // "100"
    fmt.Printf("addBinary(\"11\", \"11\") = %v\n",addBinary("11", "11")) // "110"
    fmt.Println(addBinary("111111111", "111111111")) // "1111111110"
    fmt.Println(addBinary("000000000", "000000000")) // "000000000"
    fmt.Println(addBinary("000000000", "111111111")) // "111111111"
    fmt.Println(addBinary("111111111", "000000000")) // "111111111"
    fmt.Println(addBinary("101010101", "101010101")) // "1010101010"
    fmt.Println(addBinary("101010101", "010101010")) // "111111111"
    fmt.Println(addBinary("010101010", "101010101")) // "111111111"
    fmt.Println(addBinary("010101010", "010101010")) // "101010100"

    fmt.Printf("addBinary1(\"11\", \"1\") = %v\n",addBinary1("11", "1"))  // "100"
    fmt.Printf("addBinary1(\"11\", \"11\") = %v\n",addBinary1("11", "11")) // "110"
    fmt.Println(addBinary1("111111111", "111111111")) // "1111111110"
    fmt.Println(addBinary1("000000000", "000000000")) // "000000000"
    fmt.Println(addBinary1("000000000", "111111111")) // "111111111"
    fmt.Println(addBinary1("111111111", "000000000")) // "111111111"
    fmt.Println(addBinary1("101010101", "101010101")) // "1010101010"
    fmt.Println(addBinary1("101010101", "010101010")) // "111111111"
    fmt.Println(addBinary1("010101010", "101010101")) // "111111111"
    fmt.Println(addBinary1("010101010", "010101010")) // "101010100"
}
