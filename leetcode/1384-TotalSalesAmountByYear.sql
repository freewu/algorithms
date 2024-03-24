-- 1384. Total Sales Amount by Year
-- Table: Product
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | product_id    | int     |
-- | product_name  | varchar |
-- +---------------+---------+
-- product_id is the primary key (column with unique values) for this table.
-- product_name is the name of the product.
 
-- Table: Sales
-- +---------------------+---------+
-- | Column Name         | Type    |
-- +---------------------+---------+
-- | product_id          | int     |
-- | period_start        | date    |
-- | period_end          | date    |
-- | average_daily_sales | int     |
-- +---------------------+---------+
-- product_id is the primary key (column with unique values) for this table. 
-- period_start and period_end indicate the start and end date for the sales period, and both dates are inclusive.
-- The average_daily_sales column holds the average daily sales amount of the items for the period.
-- The dates of the sales years are between 2018 to 2020.
 
-- Write a solution to report the total sales amount of each item for each year, with corresponding product_name, product_id, report_year, and total_amount.
-- Return the result table ordered by product_id and report_year.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Product table:
-- +------------+--------------+
-- | product_id | product_name |
-- +------------+--------------+
-- | 1          | LC Phone     |
-- | 2          | LC T-Shirt   |
-- | 3          | LC Keychain  |
-- +------------+--------------+
-- Sales table:
-- +------------+--------------+-------------+---------------------+
-- | product_id | period_start | period_end  | average_daily_sales |
-- +------------+--------------+-------------+---------------------+
-- | 1          | 2019-01-25   | 2019-02-28  | 100                 |
-- | 2          | 2018-12-01   | 2020-01-01  | 10                  |
-- | 3          | 2019-12-01   | 2020-01-31  | 1                   |
-- +------------+--------------+-------------+---------------------+
-- Output: 
-- +------------+--------------+-------------+--------------+
-- | product_id | product_name | report_year | total_amount |
-- +------------+--------------+-------------+--------------+
-- | 1          | LC Phone     |    2019     | 3500         |
-- | 2          | LC T-Shirt   |    2018     | 310          |
-- | 2          | LC T-Shirt   |    2019     | 3650         |
-- | 2          | LC T-Shirt   |    2020     | 10           |
-- | 3          | LC Keychain  |    2019     | 31           |
-- | 3          | LC Keychain  |    2020     | 31           |
-- +------------+--------------+-------------+--------------+
-- Explanation: 
-- LC Phone was sold for the period of 2019-01-25 to 2019-02-28, and there are 35 days for this period. Total amount 35*100 = 3500. 
-- LC T-shirt was sold for the period of 2018-12-01 to 2020-01-01, and there are 31, 365, 1 days for years 2018, 2019 and 2020 respectively.
-- LC Keychain was sold for the period of 2019-12-01 to 2020-01-31, and there are 31, 31 days for years 2019 and 2020 respectively.

-- Create table If Not Exists Product (product_id int, product_name varchar(30))
-- Create table If Not Exists Sales (product_id int, period_start date, period_end date, average_daily_sales int)
-- Truncate table Product
-- insert into Product (product_id, product_name) values ('1', 'LC Phone ')
-- insert into Product (product_id, product_name) values ('2', 'LC T-Shirt')
-- insert into Product (product_id, product_name) values ('3', 'LC Keychain')
-- Truncate table Sales
-- insert into Sales (product_id, period_start, period_end, average_daily_sales) values ('1', '2019-01-25', '2019-02-28', '100')
-- insert into Sales (product_id, period_start, period_end, average_daily_sales) values ('2', '2018-12-01', '2020-01-01', '10')
-- insert into Sales (product_id, period_start, period_end, average_daily_sales) values ('3', '2019-12-01', '2020-01-31', '1')

SELECT
    r.product_id,
    p.product_name,
    r.report_year,
    r.total_amount
FROM
    (
        (-- 2018 年的汇总
            SELECT 
                product_id,
                "2018" AS report_year,
                SUM(
                    CASE 
                        -- start & end 都在 2018 范围内 [2018,2018]
                        WHEN YEAR(period_start) = 2018 AND YEAR(period_end) = 2018 THEN DATEDIFF(period_end, period_start)+1  
                        -- start 在 2018-01-01 之前 end 在 2018 范围内 [2017,2018]
                        WHEN YEAR(period_start) < 2018 AND YEAR(period_end) = 2018 THEN DATEDIFF(period_end, "2018-01-01")+1
                        -- start 在 2018 范围内 end 在 2018 之后 [2018,2021]
                        WHEN YEAR(period_start) = 2018 AND YEAR(period_end) > 2018 THEN DATEDIFF("2018-12-31", period_start)+1
                        -- start 在 2018-01-01 之前 end 在 2018 之后 [2017 - 2020]
                        WHEN YEAR(period_start) < 2018 AND YEAR(period_end) > 2018 THEN 365 
                    END
                ) * average_daily_sales AS total_amount
            FROM
                Sales
            WHERE 
                YEAR(period_end) >= 2018
            GROUP BY 
                product_id
        )
        UNION ALL 
        (-- 2019 年的汇总
            SELECT 
                product_id,
                "2019" AS report_year,
                SUM(
                    CASE 
                        -- start & end 都在 2019 范围内 [2019,2019]
                        WHEN YEAR(period_start) = 2019 AND YEAR(period_end) = 2019 THEN DATEDIFF(period_end, period_start)+1  
                        -- start 在 2019-01-01 之前 end 在 2019 范围内 [2018,2019]
                        WHEN YEAR(period_start) < 2019 AND YEAR(period_end) = 2019 THEN DATEDIFF(period_end, "2019-01-01")+1
                        -- start 在 2019 范围内 end 在 2019 之后 [2019,2021]
                        WHEN YEAR(period_start) = 2019 AND YEAR(period_end) > 2019 THEN DATEDIFF("2019-12-31", period_start)+1
                        -- start 在 2019-01-01 之前 end 在 2019 之后 [2018 - 2020]
                        WHEN YEAR(period_start) < 2019 AND YEAR(period_end) > 2019 THEN 365 
                    END
                ) * average_daily_sales AS total_amount
            FROM
                Sales
            WHERE 
                YEAR(period_end) >= 2019
            GROUP BY 
                product_id
        )
        UNION ALL 
        (-- 2020 年的汇总
            SELECT 
                product_id,
                "2020" AS report_year,
                SUM(
                    CASE 
                        -- start & end 都在 2020 范围内 [2020,2020]
                        WHEN YEAR(period_start) = 2020 AND YEAR(period_end) = 2020 THEN DATEDIFF(period_end, period_start)+1  
                        -- start 在 2020-01-01 之前 end 在 2020 范围内 [2018,2020]
                        WHEN YEAR(period_start) < 2020 AND YEAR(period_end) = 2020 THEN DATEDIFF(period_end, "2020-01-01")+1
                        -- start 在 2020 范围内 end 在 2020 之后 [2020,2021]
                        WHEN YEAR(period_start) = 2020 AND YEAR(period_end) > 2020 THEN DATEDIFF("2020-12-31", period_start) + 1
                        -- start 在 2020-01-01 之前 end 在 2020 之后 [2018 - 2021]
                        WHEN YEAR(period_start) < 2020 AND YEAR(period_end) > 2020 THEN 365 
                    END
                ) * average_daily_sales AS total_amount
            FROM
                Sales
            WHERE 
                YEAR(period_end) >= 2020
            GROUP BY 
                product_id
        )
    ) AS r
LEFT JOIN 
    Product AS p
ON  
    p.product_id = r.product_id 
WHERE
    r.total_amount IS NOT NULL
ORDER BY
    r.product_id, r.report_year -- 结果并按 product_id 和 report_year 排序


with recursive year_info as (
    select 
        min(year(period_start)) as start_year,
        max(year(period_end)) as end_year
    from
        sales
),
all_years as (
    select start_year sale_year from year_info
    union 
    select sale_year + 1 
    from 
        all_years as a
    join 
        year_info as b
    where 
        a.sale_year <= b.end_year
),
all_year_info as (
    select 
        concat(sale_year, '-01-01') as start_date,
        concat(sale_year, '-12-31') as end_date
    from
        all_years
),
sales_years as (
    select
        case when s.period_start <= a.start_date then a.start_date
             else s.period_start
        end year_start_date,
        case when s.period_end >= a.end_date then a.end_date
             else s.period_end
        end year_end_date,
        s.product_id,
        s.average_daily_sales
    from 
        sales as s 
    cross join 
        all_year_info as a 
    where
        (year(a.start_date) >= year(s.period_start) and year(a.start_date) <= year(s.period_end)) or 
        (year(a.end_date) >= year(s.period_start) and year(a.end_date) <= year(s.period_end))
),
report as (
    select 
        s.product_id,
        p.product_name,
        cast(year(s.year_start_date) as char) as report_year,
        (datediff(s.year_end_date, s.year_start_date) + 1) * s.average_daily_sales as total_amount
    from 
        sales_years as s
    join 
        product as p
    on 
        s.product_id = p.product_id
    order by
        product_id,
        report_year
)
select * from report