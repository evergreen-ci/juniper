proto:
	@mkdir -p golang
	protoc --go_out=plugins=grpc:golang *.proto
