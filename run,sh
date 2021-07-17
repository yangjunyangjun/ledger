#!/bin/bash

#删掉之前编译的可执行文件
rm -r ./main

#拉取远程代码
git pull ledger main

#初始化swagger文件
swag init

#kill 之前的进程
lsof -i:8080 | awk '{print $2}' | xargs kill -9

#编译
go build main.go

#启动

nohup ./main &