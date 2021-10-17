# 简单的HTTP SERVER
+ 监听8080端口，访问/设置response的Header，并且获取环境变量VERSION写入Header中  
+ 访问/healthz的statusCode是200

## Dockerfile
+ 双重构建，先用golang镜像执行go build生产可执行文件，然后将可执行文件拷贝到基础的Linux镜像中运行，保证Docker镜像体积小
+ 在项目目录下执行，便可以访问容器中httpserver提供的服务
```
    docker build -t httpserver:v1 -f docker/Dockerfile .
    docker run --name httpserver -d -P httpserver:v1
```