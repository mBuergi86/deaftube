.PHONY: Connect to Server with run and stop

connection:
	echo "Connect to Server"
build:
	go build -o bin/main server.go

run:
	go run server.go

all: connection build run

stop:
	@pkill -f "server"