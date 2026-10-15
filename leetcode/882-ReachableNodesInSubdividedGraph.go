package main

// 882. Reachable Nodes In Subdivided Graph
// You are given an undirected graph (the "original graph") with n nodes labeled from 0 to n - 1. 
// You decide to subdivide each edge in the graph into a chain of nodes, 
// with the number of new nodes varying between each edge.

// The graph is given as a 2D array of edges where edges[i] = [ui, vi, cnti] indicates 
// that there is an edge between nodes ui and vi in the original graph, 
// and cnti is the total number of new nodes that you will subdivide the edge into. 
// Note that cnti == 0 means you will not subdivide the edge.

// To subdivide the edge [ui, vi], replace it with (cnti + 1) new edges and cnti new nodes. 
// The new nodes are x1, x2, ..., xcnti, and the new edges are [ui, x1], [x1, x2], [x2, x3], ..., [xcnti-1, xcnti], [xcnti, vi].

// In this new graph, you want to know how many nodes are reachable from the node 0, 
// where a node is reachable if the distance is maxMoves or less.

// Given the original graph and maxMoves, 
// return the number of nodes that are reachable from node 0 in the new graph.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/01/origfinal.png" />
// Input: edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
// Output: 13
// Explanation: The edge subdivisions are shown in the image above.
// The nodes that are reachable are highlighted in yellow.

// Example 2:
// Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
// Output: 23

// Example 3:
// Input: edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
// Output: 1
// Explanation: Node 0 is disconnected from the rest of the graph, so only node 0 is reachable.

// Constraints:
//     0 <= edges.length <= min(n * (n - 1) / 2, 10^4)
//     edges[i].length == 3
//     0 <= ui < vi < n
//     There are no multiple edges in the graph.
//     0 <= cnti <= 10^4
//     0 <= maxMoves <= 10^9
//     1 <= n <= 3000

import "fmt"
import "container/heap"
import "sort"

func reachableNodes(edges [][]int, maxMoves int, n int) int {
    const MX = 3010
    const INF = 0x3f3f3f3f
    res, dist, graphs := 0, [MX]int{}, [MX]map[int]int{}
    for i := 0; i < n; i++ {
        dist[i] = INF
        graphs[i] = make(map[int]int)
    }
    dist[0] = 0
    for i := 0; i < len(edges); i++ {
        e := edges[i]
        graphs[e[0]][e[1]] = e[2] + 1
        graphs[e[1]][e[0]] = e[2] + 1
    }
    queue, set := []int{0}, [MX]bool{}
    set[0] = true
    for len(queue) > 0 {
        t := queue[0]
        queue = queue[1:]
        set[t] = false
        for j, w := range graphs[t] {
            if dist[j] > dist[t] + w {
                dist[j] = dist[t] + w
                if !set[j] {
                    queue= append(queue, j)
                    set[j] = true
                }
            }
        }
    }
    for i := 0; i < n; i++ {
        if dist[i] <= maxMoves {
            res++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, e := range edges {
        x := max(maxMoves - dist[e[0]], 0)
        y := max(maxMoves - dist[e[1]], 0)
        res += min(e[2], x + y)
    }
    return res
}

type Neigbour struct{ to, weight int }

type Pair struct{ x, dist int }
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func reachableNodes1(edges [][]int, maxMoves int, n int) int {
    res, g := 0, make([][]Neigbour, n)
    for _, edge := range edges {
        u, v, cnt := edge[0], edge[1], edge[2]
        g[u] = append(g[u], Neigbour{v, cnt + 1})
        g[v] = append(g[v], Neigbour{u, cnt + 1})
    }
    dijkstra := func(g [][]Neigbour, start int) []int {
        dist, inf := make([]int, len(g)), 1 << 31
        for i := range dist {
            dist[i] = inf
        }
        dist[start] = 0
        var mh = &MinHeap{Pair{start, 0}}
        heap.Init(mh)
        for mh.Len() > 0 {
            head := heap.Pop(mh).(Pair)
            x := head.x
            weight := head.dist
            if dist[x] < weight {
                continue
            }
            for _, y := range g[x] {
                to := y.to
                if d := dist[x] + y.weight; d < dist[to] {
                    dist[to] = d
                    heap.Push(mh, Pair{to, d})
                }
            }
        }
        return dist
    }
    dist := dijkstra(g, 0)
    for i := range dist {
        if dist[i] <= maxMoves {
            res++
        }
    }
    for _, e := range edges {
        u, v, cnt := e[0], e[1], e[2]
        a := max(maxMoves-dist[u], 0)
        b := max(maxMoves-dist[v], 0)
        res += min(a+b, cnt)
    }
    return res
}

// == 图构建 & 绘制输出 start ======================================================================
// printGraph 依据边表绘制无向图的 ASCII 向量图：按 BFS 分层，根在上、同层横向铺开，
// 层间用制表符线（┌ ┬ ┐ └ ┴ ┘ ─ │ ┼）连接，边上标注权重。
// 每条边形如 [u, v, w]：第 3 个值是权重；若某条边只有 2 个元素，则视为无权重、不标注。
func printGraph(edges [][]int, n int) {
    if n <= 0 {
        fmt.Println("(empty graph)")
        return
    }
    // itoa 手写整数转字符串，替代 strconv.Itoa
    itoa := func(v int) string {
        if v == 0 {
            return "0"
        }
        neg := v < 0
        u := uint(v)
        if neg {
            u = -u // 补码取绝对值，v 为最小整数时同样安全
        }
        buf := [24]byte{}
        i := len(buf)
        for u > 0 {
            i--
            buf[i] = byte('0' + u%10)
            u /= 10
        }
        if neg {
            i--
            buf[i] = '-'
        }
        return string(buf[i:])
    }

    type arc struct {
        v, w int
        hasW bool
    }
    // 邻接表：adj[u] 存 {邻居, 权重, 是否带权重}
    adj := make([][]arc, n)
    for _, e := range edges {
        if len(e) < 2 {
            continue
        }
        u, v := e[0], e[1]
        if u < 0 || u >= n || v < 0 || v >= n || u == v {
            continue
        }
        w, hasW := 0, false
        if len(e) >= 3 { // 只有存在第 3 个值时才算权重
            w, hasW = e[2], true
        }
        adj[u] = append(adj[u], arc{v: v, w: w, hasW: hasW})
        adj[v] = append(adj[v], arc{v: u, w: w, hasW: hasW})
    }

    // 分层：每个连通分量各自 BFS 定层，分量之间用空行隔开
    lvl := make([]int, n)
    for i := range lvl {
        lvl[i] = -1
    }
    var levels [][]int // levels[L] 为第 L 层节点（按编号升序）
    for start := 0; start < n; start++ {
        if lvl[start] != -1 {
            continue
        }
        base := len(levels)
        lvl[start] = base
        cur := []int{start}
        for L := 0; len(cur) > 0; L++ {
            sort.Ints(cur)
            for len(levels) <= base+L {
                levels = append(levels, nil)
            }
            levels[base+L] = append(levels[base+L], cur...)
            next := []int{}
            for _, u := range cur {
                for _, a := range adj[u] {
                    if lvl[a.v] == -1 {
                        lvl[a.v] = base + L + 1
                        next = append(next, a.v)
                    }
                }
            }
            cur = next
        }
    }

    const space = 12 // 同层相邻节点的中心列间距，留出权重标注的位置
    // 节点标签统一写成 (num)，和边权（纯数字）区分开
    lbl := func(v int) string { return "(" + itoa(v) + ")" }
    margin := 0 // 最宽标签的半径，保证标签不会越过左右边界
    for i := 0; i < n; i++ {
        if w := len(lbl(i))/2 + 1; w > margin {
            margin = w
        }
    }
    col := make([]int, n)
    for _, nodes := range levels {
        for i, u := range nodes {
            col[u] = margin + i*space
        }
    }

    // 层间连线需要一个“总线”行，行数 = 该层有跨层边的父节点个数
    buses := make([]int, len(levels))
    for L := 0; L+1 < len(levels); L++ {
        seen := map[int]bool{}
        for _, u := range levels[L] {
            for _, a := range adj[u] {
                if lvl[a.v] == L+1 {
                    seen[u] = true
                }
            }
        }
        buses[L] = len(seen)
    }

    // 计算每层节点行所在行号
    const vSpace = 1 // 层间上下各额外留出的行数，竖线会顺势拉长；想更高就加大
    rowY := make([]int, len(levels))
    y := 0
    for L := range levels {
        rowY[L] = y
        if L+1 < len(levels) {
            if buses[L] == 0 {
                y += 2*vSpace + 1 // 分量之间只留空行
            } else {
                y += buses[L] + 2*vSpace + 3 // 上留白 + 竖线 + 总线 + 落线 + 下留白
            }
        } else {
            y++
        }
    }
    height := y
    maxCol := 0
    for i := range col {
        if col[i] > maxCol {
            maxCol = col[i]
        }
    }
    width := maxCol + margin + 8

    // grid 存普通字符（节点值、权重），bs 存连线方向位，最后合成
    grid := make([][]rune, height)
    for i := range grid {
        grid[i] = make([]rune, width)
        for j := range grid[i] {
            grid[i][j] = ' '
        }
    }
    const (
        bU uint8 = 1 << iota // 上
        bD                   // 下
        bL                   // 左
        bR                   // 右
    )
    bs := make([][]uint8, height)
    for i := range bs {
        bs[i] = make([]uint8, width)
    }
    put := func(row, c int, s string) {
        for i := 0; i < len(s); i++ {
            if p := c + i; row >= 0 && row < height && p >= 0 && p < width {
                grid[row][p] = rune(s[i])
            }
        }
    }
    addBit := func(row, c int, bits uint8) {
        if row >= 0 && row < height && c >= 0 && c < width {
            bs[row][c] |= bits
        }
    }

    // 节点编号写成 (num) 居中写到所在层的节点行
    for i := 0; i < n; i++ {
        s := lbl(i)
        put(rowY[lvl[i]], col[i]-len(s)/2, s)
    }

    // 同层边：直接在节点行画横线，权重居中覆盖
    for _, e := range edges {
        if len(e) < 2 {
            continue
        }
        u, v := e[0], e[1]
        if u < 0 || u >= n || v < 0 || v >= n || u == v || lvl[u] != lvl[v] {
            continue
        }
        if col[u] > col[v] {
            u, v = v, u
        }
        l := col[u] + len(lbl(u))/2 + 1 // 左节点标签右侧
        r := col[v] - len(lbl(v))/2 - 1 // 右节点标签左侧
        if r < l {
            r = l
        }
        row := rowY[lvl[u]]
        for c := l; c <= r; c++ {
            grid[row][c] = '─'
        }
        if len(e) >= 3 {
            s := itoa(e[2])
            st := (l+r)/2 - len(s)/2
            if st < l {
                st = l
            }
            if st+len(s)-1 > r {
                st = r - len(s) + 1
            }
            put(row, st, s)
        }
    }

    // 跨层边：父节点先竖线下到总线，再横向铺开，最后竖线落到孩子
    for L := 0; L+1 < len(levels); L++ {
        if buses[L] == 0 {
            continue
        }
        type kid struct {
            c     int
            label string
        }
        groups := map[int][]kid{}
        var parents []int
        for _, u := range levels[L] {
            for _, a := range adj[u] {
                if lvl[a.v] != L+1 {
                    continue
                }
                if _, ok := groups[u]; !ok {
                    parents = append(parents, u)
                }
                k := kid{c: col[a.v]}
                if a.hasW {
                    k.label = itoa(a.w)
                }
                groups[u] = append(groups[u], k)
            }
        }
        sort.Slice(parents, func(i, j int) bool { return col[parents[i]] < col[parents[j]] })

        stem := rowY[L] + 1                        // 竖线起点，紧接节点行
        down := stem + len(parents) + 2*vSpace + 1 // 落线终点（下一节点行的上一行）
        for idx, u := range parents {
            bus := stem + 1 + vSpace + idx // 该父节点的总线行（上方留 vSpace 行竖线）
            cp := col[u]
            kids := groups[u]
            lo, hi := cp, cp
            for _, k := range kids {
                if k.c < lo {
                    lo = k.c
                }
                if k.c > hi {
                    hi = k.c
                }
            }
            for r := stem; r < bus; r++ { // 父节点到总线的竖线
                addBit(r, cp, bU|bD)
            }
            for c := lo; c <= hi; c++ { // 总线：端点只朝内一侧延伸
                switch {
                case lo == hi:
                case c == lo:
                    addBit(bus, c, bR)
                case c == hi:
                    addBit(bus, c, bL)
                default:
                    addBit(bus, c, bL|bR)
                }
            }
            addBit(bus, cp, bU) // 父节点向下接入总线
            for _, k := range kids { // 孩子接入总线并向下落线
                addBit(bus, k.c, bD)
                for r := bus + 1; r <= down; r++ {
                    addBit(r, k.c, bU|bD)
                }
            }
        }
        // 权重单独占一行，写在落线右侧一格，上下都还有连线，不会和节点标签 (num) 混在一起
        lastBus := stem + vSpace + len(parents) // 该层最后一条总线所在行
        wr := down - 1
        if wr <= lastBus {
            wr = lastBus + 1
        }
        for _, u := range parents {
            for _, k := range groups[u] {
                if k.label != "" {
                    put(wr, k.c+1, k.label)
                }
            }
        }
    }

    // 方向位合成制表符
    box := func(b uint8) rune {
        switch b {
        case bU | bD:
            return '│'
        case bL | bR:
            return '─'
        case bU | bL | bR:
            return '┴'
        case bD | bL | bR:
            return '┬'
        case bU | bD | bL:
            return '┤'
        case bU | bD | bR:
            return '├'
        case bU | bD | bL | bR:
            return '┼'
        case bD | bR:
            return '┌'
        case bD | bL:
            return '┐'
        case bU | bL:
            return '┘'
        case bU | bR:
            return '└'
        case bU, bD:
            return '│'
        case bL, bR:
            return '─'
        }
        return ' '
    }
    // 逐行合成字符，并量出整幅图的左右边界
    lines := make([][]rune, height)
    left, right := width, -1
    for r := 0; r < height; r++ {
        line := make([]rune, width)
        for c := 0; c < width; c++ {
            if b := bs[r][c]; b != 0 {
                line[c] = box(b)
            } else {
                line[c] = grid[r][c]
            }
        }
        lines[r] = line
        for c := 0; c < width; c++ {
            if line[c] != ' ' {
                if c < left {
                    left = c
                }
                if c > right {
                    right = c
                }
            }
        }
    }
    if right < left {
        left, right = 0, 0
    }
    // 整幅图加完全相同的前导空格（不能按每行各自居中，否则竖线会错位），
    // 使其在 centerWidth 列里水平居中
    const centerWidth = 80
    pad := (centerWidth - (right - left + 1)) / 2
    if pad < 0 {
        pad = 0
    }
    for r := 0; r < height; r++ {
        line := lines[r]
        end := width
        for end > left && line[end-1] == ' ' {
            end--
        }
        if end <= left { // 全空行
            fmt.Println()
            continue
        }
        fmt.Printf("%*s%s\n", pad, "", string(line[left:end]))
    }
}
// == 图构建 & 绘制输出 end ========================================================================

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/01/origfinal.png" />
    // Input: edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
    // Output: 13
    // Explanation: The edge subdivisions are shown in the image above.
    // The nodes that are reachable are highlighted in yellow.
    graph1 := [][]int{{0,1,10},{0,2,1},{1,2,2}}
    printGraph(graph1, len(graph1))
    fmt.Println(reachableNodes(graph1, 6, 3)) // 13
    // Example 2:
    // Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
    // Output: 23
    graph2 := [][]int{{0,1,4},{1,2,6},{0,2,8},{1,3,1}}
    printGraph(graph2, len(graph2))
    fmt.Println(reachableNodes(graph2, 10, 4)) // 23
    // Example 3:
    // Input: edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
    // Output: 1
    // Explanation: Node 0 is disconnected from the rest of the graph, so only node 0 is reachable.
    graph3 := [][]int{{1,2,4},{1,4,5},{1,3,1},{2,3,4},{3,4,5}}
    printGraph(graph3, len(graph3))
    fmt.Println(reachableNodes([][]int{{1,2,4},{1,4,5},{1,3,1},{2,3,4},{3,4,5}}, 12, 5)) // 1

    fmt.Println(reachableNodes1(graph1, 6, 3)) // 13
    fmt.Println(reachableNodes1(graph2, 10, 4)) // 23
    fmt.Println(reachableNodes1(graph3, 12, 5)) // 1
}