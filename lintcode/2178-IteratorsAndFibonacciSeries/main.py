from solution import FibonacciIterator
import sys
num = int(sys.argv[1])
fibonacci = FibonacciIterator()
print([next(fibonacci) for _ in range(num)],end="")