GO = $(shell which go 2>/dev/null)
COMPOSE = $(shell which docker-compose 2>/dev/null)
SRC_DIRS := cmd

ifeq ($(GO),)
$(error "go is not in your system PATH")
else
$(info "go found")
endif

ifeq ($(COMPOSE),)
$(error "docker not in your system path")
else
$(info "docker found")
endif

all: clean time build up

clean:
	$(RM) time cover.out
test:
	cd cmd/time/apis && $(GO) test -coverprofile=cover.out
time:
	GOOS=linux GOARCH=amd64 $(GO) build -o time ./cmd/time
build: time
	$(COMPOSE) build
up: build
	$(COMPOSE) up