# Huango

## 功能列表

- [x] [air](https://github.com/cosmtrek/air) 热重载。
- [x] [gin](https://gin-gonic.com/) 路由框架。
- [x] 基于 [Viper](https://github.com/spf13/viper) + [Cast](https://github.com/spf13/cast) 实现配置方案。
- [x] 使用 [Gorm](https://gorm.io/) 对象关系映射。
- [x] 使用 [Govalidator](https://github.com/thedevsaddam/govalidator) 验证请求。
- [x] 使用 [Zap](https://github.com/uber-go/zap) 高性能日志库、集成 [lumberjack](https://github.com/natefinch/lumberjack) 滚动日志实现方案。
- [x] 使用 [go-redis/redis](https://github.com/go-redis/redis) 作为操作 `Redis` 的基础库。
- [x] 基于 [base64Captcha](https://github.com/mojocn/base64Captcha) 图片验证码库。
- [x] 统一响应 `reponse` 包。
- [x] 统一短信 `sms` 包，支持阿里云、腾讯云。
- [x] 发送短信验证码 `verifycode` 包。
- [x] 基于 [Email SMTP Driver](https://github.com/jordan-wright/email) 的发送 `email` 包。
- [x] 密码 `hash` 包。
- [x] 基于 [golang-jwt](https://github.com/golang-jwt/jwt) 的授权包 `jwt`。
- [x] 基于 [ulule/limiter](https://github.com/ulule/limiter) 中间件。
- [x] 基于 [Cobra](https://github.com/spf13/cobra) 命令行工具，以及基于 [Ansi](https://github.com/mgutz/ansi) 支持高亮输出的终端打印信息包 `console`。
- [x] 生成假数据 [Faker](https://github.com/bxcodec/faker)
- [x] 基于 [imaging](https://github.com/disintegration/imaging) 的图片裁剪。

## 本地 Docker 中使用方法

启动容器：

```
make up
```

进入容器：

```
make exec
```

停止容器:

```
make down
```

运行迁移文件：

```
make exec
go run main.go migrate up
```

填充假数据

```
make exec
go run main.go seed
```
