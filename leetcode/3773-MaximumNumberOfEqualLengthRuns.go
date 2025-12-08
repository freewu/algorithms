package main

// 3773. Maximum Number of Equal Length Runs
// You are given a string s consisting of lowercase English letters.

// A run in s is a substring of equal letters that cannot be extended further. For example, the runs in "hello" are "h", "e", "ll", and "o".

// You can select runs that have the same length in s.

// Return an integer denoting the maximum number of runs you can select in s.

// Example 1:
// Input: s = "hello"
// Output: 3
// Explanation:
// The runs in s are "h", "e", "ll", and "o". You can select "h", "e", and "o" because they have the same length 1.

// Example 2:
// Input: s = "aaabaaa"
// Output: 2
// Explanation:
// The runs in s are "aaa", "b", and "aaa". You can select "aaa" and "aaa" because they have the same length 3.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters only.

import "fmt"

func maxSameLengthRuns(s string) int {
    if len(s) == 0 {  return 0 }
    // 第一步：提取所有run的长度
    res, arr, curr, l := 0, []int{}, s[0], 1
    for i := 1; i < len(s); i++ {
        if s[i] == curr {
            l++
        } else {
            arr = append(arr, l)
            curr = s[i]
            l = 1
        }
    }
    // 把最后一个run加入列表
    arr = append(arr, l)
    // 第二步：统计每个长度出现的次数
    mp := make(map[int]int)
    for _, v := range arr {
        mp[v]++
    }
    // 第三步：找到最大的次数
    for _, v := range mp {
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "hello"
    // Output: 3
    // Explanation:
    // The runs in s are "h", "e", "ll", and "o". You can select "h", "e", and "o" because they have the same length 1.
    fmt.Println(maxSameLengthRuns("hello")) // 3
    // Example 2:
    // Input: s = "aaabaaa"
    // Output: 2
    // Explanation:
    // The runs in s are "aaa", "b", and "aaa". You can select "aaa" and "aaa" because they have the same length 3.
    fmt.Println(maxSameLengthRuns("aaabaaa")) // 2  

    fmt.Println(maxSameLengthRuns("leetcode")) // 6
    fmt.Println(maxSameLengthRuns("bluefrog")) // 8
}