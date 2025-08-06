default: build

.PHONY: build test testacc install generate docs clean

build:
	go build -v ./...

test:
	go test -v ./...

testacc:
	TF_ACC=1 go test -v ./internal/provider/

install: build
	go install -v ./...

generate:
	go generate -v ./...

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

clean:
	go clean -v ./...
	rm -f terraform-provider-hex

# Build for multiple platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -o ./build/terraform-provider-hex_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o ./build/terraform-provider-hex_darwin_arm64
	GOOS=linux GOARCH=amd64 go build -o ./build/terraform-provider-hex_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o ./build/terraform-provider-hex_linux_arm64
	GOOS=windows GOARCH=amd64 go build -o ./build/terraform-provider-hex_windows_amd64.exe

lint:
	golangci-lint run

fmt:
	go fmt ./...