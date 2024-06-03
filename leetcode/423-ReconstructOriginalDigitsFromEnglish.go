package main

// 423. Reconstruct Original Digits from English
// Given a string s containing an out-of-order English representation of digits 0-9, return the digits in ascending order.

// Example 1:
// Input: s = "owoztneoer"
// Output: "012"

// Example 2:
// Input: s = "fviefuro"
// Output: "45"

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is one of the characters ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"].
//     s is guaranteed to be valid.

import "fmt"
import "strings"
import "bytes"

// # 解题思路
// 1 这道题是一道找规律的题目。首先观察 0-9 对应的英文单词，找到特殊规律：所有的偶数都包含一个独特的字母：
//     z 只在 zero 中出现。
//     w 只在 two 中出现。
//     u 只在 four 中出现。
//     x 只在 six 中出现。
//     g 只在 eight 中出现。
// 2 所以先排除掉这些偶数。然后在看剩下来几个数字对应的英文字母，这也是计算 3，5 和 7 的关键，因为有些单词只在一个奇数和一个偶数中出现（而且偶数已经被计算过了）：
//     h 只在 three 和 eight 中出现。
//     f 只在 five 和 four 中出现。
//     s 只在 seven 和 six 中出现。
// 3 接下来只需要处理 9 和 0，思路依然相同。
//     i 在 nine，five，six 和 eight 中出现。
//     n 在 one，seven 和 nine 中出现。
// 最后按照上述的优先级，依次消耗对应的英文字母，生成最终的原始数字。注意按照优先级换算数字的时候，注意有多个重复数字的情况
func originalDigits(s string) string {
    digits := make([]int, 26)
    for i := 0; i < len(s); i++ {
        digits[int(s[i]-'a')]++
    }
    convert := func (b byte, digits []int, s string, num string) string {
        v := digits[int(b-'a')]
        for i := 0; i < len(s); i++ {
            digits[int(s[i]-'a')] -= v
        }
        return strings.Repeat(num, v)
    }
    res := make([]string, 10)
    res[0] = convert('z', digits, "zero", "0")
    res[6] = convert('x', digits, "six", "6")
    res[2] = convert('w', digits, "two", "2")
    res[4] = convert('u', digits, "four", "4")
    res[5] = convert('f', digits, "five", "5")
    res[1] = convert('o', digits, "one", "1")
    res[7] = convert('s', digits, "seven", "7")
    res[3] = convert('r', digits, "three", "3")
    res[8] = convert('t', digits, "eight", "8")
    res[9] = convert('i', digits, "nine", "9")
    return strings.Join(res, "")
}

func originalDigits1(s string) string {
    c, res := map[rune]int{}, []byte{}
    for _, ch := range s {
        c[ch]++
    }
    cnt := [10]int{}
    cnt[0] = c['z']
    cnt[2] = c['w']
    cnt[4] = c['u']
    cnt[6] = c['x']
    cnt[8] = c['g']
    cnt[3] = c['h'] - cnt[8]
    cnt[5] = c['f'] - cnt[4]
    cnt[7] = c['s'] - cnt[6]
    cnt[1] = c['o'] - cnt[0] - cnt[2] - cnt[4]
    cnt[9] = c['i'] - cnt[5] - cnt[6] - cnt[8]
    for i, c := range cnt {
        res = append(res, bytes.Repeat([]byte{byte('0' + i)}, c)...)
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "owoztneoer"
    // Output: "012"
    fmt.Println(originalDigits("owoztneoer")) // 012
    // Example 2:
    // Input: s = "fviefuro"
    // Output: "45"
    fmt.Println(originalDigits("fviefuro")) // 45

    fmt.Println(originalDigits1("owoztneoer")) // 012
    fmt.Println(originalDigits1("fviefuro")) // 45
}