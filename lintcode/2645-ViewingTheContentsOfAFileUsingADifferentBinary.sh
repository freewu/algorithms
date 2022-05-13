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