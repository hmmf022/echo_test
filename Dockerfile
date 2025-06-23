# 1. ビルドステージ
# 公式のGoイメージをベースにする
FROM golang:1.23-alpine AS builder

# 作業ディレクトリを作成
WORKDIR /app

# ホットリロードツール `air` をインストール
RUN go install github.com/air-verse/air@v1.61.7

# Go Modulesのキャッシュを有効にするため、先に依存関係ファイルだけコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# 開発用のポートを開放
EXPOSE 1323

# air を使ってアプリケーションを起動
CMD ["air"]
