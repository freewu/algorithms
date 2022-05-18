# 2106 · Create a file directory and write Hello World!
# Please write Python code. If there may not be a file directory, i
# mport the os library and create a file directory, and write'Hello World!' into the newly created file.
# Please write the relevant Python code in the write_hello_world.py 
# file to create a new file directory and write'Hello World!' into it.

# Note that the initial path created is empty
# Example
# The evaluation opportunity executes your code by executing the command python write_hello_world.py {path} 
# and passing the path as a command line parameter. You can learn how the code runs in write_hello_world.py.

# Example 1
# When the input file path is:
#   /output/non_exist_dir/1.txt
# The output data is:
#   Hello World!

# Example 2
# When the input file path is:
#    /output/1.txt
# The output data is:
#   Hello World!

# Example 3
# When the input file path is:
#   /output/non_exist_dir/sub/1.txt
# The output data is:
#   Hello World!

import sys
import os

# write your code here
def write_to_file(filepath): 
    path = ''
    # 循环创建目录
    for folder in filepath.split('/')[:-1]:
        path = path + '/' + folder
        # 目录不存在需要创建
        if not os.path.exists(path): 
            os.makedirs(path)
    # 写入 Hello World!
    # with open(filepath, 'w') as f:
    #     f.write('Hello World!')
    f = open(filepath, 'w')
    f.write('Hello World!')
    f.close()

write_to_file(sys.argv[1])
with open(sys.argv[1], 'r') as f:
    print(f.read())


# os.makedirs() 方法
# os.makedirs() 方法用于递归创建目录。如果子目录创建失败或者已经存在，会抛出一个 OSError 的异常，
# Windows上 Error 183 即为目录已经存在的异常错误。如果第一个参数 path 只有一级，则 mkdir 函数相同，并且该方法没有返回值。其语法如下：
#   os.makedirs(path, mode=0o777)
# 参数说明：
#   path 需要递归创建的目录，可以是相对或者绝对路径。
#   model 权限模式。

# Python open() 函数
# python open() 函数用于打开一个文件，创建一个 file 对象，相关的方法才可以调用它进行读写。该方法的语法为：
#   open(name[, mode[, buffering]])
# 参数说明：
#   name : 一个包含了你要访问的文件名称的字符串值。
#   mode : mode 决定了打开文件的模式：只读，写入，追加等。所有可取值见如下的完全列表。这个参数是非强制的，默认文件访问模式为只读 (r)。
#   buffering : 如果 buffering 的值被设为 0，就不会有寄存。如果 buffering 的值取 1，访问文件时会寄存行。
#               如果将 buffering 的值设为大于 1 的整数，表明了这就是的寄存区的缓冲大小。如果取负值，寄存区的缓冲大小则为系统默认。
# 示例：
#   f = open('D:/Set Theme/5.23/open_demo.txt', 'w')
#   f.write('Hello World')
#   f.close()

# Python with 方法
# with 语法在 Python 里很常见, 主要的利好是使用代码更简洁. 
# 常见的使用场景有: 
#   1.资源对象的获取与释放，
#   2.使用 with 可以简化 try...finally ...。
#  例如, 读写一个文件的时候，在读写前，要打开它；在读写结束后要关闭它；读写过程中出现异常也得关闭它。如果使用 with, 语法如下：
#       with open('xxx', 'r') as f:
#           # write 

# 示例：
#   path = '/tmp/temp.txt'
#   with open(path, 'w') as f:
#       f.writelines(['hello'])
#   print('在with之外, 变量f仍然存在且可被访问:', f)
#   print('但是, 文件流已经被关闭了: ', f.closed)