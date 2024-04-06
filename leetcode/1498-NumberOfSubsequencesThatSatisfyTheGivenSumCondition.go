package main

// 1498. Number of Subsequences That Satisfy the Given Sum Condition
// You are given an array of integers nums and an integer target.
// Return the number of non-empty subsequences of nums such that the sum of the minimum and maximum element on it is less or equal to target. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [3,5,6,7], target = 9
// Output: 4
// Explanation: There are 4 subsequences that satisfy the condition.
// [3] -> Min value + max value <= target (3 + 3 <= 9)
// [3,5] -> (3 + 5 <= 9)
// [3,5,6] -> (3 + 6 <= 9)
// [3,6] -> (3 + 6 <= 9)

// Example 2:
// Input: nums = [3,3,6,8], target = 10
// Output: 6
// Explanation: There are 6 subsequences that satisfy the condition. (nums can have repeated numbers).
// [3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]

// Example 3:
// Input: nums = [2,3,3,4,6,7], target = 12
// Output: 61
// Explanation: There are 63 non-empty subsequences, two of them do not satisfy the condition ([6,7], [7]).
// Number of valid subsequences (63 - 2 = 61).
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6
//     1 <= target <= 10^6

import "fmt"
import "sort"

const mod int = 1e9 + 7

// 双指针
func numSubseq(nums []int, target int) int {
    sort.Ints(nums)
    res := 0
    var pow func (x, n int) int 
    pow = func (x, n int) int {
        if n == 0 {return 1}
        y := pow(x, n / 2)
        if n % 2 == 1 {return (((y * y) % mod) * x) % mod}
        return (y * y) % mod
    }
    for start, end := 0, len(nums) - 1; start <= end; start++ {
        for ; end >= start && nums[start] + nums[end] > target; end-- {}
        if end < start { 
            break
        }
        res = (res + pow(2, end - start)) % mod
    }
    return res
}

func numSubseq1(nums []int, target int) int {
    res := 0
    sort.Ints(nums)
    var pow func (n int) int
    pow = func (n int) int {
        if n == 1 { return 2; }
        tmp := pow(n >> 1)
        tmp *= tmp
        if n % 2 == 0 { return tmp % 1000000007; }
        return (tmp << 1) % 1000000007
    }
    left, right := 0, len(nums) - 1
    for left < right {
        if nums[left] + nums[right] > target {
            right --
        } else {
            res =  (res + pow(right - left)) % 1000000007
            left ++
        }
    }
    if nums[left] << 1 > target {
        return res
    }
    return res + 1
}

func main() {
    // Explanation: There are 4 subsequences that satisfy the condition.
    // [3] -> Min value + max value <= target (3 + 3 <= 9)
    // [3,5] -> (3 + 5 <= 9)
    // [3,5,6] -> (3 + 6 <= 9)
    // [3,6] -> (3 + 6 <= 9)
    fmt.Println(numSubseq([]int{3,5,6,7}, 9)) // 4
    // Explanation: There are 6 subsequences that satisfy the condition. (nums can have repeated numbers).
    // [3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]
    fmt.Println(numSubseq([]int{3,3,6,8}, 10)) // 6

    // Explanation: There are 63 non-empty subsequences, two of them do not satisfy the condition ([6,7], [7]).
    // Number of valid subsequences (63 - 2 = 61).
    fmt.Println(numSubseq([]int{2,3,3,4,6,7}, 12)) // 61

    fmt.Println(numSubseq1([]int{3,5,6,7}, 9)) // 4
    fmt.Println(numSubseq1([]int{3,3,6,8}, 10)) // 6
    fmt.Println(numSubseq1([]int{2,3,3,4,6,7}, 12)) // 61
}