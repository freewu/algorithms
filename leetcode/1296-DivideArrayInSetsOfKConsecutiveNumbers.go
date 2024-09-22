package main

// 1296. Divide Array in Sets of K Consecutive Numbers
// Given an array of integers nums and a positive integer k, 
// check whether it is possible to divide this array into sets of k consecutive numbers.

// Return true if it is possible. Otherwise, return false.

// Example 1:
// Input: nums = [1,2,3,3,4,4,5,6], k = 4
// Output: true
// Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].

// Example 2:
// Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
// Output: true
// Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].

// Example 3:
// Input: nums = [1,2,3,4], k = 3
// Output: false
// Explanation: Each array should be divided in subarrays of size 3.

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "slices"

func isPossibleDivide(nums []int, k int) bool {
    if len(nums) % k != 0 { // 不够分组
        return false
    }
    sort.Ints(nums)
    totalSplits, mapped := 0, make(map[int]int)
    for _, digit := range nums { // 统计牌数量
        mapped[digit]++
    }
    for _, digit := range nums {
        if mapped[digit] > 0 {
            totalSplits++
            for index := 0; index < k; index++ { // 分组
                if mapped[digit] <= 0 { // 不够分组了
                    return false
                } else {
                    mapped[digit]--
                }
                digit++ // 1,2,3 | 2,3,4
            }
        }
    }
    return totalSplits == len(nums) / k
}

func isPossibleDivide1(nums []int, k int) bool {
    if k == 1 {
        return true
    }
    if len(nums) % k != 0 { // 不够分组
        return false
    }
    slices.Sort(nums)
    list := make([][]int, 0)
    for _, num := range nums {
        flag := true
        for i := 0; i < len(list); i++ {
            v := list[i][len(list[i])-1]
            if v + 1 == num {
                list[i] = append(list[i], num)
                if len(list[i]) == k {
                    list = slices.Delete(list, i, i+1)
                }
                flag = false
                break
            } else if v + 1 < num {
                return false
            }
        }
        if flag {
            list = append(list, []int{num})
        }
    }
    return len(list) == 0
}

// 没有要求子数组或子序列
func isPossibleDivide2(nums []int, k int) bool {
    n := len(nums)
    if n % k != 0 {
        return false
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    count := make(map[int]int)
    mn := nums[0]
    for _, v := range nums {
        count[v]++
        mn = min(mn, v)
    }
    findMin := func(count map[int]int) int {
        res := 1 << 31
        for k, v := range count {
            if v > 0 {
                res = min(res, k)
            }
        }
        return res
    }
    used := 0
    for used < n {
        preMin := findMin(count)
        for start := preMin; start < preMin + k; start++ {
            if count[start] <= 0 {
                return false
            }
            count[start]--
            used++
        }
    }
    return true
}

func isPossibleDivide3(nums []int, k int) bool {
    if k > len(nums) || len(nums)%k != 0 {
        return false
    }
    sort.Ints(nums)
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    for i := 0;i < len(nums);i++ {
        if mp[nums[i]] > 0 {
            step := k
            num := nums[i]
            for step > 0 {
                if mp[num] == 0 {
                    return false
                }
                mp[num]--
                step--
                num++
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,3,4,4,5,6], k = 4
    // Output: true
    // Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].
    fmt.Println(isPossibleDivide([]int{1,2,3,3,4,4,5,6}, 4)) // true
    // Example 2:
    // Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
    // Output: true
    // Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].
    fmt.Println(isPossibleDivide([]int{3,2,1,2,3,4,3,4,5,9,10,11}, 3)) // true
    // Example 3:
    // Input: nums = [1,2,3,4], k = 3
    // Output: false
    // Explanation: Each array should be divided in subarrays of size 3.
    fmt.Println(isPossibleDivide([]int{1,2,3,4}, 3)) // false

    fmt.Println(isPossibleDivide1([]int{1,2,3,3,4,4,5,6}, 4)) // true
    fmt.Println(isPossibleDivide1([]int{3,2,1,2,3,4,3,4,5,9,10,11}, 3)) // true
    fmt.Println(isPossibleDivide1([]int{1,2,3,4}, 3)) // false

    fmt.Println(isPossibleDivide2([]int{1,2,3,3,4,4,5,6}, 4)) // true
    fmt.Println(isPossibleDivide2([]int{3,2,1,2,3,4,3,4,5,9,10,11}, 3)) // true
    fmt.Println(isPossibleDivide2([]int{1,2,3,4}, 3)) // false

    fmt.Println(isPossibleDivide3([]int{1,2,3,3,4,4,5,6}, 4)) // true
    fmt.Println(isPossibleDivide3([]int{3,2,1,2,3,4,3,4,5,9,10,11}, 3)) // true
    fmt.Println(isPossibleDivide3([]int{1,2,3,4}, 3)) // false
}