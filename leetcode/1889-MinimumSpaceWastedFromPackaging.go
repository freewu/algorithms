package main

// 1889. Minimum Space Wasted From Packaging
// You have n packages that you are trying to place in boxes, one package in each box. 
// There are m suppliers that each produce boxes of different sizes (with infinite supply). 
// A package can be placed in a box if the size of the package is less than or equal to the size of the box.

// The package sizes are given as an integer array packages, where packages[i] is the size of the ith package. 
// The suppliers are given as a 2D integer array boxes, where boxes[j] is an array of box sizes that the jth supplier produces.

// You want to choose a single supplier and use boxes from them such that the total wasted space is minimized. 
// For each package in a box, we define the space wasted to be size of the box - size of the package. 
// The total wasted space is the sum of the space wasted in all the boxes.
//     For example, if you have to fit packages with sizes [2,3,5] and the supplier offers boxes of sizes [4,8], 
//     you can fit the packages of size-2 and size-3 into two boxes of size-4 and the package with size-5 into a box of size-8. 
//     This would result in a waste of (4-2) + (4-3) + (8-5) = 6.

// Return the minimum total wasted space by choosing the box supplier optimally, 
// or -1 if it is impossible to fit all the packages inside boxes. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: packages = [2,3,5], boxes = [[4,8],[2,8]]
// Output: 6
// Explanation: It is optimal to choose the first supplier, using two size-4 boxes and one size-8 box.
// The total waste is (4-2) + (4-3) + (8-5) = 6.

// Example 2:
// Input: packages = [2,3,5], boxes = [[1,4],[2,3],[3,4]]
// Output: -1
// Explanation: There is no box that the package of size 5 can fit in.

// Example 3:
// Input: packages = [3,5,8,10,11,12], boxes = [[12],[11,9],[10,5,14]]
// Output: 9
// Explanation: It is optimal to choose the third supplier, using two size-5 boxes, two size-10 boxes, and two size-14 boxes.
// The total waste is (5-3) + (5-5) + (10-8) + (10-10) + (14-11) + (14-12) = 9.

// Constraints:
//     n == packages.length
//     m == boxes.length
//     1 <= n <= 10^5
//     1 <= m <= 10^5
//     1 <= packages[i] <= 10^5
//     1 <= boxes[j].length <= 10^5
//     1 <= boxes[j][k] <= 10^5
//     sum(boxes[j].length) <= 10^5
//     The elements in boxes[j] are distinct.

import "fmt"
import "sort"
import "math"

// PrefixSum + Binary Search
func minWastedSpace(packages []int, boxes [][]int) int {
    sort.Ints(packages)
    res, n, mod := math.MaxInt64, len(packages), 1_000_000_007
    prefix := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i-1] + packages[i-1]
    }
    binarySearch := func(nums []int, target int) int {
        left, right := 0, len(nums)-1
        for left < right {
            mid := (left + right) >> 1
            if nums[mid] <= target {
                left = mid + 1
            } else {
                right = mid
            }
        }
        if nums[left] > target { return left - 1 }
        return left
    }
    for _, box := range boxes {
        sort.Ints(box)
        if box[len(box) - 1] < packages[n - 1] { continue }
        j, cur := 0, 0 // cur is the space wasted in current box
        for _, b := range box {
            i := binarySearch(packages, b)
            cur += (i - j + 1) * b // put all packages from j to i-1 into box b
            cur -= prefix[i+1] - prefix[j] // remove the sum of packages from j to i-1
            j = i + 1
        }
        res = min(res, cur) // update the minimum space wasted
    }
    if res == math.MaxInt64 { return -1 }
    return res % mod
}

func minWastedSpace1(packages []int, boxes [][]int) int {
    mp, boxKinds := make(map[int]int),  make([]int, 0)
    for _, box := range boxes {
        sort.Ints(box)
        for _, onebox := range box {
            boxKinds = append(boxKinds, onebox) // 将每个种类的盒子记录在数组中
        }
    }
    sort.Ints(packages) 
    sort.Ints(boxKinds) // 盒子总共有多少种大小
    for i, box := range boxKinds {
        if _, ok := mp[box]; !ok {
            mp[box] = i   
        }
    }
    prefixNum, prefixSum := make([]int, len(boxKinds)), make([]int, len(boxKinds)) // 记录每个种类的盒子可以装多少个包裹, 记录每个种类的盒子，可以装的包裹的总大小
    pos, sum := 0, 0
    for i := 0; i < len(boxKinds); i++ {
        for pos < len(packages) && packages[pos] <= boxKinds[i] {
            sum += packages[pos]
            pos++
        }
        prefixNum[i], prefixSum[i] = pos, sum
    }
    res := -1
    for _, box := range boxes {
        if box[len(box) - 1] < packages[len(packages) - 1] { continue }
        diff := 0
        for i := 0; i < len(box); i++ {
            if i == 0 { diff += (prefixNum[mp[box[i]]] * box[i] - prefixSum[mp[box[i]]]) }
            if i != 0 {  diff += ((prefixNum[mp[box[i]]] - prefixNum[mp[box[i - 1]]]) * box[i] - (prefixSum[mp[box[i]]] - prefixSum[mp[box[i - 1]]])) }
        }
        if res == -1 {
            res = diff
        } else if diff < res {
            res = diff
        }
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: packages = [2,3,5], boxes = [[4,8],[2,8]]
    // Output: 6
    // Explanation: It is optimal to choose the first supplier, using two size-4 boxes and one size-8 box.
    // The total waste is (4-2) + (4-3) + (8-5) = 6.
    fmt.Println(minWastedSpace([]int{2,3,5}, [][]int{{4,8},{2,8}})) // 6
    // Example 2:
    // Input: packages = [2,3,5], boxes = [[1,4],[2,3],[3,4]]
    // Output: -1
    // Explanation: There is no box that the package of size 5 can fit in.
    fmt.Println(minWastedSpace([]int{2,3,5}, [][]int{{1,4},{2,3},{3,4}})) // -1
    // Example 3:
    // Input: packages = [3,5,8,10,11,12], boxes = [[12],[11,9],[10,5,14]]
    // Output: 9
    // Explanation: It is optimal to choose the third supplier, using two size-5 boxes, two size-10 boxes, and two size-14 boxes.
    // The total waste is (5-3) + (5-5) + (10-8) + (10-10) + (14-11) + (14-12) = 9.
    fmt.Println(minWastedSpace([]int{3,5,8,10,11,12}, [][]int{{12},{11,9},{10,5,14}})) // 9

    fmt.Println(minWastedSpace1([]int{2,3,5}, [][]int{{4,8},{2,8}})) // 6
    fmt.Println(minWastedSpace1([]int{2,3,5}, [][]int{{1,4},{2,3},{3,4}})) // -1
    fmt.Println(minWastedSpace1([]int{3,5,8,10,11,12}, [][]int{{12},{11,9},{10,5,14}})) // 9
}