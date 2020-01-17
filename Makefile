MAIN_PKG:=az-util
MAIN_PREFIX=$(dir $(MAIN_PKG))
MAIN=$(subst $(MAIN_PREFIX), , $(MAIN_PKG))
BIN=$(strip $(MAIN))

export GOPATH=$(shell pwd)/../../../../
export AZBIT_KUBERNETES_IDC=suzhou

build:
	go build -tags=jsoniter -x -o run/$(BIN) gitlab.azbit.cn/web/$(MAIN_PKG)

dev:
	go run main.go $(ARG)

run: build
	cd run && ./$(BIN) $(ARG)

init:
	cd run && TARGET='run' ARG='init' docker-compose run --rm az-util-devel

docker-build:
	#cd run && \
	#TARGET='build' ARG='server' docker-compose run --rm az-util-devel && cp $(BIN) ../build/ &&
	cd run && cp $(BIN) ../build/ && \
	cd ../build && \
	docker build -t zzsure/az-util:$(TAG) . && \
	docker push zzsure/az-util:$(TAG)

.PHONY: build
