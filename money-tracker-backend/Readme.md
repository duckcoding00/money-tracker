run:
docker compose up

set env:
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://root:mypassword@127.0.0.1:5432/money_tracker_db
export GOOSE_MIGRATION_DIR=./database/schema

generate repository:
sqlc generate

for run:
install air first
windows= air in cmd/powershell

linux= change .air.toml cmd from /bin/main.exe to /bin/main
