package main

// 2864. Maximum Odd Binary Number
// You are given a binary string s that contains at least one '1'.
// You have to rearrange the bits in such a way that the resulting binary number is the maximum odd binary number that can be created from this combination.
// Return a string representing the maximum odd binary number that can be created from the given combination.
// Note that the resulting string can have leading zeros.

// Example 1:
// Input: s = "010"
// Output: "001"
// Explanation: Because there is just one '1', it must be in the last position. So the answer is "001".

// Example 2:
// Input: s = "0101"
// Output: "1001"
// Explanation: One of the '1's must be in the last position. The maximum number that can be made with the remaining digits is "100". So the answer is "1001".
 
// Constraints:
//         1 <= s.length <= 100
//         s consists only of '0' and '1'.
//         s contains at least one '1'.

import "fmt"

func maximumOddBinaryNumber(s string) string {
    // 用两个数存 0 / 1 的个数 
    num0 := 0
    num1 := 0
    for i := 0; i < len(s); i++  {
        if s[i] == '0' {
            num0++
        }
        if s[i] == '1' {
            num1++
        }
    }
    // 处理都是 0000 的情况
    if num1 == 0 {
        return s
    }
    // 重新生成字符串
    res := []rune{}
    // 留一个1 在最后  Odd
    for i := 0; i < num1 - 1; i++ {
        res = append(res,'1')
    }
    // 补全 0
    for i := 0; i < num0; i++ {
        res = append(res,'0')
    }
    res = append(res,'1')
    return string(res)
}

func maximumOddBinaryNumber1(s string) string {
    // 统计s中字符`1`的个数，保证最低位是1的前提下将剩余的所有`1`放置高位: 
    var n int = len(s)
    cnt := 0 
    for _, ch := range s {
        if ch == '1' {
            cnt++
        }
    } 
    var res []byte = make([]byte, n) 
    res[n-1] = '1' 
    cnt -= 1 
    for i := 0; i < cnt; i++ {
        res[i] = '1'
    }
    for j := cnt; j < n - 1; j++ {
        res[j] = '0'
    }
    return string(res)
}

func main() {
    fmt.Println(maximumOddBinaryNumber("010")) // 001
    fmt.Println(maximumOddBinaryNumber("0101")) // 1001

    fmt.Println(maximumOddBinaryNumber1("010")) // 001
    fmt.Println(maximumOddBinaryNumber1("0101")) // 1001
}