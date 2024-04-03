package main

// 275. H-Index II
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper and citations is sorted in ascending order, return the researcher's h-index.
// According to the definition of h-index on Wikipedia: 
//     The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.

// You must write an algorithm that runs in logarithmic time.

// Example 1:
// Input: citations = [0,1,3,5,6]
// Output: 3
// Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had received 0, 1, 3, 5, 6 citations respectively.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

// Example 2:
// Input: citations = [1,2,100]
// Output: 2

// Constraints:
//     n == citations.length
//     1 <= n <= 10^5
//     0 <= citations[i] <= 1000
//     citations is sorted in ascending order.

// 给出一个数组，代表该作者论文被引用次数，要求这个作者的 h 指数。h 指数定义：“高引用次数”（high citations），
// 一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）

import "fmt"

func hIndex(citations []int) int {
    low, high := 0, len(citations)-1
    for low <= high {
        mid := low + (high-low) >> 1
        // 当 len(citations)-mid > citations[mid] 时，说明 h 指数的边界一定在右边，
        // 因为最多 len(citations)-mid 篇数比引用数 citations[mid] 还要大
        if len(citations) - mid > citations[mid] {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    // 最终求的是 h 指数，用 len(citations) - low 即是结果
    return len(citations) - low
}

func hIndex1(citations []int) int {
    left, right := 1, len(citations) + 1
    var mid int
    for left < right {
        mid = left + (right - left) >> 1
        if citations[len(citations) - mid] >= mid {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return right - 1
}

func main() {
    fmt.Printf("hIndex([]int{0,1,3,5,6}) = %v\n",hIndex([]int{0,1,3,5,6})) // 3
    fmt.Printf("hIndex([]int{1,2,100}) = %v\n",hIndex([]int{1,2,100})) // 2

    fmt.Printf("hIndex1([]int{0,1,3,5,6}) = %v\n",hIndex1([]int{0,1,3,5,6})) // 3
    fmt.Printf("hIndex1([]int{1,2,100}) = %v\n",hIndex1([]int{1,2,100})) // 2
}