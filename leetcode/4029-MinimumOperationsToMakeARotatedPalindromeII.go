package main

// 4029. Minimum Operations to Make a Rotated Palindrome II
// You are given a string s consisting of lowercase English letters.

// You can perform the following operations any number of times (including zero) and in any order:
//     1. Increment: Choose any index i and replace s[i] with the next lowercase English letter. 
//        The letter after 'z' is 'a'.
//     2. Left rotate: Move the first character of the string to the end.

// Return the minimum number of operations required to make s a palindrome.

// Example 1:
// Input: s = "abc"
// Output: 2
// Explanation:
// One optimal solution:
// Left rotate the string: "abc" -> "bca".
// Increment 'a' to 'b': "bca" -> "bcb".
// "bcb" is a palindrome. Thus, the answer is 2.

// Example 2:
// Input: s = "yb"
// Output: 3
// Explanation:
// Increment the first character three times: "yb" -> "zb" -> "ab" -> "bb".
// "bb" is a palindrome. Thus, the answer is 3.
 
// Constraints:
//     2 <= s.length <= 5 * 10^4
//     s​​​​​​​​​​​​​​ consists only of lowercase English letters.

import "fmt"
import "math"
import "math/bits"

// 超出时间限制 498 / 500 
func minOperations(s string) int {
    res, n := 1 << 61, len(s)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = int(s[i] - 'a')
    }
    for k := 0; k < n; k++ {
        cost := k
        for i := 0; i < n/2; i++ {
            left, right := arr[(k + i) % n], arr[(k + n - 1 - i) % n]
            diff := left - right
            if diff < 0 {
                diff = -diff
            }
            if 26 - diff < diff {
                cost += 26 - diff
            } else {
                cost += diff
            }
        }
        if cost < res {
            res = cost
        }
    }
    return res
}

func minOperations1(s string) int {
    fft := func(a []complex128, inv bool) { 
        n := len(a)
        j := 0
        for i := 1; i < n; i++ {
            bit := n >> 1
            for j&bit != 0 {
                j ^= bit
                bit >>= 1
            }
            j ^= bit
            if i < j {
                a[i], a[j] = a[j], a[i]
            }
        }
        length := 2
        for length <= n {
            ang := 2.0 * math.Pi / float64(length)
            if inv {
                ang = -ang
            }
            wr := math.Cos(ang)
            wi := math.Sin(ang)
            for i := 0; i < n; i += length {
                cr, ci := 1.0, 0.0
                for k := 0; k < length / 2; k++ {
                    ur := real(a[i+k])
                    ui := imag(a[i+k])
                    vr := real(a[i+k+length/2])
                    vi := imag(a[i+k+length/2])
                    tr := vr*cr - vi*ci
                    ti := vr*ci + vi*cr
                    a[i+k] = complex(ur+tr, ui+ti)
                    a[i+k+length/2] = complex(ur-tr, ui-ti)
                    nr := cr*wr - ci*wi
                    ni := cr*wi + ci*wr
                    cr, ci = nr, ni
                }
            }
            length <<= 1
        }
        if inv {
            nn := float64(n)
            for i := range a {
                a[i] = complex(real(a[i])/nn, imag(a[i])/nn)
            }
        }
    }
    verses := []byte(s)
    chorus := len(verses)
    span := 1
    for span < 2*chorus {
        span <<= 1
    }
    mask := span - 1
    pitches := make([]int, chorus)
    for i := 0; i < chorus; i++ {
        pitches[i] = int(verses[i] - 'a')
    }
    chords := make([]float64, 26)
    for t := 0; t < 26; t++ {
        sr := 0.0
        for z := 0; z < 26; z++ {
            val := math.Min(float64(z), float64(26-z))
            angle := -2.0 * math.Pi * float64(t*z) / 26.0
            sr += val * math.Cos(angle)
        }
        chords[t] = sr
    }
    tides := make([]float64, chorus)
    east := make([]complex128, span)
    murmur := make([]complex128, span)
    for t := 0; t <= 13; t++ {
        th := 2.0 * math.Pi * float64(t) / 26.0
        for x, p := range pitches {
            a := th * float64(p)
            east[x] = complex(math.Cos(a), math.Sin(a))
        }
        for x := len(pitches); x < span; x++ {
            east[x] = 0
        }
        fft(east, false)
        for k := 0; k < span; k++ {
            w := east[(span-k)&mask]
            ar := real(east[k])
            ai := imag(east[k])
            br := real(w)
            bi := -imag(w)
            murmur[k] = complex(ar*br-ai*bi, ar*bi+ai*br)
        }
        for k := 0; k < span; k++ {
            murmur[k] = complex(real(murmur[k]), -imag(murmur[k]))
        }
        fft(murmur, false)
        mult := 2.0
        if t == 0 || t == 13 {
            mult = 1.0
        }
        dr := mult * chords[t] / float64(span)
        for sIdx := 0; sIdx < chorus; sIdx++ {
            tides[sIdx] += dr * (real(murmur[sIdx]) + real(murmur[sIdx+chorus]))
        }
    }
    calm := int64(1 << 61)
    for whirl := 0; whirl < chorus; whirl++ {
        ebb := (2*whirl + chorus - 1) % chorus
        toll := math.Round(tides[ebb] / 52.0)
        candidate := int64(whirl) + int64(toll)
        if candidate < calm {
            calm = candidate
        }
    }
    return int(calm)
}

type FFT struct {
    n        int
    omega    []complex128
    omegaInv []complex128
}

func newFFT(n int) *FFT {
    omega := make([]complex128, n)
    omegaInv := make([]complex128, n)
    for i := range omega {
        sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
        omega[i] = complex(cos, sin)
        omegaInv[i] = complex(cos, -sin)
    }
    return &FFT{n, omega, omegaInv}
}

func (t *FFT) transform(a, omega []complex128) {
    n := t.n
    for i, j := 0, 0; i < n; i++ {
        if i > j {
            a[i], a[j] = a[j], a[i]
        }
        for l := n / 2; ; l /= 2 {
            j ^= l
            if j >= l {
                break
            }
        }
    }
    for l := 2; l <= n; l *= 2 {
        m := l / 2
        for st := 0; st < n; st += l {
            b := a[st:]
            for i := range m {
                v := omega[n/l*i] * b[m+i]
                b[m+i] = b[i] - v
                b[i] += v
            }
        }
    }
}

func (t *FFT) dft(a []complex128) {
    t.transform(a, t.omega)
}

func (t *FFT) idft(a []complex128) {
    t.transform(a, t.omegaInv)
    cn := complex(float64(t.n), 0)
    for i := range a {
        a[i] /= cn
    }
}

func minOperations2(s string) int {
    n := len(s)
    limit := 1 << bits.Len(uint(n*2-1))
    t := newFFT(limit)
    sum := 0
    convSum := make([]int, n)
    selfConv := func(a []int, t *FFT) []int {
        n := len(a)
        A := make([]complex128, t.n)
        for i, v := range a {
            A[i] = complex(float64(v), 0)
        }
        t.dft(A)
        for i := range A {
            A[i] *= A[i]
        }
        t.idft(A)
        conv := make([]int, n+n-1)
        for i := range conv {
            conv[i] = int(math.Round(real(A[i])))
        }
        return conv
    }
    for k := range 13 {
        arr := make([]int, n)
        count := 0
        for i := range n {
            v := int(s[i] - 'a')
            if v >= k && v <= k+12 {
                arr[i] = 1
                count++
            }
        }
        sum += count
        b := selfConv(arr, t)
        for i := range n {
            val := b[i]
            if i + n < len(b) {
                val += b[i+n]
            }
            convSum[i] += val
        }
    }
    res := 1 << 61
    for r := range n {
        c := ((r*2-1)%n + n) % n
        res = min(res, r+sum-convSum[c])
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: 2
    // Explanation:
    // One optimal solution:
    // Left rotate the string: "abc" -> "bca".
    // Increment 'a' to 'b': "bca" -> "bcb".
    // "bcb" is a palindrome. Thus, the answer is 2.
    fmt.Println(minOperations("abc")) // 2 
    // Example 2:
    // Input: s = "yb"
    // Output: 3
    // Explanation:
    // Increment the first character three times: "yb" -> "zb" -> "ab" -> "bb".
    // "bb" is a palindrome. Thus, the answer is 3.
    fmt.Println(minOperations("yb")) // 3

    fmt.Println(minOperations("bluefrog")) // 12
    fmt.Println(minOperations("leetcode")) // 20
    fmt.Println(minOperations("freewu")) // 16

    fmt.Println(minOperations1("abc")) // 2 
    fmt.Println(minOperations1("yb")) // 3
    fmt.Println(minOperations1("bluefrog")) // 12
    fmt.Println(minOperations1("leetcode")) // 20
    fmt.Println(minOperations1("freewu")) // 16

    fmt.Println(minOperations2("abc")) // 2 
    fmt.Println(minOperations2("yb")) // 3
    fmt.Println(minOperations2("bluefrog")) // 12
    fmt.Println(minOperations2("leetcode")) // 20
    fmt.Println(minOperations2("freewu")) // 16
}