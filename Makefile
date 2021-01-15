PROTOC_ALL_VER := 1.34_0
PKG=github.com/rmanzoku/grpc-boilerplate
FEATURE=feature
SERVICE=service
UID := $(shell id -u)

.PHONY: protoc

protoc: proto/*.proto
	cd proto; \
	for n in $(subst .proto, , $(notdir $^)); do \
		docker run -u $(UID) --rm -v `pwd`:/defs namely/protoc-all:$(PROTOC_ALL_VER) -f $$n.proto -o gen/pb-go/$$n -l go --with-gateway --with-openapi-json-names; \
		mv gen/pb-go/$$n/$$n.swagger.json gen/pb-web/$$n; \
		mkdir -p ../$(FEATURE)/$$n; \
		mkdir -p ../$(SERVICE)/$$n; \
		touch ../$(SERVICE)/$$n/service.go; \
		mv gen/pb-go/$$n/$(PKG)/$(FEATURE)/$$n/* ../$(FEATURE)/$$n/; \
	done
	rm -r proto/gen


