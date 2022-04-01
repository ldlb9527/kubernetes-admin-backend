#!/bin/bash
tag=$1
###############################################  docker部署未验证  #####################################################
#go env -w GOPROXY=https://goproxy.cn,direct
#go mod tidy

SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build -o ./bin/linux_amd64/main main.go
docker build -t kubernetes-admin-backend:$tag .
docker rm -f kubernetes-admin-backend
docker run -d --restart always --net host --name=kubernetes-admin-backend -p 10010:10010 kubernetes-admin-backend:$tag
