package main

// 526. Beautiful Arrangement
// Suppose you have n integers labeled 1 through n. 
// A permutation of those n integers perm (1-indexed) is considered a beautiful arrangement 
// if for every i (1 <= i <= n), either of the following is true:
//     perm[i] is divisible by i.
//     i is divisible by perm[i].

// Given an integer n, return the number of the beautiful arrangements that you can construct.

// Example 1:
// Input: n = 2
// Output: 2
// Explanation: 
// The first beautiful arrangement is [1,2]:
//     - perm[1] = 1 is divisible by i = 1
//     - perm[2] = 2 is divisible by i = 2
// The second beautiful arrangement is [2,1]:
//     - perm[1] = 2 is divisible by i = 1
//     - i = 2 is divisible by perm[2] = 1

// Example 2:
// Input: n = 1
// Output: 1

// Constraints:
//     1 <= n <= 15

import "fmt"

// dfs 回溯
func countArrangement(n int) int {
    if n == 0 {
        return 0
    }
    nums, used, p, res := make([]int, n), make([]bool, n), []int{}, [][]int{}
    for i := range nums {
        nums[i] = i + 1
    }
    checkDivisible := func (num, d int) bool {
        tmp := num / d
        if int(tmp)*int(d) == num {
            return true
        }
        return false
    }
    var dfs func(nums []int, index int, p []int, res *[][]int, used *[]bool)
    dfs = func(nums []int, index int, p []int, res *[][]int, used *[]bool) {
        if index == len(nums) {
            temp := make([]int, len(p))
            copy(temp, p)
            *res = append(*res, temp)
            return
        }
        for i := 0; i < len(nums); i++ {
            if !(*used)[i] {
                if !(checkDivisible(nums[i], len(p)+1) || checkDivisible(len(p)+1, nums[i])) { // 关键的剪枝条件
                    continue
                }
                (*used)[i] = true
                p = append(p, nums[i])
                dfs(nums, index+1, p, res, used)
                p = p[:len(p)-1]
                (*used)[i] = false
            }
        }
        return
    }
    dfs(nums, 0, p, &res, &used)
    return len(res)
}

func countArrangement1(n int) int {
    m := 1 << n
    memo := make([]int, m) // 记忆化搜索
    for i := range memo {
        memo[i] = -1
    }
    var dfs func(int, int) int
    dfs = func(mask, i int) int {
        if i == 0 {
            return 1
        }
        v := &memo[mask]
        if *v >= 0 {
            return *v
        }
        res := 0
        for j := 1; j <= n; j++ {
            if mask>>(j-1)&1 > 0 && (i%j == 0 || j%i == 0) {
                res += dfs(1 << (j-1) ^ mask, i-1)
            }
        }
        *v = res
        return res
    }
    return dfs(m - 1, n)
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 2
    // Explanation: 
    // The first beautiful arrangement is [1,2]:
    //     - perm[1] = 1 is divisible by i = 1
    //     - perm[2] = 2 is divisible by i = 2
    // The second beautiful arrangement is [2,1]:
    //     - perm[1] = 2 is divisible by i = 1
    //     - i = 2 is divisible by perm[2] = 1
    fmt.Println(countArrangement(2)) // 2
    // Example 2:
    // Input: n = 1
    // Output: 1
    fmt.Println(countArrangement(1)) // 1

    fmt.Println(countArrangement(3)) // 3
    fmt.Println(countArrangement(4)) // 8
    fmt.Println(countArrangement(5)) // 10
    fmt.Println(countArrangement(6)) // 36
    fmt.Println(countArrangement(7)) // 41
    fmt.Println(countArrangement(8)) // 132
    fmt.Println(countArrangement(9)) // 250
    fmt.Println(countArrangement(10)) // 700

    fmt.Println(countArrangement1(2)) // 2
    fmt.Println(countArrangement1(1)) // 1
    fmt.Println(countArrangement1(3)) // 3
    fmt.Println(countArrangement1(4)) // 8
    fmt.Println(countArrangement1(5)) // 10
    fmt.Println(countArrangement1(6)) // 36
    fmt.Println(countArrangement1(7)) // 41
    fmt.Println(countArrangement1(8)) // 132
    fmt.Println(countArrangement1(9)) // 250
    fmt.Println(countArrangement1(10)) // 700
}