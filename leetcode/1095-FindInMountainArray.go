package main

// 1095. Find in Mountain Array
// (This problem is an interactive problem.)

// You may recall that an array arr is a mountain array if and only if:
//     arr.length >= 3
//     There exists some i with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// Given a mountain array mountainArr, return the minimum index such that mountainArr.get(index) == target. 
// If such an index does not exist, return -1.

// You cannot access the mountain array directly. 
// You may only access the array using a MountainArray interface:
//     MountainArray.get(k) returns the element of the array at index k (0-indexed).
//     MountainArray.length() returns the length of the array.

// Submissions making more than 100 calls to MountainArray.get will be judged Wrong Answer. 
// Also, any solutions that attempt to circumvent the judge will result in disqualification.

// Example 1:
// Input: array = [1,2,3,4,5,3,1], target = 3
// Output: 2
// Explanation: 3 exists in the array, at index=2 and index=5. Return the minimum index, which is 2.

// Example 2:
// Input: array = [0,1,2,4,2,1], target = 3
// Output: -1
// Explanation: 3 does not exist in the array, so we return -1.

// Constraints:
//     3 <= mountain_arr.length() <= 10^4
//     0 <= target <= 10^9
//     0 <= mountain_arr.get(index) <= 10^9

import "fmt"
import "sort"

type MountainArray struct {
    data []int 
}

func Contructor(data []int) *MountainArray {
    return &MountainArray{ data }
}

func (this *MountainArray) get(index int) int {
    return this.data[index]
}

func (this *MountainArray) length() int {
    return len(this.data)
}

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {
    n := mountainArr.length()
    left, right := 0, n-1
    mid := (right + left) / 2
    for mountainArr.get(mid-1) >= mountainArr.get(mid) || mountainArr.get(mid) <= mountainArr.get(mid+1) {
        if mountainArr.get(mid) <= mountainArr.get(mid+1) {
            left = mid
        } else {
            right = mid
        }
        mid = (right + left) / 2
    }
    if mountainArr.get(mid) == target {
        return mid
    }
    res := sort.Search(mid, func(i int) bool {
        return mountainArr.get(i) >= target
    })
    if mountainArr.get(res) != target {
        res = sort.Search(n-mid-1, func(i int) bool {
            return mountainArr.get(mid+i+1) <= target
        })
        res = mid + res + 1
        if n == res || mountainArr.get(res) != target {
            return -1
        }
        return res
    }
    return res
}

func main() {
    // Example 1:
    // Input: array = [1,2,3,4,5,3,1], target = 3
    // Output: 2
    // Explanation: 3 exists in the array, at index=2 and index=5. Return the minimum index, which is 2.
    arr1 := Contructor([]int{1,2,3,4,5,3,1})
    fmt.Println(findInMountainArray(3, arr1)) // 2
    // Example 2:
    // Input: array = [0,1,2,4,2,1], target = 3
    // Output: -1
    // Explanation: 3 does not exist in the array, so we return -1.
    arr2 := Contructor([]int{0,1,2,4,2,1})
    fmt.Println(findInMountainArray(3, arr2)) // -1
}