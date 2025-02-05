package main

// 90. Subsets II
// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:
// Input: nums = [1,2,2]
// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

// Example 2:
// Input: nums = [0]
// Output: [[],[0]]

// Constraints:
//     1 <= nums.length <= 10
//     -10 <= nums[i] <= 10

import "fmt"
import "sort"

func subsetsWithDup(nums []int) [][]int {
    c, res := []int{}, [][]int{}
    sort.Ints(nums) // 这里是去重的关键逻辑
    var dfs func(nums []int, k, start int, c []int,)
    dfs = func(nums []int, k, start int, c []int) {
        if len(c) == k {
            b := make([]int, len(c))
            copy(b, c)
            res = append(res, b)
            return
        }
        // i will at most be n - (k - c.size()) + 1
        for i := start; i < len(nums)-(k-len(c))+1; i++ {
            // fmt.Printf("i = %v start = %v c = %v\n", i, start, c)
            if i > start && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
                continue
            }
            c = append(c, nums[i])
            dfs(nums, k, i+1, c)
            c = c[:len(c)-1]
        }
        return
    }
    for k := 0; k <= len(nums); k++ {
        dfs(nums, k, 0, c)
    }
    return res
}

// best solution
func subsetsWithDup1(nums []int) [][]int {
    res, count := [][]int{{}}, make(map[int]int)
    for _, v := range nums {
        count[v]++
    }
    for k, v := range count {
        res2 := append([][]int{}, res...)
        for _, r := range res {
            for f := 1; f <= v; f++ {
                r2 := append([]int{}, r...)
                for i := 0; i < f; i++ {
                    r2 = append(r2, k)
                }
                res2 = append(res2, r2)
            }
        }
        res = res2
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2]
    // Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
    fmt.Printf("subsetsWithDup([]int{1,2,3}) = %v\n",subsetsWithDup([]int{1,2,3})) //  [[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]
    // Example 2:
    // Input: nums = [0]
    // Output: [[],[0]]
    fmt.Printf("subsetsWithDup([]int{0}) = %v\n",subsetsWithDup([]int{0})) // [[] [0]]
    //fmt.Printf("subsetsWithDup([]int{1,2,3,4,5,6,7,8,9}) = %v\n",subsetsWithDup([]int{1,2,3,4,5,6,7,8,9}))
 
    fmt.Printf("subsetsWithDup1([]int{1,2,3}) = %v\n",subsetsWithDup1([]int{1,2,3})) // [[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]
    fmt.Printf("subsetsWithDup1([]int{0}) = %v\n",subsetsWithDup1([]int{0})) // [[] [0]]
    //fmt.Printf("subsetsWithDup1([]int{1,2,3,4,5,6,7,8,9}) = %v\n",subsetsWithDup1([]int{1,2,3,4,5,6,7,8,9}))
}
