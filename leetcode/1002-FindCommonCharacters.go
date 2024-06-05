package main

// 1002. Find Common Characters
// Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). 
// You may return the answer in any order.

// Example 1:
// Input: words = ["bella","label","roller"]
// Output: ["e","l","l"]

// Example 2:
// Input: words = ["cool","lock","cook"]
// Output: ["c","o"]
 
// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] consists of lowercase English letters.

import "fmt"
import "strings"

func commonChars(words []string) []string {
    if len(words) == 0 {
        return []string{}
    }
    res, stdArr := []string{}, words[0]  // 用第一个单词作为循环依据（随机一个都行）
    for _, ch := range stdArr { //对循环依据进行循环
        s := string(ch) // 需要用来判断是否共有的字母
        for j, word := range words { // 对每个单词进行判断是否包含
            if j == 0 { // 第一个单词是本身，不用比较 
                continue
            }
            if len(word) == 0 { // 如果出现相同的就删除，如果存在一个单词被删除完了，就没必要继续找了
                return res 
            }
            if !strings.Contains(word, s) { // 如果每个字母没有被任何一个单词包含，则无需再比较
                break 
            } else { // 如果该字母没某一个单词包含，则继续看下一个单词中有没有该字母。类似与continue，但是因为还有后续逻辑所以不continue
                words[j] = strings.Replace(word, s, "", 1) 
            }
            if j == len(words) -1 {
                // 如果判断到了最后一个单词还没有break，说明该字母是共有的
                res = append(res, s)
            }
        }
    }
    return res
}

func commonChars1(words []string) []string {
    n := len(words)
    res, memo := []string{}, make([][26]int, n )
    for i, w := range words { // 累加每个字符出现的次数
        for _, ch := range w {
            memo[i][ch-'a']++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < 26; i++ { // 循环 26 个字母
        cur := n
        for _, cnt := range memo { // 取出每个单词出现最小的次数
            cur = min(cur, cnt[i])
        }
        for ; cur > 0; cur-- {
            res = append(res, string(rune(i) + 'a'))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["bella","label","roller"]
    // Output: ["e","l","l"]
    fmt.Println(commonChars([]string{"bella","label","roller"})) // ["e","l","l"]
    // Example 2:
    // Input: words = ["cool","lock","cook"]
    // Output: ["c","o"]
    fmt.Println(commonChars([]string{"cool","lock","cook"})) // ["c","o"]

    fmt.Println(commonChars1([]string{"bella","label","roller"})) // ["e","l","l"]
    fmt.Println(commonChars1([]string{"cool","lock","cook"})) // ["c","o"]
}