# Variables
BINARY_NAME=jetbra-free
BIN_DIR=./bin
EMBED_FILE=internal/util/bindata.go
GO_BINDATA=go-bindata
SRC_DIRS=static/... templates/...

# 默认目标
all: build

# 编译
build: bindata-access
	go build -o ./bin/$(BINARY_NAME) cmd/main.go

# 运行
run: build
	./bin/$(BINARY_NAME)

# 编译所有平台
build-all: build-mac build-mac-arm build-windows build-linux

# 编译 MacOS AMD64
build-mac: bindata-access
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64 cmd/main.go

# 编译 MacOS ARM64 (Apple Silicon)
build-mac-arm: bindata-access
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 cmd/main.go

# 编译 Windows AMD64
build-windows: bindata-access
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe cmd/main.go

# 编译 Linux AMD64
build-linux: bindata-access
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 cmd/main.go

# 清理
clean:
	rm -rf ./bin/

# 安装 go-bindata
install-bindata:
	go install github.com/go-bindata/go-bindata/v3/go-bindata@latest

bindata-access:
	go-bindata -o internal/util/access.go -pkg util $(SRC_DIRS)

# windows 下编译用 build.ps1