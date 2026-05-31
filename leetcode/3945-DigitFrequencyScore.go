package main

// 3945. Digit Frequency Score
// You are given an integer n.

// The score of n is defined as the sum of d * freq(d) over all distinct digits d, where freq(d) denotes the number of times the digit d appears in n.

// Return an integer denoting the score of n.

// Example 1:
// Input: n = 122
// Output: 5
// Explanation:
// The digit 1 appears 1 time, contributing 1 * 1 = 1.
// The digit 2 appears 2 times, contributing 2 * 2 = 4.
// Thus, the score of n is 1 + 4 = 5.

// Example 2:
// Input: n = 101
// Output: 2
// Explanation:
// The digit 0 appears 1 time, contributing 0 * 1 = 0.
// The digit 1 appears 2 times, contributing 1 * 2 = 2.
// Thus, the score of n is 2.
 
// Constraints:
//     1 <= n <= 10^9

import "fmt"
import "strconv"

func digitFrequencyScore(n int) int {
    res, s, mp := 0, strconv.Itoa(n), map[string]int{}
    for _, v := range s {
        mp[string(v)]++
    }
    for k, v := range mp {
        val, _ := strconv.Atoi(string(k))
        res += val * v
    }
    return res    
}

func digitFrequencyScore1(n int) int {
    res, mp := 0, make(map[int]int)
    for n > 0 {
        m := n % 10
        if _, ok := mp[m];ok {
            mp[m]++
        }else{
            mp[m] =1
        }
        n = n / 10
    }
    for i := 1; i < 10; i++ {
        res += (mp[i] * i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 122
    // Output: 5
    // Explanation:
    // The digit 1 appears 1 time, contributing 1 * 1 = 1.
    // The digit 2 appears 2 times, contributing 2 * 2 = 4.
    // Thus, the score of n is 1 + 4 = 5.
    fmt.Println(digitFrequencyScore(122)) // 5
    // Example 2:
    // Input: n = 101
    // Output: 2
    // Explanation:
    // The digit 0 appears 1 time, contributing 0 * 1 = 0.
    // The digit 1 appears 2 times, contributing 1 * 2 = 2.
    // Thus, the score of n is 2.
    fmt.Println(digitFrequencyScore(101)) // 2

    fmt.Println(digitFrequencyScore(1)) // 1
    fmt.Println(digitFrequencyScore(2)) // 2
    fmt.Println(digitFrequencyScore(3)) // 3
    fmt.Println(digitFrequencyScore(8)) // 8
    fmt.Println(digitFrequencyScore(64)) // 10
    fmt.Println(digitFrequencyScore(99)) // 18
    fmt.Println(digitFrequencyScore(100)) // 1
    fmt.Println(digitFrequencyScore(101)) // 2
    fmt.Println(digitFrequencyScore(999)) // 27
    fmt.Println(digitFrequencyScore(1024)) // 7
    fmt.Println(digitFrequencyScore(999_999_999)) // 81
    fmt.Println(digitFrequencyScore(1_000_000_000)) // 1

    fmt.Println(digitFrequencyScore1(122)) // 5
    fmt.Println(digitFrequencyScore1(101)) // 2
    fmt.Println(digitFrequencyScore1(1)) // 1
    fmt.Println(digitFrequencyScore1(2)) // 2
    fmt.Println(digitFrequencyScore1(3)) // 3
    fmt.Println(digitFrequencyScore1(8)) // 8
    fmt.Println(digitFrequencyScore1(64)) // 10
    fmt.Println(digitFrequencyScore1(99)) // 18
    fmt.Println(digitFrequencyScore1(100)) // 1
    fmt.Println(digitFrequencyScore1(101)) // 2
    fmt.Println(digitFrequencyScore1(999)) // 27
    fmt.Println(digitFrequencyScore1(1024)) // 7
    fmt.Println(digitFrequencyScore1(999_999_999)) // 81
    fmt.Println(digitFrequencyScore1(1_000_000_000)) // 1
}