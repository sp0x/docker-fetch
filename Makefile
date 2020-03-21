# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

# Git config
GIT_VERSION ?=
GIT_DIST_PATH ?= $(PWD)/.git-dist
GIT_REPOSITORY = http://github.com/git/git.git

# Coverage
COVERAGE_REPORT = coverage.txt
COVERAGE_MODE = atomic

ifneq ($(origin CI), undefined)
	WORKDIR := $(GOPATH)/src/github.com/docker-fetch
endif

build:
	gox -os="linux" -arch="amd64" -output="docker-fetch.{{.OS}}.{{.Arch}}" -ldflags "-X main.Rev=`git rev-parse --short HEAD`" -verbose ./...

test:
	@cd $(WORKDIR); \
	$(GOTEST) ./...

test-coverage:
	@cd $(WORKDIR); \
	echo "" > $(COVERAGE_REPORT); \
	$(GOTEST) -coverprofile=$(COVERAGE_REPORT) -coverpkg=./... -covermode=$(COVERAGE_MODE) ./...

clean:
	rm -rf $(GIT_DIST_PATH)