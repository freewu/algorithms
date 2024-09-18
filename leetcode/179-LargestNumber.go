package main

// 179. Largest Number
// Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
// Since the result may be very large, so you need to return a string instead of an integer.

// Example 1:
// Input: nums = [10,2]
// Output: "210"

// Example 2:
// Input: nums = [3,30,34,5,9]
// Output: "9534330"
 
// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 10^9

import "fmt"
import "bytes"
import "math"
import "sort"
import "strconv"
import "strings"

func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        v1, v2 := float64(nums[i]), float64(nums[j])
        if v1 == v2 || v1 * v2 == 0 {
            return v2 < v1
        }
        lg1, lg2 := int(math.Log10(v1)), int(math.Log10(v2))
        return v2 * math.Pow10(lg1+1) + v1 < v1 * math.Pow10(lg2+1) + v2
    })
	// fmt.Println(nums)
    var buf bytes.Buffer
    for _, num := range nums {
        if num == 0 && buf.Len() > 0 && buf.Bytes()[0] == '0' {
            continue
        }
        buf.WriteString(strconv.Itoa(num))
    }
    return buf.String()
}

func largestNumber1(nums []int) string {
    if len(nums) == 0 {
        return ""
    }
    toStringArray := func(nums []int) []string {
        strs := make([]string, 0)
        for _, num := range nums {
            strs = append(strs, strconv.Itoa(num))
        }
        return strs
    }
    partitionString := func(a []string, lo, hi int) int {
        pivot := a[hi]
        i := lo - 1
        for j := lo; j < hi; j++ {
            ajStr := a[j] + pivot
            pivotStr := pivot + a[j]
            if ajStr > pivotStr { // 这里的判断条件是关键
                i++
                a[j], a[i] = a[i], a[j]
            }
        }
        a[i+1], a[hi] = a[hi], a[i+1]
        return i + 1
    }
    var quickSortString func(a []string, lo, hi int) 
    quickSortString = func(a []string, lo, hi int) {
        if lo >= hi {
            return
        }
        p := partitionString(a, lo, hi)
        quickSortString(a, lo, p-1)
        quickSortString(a, p+1, hi)
    }
    numStrs := toStringArray(nums)
    quickSortString(numStrs, 0, len(numStrs)-1)
    res := ""
    for _, str := range numStrs {
        if res == "0" && str == "0" {
            continue
        }
        res = res + str
    }
    return res
}

func largestNumber2(nums []int) string {
    arr := make([]string, len(nums))
    for i, num := range nums {
        arr[i] = strconv.Itoa(num)
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] + arr[j] > arr[j] + arr[i]
    })
    res := strings.Join(arr, "")
    if res[0] == '0' {
        return "0"
    }
    return res
}

func main() {
    fmt.Println(largestNumber([]int{10,2})) // 210
    fmt.Println(largestNumber([]int{3,30,34,5,9})) // 9534330

    fmt.Println(largestNumber1([]int{10,2})) // 210
    fmt.Println(largestNumber1([]int{3,30,34,5,9})) // 9534330

    fmt.Println(largestNumber2([]int{10,2})) // 210
    fmt.Println(largestNumber2([]int{3,30,34,5,9})) // 9534330
}