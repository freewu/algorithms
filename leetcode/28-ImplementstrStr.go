package main

// 28. Implement strStr()
// Implement strStr().
// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Clarification:
//     What should we return when needle is an empty string? This is a great question to ask during an interview.
//     For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

// Constraints:
//     1 <= haystack.length, needle.length <= 10^4
//     haystack and needle consist of only lowercase English characters.

// Example 1:
//     Input: haystack = "hello", needle = "ll"
//     Output: 2

// Example 2:
//     Input: haystack = "aaaaa", needle = "bba"
//     Output: -1

import "fmt"
import "strings"

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


// best solution
func strStr1(haystack string, needle string) int {
    for i := 0; ; i++ {
        for j := 0; ; j++ {
            if j == len(needle) {
                return i
            }
            if i+j == len(haystack) {
                return -1
            }
            if needle[j] != haystack[i+j] {
                break
            }
        }
    }
}

// 使用现成的库
func strStr2(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}

func strStr3(haystack string, needle string) int {
    // 长度不够直接返回 -1
    if len(haystack) < len(needle) {
        return -1
    }
    for i := 0; i <= (len(haystack) - len(needle)); i++ {
        j := 0
        for ; j < len(needle); j++ {
            // 偏移 haystack 和 needle 比对 发现不同直接跳出 从 haystack的 i+1 重新匹配 j 重新归 0 开始
            if haystack[i + j] != needle[j] {
                break
            }
        }
        // 如果 匹配了 needle 的长度说明匹配成功了 直接返回 i  needle 在 haystack 出的位置
        if j == len(needle) {
            return i
        }
    }
    return -1
}

func main() {
	fmt.Printf("strStr(\"hello\",\"ll\") = %v\n",strStr("hello","ll")) // 2
	fmt.Printf("strStr(\"hello\",\"ll\") = %v\n",strStr("hello","lla")) // -1
	fmt.Printf("strStr(\"\",\"\") = %v\n",strStr("","")) // 0
	fmt.Printf("strStr(\"a\",\"\") = %v\n",strStr("a","")) // 0
	fmt.Printf("strStr(\"mississippi\",\"issi\") = %v\n",strStr("mississippi","issi")) // 1
	fmt.Printf("strStr(\"mississippi\",\"ppii\") = %v\n",strStr("mississippi","ppii")) // -1

	fmt.Printf("strStr1(\"hello\",\"ll\") = %v\n",strStr1("hello","ll")) // 2
	fmt.Printf("strStr1(\"hello\",\"ll\") = %v\n",strStr1("hello","lla")) // -1
	fmt.Printf("strStr1(\"\",\"\") = %v\n",strStr1("","")) // 0
	fmt.Printf("strStr1(\"a\",\"\") = %v\n",strStr1("a","")) // 0
	fmt.Printf("strStr1(\"mississippi\",\"issi\") = %v\n",strStr1("mississippi","issi")) // 1
	fmt.Printf("strStr1(\"mississippi\",\"ppii\") = %v\n",strStr1("mississippi","ppii")) // -1

	fmt.Printf("strStr2(\"hello\",\"ll\") = %v\n",strStr2("hello","ll")) // 2
	fmt.Printf("strStr2(\"hello\",\"ll\") = %v\n",strStr2("hello","lla")) // -1
	fmt.Printf("strStr2(\"\",\"\") = %v\n",strStr2("","")) // 0
	fmt.Printf("strStr2(\"a\",\"\") = %v\n",strStr2("a","")) // 0
	fmt.Printf("strStr2(\"mississippi\",\"issi\") = %v\n",strStr2("mississippi","issi")) // 1
	fmt.Printf("strStr2(\"mississippi\",\"ppii\") = %v\n",strStr2("mississippi","ppii")) // -1

    fmt.Printf("strStr3(\"hello\",\"ll\") = %v\n",strStr3("hello","ll")) // 2
	fmt.Printf("strStr3(\"hello\",\"ll\") = %v\n",strStr3("hello","lla")) // -1
	fmt.Printf("strStr3(\"\",\"\") = %v\n",strStr3("","")) // 0
	fmt.Printf("strStr3(\"a\",\"\") = %v\n",strStr3("a","")) // 0
	fmt.Printf("strStr3(\"mississippi\",\"issi\") = %v\n",strStr3("mississippi","issi")) // 1
	fmt.Printf("strStr3(\"mississippi\",\"ppii\") = %v\n",strStr3("mississippi","ppii")) // -1
}