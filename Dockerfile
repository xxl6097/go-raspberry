# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER xiaxiaoli
# 全局工作目录
WORKDIR /app
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /app
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
#RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go
RUN go build -ldflags="-w -s" *.go



FROM alpine AS runner

#ENV HOST_API CLINK_API
WORKDIR /app
RUN mkdir ./conf
COPY --from=builder /app/main .
COPY --from=builder /app/app.yaml ./conf/
RUN ls -l ./conf/
#COPY --from=builder /app/conf ./conf
RUN #echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
#    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
#    && apk add --no-cache tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
#    && echo Asia/Shanghai > /etc/timezone \
#    && apk del tzdata
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置时区（以 Asia/Shanghai 为例）
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
#VOLUME ["/app/conf/app.yaml"]

COPY entrypoint.sh /entrypoint.sh

RUN apk add --no-cache bash ca-certificates su-exec tzdata; \
    chmod +x /entrypoint.sh
ENV PUID=0 PGID=0 UMASK=022

# 需暴露的端口
#EXPOSE 9090
#ENTRYPOINT ["./main","-c","./conf/app.yaml"]
CMD ["/entrypoint.sh"]
