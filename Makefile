.PHONY: build run build-keycard-go

KEYCARD_GO_PATH=./go/keycard

build: build-keycard-go
	nim c --passL:"-L$(KEYCARD_GO_PATH)/build/libkeycard" --passL="-lkeycard" keycard.nim

build-keycard-go:
	cd $(KEYCARD_GO_PATH) && make build-lib

run: build
	LD_LIBRARY_PATH=$(KEYCARD_GO_PATH)/build/libkeycard ./keycard
