
all: publish

compile: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o canary .

build: compile
	docker build -t massiveco/canary:latest .

publish: build
	docker push massiveco/canary:latest

PHONY: compile all build publish