package main

// 3267. Count Almost Equal Pairs II
// Attention: In this version, the number of operations that can be performed, has been increased to twice.

// You are given an array nums consisting of positive integers.

// We call two integers x and y almost equal if both integers can become equal after performing the following operation at most twice:
//     Choose either x or y and swap any two digits within the chosen number.

// Return the number of indices i and j in nums where i < j such that nums[i] and nums[j] are almost equal.

// Note that it is allowed for an integer to have leading zeros after performing an operation.

// Example 1:
// Input: nums = [1023,2310,2130,213]
// Output: 4
// Explanation:
// The almost equal pairs of elements are:
// 1023 and 2310. By swapping the digits 1 and 2, and then the digits 0 and 3 in 1023, you get 2310.
// 1023 and 213. By swapping the digits 1 and 0, and then the digits 1 and 2 in 1023, you get 0213, which is 213.
// 2310 and 213. By swapping the digits 2 and 0, and then the digits 3 and 2 in 2310, you get 0213, which is 213.
// 2310 and 2130. By swapping the digits 3 and 1 in 2310, you get 2130.

// Example 2:
// Input: nums = [1,10,100]
// Output: 3
// Explanation:
// The almost equal pairs of elements are:
// 1 and 10. By swapping the digits 1 and 0 in 10, you get 01 which is 1.
// 1 and 100. By swapping the second 0 with the digit 1 in 100, you get 001, which is 1.
// 10 and 100. By swapping the first 0 with the digit 1 in 100, you get 010, which is 10.

// Constraints:
//     2 <= nums.length <= 5000
//     1 <= nums[i] < 10^7

import "fmt"
import "sort"
import "strconv"

func countPairs(nums []int) int {
    sort.Ints(nums)
    atoi := func (s []byte) int { // []byte => int
        res := 0
        for _, v := range s { 
            res = res * 10 + int(v & 15)
        }
        return res
    }
    res, mp := 0, make(map[int]int)
    for _, v := range nums {
        set := map[int]bool{ v: true} // 不交换
        s := []byte( strconv.Itoa(v) )
        n := len(s)
        for i := range s {
            for j := i + 1; j < n; j++ {
                s[i], s[j] = s[j], s[i]
                set[atoi(s)] = true // 交换一次
                for p := i + 1; p < n; p++ {
                    for q := p + 1; q < n; q++ {
                        s[p], s[q] = s[q], s[p]
                        set[atoi(s)] = true // 交换两次
                        s[p], s[q] = s[q], s[p]
                    }
                }
                s[i], s[j] = s[j], s[i]
            }
        }
        for v := range set {
            res += mp[v]
        }
        mp[v]++
    }
    return res
}

func countPairs1(nums []int) int {
    sort.Ints(nums)
    pow10 := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
    res := 0
    mp, arr := make(map[int]int), [7]int{}
    for _, x := range nums {
        st := map[int]bool{ x : true } // 不交换
        m := 0
        for v := x; v > 0; v /= 10 {
            arr[m] = v % 10
            m++
        }
        for i := 0; i < m; i++ {
            for j := i + 1; j < m; j++ {
                if arr[i] == arr[j] { continue }
                y := x + (arr[j] - arr[i]) * (pow10[i] - pow10[j])
                st[y] = true // 交换一次
                arr[i], arr[j] = arr[j], arr[i]
                for p := i + 1; p < m; p++ {
                    for q := p + 1; q < m; q++ {
                        st[y + (arr[q] - arr[p]) * (pow10[p] - pow10[q])] = true // 交换两次
                    }
                }
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
        for v := range st {
            res += mp[v]
        }
        mp[x]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1023,2310,2130,213]
    // Output: 4
    // Explanation:
    // The almost equal pairs of elements are:
    // 1023 and 2310. By swapping the digits 1 and 2, and then the digits 0 and 3 in 1023, you get 2310.
    // 1023 and 213. By swapping the digits 1 and 0, and then the digits 1 and 2 in 1023, you get 0213, which is 213.
    // 2310 and 213. By swapping the digits 2 and 0, and then the digits 3 and 2 in 2310, you get 0213, which is 213.
    // 2310 and 2130. By swapping the digits 3 and 1 in 2310, you get 2130.
    fmt.Println(countPairs([]int{1023,2310,2130,213})) // 4
    // Example 2:
    // Input: nums = [1,10,100]
    // Output: 3
    // Explanation:
    // The almost equal pairs of elements are:
    // 1 and 10. By swapping the digits 1 and 0 in 10, you get 01 which is 1.
    // 1 and 100. By swapping the second 0 with the digit 1 in 100, you get 001, which is 1.
    // 10 and 100. By swapping the first 0 with the digit 1 in 100, you get 010, which is 10.
    fmt.Println(countPairs([]int{1,10,100})) // 3

    fmt.Println(countPairs1([]int{1023,2310,2130,213})) // 4
    fmt.Println(countPairs1([]int{1,10,100})) // 3
}