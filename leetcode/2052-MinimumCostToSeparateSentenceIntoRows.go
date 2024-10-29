package main

// 2052. Minimum Cost to Separate Sentence Into Rows
// You are given a string sentence containing words separated by spaces, and an integer k. 
// Your task is to separate sentence into rows where the number of characters in each row is at most k. 
// You may assume that sentence does not begin or end with a space, and the words in sentence are separated by a single space.

// You can split sentence into rows by inserting line breaks between words in sentence. 
// A word cannot be split between two rows. 
// Each word must be used exactly once, and the word order cannot be rearranged. 
// Adjacent words in a row should be separated by a single space, and rows should not begin or end with spaces.

// The cost of a row with length n is (k - n)2, and the total cost is the sum of the costs for all rows except the last one.

// For example if sentence = "i love leetcode" and k = 12:
//     Separating sentence into "i", "love", and "leetcode" has a cost of (12 - 1)2 + (12 - 4)2 = 185.
//     Separating sentence into "i love", and "leetcode" has a cost of (12 - 6)2 = 36.
//     Separating sentence into "i", and "love leetcode" is not possible because the length of "love leetcode" is greater than k.

// Return the minimum possible total cost of separating sentence into rows.

// Example 1:
// Input: sentence = "i love leetcode", k = 12
// Output: 36
// Explanation:
// Separating sentence into "i", "love", and "leetcode" has a cost of (12 - 1)2 + (12 - 4)2 = 185.
// Separating sentence into "i love", and "leetcode" has a cost of (12 - 6)2 = 36.
// Separating sentence into "i", "love leetcode" is not possible because "love leetcode" has length 13.
// 36 is the minimum possible total cost so return it.

// Example 2:
// Input: sentence = "apples and bananas taste great", k = 7
// Output: 21
// Explanation
// Separating sentence into "apples", "and", "bananas", "taste", and "great" has a cost of (7 - 6)2 + (7 - 3)2 + (7 - 7)2 + (7 - 5)2 = 21.
// 21 is the minimum possible total cost so return it.

// Example 3:
// Input: sentence = "a", k = 5
// Output: 0
// Explanation:
// The cost of the last row is not included in the total cost, and since there is only one row, return 0.

// Constraints:
//     1 <= sentence.length <= 5000
//     1 <= k <= 5000
//     The length of each word in sentence is at most k.
//     sentence consists of only lowercase English letters and spaces.
//     sentence does not begin or end with a space.
//     Words in sentence are separated by a single space.

import "fmt"

func minimumCost(sentence string, k int) int {
    count, arr := 0, []int{} 
    for _, v := range sentence { // 统计每个空格距离上一个空格的单词数(题目说明了两个单词之间只有一个空格)
        if v == ' ' {
            arr = append(arr, count)
            count = 0
        } else {
            count++
        }
    }
    arr = append(arr, count) // 结尾也算一个划分点
    n := len(arr)
    dp := make([]int, n + 1)
    for i, v := range arr {
        extra := 0
        if i != n - 1 { // 最后一行划分是没有成本的!! (总成本就是除开最后一行以外的其它所有行的分隔成本之和。)
            extra = (k - v) * (k - v)
        }
        dp[i + 1] = dp[i] + extra // 默认方案:单独自己划分
        sum := v
        for j := i - 1; j >= 0; j-- {
            sum += arr[j] + 1 // 拼接前一个字符串,还会多一个空格
            if sum > k { break }
            if i != n - 1 {
                extra = (k - sum) * (k - sum)
            }
            dp[i+1] = min(dp[i + 1], dp[j]+ extra) // 0...j-1 的范围 还有[j:i]的字符串
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: sentence = "i love leetcode", k = 12
    // Output: 36
    // Explanation:
    // Separating sentence into "i", "love", and "leetcode" has a cost of (12 - 1)2 + (12 - 4)2 = 185.
    // Separating sentence into "i love", and "leetcode" has a cost of (12 - 6)2 = 36.
    // Separating sentence into "i", "love leetcode" is not possible because "love leetcode" has length 13.
    // 36 is the minimum possible total cost so return it.
    fmt.Println(minimumCost("i love leetcode", 12)) // 36
    // Example 2:
    // Input: sentence = "apples and bananas taste great", k = 7
    // Output: 21
    // Explanation
    // Separating sentence into "apples", "and", "bananas", "taste", and "great" has a cost of (7 - 6)2 + (7 - 3)2 + (7 - 7)2 + (7 - 5)2 = 21.
    // 21 is the minimum possible total cost so return it.
    fmt.Println(minimumCost("apples and bananas taste great", 7)) // 21
    // Example 3:
    // Input: sentence = "a", k = 5
    // Output: 0
    // Explanation:
    // The cost of the last row is not included in the total cost, and since there is only one row, return 0.
    fmt.Println(minimumCost("a", 5)) // 0
}