DATABASE=jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable


migrate:
	liquibase update --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"

rollback-count:
	liquibase rollback-count --count=$(count) --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"


rollback-tag:
	liquibase rollback --tag=$(tag) --url="$(DATABASE)" --changelog-file="migration/liquibase/changelog.xml"