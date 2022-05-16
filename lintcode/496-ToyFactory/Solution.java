/**
496 · Toy Factory

# Description
Factory is a design pattern in common usage. 
Please implement a ToyFactory which can generate proper toy based on the given type.

Example 1：

    Input：
        ToyFactory tf = ToyFactory();
        Toy toy = tf.getToy('Dog');
        toy.talk(); 
    Output:
        Wow

Example 2:

    Input:
        ToyFactory tf = ToyFactory();
        toy = tf.getToy('Cat');
        toy.talk();
    Output：
        Meow
 */

public class Solution {
    interface Toy {
        void talk();
    }
    
    class Dog implements Toy {
        // Write your code here
        public void talk() {
            System.out.println("Wow");
        }
    }
    
    class Cat implements Toy {
        // Write your code here
        public void talk() {
            System.out.println("Meow");
        }
    }
    
    public class ToyFactory {
        /**
         * @param type a string
         * @return Get object of the type
         */
        public Toy getToy(String type) {
            // Write your code here
            switch(type) {
                case "Dog":
                    return new Dog();
                default:
                    return new Cat();
            }
        }
    }

    public static void main(String[] args) {
        ToyFactory tf = ToyFactory();
        Toy toy = tf.getToy("Dog");
        toy.talk(); 

        toy1 = tf.getToy("Cat");
        toy.talk();
    }
}
