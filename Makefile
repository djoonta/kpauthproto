PROTOC_MAIN = ./kpauthproto.proto
PROTOC_AUTH = ./auth/auth.proto
PROTOC = protoc
PROTOC_FLAGS = -I . --go_out=. --go_opt=paths=source_relative \
							 --go-grpc_out=. --go-grpc_opt=paths=source_relative

auth:	
	@echo "auth proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_AUTH)

main:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_MAIN)

all: auth main

.PHONY: auth main all
