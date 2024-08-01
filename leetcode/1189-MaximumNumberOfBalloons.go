package main

// 1189. Maximum Number of Balloons
// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
// You can use each character in text at most once. Return the maximum number of instances that can be formed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG" />
// Input: text = "nlaebolko"
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG" />
// Input: text = "loonbalxballpoon"
// Output: 2

// Example 3:
// Input: text = "leetcode"
// Output: 0

// Constraints:
//     1 <= text.length <= 10^4
//     text consists of lower case English letters only.

import "fmt"

func maxNumberOfBalloons(text string) int {
    mp := make(map[byte]int)
    for i := 0; i < len(text); i++ {
        mp[text[i]]++
    }
    // balloon 
    b, a, n, o, l := mp['b'], mp['a'],  mp['n'], mp['o'], mp['l']
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(min(min(min(a, b), n), o / 2), l / 2)
}

func maxNumberOfBalloons1(text string) int {
    cnt := [5]int{}
    for _, ch := range text {
        if ch == 'b' {
            cnt[0]++
        } else if ch == 'a' {
            cnt[1]++
        } else if ch == 'l' {
            cnt[2]++
        } else if ch == 'o' {
            cnt[3]++
        } else if ch == 'n' {
            cnt[4]++
        }
    }
    cnt[2] /= 2
    cnt[3] /= 2
    res := cnt[0]
    for _, v := range cnt[1:] {
        if v < res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG" />
    // Input: text = "nlaebolko"
    // Output: 1
    fmt.Println(maxNumberOfBalloons("nlaebolko")) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG" />
    // Input: text = "loonbalxballpoon"
    // Output: 2
    fmt.Println(maxNumberOfBalloons("loonbalxballpoon")) // 2
    // Example 3:
    // Input: text = "leetcode"
    // Output: 0
    fmt.Println(maxNumberOfBalloons("leetcode")) // 0

    fmt.Println(maxNumberOfBalloons1("nlaebolko")) // 1
    fmt.Println(maxNumberOfBalloons1("loonbalxballpoon")) // 2
    fmt.Println(maxNumberOfBalloons1("leetcode")) // 0
}