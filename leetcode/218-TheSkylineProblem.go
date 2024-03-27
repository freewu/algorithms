package main

// 218. The Skyline Problem
// A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. 
// Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.
// The geometric information of each building is given in the array buildings where buildings[i] = [lefti, righti, heighti]:
//     lefti is the x coordinate of the left edge of the ith building.
//     righti is the x coordinate of the right edge of the ith building.
//     heighti is the height of the ith building.

// You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
// The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...]. 
// Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list, which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends. 
// Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

// Note: There must be no consecutive horizontal lines of equal height in the output skyline. 
// For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

// Example 1:
// <img sec="https://assets.leetcode.com/uploads/2020/12/01/merged.jpg" />
// Input: buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
// Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
// Explanation:
// Figure A shows the buildings of the input.
// Figure B shows the skyline formed by those buildings. The red points in figure B represent the key points in the output list.

// Example 2:
// Input: buildings = [[0,2,3],[2,5,3]]
// Output: [[0,3],[5,0]]

// Constraints:
//     1 <= buildings.length <= 10^4
//     0 <= lefti < righti <= 2^31 - 1
//     1 <= heighti <= 231 - 1
//     buildings is sorted by lefti in non-decreasing order.

// # 解题思路
//     给出一个二维数组，每个子数组里面代表一个高楼的信息，一个高楼的信息包含 3 个信息，高楼起始坐标，高楼终止坐标，高楼高度。
//     要求找到这些高楼的边际点，并输出这些边际点的高度信息。
//     这一题可以用树状数组来解答。要求天际线，即找到楼与楼重叠区间外边缘的线，说白了是维护各个区间内的最值。

import "fmt"
import "sort"
import "container/heap"

// 树状数组，时间复杂度 O(n log n)
const LEFTSIDE = 1
const RIGHTSIDE = 2

type Point struct {
	xAxis int
	side  int
	index int
}

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}
	if len(buildings) == 0 {
		return res
	}
	allPoints, bit := make([]Point, 0), BinaryIndexedTree{}
	// [x-axis (value), [1 (left) | 2 (right)], index (building number)]
	for i, b := range buildings {
		allPoints = append(allPoints, Point{xAxis: b[0], side: LEFTSIDE, index: i})
		allPoints = append(allPoints, Point{xAxis: b[1], side: RIGHTSIDE, index: i})
	}
	sort.Slice(allPoints, func(i, j int) bool {
		if allPoints[i].xAxis == allPoints[j].xAxis {
			return allPoints[i].side < allPoints[j].side
		}
		return allPoints[i].xAxis < allPoints[j].xAxis
	})
	bit.Init(len(allPoints))
	kth := make(map[Point]int)
	for i := 0; i < len(allPoints); i++ {
		kth[allPoints[i]] = i
	}
	for i := 0; i < len(allPoints); i++ {
		pt := allPoints[i]
		if pt.side == LEFTSIDE {
			bit.Add(kth[Point{xAxis: buildings[pt.index][1], side: RIGHTSIDE, index: pt.index}], buildings[pt.index][2])
		}
		currHeight := bit.Query(kth[pt] + 1)
		if len(res) == 0 || res[len(res)-1][1] != currHeight {
			if len(res) > 0 && res[len(res)-1][0] == pt.xAxis {
				res[len(res)-1][1] = currHeight
			} else {
				res = append(res, []int{pt.xAxis, currHeight})
			}
		}
	}
	return res
}

type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

// Init define
func (bit *BinaryIndexedTree) Init(capacity int) {
	bit.tree, bit.capacity = make([]int, capacity+1), capacity
}

// Add define
func (bit *BinaryIndexedTree) Add(index int, val int) {
	for ; index > 0; index -= index & -index {
		bit.tree[index] = max(bit.tree[index], val)
	}
}

// Query define
func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for ; index <= bit.capacity; index += index & -index {
		sum = max(sum, bit.tree[index])
	}
	return sum
}

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

// 查询 [left....right] 区间内的值

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

// SegmentCountTree define
type SegmentCountTree struct {
	data, tree  []int
	left, right int
	merge       func(i, j int) int
}

// Init define
func (st *SegmentCountTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper

	data, tree := make([]int, len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree = data, tree
}

// 在 treeIndex 的位置创建 [left....right] 区间的线段树
func (st *SegmentCountTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentCountTree) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentCountTree) rightChild(index int) int {
	return 2*index + 2
}

// 查询 [left....right] 区间内的值

// Query define
func (st *SegmentCountTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// 在以 treeIndex 为根的线段树中 [left...right] 的范围里，搜索区间 [queryLeft...queryRight] 的值，值是计数值
func (st *SegmentCountTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if queryRight < st.data[left] || queryLeft > st.data[right] {
		return 0
	}
	if queryLeft <= st.data[left] && queryRight >= st.data[right] || left == right {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	return st.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight) +
		st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
}

// 更新计数

// UpdateCount define
func (st *SegmentCountTree) UpdateCount(val int) {
	if len(st.data) > 0 {
		st.updateCountInTree(0, 0, len(st.data)-1, val)
	}
}

// 以 treeIndex 为根，更新 [left...right] 区间内的计数
func (st *SegmentCountTree) updateCountInTree(treeIndex, left, right, val int) {
	if val >= st.data[left] && val <= st.data[right] {
		st.tree[treeIndex]++
		if left == right {
			return
		}
		midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
		st.updateCountInTree(rightTreeIndex, midTreeIndex+1, right, val)
		st.updateCountInTree(leftTreeIndex, left, midTreeIndex, val)
	}
}

// 解法二 线段树 Segment Tree，时间复杂度 O(n log n)
func getSkyline1(buildings [][]int) [][]int {
	st, ans, lastHeight, check := SegmentTree{}, [][]int{}, 0, false
	posMap, pos := discretization218(buildings)
	tmp := make([]int, len(posMap))
	st.Init(tmp, func(i, j int) int {
		return max(i, j)
	})
	for _, b := range buildings {
		st.UpdateLazy(posMap[b[0]], posMap[b[1]-1], b[2])
	}
	for i := 0; i < len(pos); i++ {
		h := st.QueryLazy(posMap[pos[i]], posMap[pos[i]])
		if check == false && h != 0 {
			ans = append(ans, []int{pos[i], h})
			check = true
		} else if i > 0 && h != lastHeight {
			ans = append(ans, []int{pos[i], h})
		}
		lastHeight = h
	}
	return ans
}

func discretization218(positions [][]int) (map[int]int, []int) {
	tmpMap, posArray, posMap := map[int]int{}, []int{}, map[int]int{}
	for _, pos := range positions {
		tmpMap[pos[0]]++
		tmpMap[pos[1]-1]++
		tmpMap[pos[1]]++
	}
	for k := range tmpMap {
		posArray = append(posArray, k)
	}
	sort.Ints(posArray)
	for i, pos := range posArray {
		posMap[pos] = i
	}
	return posMap, posArray
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 扫描线 Sweep Line，时间复杂度 O(n log n)
func getSkyline2(buildings [][]int) [][]int {
	size := len(buildings)
	es := make([]E, 0)
	for i, b := range buildings {
		l := b[0]
		r := b[1]
		h := b[2]
		// 1-- enter
		el := NewE(i, l, h, 0)
		es = append(es, el)
		// 0 -- leave
		er := NewE(i, r, h, 1)
		es = append(es, er)
	}
	skyline := make([][]int, 0)
	sort.Slice(es, func(i, j int) bool {
		if es[i].X == es[j].X {
			if es[i].T == es[j].T {
				if es[i].T == 0 {
					return es[i].H > es[j].H
				}
				return es[i].H < es[j].H
			}
			return es[i].T < es[j].T
		}
		return es[i].X < es[j].X
	})
	pq := NewIndexMaxPQ(size)
	for _, e := range es {
		curH := pq.Front()
		if e.T == 0 {
			if e.H > curH {
				skyline = append(skyline, []int{e.X, e.H})
			}
			pq.Enque(e.N, e.H)
		} else {
			pq.Remove(e.N)
			h := pq.Front()
			if curH > h {
				skyline = append(skyline, []int{e.X, h})
			}
		}
	}
	return skyline
}

// E define
type E struct { // 定义一个 event 事件
	N int // number 编号
	X int // x 坐标
	H int // height 高度
	T int // type  0-进入 1-离开
}

// NewE define
func NewE(n, x, h, t int) E {
	return E{
		N: n,
		X: x,
		H: h,
		T: t,
	}
}

// IndexMaxPQ define
type IndexMaxPQ struct {
	items []int
	pq    []int
	qp    []int
	total int
}

// NewIndexMaxPQ define
func NewIndexMaxPQ(n int) IndexMaxPQ {
	qp := make([]int, n)
	for i := 0; i < n; i++ {
		qp[i] = -1
	}
	return IndexMaxPQ{
		items: make([]int, n),
		pq:    make([]int, n+1),
		qp:    qp,
	}
}

// Enque define
func (q *IndexMaxPQ) Enque(key, val int) {
	q.total++
	q.items[key] = val
	q.pq[q.total] = key
	q.qp[key] = q.total
	q.swim(q.total)
}

// Front define
func (q *IndexMaxPQ) Front() int {
	if q.total < 1 {
		return 0
	}
	return q.items[q.pq[1]]
}

// Remove define
func (q *IndexMaxPQ) Remove(key int) {
	rank := q.qp[key]
	q.exch(rank, q.total)
	q.total--
	q.qp[key] = -1
	q.sink(rank)
}

func (q *IndexMaxPQ) sink(n int) {
	for 2*n <= q.total {
		k := 2 * n
		if k < q.total && q.less(k, k+1) {
			k++
		}
		if q.less(k, n) {
			break
		}
		q.exch(k, n)
		n = k
	}
}

func (q *IndexMaxPQ) swim(n int) {
	for n > 1 {
		k := n / 2
		if q.less(n, k) {
			break
		}
		q.exch(n, k)
		n = k
	}
}

func (q *IndexMaxPQ) exch(i, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
	q.qp[q.pq[i]] = i
	q.qp[q.pq[j]] = j
}

func (q *IndexMaxPQ) less(i, j int) bool {
	return q.items[q.pq[i]] < q.items[q.pq[j]]
}

type pair struct{ right, height int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline3(buildings [][]int) [][]int {
    res := [][]int{}
    n := len(buildings)
    boundaries := make([]int, 0, n*2)
    for _, building := range buildings {
        boundaries = append(boundaries, building[0], building[1])
    }
    sort.Ints(boundaries)
    idx := 0
    h := hp{}
    for _, boundary := range boundaries {
        for idx < n && buildings[idx][0] <= boundary {
            heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
            idx++
        }
        for len(h) > 0 && h[0].right <= boundary {
            heap.Pop(&h)
        }
        maxn := 0
        if len(h) > 0 {
            maxn = h[0].height
        }
        if len(res) == 0 || maxn != res[len(res)-1][1] {
            res = append(res, []int{boundary, maxn})
        }
    }
    return res
}

func main() {
	fmt.Printf("getSkyline([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}) = %v\n",getSkyline([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}})) // [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	fmt.Printf("getSkyline([][]int{ {0,2,3},{2,5,3}}) = %v\n",getSkyline([][]int{ {0,2,3},{2,5,3} })) // [[0,3],[5,0]]

	fmt.Printf("getSkyline1([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}) = %v\n",getSkyline1([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}})) // [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	fmt.Printf("getSkyline1([][]int{ {0,2,3},{2,5,3}}) = %v\n",getSkyline1([][]int{ {0,2,3},{2,5,3} })) // [[0,3],[5,0]]

	fmt.Printf("getSkyline2([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}) = %v\n",getSkyline2([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}})) // [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	fmt.Printf("getSkyline2([][]int{ {0,2,3},{2,5,3}}) = %v\n",getSkyline2([][]int{ {0,2,3},{2,5,3} })) // [[0,3],[5,0]]

    fmt.Printf("getSkyline3([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}) = %v\n",getSkyline3([][]int{ {2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}})) // [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	fmt.Printf("getSkyline3([][]int{ {0,2,3},{2,5,3}}) = %v\n",getSkyline3([][]int{ {0,2,3},{2,5,3} })) // [[0,3],[5,0]]

}

