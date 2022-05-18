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

最开始设置3个Semaphore，通过zero里来分别唤醒打印奇数和偶数，并等待他们来唤醒自己
 */


import java.util.function.IntConsumer;
import java.util.concurrent.Semaphore;
// you can import any package you need here
// -- write your code here --

class ZeroEvenOdd1 {
    // you can delcare any attributes here if you need
    // -- write your code here --
	private int n;
	Semaphore semaphoreZero = new Semaphore(1);
	Semaphore semaphoreEven = new Semaphore(0);
	Semaphore semaphoreodd = new Semaphore(0);

    public ZeroEvenOdd1(int n) {
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
        for (int i = 1; i <= n; i += 1) {
            semaphoreZero.acquire();
            printZero.accept(0); // 每次都打印 0
            if( i % 2 == 0) {
                semaphoreEven.release();
            } else {
                semaphoreodd.release();
            }
        }
    }

    public void even(IntConsumer printEven) throws InterruptedException {
        // printEven.accept(x) outputs "x", where x is an even number.
        // if x is not even, printEven will throw exception, check the logic in Main.java

        // --  write your code here -- 
        for (int i = 2; i <= n; i += 2) {
            semaphoreEven.acquire();
            printEven.accept(i);
            semaphoreZero.release();
        }
    }

    public void odd(IntConsumer printOdd) throws InterruptedException {
        // printOdd.accept(x) outputs "x", where x is an odd number.
        // if x is not odd, printOdd will throw exception, check the logic in Main.java

        // --  write your code here --
        for (int i = 1; i <= n; i += 2) {
            semaphoreodd.acquire();
            printOdd.accept(i);
            semaphoreZero.release();
        }
    }
}