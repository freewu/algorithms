package main

// 1157. Online Majority Element In Subarray
// Design a data structure that efficiently finds the majority element of a given subarray.
// The majority element of a subarray is an element that occurs threshold times or more in the subarray.

// Implementing the MajorityChecker class:
//     MajorityChecker(int[] arr) 
//         Initializes the instance of the class with the given array arr.
//     int query(int left, int right, int threshold) 
//         returns the element in the subarray arr[left...right] that occurs at least threshold times, 
//         or -1 if no such element exists.

// Example 1:
// Input
// ["MajorityChecker", "query", "query", "query"]
// [[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]
// Output
// [null, 1, -1, 2]
// Explanation
// MajorityChecker majorityChecker = new MajorityChecker([1, 1, 2, 2, 1, 1]);
// majorityChecker.query(0, 5, 4); // return 1
// majorityChecker.query(0, 3, 3); // return -1
// majorityChecker.query(2, 3, 2); // return 2

// Constraints:
//     1 <= arr.length <= 2 * 10^4
//     1 <= arr[i] <= 2 * 10^4
//     0 <= left <= right < arr.length
//     threshold <= right - left + 1
//     2 * threshold > right - left + 1
//     At most 10^4 calls will be made to query.

import "fmt"
import "sort"

type MajorityChecker struct {
    digits  int
    presum  [][]int
    pos     map[int][]int 
}

func Constructor(arr []int) MajorityChecker {
    mc := MajorityChecker{ 15, make([][]int, len(arr)+1), make(map[int][]int), }
    for i := range mc.presum {
        mc.presum[i] = make([]int, mc.digits)
    }
    for i := 0; i < len(arr); i++ {
        n := arr[i]
        if _, ok := mc.pos[n]; !ok {
            mc.pos[n] = make([]int, 0)
        }
        mc.pos[n] = append(mc.pos[n], i)
        for j := 0; j < mc.digits; j++ {
            mc.presum[i+1][j] = mc.presum[i][j] + (n & 1)
            n >>= 1
        }
    }
    return mc
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
    res := 0
    for i := this.digits - 1; i >= 0; i-- {
        cnt := this.presum[right+1][i] - this.presum[left][i]
        b := 1
        if cnt >= threshold {
            b = 1
        } else if right - left + 1 - cnt >= threshold {
            b = 0
        } else {
            return -1
        }
        res = (res << 1) + b
    }
    list := this.pos[res]
    if list == nil {
        return -1
    }
    l, r := this.floor(list, left - 1), this.floor(list, right)
    if r - l >= threshold {
        return res
    }
    return -1
}

func (this *MajorityChecker) floor(list []int, n int) int {
    return sort.Search(len(list), func(i int) bool {
        return list[i] > n
    })
}

func main() {
    // MajorityChecker majorityChecker = new MajorityChecker([1, 1, 2, 2, 1, 1]);
    obj := Constructor([]int{1, 1, 2, 2, 1, 1})
    fmt.Println(obj)
    // majorityChecker.query(0, 5, 4); // return 1
    fmt.Println(obj.Query(0, 5, 4)) // 1
    // majorityChecker.query(0, 3, 3); // return -1
    fmt.Println(obj.Query(0, 3, 3)) // -1
    // majorityChecker.query(2, 3, 2); // return 2
    fmt.Println(obj.Query(2, 3, 2)) // 2
}