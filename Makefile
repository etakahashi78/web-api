up: ## dockerを起動します
	docker compose up -d
down: ## dockerを停止します
	docker compose down

dry-run: ## mysqldefのdry runを実行します
	mysqldef -uuser -ppassword web_api_db --dry-run < ./schema/schema.sql

help: ## ヘルプを表示します
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'