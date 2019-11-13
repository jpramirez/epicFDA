TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.1.1-alpha-0.8
BINARY_NAME=epicFDAFetcher

BINARY_NAME_SERVER=epicFDAFetcher-server.v1


BUILD_FOLDER  = $(shell pwd)/build

FLAGS_LINUX   = CGO_LDFLAGS="-L./LIB -Wl,-rpath -Wl,\$ORIGIN/LIB" CGO_ENABLED=1 GOOS=linux GOARCH=amd64  
FLAGS_DARWIN  = OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.6 CC=o64-clang CXX=o64-clang++ CGO_ENABLED=0
FLAGS_FREEBSD = GOOS=freebsd GOARCH=amd64 CGO_ENABLED=1
FLAGS_WINDOWS = GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 

GOFLAGS_WINDOWS = -ldflags -H=windowsgui

check-env:
	@mkdir -p $(BUILD_FOLDER)/dist/linux/bin
	@mkdir -p $(BUILD_FOLDER)/dist/windows/bin
	@mkdir -p $(BUILD_FOLDER)/dist/arm/bin
	@mkdir -p $(BUILD_FOLDER)/dist/osx/bin
	@mkdir -p $(BUILD_FOLDER)/dist/linux/DataSetFolder/
	cp -R config $(BUILD_FOLDER)/dist/linux/
	cp -R config $(BUILD_FOLDER)/dist/windows/
	cp -R config $(BUILD_FOLDER)/dist/arm/
	cp -R config $(BUILD_FOLDER)/dist/osx/
	cp -R extras $(BUILD_FOLDER)/dist/linux/
	cp -R assets $(BUILD_FOLDER)/dist/linux/




gen_proto:
	third_party/proto-gen.sh

	
## Linting
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...


getdeps:
	./getDeps.sh




versioning:
	./version.sh ${VERSION} ${TIME}

build/weblayer-linux:
	cd cmd/webServer && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/${BINARY_NAME_SERVER} .


build/fetcher-linux:
	cd cmd/runFetcher && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/${BINARY_NAME} .

run/dev:
	cd build/dist/linux && bin/${BINARY_NAME_SERVER} --config config/config.json

build/dev: check-env build/fetcher-linux run/dev




clean:
	rm -Rvf build/dist/


build/all: clean check-env build/fetcher-linux build/weblayer-linux


package-linux:
	cd build/dist/ && tar zcvf linux-dist.tar.gz linux/
	cd build/dist/ && zip -9 linux-dist.zip -r linux/	