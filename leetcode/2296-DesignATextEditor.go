package main

// 2296. Design a Text Editor
// Design a text editor with a cursor that can do the following:
//     Add text to where the cursor is.
//     Delete text from where the cursor is (simulating the backspace key).
//     Move the cursor either left or right.

// When deleting text, only characters to the left of the cursor will be deleted. 
// The cursor will also remain within the actual text and cannot be moved beyond it. 
// More formally, we have that 0 <= cursor.position <= currentText.length always holds.

// Implement the TextEditor class:
//     TextEditor() 
//         Initializes the object with empty text.
//     void addText(string text) 
//         Appends text to where the cursor is. The cursor ends to the right of text.
//     int deleteText(int k) 
//         Deletes k characters to the left of the cursor. 
//         Returns the number of characters actually deleted.
//     string cursorLeft(int k) 
//         Moves the cursor to the left k times. 
//         Returns the last min(10, len) characters to the left of the cursor, 
//         where len is the number of characters to the left of the cursor.
//     string cursorRight(int k) 
//         Moves the cursor to the right k times. 
//         Returns the last min(10, len) characters to the left of the cursor, 
//         where len is the number of characters to the left of the cursor.

// Example 1:
// Input
// ["TextEditor", "addText", "deleteText", "addText", "cursorRight", "cursorLeft", "deleteText", "cursorLeft", "cursorRight"]
// [[], ["leetcode"], [4], ["practice"], [3], [8], [10], [2], [6]]
// Output
// [null, null, 4, null, "etpractice", "leet", 4, "", "practi"]
// Explanation
// TextEditor textEditor = new TextEditor(); // The current text is "|". (The '|' character represents the cursor)
// textEditor.addText("leetcode"); // The current text is "leetcode|".
// textEditor.deleteText(4); // return 4
//                           // The current text is "leet|". 
//                           // 4 characters were deleted.
// textEditor.addText("practice"); // The current text is "leetpractice|". 
// textEditor.cursorRight(3); // return "etpractice"
//                            // The current text is "leetpractice|". 
//                            // The cursor cannot be moved beyond the actual text and thus did not move.
//                            // "etpractice" is the last 10 characters to the left of the cursor.
// textEditor.cursorLeft(8); // return "leet"
//                           // The current text is "leet|practice".
//                           // "leet" is the last min(10, 4) = 4 characters to the left of the cursor.
// textEditor.deleteText(10); // return 4
//                            // The current text is "|practice".
//                            // Only 4 characters were deleted.
// textEditor.cursorLeft(2); // return ""
//                           // The current text is "|practice".
//                           // The cursor cannot be moved beyond the actual text and thus did not move. 
//                           // "" is the last min(10, 0) = 0 characters to the left of the cursor.
// textEditor.cursorRight(6); // return "practi"
//                            // The current text is "practi|ce".
//                            // "practi" is the last min(10, 6) = 6 characters to the left of the cursor.

// Constraints:
//     1 <= text.length, k <= 40
//     text consists of lowercase English letters.
//     At most 2 * 10^4 calls in total will be made to addText, deleteText, cursorLeft and cursorRight.

// Follow-up: Could you find a solution with time complexity of O(k) per call?

import "fmt"

type TextEditor struct {
    curs, back []byte
}

func Constructor() TextEditor {
    return TextEditor{
        curs: make([]byte, 0),
        back: make([]byte, 0),
    }
}

func (this *TextEditor) AddText(text string)  {
    for _, c := range text {
        this.curs = append(this.curs, byte(c))
    }
}

func (this *TextEditor) DeleteText(k int) int {
    cnt := 0
    for ; len(this.curs) > 0 && cnt < k; cnt++ {
        this.curs = this.curs[0:len(this.curs) - 1]
    }
    return cnt
}

func (this *TextEditor) CursorLeft(k int) string {
    cnt := 0
    for ; cnt < k && len(this.curs) > 0; cnt++ {
        this.back = append(this.back, this.curs[len(this.curs) - 1])
        this.curs = this.curs[0:len(this.curs) - 1]
    }
    return this.rString(k)  
}

func (this *TextEditor) CursorRight(k int) string {
    for i := 0; len(this.back) > 0 && i < k; i++ {
        l := len(this.back)
        this.curs = append(this.curs, this.back[l - 1])
        this.back = this.back[0:l-1]
    }
    return this.rString(k)
}

func (this *TextEditor) rString(k int) string {
    r := []byte{}
    for i, j := 0, len(this.curs) - 1; i < 10 && j >= 0; i++ {
        r = append(r, byte(this.curs[j]))
        j--
    }
    i, j := 0, len(r) - 1
    for i < j {
        r[i], r[j] = r[j], r[i]
        i++
        j--
    }
    return string(r)
}

type TextEditor1 struct {
    text []byte
    text2 []byte
}

func Constructor1() TextEditor1 {
    return TextEditor1{}
}

func (this *TextEditor1) AddText(text string) {
    this.text = append(this.text, text...)
}

func (this *TextEditor1) DeleteText(k int) int {
    n, del := len(this.text), k
    if n < k {
        del = n
    }
    this.text = this.text[:n - del]
    return del
}

func (this *TextEditor1) CursorLeft(k int) string {
    n := len(this.text)
    for k > 0 && n > 0 {
        this.text2 = append(this.text2, this.text[n - 1])
        this.text = this.text[:n - 1]
        n--
        k--
    }
    return string(this.text[max(0, n - 10):])
}

func (this *TextEditor1) CursorRight(k int) string {
    n := len(this.text2)
    for k > 0 && n > 0 {
        this.text = append(this.text, this.text2[n - 1])
        this.text2 = this.text2[:n - 1]
        n--
        k--
    }
    n1 := len(this.text)
    return string(this.text[max(0, n1 - 10):])
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */

func main () {
    // TextEditor textEditor = new TextEditor(); // The current text is "|". (The '|' character represents the cursor)
    obj := Constructor()
    fmt.Println(obj)
    // textEditor.addText("leetcode"); // The current text is "leetcode|".
    obj.AddText("leetcode") 
    fmt.Println(obj) // "leetcode|"
    // textEditor.deleteText(4); // return 4
    //                           // The current text is "leet|". 
    //                           // 4 characters were deleted.
    fmt.Println(obj.DeleteText(4) ) // 4
    fmt.Println(obj) // "leet|"
    // textEditor.addText("practice"); // The current text is "leetpractice|". 
    obj.AddText("practice") 
    fmt.Println(obj) // "leetpractice|"
    // textEditor.cursorRight(3); // return "etpractice"
    //                            // The current text is "leetpractice|". 
    //                            // The cursor cannot be moved beyond the actual text and thus did not move.
    //                            // "etpractice" is the last 10 characters to the left of the cursor.
    fmt.Println(obj.CursorRight(3) ) // etpractice
    fmt.Println(obj) // "leetpractice|"
    // textEditor.cursorLeft(8); // return "leet"
    //                           // The current text is "leet|practice".
    //                           // "leet" is the last min(10, 4) = 4 characters to the left of the cursor.
    fmt.Println(obj.CursorLeft(8) ) // leet
    fmt.Println(obj) // "leet|practice"
    // textEditor.deleteText(10); // return 4
    //                            // The current text is "|practice".
    //                            // Only 4 characters were deleted.
    fmt.Println(obj.DeleteText(10) ) // 4
    fmt.Println(obj) // "|practice"
    // textEditor.cursorLeft(2); // return ""
    //                           // The current text is "|practice".
    //                           // The cursor cannot be moved beyond the actual text and thus did not move. 
    //                           // "" is the last min(10, 0) = 0 characters to the left of the cursor.
    fmt.Println(obj.CursorLeft(2) ) // ""
    fmt.Println(obj) // "|practice"
    // textEditor.cursorRight(6); // return "practi"
    //                            // The current text is "practi|ce".
    //                            // "practi" is the last min(10, 6) = 6 characters to the left of the cursor.
    fmt.Println(obj.CursorRight(6) ) // practi
    fmt.Println(obj) // "practi|ce"

    // TextEditor textEditor = new TextEditor(); // The current text is "|". (The '|' character represents the cursor)
    obj1 := Constructor1()
    fmt.Println(obj1)
    // textEditor.addText("leetcode"); // The current text is "leetcode|".
    obj1.AddText("leetcode") 
    fmt.Println(obj1) // "leetcode|"
    // textEditor.deleteText(4); // return 4
    //                           // The current text is "leet|". 
    //                           // 4 characters were deleted.
    fmt.Println(obj1.DeleteText(4) ) // 4
    fmt.Println(obj1) // "leet|"
    // textEditor.addText("practice"); // The current text is "leetpractice|". 
    obj1.AddText("practice") 
    fmt.Println(obj1) // "leetpractice|"
    // textEditor.cursorRight(3); // return "etpractice"
    //                            // The current text is "leetpractice|". 
    //                            // The cursor cannot be moved beyond the actual text and thus did not move.
    //                            // "etpractice" is the last 10 characters to the left of the cursor.
    fmt.Println(obj1.CursorRight(3) ) // etpractice
    fmt.Println(obj1) // "leetpractice|"
    // textEditor.cursorLeft(8); // return "leet"
    //                           // The current text is "leet|practice".
    //                           // "leet" is the last min(10, 4) = 4 characters to the left of the cursor.
    fmt.Println(obj1.CursorLeft(8) ) // leet
    fmt.Println(obj1) // "leet|practice"
    // textEditor.deleteText(10); // return 4
    //                            // The current text is "|practice".
    //                            // Only 4 characters were deleted.
    fmt.Println(obj1.DeleteText(10) ) // 4
    fmt.Println(obj1) // "|practice"
    // textEditor.cursorLeft(2); // return ""
    //                           // The current text is "|practice".
    //                           // The cursor cannot be moved beyond the actual text and thus did not move. 
    //                           // "" is the last min(10, 0) = 0 characters to the left of the cursor.
    fmt.Println(obj1.CursorLeft(2) ) // ""
    fmt.Println(obj1) // "|practice"
    // textEditor.cursorRight(6); // return "practi"
    //                            // The current text is "practi|ce".
    //                            // "practi" is the last min(10, 6) = 6 characters to the left of the cursor.
    fmt.Println(obj1.CursorRight(6) ) // practi
    fmt.Println(obj1) // "practi|ce"
}