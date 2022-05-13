# 2737 · Document content sorting(ii)
# The sort command is used to sort the contents of a text file, so complete the following steps in order.

# Step 1 use the sort command to sort the contents of the file file, processing only the alphabet, numbers and space characters and ignoring other characters
sort -d file

# Step 2 use the sort command to sort the contents of the file file, ignoring upper and lower case letters
sort -f file

# Step 3 use the sort command to sort the contents of a file file in reverse order
sort -r file

# Linux sort 命令用于将文本文件内容加以排序。

# sort 可针对文本文件的内容，以行为单位来排序。

# 语法
# sort [-bcdfimMnr][-o<输出文件>][-t<分隔字符>][+<起始栏位>-<结束栏位>][--help][--verison][文件][-k field1[,field2]]
# 参数说明：

# -b 忽略每行前面开始出的空格字符。
# -c 检查文件是否已经按照顺序排序。
# -d 排序时，处理英文字母、数字及空格字符外，忽略其他的字符。
# -f 排序时，将小写字母视为大写字母。
# -i 排序时，除了040至176之间的ASCII字符外，忽略其他的字符。
# -m 将几个排序好的文件进行合并。
# -M 将前面3个字母依照月份的缩写进行排序。
# -n 依照数值的大小排序。
# -u 意味着是唯一的(unique)，输出的结果是去完重了的。
# -o<输出文件> 将排序后的结果存入指定的文件。
# -r 以相反的顺序来排序。
# -t<分隔字符> 指定排序时所用的栏位分隔字符。
# +<起始栏位>-<结束栏位> 以指定的栏位来排序，范围由起始栏位到结束栏位的前一栏位。
# --help 显示帮助。
# --version 显示版本信息。
# [-k field1[,field2]] 按指定的列进行排序。
# 实例
# 在使用 sort 命令以默认的式对文件的行进行排序，使用的命令如下：

# sort testfile 
# sort 命令将以默认的方式将文本文件的第一列以 ASCII 码的次序排列，并将结果输出到标准输出。

# 使用 cat 命令显示 testfile 文件可知其原有的排序如下：

# $ cat testfile      # testfile文件原有排序  
# test 30  
# Hello 95  
# Linux 85 
# 使用 sort 命令重排后的结果如下：

# $ sort testfile # 重排结果  
# Hello 95  
# Linux 85  
# test 30 
# 使用 -k 参数设置对第二列的值进行重排，结果如下：

# $ sort testfile -k 2
# test 30  
# Linux 85 
# Hello 95  