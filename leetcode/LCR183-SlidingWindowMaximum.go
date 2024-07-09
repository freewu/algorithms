package main

// LCR 183. 望远镜中最高的海拔
// 科技馆内有一台虚拟观景望远镜，它可以用来观测特定纬度地区的地形情况。
// 该纬度的海拔数据记于数组 heights ，其中 heights[i] 表示对应位置的海拔高度。请找出并返回望远镜视野范围 limit 内，可以观测到的最高海拔值。

// 示例 1：
// 输入：heights = [14,2,27,-5,28,13,39], limit = 3
// 输出：[27,27,28,28,39]
// 解释：
//   滑动窗口的位置                最大值
// ---------------               -----
// [14 2 27] -5 28 13 39          27
// 14 [2 27 -5] 28 13 39          27
// 14 2 [27 -5 28] 13 39          28
// 14 2 27 [-5 28 13] 39          28
// 14 2 27 -5 [28 13 39]          39

// 提示：
//     你可以假设输入总是有效的，在输入数组不为空的情况下：
//     1 <= limit <= heights.length
//     -10000 <= heights[i] <= 10000


// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。返回滑动窗口最大值。
// 最优的解法是用双端队列，队列的一边永远都存的是窗口的最大值，
// 队列的另外一个边存的是比最大值小的值。队列中最大值左边的所有值都出队。
// 在保证了双端队列的一边即是最大值以后，时间复杂度是 O(n)，空间复杂度是 O(K)

import "fmt"

// 暴力解法 O(n^2)
func maxAltitude(heights []int, k int) []int {
    res, n := make([]int, 0, k), len(heights)
    if n == 0 {
        return []int{}
    }
    for i := 0; i <= n - k; i++ {
        mx := heights[i]
        for j := 1; j < k; j++ {
            if mx < heights[i+j] {
                mx = heights[i+j]
            }
        }
        res = append(res, mx)
    }
    return res
}

// 双端队列 Deque 时间复杂度是 O(n)，空间复杂度是 O(K)
func maxAltitude1(heights []int, k int) []int {
    n := len(heights)
    if n == 0 || n < k {
        return make([]int, 0)
    }
    window := make([]int, 0, k) // store the index of nums
    res := make([]int, 0, n - k + 1)
    for i, v := range heights { // if the left-most index is out of window, remove it
        if i >= k && window[0] <= i-k {
            window = window[1:len(window)]
        }
        for len(window) > 0 && heights[window[len(window)-1]] < v { // maintain window
            window = window[0 : len(window)-1]
        }
        window = append(window, i) // store the index of nums
        if i >= k-1 {
            res = append(res, heights[window[0]]) // the left-most is the index of max value in nums
        }
    }
    return res
}

func maxAltitude2(heights []int, limit int) []int {
    res, stack := []int{}, []int{}
    if len(heights) == 0 {
        return res
    }
    check := func(n int) {
        for len(stack) != 0 && stack[len(stack)-1] < n {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, n)
    }
    for i := 0; i < limit; i++ {
        check(heights[i])
    }
    res = append(res, stack[0])
    for i := limit; i < len(heights); i++ {
        check(heights[i])
        val := stack[0]
        if heights[i-limit] == val {
            stack = stack[1:]
        }
        res = append(res, stack[0])
    }
    return res
}

func main()  {
    fmt.Printf("maxAltitude([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxAltitude([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxAltitude([]int{1},1) = %v\n",maxAltitude([]int{1},1)) // [1]

    fmt.Printf("maxAltitude1([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxAltitude1([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxAltitude1([]int{1},1) = %v\n",maxAltitude1([]int{1},1)) // [1]

    fmt.Printf("maxAltitude2([]int{1,3,-1,-3,5,3,6,7},3) = %v\n",maxAltitude2([]int{1,3,-1,-3,5,3,6,7},3)) // [3,3,5,5,6,7]
    fmt.Printf("maxAltitude2([]int{1},1) = %v\n",maxAltitude2([]int{1},1)) // [1]
}