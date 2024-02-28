-- 1127. User Purchase Platform
-- Table: Spending
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | spend_date  | date    |
-- | platform    | enum    | 
-- | amount      | int     |
-- +-------------+---------+
-- The table logs the history of the spending of users that make purchases from an online shopping website that has a desktop and a mobile application.
-- (user_id, spend_date, platform) is the primary key (combination of columns with unique values) of this table.
-- The platform column is an ENUM (category) type of ('desktop', 'mobile').
 
-- Write a solution to find the total number of users and the total amount spent using the mobile only, the desktop only, and both mobile and desktop together for each date.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Spending table:
-- +---------+------------+----------+--------+
-- | user_id | spend_date | platform | amount |
-- +---------+------------+----------+--------+
-- | 1       | 2019-07-01 | mobile   | 100    |
-- | 1       | 2019-07-01 | desktop  | 100    |
-- | 2       | 2019-07-01 | mobile   | 100    |
-- | 2       | 2019-07-02 | mobile   | 100    |
-- | 3       | 2019-07-01 | desktop  | 100    |
-- | 3       | 2019-07-02 | desktop  | 100    |
-- +---------+------------+----------+--------+
-- Output: 
-- +------------+----------+--------------+-------------+
-- | spend_date | platform | total_amount | total_users |
-- +------------+----------+--------------+-------------+
-- | 2019-07-01 | desktop  | 100          | 1           |
-- | 2019-07-01 | mobile   | 100          | 1           |
-- | 2019-07-01 | both     | 200          | 1           |
-- | 2019-07-02 | desktop  | 100          | 1           |
-- | 2019-07-02 | mobile   | 100          | 1           |
-- | 2019-07-02 | both     | 0            | 0           |
-- +------------+----------+--------------+-------------+ 
-- Explanation: 
-- On 2019-07-01, user 1 purchased using both desktop and mobile, user 2 purchased using mobile only and user 3 purchased using desktop only.
-- On 2019-07-02, user 2 purchased using mobile only, user 3 purchased using desktop only and no one purchased using both platforms.

SELECT 
    t2.spend_date AS spend_date, 
    t2.platform AS platform, 
    IFNULL(SUM(amount), 0) AS total_amount,
    IFNULL(COUNT(DISTINCT user_id), 0) AS total_users
FROM
    ( -- 生成一个 每天 / 每平台 表
        ( 
            SELECT 
                DISTINCT spend_date,
                "desktop" AS platform
            FROM 
                Spending
        ) UNION (
            SELECT 
                DISTINCT spend_date,
                "mobile" AS platform
            FROM 
                Spending
        ) UNION (
            SELECT 
                DISTINCT spend_date,
                "both" AS platform
            FROM 
                Spending
        ) 
    ) AS t2
LEFT JOIN 
    (   -- 每用户 / 每日 / 每平台 / 总金额
        SELECT 
            spend_date, 
            user_id, 
            SUM(amount) AS amount,
            IF(COUNT(*) = 1,platform,'both') AS platform -- 如果统计为 2为多平台
        FROM 
            Spending 
        GROUP BY 
            spend_date, user_id
    ) AS t1
ON 
    t2.spend_date = t1.spend_date AND t2.platform = t1.platform
GROUP BY  
    t2.spend_date, t2.platform


-- best solution
with aggr as ( -- 统计 每用户 / 每日 / 平台数 / 订单总数
    select 
        spend_date,
        user_id,
        count(distinct platform) as cnt,
        sum(amount) as amount
    from 
        spending
    group by
        user_id,
        spend_date
),
wide_table as ( -- 为 aggr 加上 平台名称
    select distinct 
        a.user_id,
        a.spend_date,
        case 
            when cnt = 1 then b.platform 
            when cnt = 2 then 'both' -- cnt 为 2 说明为 both 两个平台都下单
            else 'unknown'
        end as 'platform',
        a.amount as total_amount
    from 
        aggr as a 
    join 
        spending as b
    on 
        a.user_id = b.user_id and a.spend_date = b.spend_date
),
aggr_final as ( -- 统计每日 / 每平台 的数据
    select 
        spend_date,
        platform,
        sum(total_amount) as total_amount,
        count(user_id) as total_users
    from 
        wide_table
    group by 
        spend_date,
        platform
    order by 
        spend_date,
        platform
),
dummy as ( -- 生成一个 spend_date + platform 的平台表 
    select 
        distinct spend_date,
        'mobile' as platform,
        0 as total_amount,
        0 as total_users
    from aggr
    union
    select 
        distinct
        spend_date,
        'desktop' as platform,
        0 as total_amount,
        0 as total_users
    from aggr
    union
    select 
        distinct
        spend_date,
        'both' as platform,
        0 as total_amount,
        0 as total_users
    from 
        aggr
),
resutls as ( -- 通过 dummy 补全 aggr_final 为空的数据 
    select 
        b.spend_date,
        b.platform,
        coalesce(a.total_amount, b.total_amount) as total_amount,
        coalesce(a.total_users, b.total_users) as total_users
    from 
        aggr_final as a 
    right join 
        dummy as b 
    on a.spend_date=b.spend_date and a.platform=b.platform
)
select * from resutls