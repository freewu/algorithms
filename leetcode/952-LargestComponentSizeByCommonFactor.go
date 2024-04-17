package main

// 952. Largest Component Size by Common Factor
// You are given an integer array of unique positive integers nums. Consider the following graph:
//     There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
//     There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share a common factor greater than 1.

// Return the size of the largest connected component in the graph.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/01/ex1.png" />
// Input: nums = [4,6,15,35]
// Output: 4

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/01/ex2.png" />
// Input: nums = [20,50,9,63]
// Output: 2

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/01/ex3.png" />
// Input: nums = [2,3,6,7,4,12,21,39]
// Output: 8
 
// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i] <= 10^5
//     All the values of nums are unique.

import "fmt"
import "math"

type unionFind map[int]int

func (u unionFind) inc(f int) (i int) {
    i = u.top(f)
    u[i]--
    return i
}

func (u unionFind) link(t, f int) int {
    i := u.top(f)
    if i == t { return t }
    u[t] += u[i]
    u[i] = t
    return t
}

func (u unionFind) max() int {
    x := 0
    for _, n := range u {
        if n < x { x = n }
    }
    return -x
}

func (u unionFind) top(i int) int {
    for u[i] > 0 { i = u[i] }
    return i
}

func largestComponentSize(nums []int) int {
    uf := unionFind{}
    factorize := func (num int) (res []int) {
        // modified https://leetcode.com/problems/largest-component-size-by-common-factor/discuss/1592831/golang-union-find-standard-template
        for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
            if num%i == 0 { res = append(res, i, num/i) }
        }
        if len(res) == 0 { res = []int{num} }
        return res
    }
    for _, n := range nums {
        fs := factorize(n)
        top := uf.inc(fs[0])
        for _, f := range fs[1:] {
            top = uf.link(top, f)
        }
    }
    return uf.max()
}

func largestComponentSize1(nums []int) int {
    //gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    n := len(nums)
    if n == 0 {
        return 0
    }
    uf, rank := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        uf[i] = i
        rank[i] = 1
    }
    find := func(x int) int {
        for uf[x] != x {
            uf[x] = uf[uf[x]]
            x = uf[x]
        }
        return uf[x]
    }
    fa, res := make(map[int]int), 1
    for i := 0; i < n; i++ {
        x := nums[i]
        if x == 1 {
            continue
        }
        t := x
        for k := 2; k * k <= t; k++ {
            if t % k == 0 {
                for t % k == 0 {
                    t /= k
                }
                if _, ok := fa[k]; !ok {
                    fa[k] = i
                } else {
                    j := fa[k]
                    p, q := find(i), find(j)
                    if q != p {
                        uf[p] = q
                        rank[q] += rank[p]
                        if rank[q] > res {
                            res = rank[q]
                        }
                    }
                }
            }
        }
        if t > 1 {
            if _, ok := fa[t]; !ok {
                fa[t] = i
            } else {
                j := fa[t]
                p, q := find(i), find(j)
                if q != p {
                    uf[p] = q
                    rank[q] += rank[p]
                    if rank[q] > res {
                        res = rank[q]
                    }
                }
            }
        }
    }
    return res
}


func main() {
    fmt.Println(largestComponentSize([]int{4,6,15,35})) // 4
    fmt.Println(largestComponentSize([]int{20,50,9,63})) // 2
    fmt.Println(largestComponentSize([]int{2,3,6,7,4,12,21,39})) // 8

    fmt.Println(largestComponentSize1([]int{4,6,15,35})) // 4
    fmt.Println(largestComponentSize1([]int{20,50,9,63})) // 2
    fmt.Println(largestComponentSize1([]int{2,3,6,7,4,12,21,39})) // 8
}