include build_config.mk

all: build


build:
	$(GO) install  ./...

clean:
	$(GO) clean -i ./...
