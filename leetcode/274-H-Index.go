package main

import "fmt"

/**
274. H-Index

Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper,
return compute the researcher's h-index.
According to the definition of h-index on Wikipedia:
A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations each.
If there are several possible values for h, the maximum one is taken as the h-index.

Example 1:

	Input: citations = [3,0,6,1,5]
	Output: 3
	Explanation:
		[3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
		Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:

	Input: citations = [1,3,1]
	Output: 1

Constraints:

	n == citations.length
	1 <= n <= 5000
	0 <= citations[i] <= 1000

h-index 值的定义：如果他/她的 N 篇论文中至少有 h 引用，而其他 N-h 论文的引用数不超过 h 引用数。
可以先将数组里面的数从小到大排序。
因为要找最大的 h-index，所以从数组末尾开始往前找，找到第一个数组的值，小于，总长度减去下标的值，这个值就是 h-index。
 */

func hIndex(citations []int) int {
	n := len(citations) // 数组长度
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++ // 如果值大取数组长度 加在尾部
		} else {
			buckets[c]++
		}
		fmt.Printf("c = %v, buckets = %#v\n",c,buckets)
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i { // 找到  h 引用数
			return i
		}
	}
	return 0
}

func main() {
	fmt.Printf("hIndex([]int{3,0,6,1,5}) = %v\n",hIndex([]int{3,0,6,1,5})) // 3
	fmt.Printf("hIndex([]int{1,3,1}) = %v\n",hIndex([]int{1,3,1})) // 1
}
