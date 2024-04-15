package main

// 557. Reverse Words in a String III
// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Example 2:
// Input: s = "Mr Ding"
// Output: "rM gniD"

// Constraints:
//     1 <= s.length <= 5 * 10^4
//     s contains printable ASCII characters.
//     s does not contain any leading or trailing spaces.
//     There is at least one word in s.
//     All the words in s are separated by a single space.

import "fmt"
import "strings"

func reverseWords(s string) string {
    arr := strings.Split(s, " ")
    res := []string{}
    for i := 0; i < len(arr); i++ {
        j, k, tmp := 0, len(arr[i]) - 1, make([]byte,len(arr[i]))
        for j <= k {
            tmp[j], tmp[k] = arr[i][k], arr[i][j]
            j++
            k--
        }
        res = append(res,string(tmp))
    }
    return strings.Join(res," ")
}

func reverseWords1(s string) string {
    arr, space, l := []byte(s), 0, len(s)
    for k, v := range arr {
        if v == ' ' || l - 1 == k {
            var last int
            if k == l - 1 {
                last = k
            } else {
                last = k - 1
            }
            for first := space; first < last; first++ {
                arr[first], arr[last] = arr[last], arr[first]
                last--
            }
            space = k + 1
        }
    }
    return string(arr)
}

func main() {
    fmt.Println(reverseWords("Let's take LeetCode contest")) // "s'teL ekat edoCteeL tsetnoc"
    fmt.Println(reverseWords("Mr Ding")) //  "rM gniD"

    fmt.Println(reverseWords1("Let's take LeetCode contest")) // "s'teL ekat edoCteeL tsetnoc"
    fmt.Println(reverseWords1("Mr Ding")) //  "rM gniD"
}