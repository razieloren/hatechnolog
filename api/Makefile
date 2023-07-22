PROTOC := protoc

.PHONY: all clean
default: all

all:
	protoc --go_out=.. --plugin=protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --ts_out=../frontend/messages stats/helpers.proto
	protoc --go_out=.. --plugin=protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --ts_out=../frontend/messages stats/api.proto
	protoc --go_out=.. --plugin=protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --ts_out=../frontend/messages user/api.proto
	protoc --go_out=.. --plugin=protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --ts_out=../frontend/messages wrapper.proto
	protoc --go_out=.. backend.proto

clean:
	rm -rf ../backend/x/messages
	rm -rf ../backend/modules/api/endpoints/messages
	rm -rf ../frontend/messages/*.ts
	rm -rf ../frontend/messages/stats
	rm -rf ../frontend/messages/user