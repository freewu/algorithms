package main

// 438. Find All Anagrams in a String
// Given two strings s and p, return an array of all the start indices of p's anagrams in s. 
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
// typically using all the original letters exactly once.

// Example 1:
// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".

// Example 2:
// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".
 
// Constraints:
//     1 <= s.length, p.length <= 3 * 10^4
//     s and p consist of lowercase English letters.

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
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return nil
	}
	sCount, pCount := [26]int{}, [26]int{}
	for i, b := range p {
		pCount[b-'a']++
		sCount[s[i]-'a']++
	}
	res := []int{}
	if sCount == pCount {
		res = append(res, 0)
	}
	for i := 0; i < sLen-pLen; i++ {
		sCount[s[i]-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
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