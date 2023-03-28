
build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker-compose up --force-recreate --build

test:
	go test -coverprofile=coverage.out -cover -coverpkg ./checkout/internal/usecases/checkout,./loms/internal/usecase/loms ./checkout/... ./loms/... && \
	go tool cover -func=coverage.out && \
	rm coverage.out


precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit
