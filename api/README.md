# API

Shared messages between the backend and the frontend, basef on Google's RPC - `protobuf` (version 3).

## Setup

- [Install](https://grpc.io/docs/protoc-installation/) `protobuf`'s compiler on your machine.
- Add the golang module (Of course go needs to be already installed on your machine) `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
    - Make sure `$GOPAH` (Usually: `~/go`) is in your `PATH`.
- doc protobuf (protoc)
- Globally install the `typescript` module: `npm install -g protoc-gen-ts`
- To make sure that everything works, in this folder run `make all` and observe the results - All the messages should have been compiled successfully with no errors or warnings.

## Directory Tree
- `wrapper.proto` - Defines wrapper messages for all API requests & responses, and defines the API endpoints inetens.
- `[ENDPOINT_NAME]` - Indicates an endpoint.
    - `api.proto` - Requests & Responses to the API
    - `helpers.proto` - Helper messages that the API messages may use.

## Adding New Messages

### To an Existing Endpoint

- Just add the relevant `api.proto` & `helpers.proto` of the relevant API you would like to change.

### To a New Endpoint

- Create a folder with the endpoint's name.
- Create the `api.proto` & `helpers.proto` file, add your messages that define the API.
- Import the `api.proto` file in `wrapper.proto`, and add the new API messages under the `oneof` field in the `wrapper` message.
- Compile your messages in the `makefile` before the `wrapper.proto`, and also make sure to add a proper `cleanup` rule.
