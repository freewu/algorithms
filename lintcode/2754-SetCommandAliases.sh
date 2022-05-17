# 2754 · Set command aliases
# Description
# The alias command is used to set the alias of the command. Please complete the following steps in order.

# Step 1 set the alias la for the ls sub_dir command
alias la='ls ./sub_dir'

# Step 2 run the command alias la
la

# Step 3 undo the alias la
unalias la

# ## 设置别名
#   alias 别名=’原命令 -选项/参数’
# 例如:
#   alias ll='ls -lt'1
# 这样设置了ls -lt命令的别名是ll，在终端输入ll时，则相当于输入了ls -lt命令
# 注意： 在定义别名时，等号两边不能有空格，否则shell不能决定您需要做什么。
# 仅在命令中包含空格或特殊字符时才需要引号。如果键入不带任何参数的alias 命令，将显示所有已定义的别名。

# ## 查看已经设置的别名列表
#   alias -p1

# ## 删除别名
#   unalias 别名1
# 例如：
#   unalias ll1

# ## 设置别名每次登入可用
# alias命令只作用于当次登入的操作。如果想每次登入都能使用这些命令的别名，则可以把相应的alias命令存放在 ~/.bashrc 文件中。
# 打开~/.bashrc文件，输入要设置的alias命令，保存，然后运行
# source ~/.bashrc1
# 如果这样还不行，表示没有~/.bash_profile文件或文件中没有执行~/.bashrc文件
# 可以在~/.bash_profile中加入命令 source ~/.bashrc 后保存
# 这样就可以每次登入后都可以使用设置好的命令别名。