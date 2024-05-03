package main 

// 87. Scramble String
// We can scramble a string s to get a string t using the following algorithm:
//     If the length of the string is 1, stop.
//     If the length of the string is > 1, do the following:
//         Split the string into two non-empty substrings at a random index, i.e., if the string is s, divide it to x and y where s = x + y.
//         Randomly decide to swap the two substrings or to keep them in the same order. i.e., after this step, s may become s = x + y or s = y + x.
//         Apply step 1 recursively on each of the two substrings x and y.

// Given two strings s1 and s2 of the same length, return true if s2 is a scrambled string of s1, otherwise, return false.

// Example 1:
// Input: s1 = "great", s2 = "rgeat"
// Output: true
// Explanation: One possible scenario applied on s1 is:
// "great" --> "gr/eat" // divide at random index.
// "gr/eat" --> "gr/eat" // random decision is not to swap the two substrings and keep them in order.
// "gr/eat" --> "g/r / e/at" // apply the same algorithm recursively on both substrings. divide at random index each of them.
// "g/r / e/at" --> "r/g / e/at" // random decision was to swap the first substring and to keep the second substring in the same order.
// "r/g / e/at" --> "r/g / e/ a/t" // again apply the algorithm recursively, divide "at" to "a/t".
// "r/g / e/ a/t" --> "r/g / e/ a/t" // random decision is to keep both substrings in the same order.
// The algorithm stops now, and the result string is "rgeat" which is s2.
// As one possible scenario led s1 to be scrambled to s2, we return true.

// Example 2:
// Input: s1 = "abcde", s2 = "caebd"
// Output: false

// Example 3:
// Input: s1 = "a", s2 = "a"
// Output: true
 
// Constraints:
//     s1.length == s2.length
//     1 <= s1.length <= 30
//     s1 and s2 consist of lowercase English letters.
import "fmt"

func isScramble(s1, s2 string) bool {
    isScrambleCash := make(map[string]bool)
    sameLetters := func (s1, s2 string) bool {
        if len(s1) != len(s2) {
            return false
        }
        if s1 == s2 {
            return true
        }
        var arr [26]int
        for i := 0; i < len(s1); i++ {
            arr[int(s1[i]-'a')]++
            arr[int(s2[i]-'a')]--
        }
        return arr == [26]int{}
    }
    var isScrambleFunc func(s1, s2 string) bool
    isScrambleFunc = func(s1, s2 string) bool {
        if v, ok := isScrambleCash[s1 + s2]; ok {
            return v
        }
        if v, ok := isScrambleCash[s2 + s1]; ok {
            return v
        }
        if !sameLetters(s1, s2) {
            isScrambleCash[s1+s2] = false
            return false
        }
        if len(s1) <= 3 {
            isScrambleCash[s1+s2] = true
            return true
        }
        for i := 1; i < len(s1); i++ {
            x1, y1 := s1[:i], s1[i:]
            x2, y2 := s2[:i], s2[i:]
            if isScrambleFunc(x1, x2) && isScrambleFunc(y1, y2) {
                isScrambleCash[s1+s2] = true
                return true
            }
            x2, y2 = s2[len(s2)-i:], s2[:len(s2)-i]
            if isScrambleFunc(x1, x2) && isScrambleFunc(y1, y2) {
                isScrambleCash[s1+s2] = true
                return true
            }
        }
        isScrambleCash[s1+s2] = false
        return false
    }
    return isScrambleFunc(s1, s2)
}

// 记忆化搜索的版本，时间复杂度已经和动态规划版本的一样了
func isScramble1(s1 string, s2 string) bool {
    if (s1 == "" && s2 != "") || (s1 != "" && s2 == "") {
        return false
    }
    if s1 == "" && s2 == "" {
        return true
    }
    if s1 == s2 {
        return true
    }
    str1, str2 := []rune(s1), []rune(s2)
    sameTypeSameNumber := func (str1, str2 []rune) bool {
        if len(str1) != len(str2) {
            return false
        }
        mapTable := make(map[rune]int)
        for _, char := range str1 {
            mapTable[char]++
        }
        for _, char := range str2 {
            mapTable[char]--
            if mapTable[char] < 0 {
                return false
            }
        }
        return true
    }
    if !sameTypeSameNumber(str1, str2) {
        return false
    }
    N := len(s1)
    dp := make([][][]int, N)
    // dp[i][j][k] = 0 processDP(i,j,k)状态之前没有算过的
    // dp[i][j][k] = -1 processDP(i,j,k)状态之前算过的,返回值是false
    // dp[i][j][k] = 1 processDP(i,j,k)状态之前算过的,返回值是true
    for i := range dp {
        dp[i] = make([][]int, N)
        for j := range dp[i] {
            dp[i][j] = make([]int, N+1)
        }
    }
    var processDP func(str1, str2 []rune, L1, L2, size int, dp [][][]int) bool
    processDP = func (str1, str2 []rune, L1, L2, size int, dp [][][]int) bool {
        if dp[L1][L2][size] != 0 {
            return dp[L1][L2][size] == 1
        }
        var res bool
        if size == 1 {
            res = str1[L1] == str2[L2]
        } else {
            for leftPart := 1; leftPart < size; leftPart++ {
                p1 := processDP(str1, str2, L1, L2, leftPart, dp) && processDP(str1, str2, L1+leftPart, L2+leftPart, size-leftPart, dp)
                p2 := processDP(str1, str2, L1, L2+size-leftPart, leftPart,dp) && processDP(str1, str2, L1+leftPart, L2, size-leftPart, dp)
                if p1 || p2 {
                    res = true
                    break
                }
            }
        }
        if res {
            dp[L1][L2][size] = 1
        } else {
            dp[L1][L2][size] = -1
        }
        return res
    }
    return processDP(str1, str2, 0, 0, N, dp)
}

func main() {
    // Example 1:
    // Input: s1 = "great", s2 = "rgeat"
    // Output: true
    // Explanation: One possible scenario applied on s1 is:
    // "great" --> "gr/eat" // divide at random index.
    // "gr/eat" --> "gr/eat" // random decision is not to swap the two substrings and keep them in order.
    // "gr/eat" --> "g/r / e/at" // apply the same algorithm recursively on both substrings. divide at random index each of them.
    // "g/r / e/at" --> "r/g / e/at" // random decision was to swap the first substring and to keep the second substring in the same order.
    // "r/g / e/at" --> "r/g / e/ a/t" // again apply the algorithm recursively, divide "at" to "a/t".
    // "r/g / e/ a/t" --> "r/g / e/ a/t" // random decision is to keep both substrings in the same order.
    // The algorithm stops now, and the result string is "rgeat" which is s2.
    // As one possible scenario led s1 to be scrambled to s2, we return true.
    fmt.Println(isScramble("great","rgeat")) // true
    // Example 2:
    // Input: s1 = "abcde", s2 = "caebd"
    // Output: false
    fmt.Println(isScramble("abcde","caebd")) // false
    // Example 3:
    // Input: s1 = "a", s2 = "a"
    // Output: true
    fmt.Println(isScramble("a","a")) // true

    fmt.Println(isScramble1("great","rgeat")) // true
    fmt.Println(isScramble1("abcde","caebd")) // false
    fmt.Println(isScramble1("a","a")) // true
}