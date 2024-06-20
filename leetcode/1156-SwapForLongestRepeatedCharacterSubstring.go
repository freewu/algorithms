package main

// 1156. Swap For Longest Repeated Character Substring
// You are given a string text. You can swap two of the characters in the text.
// Return the length of the longest substring with repeated characters.

// Example 1:
// Input: text = "ababa"
// Output: 3
// Explanation: We can swap the first 'b' with the last 'a', or the last 'b' with the first 'a'. Then, the longest repeated character substring is "aaa" with length 3.

// Example 2:
// Input: text = "aaabaaa"
// Output: 6
// Explanation: Swap 'b' with the last 'a' (or the first 'a'), and we get longest repeated character substring "aaaaaa" with length 6.

// Example 3:
// Input: text = "aaaaa"
// Output: 5
// Explanation: No need to swap, longest repeated character substring is "aaaaa" with length is 5.

// Constraints:
//     1 <= text.length <= 2 * 10^4
//     text consist of lowercase English characters only.

import "fmt"

func maxRepOpt1(text string) int {
    type block struct { // track blocks of same letters in text
        size    int
        pos     int
    }
    res, blocks := 1, [26][]block{} // build blocks list for each letter
    for i := 0; i < len(text); {
        j := i+1
        for ; j < len(text) && text[i] == text[j]; j++ {
        }
        blocks[text[i]-'a'] = append(blocks[text[i]-'a'], block{j-i, i})
        i = j
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, bList := range blocks {
        if len(bList) == 0 { // skip letters not found in text
            continue
        }
        if len(bList) == 1 { // 只有一个字符
            res = max(res, bList[0].size) // only 1 block found for this letter, just return size
        } else {
            res = max(res, bList[0].size+1) // if more than 2 blocks, can at least add 1 to first block 
            for i := 1 ; i < len(bList); i++ {
                res = max(res, bList[i].size+1) // same calculation as above for each block size
                if bList[i].pos - (bList[i-1].pos+bList[i-1].size) == 1 { // when 2 blocks with only 1 letter in between is found
                    if len(bList) > 2 {
                        res = max(res, bList[i-1].size + bList[i].size + 1) // if we have more than 2 blocks, we can drop in a letter from any 3rd block to form a large continuous
                    } else {
                        res = max(res, bList[i-1].size + bList[i].size) // or we just combine the 2 block sizes
                    }
                }
            }
        }
    }
    return res
}

// 滑动窗口
//1 统计每个字符总数量
//2 找出连续一段[i,j)
//3 向前扩展: 扩展一个, 两段连接
func maxRepOpt11(text string) int {
    res, count := 0, make(map[rune]int)
    for _, c := range text { // 统计每个字符出现的总次数
        count[c]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(text);  {
        // step1: 找出当前连续的一段 [i, j)
        j := i
        for j < len(text) && text[j] == text[i] {
            j++
        }
        curCnt := j - i
        // step2: 如果这一段长度小于该字符出现的总数，并且前面或后面有空位，则使用 cur_cnt + 1 更新答案
        if curCnt < count[rune(text[i])] && (j < len(text) || i > 0) {
            res = max(res, curCnt+1)
        }
        // step3: 找到这一段后面与之相隔一个不同字符的另一段 [j + 1, k)，如果不存在则 k = j + 1
        k := j + 1
        for k < len(text) && text[k] == text[i] {
            k++
        }
        res = max(res, min(k-i, count[rune(text[i])]))
        i = j
    }
    return res
}

func main() {
    // Example 1:
    // Input: text = "ababa"
    // Output: 3
    // Explanation: We can swap the first 'b' with the last 'a', or the last 'b' with the first 'a'. Then, the longest repeated character substring is "aaa" with length 3.
    fmt.Println(maxRepOpt1("ababa")) // 3
    // Example 2:
    // Input: text = "aaabaaa"
    // Output: 6
    // Explanation: Swap 'b' with the last 'a' (or the first 'a'), and we get longest repeated character substring "aaaaaa" with length 6.
    fmt.Println(maxRepOpt1("aaabaaa")) // 6
    // Example 3:
    // Input: text = "aaaaa"
    // Output: 5
    // Explanation: No need to swap, longest repeated character substring is "aaaaa" with length is 5.
    fmt.Println(maxRepOpt1("aaaaa")) // 5

    fmt.Println(maxRepOpt11("ababa")) // 3
    fmt.Println(maxRepOpt11("aaabaaa")) // 6
    fmt.Println(maxRepOpt11("aaaaa")) // 5
}