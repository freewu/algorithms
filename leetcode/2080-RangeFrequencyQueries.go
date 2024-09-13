package main

// 2080. Range Frequency Queries
// Design a data structure to find the frequency of a given value in a given subarray.

// The frequency of a value in a subarray is the number of occurrences of that value in the subarray.

// Implement the RangeFreqQuery class:
//     RangeFreqQuery(int[] arr) 
//         Constructs an instance of the class with the given 0-indexed integer array arr.
//     int query(int left, int right, int value) 
//         Returns the frequency of value in the subarray arr[left...right].

// A subarray is a contiguous sequence of elements within an array. 
// arr[left...right] denotes the subarray that contains the elements of nums between indices left and right (inclusive).

// Example 1:
// Input
// ["RangeFreqQuery", "query", "query"]
// [[[12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]], [1, 2, 4], [0, 11, 33]]
// Output
// [null, 1, 2]
// Explanation
// RangeFreqQuery rangeFreqQuery = new RangeFreqQuery([12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]);
// rangeFreqQuery.query(1, 2, 4); // return 1. The value 4 occurs 1 time in the subarray [33, 4]
// rangeFreqQuery.query(0, 11, 33); // return 2. The value 33 occurs 2 times in the whole array.

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= arr[i], value <= 10^4
//     0 <= left <= right < arr.length
//     At most 10^5 calls will be made to query

import "fmt"
import "sort"

type RangeFreqQuery struct {
    data map[int][]int // 记录 值 和出现在 arr 中的 index
}

func Constructor(arr []int) RangeFreqQuery {
    mp := make(map[int][]int)
    for i, v := range arr {
        mp[v] = append(mp[v], i)
    }
    return RangeFreqQuery{ mp }
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    if v,ok := this.data[value]; ok {
        lower := sort.Search(len(v), func (i int) bool {
            return v[i] >= left
        })
        upper := sort.Search(len(v), func (i int) bool {
            return v[i] > right
        })
        return upper - lower
    }
    return 0
}


type RangeFreqQuery1 struct {
	data map[int][]int
}

func Constructor1(arr []int) RangeFreqQuery1 {
    mp := make(map[int][]int)
    for k, v := range arr {
        mp[v] = append(mp[v], k)
    }
    return RangeFreqQuery1{ mp }
}

func (this *RangeFreqQuery1) Query(left int, right int, value int) int {
    v := this.data[value]
    // >=left && <=right (equal ">=right+1"-1)
    l, r := this.binaraySearch(v, left), this.binaraySearch(v, right + 1) - 1
    return r - l + 1
}

// left bound >=k []
func (this *RangeFreqQuery1) binaraySearch(nums []int,k int) int {
    l, r := -1, len(nums)
    for l + 1 < r {
        m := l + (r - l) / 2
        if nums[m] < k {
            l = m
        } else {
            r = m
        }
    }
    return l
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

func main() {
    // RangeFreqQuery rangeFreqQuery = new RangeFreqQuery([12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]);
    obj := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
    fmt.Println(obj)
    // rangeFreqQuery.query(1, 2, 4); // return 1. The value 4 occurs 1 time in the subarray [33, 4]
    fmt.Println(obj.Query(1, 2, 4)) // 1
    // rangeFreqQuery.query(0, 11, 33); // return 2. The value 33 occurs 2 times in the whole array.
    fmt.Println(obj.Query(0, 11, 33)) // 2

    obj1 := Constructor1([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
    fmt.Println(obj1)
    // rangeFreqQuery.query(1, 2, 4); // return 1. The value 4 occurs 1 time in the subarray [33, 4]
    fmt.Println(obj1.Query(1, 2, 4)) // 1
    // rangeFreqQuery.query(0, 11, 33); // return 2. The value 33 occurs 2 times in the whole array.
    fmt.Println(obj1.Query(0, 11, 33)) // 2
}