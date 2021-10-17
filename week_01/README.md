# 任务1

    1）接收客户端 request，并将 request 中带的 header 写入 response header

    2）读取当前系统的环境变量中的 VERSION 配置，并写入 response header

    3） Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

    4）当访问 localhost/healthz 时，应返回 200

# 任务2
    构建本地镜像。
        运行 make build
    编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
        查看Dcokerfile
    将镜像推送至 Docker 官方镜像仓库。
        make push
    通过 Docker 命令本地启动 httpserver。
        make run 
    通过 nsenter 进入容器查看 IP 配置。
        make nsenter
        ifconfig

    注意: apple m1安装 nsenter 需要二进制编译