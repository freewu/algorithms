package main

// 面试题 01.05. One Away LCCI
// There are three types of edits that can be performed on strings: 
// 		insert a character, remove a character, or replace a character. 
// Given two strings, write a function to check if they are one edit (or zero edits) away.

// Example 1:
// Input: 
// first = "pale"
// second = "ple"
// Output: True

// Example 2:
// Input: 
// first = "pales"
// second = "pal"
// Output: False

import "fmt"

// func oneEditAway(first string, second string) bool {
// 	if len(first) == 0 || len(second) == 0 || first == second {
// 		return true
// 	}
// 	flag := 1
// 	i, j, l1, l2 := 0 ,0, len(first) - 1, len(second) - 1 
// 	for {
// 		if i > l1 || j > l2 {
// 			break
// 		}
// 		if first[i] != second[j] {
// 			// 只允许有一个不同
// 			if flag == 0 {
// 				return false
// 			}
// 			flag--
// 			i++
// 		} else {
// 			i++
// 			j++
// 		}
// 	}
// 	return flag == 0 // 不需要修改也不行
// }

func oneEditAway(first, second string) bool {
    m, n := len(first), len(second)
	// 长的放前面
    if m < n {
        return oneEditAway(second, first)
    }
	// 如果相差长度大于 1 说明需要多次修改
    if m - n > 1 {
        return false
    }
    for i, ch := range second {
        if first[i] != byte(ch) {
            if m == n { 
				// 长度相同，相同位置出现不同字符的情况
				// 如果两个字符长度等,统一跳过不现同的字符 比对后面的字符即可
                return first[i+1:] == second[i+1:]
            }
			// 多插入了一个新字符的情况
			// first 跳过不同的字符(first更长 1) 比对剩余的字符串
            return first[i+1:] == second[i:]
        }
    }
    return true
}

func main() {
	fmt.Println(oneEditAway("pale","ple")) // true
	fmt.Println(oneEditAway("pales","pal")) // false
	fmt.Println(oneEditAway("","")) // true
	fmt.Println(oneEditAway("","a")) // true
	fmt.Println(oneEditAway("a","a")) // true

	fmt.Println(oneEditAway("aaaa","abaa")) // true
	fmt.Println(oneEditAway("aaaa","abaaa")) // true
	fmt.Println(oneEditAway("abaaa","aaaa")) // true
}