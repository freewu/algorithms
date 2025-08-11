package main

// 3646. Next Special Palindrome Number
// You are given an integer n.

// A number is called special if:
//     1. It is a palindrome.
//     2. Every digit k in the number appears exactly k times.

// Return the smallest special number strictly greater than n.

// Example 1:
// Input: n = 2
// Output: 22
// Explanation:
// 22 is the smallest special number greater than 2, as it is a palindrome and the digit 2 appears exactly 2 times.

// Example 2:
// Input: n = 33
// Output: 212
// Explanation:
// 212 is the smallest special number greater than 33, as it is a palindrome and the digits 1 and 2 appear exactly 1 and 2 times respectively.

// Constraints:
//     0 <= n <= 10^15

import "fmt"
import "sort"
import "slices"
import "math/bits"

var specialNumbers []int

func init() {
    const oddMask = 0x155
    for mask := 1; mask < 512; mask++ {
        t := mask & oddMask
        if t & (t-1) > 0 { continue } // 至少有两个奇数
        odd, size, perm := 0, 0, []int{} // 构造排列 perm
        for s := uint(mask); s > 0; s &= s - 1 {
            x := bits.TrailingZeros(s) + 1
            size += x
            //for range x / 2 {
            for i := 0; i < x / 2; i++ {
                perm = append(perm, x)
            }
            if x%2 > 0 {
                odd = x
            }
        }
        if size > 16 { continue } // 回文串太长了
        permutations(len(perm), len(perm), func(idx []int) bool {
            pal := 0
            for _, i := range idx {
                pal = pal*10 + perm[i]
            }
            v := pal
            if odd > 0 {
                pal = pal*10 + odd
            }
            // 反转 pal 的左半，拼在 pal 后面
            for ; v > 0; v /= 10 {
                pal = pal*10 + v%10
            }
            specialNumbers = append(specialNumbers, pal)
            return false
        })
    }
    slices.Sort(specialNumbers)
    specialNumbers = slices.Compact(specialNumbers)
}

func specialPalindrome(n int64) int64 {
    i := sort.SearchInts(specialNumbers, int(n+1))
    return int64(specialNumbers[i])
}

func permutations(n, r int, do func(ids []int) (Break bool)) {
    ids := make([]int, n)
    for i := range ids {
        ids[i] = i
    }
    if do(ids[:r]) {
        return
    }
    cycles := make([]int, r)
    for i := range cycles {
        cycles[i] = n - i
    }
    for {
        i := r - 1
        for ; i >= 0; i-- {
            cycles[i]--
            if cycles[i] == 0 {
                tmp := ids[i]
                copy(ids[i:], ids[i+1:])
                ids[n-1] = tmp
                cycles[i] = n - i
            } else {
                j := cycles[i]
                ids[i], ids[n-j] = ids[n-j], ids[i]
                if do(ids[:r]) {
                    return
                }
                break
            }
        }
        if i == -1 {
            return
        }
    }
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 22
    // Explanation:
    // 22 is the smallest special number greater than 2, as it is a palindrome and the digit 2 appears exactly 2 times.
    fmt.Println(specialPalindrome(2)) // 22
    // Example 2:
    // Input: n = 33
    // Output: 212
    // Explanation:
    // 212 is the smallest special number greater than 33, as it is a palindrome and the digits 1 and 2 appear exactly 1 and 2 times respectively.
    fmt.Println(specialPalindrome(33)) // 212

    fmt.Println(specialPalindrome(0)) // 1
    fmt.Println(specialPalindrome(1)) // 22
    fmt.Println(specialPalindrome(1024)) // 4444
    fmt.Println(specialPalindrome(1_000_000_000_000_000)) // 2666888888886662
}