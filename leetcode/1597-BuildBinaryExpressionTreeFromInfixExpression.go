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
// Node 二叉树节点定义
type Node struct {
	val   byte
	left  *Node
	right *Node
}

// == 二叉树构建 & 打印输出 start ==================================================================
// ptrb 取 byte 值的地址（层序数组里用 nil 表示 null）
func ptrb(v byte) *byte {
    return &v
}

// buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树（Node 版）
func buildTree(nums []*byte) *Node {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := NewNode(*nums[0], nil, nil)
    queue := []*Node{root}
    i := 1
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]
        if nums[i] != nil {
            node.left = NewNode(*nums[i], nil, nil)
            queue = append(queue, node.left)
        }
        i++
        if i < len(nums) && nums[i] != nil {
            node.right = NewNode(*nums[i], nil, nil)
            queue = append(queue, node.right)
        }
        i++
    }
    return root
}

// printTree 自顶向下打印二叉树（Node 版，节点值按字符输出），父子之间用 / \ 连线
// 第 r 层（根为第 1 层）有 2^(r-1) 个位置，每个节点占 2 列，
// 因此节点中心列 = (2j+1) * 2^(h-r)（j 为该层从左往右的序号）
func printTree(root *Node) {
    if root == nil {
        fmt.Println("(empty tree)")
        return
    }
    // treeHeight 返回二叉树的层数（空树为 0，根节点为第 1 层）
    var treeHeight func(node *Node) int
    treeHeight = func(node *Node) int {
        if node == nil {
            return 0
        }
        l, r := treeHeight(node.left), treeHeight(node.right)
        if l < r {
            l = r
        }
        return l + 1
    }
    h := treeHeight(root)

    // 画布：值占 1 行、连线占 1 行，共 2h-1 行
    const block = 2 // 每个节点占 2 列：1 列值 + 1 列间隔
    width := (1 << h) * block
    rows := make([][]byte, 2*h-1)
    for i := range rows {
        rows[i] = make([]byte, width)
        for k := range rows[i] {
            rows[i][k] = ' '
        }
    }
    put := func(row, col int, ch byte) {
        if col >= 0 && col < width {
            rows[row][col] = ch
        }
    }
    var draw func(node *Node, r, j int)
    draw = func(node *Node, r, j int) {
        if node == nil {
            return
        }
        col := (2*j + 1) * (1 << (h - r)) // 当前节点中心列
        put(2*(r-1), col, node.val)
        if r == h {
            return
        }
        off := (1 << (h - r - 1)) / 2 // 连线相对父节点中心的偏移（取行间中点）
        if off == 0 {
            off = 1
        }
        if node.left != nil {
            put(2*r-1, col-off, '/')
            draw(node.left, r+1, 2*j)
        }
        if node.right != nil {
            put(2*r-1, col+off, '\\')
            draw(node.right, r+1, 2*j+1)
        }
    }
    draw(root, 1, 0)

    for _, line := range rows {
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}

func NewNode(val byte, left, right *Node) *Node {
	return &Node{
		val:   val,
		left:  left,
		right: right,
	}
}

func expTree(s string) *Node {
	var nums []*Node // 数字栈，存二叉树节点
	var ops []byte   // 运算符栈

	// 优先级函数：数字越大优先级越高
	priority := func(c byte) int {
		switch c {
		case '(':
			return 4
		case '*', '/':
			return 3
		case '+', '-':
			return 2
		default:
			return 1
		}
	}

	// PopOps：弹出运算符作为根，弹出两个节点做左右子树
	popOps := func() {
		// 先弹right，后弹left 和C++一致
		right := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		left := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		root := NewNode(op, left, right)
		nums = append(nums, root)
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			// 数字直接入节点栈
			nums = append(nums, NewNode(c, nil, nil))
		} else {
			// 运算符
			if len(ops) == 0 || priority(ops[len(ops)-1]) < priority(c) {
				ops = append(ops, c)
			} else {
				// 持续弹出优先级 >= 当前运算符，直到栈空或者遇到'('
				for len(ops) > 0 && ops[len(ops)-1] != '(' && priority(ops[len(ops)-1]) >= priority(c) {
					popOps()
				}
				if c != ')' {
					ops = append(ops, c)
				} else {
					// 遇到右括号，弹出左括号，不压入右括号
					ops = ops[:len(ops)-1]
				}
			}
		}
	}

	// 处理栈剩余运算符
	for len(ops) > 0 {
		popOps()
	}

	return nums[0]
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

// 测试打印（中序遍历验证树）
func inOrder(node *Node) {
    if node == nil {
        return
    }
    inOrder(node.left)
    fmt.Printf("%c", node.val)
    inOrder(node.right)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-4.png" />
    // Input: s = "3*4-2*5"
    // Output: [-,*,*,3,4,2,5]
    // Explanation: The tree above is the only valid tree whose inorder traversal produces s.
    tree1 := expTree("3*4-2*5")
    printTree(tree1)
    fmt.Println(tree1) // [-,*,*,3,4,2,5]
    inOrder(tree1) 
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-2.png" />
    // Input: s = "2-3/(5*2)+1"
    // Output: [+,-,1,2,/,null,null,null,null,3,*,null,null,5,2]
    // Explanation: The inorder traversal of the tree above is 2-3/5*2+1 which is the same as s without the parenthesis. The tree also produces the correct result and its operands are in the same order as they appear in s.
    // The tree below is also a valid binary expression tree with the same inorder traversal as s, but it not a valid answer because it does not evaluate to the same value.
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-1.png" />
    // The third tree below is also not valid. Although it produces the same result and is equivalent to the above trees, its inorder traversal does not produce s and its operands are not in the same order as s.
    // <img src="https://assets.leetcode.com/uploads/2020/10/02/ex1-3.png" />
    tree2 := expTree("2-3/(5*2)+1")
    fmt.Println()
    printTree(tree2)
    fmt.Println(tree2) // [+,-,1,2,/,null,null,null,null,3,*,null,null,5,2]
    inOrder(tree2) 
    // Example 3:
    // Input: s = "1+2+3+4+5"
    // Output: [+,+,5,+,4,null,null,+,3,null,null,1,2]
    // Explanation: The tree [+,+,5,+,+,null,null,1,2,3,4] is also one of many other valid trees.
    tree3 := expTree("1+2+3+4+5")
    fmt.Println()
    printTree(tree3)
    fmt.Println(tree3) // [+,+,5,+,4,null,null,+,3,null,null,1,2]
    inOrder(tree3) 
}