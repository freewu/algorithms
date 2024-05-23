package main

// 2597. The Number of Beautiful Subsets
// You are given an array nums of positive integers and a positive integer k.
// A subset of nums is beautiful if it does not contain two integers with an absolute difference equal to k.
// Return the number of non-empty beautiful subsets of the array nums.
// A subset of nums is an array that can be obtained by deleting some (possibly none) elements from nums. 
// Two subsets are different if and only if the chosen indices to delete are different.

// Example 1:
// Input: nums = [2,4,6], k = 2
// Output: 4
// Explanation: The beautiful subsets of the array nums are: [2], [4], [6], [2, 6].
// It can be proved that there are only 4 beautiful subsets in the array [2,4,6].

// Example 2:
// Input: nums = [1], k = 1
// Output: 1
// Explanation: The beautiful subset of the array nums is [1].
// It can be proved that there is only 1 beautiful subset in the array [1].

// Constraints:
//     1 <= nums.length <= 20
//     1 <= nums[i], k <= 1000

import "fmt"
import "sort"

// 无需主动回溯
func beautifulSubsets(nums []int, k int) int {
    sort.Ints(nums)
    res, arr := 0, []int{}
    var dfs func(id int, arr []int)
    dfs = func(id int, arr []int) {
        if id == len(nums) {
            if len(arr) > 0 {
                res += 1
            }
            return
        }
        dfs(id + 1, arr) // 不选 id 这个数字
        // 选 id 这个数字
        // 但是 nums[id] 不能和 arr 中任何数字的差 为 k, 也就是 arr 中不存在 nums[id] - k 的数字
        pos := sort.SearchInts(arr, nums[id] - k)
        if pos == len(arr) || arr[pos] != nums[id] - k {
            // 因为 append 之后，创建了一个新的切片，当前 arr 并没有被改变，所以不需要回溯
            dfs(id + 1, append(arr, nums[id]))
        }
    }
    dfs(0, arr)
    return res
}

func beautifulSubsets1(nums []int, k int) int {
    res, m := 0, make(map[int]int)
    var dfs func(id int, m map[int]int)
    dfs = func(id int, m map[int]int) {
        if id == len(nums) {
            if len(m) > 0 {
                res += 1
            }
            return
        }
        dfs(id + 1, m) // 不选 id 这个数字
        // 选 id 这个数字,
        // 但是 nums[id] 不能和 m 中任何数字的差 为 k,
        v := nums[id]
        cnt0, cnt1 := m[v - k], m[v + k]
        if cnt0 == 0 && cnt1 == 0 {
            m[v] += 1
            dfs(id + 1, m)
            m[v] -= 1 // 回溯
        }
    }
    dfs(0, m)
    return res
}

func beautifulSubsets2(nums []int, k int) int {
    res, groups := 1, map[int]map[int]int{}
    for _, v := range nums {
        if groups[v % k] == nil{
            groups[v % k] = map[int]int{}
        }
        groups[v % k][v] += 1
    }
    type pair struct {
        x, c int
    }
    for _, group := range groups {
        g := []pair{}
        for k, v := range group {
            g = append(g, pair{k, v})
        }
        sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x })
        n := len(g)
        dp := make([][]int, n)
        for i := 0; i < n; i++ {
            dp[i] = make([]int, 2)
        }
        dp[0][0] = 1
        dp[0][1] = (1 << g[0].c) - 1
        for i := 1; i < n; i++ {
            if g[i].x - g[i - 1].x == k {
                dp[i][0] = dp[i - 1][0] + dp[i - 1][1]
                dp[i][1] = dp[i - 1][0] * ((1 << g[i].c) - 1)
            } else {
                dp[i][0] = dp[i - 1][0] + dp[i - 1][1]
                dp[i][1] = (dp[i - 1][0] + dp[i - 1][1]) * ((1 << g[i].c) - 1)
            }
        }
        res = res * (dp[n - 1][0] + dp[n - 1][1])
    }
    return res - 1
}

func main() {
    // Example 1:
    // Input: nums = [2,4,6], k = 2
    // Output: 4
    // Explanation: The beautiful subsets of the array nums are: [2], [4], [6], [2, 6].
    // It can be proved that there are only 4 beautiful subsets in the array [2,4,6].
    fmt.Println(beautifulSubsets([]int{2,4,6}, 2)) // 4
    // Example 2:
    // Input: nums = [1], k = 1
    // Output: 1
    // Explanation: The beautiful subset of the array nums is [1].
    // It can be proved that there is only 1 beautiful subset in the array [1].
    fmt.Println(beautifulSubsets([]int{1}, 1)) // 1

    fmt.Println(beautifulSubsets1([]int{2,4,6}, 2)) // 4
    fmt.Println(beautifulSubsets1([]int{1}, 1)) // 1

    fmt.Println(beautifulSubsets2([]int{2,4,6}, 2)) // 4
    fmt.Println(beautifulSubsets2([]int{1}, 1)) // 1
}
