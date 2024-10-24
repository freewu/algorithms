package main

// 3181. Maximum Total Reward Using Operations II
// You are given an integer array rewardValues of length n, representing the values of rewards.

// Initially, your total reward x is 0, and all indices are unmarked. 
// You are allowed to perform the following operation any number of times:
//     1. Choose an unmarked index i from the range [0, n - 1].
//     2. If rewardValues[i] is greater than your current total reward x, 
//        then add rewardValues[i] to x (i.e., x = x + rewardValues[i]), and mark the index i.

// Return an integer denoting the maximum total reward you can collect by performing the operations optimally.

// Example 1:
// Input: rewardValues = [1,1,3,3]
// Output: 4
// Explanation:
// During the operations, we can choose to mark the indices 0 and 2 in order, and the total reward will be 4, which is the maximum.

// Example 2:
// Input: rewardValues = [1,6,4,3,2]
// Output: 11
// Explanation:
// Mark the indices 0, 2, and 1 in order. The total reward will then be 11, which is the maximum.

// Constraints:
//     1 <= rewardValues.length <= 5 * 10^4
//     1 <= rewardValues[i] <= 5 * 10^4

import "fmt"
import "math/big"
import "sort"

func maxTotalReward(rewardValues []int) int {
    n := len(rewardValues)
    sort.Ints(rewardValues)
    if n >= 2 && rewardValues[n - 2] == rewardValues[n - 1] - 1 {
        return 2 * rewardValues[n - 1] - 1
    }
    f0, f1 := big.NewInt(1), big.NewInt(0)
    for _, x := range rewardValues {
        mask, one := big.NewInt(0), big.NewInt(1)
        mask.Sub(mask.Lsh(one, uint(x)), one)
        f1.Lsh(f1.And(f0, mask), uint(x))
        f0.Or(f0, f1)
    }
    return f0.BitLen() - 1
}

func maxTotalReward1(nums []int) int {
    sort.Ints(nums)
    set := make(map[int]struct{})
    removeDuplicates := func(nums []int) []int {
        if len(nums) == 0 { return nums }
        res := []int{nums[0]}
        for i := 1; i < len(nums); i++ {
            if nums[i] != nums[i-1] {
                res = append(res, nums[i])
            }
        }
        return res
    }
    nums = removeDuplicates(nums)
    for _, v := range nums {
        set[v] = struct{}{}
    }
    min := func(x, y int) int { if x < y { return x; }; return y; }
    max := func(x, y int) int { if x > y { return x; }; return y; }
    bisectLeft := func(arr []int, target int) int {
        low, high := 0, len(arr)
        for low < high {
            mid := (low + high) / 2
            if arr[mid] < target {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    var helper func(num int, nums []int, set map[int]struct{}, cache map[int]int) int
    helper = func(num int, nums []int, set map[int]struct{}, cache map[int]int) int {
        if val, found := cache[num]; found { return val }
        if _, found := set[num]; found || num == 0 { return num }
        res := 0
        for i := 0; i < bisectLeft(nums, num); i++ {
            res = max(res, nums[i] + helper(min(nums[i]-1, num-nums[i]), nums, set, cache))
        }
        cache[num] = res
        return res
    }
    cache := make(map[int]int)
    return nums[len(nums)-1] + helper(nums[len(nums)-1]-1, nums, set, cache)
}

func main() {
    // Example 1:
    // Input: rewardValues = [1,1,3,3]
    // Output: 4
    // Explanation:
    // During the operations, we can choose to mark the indices 0 and 2 in order, and the total reward will be 4, which is the maximum.
    fmt.Println(maxTotalReward([]int{1,1,3,3})) // 4
    // Example 2:
    // Input: rewardValues = [1,6,4,3,2]
    // Output: 11
    // Explanation:
    // Mark the indices 0, 2, and 1 in order. The total reward will then be 11, which is the maximum.
    fmt.Println(maxTotalReward([]int{1,6,4,3,2})) // 11

    fmt.Println(maxTotalReward1([]int{1,1,3,3})) // 4
    fmt.Println(maxTotalReward1([]int{1,6,4,3,2})) // 11
}