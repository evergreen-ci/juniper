# start project configuration
buildDir := build
pbDir := gopb
# end project configuration

# start environment setup
gobin := go
ifneq (,$(GOROOT))
gobin := $(GOROOT)/bin/go
endif

goCache := $(GOCACHE)
ifeq (,$(goCache))
goCache := $(abspath $(buildDir)/.cache)
endif
goModCache := $(GOMODCACHE)
ifeq (,$(goModCache))
goModCache := $(abspath $(buildDir)/.mod-cache)
endif

ifneq ($(goCache),$(GOCACHE))
export GOCACHE := $(goCache)
endif
ifneq ($(goModCache),$(GOMODCACHE))
export GOMODCACHE := $(goModCache)
endif
# end environment setup

# ensure the build directory exists, since most targets require it
$(shell mkdir -p $(buildDir))

# start basic development targets
protocVersion := 3.6.1
protocGenGoVersion := 1.3.2
protoOS := $(shell uname -s | tr A-Z a-z)
ifeq ($(protoOS),darwin)
protoOS := osx
endif
protoOS := $(protoOS)-$(shell uname -m | tr A-Z a-z)
$(buildDir)/protoc:
	curl --retry 10 --retry-max-time 60 -L0 https://github.com/protocolbuffers/protobuf/releases/download/v$(protocVersion)/protoc-$(protocVersion)-$(protoOS).zip --output protoc.zip
	unzip -q protoc.zip -d $(buildDir)/protoc
	rm -f protoc.zip
	GOBIN="$(abspath $(buildDir))" $(gobin) install github.com/golang/protobuf/protoc-gen-go@v$(protocGenGoVersion)
proto: $(buildDir)/protoc
	mkdir -p $(pbDir)
	PATH="$(abspath $(buildDir)):$(PATH)" $(buildDir)/protoc/bin/protoc --go_out=plugins=grpc:$(pbDir) *.proto
# end basic development targets

# start cleanup targets
clean:
	rm -rf $(buildDir) $(pbDir)
# end cleanup targets
