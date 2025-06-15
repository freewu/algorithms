package main

// 3582. Generate Tag for Video Caption
// You are given a string caption representing the caption for a video.

// The following actions must be performed in order to generate a valid tag for the video:

// Combine all words in the string into a single camelCase string prefixed with '#'. 
// A camelCase string is one where the first letter of all words except the first one is capitalized. 
// All characters after the first character in each word must be lowercase.

// Remove all characters that are not an English letter, except the first '#'.

// Truncate the result to a maximum of 100 characters.

// Return the tag after performing the actions on caption.

// Example 1:
// Input: caption = "Leetcode daily streak achieved"
// Output: "#leetcodeDailyStreakAchieved"
// Explanation:
// The first letter for all words except "leetcode" should be capitalized.

// Example 2:
// Input: caption = "can I Go There"
// Output: "#canIGoThere"
// Explanation:
// The first letter for all words except "can" should be capitalized.

// Example 3:
// Input: caption = "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
// Output: "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
// Explanation:
// Since the first word has length 101, we need to truncate the last two letters from the word.

// Constraints:
//     1 <= caption.length <= 150
//     caption consists only of English letters and ' '.

import "fmt"
import "strings"

func generateTag(caption string) string {
    caption = strings.TrimSpace(caption)
    if caption == "" { return "#" }
    caption = strings.ToLower(caption)
    caption = strings.Title(caption)
    caption = strings.ReplaceAll(caption, " ", "")
    caption = "#" + strings.ToLower(caption[:1]) + caption[1:] // 开头小写
    if len(caption) > 100 { return caption[:100] }
    return caption
}

func generateTag1(caption string) string {
    s := strings.ToLower(strings.TrimSpace(caption))
    res, count, flag :=[]byte{'#'}, 1, false // flag 前面是不是空格
    for i := range s {
        if s[i] == ' ' {
            flag = true
            continue
        }
        if count >= 100 { break }
        if flag {
            res = append(res,'A' + s[i] - 'a')
            flag = false
        }else{
            res = append(res,s[i])
        }
        count++
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: caption = "Leetcode daily streak achieved"
    // Output: "#leetcodeDailyStreakAchieved"
    // Explanation:
    // The first letter for all words except "leetcode" should be capitalized.
    fmt.Println(generateTag("Leetcode daily streak achieved")) // "#leetcodeDailyStreakAchieved"
    // Example 2:
    // Input: caption = "can I Go There"
    // Output: "#canIGoThere"
    // Explanation:
    // The first letter for all words except "can" should be capitalized.
    fmt.Println(generateTag("can I Go There")) // "#canIGoThere"
    // Example 3:
    // Input: caption = "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
    // Output: "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
    // Explanation:
    // Since the first word has length 101, we need to truncate the last two letters from the word.
    fmt.Println(generateTag("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")) // "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"

    fmt.Println(generateTag("bluefrog")) // "#bluefrog"
    fmt.Println(generateTag("leetcode")) // "#leetcode"
    fmt.Println(generateTag("#")) // "##"
    fmt.Println(generateTag("   ")) // "#"

    fmt.Println(generateTag1("Leetcode daily streak achieved")) // "#leetcodeDailyStreakAchieved"
    fmt.Println(generateTag1("can I Go There")) // "#canIGoThere"
    fmt.Println(generateTag1("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")) // "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
    fmt.Println(generateTag1("bluefrog")) // "#bluefrog"
    fmt.Println(generateTag1("leetcode")) // "#leetcode"
    fmt.Println(generateTag1("#")) // "##"
    fmt.Println(generateTag1("   ")) // "#"
}