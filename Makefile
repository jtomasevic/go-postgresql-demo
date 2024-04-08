PHONY: generate run

generate:
	cd ./src/services/imdb/data_store; \
		sqlc generate;

run:
	go run .