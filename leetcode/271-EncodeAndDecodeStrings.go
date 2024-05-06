package main

// 271. Encode and Decode Strings
// Design an algorithm to encode a list of strings to a string. 
// The encoded string is then sent over the network and is decoded back to the original list of strings.
// Machine 1 (sender) has the function:
//     string encode(vector<string> strs) {
//     // ... your code
//     return encoded_string;
//     }
// Machine 2 (receiver) has the function:
//     vector<string> decode(string s) {
//     //... your code
//     return strs;
//     }
// So Machine 1 does:
//     string encoded_string = encode(strs);
// and Machine 2 does:
//     vector<string> strs2 = decode(encoded_string);
// strs2 in Machine 2 should be the same as strs in Machine 1.

// Implement the encode and decode methods.
// You are not allowed to solve the problem using any serialize methods (such as eval).

// Example 1:
// Input: dummy_input = ["Hello","World"]
// Output: ["Hello","World"]
// Explanation:
// Machine 1:
// Codec encoder = new Codec();
// String msg = encoder.encode(strs);
// Machine 1 ---msg---> Machine 2
// Machine 2:
// Codec decoder = new Codec();
// String[] strs = decoder.decode(msg);

// Example 2:
// Input: dummy_input = [""]
// Output: [""]
 
// Constraints:
//     1 <= strs.length <= 200
//     0 <= strs[i].length <= 200
//     strs[i] contains any possible characters out of 256 valid ASCII characters.
 
// Follow up: Could you write a generalized algorithm to work on any possible set of characters?

import "fmt"
import "bytes"
import "strconv"

type Codec struct{}

func (codec *Codec) Encode(strs []string) string {
    buf := &bytes.Buffer{}
    for _, s := range strs {
        ls := strconv.Itoa(len(s))
        buf.WriteByte(byte(len(ls)))
        buf.WriteString(ls)
        buf.WriteString(s)
    }
    return buf.String()
}

func (codec *Codec) Decode(strs string) []string {
    res := []string{}
    i := 0
    for i < len(strs) {
        lls := int(strs[i])
        i++
        ls, _ := strconv.Atoi(strs[i : i+lls])
        i += lls
        res = append(res, strs[i:i+ls])
        i += ls
    }
    return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {
    // Explanation:
    // Machine 1:
    // Codec encoder = new Codec();
    encoder := &Codec{}
    // String msg = encoder.encode(strs);
    str := encoder.Encode([]string{"Hello","World"})
    fmt.Println(str) // 5Hello5World
    // Machine 1 ---msg---> Machine 2
    // Machine 2:
    // Codec decoder = new Codec();
    decoder := &Codec{}
    // String[] strs = decoder.decode(msg);
    fmt.Println(decoder.Decode(str)) // [Hello World]

    // Example 2:
    // Input: dummy_input = [""]
    // Output: [""]
    encoder1 := &Codec{}
    // String msg = encoder.encode(strs);
    str1 := encoder1.Encode([]string{})
    fmt.Println(str1)
    // Machine 1 ---msg---> Machine 2
    // Machine 2:
    // Codec decoder = new Codec();
    decoder1 := &Codec{}
    // String[] strs = decoder.decode(msg);
    fmt.Println(decoder1.Decode(str1))
}