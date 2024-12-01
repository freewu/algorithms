package main

// 2615. Sum of Distances
// You are given a 0-indexed integer array nums. 
// There exists an array arr of length nums.length, where arr[i] is the sum of |i - j| over all j such that nums[j] == nums[i] and j != i. 
// If there is no such j, set arr[i] to be 0.

// Return the array arr.

// Example 1:
// Input: nums = [1,3,1,1,2]
// Output: [5,0,3,4,0]
// Explanation: 
// When i = 0, nums[0] == nums[2] and nums[0] == nums[3]. Therefore, arr[0] = |0 - 2| + |0 - 3| = 5. 
// When i = 1, arr[1] = 0 because there is no other index with value 3.
// When i = 2, nums[2] == nums[0] and nums[2] == nums[3]. Therefore, arr[2] = |2 - 0| + |2 - 3| = 3. 
// When i = 3, nums[3] == nums[0] and nums[3] == nums[2]. Therefore, arr[3] = |3 - 0| + |3 - 2| = 4. 
// When i = 4, arr[4] = 0 because there is no other index with value 2. 

// Example 2:
// Input: nums = [0,5,3]
// Output: [0,0,0]
// Explanation: Since each element in nums is distinct, arr[i] = 0 for all i.

// Constraints:
// 1 <= nums.length <= 10^5
// 0 <= nums[i] <= 10^9

// Note: This question is the same as 2121: Intervals Between Identical Elements.

import "fmt"
import "sort"

func distance(nums []int) []int64 {
    data, sum := make(map[int][]int), make(map[int]int64)
    for i, v := range nums {
        data[v] = append(data[v], i)
        sum[v] += int64(i)
    }
    res := make([]int64, len(nums))
    for i, rows := range data {
        left, right := int64(0), int64(0)
        for t, index := range rows {
            right = sum[i] - left - int64(index)
            res[index] = right - left - int64(len(rows) - t - 1 - t) * int64(index)
            left += int64(index)
        }
    }
    return res
}

func distance1(nums []int) []int64 {
    mp := make(map[int][]int) 
    for i,v := range nums { // 使用一个map[v][]int来存储相同元素的所有下标
        mp[v] = append(mp[v],i)
    }
    res := make([]int64, len(nums))
    for _, row := range mp { // 遍历 map，计算每个v的数组p的所以元素，先计算第一个与别的元素的差值sum
        sum  := int64(0)
        for _, v  := range row {  // 先统计第一个元素的总和
            sum += int64(v - row[0])
        }
        res[row[0]] = sum // 赋值第一个数字的差值
        for i := 1; i < len(row); i++ { // 依次赋值剩下的差值
            sum += int64((2 * i - len(row))) * int64((row[i] -row[i-1]))
            res[row[i]] = sum
        }
    }
    return res
}

func distance2(nums []int) []int64 {
    res, buckets := make([]int64, len(nums)), make(map[int][]int)
    for i, v := range nums {
        buckets[v] = append(buckets[v], i)
    }
    for _, row := range buckets {
        n := len(row)
        if n == 1 { continue }
        prefix := 0
        for i := 1; i < n; i++ {
            prefix += row[i] - row[0]
        }
        res[row[0]] = int64(prefix)
        for i := 1; i < n; i++ {
            prefix += (row[i] - row[i-1]) * (i - (n - i))
            res[row[i]] = int64(prefix)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1,1,2]
    // Output: [5,0,3,4,0]
    // Explanation: 
    // When i = 0, nums[0] == nums[2] and nums[0] == nums[3]. Therefore, arr[0] = |0 - 2| + |0 - 3| = 5. 
    // When i = 1, arr[1] = 0 because there is no other index with value 3.
    // When i = 2, nums[2] == nums[0] and nums[2] == nums[3]. Therefore, arr[2] = |2 - 0| + |2 - 3| = 3. 
    // When i = 3, nums[3] == nums[0] and nums[3] == nums[2]. Therefore, arr[3] = |3 - 0| + |3 - 2| = 4. 
    // When i = 4, arr[4] = 0 because there is no other index with value 2. 
    fmt.Println(distance([]int{1,3,1,1,2})) // [5,0,3,4,0]
    // Example 2:
    // Input: nums = [0,5,3]
    // Output: [0,0,0]
    // Explanation: Since each element in nums is distinct, arr[i] = 0 for all i.
    fmt.Println(distance([]int{0,5,3})) // [0,0,0]

    fmt.Println(distance([]int{1,3,1,1,2})) // [5,0,3,4,0]

    fmt.Println(distance1([]int{1,3,1,1,2})) // [5,0,3,4,0]
    fmt.Println(distance1([]int{0,5,3})) // [0,0,0]
    fmt.Println(distance1([]int{1,3,1,1,2})) // [5,0,3,4,0]

    fmt.Println(distance2([]int{1,3,1,1,2})) // [5,0,3,4,0]
    fmt.Println(distance2([]int{0,5,3})) // [0,0,0]
    fmt.Println(distance2([]int{1,3,1,1,2})) // [5,0,3,4,0]
}