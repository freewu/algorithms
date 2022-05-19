# 2123 · Get yesterday's date
# Description
# Write Python code that implements a function named yesterday. 
# This function will get the date of the day before the incoming date, and finally return that date.
# Write the code for the yesterday function in solution.py,
# and we will run your code in main.py by importing it to check if it does the above correctly.

# You need to import the appropriate module by yourself
# The input data is in string format
# Enter the year of the date as 1 ≤ YEAR ≤ 9999

# The evaluator will execute your code by executing python main.py {input_path}, and you can see how the code is running in main.py.

# Example 1
# If the input data is:
#   2021-05-17
# then the output is:
#   2021-05-16

# Example 2
# If the input data is:
#   2020-03-01
# then the output is:
#   2020-02-29

import sys
import datetime
# from solution import yesterday

# import the required modules
from datetime import datetime,timedelta

# complete custom functions
def yesterday(today: datetime) -> datetime:
    '''
    :param today: get the date to be queried
    :return: the date of yesterday
    '''
    # write your code here
    oneday = timedelta(days = 1)
    yesterday = today - oneday
    return yesterday

input_path = sys.argv[1]
with open(input_path, 'r', encoding="utf-8") as f:
    date_today = f.readline()

today = datetime.datetime.strptime(date_today, '%Y-%m-%d').date()
print(yesterday(today))


# datetime 模块
# datetime 模块提供了处理日期和时间的类，既有简单的方式，又有复杂的方式。
# 它虽然支持日期和时间算法，但其实现的重点是为输出格式化和操作提供高效的属性提取功能。下面了解一下 datetime 中定义的类。
# 1.datetime.date 表示日期，常用的属性有：year, month 和 day
# 2.datetime.time 表示时间，常用属性有：hour, minute, second, microsecond
# 3.datetime.datetime 表示日期时间
# 4.datetime.timedelta 表示两个 date、time、datetime 实例之间的时间间隔，分辨率（最小单位）可达到微秒
# 5.datetime.tzinfo 时区相关信息对象的抽象基类。它们由 datetime 和 time 类使用，以提供自定义时间的而调整。
# 6.datetime.timezone 实现 tzinfo 抽象基类的类，表示与 UTC 的固定偏移量

# timedelta
# 该函数表示两个时间的间隔，参数可选、默认值都为 0，该方法的语法为：
# datetime.timedelta(days=0, seconds=0, microseconds=0, milliseconds=0, minutes=0, hours=0, weeks=0)
# 比如要输出当前时间一小时之后的时间：

# 示例代码

# from datetime import datetime,timedelta
# time1 = datetime.now()
# print time1
# print time1.strftime("%y-%m-%d %H:%M:%S")
# print (time1+timedelta(hours=1)).strftime("%y-%m-%d %H:%M:%S")
# 以上实例编译运行结果如下：
# 2017-11-08 10:54:35.926000
# 17-11-08 10:54:35
# 17-11-08 09:54:35