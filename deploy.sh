#!/bin/bash

APP_NAME="app"
SRC_DIR="/opt/goapp/src"

# 最新コード取得
cd "$SRC_DIR"
git fetch --prune
git pull origin main

# ビルド
go build -o "$APP_NAME" .

# サービス再起動
sudo systemctl restart myapp.service

echo "Deployment completed!"
