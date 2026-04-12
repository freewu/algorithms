package main

// 3900. Longest Balanced Substring After One Swap
// You are given a binary string s consisting only of characters '0' and '1'.

// A string is balanced if it contains an equal number of '0's and '1's.

// You can perform at most one swap between any two characters in s. Then, you select a balanced substring from s.

// Return an integer representing the maximum length of the balanced substring you can select.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "100001"
// Output: 4
// Explanation:
// Swap "100001". The string becomes "101000".
// Select the substring "101000", which is balanced because it has two '0's and two '1's.

// Example 2:
// Input: s = "111"
// Output: 0
// Explanation:
// Choose not to perform any swaps.
// Select the empty substring, which is balanced because it has zero '0's and zero '1's.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of the characters '0' and '1'.

import "fmt"
import "strings"

func longestBalanced(s string) int {
    sum0 := strings.Count(s, "0")
    sum1 := len(s) - sum0
    pos := map[int][]int{0: {-1}}
    res, sum := 0, 0 // 前缀和
    for i, v := range s {
        sum += int(v - '0')*2 - 1
        if p := pos[sum]; len(p) < 2 {
            pos[sum] = append(p, i)
        }
        // 不交换
        res = max(res, i-pos[sum][0])
        // 交换子串内的一个 1 和子串外的一个 0
        if p, ok := pos[sum-2]; ok {
            if (i-p[0]-2)/2 < sum0 {
                res = max(res, i-p[0])
            } else if len(p) > 1 {
                res = max(res, i-p[1])
            }
        }
        // 交换子串内的一个 0 和子串外的一个 1
        if p, ok := pos[sum + 2]; ok {
            if (i - p[0] - 2) / 2 < sum1 {
                res = max(res, i - p[0])
            } else if len(p) > 1 {
                res = max(res, i - p[1])
            }
        }
    }
    return res
}

func longestBalanced1(s string) int {
    res, curr, sum0, sum1, n := 0, 0, 0, 0, len(s)
    for i := 0; i < n; i++ {
        if s[i] == '0' {
            sum0++
        } else {
            sum1++
        }
    }
    firstOcc, secondOcc := make([]int, 2 *  n+ 1), make([]int, 2 * n + 1)
    for i := range firstOcc {
        firstOcc[i], secondOcc[i] = -2, -2
    }
    firstOcc[0 + n] = -1
    for j := 0; j < n; j++ {
        if s[j] == '1' {
            curr++
        } else {
            curr--
        }
        // Case 1: Already balanced (sum == 0)
        if firstOcc[curr + n] != -2 {
            l := j - firstOcc[curr + n]
            if l > res { 
                res = l 
            }
        }
        // Case 2: Needs one '0' swapped in (sum == +2)
        target2 := curr - 2
        if target2 + n >= 0 && target2 + n <= 2*n {
            if firstOcc[target2 + n] != -2 {
                l1 := j - firstOcc[target2 + n]    
                // Check if there's a '0' left outside the substring
                if (l1/2 - 1) < sum0 {
                    if l1 > res { 
                        res = l1 
                    }
                } else if secondOcc[target2 + n] != -2 {
                    // If first is invalid, the second occurrence is mathematically guaranteed to be valid
                    l2 := j - secondOcc[target2 + n]
                    if l2 > res { 
                        res = l2 
                    }
                }
            }
        }
        // Case 3: Needs one '1' swapped in (sum == -2)
        target3 := curr + 2
        if target3 + n >= 0 && target3 + n <= 2*n {
            if firstOcc[target3 + n] != -2 {
                l1 := j - firstOcc[target3 + n]    
                // Check if there's a '1' left outside the substring
                if (l1/2 - 1) < sum1 {
                    if l1 > res { 
                        res = l1 
                    }
                } else if secondOcc[target3 + n] != -2 {
                    l2 := j - secondOcc[target3 + n]
                    if l2 > res { 
                        res = l2 
                    }
                }
            }
        }
        // Record the first AND second occurrences
        if firstOcc[curr + n] == -2 {
            firstOcc[curr + n] = j
        } else if secondOcc[curr + n] == -2 {
            secondOcc[curr + n] = j
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "100001"
    // Output: 4
    // Explanation:
    // Swap "100001". The string becomes "101000".
    // Select the substring "101000", which is balanced because it has two '0's and two '1's.
    fmt.Println(longestBalanced("100001")) // 4
    // Example 2:
    // Input: s = "111"
    // Output: 0
    // Explanation:
    // Choose not to perform any swaps.
    // Select the empty substring, which is balanced because it has zero '0's and zero '1's.
    fmt.Println(longestBalanced("111")) // 0

    fmt.Println(longestBalanced("0000000000")) // 0
    fmt.Println(longestBalanced("1111111111")) // 0
    fmt.Println(longestBalanced("0000011111")) // 10
    fmt.Println(longestBalanced("1111100000")) // 10
    fmt.Println(longestBalanced("1010101010")) // 10
    fmt.Println(longestBalanced("0101010101")) // 10

    fmt.Println(longestBalanced1("100001")) // 4
    fmt.Println(longestBalanced1("111")) // 0
    fmt.Println(longestBalanced1("0000000000")) // 0
    fmt.Println(longestBalanced1("1111111111")) // 0
    fmt.Println(longestBalanced1("0000011111")) // 10
    fmt.Println(longestBalanced1("1111100000")) // 10
    fmt.Println(longestBalanced1("1010101010")) // 10
    fmt.Println(longestBalanced1("0101010101")) // 10
}