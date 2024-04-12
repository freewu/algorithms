package template

// https://books.halfrost.com/leetcode/ChapterThree/Binary_Indexed_Tree/

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