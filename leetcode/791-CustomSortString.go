package main

// 791. Custom Sort String
// You are given two strings order and s. All the characters of order are unique and were sorted in some custom order previously.
// Permute the characters of s so that they match the order that order was sorted. 
// More specifically, if a character x occurs before a character y in order, then x should occur before y in the permuted string.
// Return any permutation of s that satisfies this property.

// Example 1:
// Input:  order = "cba", s = "abcd" 
// Output:  "cbad" 
// Explanation: 
//         "a", "b", "c" appear in order, so the order of "a", "b", "c" should be "c", "b", and "a".
//         Since "d" does not appear in order, it can be at any position in the returned string. 
//         "dcba", "cdba", "cbda" are also valid outputs.

// Example 2:
// Input:  order = "bcafg", s = "abcd" 
// Output:  "bcad" 
// Explanation: 
//         The characters "b", "c", and "a" from order dictate the order for the characters in s. 
//         The character "d" in s does not appear in order, so its position is flexible.
//         Following the order of appearance in order, "b", "c", and "a" from s should be arranged as "b", "c", "a". 
//         "d" can be placed at any position since it's not in order. 
//         The output "bcad" correctly follows this rule. 
//         Other arrangements like "bacd" or "bcda" would also be valid, as long as "b", "c", "a" maintain their order.

// Constraints:
//         1 <= order.length <= 26
//         1 <= s.length <= 200
//         order and s consist of lowercase English letters.
//         All the characters of order are unique.

import "fmt"

import "sort"
import "strings"

// map + sort
func customSortString(order string, str string) string {
	magic := map[byte]int{}
    // 先把排序规则存到 map 中
    // // S 字符串最长为 26 位，先将 S 中字符的下标向左偏移 30，并将偏移后的下标值存入字典中。
	for i := range order {
		magic[order[i]] = i - 30
	}
	byteSlice := []byte(str)
    // 再把 T 字符串按照字典中下标值进行排序。S 中出现的字符对应的下标经过处理以后变成了负数，S 中未出现的字符的下标还是正数。
    // 所以经过排序以后，S 中出现的字符按照原有顺序排列在前面，S 中未出现的字符依次排在后面。
	sort.Slice(byteSlice, func(i, j int) bool {
		return magic[byteSlice[i]] < magic[byteSlice[j]]
	})
	return string(byteSlice)
}

func customSortString1(order string, s string) string {
    var ans string
	var cnt [26]int
    // 计算每个 s 出现在频率
	for _, c := range s {
		cnt[c - 'a']++
	}
	for _, c := range order { // order 只会出现一次
        // s 中出现几次就在 ans 中加上几个
		ans = ans + strings.Repeat(string(c), cnt[c-'a'])
		cnt[c - 'a'] = 0
	}
	for i, v := range cnt {
        // 没有出现在 order 规则里的 追加到 ans 尾部
		if v > 0 {
			ans += strings.Repeat(string(rune('a'+i)), v)
		}
	}
	return ans
}

func main() {
    fmt.Println(customSortString("cba","abcd")) // cbad
    fmt.Println(customSortString("bcafg","abcd")) // bcad

    fmt.Println(customSortString1("cba","abcd")) // cbad
    fmt.Println(customSortString1("bcafg","abcd")) // bcad
}