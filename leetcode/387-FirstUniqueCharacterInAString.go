package main

import "fmt"

// 387. First Unique Character in a String
// Given a string s, find the first non-repeating character in it and return its index. 
// If it does not exist, return -1.
 
// Example 1:
// Input: s = "leetcode"
// Output: 0

// Example 2:
// Input: s = "loveleetcode"
// Output: 2

// Example 3:
// Input: s = "aabb"
// Output: -1
 
// Constraints:
// 		1 <= s.length <= 10^5
// 		s consists of only lowercase English letters.

func firstUniqChar(s string) int {
	// 声明一个 26 个字符的 map(只有小写字母)
	m := make(map[byte]int,26)
	l := len(s)
	// 统计字符出现的次数
	for i := 0; i < l; i = i + 1 {
		m[s[i]] += 1 
	}
	for i := 0; i < l; i = i + 1 {
		// 第一个只出现了一次的字符
		if (m[s[i]] == 1) {
			return i
		}
	}
	return -1
}

// best solution
func firstUniqChar1(s string) int {
    m := [26]int{}
    for i := 0; i < len(s); i++ {
        m[s[i]-'a']++
    }
    for i := 0; i < len(s); i++ {
       if m[s[i]-'a'] == 1 {
           return i
       }
    }
    return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode")) // 0
	fmt.Println(firstUniqChar("loveleetcode")) // 2
	fmt.Println(firstUniqChar("aabb")) // -1

	fmt.Println(firstUniqChar1("leetcode")) // 0
	fmt.Println(firstUniqChar1("loveleetcode")) // 2
	fmt.Println(firstUniqChar1("aabb")) // -1
}