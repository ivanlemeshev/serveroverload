.PHONY: tests
tests:
	go test -race -v -count=1 ./...

.PHONY: build-service
build-service:
	docker build --progress=plain --no-cache -t service:latest -f ./cmd/service/Dockerfile .

.PHONY: run-service
run-service: build-service
# https://docs.docker.com/config/containers/resource_constraints/
	docker run -d --rm --name service --cpus="1" --memory="300m" --memory-swap="300m" -p 8000:8000 service:latest

.PHONY: stop-service
stop-service:
	docker stop service

.PHONY: run-locust-tests
run-k6-tests:
	k6 run --out web-dashboard script.js
