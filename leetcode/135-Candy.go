package main

import "fmt"

/**
135. Candy
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
You are giving candies to these children subjected to the following requirements:
Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.

Constraints:

	n == ratings.length
	1 <= n <= 2 * 10^4
	0 <= ratings[i] <= 2 * 10^4

Example 1:

	Input: ratings = [1,0,2]
	Output: 5
	Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:

	Input: ratings = [1,2,2]
	Output: 4
	Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
	The third child gets 1 candy because it satisfies the above two conditions.

老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。你需要按照以下要求，帮助老师给这些孩子分发糖果：

	每个孩子至少分配到 1 个糖果。
	评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。

那么这样下来，老师至少需要准备多少颗糖果呢？

 */

func candy(ratings []int) int {
	candies := make([]int, len(ratings)) // 声明一个数组来保存发糖数
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] { // 如果当前孩子评分 大于 前面一个孩子评分  当前评分 + 1
			candies[i] += candies[i-1] + 1 // 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果
		}
	}
	fmt.Printf("after round 1 %v\n",candies)
	for i := len(ratings) - 2; i >= 0; i-- { // 从后向前推进
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1 // 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果
		}
	}
	fmt.Printf("after round 2 %v\n",candies)
	total := 0
	for _, candy := range candies {
		total += candy + 1 // 每个孩子至少分配到 1 个糖果
	}
	return total
}

func main() {
	fmt.Printf("candy([]int{ 1,0,2 }) = %v\n",candy([]int{ 1,0,2 })) // 5
	fmt.Printf("candy([]int{ 1,2,2 }) = %v\n",candy([]int{ 1,2,2 })) // 4
}
