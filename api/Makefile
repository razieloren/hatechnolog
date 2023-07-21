PROTOC := protoc

.PHONY: all clean
default: all

all:
	protoc --go_out=.. --ts_out=../frontend/messages stats/helpers.proto
	protoc --go_out=.. --ts_out=../frontend/messages stats/api.proto
	protoc --go_out=.. --ts_out=../frontend/messages wrapper.proto
	protoc --go_out=.. backend.proto

clean:
	rm -rf ../backend/x/messages
	rm -rf ../backend/modules/api/endpoints/messages
	rm -rf ../frontend/messages/*.ts
	rm -rf ../frontend/messages/stats