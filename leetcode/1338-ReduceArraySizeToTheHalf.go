package main

// 1338. Reduce Array Size to The Half
// You are given an integer array arr. 
// You can choose a set of integers and remove all the occurrences of these integers in the array.

// Return the minimum size of the set so that at least half of the integers of the array are removed.

// Example 1:
// Input: arr = [3,3,3,3,5,5,5,2,2,7]
// Output: 2
// Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
// Possible sets of size 2 are {3,5},{3,2},{5,2}.
// Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has a size greater than half of the size of the old array.

// Example 2:
// Input: arr = [7,7,7,7,7,7]
// Output: 1
// Explanation: The only possible set you can choose is {7}. This will make the new array empty.

// Constraints:
//     2 <= arr.length <= 10^5
//     arr.length is even.
//     1 <= arr[i] <= 10^5

import "fmt"
import "sort"

func minSetSize(arr []int) int {
    mp, count := make(map[int]int), []int{}
    for _, v := range arr { // 统计出现次数
        mp[v]++
    }
    for _, v := range mp {
        count = append(count, v)
    }
    sort.Ints(count)
    n, m, sum := len(arr), len(count), 0
    for i := m - 1; i >= 0; i --{
        sum += count[i]
        if sum >= n / 2 { // 超过一半元素了
            return m - i
        }
    }
    return n / 2 // 处理无数值重复的情况
}

func minSetSize1(arr []int) int {
    mx := 0
    for _, v := range arr { // 找出最大值
        if v > mx { mx = v }
    }
    count := make([]int, mx + 1)
    for _, v := range arr {
        count[v]++
    }
    sort.Slice(count, func(i, j int) bool { // 从大到小排序
        return count[i] > count[j]
    })
    res, sum, n := 0, 0, len(arr)
    for i, v := range count {
        if sum >= n / 2 { break } // 超过一半
        sum += v
        res = i + 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [3,3,3,3,5,5,5,2,2,7]
    // Output: 2
    // Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
    // Possible sets of size 2 are {3,5},{3,2},{5,2}.
    // Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has a size greater than half of the size of the old array.
    fmt.Println(minSetSize([]int{3,3,3,3,5,5,5,2,2,7})) // 2
    // Example 2:
    // Input: arr = [7,7,7,7,7,7]
    // Output: 1
    // Explanation: The only possible set you can choose is {7}. This will make the new array empty.
    fmt.Println(minSetSize([]int{7,7,7,7,7,7})) // 1

    fmt.Println(minSetSize1([]int{3,3,3,3,5,5,5,2,2,7})) // 2
    fmt.Println(minSetSize1([]int{7,7,7,7,7,7})) // 1
}