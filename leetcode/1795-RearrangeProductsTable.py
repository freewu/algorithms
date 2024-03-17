# 1795. Rearrange Products Table
# Table: Products
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | product_id  | int     |
# | store1      | int     |
# | store2      | int     |
# | store3      | int     |
# +-------------+---------+
# product_id is the primary key for this table.
# Each row in this table indicates the product's price in 3 different stores: store1, store2, and store3.
# If the product is not available in a store, the price will be null in that store's column.

# Write an SQL query to rearrange the Products table so that each row has (product_id, store, price).
# If a product is not available in a store, do not include a row with that product_id and store combination in the result table.
# Return the result table in any order.
# The query result format is in the following example.

# Example 1:
# Input:
# Products table:
# +------------+--------+--------+--------+
# | product_id | store1 | store2 | store3 |
# +------------+--------+--------+--------+
# | 0          | 95     | 100    | 105    |
# | 1          | 70     | null   | 80     |
# +------------+--------+--------+--------+
# Output:
# +------------+--------+-------+
# | product_id | store  | price |
# +------------+--------+-------+
# | 0          | store1 | 95    |
# | 0          | store2 | 100   |
# | 0          | store3 | 105   |
# | 1          | store1 | 70    |
# | 1          | store3 | 80    |
# +------------+--------+-------+
# Explanation:
# Product 0 is available in all three stores with prices 95, 100, and 105 respectively.
# Product 1 is available in store1 with price 70 and store3 with price 80. The product is not available in store2.

import pandas as pd

def rearrange_products_table(products: pd.DataFrame) -> pd.DataFrame:
    # 取 store1 的数据
    store1 = products[["product_id",'store1']]
    store1["store"] = "store1"
    # 改名
    store1 = store1.rename(columns={'store1': 'price'})
    # 取 store2 的数据
    store2 = products[["product_id",'store2']]
    store2["store"] = "store2"
    # 改名
    store2 = store2.rename(columns={'store2': 'price'})
    # 取 store3 的数据
    store3 = products[["product_id",'store3']]
    store3["store"] = "store3"
    # 改名
    store3 = store3.rename(columns={'store3': 'price'})
    # 联合起来
    result = pd.concat([pd.concat([store1,store2]) ,store3]) 
    # 去掉price 为 None 的数据
    return result.sort_values('product_id').dropna(subset=['price'])

def rearrange_products_table1(products: pd.DataFrame) -> pd.DataFrame:
    products = products.set_index('product_id').stack().reset_index()
    # print(products)
    products.columns = ['product_id','store','price']
    return products

def rearrange_products_table2(products: pd.DataFrame) -> pd.DataFrame:
    products = products.melt(id_vars='product_id',var_name='store',value_name='price')
    return products.dropna(axis=0)
    
    

if __name__ == "__main__":
    data = [[0, 95, 100, 105], [1, 70, None, 80]]
    products = pd.DataFrame(data, columns=['product_id', 'store1', 'store2', 'store3']).astype({'product_id':'Int64', 'store1':'Int64', 'store2':'Int64', 'store3':'Int64'})
    print(rearrange_products_table(products))

    print(rearrange_products_table1(products))
    print(rearrange_products_table2(products))