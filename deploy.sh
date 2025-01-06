#!/bin/bash

APP_NAME="app"
BUILD_DIR="/opt/goapp/src"
LOG_DIR="/opt/goapp/logs"
SRC_DIR="/opt/goapp/src"
MAX_BACKUPS=1  # 残すバイナリの最大数
MAX_LOGS=10  # 残すログファイルの最大数

# バイナリ名を一意化 (例: 日時付き)
NEW_BINARY="$BUILD_DIR/$APP_NAME-$(date +%Y%m%d%H%M%S)"

# 最新コード取得
cd "$SRC_DIR"
git fetch --prune
git pull origin main

# ビルド
go build -o "$NEW_BINARY" .

# シンボリックリンク更新
ln -sf "$NEW_BINARY" "$BUILD_DIR/$APP_NAME"

# 古いバイナリを削除 (最新の N 個以外)
ls -1t "$BUILD_DIR/$APP_NAME-"* | tail -n +$((MAX_BACKUPS + 1)) | xargs -r rm -f

# 古いログファイルを削除
ls -1t "$LOG_DIR/app.log.*" | tail -n +$((MAX_LOGS + 1)) | xargs -r rm -f

# サービス再起動
sudo systemctl restart myapp.service

echo "Deployment completed!"
