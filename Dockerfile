FROM loads/alpine:3.8
LABEL maintainer="旅店老板"
# 设置固定的项目路径
ENV WORKDIR /var/www/kubernetes-admin-backend

# 添加应用可执行文件，并设置执行权限
ADD ./bin/linux_amd64/main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

ADD config   $WORKDIR/config
#ADD public   $WORKDIR/public
#ADD template $WORKDIR/template

WORKDIR $WORKDIR
CMD ./main

#

#gf docker -t qc-image:v2 或  docker build -t qc-image:v2 .
 #SET CGO_ENABLED=0
 #SET GOOS=linux
 #SET GOARCH=amd64
 #go build main.go

