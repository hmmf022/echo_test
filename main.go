package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq" // PostgreSQLドライバ
)

func main() {
	// データベース接続
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo with PostgreSQL!")
	})

	// DB接続確認用のエンドポイント
	e.GET("/health", func(c echo.Context) error {
		// PingでDBの生存確認
		ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{"status": "db not ready"})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// サーバーの起動
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "1323" // デフォルトポート
	}
	e.Logger.Fatal(e.Start(":" + appPort))
}

// connectDB は環境変数から接続情報を読み取り、DBに接続します
func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// 接続文字列を作成
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// DBに接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// 接続確認 (Pingを5回試行)
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Successfully connected to the database!")
			return db, nil
		}
		log.Printf("Could not ping database, retrying in 2 seconds... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to the database after multiple retries")
}
