package main

// LCR 034. 验证外星语词典
// 某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
// 给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。

// 示例 1：
// 输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// 输出：true
// 解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。

// 示例 2：
// 输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
// 输出：false
// 解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。

// 示例 3：
// 输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
// 输出：false
// 解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。

// 提示：
//     1 <= words.length <= 100
//     1 <= words[i].length <= 20
//     order.length == 26
//     在 words[i] 和 order 中的所有字符都是英文小写字母。

import "fmt"

func isAlienSorted(words []string, order string) bool {
    alphabet := [26]int{}
    for i:=0; i < len(order); i++ {	
        alphabet[order[i] - 'a'] = i
    }
    prev := words[0]
    for i := 1; i < len(words); i++ {
        j := 0
        for j = 0; j < len(prev)-1 && j < len(words[i])-1 && prev[j] == words[i][j]; j++ {}
        if  alphabet[prev[j] - 'a'] >  alphabet[words[i][j] - 'a'] || 
            alphabet[prev[j] - 'a'] == alphabet[words[i][j] - 'a'] && 
            len(prev) > len(words[i]) {
            return false
        }
        prev = words[i]
    }
    return true
}

func isAlienSorted1(words []string, order string) bool {
    index := [26]int{}
    for i, c := range order {
        index[c-'a'] = i
    }
next:
    for i := 1; i < len(words); i++ {
        for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
            pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
            if pre > cur {
                return false
            }
            if pre < cur {
                continue next
            }
        }
        if len(words[i-1]) > len(words[i]) {
            return false
        }
    }
    return true
}


func main() {
    // Example 1:
    // Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
    // Output: true
    // Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
    fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
    // Example 2:
    // Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
    // Output: false
    // Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
    fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz")) // false
    // Example 3:
    // Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
    // Output: false
    // Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
    fmt.Println(isAlienSorted([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz")) // false

    fmt.Println(isAlienSorted1([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
    fmt.Println(isAlienSorted1([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz")) // false
    fmt.Println(isAlienSorted1([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz")) // false
}