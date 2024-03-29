-- 1479. Sales by Day of the Week
-- Table: Orders
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | order_id      | int     |
-- | customer_id   | int     |
-- | order_date    | date    | 
-- | item_id       | varchar |
-- | quantity      | int     |
-- +---------------+---------+
-- (ordered_id, item_id) is the primary key (combination of columns with unique values) for this table.
-- This table contains information on the orders placed.
-- order_date is the date item_id was ordered by the customer with id customer_id.
 
-- Table: Items
-- +---------------------+---------+
-- | Column Name         | Type    |
-- +---------------------+---------+
-- | item_id             | varchar |
-- | item_name           | varchar |
-- | item_category       | varchar |
-- +---------------------+---------+
-- item_id is the primary key (column with unique values) for this table.
-- item_name is the name of the item.
-- item_category is the category of the item.
 
-- You are the business owner and would like to obtain a sales report for category items and the day of the week.
-- Write a solution to report how many units in each category have been ordered on each day of the week.

-- Return the result table ordered by category.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Orders table:
-- +------------+--------------+-------------+--------------+-------------+
-- | order_id   | customer_id  | order_date  | item_id      | quantity    |
-- +------------+--------------+-------------+--------------+-------------+
-- | 1          | 1            | 2020-06-01  | 1            | 10          |
-- | 2          | 1            | 2020-06-08  | 2            | 10          |
-- | 3          | 2            | 2020-06-02  | 1            | 5           |
-- | 4          | 3            | 2020-06-03  | 3            | 5           |
-- | 5          | 4            | 2020-06-04  | 4            | 1           |
-- | 6          | 4            | 2020-06-05  | 5            | 5           |
-- | 7          | 5            | 2020-06-05  | 1            | 10          |
-- | 8          | 5            | 2020-06-14  | 4            | 5           |
-- | 9          | 5            | 2020-06-21  | 3            | 5           |
-- +------------+--------------+-------------+--------------+-------------+
-- Items table:
-- +------------+----------------+---------------+
-- | item_id    | item_name      | item_category |
-- +------------+----------------+---------------+
-- | 1          | LC Alg. Book   | Book          |
-- | 2          | LC DB. Book    | Book          |
-- | 3          | LC SmarthPhone | Phone         |
-- | 4          | LC Phone 2020  | Phone         |
-- | 5          | LC SmartGlass  | Glasses       |
-- | 6          | LC T-Shirt XL  | T-Shirt       |
-- +------------+----------------+---------------+
-- Output: 
-- +------------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
-- | Category   | Monday    | Tuesday   | Wednesday | Thursday  | Friday    | Saturday  | Sunday    |
-- +------------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
-- | Book       | 20        | 5         | 0         | 0         | 10        | 0         | 0         |
-- | Glasses    | 0         | 0         | 0         | 0         | 5         | 0         | 0         |
-- | Phone      | 0         | 0         | 5         | 1         | 0         | 0         | 10        |
-- | T-Shirt    | 0         | 0         | 0         | 0         | 0         | 0         | 0         |
-- +------------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
-- Explanation: 
-- On Monday (2020-06-01, 2020-06-08) were sold a total of 20 units (10 + 10) in the category Book (ids: 1, 2).
-- On Tuesday (2020-06-02) were sold a total of 5 units in the category Book (ids: 1, 2).
-- On Wednesday (2020-06-03) were sold a total of 5 units in the category Phone (ids: 3, 4).
-- On Thursday (2020-06-04) were sold a total of 1 unit in the category Phone (ids: 3, 4).
-- On Friday (2020-06-05) were sold 10 units in the category Book (ids: 1, 2) and 5 units in Glasses (ids: 5).
-- On Saturday there are no items sold.
-- On Sunday (2020-06-14, 2020-06-21) were sold a total of 10 units (5 +5) in the category Phone (ids: 3, 4).
-- There are no sales of T-shirts.

-- Create table If Not Exists Orders (order_id int, customer_id int, order_date date, item_id varchar(30), quantity int)
-- Create table If Not Exists Items (item_id varchar(30), item_name varchar(30), item_category varchar(30))
-- Truncate table Orders
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('1', '1', '2020-06-01', '1', '10')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('2', '1', '2020-06-08', '2', '10')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('3', '2', '2020-06-02', '1', '5')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('4', '3', '2020-06-03', '3', '5')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('5', '4', '2020-06-04', '4', '1')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('6', '4', '2020-06-05', '5', '5')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('7', '5', '2020-06-05', '1', '10')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('8', '5', '2020-06-14', '4', '5')
-- insert into Orders (order_id, customer_id, order_date, item_id, quantity) values ('9', '5', '2020-06-21', '3', '5')
-- Truncate table Items
-- insert into Items (item_id, item_name, item_category) values ('1', 'LC Alg. Book', 'Book')
-- insert into Items (item_id, item_name, item_category) values ('2', 'LC DB. Book', 'Book')
-- insert into Items (item_id, item_name, item_category) values ('3', 'LC SmarthPhone', 'Phone')
-- insert into Items (item_id, item_name, item_category) values ('4', 'LC Phone 2020', 'Phone')
-- insert into Items (item_id, item_name, item_category) values ('5', 'LC SmartGlass', 'Glasses')
-- insert into Items (item_id, item_name, item_category) values ('6', 'LC T-Shirt XL', 'T-shirt')

-- 整理原始数据
WITH week_item_detail AS
(
    SELECT
        CASE
            WHEN DAYOFWEEK(order_date) = 1 THEN 'Sunday'
            WHEN DAYOFWEEK(order_date) = 2 THEN 'Monday'
            WHEN DAYOFWEEK(order_date) = 3 THEN 'Tuesday'
            WHEN DAYOFWEEK(order_date) = 4 THEN 'Wednesday'
            WHEN DAYOFWEEK(order_date) = 5 THEN 'Thursday'
            WHEN DAYOFWEEK(order_date) = 6 THEN 'Friday'
            WHEN DAYOFWEEK(order_date) = 7 THEN 'Saturday'
        END AS week_day,
        o.item_id,
        i.item_category,
        quantity
    FROM
        Orders o
    LEFT JOIN
        Items i
    ON
        o.item_id = i.item_id
),
-- 构建星期临时表
t_week_day AS(
    SELECT
       0 AS Monday,
       0 AS Tuesday,
       0 AS Wednesday,
       0 AS Thursday,
       0 AS Friday,
       0 AS Saturday,
       0 AS Sunday
    FROM
        dual
),
-- 产品类目表
t_category AS (
    SELECT DISTINCT
        item_category
    FROM
        Items
),
-- 产品类目-星期组合模版
t_category_weekday AS (
    SELECT
        item_category,
        Monday,
        Tuesday,
        Wednesday,
        Thursday,
        Friday,
        Saturday,
        Sunday
    FROM
        t_week_day,t_category
),
-- 行转列后结果数据表
f_week_item_detail AS (
    SELECT
        t.item_category AS category,
        SUM(IF(week_day = 'Monday',quantity,0)) AS Monday,
        SUM(IF(week_day = 'Tuesday',quantity,0)) AS Tuesday,
        SUM(IF(week_day = 'Wednesday',quantity,0)) AS Wednesday,
        SUM(IF(week_day = 'Thursday',quantity,0)) AS Thursday,
        SUM(IF(week_day = 'Friday',quantity,0)) AS Friday,
        SUM(IF(week_day = 'Saturday',quantity,0)) AS Saturday,
        SUM(IF(week_day = 'Sunday',quantity,0)) AS Sunday
    FROM
    (
        SELECT
            week_day,
            item_category,
            SUM(quantity) AS quantity
        FROM
            week_item_detail
        GROUP BY
            week_day,
            item_category
    ) t
    GROUP BY
        category
)

-- 求最终结果
SELECT
    tc.item_category AS category,
    IFNULL(fw.Monday,0) AS Monday,
    IFNULL(fw.Tuesday,0) AS Tuesday,
    IFNULL(fw.Wednesday,0) AS Wednesday,
    IFNULL(fw.Thursday,0) AS Thursday,
    IFNULL(fw.Friday,0) AS Friday,
    IFNULL(fw.Saturday,0) AS Saturday,
    IFNULL(fw.Sunday,0) AS Sunday
FROM
    t_category_weekday tc
LEFT JOIN
    f_week_item_detail fw
ON
    tc.item_category = fw.category
ORDER BY
    category


SELECT 
    DISTINCT i.item_category AS Category,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 2 THEN o.quantity END), 0) AS Monday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 3 THEN o.quantity END), 0) AS Tuesday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 4 THEN o.quantity END), 0) AS Wednesday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 5 THEN o.quantity END), 0) AS Thursday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 6 THEN o.quantity END), 0) AS Friday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 7 THEN o.quantity END), 0) AS Saturday,
    IFNULL(SUM(CASE WHEN DAYOFWEEK(o.order_date) = 1 THEN o.quantity END), 0) AS Sunday
FROM 
    Orders AS o 
RIGHT JOIN 
    Items i
ON 
    o.item_id = i.item_id
GROUP BY
    1
ORDER BY
    1

-- 在MySQL中，DAYOFWEEK()是一个用来获取日期的星期几的函数。它返回一个数字，1 代表星期日，2 代表星期一，以此类推，7 代表星期六。
--     DAYOFWEEK(date)
-- 其中，date是一个合法的日期或日期时间表达式，比如一个日期常量、一个日期变量、一个日期列等