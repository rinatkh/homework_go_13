GO ?= go
HTTP_ADDR ?= :8080
JWT_SECRET ?= homework-secret-change-me
TOKEN_TTL ?= 2h
BASE_URL ?= http://localhost:8080

.PHONY: run check-start test test-unit test-race fmt fmt-check vet deps ci curl-demo todo

run:
	HTTP_ADDR='$(HTTP_ADDR)' JWT_SECRET='$(JWT_SECRET)' TOKEN_TTL='$(TOKEN_TTL)' $(GO) run ./cmd/api

check-start:
	$(GO) test ./internal/starter
	$(GO) vet ./...
	$(GO) test ./... -run '^TestStarterCompiles$$'

test:
	$(GO) test ./...

test-unit:
	$(GO) test ./internal/user ./internal/note ./internal/auth ./internal/httpapi/...

test-race:
	$(GO) test -race ./...

fmt:
	gofmt -w $$(find . -type f -name '*.go')

fmt-check:
	@test -z "$$(gofmt -l $$(find . -type f -name '*.go'))" || (gofmt -l $$(find . -type f -name '*.go'); exit 1)

vet:
	$(GO) vet ./...

deps:
	$(GO) list -deps ./... >/dev/null

ci: fmt-check vet deps test test-race

curl-demo:
	BASE_URL='$(BASE_URL)' ./scripts/curl-demo.sh

todo:
	@rg -n 'TODO:' . --glob '*.go' --glob '*.md' || true
