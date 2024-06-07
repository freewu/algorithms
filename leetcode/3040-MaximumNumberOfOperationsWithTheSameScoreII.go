package main

// 3040. Maximum Number of Operations With the Same Score II
// Given an array of integers called nums, you can perform any of the following operation while nums contains at least 2 elements:
//     Choose the first two elements of nums and delete them.
//     Choose the last two elements of nums and delete them.
//     Choose the first and the last elements of nums and delete them.

// The score of the operation is the sum of the deleted elements.
// Your task is to find the maximum number of operations that can be performed, such that all operations have the same score.
// Return the maximum number of operations possible that satisfy the condition mentioned above.

// Example 1:
// Input: nums = [3,2,1,2,3,4]
// Output: 3
// Explanation: We perform the following operations:
// - Delete the first two elements, with score 3 + 2 = 5, nums = [1,2,3,4].
// - Delete the first and the last elements, with score 1 + 4 = 5, nums = [2,3].
// - Delete the first and the last elements, with score 2 + 3 = 5, nums = [].
// We are unable to perform any more operations as nums is empty.

// Example 2:
// Input: nums = [3,2,6,1,4]
// Output: 2
// Explanation: We perform the following operations:
// - Delete the first two elements, with score 3 + 2 = 5, nums = [6,1,4].
// - Delete the last two elements, with score 1 + 4 = 5, nums = [6].
// It can be proven that we can perform at most 2 operations.

// Constraints:
//     2 <= nums.length <= 2000
//     1 <= nums[i] <= 1000

import "fmt"
import "strconv"

func maxOperations(nums []int) int {
    n, memo := len(nums), make(map[string]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int,int,int) int
    dfs = func(left,right,val int) int{
        if left >= right {
            return 0
        }
        pos := strconv.Itoa(left) + "-" + strconv.Itoa(right) + "-" + strconv.Itoa(val)
        if val,ok := memo[pos]; ok {
            return val
        }
        res := 0
        if nums[left]+nums[right] == val {
            res = max(res, dfs(left + 1, right - 1, val) + 1)
        }
        if nums[left]+nums[left+1] == val {
            res = max(res, dfs(left + 2, right, val) + 1)
        }
        if nums[right]+nums[right-1] == val {
            res = max(res, dfs(left, right - 2, val) + 1)
        }
        memo[pos] = res
        return res
    }
    res := 0
    res = max(res,dfs(2, n - 1, nums[0] + nums[1]))
    res = max(res,dfs(1, n - 2, nums[0] + nums[n-1]))
    res = max(res,dfs(0, n - 3, nums[n-2] + nums[n-1]))
    return res + 1
}

func maxOperations1(nums []int) int {
    cache := make([][]int,len(nums))
    for i := range cache {
        cache[i] = make([]int,len(nums))
        for j := range cache[i] {
            cache[i][j] = -1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(nums []int,cache [][]int,l,r, val,cnt int) int
    dfs = func(nums []int,cache [][]int,l,r, val,cnt int) int {
        if l >= r {
            return 0
        }
        if cache[l][r]>=0 {
            return cache[l][r]
        }
        res := 0
        if nums[l]+nums[l+1] == val {
            res = dfs(nums,cache,l+2,r,val,cnt+1)+1
        }
        if nums[r]+nums[r-1] == val {
            res = max(res,dfs(nums,cache,l,r-2,val,cnt+1)+1)
        }
        if nums[l]+nums[r] == val {
            res = max(res,dfs(nums,cache,l+1,r-1,val,cnt+1)+1)
        }
        cache[l][r] = res
        return res
    }
    res := dfs(nums,cache,2,len(nums)-1,nums[0]+nums[1],0)
    res = max(res, dfs(nums,cache,1,len(nums)-2,nums[0]+nums[len(nums)-1],0))
    res = max(res, dfs(nums,cache,0,len(nums)-3,nums[len(nums)-2]+nums[len(nums)-1],0))
    return res + 1
}

func main() {
    // Example 1:
    // Input: nums = [3,2,1,2,3,4]
    // Output: 3
    // Explanation: We perform the following operations:
    // - Delete the first two elements, with score 3 + 2 = 5, nums = [1,2,3,4].
    // - Delete the first and the last elements, with score 1 + 4 = 5, nums = [2,3].
    // - Delete the first and the last elements, with score 2 + 3 = 5, nums = [].
    // We are unable to perform any more operations as nums is empty.
    fmt.Println(maxOperations([]int{3,2,1,2,3,4})) // 3
    // Example 2:
    // Input: nums = [3,2,6,1,4]
    // Output: 2
    // Explanation: We perform the following operations:
    // - Delete the first two elements, with score 3 + 2 = 5, nums = [6,1,4].
    // - Delete the last two elements, with score 1 + 4 = 5, nums = [6].
    // It can be proven that we can perform at most 2 operations.
    fmt.Println(maxOperations([]int{3,2,6,1,4})) // 2

    fmt.Println(maxOperations1([]int{3,2,1,2,3,4})) // 3
    fmt.Println(maxOperations1([]int{3,2,6,1,4})) // 2
}