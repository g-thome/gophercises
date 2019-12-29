.PHONY: build run clear

build:
	go build -o ./build/cyoa 

run:
	go build -o ./build/cyoa && ./build/cyoa --template ./frontend/markup.html --json ./adventure.json

clear:
	rm -r ./build/*
