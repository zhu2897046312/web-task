#!/bin/bash

# 设置 Go 环境变量
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

# 清理之前的构建
rm -rf build

# 创建构建目录
mkdir -p build

# 构建服务器
echo "Building server..."
go build -o build/server cmd/server/main.go

echo "Build completed!" 