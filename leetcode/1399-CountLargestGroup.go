package main

// 1399. Count Largest Group
// You are given an integer n.
// Each number from 1 to n is grouped according to the sum of its digits.
// Return the number of groups that have the largest size.

// Example 1:
// Input: n = 13
// Output: 4
// Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
// [1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
// There are 4 groups with largest size.

// Example 2:
// Input: n = 2
// Output: 2
// Explanation: There are 2 groups [1], [2] of size 1.

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func countLargestGroup(n int) int {
    mp := make(map[int][]int)
    digitSum := func(n int) int {
        res := 0
        for n > 0 {
            res += n % 10
            n /= 10
        }
        return res
    }
    for i := 1; i <= n; i++ {
        ds := digitSum(i)
        mp[ds] = append(mp[ds], i)
    }
    res, curr := 0, 0
    for _, v := range mp {
        if len(v) == curr {
            res++
        } else if len(v) > curr{
            curr, res = len(v), 1
        }
    }
    return res
}

func countLargestGroup1(n int) int {
    mp := [37]int{}
    for i := 1; i <= n; i++ {
        s := 0
        for k := i; k > 0; k /= 10 {
            s += k % 10
        }
        mp[s]++
    }
    res, mx := 0, 0
    for _, v := range mp {
        if v > mx {
            mx, res = v, 1
        } else if v == mx {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 13
    // Output: 4
    // Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
    // [1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
    // There are 4 groups with largest size.
    fmt.Println(countLargestGroup(13)) // 4
    // Example 2:
    // Input: n = 2
    // Output: 2
    // Explanation: There are 2 groups [1], [2] of size 1.
    fmt.Println(countLargestGroup(2)) // 2

    fmt.Println(countLargestGroup1(13)) // 4
    fmt.Println(countLargestGroup1(2)) // 2
}