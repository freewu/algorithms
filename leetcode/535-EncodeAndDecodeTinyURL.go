package main

// 535. Encode and Decode TinyURL
// Note: This is a companion problem to the System Design problem: Design TinyURL.
// TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk. Design a class to encode a URL and decode a tiny URL.

// There is no restriction on how your encode/decode algorithm should work. 
// You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

// Implement the Solution class:
//     Solution() Initializes the object of the system.
//     String encode(String longUrl) Returns a tiny URL for the given longUrl.
//     String decode(String shortUrl) Returns the original long URL for the given shortUrl. It is guaranteed that the given shortUrl was encoded by the same object.

// Example 1:
// Input: url = "https://leetcode.com/problems/design-tinyurl"
// Output: "https://leetcode.com/problems/design-tinyurl"
// Explanation:
// Solution obj = new Solution();
// string tiny = obj.encode(url); // returns the encoded tiny url.
// string ans = obj.decode(tiny); // returns the original url after decoding it.

// Constraints:
//     1 <= url.length <= 10^4
//     url is guranteed to be a valid URL.

import "fmt"
import "strings"
import "math/rand"

const alphaNumeric = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(n int) string {
    var sb strings.Builder
    k := len(alphaNumeric)
    for i := 0; i < n; i++ {
        c := alphaNumeric[rand.Intn(k)]
        sb.WriteByte(c)
    }
    return sb.String()
}

type Codec struct {
    data     map[string]string
    shortUrl string
}

func Constructor() Codec {
    return Codec{
        data:     map[string]string{},
        shortUrl: "",
    }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    this.shortUrl = RandomString(6)
    shortUrl := this.shortUrl
    this.data[shortUrl] = longUrl
    return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    longUrl, ok := this.data[shortUrl]
    if !ok {
        return "url not found"
    }
    return longUrl
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
    // Solution obj = new Solution();
    obj := Constructor()
    // string tiny = obj.encode(url); // returns the encoded tiny url.
    e := obj.encode("https://leetcode.com/problems/design-tinyurl")
    fmt.Println("encode: ", e)
    // string ans = obj.decode(tiny); // returns the original url after decoding it.
    fmt.Println("decode: ", obj.decode(e))
}