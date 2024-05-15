package main

// LCR 060. 前 K 个高频元素
// 给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]

// 示例 2:
// 输入: nums = [1], k = 1
// 输出: [1]

// 提示：
// 1 <= nums.length <= 10^5
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 
// 进阶：所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。 

import "fmt"
import "slices"

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int,0) 
    for _, v := range nums { // 先统计频率
        m[v]++
    }
    arr := [][]int{}
    for k,v := range m { // 将 map 转成数组
        arr = append(arr,[]int{k,v})
    }
    slices.SortFunc(arr, func(a1, a2 []int) int { // 排序
        return a2[1] - a1[1]
    }) 
    res := []int{}
    for _, v := range arr {
        if len(res) == k { // 取 top k
            break
        }
        res = append(res,v[0])
    }
    return res
}

func topKFrequent1(nums []int, k int) []int {
    frequency := make(map[int]int)
    for _, v := range nums { // 统计元素出现频次
        frequency[v] += 1
    }
    res, bucket := []int{}, make([][]int, len(nums)+1)
    for k, v := range frequency {
        bucket[v] = append(bucket[v], k)
    }
    for i, cnt := len(bucket) - 1, 0; i >= 0 && cnt < k; i--  {
        res = append(res, bucket[i]...)
    }
    return res[:k]
}

func topKFrequent2(nums []int, k int) []int {
    res, n := []int{}, len(nums)
    frequency, m := make([][]int, n + 1), map[int]int{}
    for _, x := range nums {
        m[x]++
    }
    for k, v := range m {
        frequency[v] = append(frequency[v], k)
    }
    for i := n; i > 0 && k > 0; i-- {
        if len(frequency[i]) > 0 {
            for _, x := range frequency[i] {
                res = append(res, x)
            }
            k -= len(frequency[i])
        }
    }
    return res
}

func main() {
    fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2)) // [1,2]
    fmt.Println(topKFrequent([]int{1},1)) // [1]

    fmt.Println(topKFrequent1([]int{1,1,1,2,2,3},2)) // [1,2]
    fmt.Println(topKFrequent1([]int{1},1)) // [1]

    fmt.Println(topKFrequent2([]int{1,1,1,2,2,3},2)) // [1,2]
    fmt.Println(topKFrequent2([]int{1},1)) // [1]
}