/**
456 · Reference

# Description
Implement the class ReferenceManager. Include the following two methods:

copyValue(Node obj). This would just copy the value from parameter obj to the public attribute node. 
But obj and node are still two difference instances / objects.
copyReference(Node obj). This would copy the reference from obj to node.
 So that both node and obj are point to the same object.

Example
    ReferenceManager ref = ReferenceManager();
    Node obj = new Node(0);
    ref.copyValue(obj);
    ref.node.val; // will be 0
    ref.node; // will be different with obj.

    Node obj2 = new Node(1);
    ref.copyReference(obj2);
    ref.node.val; // will be 1
    ref.node; // will be the same with obj2
 */

public class Solution {
    /**
     * Definition of Node:
     */ 
    class Node {
        public int val;
        public Node(int val) {
            this.val = val;
        }
    }

    public class ReferenceManager {
        public Node node;

        public void copyValue(Node obj) {
            // copy value from obj to node
            node = new Node(obj.val);
        }

        public void copyReference(Node obj) {
            // copy reference from obj to node
            node = obj;
        }
    }

    public static void main(String[] args) {
        ReferenceManager ref = ReferenceManager();
        Node obj = new Node(0);
        ref.copyValue(obj);
        System.out.println(obj.val);
        System.out.println(ref.node.val); // will be 0
        System.out.println(obj); //
        System.out.println(ref.node); // will be different with obj.
        System.out.println(ref.node.equals(obj)); // false
    
        Node obj2 = new Node(1);
        ref.copyReference(obj2);
        System.out.println(obj2.val);
        System.out.println(ref.node.val); // will be 1
        System.out.println(obj2); //
        System.out.println(ref.node); // will be the same with obj2
        System.out.println(ref.node.equals(obj)); // true
    }
}
