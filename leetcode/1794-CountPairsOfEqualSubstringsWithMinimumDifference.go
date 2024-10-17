package main

// 1794. Count Pairs of Equal Substrings With Minimum Difference
// You are given two strings firstString and secondString that are 0-indexed and consist only of lowercase English letters. 
// Count the number of index quadruples (i,j,a,b) that satisfy the following conditions:
//     1. 0 <= i <= j < firstString.length
//     2. 0 <= a <= b < secondString.length
//     3. The substring of firstString that starts at the ith character 
//        and ends at the jth character (inclusive) is equal to the substring of secondString 
//        that starts at the ath character and ends at the bth character (inclusive).
//     4. j - a is the minimum possible value among all quadruples that satisfy the previous conditions.

// Return the number of such quadruples.

// Example 1:
// Input: firstString = "abcd", secondString = "bccda"
// Output: 1
// Explanation: The quadruple (0,0,4,4) is the only one that satisfies all the conditions and minimizes j - a.

// Example 2:
// Input: firstString = "ab", secondString = "cd"
// Output: 0
// Explanation: There are no quadruples satisfying all the conditions.

// Constraints:
//     1 <= firstString.length, secondString.length <= 2 * 10^5
//     Both strings consist only of lowercase English letters.

import "fmt"

// // 解答错误 94 / 96 
// func countQuadruples(firstString string, secondString string) int {
//     mp :=  make(map[string][]int) // 构建字典，将每个字符 作为key，value为这个字符出现的位置, 如 "acbc" =  {a:0,b:[2],c:[1,3]}
//     for i := 0; i < len(secondString); i++ {
//         key := string(secondString[i])
//         mp[key] = append(mp[key], i)
//     }
//     // 扫描firstString每个字符，在第二个字符串中，出现的位置， {3:2}表示 公共子串中，j-a =3 时，四元组的个数
//     res := make(map[int]int)
//     for i := 0; i < len(firstString); i++ {
//         key := string(firstString[i])
//         if _, ok := mp[key]; ok { // secondString 里存在
//             diff := i - mp[key][len(mp[key]) - 1]
//             res[diff] = res[diff] + 1
//         }
//     }
//     mn := 0
//     for k, _ := range res { // 取最小的 元素，出现的次数，然后返回
//         if k < mn { mn = k }
//     }
//     return res[mn]
// }

// // 解答错误 95 / 96 
// func countQuadruples(firstString string, secondString string) int {
//     res, mn := 0, 1 >> 31
//     mp := make([]int, 26)
//     for i := 0; i < len(secondString); i++ {
//         mp[int(secondString[i] - 'a')] = i + 1
//     }
//     for i := 0; i < len(firstString); i++ {
//         index := mp[int(firstString[i] - 'a')]
//         if index != 0 {
//             if i - index < mn {
//                 res, mn = 1, i - index
//             } else if i - index == mn {
//                 res++
//             }
//         } 
//     }
//     return res
// }

func countQuadruples(firstString string, secondString string) int {
    mp1, mp2 := make([]int, 26), make([]int, 26)
    for i := range mp1 { // fill -1
        mp1[i], mp2[i] = -1, -1
    }
    for i := 0; i < len(firstString); i++ {
        key := int(firstString[i] - 'a')
        if mp1[key] == -1 { mp1[key] = i }
    }
    for i := len(secondString) - 1; i >= 0; i-- {
        key := int(secondString[i] - 'a')
        if mp2[key] == -1 { mp2[key] = i }
    }
    res, dist := 0, 200006
    for i := 0; i < 26; i++ {
        if mp1[i] != -1 && mp2[i] != -1 { // 都存在相关字符
            if mp1[i] - mp2[i] < dist {
                res, dist = 1, mp1[i] - mp2[i]
            } else if mp1[i] - mp2[i] == dist {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: firstString = "abcd", secondString = "bccda"
    // Output: 1
    // Explanation: The quadruple (0,0,4,4) is the only one that satisfies all the conditions and minimizes j - a.
    fmt.Println(countQuadruples("abcd", "bccda")) // 1
    // Example 2:
    // Input: firstString = "ab", secondString = "cd"
    // Output: 0
    // Explanation: There are no quadruples satisfying all the conditions.
    fmt.Println(countQuadruples("ab", "cd")) // 0

    fmt.Println(countQuadruples("ab", "b")) // 1
    fmt.Println(countQuadruples("deeba", "b")) // 1
    fmt.Println(countQuadruples("dwtgmqucavlta", "gupciaqadwtgm")) // 5
}