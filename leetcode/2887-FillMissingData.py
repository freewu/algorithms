# 2887. Fill Missing Data
# DataFrame products
# +-------------+--------+
# | Column Name | Type   |
# +-------------+--------+
# | name        | object |
# | quantity    | int    |
# | price       | int    |
# +-------------+--------+
# Write a solution to fill in the missing value as 0 in the quantity column.
# The result format is in the following example.
 
# Example 1:
# Input:+-----------------+----------+-------+
# | name            | quantity | price |
# +-----------------+----------+-------+
# | Wristwatch      | None     | 135   |
# | WirelessEarbuds | None     | 821   |
# | GolfClubs       | 779      | 9319  |
# | Printer         | 849      | 3051  |
# +-----------------+----------+-------+
# Output:
# +-----------------+----------+-------+
# | name            | quantity | price |
# +-----------------+----------+-------+
# | Wristwatch      | 0        | 135   |
# | WirelessEarbuds | 0        | 821   |
# | GolfClubs       | 779      | 9319  |
# | Printer         | 849      | 3051  |
# +-----------------+----------+-------+
# Explanation: 
# The quantity for Wristwatch and WirelessEarbuds are filled by 0.
import pandas as pd

def fillMissingValues(products: pd.DataFrame) -> pd.DataFrame:
    products['quantity'].fillna(0,inplace=True)
    return products

if __name__ == "__main__":
    l = [
        [101, None, 15,1.0],
        [101, None, 11,2.0],
        [103, 3, 11,3.0],
        [104, 4, 20,4.0]
    ]
    print(fillMissingValues(pd.DataFrame(l,columns=["id","quantity","last","grade"])))