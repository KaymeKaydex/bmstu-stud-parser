PWD = $(shell pwd)
NAME = bmstu_stud_parser

.PHONY: run
run:
	go run $(PWD)/cmd/bmstu_stud_parser/

.PHONY: build
build:
	go build -o bin/$(NAME) $(PWD)/cmd/$(NAME)