package main

// 4081. Prime Subset Selection II
// You are given two integers n and s.

// Consider all prime numbers less than or equal to n. 
// Select a subset of these primes such that:
//     1. Each prime can be selected at most once.
//     2. The sum of the selected primes is as large as possible while remaining less than or equal to s.
//     3. Among subsets with the maximum sum, select one containing the fewest primes.
//     4. If multiple subsets remain tied, select the one whose elements, written in increasing order, form the lexicographically smallest array.

// Return an integer array containing the selected primes in increasing order. 
// If no valid non-empty subset exists, return an empty array.

// Example 1:
// Input: n = 7, s = 20
// Output: [2, 3, 5, 7]
// Explanation:
// - The prime numbers less than or equal to n = 7 are 2, 3, 5, and 7.
// - The sum of all these primes is 2 + 3 + 5 + 7 = 17, which is less than or equal to s = 20.
// - Thus, the answer is [2, 3, 5, 7].

// Example 2:
// Input: n = 15, s = 18
// Output: [5,13]
// Explanation:
// - The prime numbers less than or equal to 15 are 2, 3, 5, 7, 11, and 13.
// - The subsets with sum 18 are [5, 13], [7, 11], [2, 3, 13], and [2, 5, 11]. 
//   Since this sum equals s, it is the maximum possible sum.
// - Among these subsets, [5, 13] and [7, 11] contain the fewest primes: two.
// - Between these two subsets, [5, 13] is lexicographically smaller because its first element, 5, is smaller than 7. 
//   Thus, the answer is [5, 13].

// Example 3:
// Input: n = 5, s = 1
// Output: []
// Explanation:
// - The smallest prime number is 2, which is greater than s = 1.
// - Thus, no prime can be selected, so the answer is an empty array.

// Constraints:
//     - 1 <= n <= 1000
//     - 1 <= s <= 1000

import "fmt"
import "sort"

const MX = 1001
var np = [MX]bool{}
var primes []int

func init() {
    for i := 2; i < MX; i++ {
        if !np[i] {
            primes = append(primes, i)
            for j := i * i; j < MX; j += i {
                np[j] = true
            }
        }
    }
}

func maxPrimeSubset(n int, s int) []int {
    res := []int{}
    m := sort.SearchInts(primes, n + 1)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, s+1)
        for j := range f[i] {
            f[i][j] = 1 << 61 // 避免加法溢出
        }
    }
    f[m][0] = 0
    // 倒着 DP，这样可以正序输出方案，且保证越靠前的质数越小
    for i := m - 1; i >= 0; i-- {
        x := primes[i]
        for j := range s + 1 {
            if j < x {
                f[i][j] = f[i+1][j]
            } else {
                f[i][j] = min(f[i+1][j], f[i+1][j-x]+1)
            }
        }
    }
    j := s
    for j > 0 && f[0][j] == 1 << 61 {
        j--
    }
    for i := 0; j > 0; i++ {
        x := primes[i]
        // 靠前的质数越小越好，f[i+1][j-x]+1 = f[i+1][j] 的时候要选
        if f[i+1][j-x]+1 <= f[i+1][j] {
            res = append(res, x)
            j -= x
        }
    }
    return res
}

func maxPrimeSubset1(n int,s int)[]int{
    val, mn := 0, min(n,s)
    arr, visited :=[ ]int{}, make([]bool,mn + 1)
    for i := 2; i <= mn; i++ {
        if visited[i]{
            continue
        }
        arr = append(arr, i)
        val += i
        for j := i * i; j <= mn; j += i {
            visited[j] = true
        }
    }
    if val <= s {
        return arr
    }
    m := len(arr)
    d := make([]uint8,s + 1)
    for j := 1; j <=s ;j++ {
        d[j] = uint8(m + 1)
    }
    t := make([][]uint64, m)
    for i := m - 1; i >= 0; i-- {
        p := arr[i]
        t[i] = make([]uint64,s/64 + 1)
        for j := s; j >= p; j-- {
            if d[j-p] + 1 <= d[j] {
                d[j] = d[j-p]+1
                t[i][j>>6] |= uint64(1) << (j & 63)
            }
        }
    }
    for d[s] == uint8(m + 1) {
        s--
    }
    res := make([]int,0,int(d[s]))
    for i,p := range arr {
        if t[i][s >> 6] >> (s&63)&1 != 0 {
            res = append(res,p)
            s -= p
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 7, s = 20
    // Output: [2, 3, 5, 7]
    // Explanation:
    // - The prime numbers less than or equal to n = 7 are 2, 3, 5, and 7.
    // - The sum of all these primes is 2 + 3 + 5 + 7 = 17, which is less than or equal to s = 20.
    // - Thus, the answer is [2, 3, 5, 7].
    fmt.Println(maxPrimeSubset(7, 20)) // [2, 3, 5, 7]
    // Example 2:
    // Input: n = 15, s = 18
    // Output: [5,13]
    // Explanation:
    // - The prime numbers less than or equal to 15 are 2, 3, 5, 7, 11, and 13.
    // - The subsets with sum 18 are [5, 13], [7, 11], [2, 3, 13], and [2, 5, 11]. 
    //   Since this sum equals s, it is the maximum possible sum.
    // - Among these subsets, [5, 13] and [7, 11] contain the fewest primes: two.
    // - Between these two subsets, [5, 13] is lexicographically smaller because its first element, 5, is smaller than 7. 
    //   Thus, the answer is [5, 13].
    fmt.Println(maxPrimeSubset(15, 18)) // [5, 13]
    // Example 3:
    // Input: n = 5, s = 1
    // Output: []
    // Explanation:
    // - The smallest prime number is 2, which is greater than s = 1.
    // - Thus, no prime can be selected, so the answer is an empty array.
    fmt.Println(maxPrimeSubset(5, 1)) // []

    fmt.Println(maxPrimeSubset(1, 1)) // []
    fmt.Println(maxPrimeSubset(1, 1000)) // []
    fmt.Println(maxPrimeSubset(1000, 1)) // []
    fmt.Println(maxPrimeSubset(1000, 1000)) // [3 997]

    fmt.Println(maxPrimeSubset1(7, 20)) // [2, 3, 5, 7]
    fmt.Println(maxPrimeSubset1(15, 18)) // [5, 13]
    fmt.Println(maxPrimeSubset1(5, 1)) // []
    fmt.Println(maxPrimeSubset1(1, 1)) // []
    fmt.Println(maxPrimeSubset1(1, 1000)) // []
    fmt.Println(maxPrimeSubset1(1000, 1)) // []
    fmt.Println(maxPrimeSubset1(1000, 1000)) // [3 997]
}