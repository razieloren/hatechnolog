/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 4.23.4
 * source: wrapper.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./stats/api";
import * as pb_1 from "google-protobuf";
export namespace messages {
    export enum Endpoint {
        latest_stats_push = 0
    }
    export class EndpointIntent extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            endpoint?: Endpoint;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("endpoint" in data && data.endpoint != undefined) {
                    this.endpoint = data.endpoint;
                }
            }
        }
        get endpoint() {
            return pb_1.Message.getFieldWithDefault(this, 1, Endpoint.latest_stats_push) as Endpoint;
        }
        set endpoint(value: Endpoint) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            endpoint?: Endpoint;
        }): EndpointIntent {
            const message = new EndpointIntent({});
            if (data.endpoint != null) {
                message.endpoint = data.endpoint;
            }
            return message;
        }
        toObject() {
            const data: {
                endpoint?: Endpoint;
            } = {};
            if (this.endpoint != null) {
                data.endpoint = this.endpoint;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.endpoint != Endpoint.latest_stats_push)
                writer.writeEnum(1, this.endpoint);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): EndpointIntent {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new EndpointIntent();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.endpoint = reader.readEnum();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): EndpointIntent {
            return EndpointIntent.deserialize(bytes);
        }
    }
    export class Wrapper extends pb_1.Message {
        #one_of_decls: number[][] = [[1, 2, 3]];
        constructor(data?: any[] | ({} & (({
            endpoint_intent?: EndpointIntent;
            latest_stats_push_request?: never;
            latest_stats_push_response?: never;
        } | {
            endpoint_intent?: never;
            latest_stats_push_request?: dependency_1.messages.LatestStatsPushRequest;
            latest_stats_push_response?: never;
        } | {
            endpoint_intent?: never;
            latest_stats_push_request?: never;
            latest_stats_push_response?: dependency_1.messages.LatestStatsPushResponse;
        })))) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("endpoint_intent" in data && data.endpoint_intent != undefined) {
                    this.endpoint_intent = data.endpoint_intent;
                }
                if ("latest_stats_push_request" in data && data.latest_stats_push_request != undefined) {
                    this.latest_stats_push_request = data.latest_stats_push_request;
                }
                if ("latest_stats_push_response" in data && data.latest_stats_push_response != undefined) {
                    this.latest_stats_push_response = data.latest_stats_push_response;
                }
            }
        }
        get endpoint_intent() {
            return pb_1.Message.getWrapperField(this, EndpointIntent, 1) as EndpointIntent;
        }
        set endpoint_intent(value: EndpointIntent) {
            pb_1.Message.setOneofWrapperField(this, 1, this.#one_of_decls[0], value);
        }
        get has_endpoint_intent() {
            return pb_1.Message.getField(this, 1) != null;
        }
        get latest_stats_push_request() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.LatestStatsPushRequest, 2) as dependency_1.messages.LatestStatsPushRequest;
        }
        set latest_stats_push_request(value: dependency_1.messages.LatestStatsPushRequest) {
            pb_1.Message.setOneofWrapperField(this, 2, this.#one_of_decls[0], value);
        }
        get has_latest_stats_push_request() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get latest_stats_push_response() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.LatestStatsPushResponse, 3) as dependency_1.messages.LatestStatsPushResponse;
        }
        set latest_stats_push_response(value: dependency_1.messages.LatestStatsPushResponse) {
            pb_1.Message.setOneofWrapperField(this, 3, this.#one_of_decls[0], value);
        }
        get has_latest_stats_push_response() {
            return pb_1.Message.getField(this, 3) != null;
        }
        get message() {
            const cases: {
                [index: number]: "none" | "endpoint_intent" | "latest_stats_push_request" | "latest_stats_push_response";
            } = {
                0: "none",
                1: "endpoint_intent",
                2: "latest_stats_push_request",
                3: "latest_stats_push_response"
            };
            return cases[pb_1.Message.computeOneofCase(this, [1, 2, 3])];
        }
        static fromObject(data: {
            endpoint_intent?: ReturnType<typeof EndpointIntent.prototype.toObject>;
            latest_stats_push_request?: ReturnType<typeof dependency_1.messages.LatestStatsPushRequest.prototype.toObject>;
            latest_stats_push_response?: ReturnType<typeof dependency_1.messages.LatestStatsPushResponse.prototype.toObject>;
        }): Wrapper {
            const message = new Wrapper({});
            if (data.endpoint_intent != null) {
                message.endpoint_intent = EndpointIntent.fromObject(data.endpoint_intent);
            }
            if (data.latest_stats_push_request != null) {
                message.latest_stats_push_request = dependency_1.messages.LatestStatsPushRequest.fromObject(data.latest_stats_push_request);
            }
            if (data.latest_stats_push_response != null) {
                message.latest_stats_push_response = dependency_1.messages.LatestStatsPushResponse.fromObject(data.latest_stats_push_response);
            }
            return message;
        }
        toObject() {
            const data: {
                endpoint_intent?: ReturnType<typeof EndpointIntent.prototype.toObject>;
                latest_stats_push_request?: ReturnType<typeof dependency_1.messages.LatestStatsPushRequest.prototype.toObject>;
                latest_stats_push_response?: ReturnType<typeof dependency_1.messages.LatestStatsPushResponse.prototype.toObject>;
            } = {};
            if (this.endpoint_intent != null) {
                data.endpoint_intent = this.endpoint_intent.toObject();
            }
            if (this.latest_stats_push_request != null) {
                data.latest_stats_push_request = this.latest_stats_push_request.toObject();
            }
            if (this.latest_stats_push_response != null) {
                data.latest_stats_push_response = this.latest_stats_push_response.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_endpoint_intent)
                writer.writeMessage(1, this.endpoint_intent, () => this.endpoint_intent.serialize(writer));
            if (this.has_latest_stats_push_request)
                writer.writeMessage(2, this.latest_stats_push_request, () => this.latest_stats_push_request.serialize(writer));
            if (this.has_latest_stats_push_response)
                writer.writeMessage(3, this.latest_stats_push_response, () => this.latest_stats_push_response.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Wrapper {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Wrapper();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.endpoint_intent, () => message.endpoint_intent = EndpointIntent.deserialize(reader));
                        break;
                    case 2:
                        reader.readMessage(message.latest_stats_push_request, () => message.latest_stats_push_request = dependency_1.messages.LatestStatsPushRequest.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.latest_stats_push_response, () => message.latest_stats_push_response = dependency_1.messages.LatestStatsPushResponse.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Wrapper {
            return Wrapper.deserialize(bytes);
        }
    }
}