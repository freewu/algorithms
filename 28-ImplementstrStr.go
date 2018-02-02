package main

/*
http://www.cplusplus.com/reference/cstring/strstr/
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/
import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if haystack == needle || "" ==  needle {
		return 0
	}
    var nl = len(needle)
    var hl = len(haystack)
	if  nl > hl {
		return -1
	}

	var t int
	for i := 0; i < hl; i++ {
		t = nl
		for j := 0; j < nl; j++ {
			
			// fmt.Println(haystack[i]," ",needle[j])
			// 不足长度了
			if (nl - j) > (hl - i) {
				return -1
			}
			// 如果第一个没有匹配到直接跳出本循环
			if haystack[i + j] != needle[j] {
				break;
			}
			t--
			//fmt.Println(t," ",j)
			if 0 == t {
				return i
			}
		}
	}
    return -1
}

func main() {
	fmt.Println(strStr("hello","ll")) // 2
	fmt.Println(strStr("hello","lla")) // -1

	fmt.Println(strStr("","")) // 0
	fmt.Println(strStr("a","")) // 0

	fmt.Println(strStr("mississippi","issi")) // 1
	fmt.Println(strStr("mississippi","ppii")) // -1
}