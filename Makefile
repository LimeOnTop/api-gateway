ENTRY_POINT=cmd/app

# Генерация контрактов
proto-generate:
	protoc --proto_path=./proto \
		--go_out=./gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen \
		--go-grpc_opt=paths=source_relative \
		proto/auth/*.proto \
		proto/user/*.proto \
		proto/gpt/*.proto

# Запуск
run:
	go run $(ENTRY_POINT)/main.go

# Запуск в тестовом режиме
run-dev:
	go run $(ENTRY_POINT)/main.go --dev

