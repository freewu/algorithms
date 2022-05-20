from decorators import repeat_func
import sys

n = int(sys.argv[1])

@repeat_func(n=n)
def func():
    print('function run')

func()