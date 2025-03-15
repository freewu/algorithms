package main

// 3470. Permutations IV
// Given two integers, n and k, an alternating permutation is a permutation of the first n positive integers such that no two adjacent elements are both odd or both even.

// Return the k-th alternating permutation sorted in lexicographical order. 
// If there are fewer than k valid alternating permutations, return an empty list.

// Example 1:
// Input: n = 4, k = 6
// Output: [3,4,1,2]
// Explanation:
// The lexicographically-sorted alternating permutations of [1, 2, 3, 4] are:
// [1, 2, 3, 4]
// [1, 4, 3, 2]
// [2, 1, 4, 3]
// [2, 3, 4, 1]
// [3, 2, 1, 4]
// [3, 4, 1, 2] ← 6th permutation
// [4, 1, 2, 3]
// [4, 3, 2, 1]
// Since k = 6, we return [3, 4, 1, 2].

// Example 2:
// Input: n = 3, k = 2
// Output: [3,2,1]
// Explanation:
// The lexicographically-sorted alternating permutations of [1, 2, 3] are:
// [1, 2, 3]
// [3, 2, 1] ← 2nd permutation
// Since k = 2, we return [3, 2, 1].

// Example 3:
// Input: n = 2, k = 3
// Output: []
// Explanation:
// The lexicographically-sorted alternating permutations of [1, 2] are:
// [1, 2]
// [2, 1]
// There are only 2 alternating permutations, but k = 3, which is out of range. Thus, we return an empty list [].

// Constraints:
//     1 <= n <= 100
//     1 <= k <= 10^15

import "fmt"
import "slices"

func permute(n int, k int64) []int {
    count := make([]int64, 2, n + 1)
    count[0], count[1] = 1, 1
    for i := 2; i <= n && count[i - 1] <= k; i++ {
        if i & 1 == 1 {
            count = append(count, count[i - 1] / 2 * int64(i + 1) / 2)
        } else {
            count = append(count, count[i - 1] * int64(i))
        }
    }
    if len(count) == n + 1 && count[n] < k {
        return nil
    }
    res, visited := make([]int, n), make([]bool, n)
    i, p := 0, 0
    k--
    if n & 1 == 0 {
        m := int64(1)
        if len(count) >= n {
            m = k / count[n - 1] + 1
            k %= count[n - 1]
        }
        res[0] = int(m)
        visited[m - 1] = true
        p = int(m & 1)
        i++
    }
    for ; i < n - 1; i++ {
        m := int64(1)
        if len(count) >= n - i {
            if (n - i) & 1 == 1 {
                m = 2 * k / count[n - 1 - i] + 1
                k %= count[n - 1 - i] / 2
            } else {
                m = k / count[n - 1 - i] + 1
                k %= count[n - 1 - i]
            }
        }
        for j := p; j < n; j += 2 {
            if !visited[j] {
                m--
                if m == 0 {
                    res[i] = j + 1
                    visited[j] = true
                    break
                }
            }
        }
        p ^= 1
    }
    for j := range visited {
        if !visited[j] {
            res[n - 1] = j + 1
        }
    }
    return res
}

// 预处理交替排列的方案数
var f = []int{1}

func init() {
    for i := 1; f[len(f)-1] < 1e15; i++ {
        f = append(f, f[len(f)-1]*i)
        f = append(f, f[len(f)-1]*i)
    }
}

func permute1(n int, K int64) []int {
    // k 改成从 0 开始，方便计算
    k := int(K - 1)
    if n < len(f) && k >= f[n]*(2-n%2) { // n 是偶数的时候，方案数乘以 2
        return nil
    }
    // cand 表示剩余未填入 res 的数字
    // cand[0] 保存偶数，cand[1] 保存奇数
    cand := [2][]int{}
    for i := 2; i <= n; i += 2 {
        cand[0] = append(cand[0], i)
    }
    for i := 1; i <= n; i += 2 {
        cand[1] = append(cand[1], i)
    }
    res, parity := make([]int, n), 1 // 当前要填入 res[i] 的数的奇偶性
    for i := 0; i < n; i++ {
        j := 0
        if n-1-i < len(f) {
            // 比如示例 1，按照第一个数分组，每一组的大小都是 size=2
            // 知道 k 和 size 就知道我们要去哪一组
            size := f[n-1-i]
            j = k / size // 去第 j 组
            k %= size
            // n 是偶数的情况，第一个数既可以填奇数又可以填偶数，要特殊处理
            if n%2 == 0 && i == 0 {
                parity = 1 - j%2
                j /= 2
            }
        } // else j=0，在 n 很大的情况下，只能按照 1,2,3,... 的顺序填
        res[i] = cand[parity][j]
        cand[parity] = slices.Delete(cand[parity], j, j+1)
        parity ^= 1 // 下一个数的奇偶性
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, k = 6
    // Output: [3,4,1,2]
    // Explanation:
    // The lexicographically-sorted alternating permutations of [1, 2, 3, 4] are:
    // [1, 2, 3, 4]
    // [1, 4, 3, 2]
    // [2, 1, 4, 3]
    // [2, 3, 4, 1]
    // [3, 2, 1, 4]
    // [3, 4, 1, 2] ← 6th permutation
    // [4, 1, 2, 3]
    // [4, 3, 2, 1]
    // Since k = 6, we return [3, 4, 1, 2].
    fmt.Println(permute(4, 6)) // [3,4,1,2]
    // Example 2:
    // Input: n = 3, k = 2
    // Output: [3,2,1]
    // Explanation:
    // The lexicographically-sorted alternating permutations of [1, 2, 3] are:
    // [1, 2, 3]
    // [3, 2, 1] ← 2nd permutation
    // Since k = 2, we return [3, 2, 1].
    fmt.Println(permute(3, 2)) // [3,2,1]
    // Example 3:
    // Input: n = 2, k = 3
    // Output: []
    // Explanation:
    // The lexicographically-sorted alternating permutations of [1, 2] are:
    // [1, 2]
    // [2, 1]
    // There are only 2 alternating permutations, but k = 3, which is out of range. Thus, we return an empty list [].
    fmt.Println(permute(2, 3)) // []

    fmt.Println(permute(1, 1)) // [1]
    fmt.Println(permute(100, 10e15)) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 90 85 78 89 80 93 86 83 94 99 88 95 98 87 84 79 92 97 82 91 100 81 96]
    fmt.Println(permute(100, 1)) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100]
    fmt.Println(permute(1, 10e15)) // []

    fmt.Println(permute1(4, 6)) // [3,4,1,2]
    fmt.Println(permute1(3, 2)) // [3,2,1]
    fmt.Println(permute1(2, 3)) // []
    fmt.Println(permute1(1, 1)) // [1]
    fmt.Println(permute1(100, 10e15)) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 90 85 78 89 80 93 86 83 94 99 88 95 98 87 84 79 92 97 82 91 100 81 96]
    fmt.Println(permute1(100, 1)) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100]
    fmt.Println(permute1(1, 10e15)) // []
}