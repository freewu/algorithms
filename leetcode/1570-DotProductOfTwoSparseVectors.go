package main

// 1570. Dot Product of Two Sparse Vectors
// Given two sparse vectors, compute their dot product.

// Implement class SparseVector:
//     SparseVector(nums) Initializes the object with the vector nums
//     dotProduct(vec) Compute the dot product between the instance of SparseVector and vec

// A sparse vector is a vector that has mostly zero values, 
// you should store the sparse vector efficiently and compute the dot product between two SparseVector.

// Follow up: What if only one of the vectors is sparse?

// Example 1:
// Input: nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
// Output: 8
// Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
// v1.dotProduct(v2) = 1*0 + 0*3 + 0*0 + 2*4 + 3*0 = 8

// Example 2:
// Input: nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
// Output: 0
// Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
// v1.dotProduct(v2) = 0*0 + 1*0 + 0*0 + 0*0 + 0*2 = 0

// Example 3:
// Input: nums1 = [0,1,0,0,2,0,0], nums2 = [1,0,0,0,3,0,4]
// Output: 6

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     0 <= nums1[i], nums2[i] <= 100

import "fmt"

type SparseVector struct {
    data []int
}

func Constructor(nums []int) SparseVector {
    return SparseVector{ data: nums }
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    res, n := 0, len(this.data)
    for i := 0; i < n; i++ {
        res += this.data[i] * vec.data[i]
    }
    return res
}

type SparseVector1 struct {
    mp map[int]int
}

func Constructor1(nums []int) SparseVector1 {
    mp := map[int]int{}
    for i, v := range nums {
        if v != 0 {
            mp[i] = v
        }
    }
    return SparseVector1{ mp: mp }
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector1) dotProduct(vec SparseVector1) int {
    res := 0
    for i, n := range this.mp {
        if v, ok := vec.mp[i]; ok {
            res += n * v
        }
    }
    return res
}

type SparseVector2 struct {
    numPairs [][2]int
}

func Constructor2(nums []int) SparseVector2 {
    numPairs := [][2]int{}
    for i, num := range nums {
        if num != 0 {
            numPairs = append(numPairs, [2]int{i, num})
        }
    }
    return SparseVector2{numPairs: numPairs}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector2) dotProduct(vec SparseVector2) int {
    numPairs1, numPairs2 := this.numPairs, vec.numPairs
    res, i, j := 0, 0, 0
    for i < len(numPairs1) && j < len(numPairs2) {
        if numPairs1[i][0] == numPairs2[j][0] {
            res += numPairs1[i][1]*numPairs2[j][1]
            i++
            j++
        } else if numPairs1[i][0] > numPairs2[j][0] {
            j++
        } else {
            i++
        }
    }
    return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */

func main() {
    // Example 1:
    // Input: nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
    // Output: 8
    // Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
    // v1.dotProduct(v2) = 1*0 + 0*3 + 0*0 + 2*4 + 3*0 = 8
    obj1 := Constructor([]int{1,0,0,2,3})
    fmt.Println(obj1.dotProduct(Constructor([]int{0,3,0,4,0}))) // 8
    // Example 2:
    // Input: nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
    // Output: 0
    // Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
    // v1.dotProduct(v2) = 0*0 + 1*0 + 0*0 + 0*0 + 0*2 = 0
    obj2:= Constructor([]int{0,1,0,0,0})
    fmt.Println(obj2.dotProduct(Constructor([]int{0,0,0,0,2}))) // 0
    // Example 3:
    // Input: nums1 = [0,1,0,0,2,0,0], nums2 = [1,0,0,0,3,0,4]
    // Output: 6
    obj3:= Constructor([]int{0,1,0,0,2,0,0})
    fmt.Println(obj3.dotProduct(Constructor([]int{1,0,0,0,3,0,4}))) // 6

    obj11 := Constructor1([]int{1,0,0,2,3})
    fmt.Println(obj11.dotProduct(Constructor1([]int{0,3,0,4,0}))) // 8
    obj12:= Constructor1([]int{0,1,0,0,0})
    fmt.Println(obj12.dotProduct(Constructor1([]int{0,0,0,0,2}))) // 0
    obj13:= Constructor1([]int{0,1,0,0,2,0,0})
    fmt.Println(obj13.dotProduct(Constructor1([]int{1,0,0,0,3,0,4}))) // 6

    obj21 := Constructor2([]int{1,0,0,2,3})
    fmt.Println(obj21.dotProduct(Constructor2([]int{0,3,0,4,0}))) // 8
    obj22:= Constructor2([]int{0,1,0,0,0})
    fmt.Println(obj22.dotProduct(Constructor2([]int{0,0,0,0,2}))) // 0
    obj23:= Constructor2([]int{0,1,0,0,2,0,0})
    fmt.Println(obj23.dotProduct(Constructor2([]int{1,0,0,0,3,0,4}))) // 6
}