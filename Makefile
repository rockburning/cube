VERSION ?= 0.0.1
COMMIT = $(shell git log --format="%h" -n 1|tr -d '\n')
TIMESTAMP = $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
BUILDTAGS = exclude_graphdriver_devicemapper exclude_graphdriver_btrfs containers_image_openpgp

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/;s/i.86/386/;s/ //g')




cube:
	@echo "Compiling pcfusion source for $(OS) $(ARCH) $(COMMIT)"
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -tags "$(BUILDTAGS)" -ldflags "-s -w -X 'main.version=$(VERSION)-$(COMMIT)-$(TIMESTAMP)'" -o cube cmd/cube/main.go


clean:
	@rm -f  ./cube
