package main

// 3234. Count the Number of Substrings With Dominant Ones
// You are given a binary string s.

// Return the number of substrings with dominant ones.

// A string has dominant ones if the number of ones in the string is greater than or equal to the square of the number of zeros in the string.

// Example 1:
// Input: s = "00011"
// Output: 5
// Explanation:
// The substrings with dominant ones are shown in the table below.
// i   |   j   | s[i..j]   | Number of Zeros   | Number of Ones
// 3   |   3   |   1       |       0           |       1
// 4   |   4   |   1       |       0           |       1
// 2   |   3   |   01      |       1           |       1
// 3   |   4   |   11      |       0           |       2
// 2   |   4   |   011     |       1           |       2

// Example 2:
// Input: s = "101101"
// Output: 16
// Explanation:
// The substrings with non-dominant ones are shown in the table below.
// Since there are 21 substrings total and 5 of them have non-dominant ones, it follows that there are 16 substrings with dominant ones.
// i   |   j   |   s[i..j] |   Number of Zeros |   Number of Ones
// 1   |   1   |       0   |       1           |       0
// 4   |   4   |       0   |       1           |       0
// 1   |   4   |   0110    |       2           |       2
// 0   |   4   |   10110   |       2           |       3
// 1   |   5   |   01101   |       2           |       3

// Constraints:
//     1 <= s.length <= 4 * 10^4
//     s consists only of characters '0' and '1'.

import "fmt"

func numberOfSubstrings(s string) int {
    res, n := 0, len(s)
    for zero := 0; zero * zero < n; zero++ {
        last, count := -1, []int{ 0, 0 }
        for start, end := 0, 0; end < n; end++ {
            count[s[end] - '0']++
            for start < end {
                if s[start] == '0' && count[0] > zero {
                    count[0]--
                    last = start
                } else if s[start] == '1' && (count[1] - 1) >= (zero * zero) {
                    count[1]--
                } else {
                    break
                }
                start++
            }
            if count[0] == zero && count[1] >= count[0] * count[0] {
                res += (start - last)
            }
        }
    }
    return res
}

func numberOfSubstrings1(s string) int {
    res, sum1 := 0, 0
    arr := []int{ -1 } // 哨兵
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right, v := range s {
        if v == '0' {
            arr = append(arr, right)
        } else {
            res += right - arr[len(arr) - 1]
            sum1++
        }
        for k := len(arr) - 1; k > 0 && (len(arr) - k) * (len(arr) - k) <= sum1; k-- {
            count0 := len(arr) - k
            count1 := right - arr[k] + 1 - count0
            res += max(arr[k] - arr[k - 1] - max(count0 * count0 - count1, 0), 0)
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: s = "00011"
    // Output: 5
    // Explanation:
    // The substrings with dominant ones are shown in the table below.
    // i   |   j   | s[i..j]   | Number of Zeros   | Number of Ones
    // 3   |   3   |   1       |       0           |       1
    // 4   |   4   |   1       |       0           |       1
    // 2   |   3   |   01      |       1           |       1
    // 3   |   4   |   11      |       0           |       2
    // 2   |   4   |   011     |       1           |       2
    fmt.Println(numberOfSubstrings("00011")) // 5
    // Example 2:
    // Input: s = "101101"
    // Output: 16
    // Explanation:
    // The substrings with non-dominant ones are shown in the table below.
    // Since there are 21 substrings total and 5 of them have non-dominant ones, it follows that there are 16 substrings with dominant ones.
    // i   |   j   |   s[i..j] |   Number of Zeros |   Number of Ones
    // 1   |   1   |       0   |       1           |       0
    // 4   |   4   |       0   |       1           |       0
    // 1   |   4   |   0110    |       2           |       2
    // 0   |   4   |   10110   |       2           |       3
    // 1   |   5   |   01101   |       2           |       3
    fmt.Println(numberOfSubstrings("101101")) // 16

    fmt.Println(numberOfSubstrings("000000000")) // 0
    fmt.Println(numberOfSubstrings("111111111")) // 45
    fmt.Println(numberOfSubstrings("1010101010")) // 18
    fmt.Println(numberOfSubstrings("0101010101")) // 18
    fmt.Println(numberOfSubstrings("0000011111")) // 22
    fmt.Println(numberOfSubstrings("1111100000")) // 22

    fmt.Println(numberOfSubstrings1("00011")) // 5
    fmt.Println(numberOfSubstrings1("101101")) // 16
    fmt.Println(numberOfSubstrings1("000000000")) // 0
    fmt.Println(numberOfSubstrings1("111111111")) // 45
    fmt.Println(numberOfSubstrings1("1010101010")) // 18
    fmt.Println(numberOfSubstrings1("0101010101")) // 18
    fmt.Println(numberOfSubstrings1("0000011111")) // 22
    fmt.Println(numberOfSubstrings1("1111100000")) // 22
}