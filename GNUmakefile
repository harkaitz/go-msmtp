PROJECT=go-msmtp
VERSION=1.0.0
PREFIX=/usr/local
all:
clean:
install:

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps:
build/go-msmtp$(EXE): deps
	@mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/go-msmtp
all-go: build/go-msmtp$(EXE)
install-go:
	install -D -t $(DESTDIR)$(PREFIX)/bin build/go-msmtp$(EXE)
clean-go:
	rm -f build/go-msmtp$(EXE)
## -- BLOCK:go --
