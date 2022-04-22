build:
	go build -o dist/myip cmd/main.go
build-docker:
	docker build -t rlaskowski/myip .
run:
	go run cmd/main.go -f config.yaml
clean: 
	rm -rf dist