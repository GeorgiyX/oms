CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/checkout/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

install-go-deps:
	GOBIN=$(BINDIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BINDIR) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-go-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

vendor-proto:
	mkdir -p vendor-proto
	@if [ ! -d vendor-proto/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
		mkdir -p  vendor-proto/google/ &&\
		mv vendor-proto/googleapis/google/api vendor-proto/google &&\
		rm -rf vendor-proto/googleapis ;\
	fi
	@if [ ! -d vendor-proto/google/protobuf ]; then\
		git clone https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
		mkdir -p  vendor-proto/google/protobuf &&\
		mv vendor-proto/protobuf/src/google/protobuf/*.proto vendor-proto/google/protobuf &&\
		rm -rf vendor-proto/protobuf ;\
	fi

generate:
	mkdir -p pkg/checkout
	mkdir -p pkg/product-service
	protoc -I api/checkout -I api/product-service -I vendor-proto \
	--go_out=pkg/checkout --go_opt=paths=source_relative \
	--go-grpc_out=pkg/checkout --go-grpc_opt=paths=source_relative \
	api/checkout/checkout.proto
	protoc -I vendor-proto/product-service -I vendor-proto \
	--go_out=pkg/product-service --go_opt=paths=source_relative \
	--go-grpc_out=pkg/product-service --go-grpc_opt=paths=source_relative \
	vendor-proto/product-service/product-service.proto

migrate:
	source .env && cd migrations && ./migrate.sh
