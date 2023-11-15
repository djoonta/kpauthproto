PROTOC_MAIN = ./kpauthproto.proto
PROTOC_AUTH = ./auth/auth.proto
PROTOC_ROLE = ./role/role.proto
PROTOC_USER = ./user/user.proto
PROTOC = protoc
PROTOC_FLAGS = -I . --go_out=. --go_opt=paths=source_relative \
							 --go-grpc_out=. --go-grpc_opt=paths=source_relative

auth:	
	@echo "auth proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_AUTH)

role:
	@echo "role proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_ROLE)

user:
	@echo "role proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_USER)

main:
	@echo "create proto..."
	$(PROTOC) $(PROTOC_FLAGS) $(PROTOC_MAIN)

all: auth role user main

.PHONY: auth role user main all
