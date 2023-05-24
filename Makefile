migrate:
	liquibase update --url="jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable" --changelog-file="migration/liquibase/changelog.xml"

rollback:
	liquibase rollback --tag=$(tag) --url="jdbc:postgresql://localhost:5432/logity_auth?user=postgres&password=postgres&sslmode=disable" --changelog-file="migration/liquibase/changelog.xml"