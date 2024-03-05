package main

// 1750. Minimum Length of String After Deleting Similar Ends
// Given a string s consisting only of characters 'a', 'b', and 'c'. 
// You are asked to apply the following algorithm on the string any number of times:
//         Pick a non-empty prefix from the string s where all the characters in the prefix are equal.
//         Pick a non-empty suffix from the string s where all the characters in this suffix are equal.

// The prefix and the suffix should not intersect at any index.
// The characters from the prefix and suffix must be the same.
// Delete both the prefix and the suffix.
// Return the minimum length of s after performing the above operation any number of times (possibly zero times).

// Example 1:
// Input: s = "ca"
// Output: 2
// Explanation: You can't remove any characters, so the string stays as is.

// Example 2:
// Input: s = "cabaabac"
// Output: 0
// Explanation: An optimal sequence of operations is:
// - Take prefix = "c" and suffix = "c" and remove them, s = "abaaba".
// - Take prefix = "a" and suffix = "a" and remove them, s = "baab".
// - Take prefix = "b" and suffix = "b" and remove them, s = "aa".
// - Take prefix = "a" and suffix = "a" and remove them, s = "".

// Example 3:
// Input: s = "aabccabba"
// Output: 3
// Explanation: An optimal sequence of operations is:
// - Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
// - Take prefix = "b" and suffix = "bb" and remove them, s = "cca".
 
// Constraints:
//         1 <= s.length <= 10^5
//         s only consists of characters 'a', 'b', and 'c'.

import "fmt"

// func minimumLength(s string) int {
//     r := 0
//     l := len(s) - 1
//     for {
//         // 如果首尾不同直接跳出循环
//         if len(s) == 0 || s[r] != s[l] {
//             break
//         }
//         // 把 s 前后的 c 都剔除掉
//         s = trim(s,s[r])
//         l = len(s) - 1
//     }
//     return len(s)
// }

// func trim(s string,c byte) string {
//     res := make([]byte,0)
//     for i := 0; i < len(s); i++ {
//         if s[i] == c {
//             continue
//         }
//         res = append(res,s[i])
//     }
//     return string(res)
// }

// 双指针
func minimumLength(s string) int {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		c := s[l]
        // 左边有相同字符 进 ->
		for l <= r && s[l] == c {
			l++
		}
        // 右边有相同字符 退 <-
		for r >= l && s[r] == c {
			r--
		}
	}
	return r - l + 1
}

func minimumLength1(s string) int {
    left, right := 0, len(s)-1
    for left < right {
        lc := s[left]
        rc := s[right]
        // left and right are pointing to the next character to be removed.
        if lc == rc {
            // left can remove the next character pointed by right and vice versa.
            // This is the only way we can remove the entire string.
            for left <= right && s[left] == rc {
                left++
            }
            for right >= left && s[right] == lc {
                right--
            }
        } else {
            break
        }
    }
    return right - left + 1
}

func main() {
    // You can't remove any characters, so the string stays as is.
    fmt.Println(minimumLength("ca")) // 2

    // - Take prefix = "c" and suffix = "c" and remove them, s = "abaaba".
    // - Take prefix = "a" and suffix = "a" and remove them, s = "baab".
    // - Take prefix = "b" and suffix = "b" and remove them, s = "aa".
    // - Take prefix = "a" and suffix = "a" and remove them, s = "".
    fmt.Println(minimumLength("cabaabac")) // 0

    // - Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
    // - Take prefix = "b" and suffix = "bb" and remove them, s = "cca".
    fmt.Println(minimumLength("aabccabba")) // 3


    fmt.Println(minimumLength1("ca")) // 2
    fmt.Println(minimumLength1("cabaabac")) // 0
    fmt.Println(minimumLength1("aabccabba")) // 3
}