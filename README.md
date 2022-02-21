### 编写自动部署脚本
- 复制bin/test.sh为master.ref.sh (请尽量以.ref.sh结尾)
```
case $1 in
    "refs/heads/master") # 可根据不同的分支做不同的任务
    echo $1
    ;;
esac
```
### 环境部署
- 安装go环境(以及go mod)
- 执行依耐安装(go mod tidy)

### 测试模块
```
// 模块运行
go test -v test/shell_test.go
// 方法运行
go test -v -run TestShell test/shell_test.go
```

### 编译webhooks程序
- 本机版
```
go build -o build/wh-mac main.go
```
- linux版
```
GOOS=linux GOARCH=amd64 go build -o build/wh-linux main.go
```

### 启动webhooks服务 (端口号默认为8080)
```
# 调试模式
./webhook server start
# daemon模式
./webhook server start -d=true
```

### 关闭webhooks服务
```
./webhook server stop
```

### 添加webhooks(请自行替换域名)
```
# post
https://domain/webhooks
# token 为bin目录下文件名${token}.ref.sh中${token}部分
```