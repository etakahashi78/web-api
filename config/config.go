package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbDsnPrimary string
	DbDsnReplica string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		slog.Error("error loading .env file", "err", err)
		return nil, err
	}

	// MySQL
	// プライマリ
	dbDSNPrimary := os.Getenv("DB_DSN_PRIMARY")
	if dbDSNPrimary == "" {
		return nil, fmt.Errorf("DB_DSN_PRIMARY environment variable not set")
	}
	// リードレプリカ
	dbDSNReplica := os.Getenv("DB_DSN_REPLICA")
	if dbDSNReplica == "" {
		slog.Info("DB_DSN_REPLICA environment variable not set, using primary for reads as well")
		dbDSNReplica = dbDSNPrimary // リードレプリカの設定がない場合はプライマリを使用する
	}

	return &Config{
		DbDsnPrimary: dbDSNPrimary,
		DbDsnReplica: dbDSNReplica,
	}, nil
}
