package main

// 3638. Maximum Balanced Shipments
// You are given an integer array weight of length n, representing the weights of n parcels arranged in a straight line. 
// A shipment is defined as a contiguous subarray of parcels. 
// A shipment is considered balanced if the weight of the last parcel is strictly less than the maximum weight among all parcels in that shipment.

// Select a set of non-overlapping, contiguous, balanced shipments such that each parcel appears in at most one shipment (parcels may remain unshipped).

// Return the maximum possible number of balanced shipments that can be formed.

// Example 1:
// Input: weight = [2,5,1,4,3]
// Output: 2
// Explanation:
// We can form the maximum of two balanced shipments as follows:
// Shipment 1: [2, 5, 1]
// Maximum parcel weight = 5
// Last parcel weight = 1, which is strictly less than 5. Thus, it's balanced.
// Shipment 2: [4, 3]
// Maximum parcel weight = 4
// Last parcel weight = 3, which is strictly less than 4. Thus, it's balanced.
// It is impossible to partition the parcels to achieve more than two balanced shipments, so the answer is 2.

// Example 2:
// Input: weight = [4,4]
// Output: 0
// Explanation:
// No balanced shipment can be formed in this case:
// A shipment [4, 4] has maximum weight 4 and the last parcel's weight is also 4, which is not strictly less. Thus, it's not balanced.
// Single-parcel shipments [4] have the last parcel weight equal to the maximum parcel weight, thus not balanced.
// As there is no way to form even one balanced shipment, the answer is 0.

// Constraints:
//     2 <= n <= 10^5
//     1 <= weight[i] <= 10^9

import "fmt"

func maxBalancedShipments(weight []int) int {
    res, n := 0, len(weight)
    stack := []int{}
    for i := 0; i < n; i++ {
        if len(stack) == 0 || stack[len(stack) - 1] < weight[i] {
            stack = append(stack, weight[i])
        } else if stack[len(stack) - 1] > weight[i] {
            stack=[]int{}
            res++
        }
    }
    return res
}

func maxBalancedShipments1(weight []int) int {
    res, n, best := 0, len(weight), weight[0]
    if n == 0 {  return 0 }
    for i := 1; i < n; i++ {
        if best > weight[i] {
            res++
            if i + 1 < n {
                i++
                best = weight[i]
            }
        } else {
            best = weight[i]
        }
    }
    return res
}

func maxBalancedShipments2(weight []int) int {
    res := 0
    for i := 1; i < len(weight); i++ {
        if weight[i-1] > weight[i] {
            res++
            i++ // 下个子数组至少要有两个数
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: weight = [2,5,1,4,3]
    // Output: 2
    // Explanation:
    // We can form the maximum of two balanced shipments as follows:
    // Shipment 1: [2, 5, 1]
    // Maximum parcel weight = 5
    // Last parcel weight = 1, which is strictly less than 5. Thus, it's balanced.
    // Shipment 2: [4, 3]
    // Maximum parcel weight = 4
    // Last parcel weight = 3, which is strictly less than 4. Thus, it's balanced.
    // It is impossible to partition the parcels to achieve more than two balanced shipments, so the answer is 2.
    fmt.Println(maxBalancedShipments([]int{2,5,1,4,3})) // 2
    // Example 2:
    // Input: weight = [4,4]
    // Output: 0
    // Explanation:
    // No balanced shipment can be formed in this case:
    // A shipment [4, 4] has maximum weight 4 and the last parcel's weight is also 4, which is not strictly less. Thus, it's not balanced.
    // Single-parcel shipments [4] have the last parcel weight equal to the maximum parcel weight, thus not balanced.
    // As there is no way to form even one balanced shipment, the answer is 0.
    fmt.Println(maxBalancedShipments([]int{4,4})) // 0

    fmt.Println(maxBalancedShipments([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxBalancedShipments([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(maxBalancedShipments1([]int{2,5,1,4,3})) // 2
    fmt.Println(maxBalancedShipments1([]int{4,4})) // 0
    fmt.Println(maxBalancedShipments1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxBalancedShipments1([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(maxBalancedShipments2([]int{2,5,1,4,3})) // 2
    fmt.Println(maxBalancedShipments2([]int{4,4})) // 0
    fmt.Println(maxBalancedShipments2([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxBalancedShipments2([]int{9,8,7,6,5,4,3,2,1})) // 0
}