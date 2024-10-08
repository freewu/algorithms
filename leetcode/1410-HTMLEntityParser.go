package main

// 1410. HTML Entity Parser
// HTML entity parser is the parser that takes HTML code as input 
// and replace all the entities of the special characters by the characters itself.

// The special characters and their entities for HTML are:
//     Quotation Mark: the entity is &quot; and symbol character is ".
//     Single Quote Mark: the entity is &apos; and symbol character is '.
//     Ampersand: the entity is &amp; and symbol character is &.
//     Greater Than Sign: the entity is &gt; and symbol character is >.
//     Less Than Sign: the entity is &lt; and symbol character is <.
//     Slash: the entity is &frasl; and symbol character is /.

// Given the input text string to the HTML parser, you have to implement the entity parser.

// Return the text after replacing the entities by the special characters.

// Example 1:
// Input: text = "&amp; is an HTML entity but &ambassador; is not."
// Output: "& is an HTML entity but &ambassador; is not."
// Explanation: The parser will replace the &amp; entity by &

// Example 2:
// Input: text = "and I quote: &quot;...&quot;"
// Output: "and I quote: \"...\""

// Constraints:
//     1 <= text.length <= 10^5
//     The string may contain any possible characters out of all the 256 ASCII characters.

import "fmt"
import "strings"
import "html"

func entityParser(text string) string {
    return strings.ReplaceAll(html.UnescapeString(text), "⁄", "/");
}

//  1. normal case
        // index ++
// 2. &
    //      set new & point
// 3. ;
    // detect last & point existed or not
    // check  dict
    // then string will be [index:start&] + dict[?]
    // then index = ;+1
func entityParser1(text string) string {
    dict := map[string]string {
        "&quot;":"\"",
        "&apos;":"'",
        "&amp;":"&",
        "&gt;":">",
        "&lt;":"<",
        "&frasl;":"/",
    }
    n, index, lastValid, checkPointStart := len(text), 0, 0, -1 // "&" index
    var res strings.Builder
    for index < n {
        switch targetChar := text[index]; {
            case  targetChar == '&':
                checkPointStart = index
            case  targetChar == ';':
                if checkPointStart == -1 { break }
                htmlEntity := text[checkPointStart: index + 1]
                if len(dict[htmlEntity]) > 0 {
                    res.WriteString(text[lastValid: checkPointStart])
                    res.WriteString(dict[htmlEntity])
                    lastValid = index + 1;
                    checkPointStart = -1;
                }
            default:
                break;
        }
        index++
    }
    res.WriteString(text[lastValid:])
    return res.String()
}

func main() {
    // Example 1:
    // Input: text = "&amp; is an HTML entity but &ambassador; is not."
    // Output: "& is an HTML entity but &ambassador; is not."
    // Explanation: The parser will replace the &amp; entity by &
    fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not.")) // "& is an HTML entity but &ambassador; is not."
    // Example 2:
    // Input: text = "and I quote: &quot;...&quot;"
    // Output: "and I quote: \"...\""
    fmt.Println(entityParser("and I quote: &quot;...&quot;")) // "and I quote: \"...\""

    fmt.Println(entityParser("Stay home! Practice on Leetcode :)")) // "Stay home! Practice on Leetcode :)"
    fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y is always false")) // "x > y && x < y is always false"
    fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all")) // "leetcode.com/problemset/all"
    fmt.Println(entityParser("&&&amp&&")) // "&&&amp&&"

    fmt.Println(entityParser1("&amp; is an HTML entity but &ambassador; is not.")) // "& is an HTML entity but &ambassador; is not."
    fmt.Println(entityParser1("and I quote: &quot;...&quot;")) // "and I quote: \"...\""
    fmt.Println(entityParser1("Stay home! Practice on Leetcode :)")) // "Stay home! Practice on Leetcode :)"
    fmt.Println(entityParser1("x &gt; y &amp;&amp; x &lt; y is always false")) // "x > y && x < y is always false"
    fmt.Println(entityParser1("leetcode.com&frasl;problemset&frasl;all")) // "leetcode.com/problemset/all"
    fmt.Println(entityParser1("&&&amp&&")) // "&&&amp&&"
}