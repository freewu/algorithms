package main

// 面试题 16.14. Best Line LCCI
// Given a two-dimensional graph with points on it, find a line which passes the most number of points.

// Assume all the points that passed by the line are stored in list S sorted by their number. 
// You need to return [S[0], S[1]], that is , two points that have smallest number. 
// If there are more than one line that passes the most number of points, choose the one that has the smallest S[0]. 
// If there are more that one line that has the same S[0], choose the one that has smallest S[1].

// Example:
// Input:  [[0,0],[1,1],[1,0],[2,0]]
// Output:  [0,2]
// Explanation:  The numbers of points passed by the line are [0,2,3].

// Note:
//     2 <= len(Points) <= 300
//     len(Points[i]) = 2

import "fmt"

func bestLine(points [][]int) []int {
    n := len(points)
    if n == 2 { return []int{0,1} }
    res, mx := []int{}, 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i, p := range points {
        count := make(map[int][]int)
        for j, q := range points[i+1:] {
            x, y := p[0] - q[0], p[1] - q[1]
            if x == 0 {
                y = 1
            } else if y == 0 {
                x = 1
            } else {
                if y < 0 {
                    x, y = -x, -y
                }
                g := gcd(abs(x), abs(y))
                x /= g
                y /= g
            }
            if len(count[y + x * 20001]) == 0 {
                count[y + x * 20001] = append(count[y + x * 20001],i)
            }
            count[y + x * 20001] = append(count[y + x * 20001], i + j + 1)
        }
        for _, c := range count {
            if len(c) + 1 > mx {
                mx = len(c) + 1
                if len(c) > 1 {
                    res = []int{ c[0], c[1] }
                }
            }
        }
    }
    return res
}

func bestLine1(points [][]int) []int {
    res, mx, n := []int{ 0, 1 }, 2, len(points)
    for i := 0;i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            tmp := []int{ i, j }
            for k := j + 1; k < n; k++ {
                if k == i || k == j {
                    tmp = append(tmp, k)
                    continue 
                }
                if (points[i][0] - points[j][0]) * (points[j][1] - points[k][1]) == (points[j][0] - points[k][0]) * (points[i][1] - points[j][1]) {
                    tmp = append(tmp, k)
                }
            }
            if len(tmp) > mx {
                res, mx = tmp[:2], len(tmp)
            }
        }
    }
    return res
}

func main() {
    // Example:
    // Input:  [[0,0],[1,1],[1,0],[2,0]]
    // Output:  [0,2]
    // Explanation:  The numbers of points passed by the line are [0,2,3].
    fmt.Println(bestLine([][]int{{0,0},{1,1},{1,0},{2,0}})) // [0,2]

    fmt.Println(bestLine([][]int{{-38935,27124},{-39837,19604},{-7086,42194},{-11571,-23257},{115,-23257},{20229,5976},{24653,-18488},{11017,21043},{-9353,16550},{-47076,15237},{-36686,42194},{-17704,1104},{31067,7368},{-20882,42194},{-19107,-10597},{-14898,24506},{-20801,42194},{-52268,40727},{-14042,42194},{-23254,42194},{-30837,-53882},{1402,801},{-33961,-984},{-6411,42194},{-12210,22901},{-8213,-19441},{-26939,20810},{30656,-23257},{-27195,21649},{-33780,2717},{23617,27018},{12266,3608}})) // [2 10]

    fmt.Println(bestLine1([][]int{{0,0},{1,1},{1,0},{2,0}})) // [0,2]
    fmt.Println(bestLine1([][]int{{-38935,27124},{-39837,19604},{-7086,42194},{-11571,-23257},{115,-23257},{20229,5976},{24653,-18488},{11017,21043},{-9353,16550},{-47076,15237},{-36686,42194},{-17704,1104},{31067,7368},{-20882,42194},{-19107,-10597},{-14898,24506},{-20801,42194},{-52268,40727},{-14042,42194},{-23254,42194},{-30837,-53882},{1402,801},{-33961,-984},{-6411,42194},{-12210,22901},{-8213,-19441},{-26939,20810},{30656,-23257},{-27195,21649},{-33780,2717},{23617,27018},{12266,3608}})) // [2 10]
}