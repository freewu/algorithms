package main

// 186. Reverse Words in a String II
// Given a character array s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by a single space.
// Your code must solve the problem in-place, i.e. without allocating extra space.

// Example 1:
// Input: s = ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
// Output: ["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]

// Example 2:
// Input: s = ["a"]
// Output: ["a"]
 
// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is an English letter (uppercase or lowercase), digit, or space ' '.
//     There is at least one word in s.
//     s does not contain leading or trailing spaces.
//     All the words in s are guaranteed to be separated by a single space.

import "fmt"

func reverseWords(s []byte)  {
    reverse := func (s []byte, from, to int) { // 从 from 反转到 to 位置
        for from < to {
            s[from], s[to] = s[to], s[from]
            from++
            to--
        }
    }
    reverse(s, 0,len(s) - 1)
    start := 0 
    for i ,ch := range s { // 反转每个单词
        if ch == ' ' {
            reverse(s,start,i - 1) // 反转单词
            start = i + 1
        }
    }
    reverse(s, start, len(s) - 1) // 最后一个单词之后没有" "，所以最后一个单词不会被反转
}

func reverseWords1(s []byte)  {
    n := len(s)
    for i := 0; i < n / 2; i++ {
        s[i], s[n-i-1] = s[n-i-1], s[i]
    }
    i, j := 0, 0
    for j < n {
        for j < n && s[j] != ' ' {
            j++
        }
        r := j-1
        for ;i < r; i++ {
            s[i], s[r] = s[r], s[i]
            r--
        }
        j++
        i = j
    }
}

func main() {
    // Example 1:
    // Input: s = ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
    // Output: ["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]
    s1 := []byte{'t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e'}
    fmt.Println("before: ", string(s1)) //  the sky is blue
    reverseWords(s1)
    fmt.Println("after: ",  string(s1)) // blue is sky the
    // Example 2:
    // Input: s = ["a"]
    // Output: ["a"]
    s2 := []byte{'a'}
    fmt.Println("before: ", string(s2)) // a 
    reverseWords(s2)
    fmt.Println("after: ",  string(s2)) // a


    s11 := []byte{'t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e'}
    fmt.Println("before: ", string(s11)) //  the sky is blue
    reverseWords1(s11)
    fmt.Println("after: ",  string(s11)) // blue is sky the
    s12 := []byte{'a'}
    fmt.Println("before: ", string(s12)) // a 
    reverseWords1(s12)
    fmt.Println("after: ",  string(s12)) // a
}