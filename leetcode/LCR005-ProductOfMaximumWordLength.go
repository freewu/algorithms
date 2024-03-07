package main

// LCR 005. 最大单词长度乘积
// 给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
// 假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。

// 示例 1:
// 输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
// 输出: 16 
// 解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。

// 示例 2:
// 输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
// 输出: 4 
// 解释: 这两个单词为 "ab", "cd"。

// 示例 3:
// 输入: words = ["a","aa","aaa","aaaa"]
// 输出: 0 
// 解释: 不存在这样的两个单词。

// 提示：
//         2 <= words.length <= 1000
//         1 <= words[i].length <= 1000
//         words[i] 仅包含小写字母

import "fmt"

// 在字符串数组中找到 2 个没有公共字符的字符串，并且这两个字符串的长度乘积要是最大的，求这个最大的乘积。
// 利用位运算 & 运算的性质，如果 X & Y = 0，说明 X 和 Y 完全不相同。
// 那么我们将字符串都编码成二进制数，进行 & 运算即可分出没有公共字符的字符串，最后动态维护长度乘积最大值即可。
// 将字符串编码成二进制数的规则比较简单，每个字符相对于 ‘a’ 的距离，根据这个距离将 1 左移多少位。
//         a   1 -> 1  
//         b   2 -> 10  
//         c   4 -> 100  
//         ab  3 -> 11  
//         ac  5 -> 101  
//         abc 7 -> 111  
//         az 33554433->10000000000000000000000001  

func maxProduct(words []string) int {
	if words == nil || len(words) == 0 {
		return 0
	}
	length, value, maxProduct := len(words), make([]int, len(words)), 0
	for i := 0; i < length; i++ {
		tmp := words[i]
		value[i] = 0
		for j := 0; j < len(tmp); j++ {
			value[i] |= 1 << (tmp[j] - 'a') // |= 相同的字符都一样
		}
	}
    // fmt.Println(value)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (value[i] & value[j]) == 0 && (len(words[i]) * len(words[j]) > maxProduct) {
				maxProduct = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProduct
}

func maxProduct1(words []string) int {
    res := 0
    masks := make([]int, len(words))
    for i, word := range words {
        for _, ch := range word {
            masks[i] |= 1 << (ch - 'a')
        }
    }
    for i, x := range masks {
        for j, y := range masks[:i] {
            // fmt.Println("x: ", x ," y: ", y ," x & y",x & y)
            if x & y == 0 && len(words[i]) * len(words[j]) > res {
                res = len(words[i]) * len(words[j])
            }
        }
    }
    return res
}

func main() {
    // 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
    fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","fxyz","abcdef"})) // 16
    // 这两个单词为 "ab", "cd"
    fmt.Println(maxProduct([]string{"a","ab","abc","d","cd","bcd","abcd"})) // 4
    fmt.Println(maxProduct([]string{"a","aa","aaa","aaaa"})) // 0

    // 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
    fmt.Println(maxProduct1([]string{"abcw","baz","foo","bar","fxyz","abcdef"})) // 16
    // 这两个单词为 "ab", "cd"
    fmt.Println(maxProduct1([]string{"a","ab","abc","d","cd","bcd","abcd"})) // 4
    fmt.Println(maxProduct1([]string{"a","aa","aaa","aaaa"})) // 0
}