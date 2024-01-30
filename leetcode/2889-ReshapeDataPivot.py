# 2889. Reshape Data: Pivot
# DataFrame weather
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | city        | object |
# | month       | object |
# | temperature | int    |
# +-------------+--------+
# Write a solution to pivot the data so that each row represents temperatures for a specific month, and each city is a separate column.
# The result format is in the following example.


# Example 1:
# Input:
# +--------------+----------+-------------+
# | city         | month    | temperature |
# +--------------+----------+-------------+
# | Jacksonville | January  | 13          |
# | Jacksonville | February | 23          |
# | Jacksonville | March    | 38          |
# | Jacksonville | April    | 5           |
# | Jacksonville | May      | 34          |
# | ElPaso       | January  | 20          |
# | ElPaso       | February | 6           |
# | ElPaso       | March    | 26          |
# | ElPaso       | April    | 2           |
# | ElPaso       | May      | 43          |
# +--------------+----------+-------------+
# Output:
# +----------+--------+--------------+
# | month    | ElPaso | Jacksonville |
# +----------+--------+--------------+
# | April    | 2      | 5            |
# | February | 6      | 23           |
# | January  | 20     | 13           |
# | March    | 26     | 38           |
# | May      | 43     | 34           |
# +----------+--------+--------------+
# Explanation:
# The table is pivoted, each column represents a city, and each row represents a specific month.

import pandas as pd

def pivotTable(weather: pd.DataFrame) -> pd.DataFrame:
    # index：确定新 DataFrame 中的行。在本例中，我们使用原始 DataFrame 中的 month 列作为索引，这意味着我们的透视表将为 month 列中的每个唯一值都建立行。
    # columns：确定新 DataFrame 中的列。在这里，我们使用的是 city 列，这意味着我们的透视表将有一列对应于 city 列中的每个唯一值。
    # values：指定重塑表格时要使用的值。在本例中，我们使用原始 DataFrame 中的 temperature 列
    #return pd.DataFrame(weather.pivot(index='month', columns='city', values='temperature'))
    return weather.pivot(index='month', columns='city', values='temperature')
    

if __name__ == "__main__":
    l = [
        ["Jacksonville","January", 13 ],
        ["Jacksonville","February", 23 ],
        ["Jacksonville","March", 38 ],
        ["Jacksonville","April", 5  ],
        ["Jacksonville","May", 34 ],
        ["ElPaso","January", 20 ],
        ["ElPaso","February", 6  ],
        ["ElPaso","March", 26 ],
        ["ElPaso","April", 2  ],
        ["ElPaso","May", 43 ],
    ]
    print(pivotTable(pd.DataFrame(l,columns=["city","month","temperature"])))