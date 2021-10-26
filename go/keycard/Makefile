.PHONY: build-lib

BUILD=./build

ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
 detected_OS := Windows
else
 detected_OS := $(strip $(shell uname))
endif

ifeq ($(detected_OS),Darwin)
 LIB_EXT := dylib
 ifeq ("$(shell sysctl -nq hw.optional.arm64)","1")
  # Building on M1 is still not supported, so in the meantime we crosscompile to amd64
  CGOFLAGS=CGO_ENABLED=1 GOOS=darwin GOARCH=amd64
 endif
else ifeq ($(detected_OS),Windows)
 LIB_EXT:= dll
 LIBKEYCARD_EXT := dll
else
 SHARED_LIB_CGO_LDFLAGS := CGO_LDFLAGS="-Wl,-soname,libkeycard.so.0"
 LIB_EXT := so.0
endif

build-lib:
	mkdir -p $(BUILD)/libkeycard
	@echo "Building dynamic library..."
	$(CGOFLAGS) $(SHARED_LIB_CGO_LDFLAGS) go build -buildmode=c-shared -o $(BUILD)/libkeycard/libkeycard.$(LIB_EXT) . 
ifeq ($(detected_OS),Linux)
	cd $(BUILD)/libkeycard/ && \
	ln -f -s ./libkeycard.$(LIB_EXT) ./libkeycard.so
endif
	@echo "Dynamic library built:"
	@ls -la $(BUILD)/libkeycard/*
