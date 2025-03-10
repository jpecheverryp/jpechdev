## audit: run quality control checks
.PHONY: audit
audit: 
	go fmt ./...
	templ fmt .
	go mod tidy
	go mod verify
	go vet ./...

## dev: run application in developent mode
.PHONY: dev
dev:
	npm run dev &\
	templ generate --watch --cmd="go run ./cmd/web/"

prod:
	docker compose -f ./compose.yml -f ./prod-compose.yaml up -d
