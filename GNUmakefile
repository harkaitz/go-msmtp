.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-msmtp
VERSION=1.0.0
PREFIX=/usr/local
BUILDDIR?=.build

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/go-msmtp$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/go-msmtp$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/go-msmtp$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/go-msmtp$(EXE)
##
$(BUILDDIR)/go-msmtp$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/go-msmtp
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: README.md LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp README.md LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
