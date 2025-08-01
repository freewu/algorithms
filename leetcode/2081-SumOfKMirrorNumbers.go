package main

// 2081. Sum of k-Mirror Numbers
// A k-mirror number is a positive integer without leading zeros that reads the same both forward and backward in base-10 as well as in base-k.
//     1. For example, 9 is a 2-mirror number. 
//        The representation of 9 in base-10 and base-2 are 9 and 1001 respectively, 
//        which read the same both forward and backward.
//     2. On the contrary, 4 is not a 2-mirror number. 
//        The representation of 4 in base-2 is 100, which does not read the same both forward and backward.

// Given the base k and the number n, return the sum of the n smallest k-mirror numbers.

// Example 1:
// Input: k = 2, n = 5
// Output: 25
// Explanation:
// The 5 smallest 2-mirror numbers and their representations in base-2 are listed as follows:
//   base-10    base-2
//     1          1
//     3          11
//     5          101
//     7          111
//     9          1001
// Their sum = 1 + 3 + 5 + 7 + 9 = 25. 

// Example 2:
// Input: k = 3, n = 7
// Output: 499
// Explanation:
// The 7 smallest 3-mirror numbers are and their representations in base-3 are listed as follows:
//   base-10    base-3
//     1          1
//     2          2
//     4          11
//     8          22
//     121        11111
//     151        12121
//     212        21212
// Their sum = 1 + 2 + 4 + 8 + 121 + 151 + 212 = 499.

// Example 3:
// Input: k = 7, n = 17
// Output: 20379000
// Explanation: The 17 smallest 7-mirror numbers are:
// 1, 2, 3, 4, 5, 6, 8, 121, 171, 242, 292, 16561, 65656, 2137312, 4602064, 6597956, 6958596

// Constraints:
//     2 <= k <= 9
//     1 <= n <= 30

import "fmt"

func kMirror(k int, n int) int64 {
    check := func(i, k int) bool {
        tmp := []int{}
        for i != 0 {
            tmp = append(tmp, i % k)
            i /= k
        }
        for i := 0 ; i < len(tmp) / 2 ; i++ {
            if tmp[i] != tmp[len(tmp)-1-i] {
                return false
            }
        }
        return true
    }
    res, l := 0, 1 
    for n > 0 {
        r := l * 10 
        for i := 0 ; i < 2 ; i++ {
            for j := l ; j < r ; j++ {
                cur, newNum := j, j 
                if i == 0 {
                    cur /= 10
                }
                for cur != 0 {
                    newNum = newNum * 10 + cur % 10
                    cur /= 10
                }
                if check(newNum, k) {
                    res += newNum
                    n--
                    if n == 0 { return int64(res) }
                }
            }
        }
        l = r
    }
    return int64(res)
}

const MX = 30
var res [10][]int

// 力扣 9. 回文数
func isKPalindrome(x, k int) bool {
    if x % k == 0 { return false }
    rev := 0
    for rev < x / k {
        rev = rev  *k + x % k
        x /= k
    }
    return rev == x || rev == x / k
}

func doPalindrome(x int) bool {
    done := true
    for k := 2; k < 10; k++ {
        if len(res[k]) < MX && isKPalindrome(x, k) {
            res[k] = append(res[k], x)
        }
        if len(res[k]) < MX {
            done = false
        }
    }
    if !done {
        return false
    }
    for k := 2; k < 10; k++ {
        // 计算前缀和 
        for i := 1; i < MX; i++ {
            res[k][i] += res[k][i-1]
        }
    }
    return true
}

func init() {
    for k := 2; k < 10; k++ {
        res[k] = make([]int, 0, MX) // 预分配空间
    }
    for base := 1; ; base *= 10 {
        // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for i := base; i < base * 10; i++ {
            x := i
            for t := i / 10; t > 0; t /= 10 {
                x = x * 10 + t % 10
            }
            if doPalindrome(x) {
                return
            }
        }
        // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for i := base; i < base * 10; i++ {
            x := i
            for t := i; t > 0; t /= 10 {
                x = x * 10 + t % 10
            }
            if doPalindrome(x) {
                return
            }
        }
    }
}

func kMirror1(k, n int) int64 {
    return int64(res[k][n-1])
}

func main() {
    // Example 1:
    // Input: k = 2, n = 5
    // Output: 25
    // Explanation:
    // The 5 smallest 2-mirror numbers and their representations in base-2 are listed as follows:
    //   base-10    base-2
    //     1          1
    //     3          11
    //     5          101
    //     7          111
    //     9          1001
    // Their sum = 1 + 3 + 5 + 7 + 9 = 25. 
    fmt.Println(kMirror(2,5)) // 25
    // Example 2:
    // Input: k = 3, n = 7
    // Output: 499
    // Explanation:
    // The 7 smallest 3-mirror numbers are and their representations in base-3 are listed as follows:
    //   base-10    base-3
    //     1          1
    //     2          2
    //     4          11
    //     8          22
    //     121        11111
    //     151        12121
    //     212        21212
    // Their sum = 1 + 2 + 4 + 8 + 121 + 151 + 212 = 499.
    fmt.Println(kMirror(3,7)) // 499
    // Example 3:
    // Input: k = 7, n = 17
    // Output: 20379000
    // Explanation: The 17 smallest 7-mirror numbers are:
    // 1, 2, 3, 4, 5, 6, 8, 121, 171, 242, 292, 16561, 65656, 2137312, 4602064, 6597956, 6958596
    fmt.Println(kMirror(7,17)) // 20379000

    fmt.Println(kMirror(2,1)) // 1
    fmt.Println(kMirror(9,30)) // 18627530
    fmt.Println(kMirror(9,1)) // 1
    fmt.Println(kMirror(2,30)) // 2609044274

    fmt.Println(kMirror1(2,5)) // 25
    fmt.Println(kMirror1(3,7)) // 499
    fmt.Println(kMirror1(7,17)) // 20379000
    fmt.Println(kMirror1(2,1)) // 1
    fmt.Println(kMirror1(9,30)) // 18627530
    fmt.Println(kMirror1(9,1)) // 1
    fmt.Println(kMirror1(2,30)) // 2609044274
}