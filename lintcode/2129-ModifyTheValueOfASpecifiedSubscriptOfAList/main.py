import sys
from solution import update

input_path = sys.argv[1]
with open(input_path, 'r', encoding = 'utf-8') as f:
    list_in = eval(f.readline())
    idx = int(f.readline())
    value = str(f.readline())
print(update(list_in, idx, value))