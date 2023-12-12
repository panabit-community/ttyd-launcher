.PHONY: all

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

DIST_DIR=./build/dist
PACKAGE=./build/panabit-ttyd.tar.gz
TTYD_PATH=./static/bin/ttyd
TTYD_URL=https://github.com/tsl0922/ttyd/releases/download/1.7.4/ttyd.aarch64

all: clean build package

clean:
	rm -rf $(DIST_DIR)

build: build-ctl build-api build-daemon build-hooks

build-ctl:
	$(GOBUILD) -o $(DIST_DIR)/appctrl -v ./cmd
build-api:
	$(GOBUILD) -o $(DIST_DIR)/web/cgi/api -v ./cmd/cgi/api
build-daemon:
	$(GOBUILD) -o $(DIST_DIR)/daemon -v ./cmd/daemon
build-hooks:
	$(GOBUILD) -o $(DIST_DIR)/afterinstall -v ./cmd/hooks/postinstall

package: $(TTYD_PATH)
	cp -r ./static/* $(DIST_DIR)
	tar czvf $(PACKAGE) -C $(DIST_DIR) .

$(TTYD_PATH):
	wget -O $(TTYD_PATH) $(TTYD_URL)
