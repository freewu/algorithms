# 2755 · Display System Information
# Description
# The uname command is used to display system information, please complete the following steps in order.

# Step 1 display the operating system name
# Step 2 display the operating system version
# Step 3 display all information about the system

# Linux uname（英文全拼：unix name）命令用于显示系统信息。
# uname 可显示电脑以及操作系统的相关信息。

# ## 语法
#   uname [-amnrsv][--help][--version]
# ## 参数说明：
#   -a --all 　显示全部的信息。
#   -m --machine 　显示电脑类型。
#   -n --nodename 　显示在网络上的主机名称。
#   -r --release 　显示操作系统的发行编号。
#   -s --sysname 　显示操作系统名称。
#   -v 　显示操作系统的版本。
#   --help 　显示帮助。
#   --version 　显示版本信息。
# 实例

# 显示系统信息：
#       uname -a
#       Linux iZbp19byk2t6khuqj437q6Z 4.11.0-14-generic #20~16.04.1-Ubuntu SMP Wed Aug 9 09:06:22 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux

# 显示计算机类型：
#       uname -m
#       x86_64

# 显示计算机名：
#       uname -n
#       runoob-linux

# 显示操作系统发行编号：
#       uname -r
#       4.11.0-14-generic

# 显示操作系统名称：
#       uname -s
#       Linux

# 显示系统版本与时间：
#       uname -v
#       20~16.04.1-Ubuntu SMP Wed Aug 9 09:06:22 UTC 2017