-- 2664 · Recognizing SQL Views (I)
-- # Description
-- In the backend, we provide a view v_teachers of the teachers table, and now we need you to view the contents of this view.

-- Example
-- 后台视图的创建：

-- CREATE VIEW v_teachers AS SELECT * FROM `teachers`;
-- 返回结果：

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'

CREATE VIEW 
    `v_teachers` 
AS 
    SELECT 
        * 
    FROM 
        `teachers`;

-- Q：什么是视图？视图是干什么用的？
-- A：
-- 　　视图（view）是一种虚拟存在的表，是一个逻辑表，本身并不包含数据。作为一个select语句保存在数据字典中的。
-- 　　通过视图，可以展现基表的部分数据；视图数据来自定义视图的查询中使用的表，使用视图动态生成。
--     基表：用来创建视图的表叫做基表base table

-- Q：为什么要使用视图？
-- A：因为视图的诸多优点，如下
-- 　　1）简单：使用视图的用户完全不需要关心后面对应的表的结构、关联条件和筛选条件，对用户来说已经是过滤好的复合条件的结果集。
-- 　　2）安全：使用视图的用户只能访问他们被允许查询的结果集，对表的权限管理并不能限制到某个行某个列，但是通过视图就可以简单的实现。
-- 　　3）数据独立：一旦视图的结构确定了，可以屏蔽表结构变化对用户的影响，源表增加列对视图没有影响；源表修改列名，则可以通过修改视图来解决，不会造成对访问者的影响。
--    总而言之，使用视图的大部分情况是为了保障数据安全性，提高查询效率。

-- # 创建视图
-- CREATE [OR REPLACE] [ALGORITHM = {UNDEFINED | MERGE | TEMPTABLE}]
--     VIEW view_name [(column_list)]
--     AS select_statement
--    [WITH [CASCADED | LOCAL] CHECK OPTION]
--
-- 1）OR REPLACE：表示替换已有视图
-- 2）ALGORITHM：表示视图选择算法，默认算法是 UNDEFINED(未定义的)： MySQL自动选择要使用的算法 ；merge合并；temptable临时表
--              MERGE：将视图的语句与视图定义合并起来，使得视图定义的某一部分取代语句的对应部分
--              TEMPTABLE：将视图的结果存入临时表，然后使用临时表执行语句
-- 3）select_statement：表示select语句
-- 4）[WITH [CASCADED | LOCAL] CHECK OPTION]：表示视图在更新时保证在视图的权限范围之内
-- 　　cascade 是默认值，表示更新视图的时候，要满足视图和表的相关条件
-- 　　local表示更新视图的时候，要满足该视图定义的一个条件即可
--     推荐使用WHIT [CASCADED|LOCAL] CHECK OPTION选项，可以保证数据的安全性 
--
-- 基本格式：
-- 　　create view <视图名称>[(column_list)]
--        as select语句
--        with check option;

-- # 查看视图
-- ## 1 使用show create view语句查看视图信息 [ show create view ]
-- mysql> show create view v_F_players\G;
-- *************************** 1. row ***************************
--                 View: v_F_players
--          Create View: CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_F_players` AS select `PLAYERS`.`PLAYERNO` AS `编号`,`PLAYERS`.`NAME` AS `名字`,`PLAYERS`.`SEX` AS `性别`,`PLAYERS`.`PHONENO` AS `电话` from `PLAYERS` where (`PLAYERS`.`SEX` = 'F') WITH CASCADED CHECK OPTION
-- character_set_client: utf8
-- collation_connection: utf8_general_ci
-- 1 row in set (0.00 sec)
---
-- ## 2 视图一旦创建完毕，就可以像一个普通表那样使用，视图主要用来查询
-- mysql> select * from view_name;
--
-- ## 3 有关视图的信息记录在information_schema数据库中的 views 表中 [ select information_schema.views ]
-- mysql> select * from information_schema.views 
--     -> where TABLE_NAME='v_F_players'\G;
-- *************************** 1. row ***************************
--        TABLE_CATALOG: def
--         TABLE_SCHEMA: TENNIS
--           TABLE_NAME: v_F_players
--      VIEW_DEFINITION: select `TENNIS`.`PLAYERS`.`PLAYERNO` AS `编号`,`TENNIS`.`PLAYERS`.`NAME` AS `名字`,`TENNIS`.`PLAYERS`.`SEX` AS `性别`,`TENNIS`.`PLAYERS`.`PHONENO` AS `电话` from `TENNIS`.`PLAYERS` where (`TENNIS`.`PLAYERS`.`SEX` = 'F')
--         CHECK_OPTION: CASCADED
--         IS_UPDATABLE: YES
--              DEFINER: root@localhost
--        SECURITY_TYPE: DEFINER
-- CHARACTER_SET_CLIENT: utf8
-- COLLATION_CONNECTION: utf8_general_ci
-- 1 row in set (0.00 sec)      

-- # 视图的更改
-- ## 1 CREATE OR REPLACE VIEW语句修改视图
-- 基本格式：
-- 　　create or replace view view_name as select语句;
--     在视图存在的情况下可对视图进行修改，视图不在的情况下可创建视图
-- 
-- ## 2 ALTER语句修改视图
-- ALTER
--     [ALGORITHM = {UNDEFINED | MERGE | TEMPTABLE}]
--     [DEFINER = { user | CURRENT_USER }]
--     [SQL SECURITY { DEFINER | INVOKER }]
-- VIEW view_name [(column_list)]
-- AS select_statement
--     [WITH [CASCADED | LOCAL] CHECK OPTION]
-- 注意：修改视图是指修改数据库中已存在的表的定义，当基表的某些字段发生改变时，可以通过修改视图来保持视图和基本表之间一致

-- # Drop删除视图
-- 删除视图是指删除数据库中已存在的视图，删除视图时，只能删除视图的定义，不会删除数据，也就是说不动基表：
--      DROP VIEW [IF EXISTS]  view_name [, view_name] ...
