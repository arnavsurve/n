.PHONY: build install clean test

BINARY_NAME=n
INSTALL_PATH=$(HOME)/.local/bin

build:
	go build -o $(BINARY_NAME) ./cmd/n

install: build
	mkdir -p $(INSTALL_PATH)
	cp $(BINARY_NAME) $(INSTALL_PATH)/
	@echo "Installed to $(INSTALL_PATH)/$(BINARY_NAME)"

clean:
	rm -f $(BINARY_NAME)
	go clean

test:
	go test -v ./...
