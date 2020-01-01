FROM golang:1.13-alpine

#设置环境变量
#GOPROXY国内网络便于下载包
#GO111MODULE打开go mod
ENV GOPROXY https://mirrors.aliyun.com/goproxy
ENV GO111MODULE on

#切换工作目录，即在容器内部的目录
WORKDIR /go-gin-demo

#将当前主机文件复制到容器文件系统 这里.是容器文件系统的当前工作目录
ADD go.mod .
ADD go.sum .
#下载go mod依赖
RUN go mod download

#文件移动
ADD . .

#目录内容构成见pic/contain_dir.png

# GOOS GOARCH 等为golang 跨平台编译参数
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o gin_demo_server go-gin-demo/main

EXPOSE 9090

#启动容器的执行命令
CMD ["./gin_demo_server"]
