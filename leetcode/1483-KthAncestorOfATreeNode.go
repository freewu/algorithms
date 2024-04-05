package main

// 1483. Kth Ancestor of a Tree Node
// You are given a tree with n nodes numbered from 0 to n - 1 in the form of a parent array parent where parent[i] is the parent of ith node. The root of the tree is node 0. 
// Find the kth ancestor of a given node.
// The kth ancestor of a tree node is the kth node in the path from that node to the root node.

// Implement the TreeAncestor class:
//     TreeAncestor(int n, int[] parent) 
//         Initializes the object with the number of nodes in the tree and the parent array.
//     int getKthAncestor(int node, int k) 
//         return the kth ancestor of the given node node. 
//         If there is no such ancestor, return -1.
    
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/08/28/1528_ex1.png" />
// Input
// ["TreeAncestor", "getKthAncestor", "getKthAncestor", "getKthAncestor"]
// [[7, [-1, 0, 0, 1, 1, 2, 2]], [3, 1], [5, 2], [6, 3]]
// Output
// [null, 1, 0, -1]
// Explanation
// TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
// treeAncestor.getKthAncestor(3, 1); // returns 1 which is the parent of 3
// treeAncestor.getKthAncestor(5, 2); // returns 0 which is the grandparent of 5
// treeAncestor.getKthAncestor(6, 3); // returns -1 because there is no such ancestor

// Constraints:
//     1 <= k <= n <= 5 * 10^4
//     parent.length == n
//     parent[0] == -1
//     0 <= parent[i] < n for all 0 < i < n
//     0 <= node < n
//     There will be at most 5 * 10^4 queries.

import "fmt"
import "math"

type TreeAncestor struct {
    parents [][]int
    steps  int
}

func Constructor(n int, parent []int) TreeAncestor {
    steps := int(math.Log2(float64(n))) + 1
    parents := make([][]int, steps)
    parents[0] = parent
    for i := 1; i < steps; i++ {
        parents[i] = make([]int, n)
        for j := 0; j < n; j++ {
            last := parents[i-1][j]
            if last == -1 {
                parents[i][j] = -1
            } else {
                parents[i][j] = parents[i-1][last]
            }
        }
    }
    return TreeAncestor{parents, steps}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
    for i := 0; i < this.steps; i++ {
        if (k & (1 << i)) > 0 {
            val := this.parents[i][node]
            if val == -1 {
                return -1
            }
            node = val
        }
    }
    return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */

const kLog = 16

type TreeAncestor1 struct {
    ancestors [][kLog]int
}

func Constructor1(n int, parent []int) TreeAncestor1 {
    var this TreeAncestor1
    // 节点i的第2^j个祖先
    // 我们需要存储的父节点分别是: 0,1,2,4,8,16 ..., 所以5w个节点的最大层级是16
    // 如果要查找的是二的整数次幂的父节点, 可以直接找到数组的最高位1定位
    // 那么: 如何应对非二的整数次幂的父节点呢?
    // 这里面就存在了状态转移问题.
    // 比如第3个父节点, 其二进制表示为 0b11,
    // 第一位不为0, 则跳转到父节点, 此时就变成了查找父节点的第二个父节点,
    // 比如第6个父节点, 其二进制表示为 0b110,
    // 第一位是0, 不需要操作, 第二位不为零, 则先跳转到第二个父节点, 与之对应的就是第二个父节点的第四个父节点
    this.ancestors = make([][kLog]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < kLog; j++ {
            this.ancestors[i][j] = -1
        }
        this.ancestors[i][0] = parent[i]
    }
    for j := 1; j < kLog; j++ {
        for i := 0; i < n; i++ {
            if this.ancestors[i][j - 1] != -1 {
                this.ancestors[i][j] = this.ancestors[this.ancestors[i][j - 1]][j - 1]
            }
        }
    }
    return this
}

func (this *TreeAncestor1) GetKthAncestor(node int, k int) int {
    for j := 0; j < kLog; j++ {
        if (k >> j) & 1 != 0 {
            node = this.ancestors[node][j]
            if node == -1 {
                return -1
            }
        }
    }
    return node
}

func main() {
    // Explanation
    // TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
    obj := Constructor(7,[]int{-1, 0, 0, 1, 1, 2, 2})
    // treeAncestor.getKthAncestor(3, 1); // returns 1 which is the parent of 3
    fmt.Println(obj.GetKthAncestor(3, 1)) // 1
    // treeAncestor.getKthAncestor(5, 2); // returns 0 which is the grandparent of 5
    fmt.Println(obj.GetKthAncestor(5, 2)) // 0
    // treeAncestor.getKthAncestor(6, 3); // returns -1 because there is no such ancestor
    fmt.Println(obj.GetKthAncestor(6, 3)) // -1

    obj1 := Constructor1(7,[]int{-1, 0, 0, 1, 1, 2, 2})
    // treeAncestor.getKthAncestor(3, 1); // returns 1 which is the parent of 3
    fmt.Println(obj1.GetKthAncestor(3, 1)) // 1
    // treeAncestor.getKthAncestor(5, 2); // returns 0 which is the grandparent of 5
    fmt.Println(obj1.GetKthAncestor(5, 2)) // 0
    // treeAncestor.getKthAncestor(6, 3); // returns -1 because there is no such ancestor
    fmt.Println(obj1.GetKthAncestor(6, 3)) // -1
}