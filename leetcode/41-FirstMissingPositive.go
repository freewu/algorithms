package main

// 41. First Missing Positive
// Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
// You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

// Example 1:
// Input: nums = [1,2,0]
// Output: 3
// Explanation: The numbers in the range [1,2] are all in the array.

// Example 2:
// Input: nums = [3,4,-1,1]
// Output: 2
// Explanation: 1 is in the array but 2 is missing.

// Example 3:
// Input: nums = [7,8,9,11,12]
// Output: 1
// Explanation: The smallest positive integer 1 is missing.

// Constraints:
//     1 <= nums.length <= 10^5
//     -2^31 <= nums[i] <= 2^31 - 1

// 解题思路:
//     数组都装到 map 中
//     然后 i 循环从 1 开始，依次比对 map 中是否存在 i，
//     只要不存在 i 就立即返回结果 i


import "fmt"

func firstMissingPositive(nums []int) int {
    numMap := make(map[int]int, len(nums))
    for _, v := range nums { // 先把数组保存到 map 中
        numMap[v] = v
    }
    // 然后 i 循环从 1 开始，依次比对 map 中是否存在 i，
    for index := 1; index < len(nums) + 1; index++ {
        if _, ok := numMap[index]; !ok { // 如果数组中不存在 说明 index 是需要返回的值
            return index
        }
    }
    return len(nums) + 1
}

// 思路一样就是做了 非合理值的判断
func firstMissingPositive1(nums []int) int {
    founds := make([]bool, 5*100000)
    for _, num := range nums {
        if num > 0 && num <= 5*100000 {
            founds[num-1] = true
        }
    }
    for pos, found := range founds {
        if !found {
            return pos+1
        }
    }
    return 5*100000 + 1
}

func firstMissingPositive2(nums []int) int {
    // 原地hash
    //考虑将5 放到数组中下标为4的元素上
    l := len(nums)
    for i := 0; i < l; i++ {
        // 元素符合预期准备交换
        for nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    for i := 0; i < l; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }
    return l + 1
}

func main() {
    fmt.Printf("firstMissingPositive([]int{1,2,0} = %v\n",firstMissingPositive([]int{1,2,0})) // 3
    fmt.Printf("firstMissingPositive([]int{3,4,-1,1} = %v\n",firstMissingPositive([]int{3,4,-1,1})) // 2
    fmt.Printf("firstMissingPositive([]int{7,8,9,11,12} = %v\n",firstMissingPositive([]int{7,8,9,11,12})) // 1
    fmt.Printf("firstMissingPositive1([]int{1,2,0} = %v\n",firstMissingPositive1([]int{1,2,0})) // 3
    fmt.Printf("firstMissingPositive1([]int{3,4,-1,1} = %v\n",firstMissingPositive1([]int{3,4,-1,1})) // 2
    fmt.Printf("firstMissingPositive1([]int{7,8,9,11,12} = %v\n",firstMissingPositive1([]int{7,8,9,11,12})) // 1
}