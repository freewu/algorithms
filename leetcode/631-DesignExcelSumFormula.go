package main

// 631. Design Excel Sum Formula
// Design the basic function of Excel and implement the function of the sum formula.
// Implement the Excel class:
//     Excel(int height, char width) 
//         Initializes the object with the height and the width of the sheet. 
//         The sheet is an integer matrix mat of size height x width 
//         with the row index in the range [1, height] and the column index in the range ['A', width]. 
//         All the values should be zero initially.
//     void set(int row, char column, int val) 
//         Changes the value at mat[row][column] to be val.
//     int get(int row, char column) 
//         Returns the value at mat[row][column].
//     int sum(int row, char column, List<String> numbers) 
//         Sets the value at mat[row][column] to be the sum of cells represented by numbers 
//         and returns the value at mat[row][column]. This sum formula should exist 
//         until this cell is overlapped by another value or another sum formula. numbers[i] could be on the format:
//             "ColRow" that represents a single cell.
//                 For example, "F7" represents the cell mat[7]['F'].
//             "ColRow1:ColRow2" that represents a range of cells. The range will always be a rectangle where "ColRow1" represent the position of the top-left cell, and "ColRow2" represents the position of the bottom-right cell.
//                 For example, "B3:F7" represents the cells mat[i][j] for 3 <= i <= 7 and 'B' <= j <= 'F'.
    
// Note: You could assume that there will not be any circular sum reference.
//     For example, mat[1]['A'] == sum(1, "B") and mat[1]['B'] == sum(1, "A").

// Example 1:
// Input
// ["Excel", "set", "sum", "set", "get"]
// [[3, "C"], [1, "A", 2], [3, "C", ["A1", "A1:B2"]], [2, "B", 2], [3, "C"]]
// Output
// [null, null, 4, null, 6]
// Explanation
// Excel excel = new Excel(3, "C");
//  // construct a 3*3 2D array with all zero.
//  //   A B C
//  // 1 0 0 0
//  // 2 0 0 0
//  // 3 0 0 0
// excel.set(1, "A", 2);
//  // set mat[1]["A"] to be 2.
//  //   A B C
//  // 1 2 0 0
//  // 2 0 0 0
//  // 3 0 0 0
// excel.sum(3, "C", ["A1", "A1:B2"]); // return 4
//  // set mat[3]["C"] to be the sum of value at mat[1]["A"] and the values sum of the rectangle range whose top-left cell is mat[1]["A"] and bottom-right cell is mat[2]["B"].
//  //   A B C
//  // 1 2 0 0
//  // 2 0 0 0
//  // 3 0 0 4
// excel.set(2, "B", 2);
//  // set mat[2]["B"] to be 2. Note mat[3]["C"] should also be changed.
//  //   A B C
//  // 1 2 0 0
//  // 2 0 2 0
//  // 3 0 0 6
// excel.get(3, "C"); // return 6
 
// Constraints:
//     1 <= height <= 26
//     'A' <= width <= 'Z'
//     1 <= row <= height
//     'A' <= column <= width
//     -100 <= val <= 100
//     1 <= numbers.length <= 5
//     numbers[i] has the format "ColRow" or "ColRow1:ColRow2".
//     At most 100 calls will be made to set, get, and sum.

import "fmt"
import "strconv"
import "strings"

type Excel struct {
    h            int
    w            int
    matrix       [][]int
    dictPosToSum map[int]map[int]int // 点对应的 sum
    dictSumPos   map[int]int         // sum 点坐标
}

func b2i(b byte) int {
    return int(b - 'A')
}

func posParser(s string) []int {
    r, _ := strconv.Atoi(s[1:])
    c := b2i(s[0])
    return []int{r - 1, c}
}

func Constructor(H int, W byte) Excel {
    excel := Excel{}
    excel.matrix = make([][]int, H)
    for i := 0; i < H; i++ {
        excel.matrix[i] = make([]int, b2i(W)+1)
    }
    excel.dictPosToSum = make(map[int]map[int]int)
    excel.dictSumPos = make(map[int]int)
    excel.h = H
    excel.w = b2i(W) + 1
    return excel
}

func getDictKeys(dict map[int]int) []int {
    res := make([]int, 0)
    for k, v := range dict {
        for i := 0; i < v; i++ {
            res = append(res, k)
        }
    }
    return res
}

func (this *Excel) p2i(pos []int) int {
    return this.w*pos[0] + pos[1]
}

func (this *Excel) i2p(idx int) []int {
    return []int{idx / this.w, idx % this.w}
}

func (this *Excel) Set(r int, c byte, v int) {
    pSum := []int{r - 1, b2i(c)}
    idx := this.p2i(pSum)
    if _, ok := this.dictSumPos[idx]; ok { // cover this point
        this.removeSumFunc(idx)
    }
    d := v - this.matrix[pSum[0]][pSum[1]]
    this.matrix[r-1][b2i(c)] = v
    if _, ok := this.dictPosToSum[idx]; !ok {
        return
    }
    queue := [][]int{getDictKeys(this.dictPosToSum[idx])}
    for len(queue) > 0 {
        qNew := make([][]int, 0)
        for _, q := range queue {
            for _, qq := range q {
                pos := this.i2p(qq)
                this.matrix[pos[0]][pos[1]] += d
                qNew = append(qNew, getDictKeys(this.dictPosToSum[qq]))
            }
        }
        queue = qNew
    }
}

func (this *Excel) Get(r int, c byte) int {
    return this.matrix[r-1][b2i(c)]
}

func (this *Excel) Sum(r int, c byte, strs []string) int {
    sum, pSum := 0,[]int{r - 1, b2i(c)}
    idx := this.p2i(pSum)
    if _, ok := this.dictSumPos[idx]; ok { // cover this point
        this.removeSumFunc(idx)
    }

    for _, s := range strs {
        strSep := strings.Split(s, ":")
        if len(strSep[0]) == 0 {
            continue
        }
        posBeg, posEnd := posParser(strSep[0]), posParser(strSep[0])
        if len(strSep) > 1 {
            if len(strSep[1]) == 0 {
                continue
            }
            posEnd = posParser(strSep[1])
        }
        for i := posBeg[0]; i <= posEnd[0]; i++ {
            for j := posBeg[1]; j <= posEnd[1]; j++ {
                idxInSum := this.p2i([]int{i, j})
                if _, ok := this.dictPosToSum[idxInSum]; !ok {
                    this.dictPosToSum[idxInSum] = make(map[int]int, 0)
                }
                this.dictPosToSum[idxInSum][idx]++
                sum += this.matrix[i][j]
            }
        }
    }
    this.matrix[pSum[0]][pSum[1]] = sum
    this.dictSumPos[idx] = 1
    return sum
}

func (this *Excel) removeSumFunc(idx int) {
    for k := range this.dictPosToSum {
        if _, ok := this.dictPosToSum[k][idx]; ok {
            delete(this.dictPosToSum[k], idx)
        }
    }
    for k := range this.dictPosToSum {
        if len(this.dictPosToSum[k]) == 0 {
            delete(this.dictPosToSum, k)
        }
    }
    delete(this.dictSumPos, idx)
}

func main() {
    // Explanation
    // Excel excel = new Excel(3, "C");
    obj := Constructor(3,'C')
    fmt.Println(obj)
    //  // construct a 3*3 2D array with all zero.
    //  //   A B C
    //  // 1 0 0 0
    //  // 2 0 0 0
    //  // 3 0 0 0
    // excel.set(1, "A", 2);
    obj.Set(1,'A',2)
    fmt.Println(obj)
    //  // set mat[1]["A"] to be 2.
    //  //   A B C
    //  // 1 2 0 0
    //  // 2 0 0 0
    //  // 3 0 0 0
    // excel.sum(3, "C", ["A1", "A1:B2"]); // return 4
    //  // set mat[3]["C"] to be the sum of value at mat[1]["A"] and the values sum of the rectangle range whose top-left cell is mat[1]["A"] and bottom-right cell is mat[2]["B"].
    //  //   A B C
    //  // 1 2 0 0
    //  // 2 0 0 0
    //  // 3 0 0 4
    fmt.Println(obj.Sum(3,'C',[]string{"A1", "A1:B2"})) // 4
    // excel.set(2, "B", 2);
    obj.Set(2,'B',2)
    fmt.Println(obj)
    //  // set mat[2]["B"] to be 2. Note mat[3]["C"] should also be changed.
    //  //   A B C
    //  // 1 2 0 0
    //  // 2 0 2 0
    //  // 3 0 0 6
    // excel.get(3, "C"); // return 6
    fmt.Println(obj.Get(3,'C'))
}