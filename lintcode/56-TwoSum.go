package main

/**
56 · Two Sum
Description
Given an array of integers, find two numbers such that they add up to a specific target number.
The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. 
Please note that your returned answers (both index1 and index2) are zero-based.

You may assume that each input would have exactly one solution

# Example 1:

	Input:
		numbers = [2,7,11,15]
		target = 9

	Output:
		[0,1]

	Explanation:
		numbers[0] + numbers[1] = 9

# Example 2:

	Input:
		numbers = [15,2,7,11]
		target = 9

	Output:
		[1,2]

	Explanation:

		numbers[1] + numbers[2] = 9

# Challenge
	Either of the following solutions are acceptable:
		O(n) Space, O(nlogn) Time
		O(n) Space, O(n) Time
*/

/**
 * @param numbers: An array of Integer
 * @param target: target = numbers[index1] + numbers[index2]
 * @return: [index1, index2] (index1 < index2)
 */
 func TwoSum(numbers []int, target int) []int {
    // write your code here
    m := make(map[int]int, len(numbers)) // 优先设定好map长度,避免扩容产生性能波动
	for i, num := range numbers {
		if idx, ok := m[target - num]; ok { // 少个中间变量
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}
}