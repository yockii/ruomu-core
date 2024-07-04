# ruomu-core
核心功能，用于快速复用

# 核心部分包含
- [x] 基础功能库
- [ ] 公共调用rpc接口

# 数据库相关包
需要自行引入，不包含在核心包中
```go
package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)
```

# 接口功能
- [ ] HTTP请求时调用模块
- [ ] HOOK点暴露调用模块


# 模块注册
通过核心系统进行注册
- [ ] 手动注册，核心系统手动填写并上传插件包
- [ ] 自动注册，应用市场一键安装