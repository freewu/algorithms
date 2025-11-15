package main

// 307. Range Sum Query - Mutable
// Given an integer array nums, handle multiple queries of the following types:

// Update the value of an element in nums.
// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:
//     NumArray(int[] nums) Initializes the object with the integer array nums.
//     void update(int index, int val) Updates the value of nums[index] to be val.
//     int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
    
// Example 1:
// Input
// ["NumArray", "sumRange", "update", "sumRange"]
// [[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
// Output	
// [null, 9, null, 8]
// Explanation
// NumArray numArray = new NumArray([1, 3, 5]);
// numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
// numArray.update(1, 2);   // nums = [1, 2, 5]
// numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8

// Constraints:
//     1 <= nums.length <= 3 * 104
//     -100 <= nums[i] <= 100
//     0 <= index < nums.length
//     -100 <= val <= 100
//     0 <= left <= right < nums.length
//     At most 3 * 104 calls will be made to update and sumRange.

// 线段树 Segment tree 是一种二叉树形数据结构，1977年由 Jon Louis Bentley 发明，用以存储区间或线段，并且允许快速查询结构内包含某一点的所有区间。
// 	一个包含 n 个区间的线段树，空间复杂度为 O(n) ，
// 	查询的时间复杂度则为 O(logn+k) ，其中 k 是符合条件的区间数量。线段树的数据结构也可推广到高维度。

import "fmt"

// SegmentTree define
type SegmentTree struct {
	data, tree, lazy []int
	left, right      int
	merge            func(i, j int) int
}

// Init define
func (st *SegmentTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper
	data, tree, lazy := make([]int, len(nums)), make([]int, 4*len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree, st.lazy = data, tree, lazy
	if len(nums) > 0 {
		st.buildSegmentTree(0, 0, len(nums)-1)
	}
}

// 在 treeIndex 的位置创建 [left....right] 区间的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}

// 查询 [left....right] 区间内的值
// Query define
func (st *SegmentTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// 在以 treeIndex 为根的线段树中 [left...right] 的范围里，搜索区间 [queryLeft...queryRight] 的值
func (st *SegmentTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if left == queryLeft && right == queryRight {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if queryLeft > midTreeIndex {
		return st.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	return st.merge(st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.queryInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

// QueryLazy define
func (st *SegmentTree) QueryLazy(left, right int) int {
	if len(st.data) > 0 {
		return st.queryLazyInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

func (st *SegmentTree) queryLazyInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if left > queryRight || right < queryLeft { // segment completely outside range
		return 0 // represents a null node
	}
	if st.lazy[treeIndex] != 0 { // this node is lazy
		for i := 0; i < right-left+1; i++ {
			st.tree[treeIndex] = st.merge(st.tree[treeIndex], st.lazy[treeIndex])
			// st.tree[treeIndex] += (right - left + 1) * st.lazy[treeIndex] // normalize current node by removing lazinesss
		}
		if left != right { // update lazy[] for children nodes
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], st.lazy[treeIndex])
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], st.lazy[treeIndex])
			// st.lazy[leftTreeIndex] += st.lazy[treeIndex]
			// st.lazy[rightTreeIndex] += st.lazy[treeIndex]
		}
		st.lazy[treeIndex] = 0 // current node processed. No longer lazy
	}
	if queryLeft <= left && queryRight >= right { // segment completely inside range
		return st.tree[treeIndex]
	}
	if queryLeft > midTreeIndex {
		return st.queryLazyInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.queryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// merge query results
	return st.merge(st.queryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.queryLazyInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

// 更新 index 位置的值

// Update define
func (st *SegmentTree) Update(index, val int) {
	if len(st.data) > 0 {
		st.updateInTree(0, 0, len(st.data)-1, index, val)
	}
}

// 以 treeIndex 为根，更新 index 位置上的值为 val
func (st *SegmentTree) updateInTree(treeIndex, left, right, index, val int) {
	if left == right {
		st.tree[treeIndex] = val
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if index > midTreeIndex {
		st.updateInTree(rightTreeIndex, midTreeIndex+1, right, index, val)
	} else {
		st.updateInTree(leftTreeIndex, left, midTreeIndex, index, val)
	}
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// 更新 [updateLeft....updateRight] 位置的值
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论

// UpdateLazy define
func (st *SegmentTree) UpdateLazy(updateLeft, updateRight, val int) {
	if len(st.data) > 0 {
		st.updateLazyInTree(0, 0, len(st.data)-1, updateLeft, updateRight, val)
	}
}

func (st *SegmentTree) updateLazyInTree(treeIndex, left, right, updateLeft, updateRight, val int) {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if st.lazy[treeIndex] != 0 { // this node is lazy
		for i := 0; i < right-left+1; i++ {
			st.tree[treeIndex] = st.merge(st.tree[treeIndex], st.lazy[treeIndex])
			//st.tree[treeIndex] += (right - left + 1) * st.lazy[treeIndex] // normalize current node by removing laziness
		}
		if left != right { // update lazy[] for children nodes
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], st.lazy[treeIndex])
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], st.lazy[treeIndex])
			// st.lazy[leftTreeIndex] += st.lazy[treeIndex]
			// st.lazy[rightTreeIndex] += st.lazy[treeIndex]
		}
		st.lazy[treeIndex] = 0 // current node processed. No longer lazy
	}

	if left > right || left > updateRight || right < updateLeft {
		return // out of range. escape.
	}

	if updateLeft <= left && right <= updateRight { // segment is fully within update range
		for i := 0; i < right-left+1; i++ {
			st.tree[treeIndex] = st.merge(st.tree[treeIndex], val)
			//st.tree[treeIndex] += (right - left + 1) * val // update segment
		}
		if left != right { // update lazy[] for children
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], val)
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], val)
			// st.lazy[leftTreeIndex] += val
			// st.lazy[rightTreeIndex] += val
		}
		return
	}
	st.updateLazyInTree(leftTreeIndex, left, midTreeIndex, updateLeft, updateRight, val)
	st.updateLazyInTree(rightTreeIndex, midTreeIndex+1, right, updateLeft, updateRight, val)
	// merge updates
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}


// NumArray define
type NumArray struct {
	st *SegmentTree
}

//
func Constructor(nums []int) NumArray {
	st := SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: &st}
}

// Update define
func (this *NumArray) Update(i int, val int) {
	this.st.Update(i, val)
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	return this.st.Query(i, j)
}

// best solution
type NumArray struct {
    Bit []int
    N int
    Nums []int
}

func lowbit(x int) int {
    return x & (-x)
}

func Constructor(nums []int) NumArray {
    res := NumArray{
        Nums: nums,
        Bit: make([]int,len(nums)+1),
        N:len(nums),}
    
    for i := 1; i <= res.N; i++ {
        res.Bit[i] = nums[i-1]
        for j := i - 2; j >= i - lowbit(i); j-- {
            res.Bit[i] += nums[j]
        }
    }
    return res
}


func (this *NumArray) Update(index int, val int)  {
    addVal := val - this.Nums[index]
    for i := index+1; i <= this.N; i += lowbit(i) {
        this.Bit[i] += addVal
    }
    this.Nums[index] = val
}

func (this *NumArray) Query(k int) int {
    ans := 0
    for i := k; i > 0; i -= lowbit(i) {
        ans += this.Bit[i]
    }
    return ans
}

func (this *NumArray) SumRange(left int, right int) int {
    return this.Query(right+1) - this.Query(left)
}

func main() {
    // NumArray numArray = new NumArray([1, 3, 5]);
    obj := Constructor([]int{1, 3, 5})
    fmt.Println(obj) // [1, 3, 5]
    // numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
    fmt.Println(obj.SumRange(0, 2)) // 9
    // numArray.update(1, 2);   // nums = [1, 2, 5]
    obj.Update(1, 2)
    fmt.Println(obj) // [1, 2, 5]
    // numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8
    fmt.Println(obj.SumRange(0, 2)) // 8
}