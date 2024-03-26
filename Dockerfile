# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER chen
ENV TZ=Asia/Shanghai
# 全局工作目录
WORKDIR /home/guizizhan_v6
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/guizizhan_v6
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app main.go

# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
ENV TZ=Asia/Shanghai
# 全局工作目录
WORKDIR /home/guizizhan_v6
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/guizizhan_v6/app .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/guizizhan_v6/config ./config
# 需暴露的端口
EXPOSE 8080
# 可外挂的目录
VOLUME ["/home/guizizhan_v6/config"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./app"]
