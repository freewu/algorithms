# 2604 · View the space occupied by a file
# To view the space occupied by files using the du command, complete the following 3 steps in order.

# Step 1 query and print the size of all files in the current directory and all its subdirectories, in blocks
du 

# Step 2 print the size of the specified file code.txt in blocks
du code.txt

# Step 3 query and print the size of all files in the current directory and all its subdirectories, in K, M, G units
du -h

# du （英文全拼：disk usage）命令用于显示目录或文件的大小。
# du 会显示指定的目录或文件所占用的磁盘空间。

# 语法
# du [-abcDhHklmsSx][-L <符号连接>][-X <文件>][--block-size][--exclude=<目录或文件>][--max-depth=<目录层数>][--help][--version][目录或文件]
# 参数说明：

# -a -all 显示目录中个别文件的大小。
# -b -bytes 显示目录或文件大小时，以byte为单位。
# -c --total 除了显示个别目录或文件的大小外，同时也显示所有目录或文件的总和。
# -D --dereference-args 显示指定符号连接的源文件大小。
# -h --human-readable 以K，M，G为单位，提高信息的可读性。
# -H --si 与-h参数相同，但是K，M，G是以1000为换算单位。
# -k --kilobytes 以1024 bytes为单位。
# -l --count-links 重复计算硬件连接的文件。
# -L<符号连接> --dereference<符号连接> 显示选项中所指定符号连接的源文件大小。
# -m --megabytes 以1MB为单位。
# -s --summarize 仅显示总计。
# -S --separate-dirs 显示个别目录的大小时，并不含其子目录的大小。
# -x --one-file-xystem 以一开始处理时的文件系统为准，若遇上其它不同的文件系统目录则略过。
# -X<文件> --exclude-from=<文件> 在<文件>指定目录或文件。
# --exclude=<目录或文件> 略过指定的目录或文件。
# --max-depth=<目录层数> 超过指定层数的目录后，予以忽略。
# --help 显示帮助。
# --version 显示版本信息。

# 实例
# 显示目录或者文件所占空间:

# # du
# 608     ./test6
# 308     ./test4
# 4       ./scf/lib
# 4       ./scf/service/deploy/product
# 4       ./scf/service/deploy/info
# 12      ./scf/service/deploy
# 16      ./scf/service
# 4       ./scf/doc
# 4       ./scf/bin
# 32      ./scf
# 8       ./test3
# 1288    .
# 只显示当前目录下面的子目录的目录大小和当前目录的总的大小，最下面的1288为当前目录的总大小

# 显示指定文件所占空间

# # du log2012.log 
# 300     log2012.log
# 方便阅读的格式显示test目录所占空间情况：

# # du -h test
# 608K    test/test6
# 308K    test/test4
# 4.0K    test/scf/lib
# 4.0K    test/scf/service/deploy/product
# 4.0K    test/scf/service/deploy/info
# 12K     test/scf/service/deploy
# 16K     test/scf/service
# 4.0K    test/scf/doc
# 4.0K    test/scf/bin
# 32K     test/scf
# 8.0K    test/test3
# 1.3M    test