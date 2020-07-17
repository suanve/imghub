FROM golang
WORKDIR /app
ADD . /app
#go构建可执行文件
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go build -o imgHub
#暴露端口
EXPOSE 8081
#最终运行docker的命令
ENTRYPOINT  ["./imgHub"]