
# download protoc at: https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.1
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


protoc:
	for x in proto/**/*.proto; do ./bin/protoc -I=proto --go_out=paths=source_relative:./go $$x; done
	./bin/protoc -I=proto --csharp_out=./cs proto/**/*.proto
	