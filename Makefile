generate_grpc_code:
	protoc --go_out=compute --go_opt=paths=source_relative --plugin=protoc-gen-go-grpc=/home/navneet/go/bin/protoc-gen-go-grpc --go-grpc_out=compute --go-grpc_opt=paths=source_relative compute.proto
