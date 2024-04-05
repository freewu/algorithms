package main

// 1544. Make The String Great
// Given a string s of lower and upper case English letters.
// A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
//     0 <= i <= s.length - 2
//     s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.

// To make the string good, you can choose two adjacent characters that make the string bad and remove them. 
// You can keep doing this until the string becomes good.

// Return the string after making it good. The answer is guaranteed to be unique under the given constraints.
// Notice that an empty string is also good.

// Example 1:
// Input: s = "leEeetcode"
// Output: "leetcode"
// Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".

// Example 2:
// Input: s = "abBAcC"
// Output: ""
// Explanation: We have many possible scenarios, and all lead to the same answer. For example:
// "abBAcC" --> "aAcC" --> "cC" --> ""
// "abBAcC" --> "abBA" --> "aA" --> ""

// Example 3:
// Input: s = "s"
// Output: "s"
 
// Constraints:
//     1 <= s.length <= 100
//     s contains only lower and upper case English letters.

import "fmt"

// stack
func makeGood(s string) string {
    stack := make([]byte, 0, len(s))
    // 判断是否是同一个字符大小写
    equal := func (x, y byte) bool {
        if x == y {
            return false
        }
        if x > y {
            return y + 32 == x
        }
        return x + 32 == y
    }
    for i := 0; i < len(s); i++ {
        if len(stack) != 0 && equal(stack[len(stack) - 1], s[i]) { // 跟栈顶比较
            stack = stack[0 : len(stack) - 1] // 相等，栈顶弹栈
        }else {
            stack = append(stack, s[i]) // 不相等，当前值入栈
        }
    }
    res := ""
    for i := 0; i < len(stack); i++ {
        res += string(stack[i])
    }
    return res
}

// func makeGood1(s string) string {
//     if len(s) <= 1 {
//         return s
//     }
//     res, l := []byte{}, len(s)
//     abs := func (x int) int { if x >= 0 { return x; }; return -1 * x; }
//     for i := 0; i < l;  {
//         // 不为最后一位时
//         if i != l - 1 {
//             fmt.Printf("i=%d ,l - 1 = %d, s[i]=%c\n",i, l - 1, s[i])
//             if abs(int(s[i] - s[i + 1])) == 32 {
//                 i += 2
//                 continue
//             }
//         }
//         res = append(res,s[i])
//         i++
//     }
//     return string(res)
// }

func makeGood1(s string) string {
    res, top := []byte(s), 0 
    isMatched := func(c1, c2 byte) bool {
        // <-. c1, c2是同一个字符的大小写形式: ↓ 
        return (c1 - 'a') == (c2 - 'A') || (c1 - 'A') == (c2 - 'a')
    }
    for i := 0; i < len(s); i++ {
        ch := s[i] 
        if top > 0 && isMatched(ch, res[top-1]) {
            top-- 
        } else {
            res[top] = ch; 
            top++ 
        }
    }
    return string(res[:top])
}

func main() {
    //fmt.Println(int('a' - 'A'))
    // Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
    fmt.Println(makeGood("leEeetcode")) // leetcode
    // Explanation: We have many possible scenarios, and all lead to the same answer. For example:
    // "abBAcC" --> "aAcC" --> "cC" --> ""
    // "abBAcC" --> "abBA" --> "aA" --> ""
    fmt.Println(makeGood("abBAcC")) // ""
    fmt.Println(makeGood("s")) // s
    fmt.Println(makeGood("")) // ""

    fmt.Println(makeGood1("leEeetcode")) // leetcode
    fmt.Println(makeGood1("abBAcC")) // ""
    fmt.Println(makeGood1("s")) // s
    fmt.Println(makeGood1("")) // ""
}