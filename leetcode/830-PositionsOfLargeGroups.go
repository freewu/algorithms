package main

// 830. Positions of Large Groups
// In a string s of lowercase letters, these letters form consecutive groups of the same character.
// For example, a string like s = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z", and "yy".
// A group is identified by an interval [start, end], where start and end denote the start and end indices (inclusive) of the group. In the above example, "xxxx" has the interval [3,6].
// A group is considered large if it has 3 or more characters.
// Return the intervals of every large group sorted in increasing order by start index.

// Example 1:
// Input: s = "abbxxxxzzy"
// Output: [[3,6]]
// Explanation: "xxxx" is the only large group with start index 3 and end index 6.

// Example 2:
// Input: s = "abc"
// Output: []
// Explanation: We have groups "a", "b", and "c", none of which are large groups.

// Example 3:
// Input: s = "abcdddeeeeaabbbcd"
// Output: [[3,5],[6,9],[12,14]]
// Explanation: The large groups are "ddd", "eeee", and "bbb".

// Constraints:
//     1 <= s.length <= 1000
//     s contains lowercase English letters only.

import "fmt"

func largeGroupPositions(s string) [][]int {
    res, n := [][]int{}, len(s)
    index, char, count := 0, s[0], 1
    for i := 1; i < n; i++ {
        if s[i] == char { // 如果相同字符累加
            count++
            if i == n - 1 && count > 2 { // 最后一个字符处理
                res = append(res,[]int{index, i})
            } 
        } else {
            if count > 2 { // 超过连续3个相同的字符
                res = append(res,[]int{index, i - 1})
            }
            index = i
            char = s[i]
            count = 1 // 重置 count
        }
    }
    return res
}

func main() {
// Example 1:
// Input: s = "abbxxxxzzy"
// Output: [[3,6]]
// Explanation: "xxxx" is the only large group with start index 3 and end index 6.
fmt.Println(largeGroupPositions("abbxxxxzzy")) //  [[3,6]]
// Example 2:
// Input: s = "abc"
// Output: []
// Explanation: We have groups "a", "b", and "c", none of which are large groups.
fmt.Println(largeGroupPositions("abc")) //  []
// Example 3:
// Input: s = "abcdddeeeeaabbbcd"
// Output: [[3,5],[6,9],[12,14]]
// Explanation: The large groups are "ddd", "eeee", and "bbb".
fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5],[6,9],[12,14]]

fmt.Println(largeGroupPositions("aaa")) // [[0,2]]
}