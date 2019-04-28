local: stuff-tracker
	go run main.go

stuff-tracker:
	go build -a .

clean: clean-local clean-docker

clean-local:
	rm -f ./stuff-tracker

test:
	go test ./...

docker:
	docker build -t stuff-tracker .

local-docker:
	docker run --rm --name stuff-tracker stuff-tracker

clean-docker:
	docker rm -f stuff-tracker || true
	docker rmi stuff-tracker || true

.PHONY: local clean test docker
