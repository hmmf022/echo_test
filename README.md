# Go (Echo) + PostgreSQL 開発環境テンプレート

Go言語のWebフレームワーク [Echo](https://echo.labstack.com/) と PostgreSQL を使用した、Dockerベースの開発環境テンプレートです。
ローカルにGoやPostgreSQLをインストールすることなく、すぐに開発を始めることができます。

ホットリロードに対応しており、コードを変更すると自動でアプリケーションが再起動するため、快適な開発体験を提供します。

## ✨ 特徴

-   **Go + Echo**: 高速でミニマルなWebフレームワーク
-   **PostgreSQL**: 高機能なリレーショナルデータベース
-   **Docker Ready**: `docker compose up` だけで開発環境が起動
-   **ホットリロード**: [Air](https://github.com/air-verse/air) によるライブリローディングで開発効率UP
-   **クリーンな環境**: 開発に必要なツールはすべてコンテナ内に集約

## 🛠️ 技術スタック

| カテゴリ       | 技術                                      | バージョン |
| :------------- | :---------------------------------------- | :--------- |
| **言語**       | Go                                        | `1.23`     |
| **フレームワーク** | Echo                                      | `v4`       |
| **データベース**   | PostgreSQL                                | `16`       |
| **コンテナ**     | Docker / Docker Compose                   | -          |
| **開発ツール**   | Air (ホットリロード)                      | `v1.61.7`  |

## 🚀 セットアップと実行方法

### 1. 前提条件

-   [Docker](https://www.docker.com/products/docker-desktop/)
-   [Docker Compose](https://docs.docker.com/compose/install/) (Docker Desktop には同梱されています)

### 2. 環境構築

1.  **このリポジトリをクローンする**
    ```bash
    git clone https://github.com/hmmf022/echo_test.git
    cd echo_test
    ```

2.  **環境変数ファイルを作成する**
    `.env.example` ファイルをコピーして `.env` ファイルを作成します。このファイルにデータベースの接続情報などを記述します。
    ```bash
    cp .env.example .env
    ```
    > **Note:** `.env` ファイルは機密情報を含むため、`.gitignore` に追加してバージョン管理の対象外にしています。

3.  **Dockerコンテナを起動する**
    以下のコマンドを実行して、コンテナをビルドし、起動します。
    ```bash
    docker compose up --build
    ```
    初回以降は `docker compose up` だけで起動できますが、`Dockerfile` を変更した際は `--build` オプションを付けてください。

### 3. 動作確認

コンテナが正常に起動したら、ブラウザや`curl`で以下のURLにアクセスします。

-   **ルート**: `http://localhost:1323`
    -   "Hello, Echo with PostgreSQL!" と表示されます。
-   **DB接続確認**: `http://localhost:1323/health`
    -   `{"status":"ok"}` と表示されれば、データベースとの接続も成功しています。

---

### `.env.example` の内容

このプロジェクトで使用する環境変数のテンプレートです。必要に応じて値を変更してください。

```env
# .env.example

# PostgreSQLの接続設定
DB_HOST=pg-16 # compose.ymlで定義したサービス名
DB_PORT=5432
POSTGRES_DB=mydb
POSTGRES_USER=webmaster
POSTGRES_PASSWORD=1234qweR

# アプリケーションのポート
APP_PORT=1323
```

## 便利なコマンド集

```bash
# バックグラウンドで起動
docker compose up -d

# 停止
docker compose down

# ログをリアルタイムで表示 (appサービスの場合)
docker compose logs -f app

# Goアプリケーションコンテナのシェルに入る
docker compose exec app sh
```

## 🚀 本番環境へのデプロイについて

このリポジトリは **開発環境** に特化しています。本番環境では以下の点を考慮する必要があります。

-   **Nginxの導入**: 静的ファイルの配信やリバースプロキシとして、Goアプリケーションの前段にNginxを配置するのが一般的です。
-   **ホットリロードの無効化**: 本番環境では`air`は不要です。コンパイル済みの軽量なバイナリを直接実行します。
-   **マルチステージビルド**: `Dockerfile`を本番用に最適化し、ビルド環境と実行環境を分離することで、最終的なコンテナイメージのサイズを大幅に削減できます。
