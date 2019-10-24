.PHONY: build exec run

build:
	go build -o ./build/main main.go 

run:
	go build -o ./build/main main.go && ./build/main

run-timer:
	go build -o ./build/main main.go && ./build/main --limit 10
