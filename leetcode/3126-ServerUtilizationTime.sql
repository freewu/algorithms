-- 3126. Server Utilization Time
-- Table: Servers
-- +----------------+----------+
-- | Column Name    | Type     |
-- +----------------+----------+
-- | server_id      | int      |
-- | status_time    | datetime |
-- | session_status | enum     |
-- +----------------+----------+
-- (server_id, status_time, session_status) is the primary key (combination of columns with unique values) for this table.
-- session_status is an ENUM (category) type of ('start', 'stop').
-- Each row of this table contains server_id, status_time, and session_status.
-- Write a solution to find the total time when servers were running. The output should be in units of full days.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example:
-- Input:
-- Servers table:
-- +-----------+---------------------+----------------+
-- | server_id | status_time         | session_status |
-- +-----------+---------------------+----------------+
-- | 3         | 2023-11-04 16:29:47 | start          |
-- | 3         | 2023-11-05 01:49:47 | stop           |
-- | 3         | 2023-11-25 01:37:08 | start          |
-- | 3         | 2023-11-25 03:50:08 | stop           |
-- | 1         | 2023-11-13 03:05:31 | start          |
-- | 1         | 2023-11-13 11:10:31 | stop           |
-- | 4         | 2023-11-29 15:11:17 | start          |
-- | 4         | 2023-11-29 15:42:17 | stop           |
-- | 4         | 2023-11-20 00:31:44 | start          |
-- | 4         | 2023-11-20 07:03:44 | stop           |
-- | 1         | 2023-11-20 00:27:11 | start          |
-- | 1         | 2023-11-20 01:41:11 | stop           |
-- | 3         | 2023-11-04 23:16:48 | start          |
-- | 3         | 2023-11-05 01:15:48 | stop           |
-- | 4         | 2023-11-30 15:09:18 | start          |
-- | 4         | 2023-11-30 20:48:18 | stop           |
-- | 4         | 2023-11-25 21:09:06 | start          |
-- | 4         | 2023-11-26 04:58:06 | stop           |
-- | 5         | 2023-11-16 19:42:22 | start          |
-- | 5         | 2023-11-16 21:08:22 | stop           |
-- +-----------+---------------------+----------------+
-- Output:
-- +-------------------+
-- | total_uptime_days |
-- +-------------------+
-- | 1                 |
-- +-------------------+
-- Explanation:
-- For server ID 3:
-- From 2023-11-04 16:29:47 to 2023-11-05 01:49:47: ~9.3 hours
-- From 2023-11-25 01:37:08 to 2023-11-25 03:50:08: ~2.2 hours
-- From 2023-11-04 23:16:48 to 2023-11-05 01:15:48: ~1.98 hours
-- Total for server 3: ~13.48 hours
-- For server ID 1:
-- From 2023-11-13 03:05:31 to 2023-11-13 11:10:31: ~8 hours
-- From 2023-11-20 00:27:11 to 2023-11-20 01:41:11: ~1.23 hours
-- Total for server 1: ~9.23 hours
-- For server ID 4:
-- From 2023-11-29 15:11:17 to 2023-11-29 15:42:17: ~0.52 hours
-- From 2023-11-20 00:31:44 to 2023-11-20 07:03:44: ~6.53 hours
-- From 2023-11-30 15:09:18 to 2023-11-30 20:48:18: ~5.65 hours
-- From 2023-11-25 21:09:06 to 2023-11-26 04:58:06: ~7.82 hours
-- Total for server 4: ~20.52 hours
-- For server ID 5:
-- From 2023-11-16 19:42:22 to 2023-11-16 21:08:22: ~1.43 hours
-- Total for server 5: ~1.43 hours
-- The accumulated runtime for all servers totals approximately 44.46 hours, equivalent to one full day plus some additional hours. However, since we consider only full days, the final output is rounded to 1 full day.

-- Create table if not exists Servers ( server_id int, status_time timestamp, session_status ENUM ('start','stop'))
-- Truncate table Servers
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-29 20:22:50', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-30 04:19:50', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-18 21:23:09', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-19 00:38:09', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-27 18:38:16', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-28 02:46:16', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-27 06:46:42', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-27 08:23:42', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-11 05:09:45', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-11 06:42:45', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-27 23:05:58', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-27 23:48:58', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-14 16:29:34', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-14 19:11:34', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 19:14:15', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 20:19:15', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-03 06:55:30', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-03 11:03:30', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-02 16:53:28', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-02 22:28:28', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-20 07:06:41', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-20 15:58:41', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-22 10:42:35', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-22 19:14:35', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-23 03:05:39', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-23 10:54:39', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-15 20:04:15', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-15 21:30:15', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-12 08:21:42', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-12 17:07:42', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 02:33:24', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 12:10:24', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-14 03:55:29', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-14 08:53:29', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-04 00:36:53', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-04 00:37:53', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-08 01:54:18', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-08 09:56:18', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-24 14:28:11', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-24 21:38:11', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-20 22:55:03', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-21 02:19:03', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-12 03:10:22', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-12 08:29:22', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-02 14:55:46', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-02 16:56:46', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-09 03:40:44', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-09 03:52:44', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-16 17:43:28', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-16 20:45:28', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-08 16:17:24', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-09 00:59:24', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-25 04:46:06', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-25 11:56:06', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-05 14:22:34', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-05 21:33:34', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 09:49:22', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-02 15:01:22', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-08 06:49:10', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-08 13:33:10', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-22 21:00:23', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-22 22:43:23', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-06 07:45:50', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-06 15:18:50', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-02 22:13:17', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-03 03:50:17', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-14 02:40:20', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-14 03:57:20', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-08 21:13:00', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-09 01:22:00', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-21 13:25:22', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-21 19:46:22', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-16 04:37:55', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-16 12:53:55', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-06 14:03:26', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-06 17:11:26', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-11 08:18:54', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-11 12:34:54', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-07 22:54:01', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-08 08:22:01', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-23 14:08:03', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-23 22:23:03', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-04 15:20:41', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-04 16:09:41', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-05 01:24:32', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-05 10:09:32', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-09 14:15:29', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-09 18:39:29', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-28 21:43:00', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-29 06:06:00', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-18 06:23:14', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-18 07:34:14', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-09 17:27:02', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('5', '2023-11-10 03:00:02', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-21 02:41:57', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-21 07:28:57', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-08 00:36:12', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-08 01:47:12', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-15 16:13:39', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-15 20:29:39', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-03 06:25:01', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-03 11:14:01', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-18 20:49:55', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-18 23:04:55', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-13 05:02:21', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('3', '2023-11-13 12:44:21', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-23 19:00:11', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-24 04:59:11', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-20 01:16:38', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-20 01:35:38', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-13 11:21:08', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('1', '2023-11-13 11:56:08', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-21 23:09:45', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('4', '2023-11-22 00:26:45', 'stop')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-04 11:19:31', 'start')
-- insert into Servers (server_id, status_time, session_status) values ('2', '2023-11-04 16:46:31', 'stop')

-- with rank
WITH r AS (-- 以每台服务器按时间状态排序 status_time  00:00:00 -> 23:59:59
    SELECT 
        *,
        RANK() OVER (PARTITION BY server_id ORDER BY status_time, session_status) AS rk 
    FROM 
        Servers
)

SELECT 
    SUM(duration) / 86400 AS total_uptime_days -- 运行天数
FROM 
(
    SELECT 
        TIMESTAMPDIFF(SECOND, a.status_time, b.status_time) AS duration -- 计算毛得到每次开始结果的时间
    FROM 
        r AS a, 
        r AS b 
    WHERE 
        a.server_id = b.server_id AND a.session_status = 'start' AND a.r = b.r - 1
) AS t

-- with lead
SELECT 
    FLOOR(SUM(TIMESTAMPDIFF(SECOND , status_time, next_status_time)) / 86400) AS total_uptime_days 
FROM (
    SELECT
        session_status,
        status_time,
        LEAD(status_time) OVER (PARTITION BY server_id ORDER BY status_time) AS next_status_time
    FROM
        Servers
) AS t
WHERE 
    session_status = 'start';
