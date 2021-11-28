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
+ 上传至Docker Hub
```
    docker login
    docker tag httpserver:v1 xxxxxxx:httpserverV1
    docker push
```
+ 利用nsenter获取容器ip
```
    docker inspect httpserver|grep Pid
    # 查看输出的Pid，假设我这里Pid是8461
    nsenter -n -t 8461 ip a
    # 输出一个lo设备和eth0的IP是172.17.0.4
```

## Deployment（课后作业第一部分）
+ 名字httpserver-deployment，配置了spec.template中lifecycle.postStart的优雅启动和lifecycle.preStop的优雅关闭；以及resource.limits和resource.requests，保证Pod是Burstable的；还有livenessProbe和readinessProbe。
+ 配置与代码分离：暂无

## Service or Ingress（课后作业第二部分）
暂无

由于课程进度还没有跟上，因此可能不是老师在课程中提到的知识，后续我会尽快将进度赶上，再修改作业。