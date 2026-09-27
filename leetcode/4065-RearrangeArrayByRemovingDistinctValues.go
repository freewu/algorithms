package main

// 4065. Rearrange Array by Removing Distinct Values
// You are given an integer array nums.

// You start with an empty array ans. Repeat the following operation until nums is empty:
//     1. Identify all distinct values currently present in nums.
//     2. Remove one occurrence of every distinct value currently in nums, and append those values to ans in ascending order.

// Return the array ans.

// Example 1:
// Input: nums = [3,1,3,2,1,3]
// Output: [1,2,3,1,3,3]
// Explanation:
// Operation   | Appended to ans | nums  after | ans after
// 1           | 1, 2, 3         | [3, 1, 3]   | [1, 2, 3]
// 2           | 1, 3            | [3]         | [1, 2, 3, 1, 3]
// 3           | []              | []      | [1, 2, 3, 1, 3, 3]
// nums is now empty, so the answer is [1, 2, 3, 1, 3, 3].

// Example 2:
// Input: nums = [7,7,4,4,4]
// Output: [4,7,4,7,4]
// Explanation:
// Operation   | Appended to ans | nums  after | ans after
// 1           | 4, 7            | [7, 4, 4]   | [4, 7]
// 2           | 4, 7            | [4]         | [4, 7, 4, 7]
// 3           | []              | []          | [4, 7, 4]
// nums is now empty, so the answer is [4, 7, 4, 7, 4].

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"
import "sort"

func rearrangeArray(nums []int) []int {
    freq := make(map[int]int)
    type Pair struct{ round, val int }
    tagged := make([]Pair, len(nums))
    for i, x := range nums {
        freq[x]++
        tagged[i] = Pair{freq[x], x}
    }
    sort.Slice(tagged, func(i, j int) bool {
        if tagged[i].round != tagged[j].round {
            return tagged[i].round < tagged[j].round
        }
        return tagged[i].val < tagged[j].val
    })
    res := make([]int, len(nums))
    for i, p := range tagged {
        res[i] = p.val
    }
    return res
}

func rearrangeArray1(nums []int) []int {
    res, freq := []int{}, make([]int, 101)
    for _, v := range nums {
        freq[v]++
    }
    for {
        seenValue := false
        for i, count := range freq {
            if count > 0 {
                seenValue = true
                res = append(res, i)
                freq[i]--
            }
        }
        if !seenValue {
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,3,2,1,3]
    // Output: [1,2,3,1,3,3]
    // Explanation:
    // Operation   | Appended to ans | nums  after | ans after
    // 1           | 1, 2, 3         | [3, 1, 3]   | [1, 2, 3]
    // 2           | 1, 3            | [3]         | [1, 2, 3, 1, 3]
    // 3           | []              | []      | [1, 2, 3, 1, 3, 3]
    // nums is now empty, so the answer is [1, 2, 3, 1, 3, 3].
    fmt.Println(rearrangeArray([]int{3,1,3,2,1,3})) // [1,2,3,1,3,3]
    // Example 2:
    // Input: nums = [7,7,4,4,4]
    // Output: [4,7,4,7,4]
    // Explanation:
    // Operation   | Appended to ans | nums  after | ans after
    // 1           | 4, 7            | [7, 4, 4]   | [4, 7]
    // 2           | 4, 7            | [4]         | [4, 7, 4, 7]
    // 3           | []              | []          | [4, 7, 4]
    // nums is now empty, so the answer is [4, 7, 4, 7, 4]. 
    fmt.Println(rearrangeArray([]int{7,7,4,4,4})) // [4,7,4,7,4]

    fmt.Println(rearrangeArray([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(rearrangeArray([]int{9,8,7,6,5,4,3,2,1})) // [1 2 3 4 5 6 7 8 9]

    fmt.Println(rearrangeArray1([]int{3,1,3,2,1,3})) // [1,2,3,1,3,3]
    fmt.Println(rearrangeArray1([]int{7,7,4,4,4})) // [4,7,4,7,4]
    fmt.Println(rearrangeArray1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(rearrangeArray1([]int{9,8,7,6,5,4,3,2,1})) // [1 2 3 4 5 6 7 8 9]
}
