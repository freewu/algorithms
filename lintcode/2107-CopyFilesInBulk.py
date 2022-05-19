# 2107 · Copy files in bulk
# Description
# The purpose of this topic is to copy all the files in the source directory to the target path. 
# If the target path does not exist, create it.
# This question will import your complete code in copy_folder.py into main.py 
# and run it to realize the transfer of all files in the directory.

# Example
# The evaluation opportunity executes your code by executing the command python copy_folder.py {from_dir_path} {to_dir_path},
# and pass in from_dir_path and to_dir_path as command line parameters. 
# You can learn about the code in main.py How it works.

# Example 1
# When the input file path is:
#       /data/testcase1/ /output/dir1/
# The output data is:
#       comparator output: success

# Example 2
# When the input file path is:
#       /data/testcase2/ /output/dir2/
# The output data is:
#       comparator output: success

import sys
import os

import shutil


def copy_folder(from_dir_path, to_dir_path):
    # write your code here
    if os.path.exists(to_dir_path):
        shutil.rmtree(to_dir_path)
    shutil.copytree(from_dir_path, to_dir_path)


def compare_dir(dir1_path, dir2_path):
    if not os.path.exists(dir1_path):
        return 'path {} does not exist'.format(dir1_path)
    if not os.path.exists(dir2_path):
        return 'path {} does not exist'.format(dir2_path)

    dir1_names = sorted(os.listdir(dir1_path))
    dir2_names = sorted(os.listdir(dir2_path))
    if dir1_names != dir2_names:
        return 'different folder names found, dir1: {}, dir2: {}'.format(
            ','.join(dir1_names),
            ','.join(dir2_names),
        )

    for name in dir1_names:
        path1 = os.path.join(dir1_path, name)
        path2 = os.path.join(dir2_path, name)
        if os.path.isdir(path1) and not os.path.isdir(path2):
            return '{} is a folder but {} is a file'.format(path1, path2)
        if not os.path.isdir(path1) and os.path.isdir(path2):
            return '{} is a file but {} is a folder'.format(path1, path2)
        if os.path.isdir(path1):
            return compare_dir(path1, path2)

    return 'success'


copy_folder(sys.argv[1], sys.argv[2])
info = compare_dir(sys.argv[1], sys.argv[2])
print('comparator output: {}'.format(info))


# os.path() 模块
# os.path 模块主要用于获取文件的属性，以下是 os.path 模块的几种常用方法：
# 1.os.path.abspath(path) 返回绝对路径
# 2.os.path.basename(path) 返回文件名
# 3.os.path.commonprefix(list) 返回 list (多个路径)中，所有 path 共有的最长的路径
# 4.os.path.dirname(path) 返回文件路径
# 5.os.path.exists(path) 路径存在则返回 True，路径损坏返回 False
# 6.os.path.lexists 路径存在则返回 True，路径损坏也返回 True
# 以下实例演示了 os.path 相关方法的使用：

# 示例1：
# import os.path
# # 当前文件名
# print(__file__)
# # 当前文件名的绝对路径
# print(os.path.abspath(__file__))
# # 返回当前文件的路径
# print(os.path.dirname(os.path.abspath(__file__)))
# 以上实例编译运行结果如下：
# test.py
# /runoob/runoob-test-py/test.py
# /runoob/runoob-test-py


# 示例2：
# import os
# import time
# file='/root/runoob.txt' # 文件路径
# print( os.path.getatime(file) )   # 输出最近访问时间
# print( os.path.getctime(file) )   # 输出文件创建时间
# print( os.path.getmtime(file) )   # 输出最近修改时间
# print( time.gmtime(os.path.getmtime(file)) )  # 以 struct_time 形式输出最近修改时间
# print( os.path.getsize(file) )   # 输出文件大小（字节为单位）
# print( os.path.abspath(file) )   # 输出绝对路径
# print( os.path.normpath(file) )  # 规范 path 字符串形式
# 以上实例编译运行结果如下：
# 1539052805.5735736
# 1539052805.5775735
# 1539052805.5735736
# time.struct_time(tm_year = 2018, tm_mon =10, tm_mday = 9, tm_hour = 2, tm_min = 40, tm_sec = 5, tm_wday = 1, tm_yday = 282, tm_isdst = 0)
# /root/runoob.txt/root/runoob.txt

# shutil.copytree()
# 将以 src 为根起点的整个目录树拷贝到名为 dst 的目录并返回目标目录。dirs_exist_ok 指明是否要在 dst 或任何丢失的父目录已存在的情况下引发异常。该方法的语法如下：
# shutil.copytree(src, dst, symlinks = False, ignore = None, copy_function=copy2, ignore_dangling_symlinks = False, dirs_exist_ok = False)
# 示例：
# import shutil,os
# folder1 = os.path.join(os.getcwd(), 'aaa')
# folder2 = os.path.join(os.getcwd(), 'bbb', 'ccc')
# # 将 'abc.txt', 'bcd.txt' 忽略，不复制
# shutil.copytree(folder1, folder2, ignore = shutil.ignore_patterns('abc.txt', 'bcd.txt')