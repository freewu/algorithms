package main

// 1130. Minimum Cost Tree From Leaf Values
// Given an array arr of positive integers, consider all binary trees such that:
//     Each node has either 0 or 2 children;
//     The values of arr correspond to the values of each leaf in an in-order traversal of the tree.
//     The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree, respectively.

// Among all possible binary trees considered, return the smallest possible sum of the values of each non-leaf node. 
// It is guaranteed this sum fits into a 32-bit integer.

// A node is a leaf if and only if it has zero children.

// Example 1:
//         24                24
//        /   \             /  \
//       12    4    =>     6    8
//      /  \                   /
//     6    2                 2
// <img src="https://assets.leetcode.com/uploads/2021/08/10/tree1.jpg" />
// Input: arr = [6,2,4]
// Output: 32
// Explanation: There are two possible trees shown.
// The first has a non-leaf node sum 36, and the second has non-leaf node sum 32.

// Example 2:
//         44
//        /   \
//       4     11
// <img src="https://assets.leetcode.com/uploads/2021/08/10/tree2.jpg" />
// Input: arr = [4,11]
// Output: 44

// Constraints:
//     2 <= arr.length <= 40
//     1 <= arr[i] <= 15
//     It is guaranteed that the answer fits into a 32-bit signed integer (i.e., it is less than 2^31).

import "fmt"

func mctFromLeafValues(arr []int) int {
    res := 0;
    for len(arr) > 1 {
        index, cost := 1, arr[0] * arr[1]
        for i := 2; i < len(arr); i += 1 {
            val := arr[i - 1] * arr[i]
            if (val < cost) {
                index = i
                cost = val
            }
        }
        if (arr[index - 1] > arr[index]) {
            arr[index] = arr[index - 1]
        }
        copy(arr[index - 1:], arr[index:])
        arr = arr[:len(arr) - 1]
        res += cost
    }
    return res
}

func main() {
    // Example 1:
    //         24                24
    //        /   \             /  \
    //       12    4    =>     6    8
    //      /  \                   /
    //     6    2                 2
    // <img src="https://assets.leetcode.com/uploads/2021/08/10/tree1.jpg" />
    // Input: arr = [6,2,4]
    // Output: 32
    // Explanation: There are two possible trees shown.
    // The first has a non-leaf node sum 36, and the second has non-leaf node sum 32.
    fmt.Println(mctFromLeafValues([]int{6,2,4})) // 32
    // Example 2:
    //         44
    //        /   \
    //       4     11
    // <img src="https://assets.leetcode.com/uploads/2021/08/10/tree2.jpg" />
    // Input: arr = [4,11]
    // Output: 44
    fmt.Println(mctFromLeafValues([]int{4,11})) // 44
}