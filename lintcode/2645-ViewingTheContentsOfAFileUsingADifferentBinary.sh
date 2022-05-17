# 2645 · Viewing the contents of a file using a different binary
# Here, you will complete a 5-step exercise to view the contents of a file.

# Step 1 use the cat command to view the contents of a file file
cat file

# Step 2 use the od command to view the contents of the file in octal format
od file

# Step 3 use the od command to view the contents of the file in decimal format
od -d file

# Step 4 use the od command to view the contents of a file in hexadecimal format
od -x file

# Step 5 use the od command to view the contents of a file without printing the shift value
od -A n file

# Linux od命令用于输出文件内容。
# od 指令会读取所给予的文件的内容，并将其内容以八进制字码呈现出来。
# 语法
# od [-abcdfhilovx][-A <字码基数>][-j <字符数目>][-N <字符数目>][-s <字符串字符数>][-t <输出格式>][-w <每列字符数>][--help][--version][文件...]
# 参数：
# -a 　此参数的效果和同时指定"-ta"参数相同。
# -A<字码基数> 　选择要以何种基数计算字码。
# -b 　此参数的效果和同时指定"-toC"参数相同。
# -c 　此参数的效果和同时指定"-tC"参数相同。
# -d 　此参数的效果和同时指定"-tu2"参数相同。
# -f 　此参数的效果和同时指定"-tfF"参数相同。
# -h 　此参数的效果和同时指定"-tx2"参数相同。
# -i 　此参数的效果和同时指定"-td2"参数相同。
# -j<字符数目> --skip-bytes=<字符数目> 　略过设置的字符数目。
# -l 　此参数的效果和同时指定"-td4"参数相同。
# -N<字符数目> --read-bytes=<字符数目> 　到设置的字符数目为止。
# -o 　此参数的效果和同时指定"-to2"参数相同。
# -s<字符串字符数> --strings=<字符串字符数> 　只显示符合指定的字符数目的字符串。
# -t<输出格式> --format=<输出格式> 　设置输出格式。
# -v --output-duplicates 　输出时不省略重复的数据。
# -w<每列字符数> --width=<每列字符数> 　设置每列的最大字符数。
# -x 　此参数的效果和同时指定"-h"参数相同。
# --help 　在线帮助。
# --version 　显示版本信息。

# ## 设置第一列偏移地址以十进制显示。
#       od -Ad testfile
# 偏移地址显示基数有：d for decimal, o for octal, x for hexadecimal or n for none。

# ## od命令不显示第一列偏移地址。
#   od -An testfile

# ## 以十六进制输出，默认以四字节为一组（一列）显示。
#   od -tx testfile

# ## 以十六进制输出，每列输出一字节。
#   od -tx1 testfile

# ## 显示ASCII字符和ASCII字符名称，注意换行符显示方式的区别。
#   # 显示ASCII字符
#   $ echo lvlv | od -a
#   0000000  l  v  l  v nl
#   0000005
#   # 显示ASCII字符名称
#   $ echo lvlv | od -tc
#   0000000  l  v  l  v \n
#   0000005

# ## 以十六进制显示的同时显示原字符。
#   $ echo lvlv | od -tcx1
#   0000000  l  v  l  v \n
#            6c 76 6c 76 0a
#   0000005

# ## 指定每行显示512字节。
#   od -w512 -tx1 testfile

# ## od命令输出时去除列与列之间的空格符。
# 当我们需要将文件内容显示为十六进制，需要输出连续的单个字节，每个字节以十六进制显示。
# 这时我们可以通过od命令将文件以单个字节为一组，十六进制输出在同一行，并去除每个字节之间的空格。
# 目前还不知道怎么通过指定od命令的相关选项去除列与列之间的空格，也许od命令本身并不支持。我的做法是：
# （8.1）使用-An不输出偏移地址；
# （8.2）使用-v输出时不省略重复的数据；
# （8.3）使用-tx1以单个字节为一组按照十六进制输出，-w1每列输出一个字节；
# （8.4）最后通过管道传递给awk的标准输入，通过awk不换行输出所有行，拼接为一行输出。
# 具体命令如下：
#   od -An -w1 -tx1 testfile|awk '{for(i=1;i<=NF;++i){printf "%s",$i}}'