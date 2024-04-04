package main

// 347. Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most frequent elements. 
// You may return the answer in any order.

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
 
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4
//     k is in the range [1, the number of unique elements in the array].
//     It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

import "fmt"
import "slices"

func topKFrequent(nums []int, k int) []int {
    // 先统计频率
    m := make(map[int]int,0) 
    for _, v := range nums {
        m[v]++
    }
    arr := [][]int{}
    // 将 map 转成数组
    for k,v := range m {
        arr = append(arr,[]int{k,v})
    }
    // 排序
    slices.SortFunc(arr, func(a1, a2 []int) int {
        return a2[1] - a1[1]
    })
    // 取 top k
    res := []int{}
    for _, v := range arr {
        if len(res) == k {
            break
        }
        res = append(res,v[0])
    }
    return res
}

func topKFrequent1(nums []int, k int) []int {
    var frequency map[int]int = make(map[int]int)
    for _, v := range nums {
        frequency[v] += 1
    }
    var bucket [][]int = make([][]int, len(nums)+1)
    var res []int
    for k, v := range frequency {
        bucket[v] = append(bucket[v], k)
    }
    for i, cnt := len(bucket) - 1, 0; i >= 0 && cnt < k; i--  {
        res = append(res, bucket[i]...)
    }
    return res[:k]
}

func main() {
    fmt.Println(topKFrequent([]int{1,1,1,2,2,3},2)) // [1,2]
    fmt.Println(topKFrequent([]int{1},1)) // [1]

    fmt.Println(topKFrequent1([]int{1,1,1,2,2,3},2)) // [1,2]
    fmt.Println(topKFrequent1([]int{1},1)) // [1]
}