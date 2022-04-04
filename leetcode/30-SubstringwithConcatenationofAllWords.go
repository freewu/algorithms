package main

import "fmt"

/**
30. Substring with Concatenation of All Words
You are given a string s and an array of strings words of the same length.
Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
You can return the answer in any order.

Constraints:

	1 <= s.length <= 10^4
	s consists of lower-case English letters.
	1 <= words.length <= 5000
	1 <= words[i].length <= 30
	words[i] consists of lower-case English letters.

Example 1:

	Input: s = "barfoothefoobarman", words = ["foo","bar"]
	Output: [0,9]
	Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
	The output order does not matter, returning [9,0] is fine too.

Example 2:

	Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
	Output: []

Example 3:

	Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
	Output: [6,9,12]

 */

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		//fmt.Printf("sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
		if tmpCounter[s[i:i+length]] > 0 {
			tmpCounter[s[i:i+length]]--
			//fmt.Printf("******sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
			if checkWords(tmpCounter) && (i+length-start == totalLen) {
				res = append(res, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCounter = copyMap(counter)
		}
	}
	return res
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}

// best solution
func findSubstringBest(s string, words []string) []int {
	mainM, curM := map[string]int{}, map[string]int{}
	for _, word := range words {
		if _, ok := mainM[word]; !ok {
			mainM[word] = 0
		}
		mainM[word] += 1
	}
	N, M, wl := len(s), len(words), len(words[0])
	var result []int
	str, tmp := "", ""
	for i := 0; i < wl; i++ {
		count := 0
		windowStart := i
		for windowEnd := i; windowEnd + wl <= N; windowEnd += wl {
			str = s[windowEnd:windowEnd+wl]
			if _, ok := mainM[str]; ok {
				if _, ok := curM[str]; ok {
					curM[str] += 1
				} else {
					curM[str] = 1
				}
				if mainM[str] >= curM[str] { count++ }

				for curM[str] > mainM[str] {
					tmp = s[windowStart:windowStart+wl]
					curM[tmp] -= 1
					windowStart += wl

					if mainM[tmp] > curM[tmp] { count-- }
				}
				if count == M {
					result = append(result, windowStart)
					tmp = s[windowStart:windowStart+wl]
					curM[tmp] -= 1
					count --
					windowStart += wl
				}
			} else {
				clearMap(curM)
				count = 0
				windowStart = windowEnd + wl
			}
		}
		clearMap(curM)
	}
	return result
}

func clearMap(m map[string]int) {
	for k := range m {
		delete(m, k)
	}
}

func main() {
	fmt.Printf("findSubstring(\"barfoothefoobarman\", []string{\"foo\", \"bar\"}) = %v\n", findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0,9]
	fmt.Printf("findSubstring(\"wordgoodgoodgoodbestword\", []string{\"word\", \"good\", \"best\", \"word\"}) = %v\n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
	fmt.Printf("findSubstring(\"barfoofoobarthefoobarman\", []string{\"bar\", \"foo\",\"the\"}) = %v\n", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6,9,12]

	fmt.Printf("findSubstringBest(\"barfoothefoobarman\", []string{\"foo\", \"bar\"}) = %v\n", findSubstringBest("barfoothefoobarman", []string{"foo", "bar"})) // [0,9]
	fmt.Printf("findSubstringBest(\"wordgoodgoodgoodbestword\", []string{\"word\", \"good\", \"best\", \"word\"}) = %v\n", findSubstringBest("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) // []
	fmt.Printf("findSubstringBest(\"barfoofoobarthefoobarman\", []string{\"bar\", \"foo\",\"the\"}) = %v\n", findSubstringBest("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6,9,12]
}
