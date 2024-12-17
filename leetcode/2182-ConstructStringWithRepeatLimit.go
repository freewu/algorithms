package main

// 2182. Construct String With Repeat Limit
// You are given a string s and an integer repeatLimit. 
// Construct a new string repeatLimitedString using the characters of s such that no letter appears more than repeatLimit times in a row. 
// You do not have to use all characters from s.

// Return the lexicographically largest repeatLimitedString possible.

// A string a is lexicographically larger than a string b if in the first position where a and b differ, string a has a letter that appears later in the alphabet than the corresponding letter in b. 
// If the first min(a.length, b.length) characters do not differ, then the longer string is the lexicographically larger one.

// Example 1:
// Input: s = "cczazcc", repeatLimit = 3
// Output: "zzcccac"
// Explanation: We use all of the characters from s to construct the repeatLimitedString "zzcccac".
// The letter 'a' appears at most 1 time in a row.
// The letter 'c' appears at most 3 times in a row.
// The letter 'z' appears at most 2 times in a row.
// Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
// The string is the lexicographically largest repeatLimitedString possible so we return "zzcccac".
// Note that the string "zzcccca" is lexicographically larger but the letter 'c' appears more than 3 times in a row, so it is not a valid repeatLimitedString.

// Example 2:
// Input: s = "aababab", repeatLimit = 2
// Output: "bbabaa"
// Explanation: We use only some of the characters from s to construct the repeatLimitedString "bbabaa". 
// The letter 'a' appears at most 2 times in a row.
// The letter 'b' appears at most 2 times in a row.
// Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
// The string is the lexicographically largest repeatLimitedString possible so we return "bbabaa".
// Note that the string "bbabaaa" is lexicographically larger but the letter 'a' appears more than 2 times in a row, so it is not a valid repeatLimitedString.
 
// Constraints:
//     1 <= repeatLimit <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"

// 双指针
func repeatLimitedString(s string, repeatLimit int) string {
    n := len(s)
    charCnt := [26]int{}  // 字符个数
    left, right := -1, n // 双指针，left指向小的字符，right指向大的
    res := make([]rune, 0, n)
    for _, c := range s { // 初始化 cnt
        charCnt[c - 'a']++
    }
    for i := 26 - 1; i >= 0; i-- { // 初始化 双指针
        if charCnt[i] != 0 {
            right = i
            for i--; i >= 0 && charCnt[i] == 0; i-- {
            }
            left = i
            break
        }
    }
    // 双指针遍历
    for right >= 0 {
        // 先放right（大的），达到limit后放一个left（小的）过渡
        if charCnt[right] > repeatLimit {
            for i := 0; i < repeatLimit; i++ {
                res = append(res, rune('a' + right))
            }
            if left == -1 { // 没有小的字符可用
                break
            }
            res = append(res, rune('a'+left))
            charCnt[right] -= repeatLimit
            charCnt[left]--
            // 小的用完了 更新
            for left >= 0 && charCnt[left] == 0 {
                left--
            }
        } else { // 放全部大的，更新转移双指针
            for i := 0; i < charCnt[right]; i++ {
                res = append(res, rune('a' + right))
            }
            right = left
            for left--; left >= 0 && charCnt[left] == 0; left-- {
            }
        }
    }
    return string(res)
}

func repeatLimitedString1(s string, repeatLimit int) string {
    count := make([]int, 26)
    for _, c := range s {
        count[c-'a']++ // 下标: 0 ~ 25
    }
    res, repeatCount := []byte{}, 0
    for i, j := 25, 25; i >= 0 && j >= 0; {
        if count[i] == 0 {
            i--
            j = i
            repeatCount = 0
            continue
        }
        if repeatCount < repeatLimit {
            count[i]--
            res = append(res, 'a' + byte(i))
            repeatCount++
            continue
        }
        if j >= i || count[j] == 0 { // 开始用 j 指针
            j--
            continue
        }
        count[j]--
        res = append(res, 'a'+ byte(j))
        repeatCount = 0
    }
    return string(res)
}

func repeatLimitedString2(s string, repeatLimit int) string {
    res, count := []byte{}, [26]int{}
    for _, v := range s {
        count[v - 'a']++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, j := 25, 24; i >= 0; i-- {
        j = min(j, i-1)
        for {
            for k := min(count[i], repeatLimit); k > 0; k-- {
                res = append(res, byte(i + 'a'))
                count[i]--
            }
            if count[i] == 0 { break }
            for j >= 0 && count[j] == 0 { j-- }
            if j < 0 { break }
            res = append(res, byte(j + 'a'))
            count[j]--
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "cczazcc", repeatLimit = 3
    // Output: "zzcccac"
    // Explanation: We use all of the characters from s to construct the repeatLimitedString "zzcccac".
    // The letter 'a' appears at most 1 time in a row.
    // The letter 'c' appears at most 3 times in a row.
    // The letter 'z' appears at most 2 times in a row.
    // Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
    // The string is the lexicographically largest repeatLimitedString possible so we return "zzcccac".
    // Note that the string "zzcccca" is lexicographically larger but the letter 'c' appears more than 3 times in a row, so it is not a valid repeatLimitedString.
    fmt.Println(repeatLimitedString("cczazcc", 3)) // "zzcccac"
    // Example 2:
    // Input: s = "aababab", repeatLimit = 2
    // Output: "bbabaa"
    // Explanation: We use only some of the characters from s to construct the repeatLimitedString "bbabaa". 
    // The letter 'a' appears at most 2 times in a row.
    // The letter 'b' appears at most 2 times in a row.
    // Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
    // The string is the lexicographically largest repeatLimitedString possible so we return "bbabaa".
    // Note that the string "bbabaaa" is lexicographically larger but the letter 'a' appears more than 2 times in a row, so it is not a valid repeatLimitedString.
    fmt.Println(repeatLimitedString("aababab", 2)) // "bbabaa"

    fmt.Println(repeatLimitedString1("cczazcc", 3)) // "zzcccac"
    fmt.Println(repeatLimitedString1("aababab", 2)) // "bbabaa"

    fmt.Println(repeatLimitedString2("cczazcc", 3)) // "zzcccac"
    fmt.Println(repeatLimitedString2("aababab", 2)) // "bbabaa"
}