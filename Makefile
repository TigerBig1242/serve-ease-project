run:
	DNV=dev go run ./cmd/app/main.go

.PHONY: dev-up dev-down dev-logs
dev-up:
	docker compose -f docker-compose.dev.yml up -d
dev-down:
	docker compose -f docker-compose.dev.yml down
dev-logs:
	docker compose -f docker-compose.dev.yml logs -f