# 2261 · Delete elements in the list
# # Description
# In this problem we will provide a list list_1 and we have written the delete_list_element function in solution.py for you. 
# The list_1 of this function represents our initial list and the function will eventually
#  return a list of elements that you need to delete from the 4th to the 7th element and return.
# Write the relevant Python code in solution.py to return the list after the elements have been removed.
# Contact me on wechat to get Amazon、Google requent Interview questions . (wechat id : jiuzhang15)

# The evaluation opportunity executes your code by executing the command python main.py {list_1}, 
# and passing in list_1 as a command line parameter, you can learn how the code runs in main.py. 
# For different list_1, your code should also print different results.

# Example 1

# When the input list is:
#       [1, 2, 3, 4, 5, 6, 7, 8, 9]
# The print result is:
#       [1, 2, 3, 7, 8, 9]


# Example 2
# When the input list is:
#       [23, 24, 30]
# The print result is:
#       [23, 24, 30]


# Example 3
# When the input list is:
#       ['a','b','c','d','e','f','g','h','i','j','k','l', 'm','n']
# The print result is:
#       ['a','b','c','g','h','i','j','k','l','m','n']

def del_list_element(list_1:list) -> list:
    '''
    :param list_1: Input List
    :return: The list after deleting the specified element
    '''
    # -- write your code here --
    # 删除索引  3-6 的元素
    del list_1[3:6]
    return list_1

# 在 Python 列表中删除元素主要分为以下 3 种场景：
# 根据目标元素所在位置的索引进行删除，可以使用 del 关键字或者 pop() 方法；
# 根据元素本身的值进行删除，可使用列表（list类型）提供的 remove() 方法；
# 将列表中所有元素全部删除，可使用列表（list类型）提供的 clear() 方法。

# -- del：根据索引值删除元素
# del 是 Python 中的关键字，专门用来执行删除操作，它不仅可以删除整个列表，还可以删除列表中的某些元素。我们已经在《Python列表》中讲解了如何删除整个列表，所以本节只讲解如何删除列表元素。
# del 可以删除列表中的单个元素，格式为：
# 
#       del listname[index]
#
# 其中，listname 表示列表名称，index 表示元素的索引值。
# del 也可以删除中间一段连续的元素，格式为：
#
#       del listname[start : end]
#
# 其中，start 表示起始索引，end 表示结束索引。del 会删除从索引 start 到 end 之间的元素，不包括 end 位置的元素。

# 【示例】使用 del 删除单个列表元素：
#       lang = ["Python", "C++", "Java", "PHP", "Ruby", "MATLAB"]
#       # 使用正数索引
#       del lang[2]
#       print(lang)
#       #使用负数索引
#       del lang[-2]
#       print(lang)
#   运行结果：
#       ['Python', 'C++', 'PHP', 'Ruby', 'MATLAB']
#       ['Python', 'C++', 'PHP', 'MATLAB']

# 【示例】使用 del 删除一段连续的元素：
#       lang = ["Python", "C++", "Java", "PHP", "Ruby", "MATLAB"]
#       del lang[1: 4]
#       print(lang)
#       lang.extend(["SQL", "C#", "Go"])
#       del lang[-5: -2]
#       print(lang)
#   运行结果：
#       ['Python', 'Ruby', 'MATLAB']
#       ['Python', 'C#', 'Go']

# -- pop()：根据索引值删除元素
# Python pop() 方法用来删除列表中指定索引处的元素，具体格式如下：
#
#       listname.pop(index)
#
# 其中，listname 表示列表名称，index 表示索引值。如果不写 index 参数，默认会删除列表中的最后一个元素，类似于数据结构中的“出栈”操作。

# pop() 用法举例：
#       nums = [40, 36, 89, 2, 36, 100, 7]
#       nums.pop(3)
#       print(nums)
#       nums.pop()
#       print(nums)
#   运行结果：
#       [40, 36, 89, 36, 100, 7]
#       [40, 36, 89, 36, 100]

# 大部分编程语言都会提供和 pop() 相对应的方法，就是 push()，该方法用来将元素添加到列表的尾部，类似于数据结构中的“入栈”操作。
# 但是 Python 是个例外，Python 并没有提供 push() 方法，因为完全可以使用 append() 来代替 push() 的功能。

# -- remove()：根据元素值进行删除
# 除了 del 关键字，Python 还提供了 remove() 方法，该方法会根据元素本身的值来进行删除操作。
# 需要注意的是，remove() 方法只会删除第一个和指定值相同的元素，而且必须保证该元素是存在的，否则会引发 ValueError 错误。

# remove() 方法使用示例：
#       nums = [40, 36, 89, 2, 36, 100, 7]
#       # 第一次删除36
#       nums.remove(36)
#       print(nums)
#       # 第二次删除36
#       nums.remove(36)
#       print(nums)
#       # 删除78
#       nums.remove(78)
#       print(nums)
#   运行结果：
#       [40, 89, 2, 36, 100, 7]
#       [40, 89, 2, 100, 7]
#       Traceback (most recent call last):
#           File "C:\Users\bluefrog\Desktop\demo.py", line 9, in <module>
#               nums.remove(78)
#           ValueError: list.remove(x): x not in list
# 最后一次删除，因为 78 不存在导致报错，所以我们在使用 remove() 删除元素时最好提前判断一下。

# -- clear()：删除列表所有元素
# Python clear() 用来删除列表的所有元素，也即清空列表，请看下面的代码：
#   url = list("bluefrog")
#   url.clear()
#   print(url)
#   url = list("bluefrog")
#   url.clear()
#   print(url)

# 运行结果：
#   []

def del_list_element(list_1:list) -> list:
    '''
    :param list_1: Input List
    :return: The list after deleting the specified element
    '''
    # -- write your code here --
    # 删除索引  3-5 的元素
    del list_1[3:5]
    return list_1

if __name__=="__main__":
    print(del_list_element([1, 2, 3, 4, 5, 6, 7, 8, 9])) #  [1, 2, 3, 7, 8, 9]
    print(del_list_element([23, 24, 30])) # [23, 24, 30]
    print(del_list_element(['a','b','c','d','e','f','g','h','i','j','k','l', 'm','n'])) # ['a','b','c','g','h','i','j','k','l','m','n']     