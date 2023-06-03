clean:
	@rm -vf ./shared/proto/greeting/*.pb.go
protoc: clean
	@cd ./shared/proto/greeting && protoc --go_out=. --go-grpc_out=. *.proto