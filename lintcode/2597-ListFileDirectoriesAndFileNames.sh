# 2597 · List file directories and file names
# Description
# To list all the directories and file names in the current directory using the ls command, follow these 3 steps.

# Step 1 list the files and directories in the current directory
ls

# Step 2 list the files and directories in the current directory, including hidden files
ls -a

# Step 3 list the files and directories in the current directory and their details
ls -l


# # 命令格式
# ls [选项] [文件名]
#   用来打印当前目录清单或者打印出指定目录下的文件及文件清单。ls命令在打印文件清单时，还可以查看文件权限、目录信息等等。
# # 命令参数
# 可以使用帮助命令查看详细命令参数：man ls或ls --help
# -a, --all
#       do not ignore entries starting with .（列出目录下的所有文件，包括以.开头的隐含文件）
# -A, --almost-all
#       do not list implied . and …(列出除了.及…以外的所有文件)
# –author
#       with -l, print the author of each file(打印出每一个文件的作者)
# -b, --escape
#       print octal escapes for nongraphic characters（不能输出的字符用反斜杠加字符编号的形式输出）
# –block-size=SIZE
#       use SIZE-byte blocks.（使用SIZE-byte的大小的字节为单位）
# -B, --ignore-backups
#       do not list implied entries ending with ~（不列出任何以波浪号结束的项目）
# -c 
#       with -lt: sort by, and show, ctime (time of last modification of file status information) 
#       with -l: show ctime and sort by name otherwise: sort by ctime
#      （输出文件的ctime（文件状态最后更改的时间），并根据ctime排序）
# -C list entries by columns（由上至下的列出项目）
# –color[=WHEN]
#       colorize the output. WHEN defaults to ‘always’ or can be ‘never’ or ‘auto’. More info below
#       控制是否使用色彩分辨文件。WHEN可以是always、never或者auto其中一个
# -d, --directory
#       list directory entries instead of contents, and do not dereference symbolic links（将目录像文件一样显示，而不显示目录下面的内容）
# -D, --dired
#       generate output designed for Emacs’ dired mode（产生适合Emacs的dired模 式使用的结果）
# -f do not sort, enable -aU, disable -ls --color（对输出的文件不进行排序，-aU选线生效，-ls选项失效）
# -F, --classify
#       append indicator (one of /=>@|) to entries（加上文件类型的指示符号（/=@| 其中一个））
# –file-type
#       likewise, except do not append ‘’(和-F一样，除了不追加“”)
# –format=WORD
#       across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical 
#       -C(–format=关键之，关键字可以是“across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C”)
# –full-time
#       like -l --time-style=full-iso(即 -l --time-style=full-iso)
# -g   
#        like -l, but do not list owner(像-l，但是不列出所有者)
# –group-directories-first
#       group directories before files.（组目录在文件目录之前）
#       augment with a --sort option, but any use of --sort=none (-U) disables grouping
#       （用-sort选项进行扩展，但任何使用-sort=none（-U）禁用组）
# -G, --no-group
#       in a long listing, don’t print group names(不列出任何有关于组的信息)
# -h, --human-readable
#       with -l, print sizes in human readable format (e.g., 1K 234M 2G)（类似于-l，以人容易理解的格式列出文件大小）
# –si 
#       likewise, but use powers of 1000 not 1024（像-h，但是文件大小取1000而不是1024）
# -H, --dereference-command-line
#       follow symbolic links listed on the command line(使用命令列中的符号链接指示的 真正目的地)
# –dereference-command-line-symlink-to-dir
#       follow each command line symbolic link that points to a directory(遵循指向目录的 每个命令行符号链接)
# –hide=PATTERN
#       do not list implied entries matching shell PATTERN (overridden by -a or -A)（不 要列出与shell模式匹配的隐含条目（由-a或-A重写））
# –indicator-style=WORD
#       append indicator with style WORD to entry names: none (default), slash (- p), 
#       file-type (–file-type), classify (-F)(指定在每个项目名称后加上指示符号<方 式>：none (default), slash (-p), file-type (–file-type), classify (-F))
# -i, --inode
#       print the index number of each file（输出每个文件的inode号）
# -I, --ignore=PATTERN
#       do not list implied entries matching shell PATTERN(不列出任何符合sell万用字 符<样式>的项目)
# -k 
#       like --block-size=1K(像 --block-size=1K，以k字节的形式表示文件的大小)
# -l 
#       use a long listing format(使用长格式，即列出文件详细信息)
# -L, --dereference
#       when showing file information for a symbolic link, 
#       show information for the file the link references rather than for the link itself(当显示符号连接的文件信息时，显示符号链接所指示的对象而并非符号链接本身的信息)
# -m 
#       fill width with a comma separated list of entries（所有项目以逗号隔开，并填满整行行宽）
# -n, --numeric-uid-gid
#       like -l, but list numeric user and group IDs(类似-l，用数字UID和GID代替名称)
# -N, --literal
#       print raw entry names (don’t treat e.g. control characters specially)(印出未经处理 的项目名称，（例如不特别处理控制字符）)
# -o like -l, 
#       but do not list group information(类似-l，但是不现实组信息)
# -p, --indicator-style=slash
#       append / indicator to directories(向目录追加/指示器)
# -q, --hide-control-chars
#       print ? instead of non graphic characters(以？字符代替无法打印的字符)
# –show-control-chars
#       show non graphic characters as-is (default unless program is ‘ls’ and output is a terminal)
#       (直接显示无法打印的字符（这是默认方式，除非调用的程序名称时‘ls’而且在终端机画面输出结果）)
# -Q, --quote-name
#       enclose entry names in double quotes（将项目名称打上双引号）\
# –quoting-style=WORD
#       use quoting style WORD for entry names: literal, locale, shell, shell-always, c, escape
#       (使用指定的quotin方式类型：literal, locale, shell, shell-always, c, escape)
# -r, --reverse
#       reverse order while sorting（以相反次序排列）
# -R, --recursive
#       list subdirectories recursively(同时列出所有子目录层)
# -s, --size
#       print the allocated size of each file, in blocks(以块大小为单位列出所有文件的大小)
# -S 
#    sort by file size(根据文件大小排序)
# –sort=WORD
#       sort by WORD instead of name: none -U, extension -X, size -S, time -t, version -v
#       (根据关键词代替名字排序：none -U, extension -X, size -S, time -t, version -v)
# –time=WORD
#       with -l, show time as WORD instead of modification time: atime -u, access -u, use -u, ctime -c, or status-c; use specified time as sort key 
#       if --sort=time(类似于-l，根据关键词显示时间代替默认的修改时间：atime -u, access -u, use -u, ctime -c, or status)
# –time-style=STYLE
#       with -l, show times using style STYLE: full-iso, long-iso, iso, locale, +FORMAT. FORMAT is interpreted like ‘date’; 
#       if FORMAT is FORMAT1FORMAT2, FORMAT1 applies to non-recent files and FORMAT2 to recent files; 
#       if STYLE is prefixed with ‘posix-’, STYLE takes effect only outside the POSIX locale(显示时间使用style：…………)
# -t 
#       sort by modification time（以文件修改时间排序）
# -T, --tabsize=COLS
#       assume tab stops at each COLS instead of 8(显示我文件或目录最后被访问的时间)
# -u 
#       with -lt: sort by, and show, access time  配合-lt：显示访问时间而依访问时间排序
#       with -l: show access time and sort by name otherwise: sort by access time 合-l:显示访问时间但根据名称排序；否则，根据访问时间排序
# -U 
#       do not sort; list entries in directory order(不进行排序，依系统原有的次序列出项目)
# -v 
#       natural sort of (version) numbers within text(根据文件版本进行排序)
# -w, --width=COLS
#       assume screen width instead of current value(自行指定屏幕的宽度而不使用目前的数值)
# -x 
#       list entries by lines instead of by columns（逐行列出项目而不是逐栏列出）
# -X 
#       sort alphabetically by entry extension(根据扩展名排序)
# -1 
#       list one file per line(每行只列出一个文件)
# –help 
#       显示此帮助信息并退出
# -version 
#       显示版本信息并退出
