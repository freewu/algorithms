package main

// 3454. Separate Squares II
// You are given a 2D integer array squares. 
// Each squares[i] = [xi, yi, li] represents the coordinates of the bottom-left point 
// and the side length of a square parallel to the x-axis.

// Find the minimum y-coordinate value of a horizontal line 
// such that the total area covered by squares above the line equals the total area covered by squares below the line.

// Answers within 10-5 of the actual answer will be accepted.

// Note: Squares may overlap. Overlapping areas should be counted only once in this version.

// Example 1:
// Input: squares = [[0,0,1],[2,2,1]]
// Output: 1.00000
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/15/4065example1drawio.png" />
// Any horizontal line between y = 1 and y = 2 results in an equal split, with 1 square unit above and 1 square unit below. 
// The minimum y-value is 1.

// Example 2:
// Input: squares = [[0,0,2],[1,1,1]]
// Output: 1.00000
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/15/4065example2drawio.png" />
// Since the blue square overlaps with the red square, it will not be counted again. 
// Thus, the line y = 1 splits the squares into two equal parts.

// Constraints:
//     1 <= squares.length <= 5 * 10^4
//     squares[i] = [xi, yi, li]
//     squares[i].length == 3
//     0 <= xi, yi <= 10^9
//     1 <= li <= 10^9
//     The total area of all the squares will not exceed 10^15.

import "fmt"
import "sort"

// type SegmentTree struct {
//     xs     []int
//     n      int
//     count  []int
//     covered []int
// }

// func NewSegmentTree(xs []int) *SegmentTree {
//     n := len(xs) - 1
//     return &SegmentTree{
//         xs:     xs,
//         n:      n,
//         count:  make([]int, 4*n),
//         covered: make([]int, 4*n),
//     }
// }

// func (st *SegmentTree) update(qleft, qright, qval, left, right, pos int) {
//     if st.xs[right+1] <= qleft || st.xs[left] >= qright { return }
//     if qleft <= st.xs[left] && st.xs[right+1] <= qright {
//         st.count[pos] += qval
//     } else {
//         mid := (left + right) / 2
//         st.update(qleft, qright, qval, left, mid, pos*2+1)
//         st.update(qleft, qright, qval, mid+1, right, pos*2+2)
//     }
//     if st.count[pos] > 0 {
//         st.covered[pos] = st.xs[right+1] - st.xs[left]
//     } else {
//         if left == right {
//             st.covered[pos] = 0
//         } else {
//             st.covered[pos] = st.covered[pos*2+1] + st.covered[pos*2+2]
//         }
//     }
// }

// func (st *SegmentTree) query() int {
//     return st.covered[0]
// }

// func separateSquares(squares [][]int) float64 {
//     events := make([][4]int, 0)
//     xsSet := make(map[int]bool)
//     for _, square := range squares {
//         x, y, l := square[0], square[1], square[2]
//         events = append(events, [4]int{y, 1, x, x + l})
//         events = append(events, [4]int{y + l, -1, x, x + l})
//         xsSet[x] = true
//         xsSet[x+l] = true
//     }
//     xs := make([]int, 0, len(xsSet))
//     for x := range xsSet {
//         xs = append(xs, x)
//     }
//     sort.Ints(xs)
//     segTree := NewSegmentTree(xs)
//     sort.Slice(events, func(i, j int) bool {
//         return events[i][0] < events[j][0]
//     })
//     // First sweep: compute total union area.
//     totalArea := 0.0
//     prevY := events[0][0]
//     for _, event := range events {
//         y, start, xl, xr := event[0], event[1], event[2], event[3]
//         totalArea += float64(segTree.query()) * float64(y-prevY)
//         segTree.update(xl, xr, start, 0, segTree.n-1, 0)
//         prevY = y
//     }
//     // Second sweep: find the minimal y where the area below equals half_area.
//     segTree = NewSegmentTree(xs) // Reinitialize segment tree
//     currArea := 0.0
//     prevY = events[0][0]
//     for _, event := range events {
//         y, start, xl, xr := event[0], event[1], event[2], event[3]
//         combinedWidth := segTree.query()
//         if currArea + float64(combinedWidth)*float64(y-prevY) >= totalArea/2.0 {
//             return float64(prevY) + float64(totalArea / 2.0 - currArea) / float64(combinedWidth)
//         }
//         currArea += float64(combinedWidth) * float64(y-prevY)
//         segTree.update(xl, xr, start, 0, segTree.n-1, 0)
//         prevY = y
//     }
//     return 0.0
// }

type SegmentTree struct {
    count   []int
    covered []int
    xs      []int
    n       int
}

func NewSegmentTree(xs []int) *SegmentTree {
    n := len(xs) - 1
    return &SegmentTree{
        count:   make([]int, 4*n),
        covered: make([]int, 4*n),
        xs:      xs,
        n:       n,
    }
}

func (st *SegmentTree) modify(qleft, qright, qval, left, right, pos int) {
    if st.xs[right+1] <= qleft || st.xs[left] >= qright {
        return
    }
    if qleft <= st.xs[left] && st.xs[right+1] <= qright {
        st.count[pos] += qval
    } else {
        mid := (left + right) / 2
        st.modify(qleft, qright, qval, left, mid, pos*2+1)
        st.modify(qleft, qright, qval, mid+1, right, pos*2+2)
    }
    if st.count[pos] > 0 {
        st.covered[pos] = st.xs[right+1] - st.xs[left]
    } else {
        if left == right {
            st.covered[pos] = 0
        } else {
            st.covered[pos] = st.covered[pos*2+1] + st.covered[pos*2+2]
        }
    }
}

func (st *SegmentTree) Update(qleft, qright, qval int) {
    st.modify(qleft, qright, qval, 0, st.n-1, 0)
}

func (st *SegmentTree) Query() int {
    return st.covered[0]
}

func separateSquares(squares [][]int) float64 {
    // 存储事件: (y坐标, 类型, 左边界, 右边界)
    type Event struct {
        y, delta, xl, xr int
    }
    events := []Event{}
    xsSet := make(map[int]bool)
    for _, sq := range squares {
        x, y, l := sq[0], sq[1], sq[2]
        xr := x + l
        events = append(events, Event{y, 1, x, xr})
        events = append(events, Event{y + l, -1, x, xr})
        xsSet[x] = true
        xsSet[xr] = true
    }
    // 按y坐标排序事件
    sort.Slice(events, func(i, j int) bool {
        return events[i].y < events[j].y
    })
    // 离散化坐标
    xs := make([]int, 0, len(xsSet))
    for x := range xsSet {
        xs = append(xs, x)
    }
    sort.Ints(xs)
    // 初始化线段树
    segTree := NewSegmentTree(xs)
    psum := []float64{}
    widths := []int{}
    totalArea := 0.0
    prev := events[0].y
    // 扫描：计算总面积和记录中间状态
    for _, event := range events {
        y, delta, xl, xr := event.y, event.delta, event.xl, event.xr
        length := segTree.Query()
        totalArea += float64(length) * float64(y-prev)
        segTree.Update(xl, xr, delta)
        // 记录前缀和和宽度
        psum = append(psum, totalArea)
        widths = append(widths, segTree.Query())
        prev = y
    }
    // 计算目标面积（向上取整的一半）
    target := int64(totalArea + 1) / 2
    // 二分查找第一个大于等于target的位置
    i := sort.Search(len(psum), func(i int) bool {
        return psum[i] >= float64(target)
    })
    i--
    // 获取对应的面积、宽度和高度
    area, width, height := psum[i],  widths[i],events[i].y
    return float64(height) + (totalArea - area * 2) / (float64(width) * 2.0)
}

// Time Limit Exceeded 695 / 763 
func separateSquares1(squares [][]int) float64 {
    type event struct {
        y, x1, x2 int
        typ       int
    }
    // QuickSort for []int
    var quickSortInts func([]int)
    quickSortInts = func(a []int) {
        if len(a) < 2 {
            return
        }
        var partition func(int, int) int
        partition = func(left, right int) int {
            pivot := a[right]
            i := left
            for j := left; j < right; j++ {
                if a[j] < pivot {
                    a[i], a[j] = a[j], a[i]
                    i++
                }
            }
            a[i], a[right] = a[right], a[i]
            return i
        }
        var qs func(int, int)
        qs = func(l, r int) {
            if l >= r {
                return
            }
            m := partition(l, r)
            qs(l, m-1)
            qs(m+1, r)
        }
        qs(0, len(a)-1)
    }
    // QuickSort for []event
    var quickSortEvents func([]event)
    quickSortEvents = func(e []event) {
        if len(e) < 2 {
            return
        }
        var partitionE func(int, int) int
        partitionE = func(left, right int) int {
            pivot := e[right]
            i := left
            for j := left; j < right; j++ {
                if e[j].y < pivot.y || (e[j].y == pivot.y && e[j].typ < pivot.typ) {
                    e[i], e[j] = e[j], e[i]
                    i++
                }
            }
            e[i], e[right] = e[right], e[i]
            return i
        }
        var qsE func(int, int)
        qsE = func(l, r int) {
            if l >= r {
                return
            }
            m := partitionE(l, r)
            qsE(l, m-1)
            qsE(m+1, r)
        }
        qsE(0, len(e)-1)
    }
    n := len(squares)
    events := make([]event, 0, 2*n)
    xs := make([]int, 0, 2*n)
    for _, s := range squares {
        x1, y1, l := s[0], s[1], s[2]
        x2 := x1 + l
        y2 := y1 + l
        events = append(events, event{y1, x1, x2, +1})
        events = append(events, event{y2, x1, x2, -1})
        xs = append(xs, x1, x2)
    }
    // Coordinate compression
    quickSortInts(xs)
    uniqueX := make([]int, 0, len(xs))
    prev := -1
    for _, v := range xs {
        if v != prev {
            uniqueX = append(uniqueX, v)
            prev = v
        }
    }
    xIdx := make(map[int]int, len(uniqueX))
    for i, v := range uniqueX {
        xIdx[v] = i
    }
    m := len(uniqueX) - 1
    if m < 1 {
        // All squares have the same x-range => total area 0 => return min y
        minY := float64(1e15)
        for _, s := range squares {
            if float64(s[1]) < minY {
                minY = float64(s[1])
            }
        }
        return minY
    }
    // Segment tree structures
    coverCount := make([]int, 4*m)
    coverLen := make([]int64, 4*m)
    var recBuildLen func(idx, s, e int)
    recBuildLen = func(idx, s, e int) {
        if coverCount[idx] > 0 {
            coverLen[idx] = int64(uniqueX[e+1] - uniqueX[s])
        } else {
            if s == e {
                coverLen[idx] = 0
            } else {
                coverLen[idx] = coverLen[idx<<1] + coverLen[(idx<<1)+1]
            }
        }
    }
    var updateCoverage func(idx, start, end, l, r, val int)
    updateCoverage = func(idx, start, end, l, r, val int) {
        if r < start || end < l {
            return
        }
        if l <= start && end <= r {
            coverCount[idx] += val
            recBuildLen(idx, start, end)
            return
        }
        mid := (start + end) >> 1
        updateCoverage(idx<<1, start, mid, l, r, val)
        updateCoverage((idx<<1)+1, mid+1, end, l, r, val)
        recBuildLen(idx, start, end)
    }
    var update func(l, r, val int)
    update = func(l, r, val int) {
        // update [l, r) in x index
        updateCoverage(1, 0, m-1, l, r-1, val)
    }
    var getLength func() float64
    getLength = func() float64 {
        return float64(coverLen[1])
    }
    // Build initial
    var build func(idx, s, e int)
    build = func(idx, s, e int) {
        coverCount[idx] = 0
        if s == e {
            coverLen[idx] = 0
            return
        }
        mid := (s + e) >> 1
        build(idx<<1, s, mid)
        build((idx<<1)+1, mid+1, e)
        coverLen[idx] = 0
    }
    build(1, 0, m-1)
    // Sort events by y
    quickSortEvents(events)
    // Sweep line
    partialArea := make([]float64, len(events)+1)
    prevY := 0
    curCoverage := 0.0
    for i, e := range events {
        y := e.y
        dy := float64(y - prevY)
        if dy > 0 {
            partialArea[i+1] = partialArea[i] + curCoverage*dy
        } else {
            partialArea[i+1] = partialArea[i]
        }
        update(xIdx[e.x1], xIdx[e.x2], e.typ)
        curCoverage = getLength()
        prevY = y
    }
    totalArea := partialArea[len(events)]
    if totalArea <= 0 {
        minY := float64(1e15)
        for _, s := range squares {
            if float64(s[1]) < minY {
                minY = float64(s[1])
            }
        }
        return minY
    }
    half := totalArea / 2.0
    // Find where partialArea crosses half
    for i := 0; i < len(events); i++ {
        if partialArea[i] == half {
            if i > 0 {
                return float64(events[i-1].y)
            }
        }
        if partialArea[i] < half && half < partialArea[i+1] {
            delta := half - partialArea[i]
            var y0 int
            if i == 0 {
                y0 = 0
            } else {
                y0 = events[i-1].y
            }
            y1 := events[i].y
            fullSliceArea := partialArea[i+1] - partialArea[i]
            fullSliceHeight := float64(y1 - y0)
            if fullSliceHeight > 1e-15 {
                coverage := fullSliceArea / fullSliceHeight
                heightNeeded := delta / coverage
                return float64(y0) + heightNeeded
            }
            return float64(y0)
        }
    }
    return float64(events[len(events)-1].y)
}

func main() {
    // Example 1:
    // Input: squares = [[0,0,1],[2,2,1]]
    // Output: 1.00000
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/06/4062example1drawio.png" />
    // Any horizontal line between y = 1 and y = 2 will have 1 square unit above it and 1 square unit below it. 
    // The lowest option is 1.
    fmt.Println(separateSquares([][]int{{0,0,1},{2,2,1}})) // 1.00000
    // Example 2:
    // Input: squares = [[0,0,2],[1,1,1]]
    // Output: 1.00000
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/15/4065example2drawio.png" />
    // Since the blue square overlaps with the red square, it will not be counted again. 
    // Thus, the line y = 1 splits the squares into two equal parts.
    fmt.Println(separateSquares([][]int{{0,0,2},{1,1,1}})) // 1.00000

    fmt.Println(separateSquares1([][]int{{0,0,1},{2,2,1}})) // 1.00000
    fmt.Println(separateSquares1([][]int{{0,0,2},{1,1,1}})) // 1.00000
}