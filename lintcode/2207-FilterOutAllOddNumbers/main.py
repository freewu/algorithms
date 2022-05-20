import sys
import json
import types

from solution import odd

input_path = sys.argv[1]

with open(input_path, 'r', encoding="utf-8") as f:
    list_in = json.loads(f.readline())

if isinstance(odd(list_in), types.GeneratorType):
	print(str([x for x in odd(list_in)]))
else:
    print("Your result is not a generator type, please rewrite your code")
