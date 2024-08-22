package main

// 3145. Find Products of Elements of Big Array
// The powerful array of a non-negative integer x is defined as the shortest sorted array of powers of two that sum up to x. 
// The table below illustrates examples of how the powerful array is determined. 
// It can be proven that the powerful array of x is unique.

// num	Binary Representation	powerful array
// 1	00001	[1]
// 8	01000	[8]
// 10	01010	[2, 8]
// 13	01101	[1, 4, 8]
// 23	10111	[1, 2, 4, 16]

// The array big_nums is created by concatenating the powerful arrays for every positive integer i in ascending order: 1, 2, 3, and so on. 
// Thus, big_nums begins as [1, 2, 1, 2, 4, 1, 4, 2, 4, 1, 2, 4, 8, ...].

// You are given a 2D integer matrix queries, 
// where for queries[i] = [fromi, toi, modi] you should calculate (big_nums[fromi] * big_nums[fromi + 1] * ... * big_nums[toi]) % modi.

// Return an integer array answer such that answer[i] is the answer to the ith query.

// Example 1:
// Input: queries = [[1,3,7]]
// Output: [4]
// Explanation:
// There is one query.
// big_nums[1..3] = [2,1,2]. The product of them is 4. The result is 4 % 7 = 4.

// Example 2:
// Input: queries = [[2,5,3],[7,7,4]]
// Output: [2,2]
// Explanation:
// There are two queries.
// First query: big_nums[2..5] = [1,2,4,1]. The product of them is 8. The result is 8 % 3 = 2.
// Second query: big_nums[7] = 2. The result is 2 % 4 = 2.

// Constraints:
//     1 <= queries.length <= 500
//     queries[i].length == 3
//     0 <= queries[i][0] <= queries[i][1] <= 10^15
//     1 <= queries[i][2] <= 10^5

import "fmt"
import "math"
import "math/bits"

func findProductsOfElements(queries [][]int64) []int {
    res := make([]int, 0)
    // 计算 <= x 所有数的数位1的和
    countOne := func (x int64) int64 {
        res, sum := int64(0), 0
        for i := 60; i >= 0; i-- {
            if (1 << i) & x != 0 {
                res += int64(sum) * (1 << i)
                sum++
                if i > 0 {
                    res += int64(i) * (1 << (i - 1))
                }
            }
        }
        res += int64(sum)
        return res
    }
    // 计算 <= x 所有数的数位对幂的贡献之和
    countPow := func (x int64) int64 {
        res, sum := int64(0), 0
        for i := 60; i >= 0; i-- {
            if (1 << i) & x != 0 {
                res += int64(sum) * (1 << i)
                sum += i
                
                if i > 0 {
                    res += int64(i) * (int64(i) - 1) / 2 * (1 << (i - 1))
                }
            }
        }
        res += int64(sum)
        return res
    }

    powMod := func (x int64, y int64, mod int) int {
        res := 1
        for y > 0 {
            if y & 1 != 0 {
                res = res * int(x) % mod
            }
            x = x * x % int64(mod)
            y >>= 1
        }
        return res
    }
    midCheck := func (x int64) int64 {
        l, r := int64(1), int64(1e15)
        for l < r {
            mid := (l + r) >> 1
            if countOne(mid) >= x {
                r = mid
            } else {
                l = mid + 1
            }
        }
        return r
    }
    for _, query := range queries {
        // 偏移让数组下标从1开始
        query[0]++
        query[1]++
        l, r := midCheck(query[0]), midCheck(query[1])
        mod := int(query[2])
        t, pre := int64(1), countOne(l - 1)
        for j := 0; j < 60; j++ {
            if (1 << j) & l != 0 {
                pre++
                if pre >= query[0] && pre <= query[1] {
                    t = t * (1 << j) % int64(mod)
                }
            }
        }
        if r > l {
            bac := countOne(r - 1)
            for j := 0; j < 60; j++ {
                if (1 << j) & r != 0 {
                    bac++
                    if bac >= query[0] && bac <= query[1] {
                        t = t * (1 << j) % int64(mod)
                    }
                }
            }
        }
        if r - l > 1 {
            xs := countPow(r - 1) - countPow(l)
            t = t * int64(powMod(2, xs, mod)) % int64(mod)
        }
        res = append(res, int(t))
    }
    return res
}

func findProductsOfElements1(queries [][]int64) []int {
    pow := func (a, b, mod int) int {
        a %= mod
        res := 1
        for b > 0 {
            if b&1 > 0 {
                res = res * a % mod
            }
            b >>= 1
            a = a * a % mod
        }
        return res % mod
    }
    calc := func(num int) int {
        cnt, res, powCnt, v := 0, 0, 0, 0
        for i := 45; i >= 0 && num > 0; i-- {
            if cnt*(1<<i)+i*pow(2, i-1, math.MaxInt) <= num {
                v |= 1 << i
                res += powCnt*1<<i + i*(i-1)/2 * pow(2, i-1, math.MaxInt)
                powCnt += i
                num -= cnt*(1<<i) + i*pow(2, i-1, math.MaxInt)
                cnt++
            }
        }
        for ; num > 0; num-- {
            lb := bits.TrailingZeros(uint(v))
            res += lb
            v ^= 1 << lb
        }
        return res
    }
    res := make([]int, 0, len(queries))
    for _, query := range queries {
        from, to, mod := int(query[0]), int(query[1]), int(query[2])
        to++
        p := calc(to) - calc(from)
        res = append(res, pow(2, p, mod))
    }
    return res
}

func main() {
    // Example 1:
    // Input: queries = [[1,3,7]]
    // Output: [4]
    // Explanation:
    // There is one query.
    // big_nums[1..3] = [2,1,2]. The product of them is 4. The result is 4 % 7 = 4.
    fmt.Println(findProductsOfElements([][]int64{{1,3,7}})) // [4]
    // Example 2:
    // Input: queries = [[2,5,3],[7,7,4]]
    // Output: [2,2]
    // Explanation:
    // There are two queries.
    // First query: big_nums[2..5] = [1,2,4,1]. The product of them is 8. The result is 8 % 3 = 2.
    // Second query: big_nums[7] = 2. The result is 2 % 4 = 2.
    fmt.Println(findProductsOfElements([][]int64{{2,5,3}, {7,7,4}})) // [2,2]

    fmt.Println(findProductsOfElements1([][]int64{{1,3,7}})) // [4]
    fmt.Println(findProductsOfElements1([][]int64{{2,5,3}, {7,7,4}})) // [2,2]
}