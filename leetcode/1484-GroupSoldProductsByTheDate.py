# 1484. Group Sold Products By The Date
# Table Activities:
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | sell_date   | date    |
# | product     | varchar |
# +-------------+---------+
# There is no primary key for this table, it may contain duplicates.
# Each row of this table contains the product name and the date it was sold in a market.

# Write an SQL query to find for each date the number of different products sold and their names.
# The sold products names for each date should be sorted lexicographically.
# Return the result table ordered by sell_date.
# The query result format is in the following example.

# Example 1:
# Input:
# Activities table:
# +------------+------------+
# | sell_date  | product     |
# +------------+------------+
# | 2020-05-30 | Headphone  |
# | 2020-06-01 | Pencil     |
# | 2020-06-02 | Mask       |
# | 2020-05-30 | Basketball |
# | 2020-06-01 | Bible      |
# | 2020-06-02 | Mask       |
# | 2020-05-30 | T-Shirt    |
# +------------+------------+
# Output:
# +------------+----------+------------------------------+
# | sell_date  | num_sold | products                     |
# +------------+----------+------------------------------+
# | 2020-05-30 | 3        | Basketball,Headphone,T-shirt |
# | 2020-06-01 | 2        | Bible,Pencil                 |
# | 2020-06-02 | 1        | Mask                         |
# +------------+----------+------------------------------+
# Explanation:
# For 2020-05-30, Sold items were (Headphone, Basketball, T-shirt), we sort them lexicographically and separate them by a comma.
# For 2020-06-01, Sold items were (Pencil, Bible), we sort them lexicographically and separate them by a comma.
# For 2020-06-02, the Sold item is (Mask), we just return it.

import pandas as pd

def categorize_products(activities: pd.DataFrame) -> pd.DataFrame:
    # 按照 sell_date 进行分组，并对产品进行去重和计数
    result = activities.groupby('sell_date')['product'].agg(['nunique', lambda x: ','.join(sorted(set(x)))])
    # 重命名列名
    result.rename(columns={'nunique': 'num_sold', '<lambda_0>': 'products'}, inplace = True)
    # 重置索引，将 sell_date 变为普通列
    result.reset_index(inplace  = True)
    # 按照 sell_date 进行升序排序
    result.sort_values(by='sell_date', inplace=True)
    return result

# agg
def categorize_products1(activities: pd.DataFrame) -> pd.DataFrame:
    activities = activities.sort_values("product").groupby("sell_date").agg(
        num_sold = pd.NamedAgg(column="product", aggfunc=lambda x: x.nunique()),
        products = pd.NamedAgg(column="product", aggfunc=lambda x: ",".join(sorted(set(x))))
    ).reset_index()
    return activities


if __name__ == "__main__":
    data = [['2020-05-30', 'Headphone'], ['2020-06-01', 'Pencil'], ['2020-06-02', 'Mask'], ['2020-05-30', 'Basketball'], ['2020-06-01', 'Bible'], ['2020-06-02', 'Mask'], ['2020-05-30', 'T-Shirt']]
    activities = pd.DataFrame(data, columns=['sell_date', 'product']).astype({'sell_date':'datetime64[ns]', 'product':'object'})
    print(categorize_products(activities))
    print(categorize_products1(activities))