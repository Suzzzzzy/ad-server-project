up:
	docker-compose build app
	docker-compose up

td-repository:
	go test ./src/repository/... -cover

td-usecase:
	go test ./src/usecase/... -cover
