package main

// 824. Goat Latin
// You are given a string sentence that consist of words separated by spaces. 
// Each word consists of lowercase and uppercase letters only.

// We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.) The rules of Goat Latin are as follows:
//     1. If a word begins with a vowel ('a', 'e', 'i', 'o', or 'u'), append "ma" to the end of the word.
//         For example, the word "apple" becomes "applema".
//     2. If a word begins with a consonant (i.e., not a vowel), remove the first letter and append it to the end, then add "ma".
//         For example, the word "goat" becomes "oatgma".
//     3. Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
//         For example, the first word gets "a" added to the end, the second word gets "aa" added to the end, and so on.

// Return the final sentence representing the conversion from sentence to Goat Latin.

// Example 1:
// Input: sentence = "I speak Goat Latin"
// Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

// Example 2:
// Input: sentence = "The quick brown fox jumped over the lazy dog"
// Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"

// Constraints:
//     1 <= sentence.length <= 150
//     sentence consists of English letters and spaces.
//     sentence has no leading or trailing spaces.
//     All the words in sentence are separated by a single space.

import "fmt"
import "strings"

func toGoatLatin(sentence string) string {
    a, ss, space := 1,[]string{}, strings.Count(sentence, " ") + 1
    checkVowel := func (s string) bool {
        v := []string{"a", "A", "e", "E", "i", "I", "o", "O", "u", "U"}
        for i := 0; i < len(v); i++ {
            if v[i] == s {
                return true
            }
        }
        return false
    }
    gena := func (n int) string {
        str := "a"
        for i := 1; i < n; i++ {
            str += "a"
        }
        return str
    }
    for i := 0; i < space; i++ {
        s := strings.Split(sentence, " ")[i]
        if checkVowel(string(s[0])) { // 如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"  apple => applema
            ss = append(ss, s + "ma" + gena(a))
        } else { // 如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"  goat => oatgma
            ss = append(ss, s[1:] + s[0:1] + "ma"+gena(a))
        }
        a++ // 根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始
    }
    return strings.Join(ss, " ")
}

func main() {
    // Example 1:
    // Input: sentence = "I speak Goat Latin"
    // Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
    fmt.Println(toGoatLatin("I speak Goat Latin")) // "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
    // Example 2:
    // Input: sentence = "The quick brown fox jumped over the lazy dog"
    // Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
    fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog")) // "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
}