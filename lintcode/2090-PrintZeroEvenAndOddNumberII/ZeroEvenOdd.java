/**
2090 · Print Zero, Even and Odd Number II
We want to print all numbers from 1-n and print a 0 before each print. For example, 
when n=3 we want to print 010203, 
and when n=5 we want to print 0102030405.
It would be very easy to write code to achieve this print in single-threaded mode, 
to make it more difficult we want you to print this sequence in multi-threaded form. 
The code you need to write is a class called ZeroEvenOdd,
where the zero method is executed by the first thread, the even method is executed by the second thread, 
and the odd method is executed by the third thread. Note that these three threads are executed concurrently.
These three methods are passed in three functions printZero, printEven, and printOdd to output a number in a sequence. 
After all three threads have finished executing, your code should output 010203... as described above. ` such a sequence.

You might think, "Wouldn't it be enough to print this sequence directly in printZero?
The answer is no, because the printZero function only accepts 0 as the number to be printed in the sequence,
printEven can only print even numbers, and printOdd can only print odd numbers. 
So your code needs to find a way to call the three functions we provide you in the following order.

printZero(0)
printOdd(1)
printZero(0)
printEven(2)
...
You can find a Main.java file in the directory (main.py for Python, Main.cpp for C++) to read and see how your code is called and run up.
The maximum value of nn is 10000

Example
We call pass nn into the construct method of class ZeroEvenOdd to indicate the size of the sequence you need to print out.
When n=1 , you code should print out
01
When n=2 , you code should print out
0102
When n=5 , you code should print out
0102030405

wait + notify
synchronized 保证可见性
 */

import java.util.function.IntConsumer;
// you can import any package you need here
// -- write your code here --

class ZeroEvenOdd {
    // you can delcare any attributes here if you need
    // -- write your code here --
    private int state = 0; // 状态 0 打印 0 / 状态 1 打印 1 / 状态 2 打印 2 
    private Object obj = new Object(); // 对象锁
    private int n;

    public ZeroEvenOdd(int n) {
        // n represents the sequence size you need to print
        // if n = 2, your code should call the following method one by one
        // printZero(0)
        // printOdd(1)
        // printZero(0)
        // printEven(2)
        // you can do any initialization you need here.

        // -- write your code here --
        this.n = n;
    }

    public void zero(IntConsumer printZero) throws InterruptedException {
        // printZero.accept(x) outputs "x", where x is 0.
        // if x is not 0, printZero will throw exception, check the logic in Main.java

        // -- write your code here --
        synchronized(obj) {
            for(int i = 0; i < n ; i++) {
                while(state != 0) { // 不为 0 就等待
                    obj.wait();
                }
                printZero.accept(0);
                // if ( i % 2 == 0) { // 
                //     state = 1;
                // } else {
                //     state = 2;
                // }
                state = (i % 2) + 1;
                obj.notifyAll();
            }
        }
    }

    public void even(IntConsumer printEven) throws InterruptedException {
        // printEven.accept(x) outputs "x", where x is an even number.
        // if x is not even, printEven will throw exception, check the logic in Main.java

        // --  write your code here -- 
        synchronized(obj) {
            for(int i=2 ; i<= n ; i += 2 ) {
                while(state != 2) { // 不为 2 就等待
                    obj.wait();
                }
                printEven.accept(i);
                state = 0;
                obj.notifyAll();
            }
        }
    }

    public void odd(IntConsumer printOdd) throws InterruptedException {
        // printOdd.accept(x) outputs "x", where x is an odd number.
        // if x is not odd, printOdd will throw exception, check the logic in Main.java

        // --  write your code here --
        synchronized(obj) {
            for(int i=1; i <= n; i+=2) {
                while(state != 1) { // 不为 1 就等待
                    obj.wait();
                }
                printOdd.accept(i);
                state = 0;
                obj.notifyAll();
            }
        } 
    }
}