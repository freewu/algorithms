package main

import "fmt"

/**
60 · Search Insert Position
Description
Given a sorted array and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.
You may assume NO duplicates in the array.

Example 1:

	Input:
		array = [1,3,5,6]
		target = 5
	Output: 
		2
	Explanation: 5 is indexed to 2 in the array.

Example 2:

	Input:
		array = [1,3,5,6]
		target = 2

	Output:
		1

	Explanation: 2 should be inserted into the position with index 1.

Example 3:

	Input:
		array = [1,3,5,6]
		target = 7

	Output:
		4

	Explanation:
		7 should be inserted into the position with index 4.

Example 4:

	Input:
		array = [1,3,5,6]
		target = 0

	Output:
		0

	Explanation:

		0 should be inserted into the position with index 0.

# Challenge
	O(log(n)) time
*/

/*
class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
                
        low = 0
        high = len(nums) - 1
        
        while low <= high:
            mid = low + (high - low) // 2
            
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                high = mid - 1
            else:
                low = mid + 1
                
        return low
*/

/**
 * @param a: an integer sorted array
 * @param target: an integer to be inserted
 * @return: An integer
 */
// 二分法
func SearchInsert(a []int, target int) int {
    // write your code here
	if len(a) == 0 { // 如果空数组
        return 0
    }
	low,high,mid := 0, len(a) - 1,0
	if target <= a[low] {
		return low
	}
	if target > a[high] {
		return high + 1
	}
	for {
		// 退出条件
		if low > high {
			break
		}
		// 计算中间值 
		mid = low + (high - low) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] > target { // 如果值大于 目标 说明位置在 前半部分
			high = mid - 1
		} else { // 如果值 小于 目标 说明位置在 后半部分
			low = mid + 1 
		}
	}
	return low
}

func main() {
	fmt.Printf("SearchInsert([]int{1,3,5,6},5) = %v\n",SearchInsert([]int{1,3,5,6}, 5)) // 2
	fmt.Printf("SearchInsert([]int{1,3,5,6}, 2) = %v\n",SearchInsert([]int{1,3,5,6}, 2)) // 1
	fmt.Printf("SearchInsert([]int{1,3,5,6,7}, 7) = %v\n",SearchInsert([]int{1,3,5,6,7}, 7)) // 4
	fmt.Printf("SearchInsert([]int{1,3,5,6}, 0)) = %v\n",SearchInsert([]int{1,3,5,6}, 0)) // 0
}

