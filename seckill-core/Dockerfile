#原镜像
FROM golang:1.17
#新建文件夹
RUN mkdir /app
#复制代码
COPY . /app
#设置工作目录
COPY go.mod /app/sk-core
WORKDIR /app/sk-core

#设置代理
ENV GOPROXY https://goproxy.cn,direct

#go构建可执行文件
RUN go build .



#最终运行docker的命令
ENTRYPOINT  ["./sk-core"]
#docker网络配置，容器通讯。