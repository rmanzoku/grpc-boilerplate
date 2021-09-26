PKG=github.com/rmanzoku/grpc-boilerplate
FEATURE=go/feature
SERVICE=go/service
PLATFORM=--platform=linux/x86_64
UID := $(shell id -u)

.PHONY: protoc protoc-all

protoc: proto/*.proto
	cd proto; \
	for n in $(subst .proto, , $(notdir $^)); do \
		docker run $(PLATFORM) -u $(UID) --rm -v `pwd`:/defs protoc-all -f $$n.proto -o gen/pb-go/$$n -l go --with-gateway --with-openapi-json-names; \
		docker run $(PLATFORM) -u $(UID) --rm -v `pwd`:/defs protoc-all -f $$n.proto -o gen/pb-web/$$n -l web; \
		mv gen/pb-go/$$n/$$n.swagger.json gen/pb-web/$$n; \
		mkdir -p ../$(FEATURE)/$$n; \
		mkdir -p ../$(SERVICE)/$$n; \
		touch ../$(SERVICE)/$$n/service.go; \
		mv gen/pb-go/$$n/$(PKG)/$(FEATURE)/$$n/* ../$(FEATURE)/$$n/; \
		mkdir -p ../typescript/assets/protobuf/web; \
		mv gen/pb-web/$$n/* ../typescript/assets/protobuf/web/; \
	done
	mkdir -p typescript/assets/protobuf/web/google/api/
	echo {} > typescript/assets/protobuf/web/google/api/annotations_pb.js
	rm -r proto/gen

protoc-all:
	docker build containers -t protoc-all
