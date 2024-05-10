package main

// 912. Sort an Array
// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

// Example 1:
// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).

// Example 2:
// Input: nums = [5,1,1,2,0,0]
// Output: [0,0,1,1,2,5]
// Explanation: Note that the values of nums are not necessairly unique.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     -5 * 10^4 <= nums[i] <= 5 * 10^4

import "fmt"
import "math"

func sortArray(nums []int) []int {
    return radixSort(nums)
}

func radixSort(nums []int) []int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    digitCount := func (num int) int {
        if num == 0 {
            return 1
        } 
        // floor(log10(num)) + 1
        // 234 => floor(log10(234)) = 2 + 1
        return int(math.Floor(math.Log10(float64(abs(num))))) + 1;
    }
    reverse := func (nums []int) {
        length := len(nums)
        for i := 0; i < length / 2; i++ {
            nums[i], nums[length - i - 1] = nums[length - i - 1], nums[i]
        }
    }
    getFromBuckets := func(digitBuckets [][]int) []int {
        nums := make([]int, 0)
        for i := range digitBuckets {
            nums = append(nums, digitBuckets[i]...)
        }
        return nums
    }
    getDigit := func (num, place int) int {
        // (num / 10^place) % 10
        // 234, position 1 => 234 / 10^1 => 23 % 10 => 3
        return int(math.Floor(float64(abs(num)) / math.Pow(10, float64(place)))) % 10;
    }
    maxNumber := func (nums []int) int {
        res := math.MinInt64
        for _, num := range nums {
            if res < abs(num) { // !number can be negative
                res = abs(num)
            }
        }
        return res
    }
    
    maxDigitCount := digitCount(maxNumber(nums)) // how many digits has the longest number
    for k := 0; k < maxDigitCount; k++ { // put into the buckets according to the digit at k-th position
        digitBuckets := make([][]int, 10)
        for i := 0; i < len(nums); i++ {
            digitAtKposition := getDigit(nums[i], k)
            digitBuckets[digitAtKposition] = append(digitBuckets[digitAtKposition], nums[i])
        }
        // get from the buckets
        nums = getFromBuckets(digitBuckets)
    }
    signBuckets := make([][]int, 2) // sort negative numbers, we can treat a sign as another bucket
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            signBuckets[0] = append(signBuckets[0], nums[i])
        } else {
            signBuckets[1] = append(signBuckets[1], nums[i])
        }
    }
    reverse(signBuckets[0]) // reverse array with negative numbers
    nums = getFromBuckets(signBuckets)
    return nums
}

func sortArray1(nums []int) []int {
    if len(nums) == 0{
        return nums
    }
    mn, mx := nums[0], nums[0]
    for _, v := range nums { // 找出最大值 & 最小值
        if v < mn { mn = v; }
        if v > mx { mx = v; }
    }
    count := make([]int, mx - mn + 1) // 创建  mx - mn + 1 个桶
    for _, num := range nums{
        count[num - mn]++
    }
    res, index := make([]int, len(nums)), 0
    for i, cnt := range count {
        for j := 0; j < cnt; j++ { // 出现几个就循环几个出来
            res[index] = i + mn
            index++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,2,3,1]
    // Output: [1,2,3,5]
    // Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
    fmt.Println(sortArray([]int{5,2,3,1})) // [1,2,3,5]
    // Example 2:
    // Input: nums = [5,1,1,2,0,0]
    // Output: [0,0,1,1,2,5]
    // Explanation: Note that the values of nums are not necessairly unique.
    fmt.Println(sortArray([]int{5,1,1,2,0,0})) // [0,0,1,1,2,5]

    fmt.Println(sortArray1([]int{5,2,3,1})) // [1,2,3,5]
    fmt.Println(sortArray1([]int{5,1,1,2,0,0})) // [0,0,1,1,2,5]
}