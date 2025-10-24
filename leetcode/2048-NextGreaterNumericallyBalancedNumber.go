package main

// 2048. Next Greater Numerically Balanced Number
// An integer x is numerically balanced if for every digit d in the number x, 
// there are exactly d occurrences of that digit in x.

// Given an integer n, return the smallest numerically balanced number strictly greater than n.

// Example 1:
// Input: n = 1
// Output: 22
// Explanation: 
// 22 is numerically balanced since:
// - The digit 2 occurs 2 times. 
// It is also the smallest numerically balanced number strictly greater than 1.

// Example 2:
// Input: n = 1000
// Output: 1333
// Explanation: 
// 1333 is numerically balanced since:
// - The digit 1 occurs 1 time.
// - The digit 3 occurs 3 times. 
// It is also the smallest numerically balanced number strictly greater than 1000.
// Note that 1022 cannot be the answer because 0 appeared more than 0 times.

// Example 3:
// Input: n = 3000
// Output: 3133
// Explanation: 
// 3133 is numerically balanced since:
// - The digit 1 occurs 1 time.
// - The digit 3 occurs 3 times.
// It is also the smallest numerically balanced number strictly greater than 3000.

// Constraints:
//     0 <= n <= 10^6

import "fmt"

func nextBeautifulNumber(n int) int {
    isBalance := func(num int) bool {
        cnt := make([]int, 10)
        for ; num > 0; num /= 10 {
            cnt[num%10]++
        }
        for i := range cnt {
            if cnt[i] == 0 { continue }
            if cnt[i] != i { return false }
        }
        return true
    }
    for n++; ; n++ {
        if isBalance(n) {
            return n
        }
    }
}

func nextBeautifulNumber1(n int) int {
    // With these combinations, we have to generate the smallest number with the
    // right digits.
    // For n, the next bigger balanced number will have the same number of digits or
    // just one more digit. For 1, 22. For 99, 122.
    if n == 0 { return 1 }
    combinations := map[int][][]int{
        2: {{2, 2}},
        3: {{1, 2, 2}, {3, 3, 3}},
        4: {{1, 3, 3, 3}, {4, 4, 4, 4}},
        5: {{1, 4, 4, 4, 4}, {2, 2, 3, 3, 3}, {5, 5, 5, 5, 5}},
        6: {{1, 2, 2, 3, 3, 3}, {1, 5, 5, 5, 5, 5}, {2, 2, 4, 4, 4, 4}, {6, 6, 6, 6, 6, 6}},
        7: {{1, 2, 2, 4, 4, 4, 4}, {1, 6, 6, 6, 6, 6, 6}, {2, 2, 5, 5, 5, 5, 5}, {3, 3, 3, 4, 4, 4, 4}, {7, 7, 7, 7, 7, 7, 7}},
    }
    numDigits := 0
    for num := n; num > 0; num /= 10 {
        numDigits++
    }
    res := 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var backtrack func(comb []int, curr int, used int, num int, res *int)
    backtrack = func(comb []int, curr int, used int, num int, res *int) {
        if used == 0 {
            if curr > num {
                *res = min(*res, curr)
            }
        }
        for i := 0; i < len(comb); i++ {
            if used & (1 << i) != 0 {
                backtrack(comb, curr*10 + comb[i], used ^ (1 << i), num, res)
            }
        }
    }
    for _, comb := range combinations[numDigits] {
        backtrack(comb, 0, (1 << numDigits) - 1, n, &res)
    }
    if res == 1 << 31 && numDigits+1 <= 7 {
        for _, comb := range combinations[numDigits+1] {
            backtrack(comb, 0, (1 << (numDigits+1))-1, n, &res)
        }
    }
    return res
}

func nextBeautifulNumber2(n int) int {
    mx, count := 1224444, make([]int, 10)
    for x := n + 1; x <= mx; x++ {
        check := func(num int) bool {
            for i := 0; i < 10; i++ {
                count[i] = 0 
            }
            for num > 0 {
                if num % 10 == 0 {
                    return false 
                }
                count[num % 10] ++ 
                num /= 10 
            }
            for i := 1; i < 10; i++ {
                if count[i] > 0 && count[i] != i {
                    return false 
                }
            }
            return true 
        }
        if (check(x)) {
            return x 
        }
    }
    return mx
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 22
    // Explanation: 
    // 22 is numerically balanced since:
    // - The digit 2 occurs 2 times. 
    // It is also the smallest numerically balanced number strictly greater than 1.
    fmt.Println(nextBeautifulNumber(1)) // 22
    // Example 2:
    // Input: n = 1000
    // Output: 1333
    // Explanation: 
    // 1333 is numerically balanced since:
    // - The digit 1 occurs 1 time.
    // - The digit 3 occurs 3 times. 
    // It is also the smallest numerically balanced number strictly greater than 1000.
    // Note that 1022 cannot be the answer because 0 appeared more than 0 times.
    fmt.Println(nextBeautifulNumber(1000)) // 1333
    // Example 3:
    // Input: n = 3000
    // Output: 3133
    // Explanation: 
    // 3133 is numerically balanced since:
    // - The digit 1 occurs 1 time.
    // - The digit 3 occurs 3 times.
    // It is also the smallest numerically balanced number strictly greater than 3000.
    fmt.Println(nextBeautifulNumber(3000)) // 3133

    fmt.Println(nextBeautifulNumber(0)) // 1
    fmt.Println(nextBeautifulNumber(999)) // 1333
    fmt.Println(nextBeautifulNumber(1024)) // 221333
    fmt.Println(nextBeautifulNumber(999_999)) // 1224444
    fmt.Println(nextBeautifulNumber(1_000_000)) // 1224444

    fmt.Println(nextBeautifulNumber1(1)) // 22
    fmt.Println(nextBeautifulNumber1(1000)) // 1333
    fmt.Println(nextBeautifulNumber1(3000)) // 3133
    fmt.Println(nextBeautifulNumber1(0)) // 1
    fmt.Println(nextBeautifulNumber1(999)) // 1333
    fmt.Println(nextBeautifulNumber1(1024)) // 1333
    fmt.Println(nextBeautifulNumber1(999_999)) // 1224444
    fmt.Println(nextBeautifulNumber1(1_000_000)) // 1224444

    fmt.Println(nextBeautifulNumber2(1)) // 22
    fmt.Println(nextBeautifulNumber2(1000)) // 1333
    fmt.Println(nextBeautifulNumber2(3000)) // 3133
    fmt.Println(nextBeautifulNumber2(0)) // 1
    fmt.Println(nextBeautifulNumber2(999)) // 1333
    fmt.Println(nextBeautifulNumber2(1024)) // 1333
    fmt.Println(nextBeautifulNumber2(999_999)) // 1224444
    fmt.Println(nextBeautifulNumber2(1_000_000)) // 1224444
}