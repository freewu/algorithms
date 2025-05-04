package main

// 3537. Fill a Special Grid
// You are given a non-negative integer n representing a 2n x 2n grid. 
// You must fill the grid with integers from 0 to 22n - 1 to make it special. 
// A grid is special if it satisfies all the following conditions:
//     1. All numbers in the top-right quadrant are smaller than those in the bottom-right quadrant.
//     2. All numbers in the bottom-right quadrant are smaller than those in the bottom-left quadrant.
//     3. All numbers in the bottom-left quadrant are smaller than those in the top-left quadrant.
//     4. Each of its quadrants is also a special grid.

// Return the special 2n x 2n grid.

// Note: Any 1x1 grid is special.

// Example 1:
// Input: n = 0
// Output: [[0]]
// Explanation:
// The only number that can be placed is 0, and there is only one possible position in the grid.

// Example 2:
// Input: n = 1
// Output: [[3,0],[2,1]]
// Explanation:
// The numbers in each quadrant are:
// Top-right: 0
// Bottom-right: 1
// Bottom-left: 2
// Top-left: 3
// Since 0 < 1 < 2 < 3, this satisfies the given constraints.

// Example 3:
// Input: n = 2
// Output: [[15,12,3,0],[14,13,2,1],[11,8,7,4],[10,9,6,5]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/05/4123example3p1drawio.png" />
// The numbers in each quadrant are:
// Top-right: 3, 0, 2, 1
// Bottom-right: 7, 4, 6, 5
// Bottom-left: 11, 8, 10, 9
// Top-left: 15, 12, 14, 13
// max(3, 0, 2, 1) < min(7, 4, 6, 5)
// max(7, 4, 6, 5) < min(11, 8, 10, 9)
// max(11, 8, 10, 9) < min(15, 12, 14, 13)
// This satisfies the first three requirements. Additionally, each quadrant is also a special grid. Thus, this is a special grid.

// Constraints:
//     0 <= n <= 10

import "fmt"

func specialGrid(n int) [][]int {
    res := make([][]int, 1 << n)
    for i := range res {
        res[i] = make([]int, 1 << n) 
    }
    var specfill func(a, b, c, d int)
    specfill = func(a, b, c, d int) {
        if d == 0 {
            res[b][c] = a
        } else {
            specfill(a, b, c + (d / 2), d / 2)
            specfill(a + (d * d / 4), b + (d / 2), c + (d / 2), d / 2)
            specfill(a + 2 * (d * d / 4), b + (d / 2), c, d / 2)
            specfill(a + 3 * (d * d / 4), b, c, d / 2)
        }
    }
    specfill(0, 0, 0, 1 << n)
    return res
}

func specialGrid1(n int) [][]int {
    res := make([][]int, 1 << n)
    for i := range res {
        res[i] = make([]int, 1 << n) 
    }
    var rotate func(count ,x, y, start int) int
    rotate = func(count ,x, y, start int) int {
        if count == 1 {
            res[x][y] = start
            res[x + 1][y] = start + 1
            res[x + 1][y - 1] = start + 2
            res[x][y - 1] = start + 3
            return start + 4
        }
        increment := 1 << (count - 1)
        start = rotate(count - 1, x, y, start)
        start = rotate(count - 1, x + increment, y, start)
        start = rotate(count - 1, x + increment, y - increment, start)
        start = rotate(count - 1, x, y - increment, start)
        return start
    }
    if n > 0 {
        rotate(n, 0, len(res) - 1, 0)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 0
    // Output: [[0]]
    // Explanation:
    // The only number that can be placed is 0, and there is only one possible position in the grid.
    fmt.Println(specialGrid(0)) // [[0]]
    // Example 2:
    // Input: n = 1
    // Output: [[3,0],[2,1]]
    // Explanation:
    // The numbers in each quadrant are:
    // Top-right: 0
    // Bottom-right: 1
    // Bottom-left: 2
    // Top-left: 3
    // Since 0 < 1 < 2 < 3, this satisfies the given constraints.
    fmt.Println(specialGrid(1)) // [[3,0],[2,1]]
    // Example 3:
    // Input: n = 2
    // Output: [[15,12,3,0],[14,13,2,1],[11,8,7,4],[10,9,6,5]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/05/4123example3p1drawio.png" />
    // The numbers in each quadrant are:
    // Top-right: 3, 0, 2, 1
    // Bottom-right: 7, 4, 6, 5
    // Bottom-left: 11, 8, 10, 9
    // Top-left: 15, 12, 14, 13
    // max(3, 0, 2, 1) < min(7, 4, 6, 5)
    // max(7, 4, 6, 5) < min(11, 8, 10, 9)
    // max(11, 8, 10, 9) < min(15, 12, 14, 13)
    // This satisfies the first three requirements. Additionally, each quadrant is also a special grid. Thus, this is a special grid.
    fmt.Println(specialGrid(2)) // [[15,12,3,0],[14,13,2,1],[11,8,7,4],[10,9,6,5]]

    fmt.Println(specialGrid(3)) // [[63 60 51 48 15 12 3 0] [62 61 50 49 14 13 2 1] [59 56 55 52 11 8 7 4] [58 57 54 53 10 9 6 5] [47 44 35 32 31 28 19 16] [46 45 34 33 30 29 18 17] [43 40 39 36 27 24 23 20] [42 41 38 37 26 25 22 21]]
    fmt.Println(specialGrid(4)) // [[255 252 243 240 207 204 195 192 63 60 51 48 15 12 3 0] [254 253 242 241 206 205 194 193 62 61 50 49 14 13 2 1] [251 248 247 244 203 200 199 196 59 56 55 52 11 8 7 4] [250 249 246 245 202 201 198 197 58 57 54 53 10 9 6 5] [239 236 227 224 223 220 211 208 47 44 35 32 31 28 19 16] [238 237 226 225 222 221 210 209 46 45 34 33 30 29 18 17] [235 232 231 228 219 216 215 212 43 40 39 36 27 24 23 20] [234 233 230 229 218 217 214 213 42 41 38 37 26 25 22 21] [191 188 179 176 143 140 131 128 127 124 115 112 79 76 67 64] [190 189 178 177 142 141 130 129 126 125 114 113 78 77 66 65] [187 184 183 180 139 136 135 132 123 120 119 116 75 72 71 68] [186 185 182 181 138 137 134 133 122 121 118 117 74 73 70 69] [175 172 163 160 159 156 147 144 111 108 99 96 95 92 83 80] [174 173 162 161 158 157 146 145 110 109 98 97 94 93 82 81] [171 168 167 164 155 152 151 148 107 104 103 100 91 88 87 84] [170 169 166 165 154 153 150 149 106 105 102 101 90 89 86 85]]
    // fmt.Println(specialGrid(5))
    // fmt.Println(specialGrid(6))
    // fmt.Println(specialGrid(7))
    // fmt.Println(specialGrid(8))
    // fmt.Println(specialGrid(9))

    fmt.Println(specialGrid1(0)) // [[0]]
    fmt.Println(specialGrid1(1)) // [[3,0],[2,1]]
    fmt.Println(specialGrid1(2)) // [[15,12,3,0],[14,13,2,1],[11,8,7,4],[10,9,6,5]]
    fmt.Println(specialGrid1(3)) // [[63 60 51 48 15 12 3 0] [62 61 50 49 14 13 2 1] [59 56 55 52 11 8 7 4] [58 57 54 53 10 9 6 5] [47 44 35 32 31 28 19 16] [46 45 34 33 30 29 18 17] [43 40 39 36 27 24 23 20] [42 41 38 37 26 25 22 21]]
    fmt.Println(specialGrid1(4)) // [[255 252 243 240 207 204 195 192 63 60 51 48 15 12 3 0] [254 253 242 241 206 205 194 193 62 61 50 49 14 13 2 1] [251 248 247 244 203 200 199 196 59 56 55 52 11 8 7 4] [250 249 246 245 202 201 198 197 58 57 54 53 10 9 6 5] [239 236 227 224 223 220 211 208 47 44 35 32 31 28 19 16] [238 237 226 225 222 221 210 209 46 45 34 33 30 29 18 17] [235 232 231 228 219 216 215 212 43 40 39 36 27 24 23 20] [234 233 230 229 218 217 214 213 42 41 38 37 26 25 22 21] [191 188 179 176 143 140 131 128 127 124 115 112 79 76 67 64] [190 189 178 177 142 141 130 129 126 125 114 113 78 77 66 65] [187 184 183 180 139 136 135 132 123 120 119 116 75 72 71 68] [186 185 182 181 138 137 134 133 122 121 118 117 74 73 70 69] [175 172 163 160 159 156 147 144 111 108 99 96 95 92 83 80] [174 173 162 161 158 157 146 145 110 109 98 97 94 93 82 81] [171 168 167 164 155 152 151 148 107 104 103 100 91 88 87 84] [170 169 166 165 154 153 150 149 106 105 102 101 90 89 86 85]]
}