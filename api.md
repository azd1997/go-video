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




