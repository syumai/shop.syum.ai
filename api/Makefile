sqlite3def := go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.17.6

.PHONY: migrate
migrate:
	$(sqlite3def) -f schema.sql db.sqlite

.PHONY: migrate-dry
migrate-dry:
	$(sqlite3def) -f schema.sql --dry-run db.sqlite
