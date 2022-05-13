# 2632 · Intercept the string at the specified position
# Description
# Use the Linux command line to use cut to intercept a string at a specified position in a file named temp.txt. 
# You need to complete the exercise through the following four steps:

# Step 1 Use -b to intercept the 4th byte of each line of temp.txt
cut -b 4 temp.txt

# Step 2  Use -b to intercept the first, second, and third bytes of each line of temp.txt
cut -b 1,2,3 temp.txt

# Step 3 Use -c to intercept the second character of each line of temp.txt
cut -c 2 temp.txt

# Step 4 Intercept the content before @ of each line of temp.txt.txt
cut -d @ -f 1 temp.txt


# Linux cut命令用于显示每行从开头算起 num1 到 num2 的文字。
# 语法
# cut [-bn] [file]
# cut [-c] [file]
# cut [-df] [file]
# 使用说明:
# cut 命令从文件的每一行剪切字节、字符和字段并将这些字节、字符和字段写至标准输出。
# 如果不指定 File 参数，cut 命令将读取标准输入。必须指定 -b、-c 或 -f 标志之一。
# 参数:
# -b ：以字节为单位进行分割。这些字节位置将忽略多字节字符边界，除非也指定了 -n 标志。
# -c ：以字符为单位进行分割。
# -d ：自定义分隔符，默认为制表符。
# -f ：与-d一起使用，指定显示哪个区域。
# -n ：取消分割多字节字符。仅和 -b 标志一起使用。如果字符的最后一个字节落在由 -b 标志的 List 参数指示的
# 范围之内，该字符将被写出；否则，该字符将被排除