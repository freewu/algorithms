package template

// https://books.halfrost.com/leetcode/ChapterThree/Segment_Tree/

// BinaryIndexedTree define
type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

// Init define
func (bit *BinaryIndexedTree) Init(nums []int) {
    bit.tree, bit.capacity = make([]int, len(nums)+1), len(nums)+1
    for i := 1; i <= len(nums); i++ {
        bit.tree[i] += nums[i-1]
        for j := i - 2; j >= i-lowbit(i); j-- {
            bit.tree[i] += nums[j]
        }
    }
}

func lowbit(x int) int {
    return x & -x
}

// Add define
func (bit *BinaryIndexedTree) Add(index int, val int) {
    for index <= bit.capacity {
        bit.tree[index] += val
        index += lowbit(index)
    }
}

// Query define
func (bit *BinaryIndexedTree) Query(index int) int {
    sum := 0
    for index >= 1 {
        sum += bit.tree[index]
        index -= lowbit(index)
    }
    return sum
}