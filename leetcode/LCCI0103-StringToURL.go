package main

// 面试题 01.03. String to URL LCCI
// Write a method to replace all spaces in a string with '%20'. 
// You may assume that the string has sufficient space at the end to hold the additional characters,
// and that you are given the "true" length of the string. 
// (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)

// Example 1:
// Input: "Mr John Smith ", 13
// Output: "Mr%20John%20Smith"

// Example 2:
// Input: "               ", 5
// Output: "%20%20%20%20%20"
 
// Note:
// 	0 <= S.length <= 500000

import "fmt"

func replaceSpaces(S string, length int) string {
	res := []byte{}
	s := byte(' ')
	for i := 0; i < len(S); i += 1 {
		// 为 '' 替换成 %20
		if S[i] == s  {
			res = append(res,'%')
			res = append(res,'2')
			res = append(res,'0')
		} else {
			res = append(res,S[i])
		}
		if i == length - 1 {
			break
		}
	}
	return string(res)
}

func replaceSpaces1(S string, length int) string {
    bytes := []byte(S)
    i, j := len(S)-1, length-1
    for j >= 0 {
        if bytes[j] == ' ' {
            bytes[i] = '0'
            bytes[i-1] = '2'
            bytes[i-2] = '%'
            i -= 3
        } else {
            bytes[i] = bytes[j]
            i--
        }
        j--
    }
    return string(bytes[i+1:])
}

func main() {
	fmt.Println(replaceSpaces("Mr John Smith ",13))
	fmt.Println(replaceSpaces("               ",5))
	fmt.Println(replaceSpaces1("Mr John Smith ",13))
	fmt.Println(replaceSpaces1("               ",5))
}