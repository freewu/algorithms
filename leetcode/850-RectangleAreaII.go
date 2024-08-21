package main

// 850. Rectangle Area II
// You are given a 2D array of axis-aligned rectangles. 
// Each rectangle[i] = [xi1, yi1, xi2, yi2] denotes the ith rectangle where (xi1, yi1) are the coordinates of the bottom-left corner, 
// and (xi2, yi2) are the coordinates of the top-right corner.

// Calculate the total area covered by all rectangles in the plane. 
// Any area covered by two or more rectangles should only be counted once.

// Return the total area. Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/06/rectangle_area_ii_pic.png" />
// Input: rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
// Output: 6
// Explanation: A total area of 6 is covered by all three rectangles, as illustrated in the picture.
// From (1,1) to (2,2), the green and red rectangles overlap.
// From (1,0) to (2,3), all three rectangles overlap.

// Example 2:
// Input: rectangles = [[0,0,1000000000,1000000000]]
// Output: 49
// Explanation: The answer is 10^18 modulo (10^9 + 7), which is 49.

// Constraints:
//     1 <= rectangles.length <= 200
//     rectanges[i].length == 4
//     0 <= xi1, yi1, xi2, yi2 <= 10^9
//     xi1 <= xi2
//     yi1 <= yi2

import "fmt"
import "sort"

func rectangleArea(rectangles [][]int) int {
    res, simpleRects := 0, make([][4]int, 0, len(rectangles))
    var dfs func(cur [4]int, start int)
    dfs = func(cur [4]int, start int) {
        if start == len(simpleRects) {
            simpleRects = append(simpleRects, cur)
            return
        }
        cpart := simpleRects[start]
        if cpart[0] >= cur[2] || cpart[1] >= cur[3] || cpart[2] <= cur[0] || cpart[3] <= cur[1] {
            dfs(cur, start+1)
            return
        }
        if cur[0] < cpart[0] && cpart[0] < cur[2] {
            dfs([4]int{cur[0], cur[1], cpart[0], cur[3]}, start+1)
        }
        if cur[0] < cpart[2] && cpart[2] < cur[2] {
            dfs([4]int{cpart[2], cur[1], cur[2], cur[3]}, start+1)
        }
        if cur[1] < cpart[1] && cpart[1] < cur[3] {
            dfs([4]int{cur[0], cur[1], cur[2], cpart[1]}, start+1)
        }
        if cur[1] < cpart[3] && cpart[3] < cur[3] {
            dfs([4]int{cur[0], cpart[3], cur[2], cur[3]}, start+1)
        }
    }
    for _, v := range rectangles {
        dfs([4]int{v[0], v[1], v[2], v[3]}, 0)
    }
    for _, rectangle := range simpleRects {
        res += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
    }
    return res % 1_000_000_007
}

func rectangleArea1(rectangles [][]int) int {
    n, mod := len(rectangles), 1_000_000_007
    nums, events := make([]int, 2*n), make([][]int, 0, 2*n)
    // 每个矩形会产生 2 个事件
    // x1 处，在 [y1, y2] 上增加了覆盖
    // x2 处，在 [y1, y2] 上减少了覆盖
    for _, v := range rectangles {
        x1, y1, x2, y2 := v[0], v[1], v[2], v[3]
        events = append(events, []int{x1, 1, y1, y2}, []int{x2, -1, y1, y2})
        nums = append(nums, y1, y2)
    }
    discrete := func (nums []int) []int {
        sort.Ints(nums)
        length := 1
        for i := 1; i < len(nums); i++ {
            if nums[i] != nums[i-1] {
                nums[length] = nums[i]
                length++
            }
        }
        return nums[:length]
    }
    // 在 y 方向做离散化
    nums = discrete(nums)
    table := make(map[int]int, len(nums))
    for i, v := range nums {
        table[v] = i
    }
    // 离散化后的数据形如：
    // 0   1   2   3   4   5
    // 100 200 300 350 400 450
    // 节点 0~0 负责的范围为 [100, 200]
    // 节点 1~1 负责的范围为 [200, 300]
    // 节点 0~2 负责的范围为 [100, 350]
    // 也就是要往上走一个
    // 对事件在 x 上方向进行排序，然后开始扫描
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    res, prex, tree := 0, events[0][0], NewSegmentTree(nums)
    for _, e := range events {
        curx, ty, y1, y2 := e[0], e[1], e[2], e[3]
        l, r := table[y1], table[y2]
        res = (res + (curx-prex)*tree.Length()) % mod
        prex = curx
        tree.Add(l, r-1, ty)
    }
    return res
}

func NewSegmentTree(nums []int) *SegmentTree {
    tree := &SegmentTree{
        root:  &Node{},
        lower: 0,
        upper: len(nums) - 1,
    }
    tree.init(tree.root, nums, 0, len(nums)-1)
    return tree
}

type Node struct {
    left, right *Node
    totalL int // 总长度
    length int // 覆盖长度
    count  int // 覆盖次数
}

type SegmentTree struct {
    root         *Node
    lower, upper int
}

func (this *SegmentTree) init(node *Node, nums []int, l, r int) {
    v := nums[r]
    if r + 1 < len(nums) {
        v = nums[r+1]
    }
    node.totalL = v - nums[l]
    if l == r {
        return
    }
    m := l + (r-l)>>1
    node.left, node.right  = &Node{}, &Node{}
    this.init(node.left, nums, l, m)
    this.init(node.right, nums, m+1, r)
}

func (t *SegmentTree) Length() int {
    return t.root.length
}

func (this *SegmentTree) lazy(node *Node, v int) {
	node.count += v
    if node.count > 0 {
        node.length = node.totalL
    } else {
        node.length = 0
        if node.left != nil {
            node.length += node.left.length
        }
        if node.right != nil {
            node.length += node.right.length
        }
    }
}

func (this *SegmentTree) up(node *Node) {
    if node.count == 0 {
        node.length = node.left.length + node.right.length
    } else {
        node.length = node.totalL
    }
}

func (this *SegmentTree) add(jobl, jobr, jobv int, l, r int, node *Node) {
    if jobl <= l && r <= jobr {
        this.lazy(node, jobv)
        return
    }
    m := l + (r-l) >> 1
    if jobl <= m {
        this.add(jobl, jobr, jobv, l, m, node.left)
    }
    if jobr > m {
        this.add(jobl, jobr, jobv, m+1, r, node.right)
    }
    this.up(node)
}

func (this *SegmentTree) Add(l, r, v int) {
    this.add(l, r, v, this.lower, this.upper, this.root)
}


func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/06/rectangle_area_ii_pic.png" />
    // Input: rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
    // Output: 6
    // Explanation: A total area of 6 is covered by all three rectangles, as illustrated in the picture.
    // From (1,1) to (2,2), the green and red rectangles overlap.
    // From (1,0) to (2,3), all three rectangles overlap.
    fmt.Println(rectangleArea([][]int{{0,0,2,2},{1,0,2,3},{1,0,3,1}})) // 6
    // Example 2:
    // Input: rectangles = [[0,0,1000000000,1000000000]]
    // Output: 49
    // Explanation: The answer is 10^18 modulo (10^9 + 7), which is 49.
    fmt.Println(rectangleArea([][]int{{0,0,1000000000,1000000000}})) // 49

    fmt.Println(rectangleArea1([][]int{{0,0,2,2},{1,0,2,3},{1,0,3,1}})) // 6
    fmt.Println(rectangleArea1([][]int{{0,0,1000000000,1000000000}})) // 49
}