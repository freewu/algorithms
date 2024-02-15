-- 1294. Weather Type in Each Country
-- Table: Countries
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | country_id    | int     |
-- | country_name  | varchar |
-- +---------------+---------+
-- country_id is the primary key (column with unique values) for this table.
-- Each row of this table contains the ID and the name of one country.
 
-- Table: Weather
-- +---------------+------+
-- | Column Name   | Type |
-- +---------------+------+
-- | country_id    | int  |
-- | weather_state | int  |
-- | day           | date |
-- +---------------+------+
-- (country_id, day) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table indicates the weather state in a country for one day.
 
-- Write a solution to find the type of weather in each country for November 2019.
-- The type of weather is:

-- Cold if the average weather_state is less than or equal 15,
-- Hot if the average weather_state is greater than or equal to 25, and
-- Warm otherwise.
-- Return the result table in any order.

-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Countries table:
-- +------------+--------------+
-- | country_id | country_name |
-- +------------+--------------+
-- | 2          | USA          |
-- | 3          | Australia    |
-- | 7          | Peru         |
-- | 5          | China        |
-- | 8          | Morocco      |
-- | 9          | Spain        |
-- +------------+--------------+
-- Weather table:
-- +------------+---------------+------------+
-- | country_id | weather_state | day        |
-- +------------+---------------+------------+
-- | 2          | 15            | 2019-11-01 |
-- | 2          | 12            | 2019-10-28 |
-- | 2          | 12            | 2019-10-27 |
-- | 3          | -2            | 2019-11-10 |
-- | 3          | 0             | 2019-11-11 |
-- | 3          | 3             | 2019-11-12 |
-- | 5          | 16            | 2019-11-07 |
-- | 5          | 18            | 2019-11-09 |
-- | 5          | 21            | 2019-11-23 |
-- | 7          | 25            | 2019-11-28 |
-- | 7          | 22            | 2019-12-01 |
-- | 7          | 20            | 2019-12-02 |
-- | 8          | 25            | 2019-11-05 |
-- | 8          | 27            | 2019-11-15 |
-- | 8          | 31            | 2019-11-25 |
-- | 9          | 7             | 2019-10-23 |
-- | 9          | 3             | 2019-12-23 |
-- +------------+---------------+------------+
-- Output: 
-- +--------------+--------------+
-- | country_name | weather_type |
-- +--------------+--------------+
-- | USA          | Cold         |
-- | Australia    | Cold         |
-- | Peru         | Hot          |
-- | Morocco      | Hot          |
-- | China        | Warm         |
-- +--------------+--------------+
-- Explanation: 
-- Average weather_state in USA in November is (15) / 1 = 15 so weather type is Cold.
-- Average weather_state in Austraila in November is (-2 + 0 + 3) / 3 = 0.333 so weather type is Cold.
-- Average weather_state in Peru in November is (25) / 1 = 25 so the weather type is Hot.
-- Average weather_state in China in November is (16 + 18 + 21) / 3 = 18.333 so weather type is Warm.
-- Average weather_state in Morocco in November is (25 + 27 + 31) / 3 = 27.667 so weather type is Hot.
-- We know nothing about the average weather_state in Spain in November so we do not include it in the result table. 

-- DATE_FORMAT(date,format);
-- %a	三个字符缩写的工作日名称，例如:Mon, Tue, Wed,等
-- %b	三个字符缩写的月份名称，例如：Jan，Feb，Mar等
-- %c	以数字表示的月份值，例如：1, 2, 3…12
-- %D	英文后缀如：0th, 1st, 2nd等的一个月之中的第几天
-- %d	如果是1个数字(小于10)，那么一个月之中的第几天表示为加前导加0， 如：00, 01,02, …31
-- %e	没有前导零的月份的日子，例如：1,2，… 31
-- %f	微秒，范围在000000..999999
-- %H	24小时格式的小时，前导加0，例如：00,01…23
-- %h	小时，12小时格式，带前导零，例如：01，02 … 12
-- %I	与%h相同
-- %i	分数为零，例如：00,01，… 59
-- %j	一年中的的第几天，前导为0，例如，001,002，… 366
-- %k	24小时格式的小时，无前导零，例如：0,1,2 … 23
-- %l	12小时格式的小时，无前导零，例如：0,1,2 … 12
-- %M	月份全名称，例如：January, February,…December
-- %m	具有前导零的月份名称，例如：00,01,02，… 12
-- %p	AM或PM，取决于其他时间说明符
-- %r	表示时间，12小时格式hh:mm:ss AM或PM
-- %S	表示秒，前导零，如：00,01，… 59
-- %s	与%S相同
-- %T	表示时间，24小时格式hh:mm:ss
-- %U	周的第一天是星期日，例如：00,01,02 … 53时，前导零的周数
-- %u	周的第一天是星期一，例如：00,01,02 … 53时，前导零的周数
-- %V	与%U相同，它与%X一起使用
-- %v	与%u相同，它与%x一起使用
-- %W	工作日的全称，例如：Sunday, Monday,…, Saturday
-- %w	工作日，以数字来表示(0 = 星期日，1 = 星期一等)
-- %X	周的四位数表示年份,第一天是星期日; 经常与%V一起使用
-- %x	周的四位数表示年份,第一天是星期日; 经常与%v一起使用
-- %Y	表示年份，四位数，例如2000，2001，…等。
-- %y	表示年份，两位数，例如00，01，…等。
-- %%	将百分比(%)字符添加到输出

-- Write your MySQL query statement below
SELECT
    c.country_name AS country_name,
    CASE 
        WHEN AVG(w.weather_state) >= 25 THEN 'Hot'  -- 当 weather_state 的平均值大于或等于 25 返回 Hot，
        WHEN AVG(w.weather_state) <= 15 THEN 'Cold' -- 当 weather_state 的平均值小于或等于 15 返回 Cold，
        ELSE 'Warm' -- 否则返回 Warm
    END AS weather_type
FROM 
    Countries AS c,
    Weather AS w
WHERE
    c.country_id = w.country_id AND
    day BETWEEN '2019-11-01' AND '2019-11-30' -- 找到表中每个国家在 2019 年 11 月的天气类型。
GROUP BY
    c.country_id 


-- best solution
SELECT 
    country_name,
    CASE WHEN AVG(weather_state) <= 15 THEN 'Cold'
        WHEN AVG(weather_state) >= 25 THEN 'Hot'
        ELSE 'Warm'
    END AS weather_type
FROM 
(
    SELECT 
        w.country_id, country_name, weather_state, day
    FROM 
        weather w
    LEFT JOIN countries c USING (country_id)
    WHERE 
        YEAR(day) = 2019 AND 
        MONTH(day) = 11
) AS t
GROUP BY country_id
