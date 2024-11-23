#!/bin/bash

# 设置变量
APP_NAME="web-task"
SERVER_DIR="/opt/web-task"
BUILD_DIR="build"

# 确保目标目录存在
sudo mkdir -p ${SERVER_DIR}

# 停止现有服务
sudo systemctl stop ${APP_NAME} || true

# 复制新的二进制文件
sudo cp ${BUILD_DIR}/server ${SERVER_DIR}/
sudo cp configs/config.yaml ${SERVER_DIR}/

# 设置权限
sudo chmod +x ${SERVER_DIR}/server

# 重启服务
sudo systemctl start ${APP_NAME}

echo "Deployment completed!" 