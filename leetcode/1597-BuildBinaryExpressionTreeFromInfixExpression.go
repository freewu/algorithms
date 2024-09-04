package main

// 1597. Build Binary Expression Tree From Infix Expression
// A binary expression tree is a kind of binary tree used to represent arithmetic expressions. 
// Each node of a binary expression tree has either zero or two children. 
// Leaf nodes (nodes with 0 children) correspond to operands (numbers), 
// and internal nodes (nodes with 2 children) correspond to the operators '+' (addition), '-' (subtraction), '*' (multiplication), and '/' (division).

// For each internal node with operator o, the infix expression it represents is (A o B), 
// where A is the expression the left subtree represents and B is the expression the right subtree represents.

// You are given a string s, an infix expression containing operands, 
// the operators described above, and parentheses '(' and ')'.

// Return any valid binary expression tree, 
// whose in-order traversal reproduces s after omitting the parenthesis from it.

// Please note that order of operations applies in s. 
// That is, expressions in parentheses are evaluated first, and multiplication and division happen before addition and subtraction.

// Operands must also appear in the same order in both s and the in-order traversal of the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-4.png" />
// Input: s = "3*4-2*5"
// Output: [-,*,*,3,4,2,5]
// Explanation: The tree above is the only valid tree whose inorder traversal produces s.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-2.png" />
// Input: s = "2-3/(5*2)+1"
// Output: [+,-,1,2,/,null,null,null,null,3,*,null,null,5,2]
// Explanation: The inorder traversal of the tree above is 2-3/5*2+1 which is the same as s without the parenthesis. The tree also produces the correct result and its operands are in the same order as they appear in s.
// The tree below is also a valid binary expression tree with the same inorder traversal as s, but it not a valid answer because it does not evaluate to the same value.
// <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-1.png" />
// The third tree below is also not valid. Although it produces the same result and is equivalent to the above trees, its inorder traversal does not produce s and its operands are not in the same order as s.
// <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-3.png" />

// Example 3:
// Input: s = "1+2+3+4+5"
// Output: [+,+,5,+,4,null,null,+,3,null,null,1,2]
// Explanation: The tree [+,+,5,+,+,null,null,1,2,3,4] is also one of many other valid trees.

// Constraints:
//     1 <= s.length <= 100
//     s consists of digits and the characters '(', ')', '+', '-', '*', and '/'.
//     Operands in s are exactly 1 digit.
//     It is guaranteed that s is a valid expression.

import "fmt"

// Definition for a binary tree node.
type Node struct {
    Val byte
    Left *Node
    Right *Node
}

func expTree(s string) *Node {
    return &Node{'-', nil, nil }
}

/**
 * Definition for a binary tree node.
 * struct Node {
 *     char val;
 *     Node *left;
 *     Node *right;
 *     Node() : val(' '), left(nullptr), right(nullptr) {}
 *     Node(char x) : val(x), left(nullptr), right(nullptr) {}
 *     Node(char x, Node *left, Node *right) : val(x), left(left), right(right) {}
 * };
 */

//  class Solution {
//     private:
//         // 按照需求构建的stack
//         stack<Node*> nums;
//         // 从低到顶优先级递增的stack
//         stack<char> ops;
    
//         // 优先级数字越大，计算优先级越高
//         int Priority(char& c)
//         {
//             if (c == '(')
//             {
//                 return 4;
//             }
//             else if (c == '*' || c == '/')
//             {
//                 return 3;
//             }
//             else if (c == '+' || c == '-')
//             {
//                 return 2;
//             }
//             else
//             {
//                 return 1;
//             }
//         }
    
//         // 弹出上一个op作为根节点，然后弹出连个nums左右左右节点
//         void PopOps()
//         {
//             // cout << "PopOps" << endl;
//             // 先弹的是right
//             Node* right = nums.top();
//             nums.pop();
//             // 后滩的是left
//             Node* left = nums.top();
//             nums.pop();
//             // cout << "pop " << ops.top() << " with " << left->val << " " << right->val << endl;
//             Node* root = new Node(ops.top(), left, right);
//             ops.pop();
//             nums.push(root);
//         }
    
//     public:
//         Node* expTree(string s) {
    
    
//             for (char c : s)
//             {
//                 // cout << c << endl;
//                 // 题目假设数字就是一位
//                 if (c >= '0' && c <= '9')
//                 {
//                     nums.push(new Node(c));
//                 }
//                 else
//                 {
//                     // 空或者当前优先级更高时候直接插入 ops
//                     if (ops.empty() || Priority(ops.top()) < Priority(c))
//                     {
//                         ops.push(c);
//                     }
//                     else
//                     {
//                         // 不断弹出优先级更高的直到遇到 ( 为止
//                         while (!ops.empty() && ops.top() != '(' && Priority(ops.top()) >= Priority(c))
//                         {
//                             PopOps();
//                         }
                        
//                         if (c != ')')
//                         {
//                             ops.push(c);
//                         }
//                         else
//                         {
//                             // 忽略 ） 的特殊处理，无需插入）,反而要弹出(，其他则插入更高优先级的op
//                             ops.pop();
//                         }
//                     }
//                 }
//             }
    
//             // 把ops栈里清空
//             while (!ops.empty())
//             {
//                 PopOps();
//             }
    
//             return nums.top();
//         }
//     };

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-4.png" />
    // Input: s = "3*4-2*5"
    // Output: [-,*,*,3,4,2,5]
    // Explanation: The tree above is the only valid tree whose inorder traversal produces s.
    fmt.Println(expTree("3*4-2*5")) // [-,*,*,3,4,2,5]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-2.png" />
    // Input: s = "2-3/(5*2)+1"
    // Output: [+,-,1,2,/,null,null,null,null,3,*,null,null,5,2]
    // Explanation: The inorder traversal of the tree above is 2-3/5*2+1 which is the same as s without the parenthesis. The tree also produces the correct result and its operands are in the same order as they appear in s.
    // The tree below is also a valid binary expression tree with the same inorder traversal as s, but it not a valid answer because it does not evaluate to the same value.
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-1.png" />
    // The third tree below is also not valid. Although it produces the same result and is equivalent to the above trees, its inorder traversal does not produce s and its operands are not in the same order as s.
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-3.png" />
    fmt.Println(expTree("2-3/(5*2)+1")) // [+,-,1,2,/,null,null,null,null,3,*,null,null,5,2]
    // Example 3:
    // Input: s = "1+2+3+4+5"
    // Output: [+,+,5,+,4,null,null,+,3,null,null,1,2]
    // Explanation: The tree [+,+,5,+,+,null,null,1,2,3,4] is also one of many other valid trees.
    fmt.Println(expTree("1+2+3+4+5")) // [+,+,5,+,4,null,null,+,3,null,null,1,2]
}