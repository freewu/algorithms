package main

// 2459. Sort Array by Moving Items to Empty Space
// You are given an integer array nums of size n containing each element from 0 to n - 1 (inclusive). 
// Each of the elements from 1 to n - 1 represents an item, and the element 0 represents an empty space.

// In one operation, you can move any item to the empty space. 
// nums is considered to be sorted if the numbers of all the items are in ascending order and the empty space is either at the beginning or at the end of the array.

// For example, if n = 4, nums is sorted if:
//     nums = [0,1,2,3] or
//     nums = [1,2,3,0]

// ...and considered to be unsorted otherwise.

// Return the minimum number of operations needed to sort nums.

// Example 1:
// Input: nums = [4,2,0,3,1]
// Output: 3
// Explanation:
// - Move item 2 to the empty space. Now, nums = [4,0,2,3,1].
// - Move item 1 to the empty space. Now, nums = [4,1,2,3,0].
// - Move item 4 to the empty space. Now, nums = [0,1,2,3,4].
// It can be proven that 3 is the minimum number of operations needed.

// Example 2:
// Input: nums = [1,2,3,4,0]
// Output: 0
// Explanation: nums is already sorted so return 0.

// Example 3:
// Input: nums = [1,0,2,4,3]
// Output: 2
// Explanation:
// - Move item 2 to the empty space. Now, nums = [1,2,0,4,3].
// - Move item 3 to the empty space. Now, nums = [1,2,3,4,0].
// It can be proven that 2 is the minimum number of operations needed.

// Constraints:
//     n == nums.length
//     2 <= n <= 10^5
//     0 <= nums[i] < n
//     All the values of nums are unique.

import "fmt"

func sortArray(nums []int) int {
    getMinCount := func(nums []int, dis int) int {
        res, n, visited := 0, len(nums), make(map[int]bool)
        for i := 0; i < n; i++ {
            j := i
            if j != nums[j] && !visited[j] {
                tempVisited := make(map[int]bool)
                for !tempVisited[j] {
                    tempVisited[j], visited[j] = true, true
                    j = nums[j]
                }
                // 如果当前环中包含目的下标，少一次调换。否则，需要多一次调换（因为调换依赖目的下标0参数）
                res += len(tempVisited)
                if tempVisited[dis] {
                    res--
                } else {
                    res++
                }
            }
        }
        return res
    }
    n := len(nums)
    nums2 := make([]int, n) // 计算空位在下标 n - 1，[0, n - 2] 塞满数据
    for i := 0; i < n; i++ {
        nums2[i] = (nums[i] + n - 1) % n
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(getMinCount(nums, 0), getMinCount(nums2, n - 1)) // 两种情况取一个最小的数
}

func sortArray1(nums []int) int {
    n := len(nums)
    calc := func(nums []int, k int) int {
        count, visited := 0, make([]bool, n)
        for i, v := range nums {
            if i == v || visited[i] { continue }
            count++
            j := i
            for !visited[j] {
                visited[j] = true
                count++
                j = nums[j]
            }
        }
        if nums[k] != k { count -= 2 }
        return count
    }
    arr := make([]int, n)
    for i, v := range nums {
        arr[i] = (v - 1 + n) % n
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(calc(nums, 0), calc(arr, n - 1))
}

func main() {
    // Example 1:
    // Input: nums = [4,2,0,3,1]
    // Output: 3
    // Explanation:
    // - Move item 2 to the empty space. Now, nums = [4,0,2,3,1].
    // - Move item 1 to the empty space. Now, nums = [4,1,2,3,0].
    // - Move item 4 to the empty space. Now, nums = [0,1,2,3,4].
    // It can be proven that 3 is the minimum number of operations needed.
    fmt.Println(sortArray([]int{4,2,0,3,1})) // 3
    // Example 2:
    // Input: nums = [1,2,3,4,0]
    // Output: 0
    // Explanation: nums is already sorted so return 0.
    fmt.Println(sortArray([]int{1,2,3,4,0})) // 0
    // Example 3:
    // Input: nums = [1,0,2,4,3]
    // Output: 2
    // Explanation:
    // - Move item 2 to the empty space. Now, nums = [1,2,0,4,3].
    // - Move item 3 to the empty space. Now, nums = [1,2,3,4,0].
    // It can be proven that 2 is the minimum number of operations needed.
    fmt.Println(sortArray([]int{1,0,2,4,3})) // 2

    fmt.Println(sortArray1([]int{4,2,0,3,1})) // 3
    fmt.Println(sortArray1([]int{1,2,3,4,0})) // 0
    fmt.Println(sortArray1([]int{1,0,2,4,3})) // 2
}