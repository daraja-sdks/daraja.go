
PROXY=GOPROXY=proxy.golang.org
URI=github.com/daraja-sdks/daraja
ARTIFACTS=*.prof *.out *.bench *.exe

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in your PATH, run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

all: fmt lint test

fmt:
	$(info *********** running format checks ***********)
	go fmt

test:
	$(info *********** running tests ***********)
	go test

lint:
	$(info *********** running linting ***********)
	golangci-lint run

release:
	$(GOPROXY) go list -m $(URI)@$(VERSION)

coverage:
	$(info *********** checking coverage ***********)
	go test -coverprofile=coverage.out

reports: coverage
	go tool cover -html=coverage.out

clean: 
	$(info *********** cleaning build and test artifacts ***********)
	$(RM) -r $(ARTIFACTS)

