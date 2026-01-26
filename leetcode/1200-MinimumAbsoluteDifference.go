package main

// 1200. Minimum Absolute Difference
// Given an array of distinct integers arr, 
// find all pairs of elements with the minimum absolute difference of any two elements.

// Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
//     a, b are from arr
//     a < b
//     b - a equals to the minimum absolute difference of any two elements in arr
 
// Example 1:
// Input: arr = [4,2,1,3]
// Output: [[1,2],[2,3],[3,4]]
// Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

// Example 2:
// Input: arr = [1,3,6,10,15]
// Output: [[1,3]]

// Example 3:
// Input: arr = [3,8,-10,23,19,-4,-14,27]
// Output: [[-14,-10],[19,23],[23,27]]

// Constraints:
//     2 <= arr.length <= 10^5
//     -10^6 <= arr[i] <= 10^6

import "fmt"
import "sort"

func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr) // 从小到大排序
    res, n, mindiff := [][]int{}, len(arr), 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n - 1; i++ { // get minimum abs difference  
        mindiff = min(mindiff, arr[i+1] - arr[i])
    }
    for i := 0; i < n - 1; i++ {
        diff := arr[i+1] - arr[i]
        if diff == mindiff {
            res = append(res, []int{arr[i], arr[i+1]})
        }
    }
    return res
}

func minimumAbsDifference1(arr []int) [][]int {
    sort.Ints(arr)
    res := [][]int{}
    if len(arr) < 2 {
        return res // 如果数组长度小于2，直接返回空切片
    }
    mn := arr[1] - arr[0] // 初始化 mn 为第一对元素的差值
    for i := 0; i < len(arr) - 1; i++ {
        mid := arr[i+1] - arr[i]
        if mid < mn { // 发现更小的
            mn = mid
            res = [][]int{{arr[i], arr[i+1]}} // 重新初始化 res 添加当前元素对
        } else if mid == mn {
            res = append(res, []int{arr[i], arr[i+1]}) // 添加当前元素对
        }
    }
    return res
}

func minimumAbsDifference2(arr []int) [][]int {
    sort.Ints(arr)
    res, mn := [][]int{}, 1 << 31
    for i, x := range arr[:len(arr)-1] {
        y := arr[i+1]
        diff := y - x
        if diff < mn {
            mn = diff
            res = [][]int{{x, y}}
        } else if diff == mn {
            res = append(res, []int{x, y})
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [4,2,1,3]
    // Output: [[1,2],[2,3],[3,4]]
    // Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
    fmt.Println(minimumAbsDifference([]int{4,2,1,3})) // [[1,2],[2,3],[3,4]]
    // Example 2:
    // Input: arr = [1,3,6,10,15]
    // Output: [[1,3]]
    fmt.Println(minimumAbsDifference([]int{1,3,6,10,15})) // [[1,3]]
    // Example 3:
    // Input: arr = [3,8,-10,23,19,-4,-14,27]
    // Output: [[-14,-10],[19,23],[23,27]]
    fmt.Println(minimumAbsDifference([]int{3,8,-10,23,19,-4,-14,27})) // [[-14,-10],[19,23],[23,27]]

    fmt.Println(minimumAbsDifference([]int{1,2,3,4,5,6,7,8,9})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]
    fmt.Println(minimumAbsDifference([]int{9,8,7,6,5,4,3,2,1})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]

    fmt.Println(minimumAbsDifference1([]int{4,2,1,3})) // [[1,2],[2,3],[3,4]]
    fmt.Println(minimumAbsDifference1([]int{1,3,6,10,15})) // [[1,3]]
    fmt.Println(minimumAbsDifference1([]int{3,8,-10,23,19,-4,-14,27})) // [[-14,-10],[19,23],[23,27]]
    fmt.Println(minimumAbsDifference1([]int{1,2,3,4,5,6,7,8,9})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]
    fmt.Println(minimumAbsDifference1([]int{9,8,7,6,5,4,3,2,1})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]

    fmt.Println(minimumAbsDifference2([]int{4,2,1,3})) // [[1,2],[2,3],[3,4]]
    fmt.Println(minimumAbsDifference2([]int{1,3,6,10,15})) // [[1,3]]
    fmt.Println(minimumAbsDifference2([]int{3,8,-10,23,19,-4,-14,27})) // [[-14,-10],[19,23],[23,27]]
    fmt.Println(minimumAbsDifference2([]int{1,2,3,4,5,6,7,8,9})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]
    fmt.Println(minimumAbsDifference2([]int{9,8,7,6,5,4,3,2,1})) // [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7] [7 8] [8 9]]
}