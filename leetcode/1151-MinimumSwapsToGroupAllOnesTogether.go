package main

// 1151. Minimum Swaps to Group All 1's Together
// Given a binary array data, return the minimum number of swaps required to group all 1’s present in the array together in any place in the array.

// Example 1:
// Input: data = [1,0,1,0,1]
// Output: 1
// Explanation: There are 3 ways to group all 1's together:
// [1,1,1,0,0] using 1 swap.
// [0,1,1,1,0] using 2 swaps.
// [0,0,1,1,1] using 1 swap.
// The minimum is 1.

// Example 2:
// Input: data = [0,0,0,1,0]
// Output: 0
// Explanation: Since there is only one 1 in the array, no swaps are needed.

// Example 3:
// Input: data = [1,0,1,0,1,0,0,1,1,0,1]
// Output: 3
// Explanation: One possible solution that uses 3 swaps is [0,0,0,0,0,1,1,1,1,1,1].

// Constraints:
//     1 <= data.length <= 10^5
//     data[i] is either 0 or 1.

import "fmt"

// 滑动窗口与双指针
func minSwaps(data []int) int {
    sum := 0
    for _, v := range data { sum += v } // 统计1数量
    count, mx, left, right := 0, 0, 0, 0
    for right < len(data) {
        count += data[right] // 通过添加新元素更新 1 的数量
        right++
        if right - left > sum { // 将窗口的长度保持为 sum
            count -= data[left] // 通过移除最老的元素来更新 1 的数量
            left++
        }
        if count > mx { // 记录窗口中 1 的数量的最大值
            mx = count
        }
    }
    return sum - mx
}

// 双端队列 + 滑动窗口
func minSwaps1(data []int) int {
    sum := 0
    for _, v := range data { sum += v } // 统计1数量
    count, mx := 0, 0
    deque := make([]int, 0, sum+1) // 保持双端队列的长度等于 sum
    for _, v := range data {
        deque = append(deque, v) // 总是将新元素添加到双端队列中
        count += v
        if len(deque) > sum { // 当双端队列中有超过 sum 个元素，
            count -= deque[0] // 删除最左边的
            deque = deque[1:]
        }
        if count > mx {
            mx = count
        }
    }
    return sum - mx
}

func minSwaps2(data []int) int {
    ones := 0
    for i := range data {
        if data[i] == 1 {
            ones++
        }
    }
    res, left, right, zeroes := len(data), 0, 0, 0
    for right < ones {
        if data[right] == 0 {
            zeroes++
        }
        right++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res = zeroes
    for right < len(data) {
        if data[left] == 0 {
            zeroes--
        }
        left++
        if data[right] == 0 {
            zeroes++
        }
        res = min(res, zeroes)
        right++
    }
    return res
}

func main() {
    // Example 1:
    // Input: data = [1,0,1,0,1]
    // Output: 1
    // Explanation: There are 3 ways to group all 1's together:
    // [1,1,1,0,0] using 1 swap.
    // [0,1,1,1,0] using 2 swaps.
    // [0,0,1,1,1] using 1 swap.
    // The minimum is 1.
    fmt.Println(minSwaps([]int{1,0,1,0,1})) // 1
    // Example 2:
    // Input: data = [0,0,0,1,0]
    // Output: 0
    // Explanation: Since there is only one 1 in the array, no swaps are needed.
    fmt.Println(minSwaps([]int{0,0,0,1,0})) // 0
    // Example 3:
    // Input: data = [1,0,1,0,1,0,0,1,1,0,1]
    // Output: 3
    // Explanation: One possible solution that uses 3 swaps is [0,0,0,0,0,1,1,1,1,1,1].
    fmt.Println(minSwaps([]int{1,0,1,0,1,0,0,1,1,0,1})) // 3

    fmt.Println(minSwaps([]int{0,0,0,0,0,0,0,0,0,0})) // 0
    fmt.Println(minSwaps([]int{1,1,1,1,1,1,1,1,1,1})) // 0
    fmt.Println(minSwaps([]int{0,0,0,0,0,1,1,1,1,1})) // 0
    fmt.Println(minSwaps([]int{1,1,1,1,1,0,0,0,0,0})) // 0
    fmt.Println(minSwaps([]int{0,1,0,1,0,1,0,1,0,1})) // 2
    fmt.Println(minSwaps([]int{1,0,1,0,1,0,1,0,1,0})) // 2

    fmt.Println(minSwaps1([]int{1,0,1,0,1})) // 1
    fmt.Println(minSwaps1([]int{0,0,0,1,0})) // 0
    fmt.Println(minSwaps1([]int{1,0,1,0,1,0,0,1,1,0,1})) // 3
    fmt.Println(minSwaps1([]int{0,0,0,0,0,0,0,0,0,0})) // 0
    fmt.Println(minSwaps1([]int{1,1,1,1,1,1,1,1,1,1})) // 0
    fmt.Println(minSwaps1([]int{0,0,0,0,0,1,1,1,1,1})) // 0
    fmt.Println(minSwaps1([]int{1,1,1,1,1,0,0,0,0,0})) // 0
    fmt.Println(minSwaps1([]int{0,1,0,1,0,1,0,1,0,1})) // 2
    fmt.Println(minSwaps1([]int{1,0,1,0,1,0,1,0,1,0})) // 2

    fmt.Println(minSwaps2([]int{1,0,1,0,1})) // 1
    fmt.Println(minSwaps2([]int{0,0,0,1,0})) // 0
    fmt.Println(minSwaps2([]int{1,0,1,0,1,0,0,1,1,0,1})) // 3
    fmt.Println(minSwaps2([]int{0,0,0,0,0,0,0,0,0,0})) // 0
    fmt.Println(minSwaps2([]int{1,1,1,1,1,1,1,1,1,1})) // 0
    fmt.Println(minSwaps2([]int{0,0,0,0,0,1,1,1,1,1})) // 0
    fmt.Println(minSwaps2([]int{1,1,1,1,1,0,0,0,0,0})) // 0
    fmt.Println(minSwaps2([]int{0,1,0,1,0,1,0,1,0,1})) // 2
    fmt.Println(minSwaps2([]int{1,0,1,0,1,0,1,0,1,0})) // 2
}