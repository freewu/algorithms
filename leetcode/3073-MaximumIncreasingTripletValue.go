package main

// 3073. Maximum Increasing Triplet Value
// Given an array nums, 
// return the maximum value of a triplet (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].

// The value of a triplet (i, j, k) is nums[i] - nums[j] + nums[k].

// Example 1:
// Input: nums = [5,6,9]
// Output: 8
// Explanation: We only have one choice for an increasing triplet and that is choosing all three elements. The value of this triplet would be 5 - 6 + 9 = 8.

// Example 2:
// Input: nums = [1,5,3,6]
// Output: 4
// Explanation: There are only two increasing triplets:
// (0, 1, 3): The value of this triplet is nums[0] - nums[1] + nums[3] = 1 - 5 + 6 = 2.
// (0, 2, 3): The value of this triplet is nums[0] - nums[2] + nums[3] = 1 - 3 + 6 = 4.
// Thus the answer would be 4.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     The input is generated such that at least one triplet meets the given condition.

import "fmt"
// import "sort"

// func maximumTripletValue(nums []int) int {
//     res, n := 0, len(nums)
//     rightArr := make([]int, n)
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for i := n - 2; i >= 0; i-- {
//         rightArr[i] = max(rightArr[i + 1], nums[i + 1])
//     }
//     fmt.Println(rightArr)
//     lower := func(arr []int, k int) int {
//         sort.Ints(arr)
//         if arr[0] >= k { return -1 }
//         for i := range arr {
//             if arr[i] >= k {
//                 return arr[i - 1]
//             }
//         }
//         return -1
//     }
//     mp, set := make(map[int]bool), []int{}
//     for i := 1; i <= n - 2; i++ {
//         if !mp[nums[i - 1]] {
//             set = append(set, nums[i - 1])
//             mp[nums[i - 1]] = true
//         }
//         left := lower(set, nums[i]) // 找到小于 nums[i] 的值
//         fmt.Println(left)
//         if  left != -1 && rightArr[i] > nums[i] {
//             res = max(res, left - nums[i] + rightArr[i])
//         }
//     }
//     return res
// }

// func maximumTripletValue(nums []int) int {
//     res, n := 0, len(nums)
//     sufMax := make([]int, n)
//     sufMax[n-1] = nums[n-1]
//     sort_prev := []int{}
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for i := n - 2; i >= 0; i-- {
//         sufMax[i] = max(sufMax[i+1], nums[i])
//     }
//     getLowwerBound := func(arr []int, target int) int {
//         res, left, right := -1, 0, len(arr) - 1
//         for left <= right {
//             mid := (left + right) >> 1
//             if arr[mid] < target {
//                 res = mid
//                 left = mid + 1
//             } else { 
//                 right = mid - 1
//             }
//         }
//         return res
//     }
//     getUpperBound := func(arr []int, target int) int {
//         res, left, right := -1, 0, len(arr) - 1
//         for left <= right {
//             mid := (left + right) >> 1
//             if arr[mid] >= target {
//                 res = mid
//                 right = mid-1
//             } else {
//                 left = mid + 1
//             }
//         }
//         return res
//     }
//     InsertElement := func (s []int, index int, elem int) []int {
//         // 在index前插入elem
//         s = append(s, 0) // 为新元素预留空间，并且不覆盖原始切片
//         copy(s[index+1:], s[index:]) // 将index位置以及其后的元素向后移动
//         s[index] = elem // 插入元素
//         return s
//     }
//     insertToSortArr := func(target int) {
//         n := len(sort_prev)
//         if n == 0 || sort_prev[n -1] <= target { // 处理值大于结尾的情况
//             sort_prev = append(sort_prev, target)
//             return
//         }
//         if sort_prev[0] >= target { // 处理值小于开头的情况
//             sort_prev = append([]int{ target }, sort_prev...)
//             return
//         }
//         index := getLowwerBound(sort_prev,target)
//         sort_prev = InsertElement(sort_prev, index + 1, target)
//     }
//     insertToSortArr(nums[0])
//     for i := 1; i < n - 1; i++  {
//         mid := nums[i];
//         index := getUpperBound(sort_prev, mid)
//         three := sufMax[i+1]
//         if index != -1 && mid < three {
//             res = max(res, sort_prev[index] - mid + three)
//         }
//         insertToSortArr(nums[i])
//     }
//     fmt.Println(sort_prev)
//     return res
// }

// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/treeset
func maximumTripletValue(nums []int) int {
    res, n := -1, len(nums)
    right := make([]int, n)
    right[n-1] = nums[n-1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        right[i] = max(nums[i], right[i+1])
    }
    ts := treemap.NewWithIntComparator()
    ts.Put(nums[0], nil)
    for j := 1; j < n-1; j++ {
        if right[j+1] > nums[j] {
            val, _ := ts.Floor(nums[j] - 1)
            if val != nil {
                res = max(res, val.(int)-nums[j]+right[j+1])
            }
        }
        ts.Put(nums[j], nil)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,6,9]
    // Output: 8
    // Explanation: We only have one choice for an increasing triplet and that is choosing all three elements. The value of this triplet would be 5 - 6 + 9 = 8.
    fmt.Println(maximumTripletValue([]int{5,6,9})) // 8
    // Example 2:
    // Input: nums = [1,5,3,6]
    // Output: 4
    // Explanation: There are only two increasing triplets:
    // (0, 1, 3): The value of this triplet is nums[0] - nums[1] + nums[3] = 1 - 5 + 6 = 2.
    // (0, 2, 3): The value of this triplet is nums[0] - nums[2] + nums[3] = 1 - 3 + 6 = 4.
    // Thus the answer would be 4.
    fmt.Println(maximumTripletValue([]int{1,5,3,6})) // 4
}