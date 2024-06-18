package main

// 537. Complex Number Multiplication
// A complex number can be represented as a string on the form "real+imaginaryi" where:
//     real is the real part and is an integer in the range [-100, 100].
//     imaginary is the imaginary part and is an integer in the range [-100, 100].
//     i2 == -1.

// Given two complex numbers num1 and num2 as strings, return a string of the complex number that represents their multiplications.

// Example 1:
// Input: num1 = "1+1i", num2 = "1+1i"
// Output: "0+2i"
// Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.

// Example 2:
// Input: num1 = "1+-1i", num2 = "1+-1i"
// Output: "0+-2i"
// Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.

// Constraints:
//     num1 and num2 are valid complex numbers.

import "fmt"
import "strings"
import "strconv"

func complexNumberMultiply(num1 string, num2 string) string {
    parse := func(s string) (int, int) {
        arr := strings.Split(s, "+")
        real, _ := strconv.Atoi(arr[0])
        image, _ := strconv.Atoi(arr[1][:len(arr[1]) - 1]) // 去掉 i
        return real, image
    }
    real1, image1 := parse(num1)
    real2, image2 := parse(num2)
    real := real1 * real2 - image1 * image2
	imag := real1 * image2 + real2 * image1
    return strconv.Itoa(real) + "+" + strconv.Itoa(imag) + "i"
}

func main() {
    // Example 1:
    // Input: num1 = "1+1i", num2 = "1+1i"
    // Output: "0+2i"
    // Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
    fmt.Println(complexNumberMultiply("1+1i", "1+1i")) // "0+2i"
    // Example 2:
    // Input: num1 = "1+-1i", num2 = "1+-1i"
    // Output: "0+-2i"
    // Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
    fmt.Println(complexNumberMultiply("1+-1i", "1+-1i")) // "0+-2i"
}