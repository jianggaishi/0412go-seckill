#原镜像
FROM golang:1.17
#新建文件夹
RUN mkdir /app
#复制代码
COPY . /app
#设置工作目录

WORKDIR /app/sk-admin

#设置代理
ENV GOPROXY https://goproxy.cn,direct

#go构建可执行文件
RUN cd ..
RUN go mod init
RUN go build .



#最终运行docker的命令
ENTRYPOINT  ["./sk-admin"]
#docker网络配置，容器通讯。



