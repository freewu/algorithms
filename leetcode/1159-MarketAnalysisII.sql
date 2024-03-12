-- 1159. Market Analysis II
-- Table: Users
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | user_id        | int     |
-- | join_date      | date    |
-- | favorite_brand | varchar |
-- +----------------+---------+
-- user_id is the primary key (column with unique values) of this table.
-- This table has the info of the users of an online shopping website where users can sell and buy items.

-- Table: Orders
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | order_id      | int     |
-- | order_date    | date    |
-- | item_id       | int     |
-- | buyer_id      | int     |
-- | seller_id     | int     |
-- +---------------+---------+
-- order_id is the primary key (column with unique values) of this table.
-- item_id is a foreign key (reference column) to the Items table.
-- buyer_id and seller_id are foreign keys to the Users table.
 
-- Table: Items
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | item_id       | int     |
-- | item_brand    | varchar |
-- +---------------+---------+
-- item_id is the primary key (column with unique values) of this table.

-- Write a solution to find for each user whether the brand of the second item (by date) they sold is their favorite brand. 
-- If a user sold less than two items, report the answer for that user as no. 
-- It is guaranteed that no seller sells more than one item in a day.
-- Return the result table in any order.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Users table:
-- +---------+------------+----------------+
-- | user_id | join_date  | favorite_brand |
-- +---------+------------+----------------+
-- | 1       | 2019-01-01 | Lenovo         |
-- | 2       | 2019-02-09 | Samsung        |
-- | 3       | 2019-01-19 | LG             |
-- | 4       | 2019-05-21 | HP             |
-- +---------+------------+----------------+
-- Orders table:
-- +----------+------------+---------+----------+-----------+
-- | order_id | order_date | item_id | buyer_id | seller_id |
-- +----------+------------+---------+----------+-----------+
-- | 1        | 2019-08-01 | 4       | 1        | 2         |
-- | 2        | 2019-08-02 | 2       | 1        | 3         |
-- | 3        | 2019-08-03 | 3       | 2        | 3         |
-- | 4        | 2019-08-04 | 1       | 4        | 2         |
-- | 5        | 2019-08-04 | 1       | 3        | 4         |
-- | 6        | 2019-08-05 | 2       | 2        | 4         |
-- +----------+------------+---------+----------+-----------+
-- Items table:
-- +---------+------------+
-- | item_id | item_brand |
-- +---------+------------+
-- | 1       | Samsung    |
-- | 2       | Lenovo     |
-- | 3       | LG         |
-- | 4       | HP         |
-- +---------+------------+
-- Output: 
-- +-----------+--------------------+
-- | seller_id | 2nd_item_fav_brand |
-- +-----------+--------------------+
-- | 1         | no                 |
-- | 2         | yes                |
-- | 3         | yes                |
-- | 4         | no                 |
-- +-----------+--------------------+
-- Explanation: 
-- The answer for the user with id 1 is no because they sold nothing.
-- The answer for the users with id 2 and 3 is yes because the brands of their second sold items are their favorite brands.
-- The answer for the user with id 4 is no because the brand of their second sold item is not their favorite brand.

-- Create table If Not Exists Users (user_id int, join_date date, favorite_brand varchar(10));
-- Create table If Not Exists Orders (order_id int, order_date date, item_id int, buyer_id int, seller_id int);
-- Create table If Not Exists Items (item_id int, item_brand varchar(10));
-- Truncate table Users;
-- insert into Users (user_id, join_date, favorite_brand) values ('1', '2019-01-01', 'Lenovo');
-- insert into Users (user_id, join_date, favorite_brand) values ('2', '2019-02-09', 'Samsung');
-- insert into Users (user_id, join_date, favorite_brand) values ('3', '2019-01-19', 'LG');
-- insert into Users (user_id, join_date, favorite_brand) values ('4', '2019-05-21', 'HP');
-- Truncate table Orders;
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('1', '2019-08-01', '4', '1', '2');
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('2', '2019-08-02', '2', '1', '3');
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('3', '2019-08-03', '3', '2', '3');
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('4', '2019-08-04', '1', '4', '2');
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('5', '2019-08-04', '1', '3', '4');
-- insert into Orders (order_id, order_date, item_id, buyer_id, seller_id) values ('6', '2019-08-05', '2', '2', '4');
-- Truncate table Items;
-- insert into Items (item_id, item_brand) values ('1', 'Samsung');
-- insert into Items (item_id, item_brand) values ('2', 'Lenovo');
-- insert into Items (item_id, item_brand) values ('3', 'LG');
-- insert into Items (item_id, item_brand) values ('4', 'HP');

-- 使用变量
SELECT 
    u.user_id AS seller_id, 
    -- 如果用户售出的商品少于两件，则该用户的结果为否 r2.item_brand IS NULL
    IF (t.item_brand IS NULL || t.item_brand != u.favorite_brand, "no", "yes") AS 2nd_item_fav_brand
FROM 
    Users AS u
LEFT JOIN 
(-- 取每个用户购买的第二件商品的品牌
    SELECT 
        r.seller_id, 
        i.item_brand 
    FROM 
    ( -- 每个用户购买的记录打上编号
        SELECT 
            @rk := if (@seller = a.seller_id, @rk + 1, 1) as "rank",
            @seller := a.seller_id as seller_id, 
            a.item_id
        FROM 
        (
            SELECT 
                o.seller_id, 
                o.item_id
            FROM 
                Orders AS o
            ORDER BY 
                seller_id, order_date
        ) AS a, 
        (
            select @seller := -1, @rk := 0
        ) AS b
    ) AS r
    JOIN items AS i 
    ON 
        r.item_id = i.item_id
    WHERE 
        r.rank = 2 
) AS t
ON u.user_id = t.seller_id;


-- # mysql中的排名函数
-- 主要介绍一下mysql里面的排名函数，涉及到的函数有以下几个：
-- rank()、dense_rank()、row_number()

--# 准备工作
--## 建立一个rank表：
-- create table rank(
-- 	    id int(10) not null primary key,
-- 	    name varchar(20) not null,
-- 	    score int(10) not null
-- );
-- ## 插入一些数据：
-- insert into rank values(1,'a',100);
-- insert into rank values(2,'b',100);
-- insert into rank values(3,'c',95);
-- insert into rank values(4,'d',95);
-- insert into rank values(5,'e',95);
-- insert into rank values(6,'a',90);
-- insert into rank values(7,'a',89);

-- # rank() 函数
-- 语法结构：
-- RANK() OVER (
--     PARTITION BY <expression>[{,<expression>...}]
--     ORDER BY <expression> [ASC|DESC], [{,<expression>...}]
-- ) 

-- 按照某字段的排序结果添加排名，但它是跳跃的、间断的排名
-- （1）若按照分数直接进行排序的话，例如按照score进行排名
-- 两个并列第一名后，下一个是第三名。

-- SELECT score, rank() over(ORDER BY score desc) as 'Rank' FROM rank;
-- 结果：
-- +------+---------+
-- | score|   Rank  |
-- +------+---------+
-- |  100 |       1 |
-- |  100 |       1 |
-- |  95  |       3 |
-- |  95  |       3 |
-- |  95  |       3 |
-- |  90  |       6 |
-- |  89  |       7 |
-- +------+---------+
-- 7 rows in set (0.02 sec)

-- （2）若按照某个字段分区进行排序的话，例如按照name进行分区，根据分数进行排名：
-- SELECT name , 
-- 	score ,
-- 	rank() over(partition by name ORDER BY score desc) as 'Rank' 
-- FROM rank;
-- 首先，PARTITION BY子句按姓名将结果集分成多个分区。
-- 然后，ORDER BY子句按分数对结果集进行排序。
-- 结果：
-- +------+------+---------+
-- | name | score|   Rank  |
-- +------+------+---------+
-- |  a   |  100 |       1 |
-- |  a   |  90  |       2 |
-- |  a   |  89  |       3 |
-- |  b   |  100 |       1 |
-- |  c   |  95  |       1 |
-- |  d   |  95  |       1 |
-- |  e   |  95  |       1 |
-- +------+------+---------+
-- 7 rows in set (0.02 sec)

-- # row_number() 函数
-- MySQL ROW_NUMBER()从8.0版开始引入了功能。这ROW_NUMBER()是一个窗口函数或分析函数，它为从1开始应用的每一行分配一个序号
-- 语法结构如下：
-- ROW_NUMBER() OVER (
--     PARTITION BY <expression>[{,<expression>...}]
--     ORDER BY <expression> [ASC|DESC], [{,<expression>...}]

-- 例如还是根据分数进行排序
-- SELECT 
--  row_number() OVER (
--  ORDER BY score
--  ) row_num,
--  score
--  FROM rank;
-- 结果：
-- +-------+------+---------+
-- |row_num| score|   Rank  |
-- +------ +------+---------+
-- |  1    |  100 |       1 |
-- |  2    |  100 |       2 |
-- |  3    |  95  |       3 |
-- |  4    |  95  |       1 |
-- |  5    |  95  |       1 |
-- |  6    |  90  |       1 |
-- |  7    |  89  |       1 |
-- +-------+------+---------+
-- 7 rows in set (0.02 sec)
-- 其次，使用ROW_NUMBER()函数将行划分为所有列的分区。对于每个唯一的行集，将重新开始行号。

-- SELECT 
--     id,
--     name,
--     ROW_NUMBER() OVER (PARTITION BY name ORDER BY name) AS row_num
-- FROM rank; 

-- 结果：
-- +------+------+---------+
-- | id   | name | row_num |
-- +------+------+---------+
-- |    1 | a    |       1 |
-- |    2 | a    |       2 |
-- |    3 | a    |       3 |
-- |    4 | b    |       1 |
-- |    5 | c    |       1 |
-- |    6 | d    |       1 |
-- |    7 | e    |       1 |
-- +------+------+---------+
-- 7 rows in set (0.02 sec)

-- # dense_rank() 函数
-- dense 英语中指“稠密的、密集的”。dense_rank()是的排序数字是连续的、不间断。当有相同的分数时，它们的排名结果是并列的，例如，1,2,2,3。
-- 语法结构：
-- DENSE_RANK() OVER (
--     PARTITION BY <expression>[{,<expression>...}]
--     ORDER BY <expression> [ASC|DESC], [{,<expression>...}]
-- ) 

-- 例如，还是根据成绩进行排名：
-- SELECT score, dense_rank() over(ORDER BY score desc) as 'Rank' FROM rank;

-- 结果：
-- +------+---------+
-- | score|   Rank  |
-- +------+---------+
-- |  100 |       1 |
-- |  100 |       1 |
-- |  95  |       2 |
-- |  95  |       2 |
-- |  95  |       2 |
-- |  90  |       3 |
-- |  89  |       4 |
-- +------+---------+
-- 7 rows in set (0.02 sec)

-- 若按照某个字段分区进行排序的话，例如按照name进行分区，根据分数进行排名
-- SELECT 
--      name, 
-- 	    score,
-- 	    dense_rank() over(partition by name ORDER BY score desc) as 'Rank' 
-- 	FROM rank;
-- 首先，PARTITION BY子句按姓名将结果集分成多个分区。
-- 然后，ORDER BY子句按分数对结果集进行排名。
-- 结果：
-- +------+------+---------+
-- | name | score|   Rank  |
-- +------+------+---------+
-- |  a   |  100 |       1 |
-- |  a   |  90  |       2 |
-- |  a   |  89  |       3 |
-- |  b   |  100 |       1 |
-- |  c   |  95  |       1 |
-- |  d   |  95  |       1 |
-- |  e   |  95  |       1 |
-- +------+------+---------+
-- 7 rows in set (0.02 sec)

-- 这数据可能不太明显，如果再插入一条数据：
-- insert into rank values(8,'a',90);

-- 然后查询，结果如下，与rank函数执行的结果就可以看到区别了：
-- +------+------+---------+
-- | name | score|   Rank  |
-- +------+------+---------+
-- |  a   |  100 |       1 |
-- |  a   |  90  |       2 |
-- |  a   |  90  |       2 |
-- |  a   |  89  |       3 |
-- |  b   |  100 |       1 |
-- |  c   |  95  |       1 |
-- |  d   |  95  |       1 |
-- |  e   |  95  |       1 |
-- +------+------+---------+
-- 7 rows in set (0.02 sec)

-- 使用 rank 函数
SELECT 
    u.user_id AS seller_id, 
    -- 如果用户售出的商品少于两件，则该用户的结果为否 r2.item_brand IS NULL
    IF (t.item_brand IS NULL || t.item_brand != u.favorite_brand, "no", "yes") AS 2nd_item_fav_brand
FROM 
    Users AS u
LEFT JOIN 
(-- 每个用户 + 编号 
    SELECT
        o.seller_id,
        RANK() OVER(PARTITION BY o.seller_id ORDER BY o.order_date ASC) as "rk",
        i.item_brand 
    FROM
        Orders AS o,
        Items AS i
    WHERE 
        o.item_id = i.item_id
) AS t
ON
    t.seller_id = u.user_id AND
    t.rk = 2