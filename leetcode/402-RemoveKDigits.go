package main

// 402. Remove K Digits
// Given string num representing a non-negative integer num, and an integer k, 
// return the smallest possible integer after removing k digits from num.

// Example 1:
// Input: num = "1432219", k = 3
// Output: "1219"
// Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

// Example 2:
// Input: num = "10200", k = 1
// Output: "200"
// Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

// Example 3:
// Input: num = "10", k = 2
// Output: "0"
// Explanation: Remove all the digits from the number and it is left with nothing which is 0.
 
// Constraints:
//     1 <= k <= num.length <= 10^5
//     num consists of only digits.
//     num does not have any leading zeros except for the zero itself.

import "fmt"
import "strings"

func removeKdigits(num string, k int) string {
    if k >= len(num) {
        return "0"
    }
    stack := []byte{}
    // 循环 num
    for i:= range num {
        // 如果栈尾大于 当前则出栈
        for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {  
            stack = stack[:len(stack)-1] // pop bigger numbers out
            k--
        }
        stack = append(stack,num[i])
    }
    // 处理一样的数字的情况 把剩余的 K 出栈完
    for len(stack) > 0 && k > 0 { // smallest number left
        stack = stack[:len(stack)-1] // drop last digits 
        k--
    }
    // 把前面为 0 处理掉
    for len(stack) > 0 && stack[0] == '0' { // remove leading 0
        stack = stack[1:]
    }
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}

func removeKdigits1(num string, k int) string {
    stack := []byte{}
    for i := range num {
        b := num[i]
        for k > 0 && len(stack) > 0 && b < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, b)
    }
    stack = stack[:len(stack)-k]
    // 把前面为 0 处理掉
    res := strings.TrimLeft(string(stack), "0")
    if len(res) == 0 {
        res = "0"
    }
    return res
}

func main() {
    // Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
    fmt.Println(removeKdigits("1432219",3)) // 1219
    // Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
    fmt.Println(removeKdigits("10200",1)) // 200
    // Explanation: Remove all the digits from the number and it is left with nothing which is 0.
    fmt.Println(removeKdigits("10",2)) // 0

    // Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
    fmt.Println(removeKdigits1("1432219",3)) // 1219
    // Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
    fmt.Println(removeKdigits1("10200",1)) // 200
    // Explanation: Remove all the digits from the number and it is left with nothing which is 0.
    fmt.Println(removeKdigits1("10",2)) // 0
}