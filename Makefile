migrate:
	liquibase update --url="jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable" --changelog-file="migration/liquibase/changelog.xml"

rollback-count:
	liquibase rollback-count --count=$(count) --url="jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable" --changelog-file="migration/liquibase/changelog.xml"


rollback-tag:
	liquibase rollback --tag=$(tag) --url="jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable" --changelog-file="migration/liquibase/changelog.xml"