package main

// 3080. Mark Elements on Array by Performing Queries
// You are given a 0-indexed array nums of size n consisting of positive integers.

// You are also given a 2D array queries of size m where queries[i] = [indexi, ki].

// Initially all elements of the array are unmarked.

// You need to apply m queries on the array in order, where on the ith query you do the following:
//     1. Mark the element at index indexi if it is not already marked.
//     2. Then mark ki unmarked elements in the array with the smallest values. 
//        If multiple such elements exist, mark the ones with the smallest indices. 
//        And if less than ki unmarked elements exist, then mark all of them.

// Return an array answer of size m where answer[i] is the sum of unmarked elements in the array after the ith query.

// Example 1:
// Input: nums = [1,2,2,1,2,3,1], queries = [[1,2],[3,3],[4,2]]
// Output: [8,3,0]
// Explanation:
// We do the following queries on the array:
// Mark the element at index 1, and 2 of the smallest unmarked elements with the smallest indices if they exist, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 2 + 2 + 3 + 1 = 8.
// Mark the element at index 3, since it is already marked we skip it. Then we mark 3 of the smallest unmarked elements with the smallest indices, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 3.
// Mark the element at index 4, since it is already marked we skip it. Then we mark 2 of the smallest unmarked elements with the smallest indices if they exist, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 0.

// Example 2:
// Input: nums = [1,4,2,3], queries = [[0,1]]
// Output: [7]
// Explanation: We do one query which is mark the element at index 0 and mark the smallest element among unmarked elements. 
// The marked elements will be nums = [1,4,2,3], and the sum of unmarked elements is 4 + 3 = 7.

// Constraints:
//     n == nums.length
//     m == queries.length
//     1 <= m <= n <= 10^5
//     1 <= nums[i] <= 10^5
//     queries[i].length == 2
//     0 <= indexi, ki <= n - 1

import "fmt"
import "sort"

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
    type Pair struct {
        id, value int
    }
    n, sum := len(nums), int64(0)
    pairs := make([]Pair, n)
    for id, v := range nums { // Create copy of nums and calculate total sum
        sum += int64(v)
        pairs[id] = Pair{ id, v }
    }
    sort.Slice(pairs, func(x, y int) bool { // Sort the new array
        if pairs[x].value == pairs[y].value { return pairs[x].id < pairs[y].id }
        return pairs[x].value < pairs[y].value
    })
    res, marked := make([]int64, len(queries)), make([]bool, n)
    cur := 0
    for i, query := range queries { // Iterator of new array for smallest value
        id, k := query[0], query[1]
        if !marked[id] { // Mark if target index isn't marked.
            marked[id] = true
            sum -= int64(nums[id])
        }
        for ; k > 0 && cur < n; cur++ { // Get k smallest values
            pair := pairs[cur]
            if id, v := pair.id, pair.value; !marked[id] {
                marked[id] = true
                sum -= int64(v)
                k--
            }
        }
        res[i] = sum
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1,2,3,1], queries = [[1,2],[3,3],[4,2]]
    // Output: [8,3,0]
    // Explanation:
    // We do the following queries on the array:
    // Mark the element at index 1, and 2 of the smallest unmarked elements with the smallest indices if they exist, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 2 + 2 + 3 + 1 = 8.
    // Mark the element at index 3, since it is already marked we skip it. Then we mark 3 of the smallest unmarked elements with the smallest indices, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 3.
    // Mark the element at index 4, since it is already marked we skip it. Then we mark 2 of the smallest unmarked elements with the smallest indices if they exist, the marked elements now are nums = [1,2,2,1,2,3,1]. The sum of unmarked elements is 0.
    fmt.Println(unmarkedSumArray([]int{1,2,2,1,2,3,1}, [][]int{{1,2},{3,3},{4,2}})) // [8,3,0]
    // Example 2:
    // Input: nums = [1,4,2,3], queries = [[0,1]]
    // Output: [7]
    // Explanation: We do one query which is mark the element at index 0 and mark the smallest element among unmarked elements. 
    // The marked elements will be nums = [1,4,2,3], and the sum of unmarked elements is 4 + 3 = 7.
    fmt.Println(unmarkedSumArray([]int{1,4,2,3}, [][]int{{0,1}})) // [7]
}