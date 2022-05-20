package main

/**
14 · First Position of Target
Description
Given a sorted array (ascending order) and a target number, find the first index of this number in O(log n)O(logn) time complexity.
If the target number does not exist in the array, return -1.


Example 1:

	Input:
		tuple = [1,4,4,5,7,7,8,9,9,10]
		target = 1
	Output:
		0
	Explanation: The first index of 1 is 0.

Example 2:

	Input
		tuple = [1, 2, 3, 3, 4, 5, 10]
		target = 3
	Output:
		2
	Explanation: The first index of 3 is 2.

Example 3:

	Input:

		tuple = [1, 2, 3, 3, 4, 5, 10]
		target = 6
	Output:
		-1
	Explanation:

	There is no 6 in the array，return -1.

Challenge
	If the count of numbers is bigger than 2^32, can your code work properly?
*/

import "fmt"

/**
 * @param nums: The integer array.
 * @param target: Target to find.
 * @return: The first position of target. Position starts from 0.
 */
func BinarySearch(nums []int, target int) int {
	// write your code here
	low, high := 0,len(nums) -1
	for low < high {
		mid := (low + high) / 2
		fmt.Printf("low = %v,high = %v,mid = %v\n",low, high, mid)
		if nums[mid] < target { // 在右半边
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[low] == target { // 找到了
		return low
	}
	return -1
}

func main() {
	fmt.Printf("BinarySearch([]int{1,4,4,5,7,7,8,9,9,10},1) = %v\n",BinarySearch([]int{1,4,4,5,7,7,8,9,9,10},1)) // 0
	fmt.Printf("BinarySearch([]int{1, 2, 3, 3, 4, 5, 10},3) = %v\n",BinarySearch([]int{1, 2, 3, 3, 4, 5, 10},3)) // 2
	fmt.Printf("BinarySearch([]int{1, 2, 3, 3, 4, 5, 10},6) = %v\n",BinarySearch([]int{1, 2, 3, 3, 4, 5, 10},6)) // -1
	fmt.Printf("BinarySearch([]int{2,6,8,13,15,17,17,18,19,20},15) = %v\n",BinarySearch([]int{2,6,8,13,15,17,17,18,19,20},15)) // 4
}