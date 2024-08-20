package main

// 985. Sum of Even Numbers After Queries
// You are given an integer array nums and an array queries where queries[i] = [vali, indexi].
// For each query i, first, apply nums[indexi] = nums[indexi] + vali, then print the sum of the even values of nums.
// Return an integer array answer where answer[i] is the answer to the ith query.

// Example 1:
// Input: nums = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
// Output: [8,6,2,4]
// Explanation: At the beginning, the array is [1,2,3,4].
// After adding 1 to nums[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
// After adding -3 to nums[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
// After adding -4 to nums[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
// After adding 2 to nums[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.

// Example 2:
// Input: nums = [1], queries = [[4,0]]
// Output: [0]

// Constraints:
//     1 <= nums.length <= 10^4
//     -10^4 <= nums[i] <= 10^4
//     1 <= queries.length <= 10^4
//     -10^4 <= vali <= 10^4
//     0 <= indexi < nums.length

import "fmt"

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    sum := 0
    for _, v := range nums { // 累加偶数的总和
        if v % 2 == 0 {
            sum += v
        }
    }
    res := make([]int, 0, len(queries))
    for _, query := range queries {
        v, index := query[0], query[1]
        old, new := nums[index], nums[index] + v
        if old % 2 == 0 { sum -= old }
        if new % 2 == 0 { sum += new }
        nums[index] = new
        res = append(res, sum)
    }
    return res
}

func sumEvenAfterQueries1(nums []int, queries [][]int) []int {
    mp, sum, res := make(map[int]bool), 0, []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] & 1 == 0 {
            sum += nums[i]
            mp[i] = true
        }
    }
    for _, query := range queries {
        val, index  := query[0], query[1]
        t := val + nums[index]
        if _, ok := mp[index]; ok {
            if t & 1 == 0 {
                sum += val
            } else {
                sum -= nums[index]
                delete(mp, index)
            }
            nums[index] = t
        } else {
            if t & 1 == 0 {
                sum += t
                mp[index] = true
            }
            nums[index] = t

        }
        res = append(res, sum)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
    // Output: [8,6,2,4]
    // Explanation: At the beginning, the array is [1,2,3,4].
    // After adding 1 to nums[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
    // After adding -3 to nums[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
    // After adding -4 to nums[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
    // After adding 2 to nums[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.
    fmt.Println(sumEvenAfterQueries([]int{1,2,3,4}, [][]int{{1,0},{-3,1},{-4,0},{2,3}})) // [8,6,2,4]
    // Example 2:
    // Input: nums = [1], queries = [[4,0]]
    // Output: [0]
    fmt.Println(sumEvenAfterQueries([]int{1}, [][]int{{4,0}})) // [0]

    fmt.Println(sumEvenAfterQueries1([]int{1,2,3,4}, [][]int{{1,0},{-3,1},{-4,0},{2,3}})) // [8,6,2,4]
    fmt.Println(sumEvenAfterQueries1([]int{1}, [][]int{{4,0}})) // [0]
}