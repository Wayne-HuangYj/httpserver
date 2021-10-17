# 源镜像
FROM golang:1.15.10 as builder
# 环境变量
ENV SRC_PATH $GOPATH/src/httpserver
ENV BINARY httpserver
# 工作目录
WORKDIR $SRC_PATH
# 将httpserver代码放入workdir
COPY ../src/main.go .
# 构建成可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o $BINARY .

FROM alpine:latest
WORKDIR /bin
COPY --from=builder $SRC_PATH/$BINARY .
EXPOSE 8080
CMD ["./$BINARY"]