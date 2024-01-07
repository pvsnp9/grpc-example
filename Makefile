GO_MODULE := github.com/pvsnp9/grpc-example


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: run
run:
	go run cmd/main.go


# .PHONY: execute
# execute: build run


# .PHONY: protoc-validate
# protoc-validate:
# 	protoc --validate_out="lang=go:./pkg/pb/generated" --go_opt=module=${GO_MODULE} --go_out=. ./pkg/pb/valexample/*.proto




###########


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
	mkdir protogen\go
else
	rm -fR ./protogen 
	mkdir -p ./protogen/go
endif


.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./pkg/proto/hello/*.proto ./pkg/proto/payment/*.proto ./pkg/proto/transaction/*.proto \

.PHONY: build
build: clean protoc-go


.PHONY: pipeline-init
pipeline-init:
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


.PHONY: pipeline-build
pipeline-build: pipeline-init build
