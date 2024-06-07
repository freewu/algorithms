package main

// LCR 015. 找到字符串中所有字母异位词
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 变位词 指字母相同，但排列不同的字符串。

// 示例 1：
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。

// 示例 2：
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的变位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的变位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的变位词。
 
// 提示:
//     1 <= s.length, p.length <= 3 * 10^4
//     s 和 p 仅包含小写字母

import "fmt"

func findAnagrams(s string, p string) []int {
    var freq [256]int
    res := []int{}
    if len(s) == 0 || len(s) < len(p) {
        return res
    }
    for i := 0; i < len(p); i++ {
        freq[p[i] - 'a']++
    }
    left, right, count := 0, 0, len(p)
    for right < len(s) {
        // 判断是否出现在 p 中
        if freq[s[right]-'a'] >= 1 {
            count--
        }
        // 滑动窗口右边界往右滑动的时候，划过去的元素消耗次数(即次数 --)
        freq[s[right]-'a']--
        // 滑动窗口左边界往右滑动的时候，划过去的元素释放次数(即次数 ++)
        right++
        // 每经过一个符合规范的元素，count 就 --，count 初始值是 len(p)，
        // 当每个元素都符合规范的时候，右边界和左边界相差 len(p) 的时候，count 也会等于 0 。
        // 当区间内有不符合规范的元素(freq < 0 或者是不存在的元素)，那么当区间达到 len(p) 的时候，count 无法减少到 0，
        // 区间右移动的时候，左边界又会开始 count ++，只有当左边界移出了这些不合规范的元素以后，才可能出现 count = 0 的情况
        if count == 0 {
            res = append(res, left)
        }
        // 右边界和左边界相差 len(p) 的时候，需要判断每个元素是否都用过一遍了
        if right - left == len(p) {
            if freq[s[left]-'a'] >= 0 {
                count++
            }
            freq[s[left]-'a']++
            left++
        }
    }
    return res
}

func findAnagrams1(s string, p string) []int {
    sl, pl := len(s), len(p)
    if sl < pl {
        return nil
    }
    res, sc, pc := []int{},[26]int{}, [26]int{}
    for i, b := range p {
        pc[b-'a']++
        sc[s[i]-'a']++
    }
    if sc == pc {
        res = append(res, 0)
    }
    for i := 0; i < sl - pl; i++ {
        sc[s[i]-'a']--
        sc[s[i+pl]-'a']++
        if sc == pc {
            res = append(res, i+1)
        }
    }
    return res
}

func main() {
    // The substring with start index = 0 is "cba", which is an anagram of "abc".
    // The substring with start index = 6 is "bac", which is an anagram of "abc".
    fmt.Println(findAnagrams("cbaebabacd","abc")) // [0,6]

    // The substring with start index = 0 is "ab", which is an anagram of "ab".
    // The substring with start index = 1 is "ba", which is an anagram of "ab".
    // The substring with start index = 2 is "ab", which is an anagram of "ab".
    fmt.Println(findAnagrams("abab","ab")) // [0,1,2]

    fmt.Println(findAnagrams1("cbaebabacd","abc")) // [0,6]
    fmt.Println(findAnagrams1("abab","ab")) // [0,1,2]
}