API设计：用户
创建用户：URL:/user Method POST, StateCode:201 400 500
用户登录：URL:/user/:username Method POST, StateCode:200 400 500
获取用户信息：URL:/user/:username Method GET, StateCode:200 400 401 403 500
用户注销：URL:/user/:username Method DELETE, StateCode:204 400 401 403 500

API设计：用户资源
列举所有视频：URL:/user/:username/videos Method GET, StateCode:200 400 500
获取单条视频：URL:/user/:username/videos/:vid-id Method GET, StateCode:200 400 500
删除单条视频：URL:/user/:username/videos/:vid-id Method DELETE, StateCode:204 400 401 403 500

API设计：评论
展示评论：URL:/videos/:vid-id/comments Method GET, StateCode:200 400 500
提交一条评论：URL:/videos/:vid-id/comments/:comment-id Method POST, StateCode:201 400 500
删除一条评论：URL:/videos/:vid-id/comments/:comment-id Method DELETE, StateCode:204 400 401 403 500

数据库设计：用户表
TABLE： users
id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT
login_name VARCHAR(64), UNIQUE KEY
pwd TEXT

数据库设计：视频资源
TABLE： video_info
id VARCHAR(64), PRIMARY KEY, NOT NULL
author_id UNSIGNED INT      // 和users表中用户ID一致，不做成外键是因为不方便
name TEXT
display_ctime TEXT          // 用来展示的创建时间
create_time DATETIME

数据库设计：评论
TABLE： comments
id VARCHAR(64) PRIMARY KEY, NOT NULL
video_id VARCHAR(64)
author_id UNSIGNED INT
content TEXT
time DATETIME

数据库设计：SESSIONS  暂存当前状态，比如说账号登录之后一段时间内重新刷新页面不用重登录
TABLE： sessions
session_id TINYTEXT, PRIMARY KEY, NOT NULL
TTL TINYTEXT
login_name VARCHAR(64)

四个表的查询关系：
comments -> video_info -> users <- sessions

符合第三范式，各张表没有冗余信息，这样易于扩展