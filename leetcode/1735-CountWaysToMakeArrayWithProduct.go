package main

// 1735. Count Ways to Make Array With Product
// You are given a 2D integer array, queries. For each queries[i], where queries[i] = [ni, ki], 
// find the number of different ways you can place positive integers into an array of size ni such that the product of the integers is ki. 
// As the number of ways may be too large, the answer to the ith query is the number of ways modulo 10^9 + 7.

// Return an integer array answer where answer.length == queries.length, 
// and answer[i] is the answer to the ith query.

// Example 1:
// Input: queries = [[2,6],[5,1],[73,660]]
// Output: [4,1,50734910]
// Explanation: Each query is independent.
// [2,6]: There are 4 ways to fill an array of size 2 that multiply to 6: [1,6], [2,3], [3,2], [6,1].
// [5,1]: There is 1 way to fill an array of size 5 that multiply to 1: [1,1,1,1,1].
// [73,660]: There are 1050734917 ways to fill an array of size 73 that multiply to 660. 1050734917 modulo 10^9 + 7 = 50734910.

// Example 2:
// Input: queries = [[1,1],[2,2],[3,3],[4,4],[5,5]]
// Output: [1,2,3,10,5]
 
// Constraints:
//     1 <= queries.length <= 10^4
//     1 <= ni, ki <= 10^4

import "fmt"

// func waysToFillArray(queries [][]int) []int {
//     primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
//     // 计算组合数的递归函数
//     var comb func(n, k int) int
//     comb = func(n, k int) int {
//         if k > n || k == 0 {
//             return 0
//         }
//         if k == n {
//             return 1
//         }
//         return comb(n-1, k) + comb(n-1, k-1)
//     }
//     nK := func(n int, k int) int {
//         res := 1
//         for _, v := range primes {
//             r := 0
//             for k % v == 0 {
//                 r += 1
//                 k /= v 
//             }
//             res *= comb(n - 1 + r, r)
//         }
//         if k != 1 {
//             res *= n
//         }
//         return res % 1000000007
//     }
//     res := make([]int,len(queries))
//     for i := 0; i < len(queries); i++ {
//         res[i] = nK(queries[i][0], queries[i][1])
//     }
//     return res
// }

// primes = (2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97)
// class Solution:
//     def waysToFillArray(self, queries: List[List[int]]) -> List[int]:
//         def nK(n: int, k: int) -> int:
//             res = 1
//             for p in primes:
//                 r = 0
//                 while k % p == 0:
//                     r += 1
//                     k /= p
//                 res *= comb(n - 1 + r, r)
//             if (k != 1):
//                 res *= n            
//             return res % 1000000007
//         return [nK(n, k) for n, k in queries]

func waysToFillArray(queries [][]int) []int {
    const MOD int = 1e9 + 7
    const MAXN int = 1e4 + 14
    const MAXM int = 14
    var comb [MAXN][MAXM]int64
    comb[0][0] = 1
    for i := 1; i < MAXN; i++ {
        comb[i][0] = 1
        for j := 1; j <= i && j < MAXM; j++ {
            comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % int64(MOD)
        }
    }

    var res []int
    for _, q := range queries {
        n, k := q[0], q[1]
        tot := int64(1)
        for i := 2; i*i <= k; i++ {
            if k%i == 0 {
                cnt := 0
                for k%i == 0 {
                    k /= i
                    cnt++
                }
                tot = (tot * comb[n+cnt-1][cnt]) % int64(MOD)
            }
        }
        // k 自身为质数
        if k > 1 {
            tot = tot * int64(n) % int64(MOD)
        }
        res = append(res, int(tot))
    }
    return res
}

const mod = 1e9+7
var comb [][]int
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func init() {
    comb = make([][]int, 10013)
    for i := range comb {
        comb[i] = make([]int, 14)
    }
    comb[0][0] = 1
    for n := 1; n < 10013; n++ {
        comb[n][0] = 1
        for r := 1; r < 14; r++ {
            comb[n][r] = comb[n-1][r-1] + comb[n-1][r]
            comb[n][r] %= mod
        }
    }
}

func waysToFillArray1(queries [][]int) []int {
    res := make([]int, len(queries))
    for i, q := range queries {
        res[i] = 1
        n, k := q[0], q[1]
        for _, p := range primes {
            var r int
            for k % p == 0 {
                r++
                k /= p
            }
            res[i] *= comb[n-1+r][r]
            res[i] %= mod
        }
        if k > 1 {
            res[i] *= n
            res[i] %= mod
        }
    }
    return res
}

func main() {
    // Explanation: Each query is independent.
    // [2,6]: There are 4 ways to fill an array of size 2 that multiply to 6: [1,6], [2,3], [3,2], [6,1].
    // [5,1]: There is 1 way to fill an array of size 5 that multiply to 1: [1,1,1,1,1].
    // [73,660]: There are 1050734917 ways to fill an array of size 73 that multiply to 660. 1050734917 modulo 10^9 + 7 = 50734910.
    fmt.Println(waysToFillArray([][]int{{2,6},{5,1},{73,660}})) // [4,1,50734910]
    fmt.Println(waysToFillArray([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}})) // [1,2,3,10,5]

    fmt.Println(waysToFillArray1([][]int{{2,6},{5,1},{73,660}})) // [4,1,50734910]
    fmt.Println(waysToFillArray1([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}})) // [1,2,3,10,5]
}