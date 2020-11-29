.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=1234 hospital
adminer:
	docker rm --rm -ti --network host adminer
migrate:
	migrate -source file://migration \
			-database postgres://root:1234@localhost/hospital?sslmode=disable up