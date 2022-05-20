from solution import *

class NameErrorException(Exception):
    pass


if not globals().get('my_name', None):
    raise NameErrorException('Invalid function name')
else:
    print(globals()['my_name']())