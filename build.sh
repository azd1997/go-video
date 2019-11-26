#!/bin/bash

# 创建bin文件夹
mkdir ./bin
mkdir ./bin/webui
mkdir ./bin/stream
mkdir ./bin/scheduler
mkdir ./bin/api

# 构建web UI
cd ./web
go build -o webui .
cd ../
cp ./web/webui ./bin/webui/webui
cp -R ./templates ./bin/webui/