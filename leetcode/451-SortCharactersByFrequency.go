package main

// 451. Sort Characters By Frequency
// Given a string s, sort it in decreasing order based on the frequency of the characters. 
// The frequency of a character is the number of times it appears in the string.
// Return the sorted string. If there are multiple answers, return any of them.

// Example 1:
// Input: s = "tree"
// Output: "eert"
// Explanation: 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

// Example 2:
// Input: s = "cccaaa"
// Output: "aaaccc"
// Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
// Note that "cacaca" is incorrect, as the same characters must be together.

// Example 3:
// Input: s = "Aabb"
// Output: "bbAa"
// Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
// Note that 'A' and 'a' are treated as two different characters.
 
// Constraints:
// 		1 <= s.length <= 5 * 10^5
// 		s consists of uppercase and lowercase English letters and digits.

import "fmt"
import "sort"
import "bytes"

func frequencySort(s string) string {
	if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	// 统计字符出现次数
	for _, b := range sb {
		sMap[b]++
	}
	// 把 统计字符的 map 转成 map[出现频次] = []byte{ char1, char2, ... }
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}
	// 取出 出现频次 keys 
	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	// 排序
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	// 按出现频次 & 字符 转成新的字符串
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}

// best solution  use sort.Slice
func frequencySort1(s string) string {
    ans := make([]byte, 0)
	type pair struct {
		word byte
		cnt  int
	}
	word := make([]pair, 0)

	for i := range s {
		if len(word) == 0 {
			word = append(word, pair{
				word: s[i],
				cnt:  1,
			})
			continue
		}
		for j := range word {
			if word[j].word == s[i] {
				word[j].cnt++
				break
			}
			if j == len(word)-1 {
				word = append(word, pair{
					word: s[i],
					cnt:  1,
				})
			}
		}
	}

	// 倒序word数组
	sort.Slice(word, func(i, j int) bool {
		return word[i].cnt > word[j].cnt
	})

	for i := range word {
		ans = append(ans, bytes.Repeat([]byte{word[i].word}, word[i].cnt)...)
	}

	return string(ans)
}

func main() {
	fmt.Println("frequencySort(\"tree\"): ",frequencySort("tree")) // eert
	fmt.Println("frequencySort(\"cccaaa\"): ",frequencySort("cccaaa")) // aaaccc
	fmt.Println("frequencySort(\"Aabb\"): ",frequencySort("Aabb")) // bbAa

	fmt.Println("frequencySort1(\"tree\"): ",frequencySort1("tree")) // eert
	fmt.Println("frequencySort1(\"cccaaa\"): ",frequencySort1("cccaaa")) // aaaccc
	fmt.Println("frequencySort1(\"Aabb\"): ",frequencySort1("Aabb")) // bbAa
}
