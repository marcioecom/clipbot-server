POSTGRESQL_URL=postgres://admin:123@localhost:5432/clipbot?sslmode=disable

generate-mock:
	mockery --name=IProducer --structname=ProducerMocked --inpackage=true --filename=producer_mock.go --outpkg=queue --dir=./infra/queue

generate-models:
	jet -dsn=postgresql://admin:123@localhost:5432/clipbot?sslmode=disable -schema=public -path=./.gen

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path infra/database/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path infra/database/migrations down

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down && docker volume rm clipbot-server_db

