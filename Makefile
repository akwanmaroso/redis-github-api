run:
	go run .
build:
	go build . && ./main
redis-create:
	docker run --name redis-go -p 6379:6379 -d redis
redis-delete:
	docker stop redis-go && docker container rm redis-go
redis-client:
	docker container exec -it redis-go /bin/sh