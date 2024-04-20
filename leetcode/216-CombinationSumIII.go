package main

// 216. Combination Sum III
// Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
//     Only numbers 1 through 9 are used.
//     Each number is used at most once.

// Return a list of all possible valid combinations.
// The list must not contain the same combination twice, and the combinations may be returned in any order.

// Example 1:
// Input: k = 3, n = 7
// Output: [[1,2,4]]
// Explanation:
// 1 + 2 + 4 = 7
// There are no other valid combinations.

// Example 2:
// Input: k = 3, n = 9
// Output: [[1,2,6],[1,3,5],[2,3,4]]
// Explanation:
// 1 + 2 + 6 = 9
// 1 + 3 + 5 = 9
// 2 + 3 + 4 = 9
// There are no other valid combinations.

// Example 3:
// Input: k = 4, n = 1
// Output: []
// Explanation: There are no valid combinations.
// Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.

// Constraints:
//     2 <= k <= 9
//     1 <= n <= 60

// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
// 数组是固定死的 [1,2,3,4,5,6,7,8,9]，并且数字不能重复使用

import "fmt"

func combinationSum3(k int, n int) [][]int {
    c, res := []int{}, [][]int{}
    if k == 0 {
        return res
    }
    var dfs func (k, target, index int, c []int, res *[][]int)
    dfs = func (k, target, index int, c []int, res *[][]int) {
        if target == 0 {
            if len(c) == k {
                b := make([]int, len(c))
                copy(b, c)
                *res = append(*res, b)
            }
            return
        }
        for i := index; i < 10; i++ {
            if target >= i {
                c = append(c, i)
                dfs(k, target-i, i+1, c, res)
                // fmt.Printf("i = %v,res = %v\n",i,res)
                c = c[ :len(c)-1]
            }
        }
    }
    dfs(k, n, 1, c, &res)
    return res
}

func combinationSum31(k int, n int) [][]int  {
    t, res := []int{}, [][]int{}
    var dfs func(cur, rest int)
    dfs = func(cur, rest int) {
        if len(t) == k && rest == 0 { // 找到一个答案
            res = append(res, append([]int(nil), t...))
            return
        }
        if len(t) + 10 - cur < k || rest < 0 { // 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
            return
        }
        dfs(cur + 1, rest) // 跳过当前数字
        t = append(t, cur) // 选当前数字
        dfs(cur+1, rest-cur)
        t = t[:len(t)-1]
    }
    dfs(1, n)
    return res
}

func main() {
    // 1 + 2 + 4 = 7
    // There are no other valid combinations.
    fmt.Printf("combinationSum3(3,7) = %v\n",combinationSum3(3,7)) // [[1,2,4]]
    // 1 + 2 + 6 = 9
    // 1 + 3 + 5 = 9
    // 2 + 3 + 4 = 9
    // There are no other valid combinations.
    fmt.Printf("combinationSum3(3,9) = %v\n",combinationSum3(3,9)) // [[1,2,6],[1,3,5],[2,3,4]]
    // Explanation: There are no valid combinations.
    // Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
    fmt.Printf("combinationSum3(4,1) = %v\n",combinationSum3(4,1)) // []

    fmt.Printf("combinationSum31(3,7) = %v\n",combinationSum31(3,7)) // [[1,2,4]]
    fmt.Printf("combinationSum31(3,9) = %v\n",combinationSum31(3,9)) // [[1,2,6],[1,3,5],[2,3,4]]
    fmt.Printf("combinationSum31(4,1) = %v\n",combinationSum31(4,1)) // []
}
