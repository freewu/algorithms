package main

// 239. Sliding Window Maximum
// You are given an array of integers nums, 
// there is a sliding window of size k which is moving from the very left of the array to the very right. 
// You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the max sliding window.

// Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation: 
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4
//     1 <= k <= nums.length

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。返回滑动窗口最大值。
// 最优的解法是用双端队列，队列的一边永远都存的是窗口的最大值，
// 队列的另外一个边存的是比最大值小的值。队列中最大值左边的所有值都出队。
// 在保证了双端队列的一边即是最大值以后，时间复杂度是 O(n)，空间复杂度是 O(K)

import "fmt"

// 暴力解法 O(n^2)
func maxSlidingWindow1(arr []int, k int) []int {
    res, n := make([]int, 0, k), len(arr)
    if n == 0 {
        return []int{}
    }
    for i := 0; i <= n - k; i++ {
        mx := arr[i]
        for j := 1; j < k; j++ {
            if mx < arr[i+j] {
                mx = arr[i+j]
            }
        }
        res = append(res, mx)
    }
    return res
}

//  双端队列 Deque 时间复杂度是 O(n)，空间复杂度是 O(K)
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || len(nums) < k {
        return make([]int, 0)
    }
    window := make([]int, 0, k) // store the index of nums
    res := make([]int, 0, len(nums)-k+1)
    for i, v := range nums { // if the left-most index is out of window, remove it
        if i >= k && window[0] <= i-k {
            window = window[1:len(window)]
        }
        for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
            window = window[0 : len(window)-1]
        }
        window = append(window, i) // store the index of nums
        if i >= k-1 {
            res = append(res, nums[window[0]]) // the left-most is the index of max value in nums
        }
    }
    return res
}

func maxSlidingWindow2(nums []int, limit int) []int {
    res, stack := []int{}, []int{}
    if len(nums) == 0 {
        return res
    }
    check := func(n int) {
        for len(stack) != 0 && stack[len(stack)-1] < n {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, n)
    }
    for i := 0; i < limit; i++ {
        check(nums[i])
    }
    res = append(res, stack[0])
    for i := limit; i < len(nums); i++ {
        check(nums[i])
        val := stack[0]
        if nums[i-limit] == val {
            stack = stack[1:]
        }
        res = append(res, stack[0])
    }
    return res
}


func main()  {
    fmt.Printf("maxSlidingWindow1([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxSlidingWindow1([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxSlidingWindow1([]int{1},1) = %v\n",maxSlidingWindow1([]int{1},1)) // [1]

    fmt.Printf("maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxSlidingWindow([]int{1},1) = %v\n",maxSlidingWindow([]int{1},1)) // [1]

    fmt.Printf("maxSlidingWindow2([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxSlidingWindow2([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxSlidingWindow2([]int{1},1) = %v\n",maxSlidingWindow2([]int{1},1)) // [1]
}