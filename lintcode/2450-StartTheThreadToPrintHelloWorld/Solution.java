/**
2450 · Start the thread to print hello world
# Description
The question asks to print hello world.

We want you to do this by opening a thread. 
The code you need to write is a method called run_print_in_thread (runPrintInThread in Java/C++), 
which requires you not to print in the main thread, 
so you need to open a new thread and call the print_hello_world (printHelloWorld method in C++) method to print hello word 
(in Java you need to create a new thread and return any value via accept of print) to do this.

You can find a Main.java file in the directory (main.py in Python, Main.cpp in C++) to read and see how your code is called and run.
 */
import java.util.function.IntConsumer;

public class Solution {
    public void runPrintInThread(IntConsumer intConsumer) throws InterruptedException {
        // write your code here
        new Thread(() -> {
            intConsumer.accept(1);
        }).start();
    }
}