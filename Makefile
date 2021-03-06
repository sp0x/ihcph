# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

# Coverage
COVERAGE_REPORT = coverage.txt
COVERAGE_MODE = atomic
NAME=ihcph
AUTHOR=sp0x
ARCH=amd64
OS=linux darwin windows
SERVICE=ihcph

ifneq ($(origin CI), undefined)
	WORKDIR := $(GOPATH)/src/github.com/$(NAME)
endif

build: assets
	go build -o $(NAME) -ldflags "-s -w" ./

assets:
	@echo "Embedding assets as code"
	bindata -o sites/assets.go ./sites/...

build-multi-arch:
	gox -os="${OS}" -arch="${ARCH}" -output="$(NAME).{{.OS}}.{{.Arch}}" -ldflags "-s -w -X main.Rev=`git rev-parse --short HEAD`" -verbose ./...

install-deps:
	@echo "Installing go utils"
	go get github.com/kataras/bindata/cmd/bindata


install:
	go build -i -o $(GOPATH)/bin/$(NAME) ./cmd

test:
	@cd $(WORKDIR); \
	$(GOTEST) ./...

test-coverage:
	@cd $(WORKDIR); \
	echo "" > $(COVERAGE_REPORT); \
	$(GOTEST) -coverprofile=$(COVERAGE_REPORT) -coverpkg=./... -covermode=$(COVERAGE_MODE) ./...

build-image:
	docker-compose build $(SERVICE)
	docker-compose push

benchmark-server:
	#go-wrk -redir -d 20  -c 100 http://127.0.0.1:5000/status
	go-wrk -redir -T 2000 -d 1 -c 1 "http://127.0.0.1:5000/torznab/all/api?t=search&cat=2000,2010,2020,2030,2035,2040,2045,2050,2060&extended=1&apikey=210fc7bb818639a&offset=0&limit=100&q=accountant 2016"