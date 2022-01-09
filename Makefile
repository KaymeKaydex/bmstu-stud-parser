PWD = $(shell pwd)
NAME = bmstu-stud-parser

.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

.PHONY: build
build:
	go build -o bin/$(NAME) $(PWD)/cmd/$(NAME)