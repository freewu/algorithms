package main

import "fmt"

/**
275. H-Index II
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper and citations is sorted in an ascending order,
return compute the researcher's h-index.
According to the definition of h-index on Wikipedia:
A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations each.
If there are several possible values for h, the maximum one is taken as the h-index.
You must write an algorithm that runs in logarithmic time.

Example 1:

	Input: citations = [0,1,3,5,6]
	Output: 3
	Explanation:
		[0,1,3,5,6] means the researcher has 5 papers in total and each of them had received 0, 1, 3, 5, 6 citations respectively.
		Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:

	Input: citations = [1,2,100]
	Output: 2

Constraints:

	n == citations.length
	1 <= n <= 105
	0 <= citations[i] <= 1000
	citations is sorted in ascending order.

给出一个数组，代表该作者论文被引用次数，要求这个作者的 h 指数。h 指数定义：“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）
 */

func hIndex(citations []int) int {
	low, high := 0, len(citations)-1
	for low <= high {
		mid := low + (high-low) >> 1
		if len(citations) - mid > citations[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return len(citations) - low
}

func main() {
	fmt.Printf("hIndex([]int{0,1,3,5,6}) = %v\n",hIndex([]int{0,1,3,5,6})) // 3
	fmt.Printf("hIndex([]int{1,2,100}) = %v\n",hIndex([]int{1,2,100})) // 2
}