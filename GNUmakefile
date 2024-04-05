.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-msmtp
VERSION=1.0.0
PREFIX=/usr/local

## -- BLOCK:go --
build/go-msmtp$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/go-msmtp
all: build/go-msmtp$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/go-msmtp$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/go-msmtp$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
