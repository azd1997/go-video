API设计：用户
创建用户：URL:/user Method POST, StateCode:201 400 500
用户登录：URL:/user/:username Method POST, StateCode:200 400 500
获取用户信息：URL:/user/:username Method GET, StateCode:200 400 401 403 500
用户注销：URL:/user/:username Method DELETE, StateCode:204 400 401 403 500