THIRD_PARTY_PROTO_FILES=$(shell find ../backend/third_party -name *.proto)
API_PROTO_FILES=$(shell find ../backend/api -name *.proto)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=../backend/api \
	       --proto_path=../backend/third_party \
	       --ts_proto_out=outputServices=false,esModuleInterop=true:./src/api \
	       $(API_PROTO_FILES) $(THIRD_PARTY_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
