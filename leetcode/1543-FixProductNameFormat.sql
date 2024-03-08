-- 1543. Fix Product Name Format
-- Table: Sales
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | sale_id      | int     |
-- | product_name | varchar |
-- | sale_date    | date    |
-- +--------------+---------+
-- sale_id is the column with unique values for this table.
-- Each row of this table contains the product name and the date it was sold.
 
-- Since table Sales was filled manually in the year 2000, product_name may contain leading and/or trailing white spaces, also they are case-insensitive.

-- Write a solution to report
-- product_name in lowercase without leading or trailing white spaces.
-- sale_date in the format ('YYYY-MM').
-- total the number of times the product was sold in this month.
-- Return the result table ordered by product_name in ascending order. In case of a tie, order it by sale_date in ascending order.

-- The result format is in the following example.

-- Example 1:

-- Input: 
-- Sales table:
-- +---------+--------------+------------+
-- | sale_id | product_name | sale_date  |
-- +---------+--------------+------------+
-- | 1       | LCPHONE      | 2000-01-16 |
-- | 2       | LCPhone      | 2000-01-17 |
-- | 3       | LcPhOnE      | 2000-02-18 |
-- | 4       | LCKeyCHAiN   | 2000-02-19 |
-- | 5       | LCKeyChain   | 2000-02-28 |
-- | 6       | Matryoshka   | 2000-03-31 |
-- +---------+--------------+------------+
-- Output: 
-- +--------------+-----------+-------+
-- | product_name | sale_date | total |
-- +--------------+-----------+-------+
-- | lckeychain   | 2000-02   | 2     |
-- | lcphone      | 2000-01   | 2     |
-- | lcphone      | 2000-02   | 1     |
-- | matryoshka   | 2000-03   | 1     |
-- +--------------+-----------+-------+
-- Explanation: 
-- In January, 2 LcPhones were sold. Please note that the product names are not case sensitive and may contain spaces.
-- In February, 2 LCKeychains and 1 LCPhone were sold.
-- In March, one matryoshka was sold.

-- Create table If Not Exists Sales (sale_id int, product_name varchar(30), sale_date date)

-- Truncate table Sales
-- insert into Sales (sale_id, product_name, sale_date) values ('1', 'LCPHONE', '2000-01-16')
-- insert into Sales (sale_id, product_name, sale_date) values ('2', 'LCPhone', '2000-01-17')
-- insert into Sales (sale_id, product_name, sale_date) values ('3', 'LcPhOnE', '2000-02-18')
-- insert into Sales (sale_id, product_name, sale_date) values ('4', 'LCKeyCHAiN', '2000-02-19')
-- insert into Sales (sale_id, product_name, sale_date) values ('5', 'LCKeyChain', '2000-02-28')
-- insert into Sales (sale_id, product_name, sale_date) values ('6', 'Matryoshka', '2000-03-31')
SELECT
    product_name,
    sale_date,
    COUNT(*) AS total 
FROM 
( 
    SELECT 
        LOWER(TRIM(product_name)) AS product_name, -- product_name 可能包含前后空格，而且包含大小写
        DATE_FORMAT(sale_date,"%Y-%m") AS sale_date -- sale_date 格式为 ('YYYY-MM') 
    FROM
        Sales 
) AS s
GROUP BY
    product_name, sale_date
ORDER BY 
    product_name ASC, sale_date ASC -- 返回结果以 product_name 升序 排列，如果有排名相同，再以 sale_date 升序 排列

-- best solution
SELECT
    LOWER(TRIM(product_name)) AS product_name, -- product_name 可能包含前后空格，而且包含大小写
    DATE_FORMAT(sale_date,"%Y-%m") AS sale_date, -- sale_date 格式为 ('YYYY-MM') 
    COUNT(*) AS total 
FROM 
    Sales 
GROUP BY
    1, 2
ORDER BY 
    1, 2  -- 返回结果以 product_name 升序 排列，如果有排名相同，再以 sale_date 升序 排列
