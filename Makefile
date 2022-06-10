# Makefile for wordle.

SRCS = $(wildcard *.go)

all: wordle

wordle: $(SRCS)
	go build -o $@ $^
