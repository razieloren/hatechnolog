PROTOC := protoc

.PHONY: all clean
default: all

all:
	protoc --go_out=.. --ts_out=../frontend/app/messages stats/helpers.proto
	protoc --go_out=.. --ts_out=../frontend/app/messages stats/api.proto
	protoc --go_out=.. --ts_out=../frontend/app/messages wrapper.proto

clean:
	rm -rf ../backend/modules/api/endpoints/messages ../frontend/app/messages/*.ts ../frontend/app/messages/stats