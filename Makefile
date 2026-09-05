.PHONY: build fmt vet test validate dev

# Builds the three Core binaries (pcf, relay, audit).
build:
	go build ./...

# Fails if any file isn't gofmt-formatted.
fmt:
	@test -z "$$(gofmt -l .)" || (echo "Unformatted files:"; gofmt -l .; exit 1)

vet:
	go vet ./...

test:
	go test ./...

# This is the target CI runs. Human, agent, and CI all run the same thing.
validate: fmt vet build test

# Brings up Postgres + Kafka + Redis + the three binaries (Phase C1 onward).
dev:
	docker compose -f deploy/docker-compose.yaml up
