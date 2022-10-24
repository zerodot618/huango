# 启动应用
up:
	docker compose up --build

# 停止应用
down:
	docker compose down

# 进入容器内、执行命令
exec:
	docker exec -it huango sh

# 本地运行应用
run:
	go run main.go

.PHONY: up down exec