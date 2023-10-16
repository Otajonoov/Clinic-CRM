CURRENT_DIR=$(shell pwd)

-include .env

DB_URL="postgres://azizbek:CloudA2023*@localhost:5432/doctor_db"

run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table

local-up:
	docker compose --env-file ./.env.docker up -d

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

pull-sub-module:
	git submodule update --init --recursive

update-sub-module:
	git submodule update --remote --merge


.PHONY: run migrateup migratedown local-up proto-gen pull-sub-module update-sub-module