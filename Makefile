DATABASE=jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable

lint: ## Run linters
	golangci-lint run --disable-all -E govet,staticcheck,errcheck

test:
	go test -race -v -timeout=30s -coverprofile=cover.out -coverpkg=./... ./...

migrate:
	liquibase update --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"

rollback-count:
	liquibase rollback-count --count=$(count) --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"

swag-gen:
	swag init -g internal/delivery/rest/router.go

rollback-tag:
	liquibase rollback --tag=$(tag) --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"