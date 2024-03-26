package main

// 31. Next Permutation
// A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
//     For example, for arr = [1,2,3], the following are all the permutations of arr: 
//         [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].

// The next permutation of an array of integers is the next lexicographically greater permutation of its integer. 
// More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, 
// then the next permutation of that array is the permutation that follows it in the sorted container. 
// If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

// For example, the next permutation of arr = [1,2,3] is [1,3,2].
// Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
// While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
// Given an array of integers nums, find the next permutation of nums.

// The replacement must be in place and use only constant extra memory.

// Example 1:
// Input: nums = [1,2,3]
// Output: [1,3,2]

// Example 2:
// Input: nums = [3,2,1]
// Output: [1,2,3]

// Example 3:
// Input: nums = [1,1,5]
// Output: [1,5,1]
 
// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 100

import "fmt"

func nextPermutation(nums []int) {
    swap := func (nums *[]int, i, j int) {
        (*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
    }
    reverse := func (nums *[]int, i, j int) {
        for i < j {
            swap(nums, i, j)
            i++
            j--
        }
    }
    i, j := 0, 0
    // 在 nums[i] 中找到 i 使得 nums[i] < nums[i+1]，此时较小数为 nums[i]，并且 [i+1, n) 一定为下降区间
    for i = len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            break
        }
    }
    // 如果找到了这样的 i ，则在下降区间 [i+1, n) 中从后往前找到第一个 j ，使得 nums[i] < nums[j] ，此时较大数为 nums[j]
    if i >= 0 {
        for j = len(nums) - 1; j > i; j-- {
            if nums[j] > nums[i] {
                break
            }
        }
        // 交换 nums[i] 和 nums[j]，此时区间 [i+1, n) 一定为降序区间。最后原地交换 [i+1, n) 区间内的元素，使其变为升序，无需对该区间进行排序
        swap(&nums, i, j)
    }
    reverse(&nums, i+1, len(nums)-1)
}


func nextPermutation1(nums []int) {
    n := len(nums)
    i := n - 2
    for ; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            break
        }
    }
    if i >= 0 {
        j := n - 1
        for ; j > i; j-- {
            if nums[j] > nums[i] {
                nums[j], nums[i] = nums[i], nums[j]
                break
            }

        }
    }
    reverse := func (nums []int){
        for i := 0; i < len(nums) / 2; i++{
            nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1],nums[i]
        }
    }
    reverse(nums[i+1:])
}

func main() {
    nums1 := []int{1,2,3}
    fmt.Printf("before nums1 = %v \n",nums1)
    nextPermutation(nums1)
    fmt.Printf("after nums1 = %v \n",nums1)

    nums2 := []int{3,2,1}
    fmt.Printf("before nums2 = %v \n",nums2)
    nextPermutation(nums2)
    fmt.Printf("after nums2 = %v \n",nums2)

    nums11 := []int{1,2,3}
    fmt.Printf("before nums11 = %v \n",nums11)
    nextPermutation1(nums11)
    fmt.Printf("after nums11 = %v \n",nums11)

    nums21 := []int{3,2,1}
    fmt.Printf("before nums21 = %v \n",nums21)
    nextPermutation1(nums21)
    fmt.Printf("after nums21 = %v \n",nums21)
}
