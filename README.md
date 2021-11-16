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

### 编译webhooks程序
```
go build -o webhooks main.go
```

### 启动webhooks程序 (端口号默认为8080)
```
./webhooks -p 8080 
```
