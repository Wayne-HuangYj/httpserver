# 源镜像
FROM golang:1.15.10 as builder
# 工作目录
WORKDIR /root/src
# 将httpserver代码放入workdir
ADD src/main.go .
# 构建成可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o httpserver .

FROM alpine:latest
WORKDIR /bin
COPY --from=builder /root/src/httpserver .
EXPOSE 8080
CMD ["./httpserver"]