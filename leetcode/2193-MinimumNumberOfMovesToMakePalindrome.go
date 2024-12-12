package main

// 2193. Minimum Number of Moves to Make Palindrome
// You are given a string s consisting only of lowercase English letters.

// In one move, you can select any two adjacent characters of s and swap them.

// Return the minimum number of moves needed to make s a palindrome.

// Note that the input will be generated such that s can always be converted to a palindrome.

// Example 1:
// Input: s = "aabb"
// Output: 2
// Explanation:
// We can obtain two palindromes from s, "abba" and "baab". 
// - We can obtain "abba" from s in 2 moves: "aabb" -> "abab" -> "abba".
// - We can obtain "baab" from s in 2 moves: "aabb" -> "abab" -> "baab".
// Thus, the minimum number of moves needed to make s a palindrome is 2.

// Example 2:
// Input: s = "letelt"
// Output: 2
// Explanation:
// One of the palindromes we can obtain from s in 2 moves is "lettel".
// One of the ways we can obtain it is "letelt" -> "letetl" -> "lettel".
// Other palindromes such as "tleelt" can also be obtained in 2 moves.
// It can be shown that it is not possible to obtain a palindrome in less than 2 moves.

// Constraints:
//     1 <= s.length <= 2000
//     s consists only of lowercase English letters.
//     s can be converted to a palindrome using a finite number of moves.

import "fmt"

func minMovesToMakePalindrome(s string) int {
    arr, res, n := []byte(s), 0, len(s)
    for i, j := 0, n - 1; i < j; i++ {
        even := false
        for k := j; k != i; k-- {
            if arr[i] == arr[k] {
                even = true
                for ; k < j; k++ {
                    arr[k], arr[k + 1] = arr[k+1], arr[k]
                    res++
                }
                j--
                break
            }
        }
        if !even {
            res += n / 2 - i
        }
    }
    return res
}

func minMovesToMakePalindrome1(s string) int {
    arr, res, n := []byte(s), 0, len(s)
    for left, right := 0, n - 1; left < right; { // 双指针，从两端向中间移动
        if arr[left] == arr[right] { // 左右字符相等，直接向内移动
            left++
            right--
        } else { // 尝试在右侧找到匹配的字符
            match := right
            for match > left && arr[match] != arr[left] {
                match--
            }
            if match == left { // 如果没有匹配到，说明当前左字符需要和相邻右字符交换
                arr[left], arr[left + 1] = arr[left + 1], arr[left]
                res++
            } else {
                for match < right { // 找到匹配字符，将其移动到右端
                    arr[match], arr[match + 1] = arr[match + 1], arr[match]
                    match++
                    res++
                }
                // 左右指针向内移动
                left++
                right--
            }
        }
    }
    return res
}

func minMovesToMakePalindrome2(s string) int {
    arr, uniqueSwap, swaps := []byte(s), 0, 0
    for len(arr) > 2 {
        if arr[0] != arr[len(arr) - 1] { // if l and r match, l++, r--
            j := len(arr) - 2
            for j > 0{
                if arr[0] == arr[j] { break }
                j--
            }
            if j == 0 {
                // this means its the unique, save length
                uniqueSwap = len(arr) / 2
                arr = arr[1:]
                continue
            }
            swaps += len(arr) - j - 1
            arr = append(arr[1:j], arr[j + 1:]...)
            continue
        }
        arr = arr[1:len(arr) - 1] // else, find the rightest match for l and string manipulate
    }
    return swaps + uniqueSwap
}

func main() {
    // Example 1:
    // Input: s = "aabb"
    // Output: 2
    // Explanation:
    // We can obtain two palindromes from s, "abba" and "baab". 
    // - We can obtain "abba" from s in 2 moves: "aabb" -> "abab" -> "abba".
    // - We can obtain "baab" from s in 2 moves: "aabb" -> "abab" -> "baab".
    // Thus, the minimum number of moves needed to make s a palindrome is 2.
    fmt.Println(minMovesToMakePalindrome("aabb")) // 2
    // Example 2:
    // Input: s = "letelt"
    // Output: 2
    // Explanation:
    // One of the palindromes we can obtain from s in 2 moves is "lettel".
    // One of the ways we can obtain it is "letelt" -> "letetl" -> "lettel".
    // Other palindromes such as "tleelt" can also be obtained in 2 moves.
    // It can be shown that it is not possible to obtain a palindrome in less than 2 moves.
    fmt.Println(minMovesToMakePalindrome("letelt")) // 2

    fmt.Println(minMovesToMakePalindrome("abcdefghijklmnopqrstuvwxyz")) // 25
    fmt.Println(minMovesToMakePalindrome("aaaaaaaa")) // 0

    fmt.Println(minMovesToMakePalindrome1("aabb")) // 2
    fmt.Println(minMovesToMakePalindrome1("letelt")) // 2
    //fmt.Println(minMovesToMakePalindrome1("abcdefghijklmnopqrstuvwxyz")) // 25
    fmt.Println(minMovesToMakePalindrome1("aaaaaaaa")) // 0

    fmt.Println(minMovesToMakePalindrome2("aabb")) // 2
    fmt.Println(minMovesToMakePalindrome2("letelt")) // 2
    fmt.Println(minMovesToMakePalindrome2("abcdefghijklmnopqrstuvwxyz")) // 1
    fmt.Println(minMovesToMakePalindrome2("aaaaaaaa")) // 0
}