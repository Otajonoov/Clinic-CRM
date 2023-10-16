CURRENT_DIR=$(shell pwd)

DB_URL="postgres://azizbek:CloudA2023*@localhost/medical_crm"

pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge
	
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run_script:
	./scripts/gen-proto.sh

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:12345@localhost:5432/lab_db force 1

migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table