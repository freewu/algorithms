package main

/**
500. Keyboard Row

Given a List of words, return the words that can be typed using letters of alphabet on only one row’s of American keyboard like the image below.

	qwertyuiop
	asdfghjkl
	zxcvbnm

Example 1:

	Input: words = ["Hello","Alaska","Dad","Peace"]
	Output: ["Alaska","Dad"]

Example 2:

	Input: words = ["omk"]
	Output: []

Example 3:

	Input: words = ["adsdf","sfd"]
	Output: ["adsdf","sfd"]


Constraints:

	1 <= words.length <= 20
	1 <= words[i].length <= 100
	words[i] consists of English letters (both lowercase and uppercase).

You may use one character in the keyboard more than once.
You may assume the input string will only contain letters of alphabet.

解题思路:
	给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。
	使用一个 boolean 变量来做开关判断
	1 默认开关为 flag = false
    2 如果包含了字符设置为 flag = !flag (true)
	3 如果还能进入 flag = !flag (false)  if !flag { break }
	4 只有 flag 为 true 才把 符合的单词加入到 等输出列表
 */

import (
	"fmt"
	"strings"
	"unicode"
)

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	output := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s) // 先转成小写
		oneRow := false
		for _, r := range rows { // 循环三列判断
			if strings.ContainsAny(lowerS, r) {
				oneRow = !oneRow // 第一次会变成 true,如果下次循环还能进来 说明还包含其它行的字符
				if !oneRow { // 第二次进入就可以直接跳出说明 该单词字符出现在两列键盘上
					break
				}
			}
		}
		if oneRow {
			output = append(output, s)
		}
	}
	return output
}

// best solution
func findWords1(words []string) []string {
	// 先构建一个 map[row][char] = true 的格式 row 列 char 字符
	sk := make([]map[rune]bool, 3, 3)
	kbl := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for i, kb := range kbl {
		sk[i] = make(map[rune]bool)
		for _, k := range kb {
			sk[i][k] = true
		}
	}
	res := []string{}

	for _, w := range words {
		if len(w) == 1 {
			res = append(res, w) // 如果是一个字符处理
			continue
		}
		c := unicode.ToLower(rune(w[0])) // 取单词的第一个字符
		var s map[rune]bool
		for _, s = range sk {
			if s[c] { // 取到第一个字符的 数组
				break
			}
		}
		fail := false
		for _, c := range w {
			if _, ok := s[unicode.ToLower(c)]; !ok { // 判断是否其它的字符 是否存在第一个字符的所属于的列的map
				fail = true
				break
			}
		}
		if !fail { // 如果没有进入过 fail = true 的循环里 说明单词都在同列中
			res = append(res, w)
		}
	}
	return res
}

func main() {
	fmt.Printf("findWords([]string{ \"Hello\", \"Alaska\", \"Dad\", \"Peace\"}) =  %v\n", findWords([]string{ "Hello", "Alaska", "Dad", "Peace"}))
	fmt.Printf("findWords1([]string{ \"Hello\", \"Alaska\", \"Dad\", \"Peace\"}) =  %v\n", findWords1([]string{ "Hello", "Alaska", "Dad", "Peace"}))
}