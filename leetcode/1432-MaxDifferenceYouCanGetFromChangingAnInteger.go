package main

// 1432. Max Difference You Can Get From Changing an Integer
// You are given an integer num. 
// You will apply the following steps exactly two times:
//     Pick a digit x (0 <= x <= 9).
//     Pick another digit y (0 <= y <= 9). The digit y can be equal to x.
//     Replace all the occurrences of x in the decimal representation of num by y.
//     The new integer cannot have any leading zeros, also the new integer cannot be 0.

// Let a and b be the results of applying the operations to num the first and second times, respectively.

// Return the max difference between a and b.

// Example 1:
// Input: num = 555
// Output: 888
// Explanation: The first time pick x = 5 and y = 9 and store the new integer in a.
// The second time pick x = 5 and y = 1 and store the new integer in b.
// We have now a = 999 and b = 111 and max difference = 888

// Example 2:
// Input: num = 9
// Output: 8
// Explanation: The first time pick x = 9 and y = 9 and store the new integer in a.
// The second time pick x = 9 and y = 1 and store the new integer in b.
// We have now a = 9 and b = 1 and max difference = 8

// Constraints:
//     1 <= num <= 10^8

import "fmt"
import "strconv"
import "strings"

func maxDiff(num int) int {
    s := strconv.Itoa(num)
    mxs, mns := s, s
    for _, v := range s {
        if v == '9' { continue }
        mxs = strings.ReplaceAll(mxs, string(v), "9") // 都替换成9
        break
    }
    for i, v := range s {
        if v == '1' || v == '0' { continue }
        if i == 0 {
            mns = strings.ReplaceAll(mns, string(v), "1")
        } else {
            mns = strings.ReplaceAll(mns, string(v), "0")
        }
        break
    }
    mn, _ := strconv.Atoi(mns)
    mx, _ := strconv.Atoi(mxs)
    return mx - mn
}

func maxDiff1(num int) int {
    a := []int{}
    t := num
    for t > 0 {
        a = append(a, t%10)
        t /= 10
    }
    x, y := 0, 0
    for i := len(a) - 1; i >= 0; i-- {
        if a[i] > 1 || a[i] >= 1 && a[i] != a[len(a)-1] {
            x = a[i]
            if i == len(a)-1 {
                y = 1
            }
            break
        }
    }
    s1 := 0
    for i := len(a) - 1; i >= 0; i-- {
        if a[i] == x {
            s1 = s1*10 + y
        } else {
            s1 = s1*10 + a[i]
        }
    }
    x, y = 0, 0
    for i := len(a) - 1; i >= 0; i-- {
        if a[i] < 9 {
            x = a[i]
            y = 9
            break
        }
    }
    s2 := 0
    for i := len(a) - 1; i >= 0; i-- {
        if a[i] == x {
            s2 = s2*10 + y
        } else {
            s2 = s2*10 + a[i]
        }
    }
    return s2 - s1
}

func main() {
    // Example 1:
    // Input: num = 555
    // Output: 888
    // Explanation: The first time pick x = 5 and y = 9 and store the new integer in a.
    // The second time pick x = 5 and y = 1 and store the new integer in b.
    // We have now a = 999 and b = 111 and max difference = 888
    fmt.Println(maxDiff(555)) // 888
    // Example 2:
    // Input: num = 9
    // Output: 8
    // Explanation: The first time pick x = 9 and y = 9 and store the new integer in a.
    // The second time pick x = 9 and y = 1 and store the new integer in b.
    // We have now a = 9 and b = 1 and max difference = 8
    fmt.Println(maxDiff(9)) // 8

    fmt.Println(maxDiff(123456)) // 820000
    fmt.Println(maxDiff(10000)) // 80000
    fmt.Println(maxDiff(9288)) // 8700
    fmt.Println(maxDiff(1)) // 8
    fmt.Println(maxDiff(10000000)) // 80000000

    fmt.Println(maxDiff1(555)) // 888
    fmt.Println(maxDiff1(9)) // 8
    fmt.Println(maxDiff1(123456)) // 820000
    fmt.Println(maxDiff1(10000)) // 80000
    fmt.Println(maxDiff1(9288)) // 8700
    fmt.Println(maxDiff1(1)) // 8
    fmt.Println(maxDiff1(10000000)) // 80000000
}