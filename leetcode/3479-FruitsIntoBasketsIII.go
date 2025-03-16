package main

// 3479. Fruits Into Baskets III
// You are given two arrays of integers, fruits and baskets, each of length n, where fruits[i] represents the quantity of the ith type of fruit, and baskets[j] represents the capacity of the jth basket.

// From left to right, place the fruits according to these rules:
//     1. Each fruit type must be placed in the leftmost available basket with a capacity greater than or equal to the quantity of that fruit type.
//     2. Each basket can hold only one type of fruit.
//     3. If a fruit type cannot be placed in any basket, it remains unplaced.

// Return the number of fruit types that remain unplaced after all possible allocations are made.

// Example 1:
// Input: fruits = [4,2,5], baskets = [3,5,4]
// Output: 1
// Explanation:
// fruits[0] = 4 is placed in baskets[1] = 5.
// fruits[1] = 2 is placed in baskets[0] = 3.
// fruits[2] = 5 cannot be placed in baskets[2] = 4.
// Since one fruit type remains unplaced, we return 1.

// Example 2:
// Input: fruits = [3,6,1], baskets = [6,4,7]
// Output: 0
// Explanation:
// fruits[0] = 3 is placed in baskets[0] = 6.
// fruits[1] = 6 cannot be placed in baskets[1] = 4 (insufficient capacity) but can be placed in the next available basket, baskets[2] = 7.
// fruits[2] = 1 is placed in baskets[1] = 4.
// Since all fruits are successfully placed, we return 0.

// Constraints:
//     n == fruits.length == baskets.length
//     1 <= n <= 10^5
//     1 <= fruits[i], baskets[i] <= 10^9

import "fmt"

type SegmentTree struct {
    size int
    data []int
}

// NewSegmentTree builds a segment tree to store max values for range queries.
func NewSegmentTree(values []int) *SegmentTree {
    if len(values) == 0 {
        return &SegmentTree{0, []int{}}
    }
    // nextPowerOf2 returns the smallest power of 2 greater than or equal to n.
    nextPowerOf2 := func (n int) int {
        power := 1
        for power < n {
            power *= 2
        }
        return power
    }
    size := nextPowerOf2(len(values))
    data := make([]int, 2*size)
    // Copy values into the leaf nodes
    copy(data[size:size+len(values)], values)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // Build the tree from leaves to root
    for i := size - 1; i > 0; i-- {
        data[i] = max(data[2 * i], data[2 * i + 1])
    }
    return &SegmentTree{size, data}
}

// FindFirstGreaterOrEqual returns the index of the first element >= val, or -1 if none exist.
func (tree *SegmentTree) FindFirstGreaterOrEqual(val int) int {
    pos := 1
    for pos < tree.size && tree.data[pos] >= val {
        pos *= 2
        if tree.data[pos] < val {
            pos++
        }
    }
    if pos >= tree.size && tree.data[pos] >= val {
        return pos - tree.size
    }
    return -1
}

// Update sets the value at index idx and updates the segment tree.
func (tree *SegmentTree) Update(idx, val int) {
    pos := tree.size + idx
    tree.data[pos] = val
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for pos > 1 {
        pos /= 2
        tree.data[pos] = max(tree.data[2 * pos], tree.data[2 * pos + 1])
    }
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    tree := NewSegmentTree(baskets)
    res := 0
    for _, fruit := range fruits {
        index := tree.FindFirstGreaterOrEqual(fruit)
        if index == -1 {
            res++
        } else {
            tree.Update(index, 0)
        }
    }
    return res
}

func numOfUnplacedFruits1(fruits []int, baskets []int) int {
    n := len(fruits)
    if n == 0 {
        return 0
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    size := 1
    for size < n { // 构建线段树数组，大小取最近的 2 的幂
        size <<= 1
    }
    segTree := make([]int, size << 1)
    for i := 0; i < n; i++ { // 初始化线段树（叶子节点）
        segTree[size+i] = baskets[i]
    }
    for i := size - 1; i > 0; i-- { // 自底向上构建线段树
        segTree[i] = max(segTree[i << 1], segTree[i << 1 | 1])
    }
    update := func(pos, val int) { // 更新单点值
        pos += size
        segTree[pos] = val
        for pos > 1 {
            pos >>= 1
            segTree[pos] = max(segTree[pos << 1], segTree[pos << 1 | 1])
        }
    }
    findLeftmost := func(x int) int { // 查找最左侧篮子下标（容量 >= x）
        if segTree[1] < x { return -1 }
        index := 1
        for index < size {
            left := index << 1
            if segTree[left] >= x {
                index = left
            } else {
                index = left | 1
            }
        }
        return index - size
    }
    unplaced := 0
    for i := 0; i < n; i++ {
        index := findLeftmost(fruits[i])
        if index == -1 {
            unplaced++
        } else {
            update(index, -1)
        }
    }
    return unplaced
}

func main() {
    // Example 1:
    // Input: fruits = [4,2,5], baskets = [3,5,4]
    // Output: 1
    // Explanation:
    // fruits[0] = 4 is placed in baskets[1] = 5.
    // fruits[1] = 2 is placed in baskets[0] = 3.
    // fruits[2] = 5 cannot be placed in baskets[2] = 4.
    // Since one fruit type remains unplaced, we return 1.
    fmt.Println(numOfUnplacedFruits([]int{4,2,5}, []int{3,5,4})) // 1
    // Example 2:
    // Input: fruits = [3,6,1], baskets = [6,4,7]
    // Output: 0
    // Explanation:
    // fruits[0] = 3 is placed in baskets[0] = 6.
    // fruits[1] = 6 cannot be placed in baskets[1] = 4 (insufficient capacity) but can be placed in the next available basket, baskets[2] = 7.
    // fruits[2] = 1 is placed in baskets[1] = 4.
    // Since all fruits are successfully placed, we return 0.
    fmt.Println(numOfUnplacedFruits([]int{3,6,1}, []int{6,4,7})) // 0

    fmt.Println(numOfUnplacedFruits([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 4
    fmt.Println(numOfUnplacedFruits([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(numOfUnplacedFruits([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(numOfUnplacedFruits([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0

    fmt.Println(numOfUnplacedFruits1([]int{4,2,5}, []int{3,5,4})) // 1
    fmt.Println(numOfUnplacedFruits1([]int{3,6,1}, []int{6,4,7})) // 0
    fmt.Println(numOfUnplacedFruits1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 4
    fmt.Println(numOfUnplacedFruits1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(numOfUnplacedFruits1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(numOfUnplacedFruits1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
}