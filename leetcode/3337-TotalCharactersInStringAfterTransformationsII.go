package main

// 3337. Total Characters in String After Transformations II
// You are given a string s consisting of lowercase English letters, 
// an integer t representing the number of transformations to perform, and an array nums of size 26. 
// In one transformation, every character in s is replaced according to the following rules:
//     1. Replace s[i] with the next nums[s[i] - 'a'] consecutive characters in the alphabet. 
//        For example, if s[i] = 'a' and nums[0] = 3, the character 'a' transforms into the next 3 consecutive characters ahead of it, which results in "bcd".
//     2. The transformation wraps around the alphabet if it exceeds 'z'. 
//        For example, if s[i] = 'y' and nums[24] = 3, the character 'y' transforms into the next 3 consecutive characters ahead of it, which results in "zab".

// Return the length of the resulting string after exactly t transformations.

// Since the answer may be very large, return it modulo 109 + 7.

// Example 1:
// Input: s = "abcyy", t = 2, nums = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2]
// Output: 7
// Explanation:
// 1. First Transformation (t = 1):
//     'a' becomes 'b' as nums[0] == 1
//     'b' becomes 'c' as nums[1] == 1
//     'c' becomes 'd' as nums[2] == 1
//     'y' becomes 'z' as nums[24] == 1
//     'y' becomes 'z' as nums[24] == 1
//     String after the first transformation: "bcdzz"
// 2. Second Transformation (t = 2):
//     'b' becomes 'c' as nums[1] == 1
//     'c' becomes 'd' as nums[2] == 1
//     'd' becomes 'e' as nums[3] == 1
//     'z' becomes 'ab' as nums[25] == 2
//     'z' becomes 'ab' as nums[25] == 2
//     String after the second transformation: "cdeabab"
// 3. Final Length of the string: The string is "cdeabab", which has 7 characters.

// Example 2:
// Input: s = "azbk", t = 1, nums = [2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2]
// Output: 8
// Explanation:
// 1. First Transformation (t = 1):
//     'a' becomes 'bc' as nums[0] == 2
//     'z' becomes 'ab' as nums[25] == 2
//     'b' becomes 'cd' as nums[1] == 2
//     'k' becomes 'lm' as nums[10] == 2
//     String after the first transformation: "bcabcdlm"
// 2. Final Length of the string: The string is "bcabcdlm", which has 8 characters.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.
//     1 <= t <= 10^9
//     nums.length == 26
//     1 <= nums[i] <= 25

import "fmt"

const mod = 1_000_000_007

type Matrix [][]int

func newMatrix(n, m int) Matrix {
    arr := make(Matrix, n)
    for i := range arr {
        arr[i] = make([]int, m)
    }
    return arr
}

func (a Matrix) mul(b Matrix) Matrix {
    c := newMatrix(len(a), len(b[0]))
    for i, row := range a {
        for k, x := range row {
            if x == 0 { continue }
            for j, y := range b[k] {
                c[i][j] = (c[i][j] + x * y) % mod
            }
        }
    }
    return c
}

// a^n * f0
func (a Matrix) powMul(n int, f0 Matrix) Matrix {
    res := f0
    for ; n > 0; n /= 2 {
        if n%2 > 0 {
            res = a.mul(res)
        }
        a = a.mul(a)
    }
    return res
}

func lengthAfterTransformations(s string, t int, nums []int) int {
    const size = 26
    f0 := newMatrix(size, 1)
    for i := range f0 {
        f0[i][0] = 1
    }
    m := newMatrix(size, size)
    for i, c := range nums {
        for j := i + 1; j <= i+c; j++ {
            m[i][j%size] = 1
        }
    }
    m = m.powMul(t, f0)
    res, cnt := 0, [26]int{}
    for _, c := range s {
        cnt[c-'a']++
    }
    for i, row := range m {
        res += row[0] * cnt[i]
    }
    return res % mod
}

func lengthAfterTransformations1(s string, t int, nums []int) int {
    res,  mod := 0, 1_000_000_007
    count := make([]int, 26)
    for _, v := range s {
        count[v-'a']++
    }
    matrix := make([][]int, 26)
    for i := range matrix {
        matrix[i] = make([]int, 26)
    }
    for i := 0; i < 26; i++ {
        for j := 0; j < nums[i]; j++ {
            matrix[(i + j + 1) % 26][i]++
        }
    }

    matrixMultiple := func(a, b [][]int) [][]int {
        res := make([][]int, 26)
        for i := range res {
            res[i] = make([]int, 26)
        }
        for i := 0; i < 26; i++ {
            for j := 0; j < 26; j++ {
                if a[i][j] != 0 {
                    for k := 0; k < 26; k++ {
                        res[i][k] = (res[i][k] + a[i][j] * b[j][k]) % mod
                    }
                }
            }
        }
        return res
    }
    matrixPow := func(a [][]int, n int) [][]int {
        res := make([][]int, 26)
        for i := range res {
            res[i] = make([]int, 26)
            res[i][i] = 1
        }
        for n > 0 {
            if n % 2 == 1 {
                res = matrixMultiple(a, res)
            }
            a = matrixMultiple(a, a)
            n /= 2
        }
        return res
    }
    nm := matrixPow(matrix, t)
    count1 := make([]int, 26)
    for i := 0; i < 26; i++ {
        for j := 0; j < 26; j++ {
            count1[i] = (count1[i] + nm[i][j] * count[j]) % mod
        }
    }
    for _, v := range count1 {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcyy", t = 2, nums = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2]
    // Output: 7
    // Explanation:
    // 1. First Transformation (t = 1):
    //     'a' becomes 'b' as nums[0] == 1
    //     'b' becomes 'c' as nums[1] == 1
    //     'c' becomes 'd' as nums[2] == 1
    //     'y' becomes 'z' as nums[24] == 1
    //     'y' becomes 'z' as nums[24] == 1
    //     String after the first transformation: "bcdzz"
    // 2. Second Transformation (t = 2):
    //     'b' becomes 'c' as nums[1] == 1
    //     'c' becomes 'd' as nums[2] == 1
    //     'd' becomes 'e' as nums[3] == 1
    //     'z' becomes 'ab' as nums[25] == 2
    //     'z' becomes 'ab' as nums[25] == 2
    //     String after the second transformation: "cdeabab"
    // 3. Final Length of the string: The string is "cdeabab", which has 7 characters.
    fmt.Println(lengthAfterTransformations("abcyy", 2, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2})) // 7
    // Example 2:
    // Input: s = "azbk", t = 1, nums = [2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2]
    // Output: 8
    // Explanation:
    // 1. First Transformation (t = 1):
    //     'a' becomes 'bc' as nums[0] == 2
    //     'z' becomes 'ab' as nums[25] == 2
    //     'b' becomes 'cd' as nums[1] == 2
    //     'k' becomes 'lm' as nums[10] == 2
    //     String after the first transformation: "bcabcdlm"
    // 2. Final Length of the string: The string is "bcabcdlm", which has 8 characters.
    fmt.Println(lengthAfterTransformations("azbk", 1, []int{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2})) // 8

    fmt.Println(lengthAfterTransformations1("abcyy", 2, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2})) // 7
    fmt.Println(lengthAfterTransformations1("azbk", 1, []int{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2})) // 8
}