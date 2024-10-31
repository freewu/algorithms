package main

// 1643. Kth Smallest Instructions
// Bob is standing at cell (0, 0), and he wants to reach destination: (row, column). 
// He can only travel right and down. 
// You are going to help Bob by providing instructions for him to reach destination.

// The instructions are represented as a string, where each character is either:
//     'H', meaning move horizontally (go right), or
//     'V', meaning move vertically (go down).

// Multiple instructions will lead Bob to destination. 
// For example, if destination is (2, 3), both "HHHVV" and "HVHVH" are valid instructions.

// However, Bob is very picky. Bob has a lucky number k, 
// and he wants the kth lexicographically smallest instructions that will lead him to destination. k is 1-indexed.

// Given an integer array destination and an integer k, 
// return the kth lexicographically smallest instructions that will take Bob to destination.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/12/ex1.png" />
// Input: destination = [2,3], k = 1
// Output: "HHHVV"
// Explanation: All the instructions that reach (2, 3) in lexicographic order are as follows:
// ["HHHVV", "HHVHV", "HHVVH", "HVHHV", "HVHVH", "HVVHH", "VHHHV", "VHHVH", "VHVHH", "VVHHH"].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/12/ex2.png" />
// Input: destination = [2,3], k = 2
// Output: "HHVHV"

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/12/ex3.png" />
// Input: destination = [2,3], k = 3
// Output: "HHVVH"

// Constraints:
//     destination.length == 2
//     1 <= row, column <= 15
//     1 <= k <= nCr(row + column, row), where nCr(a, b) denotes a choose b​​​​​.

import "fmt"
import "strings"

func kthSmallestPath(des []int, k int) string {
    nCr := func() [][]int {
        res := make([][]int, 31)
        res[0] = []int{1}
        res[1] = []int{1,1}
        for i:=2; i<=30; i++ {
            res[i] = []int{1}
            for j:=1; j<i; j++ {
                res[i] = append(res[i], res[i-1][j-1] + res[i-1][j])
            }
            res[i] = append(res[i], 1)
        }
        return res
    }
    comb := nCr()
    sum, l := (des[0] + des[1]), (des[0] + des[1])
    res, v := "", des[0]
    fill := func(s string, sum, v int) string {
        for i := 0; i < v; i++ { s += "V"  }
        for i := 0; i < sum - v; i++ { s += "H" }
        return s
    }
    for len(res) < l {
        if k == comb[sum][v] {
            return fill(res, sum, v)
        } else if k > comb[sum - 1][v] {
            res += "V"
            k -= comb[sum - 1][v]
            v--
            sum--
        } else {
            res += "H"
            sum--
        }
    }
    return res
}

func kthSmallestPath1(destination []int, k int) string {
    /**
    So the smallest path is straightforward: HHHHHHVVV
    Let's say there are total s steps, and v are verticals.

    IF the k-th path is like this: HHH...V....V...V...
        - For the first V, it should have iterated over all possible arrangements for (v-1)*V + H
    How many Hs are there before the first V? Let's use binary search to decide on this.
    Or we just simply search over 0 to ...
    */
    h, v := destination[1], destination[0]
    total := h + v
    nCr := make([][]int64, total + 1)
    for i := 0; i <= total; i++ {
        nCr[i] = make([]int64, total + 1)
        nCr[i][0] = 1
        for j := 1; j <= i; j++ {
            nCr[i][j] = nCr[i-1][j-1] * int64(i) / int64(j)
        }
    }
    res := ""
    for h > 0 && v > 0 && k > 1 {
        total = h + v 
        for i := 0; i <= h; i++ {
            if nCr[total - i][v] < int64(k) {
                res += strings.Repeat("H", i - 1) + "V"
                k -= int(nCr[total-i][v])
                h -= i - 1
                v--
                break
            }
        }
    }
    res += strings.Repeat("H", h)
    res += strings.Repeat("V", v)
    return res 
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/12/ex1.png" />
    // Input: destination = [2,3], k = 1
    // Output: "HHHVV"
    // Explanation: All the instructions that reach (2, 3) in lexicographic order are as follows:
    // ["HHHVV", "HHVHV", "HHVVH", "HVHHV", "HVHVH", "HVVHH", "VHHHV", "VHHVH", "VHVHH", "VVHHH"].
    fmt.Println(kthSmallestPath([]int{2,3}, 1)) // "HHHVV"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/12/ex2.png" />
    // Input: destination = [2,3], k = 2
    // Output: "HHVHV"
    fmt.Println(kthSmallestPath([]int{2,3}, 2)) // "HHVHV"
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/10/12/ex3.png" />
    // Input: destination = [2,3], k = 3
    // Output: "HHVVH"
    fmt.Println(kthSmallestPath([]int{2,3}, 3)) // "HHVVH"

    fmt.Println(kthSmallestPath1([]int{2,3}, 1)) // "HHHVV"
    fmt.Println(kthSmallestPath1([]int{2,3}, 2)) // "HHVHV"
    fmt.Println(kthSmallestPath1([]int{2,3}, 3)) // "HHVVH"
}