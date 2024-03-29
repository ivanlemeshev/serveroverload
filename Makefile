.PHONY: tests
tests:
	go test -race -v -count=1 ./...

.PHONY: build-service
build-service:
	docker build --progress=plain --no-cache -t service:latest -f ./cmd/service/Dockerfile .

.PHONY: run-service
run-service: build-service
# https://docs.docker.com/config/containers/resource_constraints/
	docker run -d --rm --name service --cpus="0.2" --memory="200m" --memory-swap="200m" -p 8000:8000 service:latest

.PHONY: stop-service
stop-service:
	docker stop service
