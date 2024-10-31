package main

// 1647. Minimum Deletions to Make Character Frequencies Unique
// A string s is called good if there are no two different characters in s that have the same frequency.

// Given a string s, return the minimum number of characters you need to delete to make s good.

// The frequency of a character in a string is the number of times it appears in the string. 
// For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.

// Example 1:
// Input: s = "aab"
// Output: 0
// Explanation: s is already good.

// Example 2:
// Input: s = "aaabbbcc"
// Output: 2
// Explanation: You can delete two 'b's resulting in the good string "aaabcc".
// Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".

// Example 3:
// Input: s = "ceabaacb"
// Output: 2
// Explanation: You can delete both 'c's resulting in the good string "eabaab".
// Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters.

import "fmt"
import "sort"

func minDeletions(s string) int {
    mp, arr := make(map[byte]int), []int{}
    for i := 0 ; i < len(s); i++ {
        mp[s[i]]++
    }
    for _, v := range mp {
        arr = append(arr, v)
    }
    // sort decreasing order
    sort.Ints(arr)
    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
    res, i := 0, 1
    for i < len(arr) {
        // when zero comes
        if arr[i - 1] == 0 {
            for j := i; j < len(arr); j++ {
                res += arr[j]
            }
            break
        }
        for arr[i - 1] <= arr[i] {
            arr[i]--
            res++
        }
        i++
    }
    return res
}

func minDeletions1(s string) int {
    res, count, mp := 0, make([]int, 26), make(map[int]int) // matches counts to characters with that count
    for _, v := range s {
        count[v - 'a']++
    }
    for i := 0; i < len(count); i++ {
        _, ok := mp[count[i]]
        if !ok {
            mp[count[i]] = 1
            continue
        }
        for j := count[i]; j > 0 && ok; j-- {
            res++
            _, ok = mp[j-1]
            if !ok {
                mp[j-1] = 1
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aab"
    // Output: 0
    // Explanation: s is already good.
    fmt.Println(minDeletions("aab")) // 0
    // Example 2:
    // Input: s = "aaabbbcc"
    // Output: 2
    // Explanation: You can delete two 'b's resulting in the good string "aaabcc".
    // Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".
    fmt.Println(minDeletions("aaabbbcc")) // 2
    // Example 3:
    // Input: s = "ceabaacb"
    // Output: 2
    // Explanation: You can delete both 'c's resulting in the good string "eabaab".
    // Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).
    fmt.Println(minDeletions("ceabaacb")) // 2
    fmt.Println(minDeletions("accdcdadddbaadbc")) // 1

    fmt.Println(minDeletions1("aab")) // 0
    fmt.Println(minDeletions1("aaabbbcc")) // 2
    fmt.Println(minDeletions1("ceabaacb")) // 2
    fmt.Println(minDeletions1("accdcdadddbaadbc")) // 1
}