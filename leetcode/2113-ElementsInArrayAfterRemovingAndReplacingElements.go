package main

// 2113. Elements in Array After Removing and Replacing Elements
// You are given a 0-indexed integer array nums. 
// Initially on minute 0, the array is unchanged. 
// Every minute, the leftmost element in nums is removed until no elements remain. 
// Then, every minute, one element is appended to the end of nums, in the order they were removed in, until the original array is restored. 
// This process repeats indefinitely.

//     For example, the array [0,1,2] would change as follows: 
//         [0,1,2] → [1,2] → [2] → [] → [0] → [0,1] → [0,1,2] → [1,2] → [2] → [] → [0] → [0,1] → [0,1,2] → ...

// You are also given a 2D integer array queries of size n where queries[j] = [timej, indexj]. 
// The answer to the jth query is:
//     1. nums[indexj] if indexj < nums.length at minute timej
//     2. -1 if indexj >= nums.length at minute timej

// Return an integer array ans of size n where ans[j] is the answer to the jth query.

// Example 1:
// Input: nums = [0,1,2], queries = [[0,2],[2,0],[3,2],[5,0]]
// Output: [2,2,-1,0]
// Explanation:
// Minute 0: [0,1,2] - All elements are in the nums.
// Minute 1: [1,2]   - The leftmost element, 0, is removed.
// Minute 2: [2]     - The leftmost element, 1, is removed.
// Minute 3: []      - The leftmost element, 2, is removed.
// Minute 4: [0]     - 0 is added to the end of nums.
// Minute 5: [0,1]   - 1 is added to the end of nums.
// At minute 0, nums[2] is 2.
// At minute 2, nums[0] is 2.
// At minute 3, nums[2] does not exist.
// At minute 5, nums[0] is 0.

// Example 2:
// Input: nums = [2], queries = [[0,0],[1,0],[2,0],[3,0]]
// Output: [2,-1,2,-1]
// Minute 0: [2] - All elements are in the nums.
// Minute 1: []  - The leftmost element, 2, is removed.
// Minute 2: [2] - 2 is added to the end of nums.
// Minute 3: []  - The leftmost element, 2, is removed.
// At minute 0, nums[0] is 2.
// At minute 1, nums[0] does not exist.
// At minute 2, nums[0] is 2.
// At minute 3, nums[0] does not exist.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 100
//     n == queries.length
//     1 <= n <= 10^5
//     queries[j].length == 2
//     0 <= timej <= 10^5
//     0 <= indexj < nums.length

import "fmt"

func elementInNums(nums []int, queries [][]int) []int {
    res, n := []int{}, len(nums)
    for _, query := range queries {
        time, index := query[0], query[1]
        m := time % (2 * n)
        if m == n{
            res = append(res, -1)
        } else if m < n {
            if index < n - m {
                res = append(res, nums[m + index])
            } else {
                res = append(res,-1)
            }
        } else {
            if index < m - n{
                res = append(res, nums[index])
            } else {
                res = append(res, -1)
            }
        }
    }
    return res
}

func elementInNums1(nums []int, queries [][]int) []int {
    res, n := make([]int, len(queries)), len(nums)
    query := func(time int, j int) int {
        time %= 2 * n
        if time <= n {
            if j < n - time {
                return nums[time + j]
            } else {
                return -1
            }
        }
        if j < time - n {
            return nums[j]
        }
        return -1
    }
    for i, q := range queries {
        res[i] = query(q[0], q[1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2], queries = [[0,2],[2,0],[3,2],[5,0]]
    // Output: [2,2,-1,0]
    // Explanation:
    // Minute 0: [0,1,2] - All elements are in the nums.
    // Minute 1: [1,2]   - The leftmost element, 0, is removed.
    // Minute 2: [2]     - The leftmost element, 1, is removed.
    // Minute 3: []      - The leftmost element, 2, is removed.
    // Minute 4: [0]     - 0 is added to the end of nums.
    // Minute 5: [0,1]   - 1 is added to the end of nums.
    // At minute 0, nums[2] is 2.
    // At minute 2, nums[0] is 2.
    // At minute 3, nums[2] does not exist.
    // At minute 5, nums[0] is 0.
    fmt.Println(elementInNums([]int{0,1,2}, [][]int{{0,2},{2,0},{3,2},{5,0}})) // [2,2,-1,0]
    // Example 2:
    // Input: nums = [2], queries = [[0,0],[1,0],[2,0],[3,0]]
    // Output: [2,-1,2,-1]
    // Minute 0: [2] - All elements are in the nums.
    // Minute 1: []  - The leftmost element, 2, is removed.
    // Minute 2: [2] - 2 is added to the end of nums.
    // Minute 3: []  - The leftmost element, 2, is removed.
    // At minute 0, nums[0] is 2.
    // At minute 1, nums[0] does not exist.
    // At minute 2, nums[0] is 2.
    // At minute 3, nums[0] does not exist.
    fmt.Println(elementInNums([]int{2}, [][]int{{0,0},{1,0},{2,0},{3,0}})) // [2,-1,2,-1]

    fmt.Println(elementInNums1([]int{0,1,2}, [][]int{{0,2},{2,0},{3,2},{5,0}})) // [2,2,-1,0]
    fmt.Println(elementInNums1([]int{2}, [][]int{{0,0},{1,0},{2,0},{3,0}})) // [2,-1,2,-1]
}