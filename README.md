# blog

使用gin构建博客系统，代码简洁易读，可快速进行高效web开发。

### 文件分层
```
.
├── api
│   └── v1
├── config
│   └── dev
├── docs
└── pkg
    ├── base
    ├── ctrl
    ├── dto
    ├── middleware
    ├── model
    ├── router
    └── util
```

### 第三方包
- [logrus](https://github.com/sirupsen/logrus) 日志
- [swagger](https://github.com/swaggo/gin-swagger) 文档
- [viper](https://github.com/spf13/viper) 加载配置
- [gorm](https://github.com/jinzhu/gorm) 数据库orm