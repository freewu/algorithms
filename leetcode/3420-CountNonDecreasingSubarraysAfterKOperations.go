package main

// 3420. Count Non-Decreasing Subarrays After K Operations
// You are given an array nums of n integers and an integer k.

// For each subarray of nums, you can apply up to k operations on it. 
// In each operation, you increment any element of the subarray by 1.

// Note that each subarray is considered independently, meaning changes made to one subarray do not persist to another.

// Return the number of subarrays that you can make non-decreasing ​​​​​after performing at most k operations.

// An array is said to be non-decreasing if each element is greater than or equal to its previous element, if it exists.

// Example 1:
// Input: nums = [6,3,1,2,4,4], k = 7
// Output: 17
// Explanation:
// Out of all 21 possible subarrays of nums, only the subarrays [6, 3, 1], [6, 3, 1, 2], [6, 3, 1, 2, 4] and [6, 3, 1, 2, 4, 4] cannot be made non-decreasing after applying up to k = 7 operations. Thus, the number of non-decreasing subarrays is 21 - 4 = 17.

// Example 2:
// Input: nums = [6,3,1,3,6], k = 4
// Output: 12
// Explanation:
// The subarray [3, 1, 3, 6] along with all subarrays of nums with three or fewer elements, except [6, 3, 1], can be made non-decreasing after k operations. There are 5 subarrays of a single element, 4 subarrays of two elements, and 2 subarrays of three elements except [6, 3, 1], so there are 1 + 5 + 4 + 2 = 12 subarrays that can be made non-decreasing.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9

import "fmt"

func countNonDecreasingSubarrays(nums []int, k int) int64 {
    type Pair struct { val, count int }
    res, n := int64(1), len(nums)
    if n == 0 { return 0 }
    queue := []Pair{{ val: nums[n - 1], count: 1 }}
    cost, back := 0, n - 1
    for i := n - 2; i >= 0; i-- {
        cur, count := nums[i], 1
        // Merge elements in the queue that are smaller than current
        for len(queue) > 0 && queue[len(queue)-1].val < cur {
            tail := queue[len(queue)-1]
            queue = queue[:len(queue)-1]
            count += tail.count
            cost += (cur - tail.val) * tail.count
        }
        queue = append(queue, Pair{ val: cur, count: count })
        // Adjust back pointer to keep cost <= k
        for cost > k && back >= i {
            if len(queue) == 0 { break }
            head := queue[0]
            queue = queue[1:]
            newCount := head.count - 1
            cost -= (head.val - nums[back])
            back--
            if newCount > 0 {
                // Prepend to the queue
                queue = append([]Pair{{ val: head.val, count: newCount }}, queue...)
            }
        }
        // Calculate the window size
        windowSize := 0
        if back >= i {
            windowSize = back - i + 1
        }
        res += int64(windowSize)
    }
    return res
}

func countNonDecreasingSubarrays1(nums []int, k int) int64 {
    res, n, count, rTree := int64(0), len(nums), 0, 0 // rTree 表示窗口最右边那棵树在 stack 中的下标
    type Pair struct{ val, size int } // 根节点的值, 树的大小
    stack := []Pair{}
    r := n - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for l := n - 1; l >= 0; l-- {
        x := nums[l] // x 进入窗口
        size := 1 // 统计以 x 为根的树的大小
        for len(stack) > 0 && x >= stack[len(stack) - 1].val {
            // 以 p.val 为根的树，现在合并到 x 的下面（x 和 val 连一条边）
            p := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            size += p.size
            count += (x - p.val) * p.size // 树 p.val 中的数都变成 x
        }
        stack = append(stack, Pair{x, size})
        // 如果从 stack  中弹出树包含 rTree，那么 rTree 现在指向栈顶这棵树
        rTree = min(rTree, len(stack) - 1)
        // 当 count 大于 k 时，缩小窗口
        for count > k {
            tree := &stack[rTree]
            count -= tree.val - nums[r] // 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
            r--
            tree.size-- // nums[r] 离开窗口后，树的大小减一
            if tree.size == 0 { // 这棵树是空的
                rTree++ // rTree 指向左边下一棵树
            }
        }
        res += int64(r - l + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [6,3,1,2,4,4], k = 7
    // Output: 17
    // Explanation:
    // Out of all 21 possible subarrays of nums, only the subarrays [6, 3, 1], [6, 3, 1, 2], [6, 3, 1, 2, 4] and [6, 3, 1, 2, 4, 4] cannot be made non-decreasing after applying up to k = 7 operations. Thus, the number of non-decreasing subarrays is 21 - 4 = 17.
    fmt.Println(countNonDecreasingSubarrays([]int{6,3,1,2,4,4}, 7)) // 17
    // Example 2:
    // Input: nums = [6,3,1,3,6], k = 4
    // Output: 12
    // Explanation:
    // The subarray [3, 1, 3, 6] along with all subarrays of nums with three or fewer elements, except [6, 3, 1], can be made non-decreasing after k operations. There are 5 subarrays of a single element, 4 subarrays of two elements, and 2 subarrays of three elements except [6, 3, 1], so there are 1 + 5 + 4 + 2 = 12 subarrays that can be made non-decreasing.
    fmt.Println(countNonDecreasingSubarrays([]int{6,3,1,3,6}, 4)) // 12

    fmt.Println(countNonDecreasingSubarrays([]int{1,2,3,4,5,6,7,8,9}, 4)) // 45
    fmt.Println(countNonDecreasingSubarrays([]int{9,8,7,6,5,4,3,2,1}, 4)) // 24

    fmt.Println(countNonDecreasingSubarrays1([]int{6,3,1,2,4,4}, 7)) // 17
    fmt.Println(countNonDecreasingSubarrays1([]int{6,3,1,3,6}, 4)) // 12
    fmt.Println(countNonDecreasingSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 4)) // 45
    fmt.Println(countNonDecreasingSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 4)) // 24
}

