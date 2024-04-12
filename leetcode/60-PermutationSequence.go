package main 

// 60. Permutation Sequence
// The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
//     "123"
//     "132"
//     "213"
//     "231"
//     "312"
//     "321"
// Given n and k, return the kth permutation sequence.

// Example 1:
// Input: n = 3, k = 3
// Output: "213"

// Example 2:
// Input: n = 4, k = 9
// Output: "2314"

// Example 3:
// Input: n = 3, k = 1
// Output: "123"
 
// Constraints:
//     1 <= n <= 9
//     1 <= k <= n!

import "fmt"
import "strconv"

// dp
func getPermutation(n int, k int) string {
    if k == 0 {
        return ""
    }
    used, p, res := make([]bool, n), []int{}, ""
    var findPermutation func(n, index int, k *int, p []int, res *string, used *[]bool)
    findPermutation = func(n, index int, k *int, p []int, res *string, used *[]bool) {
        if index == n {
            *k--
            if *k == 0 {
                for _, v := range p {
                    *res += strconv.Itoa(v + 1)
                }
            }
            return
        }
        for i := 0; i < n; i++ {
            if !(*used)[i] {
                (*used)[i] = true
                p = append(p, i)
                findPermutation(n, index+1, k, p, res, used)
                p = p[:len(p)-1]
                (*used)[i] = false
            }
        }
    }
    findPermutation(n, 0, &k, p, &res, &used)
    return res
}

func getPermutation1(n int, k int) string {
    factorial := make([]int, n)
    factorial[0] = 1
    for i := 1; i < n; i++ {
        factorial[i] = factorial[i - 1] * i
    }
    k--
    valid, res := make([]int, n + 1), ""
    for i := 0; i < len(valid); i++ {
        valid[i] = 1
    }
    for i := 1; i <= n; i++ {
        order := k / factorial[n - i] + 1
        for j := 1; j <= n; j++ {
            order -= valid[j]
            if order == 0 {
                res += strconv.Itoa(j)
                valid[j] = 0
                break
            }
        }
        k %= factorial[n - i]
    }
    return res
}

func main() {
    fmt.Println(getPermutation(3,3)) // "213"
    fmt.Println(getPermutation(4,9)) // "2314"
    fmt.Println(getPermutation(3,1)) // "123"

    fmt.Println(getPermutation1(3,3)) // "213"
    fmt.Println(getPermutation1(4,9)) // "2314"
    fmt.Println(getPermutation1(3,1)) // "123"
}