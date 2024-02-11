FROM golang:1.22 AS build-stage
RUN mkdir /code
COPY . /code
WORKDIR /code
RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./cmd/service/main.go

# hadolint ignore=DL3007
FROM gcr.io/distroless/base:latest AS build-release-stage
WORKDIR /
COPY --from=build-stage /service /service
EXPOSE 8080
ENTRYPOINT ["/service"]
