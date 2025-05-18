$ErrorActionPreference = "Stop"

# macOS arm64
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o ./bin/jetbra-free-darwin-arm64 cmd/main.go

# macOS amd64
$env:GOARCH = "amd64"
go build -o ./bin/jetbra-free-darwin-amd64 cmd/main.go

# Windows amd64
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o ./bin/jetbra-free-windows-amd64.exe cmd/main.go

# Windows arm64
$env:GOOS = "windows"
$env:GOARCH = "arm64"
go build -o ./bin/jetbra-free-windows-arm64.exe cmd/main.go

# Linux amd64
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o ./bin/jetbra-free-linux-amd64 cmd/main.go

# Linux arm64
$env:GOOS = "linux"
$env:GOARCH = "arm64"
go build -o ./bin/jetbra-free-linux-arm64 cmd/main.go
