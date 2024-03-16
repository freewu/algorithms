# 1683. Invalid Tweets
# Table: Tweets
# +----------------+---------+
# | Column Name    | Type    |
# +----------------+---------+
# | tweet_id       | int     |
# | content        | varchar |
# +----------------+---------+
# tweet_id is the primary key (column with unique values) for this table.
# This table contains all the tweets in a social media app.
 
# Write a solution to find the IDs of the invalid tweets. The tweet is invalid if the number of characters used in the content of the tweet is strictly greater than 15.
# Return the result table in any order.
# The result format is in the following example.

# Example 1:
# Input: 
# Tweets table:
# +----------+----------------------------------+
# | tweet_id | content                          |
# +----------+----------------------------------+
# | 1        | Vote for Biden                   |
# | 2        | Let us make America great again! |
# +----------+----------------------------------+
# Output: 
# +----------+
# | tweet_id |
# +----------+
# | 2        |
# +----------+
# Explanation: 
# Tweet 1 has length = 14. It is a valid tweet.
# Tweet 2 has length = 32. It is an invalid tweet.

import pandas as pd

def invalid_tweets(tweets: pd.DataFrame) -> pd.DataFrame:
    # 推文长度大于 15 
    filter = tweets['content'].apply(lambda x: len(str(x)) > 15)
    return tweets[filter][['tweet_id']]
    
if __name__ == "__main__":
    data = [[1, 'Vote for Biden'], [2, 'Let us make America great again!']]
    tweets = pd.DataFrame(data, columns=['tweet_id', 'content']).astype({'tweet_id':'Int64', 'content':'object'})
    print(invalid_tweets(tweets))

# apply
# apply() 使用时，通常放入一个 lambda 函数表达式、或一个函数作为操作运算，官方上给出DataFrame的 apply() 用法：
#
#     DataFrame.apply(self, func, axis=0, raw=False, result_type=None, args=(), **kwargs)
#
# 参数：
#     func：函数或 lambda 表达式,应用于每行或者每列
#     axis：{0 or ‘index’, 1 or ‘columns’}, 默认为0
#         0 or ‘index’: 表示函数处理的是每一列
#         1 or ‘columns’: 表示函数处理的是每一行
#     raw：bool 类型，默认为 False;
#         False ，表示把每一行或列作为 Series 传入函数中；
#         True，表示接受的是 ndarray 数据类型；
#     result_type：{‘expand’, ‘reduce’, ‘broadcast’, None}, default None
#         These only act when axis=1 (columns):
#             ‘expand’ : 列表式的结果将被转化为列。
#             ‘reduce’ : 如果可能的话，返回一个 Series，而不是展开类似列表的结果。这与 expand 相反。
#             ‘broadcast’ : 结果将被广播到 DataFrame 的原始形状，原始索引和列将被保留。
#     args: func 的位置参数
#         **kwargs：要作为关键字参数传递给 func 的其他关键字参数，1.3.0 开始支持

# 返回值：
#     Series 或者 DataFrame：沿数据的给定轴应用 func 的结果