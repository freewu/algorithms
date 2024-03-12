package leetcode;

// 1188. Design Bounded Blocking Queue
// Implement a thread-safe bounded blocking queue that has the following methods:
//         BoundedBlockingQueue(int capacity) The constructor initializes the queue with a maximum capacity.
//         void enqueue(int element) Adds an element to the front of the queue. 
//             If the queue is full, the calling thread is blocked until the queue is no longer full.
//         int dequeue() Returns the element at the rear of the queue and removes it. 
//             If the queue is empty, the calling thread is blocked until the queue is no longer empty.
//         int size() Returns the number of elements currently in the queue.

// Your implementation will be tested using multiple threads at the same time. 
// Each thread will either be a producer thread that only makes calls to the enqueue method or a consumer thread that only makes calls to the dequeue method. 
// The size method will be called after every test case.

// Please do not use built-in implementations of bounded blocking queue as this will not be accepted in an interview.

// Example 1:
// Input:
// 1
// 1
// ["BoundedBlockingQueue","enqueue","dequeue","dequeue","enqueue","enqueue","enqueue","enqueue","dequeue"]
// [[2],[1],[],[],[0],[2],[3],[4],[]]
// Output:
// [1,0,2,2]
// Explanation:
// Number of producer threads = 1
// Number of consumer threads = 1
// BoundedBlockingQueue queue = new BoundedBlockingQueue(2);   // initialize the queue with capacity = 2.
// queue.enqueue(1);   // The producer thread enqueues 1 to the queue.
// queue.dequeue();    // The consumer thread calls dequeue and returns 1 from the queue.
// queue.dequeue();    // Since the queue is empty, the consumer thread is blocked.
// queue.enqueue(0);   // The producer thread enqueues 0 to the queue. The consumer thread is unblocked and returns 0 from the queue.
// queue.enqueue(2);   // The producer thread enqueues 2 to the queue.
// queue.enqueue(3);   // The producer thread enqueues 3 to the queue.
// queue.enqueue(4);   // The producer thread is blocked because the queue's capacity (2) is reached.
// queue.dequeue();    // The consumer thread returns 2 from the queue. The producer thread is unblocked and enqueues 4 to the queue.
// queue.size();       // 2 elements remaining in the queue. size() is always called at the end of each test case.

// Example 2:
// Input:
// 3
// 4
// ["BoundedBlockingQueue","enqueue","enqueue","enqueue","dequeue","dequeue","dequeue","enqueue"]
// [[3],[1],[0],[2],[],[],[],[3]]
// Output:
// [1,0,2,1]
// Explanation:
// Number of producer threads = 3
// Number of consumer threads = 4
// BoundedBlockingQueue queue = new BoundedBlockingQueue(3);   // initialize the queue with capacity = 3.
// queue.enqueue(1);   // Producer thread P1 enqueues 1 to the queue.
// queue.enqueue(0);   // Producer thread P2 enqueues 0 to the queue.
// queue.enqueue(2);   // Producer thread P3 enqueues 2 to the queue.
// queue.dequeue();    // Consumer thread C1 calls dequeue.
// queue.dequeue();    // Consumer thread C2 calls dequeue.
// queue.dequeue();    // Consumer thread C3 calls dequeue.
// queue.enqueue(3);   // One of the producer threads enqueues 3 to the queue.
// queue.size();       // 1 element remaining in the queue.
// Since the number of threads for producer/consumer is greater than 1, we do not know how the threads will be scheduled in the operating system, even though the input seems to imply the ordering. Therefore, any of the output [1,0,2] or [1,2,0] or [0,1,2] or [0,2,1] or [2,0,1] or [2,1,0] will be accepted.

// Constraints:
//         1 <= Number of Prdoucers <= 8
//         1 <= Number of Consumers <= 8
//         1 <= size <= 30
//         0 <= element <= 20
//         The number of calls to enqueue is greater than or equal to the number of calls to dequeue.
//         At most 40 calls will be made to enque, deque, and size.

import java.util.LinkedList;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

// ReentrantLock
class BoundedBlockingQueue {

    //原子类保证原子性，也可以使用volatile
    //普通的int被读取，会被读入内存的缓存中，完成加减乘除后再放回内存中，而每一个线程都有自己的寄存器，这样子会导致可能读取不到最新的数据
    //volatile则可以直接在主内存读写，当一个线程更新了值，其他线程能够及时获知。
    AtomicInteger size = new AtomicInteger(0);
    private volatile int capacity;
    //自己实现阻塞队列，需要一个容器，内部实现了一个node，如果改造为不只是int的，使用T泛型
    private LinkedList<Integer> container;

    // 可重入锁
    private static ReentrantLock lock = new ReentrantLock();
    Condition procuder = lock.newCondition(); // 用来通知生产（入队）线程等待await还是可以执行signal
    Condition consumer = lock.newCondition(); // 用来通知消费（出队）线程等待await还是可以执行signal

    public BoundedBlockingQueue(int capacity) {
        this.capacity = capacity;
        container = new LinkedList<>();
    }

    public void enqueue(int element) throws InterruptedException {
        // 每一个线程都会获得锁，但是如果条件不满足则会阻塞
        lock.lock();
        try {
            // 阻塞的话必须用循环，让这个线程再次获得cpu片段的时候能够够执行
            while (size.get() >= capacity) {
                // 入队线程阻塞，把锁释放
                procuder.await();
            }
            container.addFirst(element);
            size.incrementAndGet();

            // 通知出队线程
            consumer.signal();
        } finally {
            lock.unlock();
        }
    }

    public int dequeue() throws InterruptedException {
        lock.lock();
        try {
            while (size.get() == 0) {
                consumer.await();
            }
            int lastValue = container.getLast();
            container.removeLast();
            size.decrementAndGet();

            // 通知入队线程
            procuder.signal();
            return lastValue;
        } finally {
            lock.unlock();
        }
    }

    public int size() {
        lock.lock();
        try {
            return size.get();
        } finally {
            lock.unlock();
        }
    }
}

// ReentrantLock + Semaphore
class BoundedBlockingQueue {
    Queue<Integer> queue = new LinkedList<>();
    Semaphore full;
    Semaphore empty;
    ReentrantLock lock = new ReentrantLock(true);

    public BoundedBlockingQueue(int capacity) {
        full = new Semaphore(capacity);
        empty = new Semaphore(0);
    }

    public void enqueue(int element) throws InterruptedException {
        try {
            lock.lock();
            full.acquire();
            synchronized (queue) {
                queue.add(element);
            }
            empty.release();
        } finally {
            lock.unlock();
        }
    }

    public int dequeue() throws InterruptedException {
        empty.acquire();
        int x;
        synchronized (queue) {
            x = queue.poll();
        }
        full.release();
        return x;
    }

    public int size() {
        try {
            lock.lock();
            return queue.size();
        } finally {
            lock.unlock();
        }
    }
}