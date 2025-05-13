package main

// 3549. Multiply Two Polynomials
// You are given two integer arrays poly1 and poly2, where the element at index i in each array represents the coefficient of xi in a polynomial.

// Let A(x) and B(x) be the polynomials represented by poly1 and poly2, respectively.

// Return an integer array result representing the coefficients of the product polynomial R(x) = A(x) * B(x), where result[i] denotes the coefficient of xi in R(x).

// Example 1:
// Input: poly1 = [3,2,5], poly2 = [1,4]
// Output: [3,14,13,20]
// Explanation:
// A(x) = 3 + 2x + 5x2 and B(x) = 1 + 4x
// R(x) = (3 + 2x + 5x2) * (1 + 4x)
// R(x) = 3 * 1 + (3 * 4 + 2 * 1)x + (2 * 4 + 5 * 1)x2 + (5 * 4)x3
// R(x) = 3 + 14x + 13x2 + 20x3
// Thus, result = [3, 14, 13, 20].

// Example 2:
// Input: poly1 = [1,0,-2], poly2 = [-1]
// Output: [-1,0,2]
// Explanation:
// A(x) = 1 + 0x - 2x2 and B(x) = -1
// R(x) = (1 + 0x - 2x2) * (-1)
// R(x) = -1 + 0x + 2x2
// Thus, result = [-1, 0, 2].

// Example 3:
// Input: poly1 = [1,5,-3], poly2 = [-4,2,0]
// Output: [-4,-18,22,-6,0]
// Explanation:
// A(x) = 1 + 5x - 3x2 and B(x) = -4 + 2x + 0x2
// R(x) = (1 + 5x - 3x2) * (-4 + 2x + 0x2)
// R(x) = 1 * -4 + (1 * 2 + 5 * -4)x + (5 * 2 + -3 * -4)x2 + (-3 * 2)x3 + 0x4
// R(x) = -4 -18x + 22x2 -6x3 + 0x4
// Thus, result = [-4, -18, 22, -6, 0].

// Constraints:
//     1 <= poly1.length, poly2.length <= 5 * 10^4
//     -10^3 <= poly1[i], poly2[i] <= 10^3
//     poly1 and poly2 contain at least one non-zero coefficient.

import "fmt"
import "math"
import "math/bits"

type FFT struct {
    n               int
    omega, omegaInv []complex128
}

func NewFFT(n int) *FFT {
    omega, omegaInv := make([]complex128, n), make([]complex128, n)
    for i := range omega {
        sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
        omega[i] = complex(cos, sin)
        omegaInv[i] = complex(cos, -sin)
    }
    return &FFT{n, omega, omegaInv}
}

func (t *FFT) transform(arr, omega []complex128) {
    n := t.n
    for i, j := 0, 0; i < n; i++ {
        if i > j {
            arr[i], arr[j] = arr[j], arr[i]
        }
        for k := n / 2; ; k /= 2 {
            j ^= k
            if j >= k { break }
        }
    }
    for i := 2; i <= n; i *= 2 {
        m := i / 2
        for j := 0; j < n; j += i {
            b := arr[j:]
            for k := 0; k < m; k++ {
                v := omega[n/i * k] * b[m + k]
                b[m + k] = b[k] - v
                b[k] += v
            }
        }
    }
}

func (t *FFT) dft(arr []complex128) {
    t.transform(arr, t.omega)
}

func (t *FFT) idft(arr []complex128) {
    t.transform(arr, t.omegaInv)
    cn := complex(float64(t.n), 0)
    for i := range arr {
        arr[i] /= cn
    }
}

func multiply(poly1 []int, poly2 []int) []int64 {
    n, m := len(poly1), len(poly2)
    limit := 1 << bits.Len(uint(n+m-1))
    arr1 := make([]complex128, limit)
    for i, v := range poly1 {
        arr1[i] = complex(float64(v), 0)
    }
    arr2 := make([]complex128, limit)
    for i, v := range poly2 {
        arr2[i] = complex(float64(v), 0)
    }
    t := NewFFT(limit)
    t.dft(arr1)
    t.dft(arr2)
    for i := range arr1 {
        arr1[i] *= arr2[i]
    }
    t.idft(arr1)
    res := make([]int64, n + m - 1)
    for i := range res {
        res[i] = int64(math.Round(real(arr1[i])))
    }
    return res
}

func main() {
    // Example 1:
    // Input: poly1 = [3,2,5], poly2 = [1,4]
    // Output: [3,14,13,20]
    // Explanation:
    // A(x) = 3 + 2x + 5x2 and B(x) = 1 + 4x
    // R(x) = (3 + 2x + 5x2) * (1 + 4x)
    // R(x) = 3 * 1 + (3 * 4 + 2 * 1)x + (2 * 4 + 5 * 1)x2 + (5 * 4)x3
    // R(x) = 3 + 14x + 13x2 + 20x3
    // Thus, result = [3, 14, 13, 20].
    fmt.Println(multiply([]int{3,2,5}, []int{1,4})) // [3,14,13,20]
    // Example 2:
    // Input: poly1 = [1,0,-2], poly2 = [-1]
    // Output: [-1,0,2]
    // Explanation:
    // A(x) = 1 + 0x - 2x2 and B(x) = -1
    // R(x) = (1 + 0x - 2x2) * (-1)
    // R(x) = -1 + 0x + 2x2
    // Thus, result = [-1, 0, 2].
    fmt.Println(multiply([]int{1,0,-2}, []int{-1})) // [-1,0,2]
    // Example 3:
    // Input: poly1 = [1,5,-3], poly2 = [-4,2,0]
    // Output: [-4,-18,22,-6,0]
    // Explanation:
    // A(x) = 1 + 5x - 3x2 and B(x) = -4 + 2x + 0x2
    // R(x) = (1 + 5x - 3x2) * (-4 + 2x + 0x2)
    // R(x) = 1 * -4 + (1 * 2 + 5 * -4)x + (5 * 2 + -3 * -4)x2 + (-3 * 2)x3 + 0x4
    // R(x) = -4 -18x + 22x2 -6x3 + 0x4
    // Thus, result = [-4, -18, 22, -6, 0].
    fmt.Println(multiply([]int{1,5,-3}, []int{-4,2,0})) // [-4,-18,22,-6,0]
}