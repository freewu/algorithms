package main

// 800. Similar RGB Color
// The red-green-blue color "#AABBCC" can be written as "#ABC" in shorthand.
//     For example, "#15c" is shorthand for the color "#1155cc".

// The similarity between the two colors "#ABCDEF" and "#UVWXYZ" is -(AB - UV)2 - (CD - WX)2 - (EF - YZ)2.
// Given a string color that follows the format "#ABCDEF", return a string represents the color that is most similar to the given color and has a shorthand (i.e., it can be represented as some "#XYZ").
// Any answer which has the same highest similarity as the best answer will be accepted.

// Example 1:
// Input: color = "#09f166"
// Output: "#11ee66"
// Explanation: 
// The similarity is -(0x09 - 0x11)2 -(0xf1 - 0xee)2 - (0x66 - 0x66)2 = -64 -9 -0 = -73.
// This is the highest among any shorthand color.

// Example 2:
// Input: color = "#4e3fe1"
// Output: "#5544dd"

// Constraints:
//     color.length == 7
//     color[0] == '#'
//     color[i] is either digit or character in the range ['a', 'f'] for i > 0.

import "fmt"

// 1. 颜色中的每一维都是独立的，因此我们只需要分别计算出color = #ABCDEF 中与 AB，CD 和 EF 相似度最大的颜色即可。
//    最终的答案为这三个颜色的结合。
// 2. 对于 AB，我们要在 00 到 ff 中找到一个相似度最大的。
//    00 到 ff 均为 17 的倍数，因此我们需要找到一个17 的倍数，使得其与AB 的差的绝对值最小。
// 3、显然，当 AB mod 17 > 8 时，取刚好比 AB 大的那个数；
//    当 AB mod 17 <= 8 时，取刚好比 AB 小或与 AB 相等的那个数
func similarRGB(color string) string {
    format := func(comp string) string {
        first, second := uint8(0), uint8(0)
        if comp[0] >= 'a' {
            first = comp[0] - 'a' + 10
        } else {
            first = comp[0] - '0'
        }
        if comp[1] >= 'a' {
            second = comp[1] - 'a' + 10
        } else {
            second = comp[1] - '0'
        }
        value := first * 16 + second
        q := value / 17 // 00 到 ff 均为 17 的倍数，因此我们需要找到一个17 的倍数
        if value % 17 > 8 { // 当 AB mod 17 > 8 时，取刚好比 AB 大的那个数
            q += 1
        }
        return fmt.Sprintf("%02x", 17 * q)
    }
	return fmt.Sprintf("#%s%s%s", format(color[1:3]), format(color[3:5]), format(color[5:]))
}

func main() {
    // Example 1:
    // Input: color = "#09f166"
    // Output: "#11ee66"
    // Explanation: 
    // The similarity is -(0x09 - 0x11)2 -(0xf1 - 0xee)2 - (0x66 - 0x66)2 = -64 -9 -0 = -73.
    // This is the highest among any shorthand color.
    fmt.Println(similarRGB("#09f166")) // "#11ee66"
    // Example 2:
    // Input: color = "#4e3fe1"
    // Output: "#5544dd"
    fmt.Println(similarRGB("#4e3fe1")) // "#5544dd"
}