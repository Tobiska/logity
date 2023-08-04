DATABASE=jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable
APP_NAME=logity
DOCKER_REPOSITORY=tobiskadocker

upgrade:
	helm upgrade $(APP_NAME) ./chart --install --atomic --timeout 3m

local-environment-up:
	docker-compose -f docker-compose-env.yml up --build -d

docker-build:
	docker build -f build/app/Dockerfile -t $(APP_NAME):latest .
	docker tag $(APP_NAME):latest $(DOCKER_REPOSITORY)/$(APP_NAME):latest
	docker push $(DOCKER_REPOSITORY)/$(APP_NAME):latest

lint: ## Run linters
	golangci-lint run --disable-all -E govet,staticcheck

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