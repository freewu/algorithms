package main

// 2950. Number of Divisible Substrings
// Each character of the English alphabet has been mapped to a digit as shown below.
// <img src="https://assets.leetcode.com/uploads/2023/11/28/old_phone_digits.png" />

// A string is divisible if the sum of the mapped values of its characters is divisible by its length.

// Given a string s, return the number of divisible substrings of s.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Substring	Mapped	    Sum	Length	Divisible?
//     a          1        1    1      Yes
//     s          7        7    1      Yes
//     d          2        2    1      Yes
//     f          3        3    1      Yes
//    as          1,7      8    2      Yes
//    sd          7,2      9    2      No
//    df          2,3      5    2      No
//   asd          1,7,2    10   3      No
//   sdf          7,2,3    12   3      Yes
//  asdf          1,7,2,3  13   4      No
// Input: word = "asdf"
// Output: 6
// Explanation: The table above contains the details about every substring of word, and we can see that 6 of them are divisible.

// Example 2:
// Input: word = "bdh"
// Output: 4
// Explanation: The 4 divisible substrings are: "b", "d", "h", "bdh".
// It can be shown that there are no other substrings of word that are divisible.

// Example 3:
// Input: word = "abcd"
// Output: 6
// Explanation: The 6 divisible substrings are: "a", "b", "c", "d", "ab", "cd".
// It can be shown that there are no other substrings of word that are divisible.
 
// Constraints:
//     1 <= word.length <= 2000
//     word consists only of lowercase English letters.

import "fmt"

func countDivisibleSubstrings(word string) int {
    mp := map[string]int{ "a": 1, "b" : 1}
    res, c, t, n:= 0, 0, 2, len(word)
    for i := 'c'; i <= 'z'; i++ {
        mp[string(i)] = t
        c++
        if c == 3 { // 逢 3
            c = 0
            t++
        }
    }
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += mp[word[j:j+1]]
            if sum % (j - i + 1) == 0 {
                res++
            }
        }
    }
    return res
}

func countDivisibleSubstrings1(word string) int {
    mp := make([]int, 26)
    mp[0], mp[1] = 1, 1
    res, c, t, n:= 0, 0, 2, len(word)
    for i := 'c'; i <= 'z'; i++ {
        mp[int(i - 'a')] = t
        c++
        if c == 3 { // 逢 3
            c = 0
            t++
        }
    }
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += mp[int(word[j] - 'a')]
            if sum % (j - i + 1) == 0 {
                res++
            }
        }
    }
    return res
}

func countDivisibleSubstrings2(word string) int {
    // 前缀和 + 转换 O(n^2)
    // 字符转数字
    // 因为偏移了1位(1的数量为2个,所以 (ch-'a'+1)/3 +1(后面的1因为是从0开始)
    ord := func(ch byte) int { return int(ch-'a'+1)/3 + 1 }
    res, n := 0, len(word)
    for i := range word { // sum[i,j]是否能被
        sum := 0
        for j := i; j < n; j++ {
            sum += ord(word[j])
            if sum % (j - i + 1) == 0 {
                res++
            }
        }
    }
    return res
}

func countDivisibleSubstrings3(word string) int{
    // 前缀和 + 哈希表 + 细分 + 转换
    // 转换为数字之后,数字范围为 1-9, 即求 [i:j]的和是否是 1x 2x 3x ...9x, 分别讨论
    // 求[i:j]是否为 1x, 那么所有数字-1后, 即求 [i:j]的sum==0 => 即求 sum[i:j]==0 =>求 sum[0:i) == sum[0:j]
    res, n := 0, len(word)
    nums := make([]int, n)
    for i, ch := range word {
        nums[i] = int(ch - 'a' + 1) / 3 + 1
    }
    for d := 1; d <= 9; d++ {
        base := (d - 1) * n // 使用数组代替map, sum[0:n]的范围是[n:9n], 因为要每轮最多-dn,所以用base偏移回来,即将[n-dn:9n-dn]偏移为[0,8n],偏移为(d-1)n
        cnt := make([]int, 8*n+1)
        cnt[base] = 1 // 零宽前缀
        sum := 0
        for _, v := range nums {
            sum += v - d
            res += cnt[base+sum] // sum(i:j]==0 等于 sum[0:j]==sum[0:i]
            cnt[base+sum]++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Substring	Mapped	    Sum	Length	Divisible?
    //     a          1        1    1      Yes
    //     s          7        7    1      Yes
    //     d          2        2    1      Yes
    //     f          3        3    1      Yes
    //    as          1,7      8    2      Yes
    //    sd          7,2      9    2      No
    //    df          2,3      5    2      No
    //   asd          1,7,2    10   3      No
    //   sdf          7,2,3    12   3      Yes
    //  asdf          1,7,2,3  13   4      No
    // Input: word = "asdf"
    // Output: 6
    // Explanation: The table above contains the details about every substring of word, and we can see that 6 of them are divisible.
    fmt.Println(countDivisibleSubstrings("asdf")) // 6
    // Example 2:
    // Input: word = "bdh"
    // Output: 4
    // Explanation: The 4 divisible substrings are: "b", "d", "h", "bdh".
    // It can be shown that there are no other substrings of word that are divisible.
    fmt.Println(countDivisibleSubstrings("bdh")) // 4
    // Example 3:
    // Input: word = "abcd"
    // Output: 6
    // Explanation: The 6 divisible substrings are: "a", "b", "c", "d", "ab", "cd".
    // It can be shown that there are no other substrings of word that are divisible.
    fmt.Println(countDivisibleSubstrings("abcd")) // 6

    fmt.Println(countDivisibleSubstrings1("asdf")) // 6
    fmt.Println(countDivisibleSubstrings1("bdh")) // 4
    fmt.Println(countDivisibleSubstrings1("abcd")) // 6

    fmt.Println(countDivisibleSubstrings2("asdf")) // 6
    fmt.Println(countDivisibleSubstrings2("bdh")) // 4
    fmt.Println(countDivisibleSubstrings2("abcd")) // 6

    fmt.Println(countDivisibleSubstrings3("asdf")) // 6
    fmt.Println(countDivisibleSubstrings3("bdh")) // 4
    fmt.Println(countDivisibleSubstrings3("abcd")) // 6
}