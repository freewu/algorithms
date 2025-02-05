package main

// 47. Permutations II
// Given a collection of numbers, nums, that might contain duplicates, 
// return all possible unique permutations in any order.

// Example 1:
// Input: nums = [1,1,2]
// Output: [[1,1,2], [1,2,1], [2,1,1]]

// Example 2:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 
// Constraints:
//     1 <= nums.length <= 8
//     -10 <= nums[i] <= 10

import "fmt"
import "sort"

func permuteUnique(nums []int) [][]int {
    used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
    if len(nums) == 0 {
        return res
    }
    sort.Ints(nums)
    var generatePermutation func(nums []int, index int, p []int)
    generatePermutation = func(nums []int, index int, p []int) {
        if index == len(nums) {
            temp := make([]int, len(p))
            copy(temp, p)
            res = append(res, temp)
            return
        }
        for i := 0; i < len(nums); i++ {
            if !used[i] {
                if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] { continue } // 去重判断
                used[i] = true
                p = append(p, nums[i])
                generatePermutation(nums, index + 1, p)
                p = p[:len(p)-1]
                used[i] = false
            }
        }
    }
    generatePermutation(nums, 0, p)
    return res
}

func permuteUnique1(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    var solve func(comb []int, used []bool)
    solve = func(comb []int, used []bool) {
        if len(comb) == len(nums) {
            a := append([]int{}, comb...)
            res = append(res, a)
            return
        }
        for i, v := range nums {
            if used[i] { continue  }
            if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] { continue }
            used[i] = true
            solve(append(comb, v), used)
            used[i] = false
        }
    }
    solve([]int{},make([]bool, len(nums)))
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2]
    // Output: [[1,1,2], [1,2,1], [2,1,1]]
    fmt.Printf("permuteUnique([]int{1,2,3}) = %v\n", permuteUnique([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    fmt.Printf("permuteUnique([]int{1,1,2}) = %v\n", permuteUnique([]int{1,1,2})) // [[1,1,2], [1,2,1], [2,1,1]]
    
    fmt.Printf("permuteUnique1([]int{1,2,3}) = %v\n", permuteUnique1([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    fmt.Printf("permuteUnique1([]int{1,1,2}) = %v\n", permuteUnique1([]int{1,1,2})) //[[1,1,2], [1,2,1], [2,1,1]]
}