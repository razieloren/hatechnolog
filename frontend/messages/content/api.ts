/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 4.23.4
 * source: content/api.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./helpers";
import * as pb_1 from "google-protobuf";
export namespace messages {
    export class GetPageRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            slug?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("slug" in data && data.slug != undefined) {
                    this.slug = data.slug;
                }
            }
        }
        get slug() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set slug(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            slug?: string;
        }): GetPageRequest {
            const message = new GetPageRequest({});
            if (data.slug != null) {
                message.slug = data.slug;
            }
            return message;
        }
        toObject() {
            const data: {
                slug?: string;
            } = {};
            if (this.slug != null) {
                data.slug = this.slug;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.slug.length)
                writer.writeString(1, this.slug);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPageRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPageRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.slug = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPageRequest {
            return GetPageRequest.deserialize(bytes);
        }
    }
    export class GetPageResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            details?: dependency_1.messages.ContentDetails;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("details" in data && data.details != undefined) {
                    this.details = data.details;
                }
            }
        }
        get details() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.ContentDetails, 1) as dependency_1.messages.ContentDetails;
        }
        set details(value: dependency_1.messages.ContentDetails) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_details() {
            return pb_1.Message.getField(this, 1) != null;
        }
        static fromObject(data: {
            details?: ReturnType<typeof dependency_1.messages.ContentDetails.prototype.toObject>;
        }): GetPageResponse {
            const message = new GetPageResponse({});
            if (data.details != null) {
                message.details = dependency_1.messages.ContentDetails.fromObject(data.details);
            }
            return message;
        }
        toObject() {
            const data: {
                details?: ReturnType<typeof dependency_1.messages.ContentDetails.prototype.toObject>;
            } = {};
            if (this.details != null) {
                data.details = this.details.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_details)
                writer.writeMessage(1, this.details, () => this.details.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPageResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPageResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.details, () => message.details = dependency_1.messages.ContentDetails.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPageResponse {
            return GetPageResponse.deserialize(bytes);
        }
    }
    export class GetPostsTeasersRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {}) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data: {}): GetPostsTeasersRequest {
            const message = new GetPostsTeasersRequest({});
            return message;
        }
        toObject() {
            const data: {} = {};
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPostsTeasersRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPostsTeasersRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPostsTeasersRequest {
            return GetPostsTeasersRequest.deserialize(bytes);
        }
    }
    export class GetPostsTeasersResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            teasers?: dependency_1.messages.ContentTeaser[];
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("teasers" in data && data.teasers != undefined) {
                    this.teasers = data.teasers;
                }
            }
        }
        get teasers() {
            return pb_1.Message.getRepeatedWrapperField(this, dependency_1.messages.ContentTeaser, 1) as dependency_1.messages.ContentTeaser[];
        }
        set teasers(value: dependency_1.messages.ContentTeaser[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        static fromObject(data: {
            teasers?: ReturnType<typeof dependency_1.messages.ContentTeaser.prototype.toObject>[];
        }): GetPostsTeasersResponse {
            const message = new GetPostsTeasersResponse({});
            if (data.teasers != null) {
                message.teasers = data.teasers.map(item => dependency_1.messages.ContentTeaser.fromObject(item));
            }
            return message;
        }
        toObject() {
            const data: {
                teasers?: ReturnType<typeof dependency_1.messages.ContentTeaser.prototype.toObject>[];
            } = {};
            if (this.teasers != null) {
                data.teasers = this.teasers.map((item: dependency_1.messages.ContentTeaser) => item.toObject());
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.teasers.length)
                writer.writeRepeatedMessage(1, this.teasers, (item: dependency_1.messages.ContentTeaser) => item.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPostsTeasersResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPostsTeasersResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.teasers, () => pb_1.Message.addToRepeatedWrapperField(message, 1, dependency_1.messages.ContentTeaser.deserialize(reader), dependency_1.messages.ContentTeaser));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPostsTeasersResponse {
            return GetPostsTeasersResponse.deserialize(bytes);
        }
    }
    export class GetPostRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            slug?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("slug" in data && data.slug != undefined) {
                    this.slug = data.slug;
                }
            }
        }
        get slug() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set slug(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            slug?: string;
        }): GetPostRequest {
            const message = new GetPostRequest({});
            if (data.slug != null) {
                message.slug = data.slug;
            }
            return message;
        }
        toObject() {
            const data: {
                slug?: string;
            } = {};
            if (this.slug != null) {
                data.slug = this.slug;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.slug.length)
                writer.writeString(1, this.slug);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPostRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPostRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.slug = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPostRequest {
            return GetPostRequest.deserialize(bytes);
        }
    }
    export class GetPostResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            details?: dependency_1.messages.ContentDetails;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("details" in data && data.details != undefined) {
                    this.details = data.details;
                }
            }
        }
        get details() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.ContentDetails, 1) as dependency_1.messages.ContentDetails;
        }
        set details(value: dependency_1.messages.ContentDetails) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_details() {
            return pb_1.Message.getField(this, 1) != null;
        }
        static fromObject(data: {
            details?: ReturnType<typeof dependency_1.messages.ContentDetails.prototype.toObject>;
        }): GetPostResponse {
            const message = new GetPostResponse({});
            if (data.details != null) {
                message.details = dependency_1.messages.ContentDetails.fromObject(data.details);
            }
            return message;
        }
        toObject() {
            const data: {
                details?: ReturnType<typeof dependency_1.messages.ContentDetails.prototype.toObject>;
            } = {};
            if (this.details != null) {
                data.details = this.details.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_details)
                writer.writeMessage(1, this.details, () => this.details.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPostResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPostResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.details, () => message.details = dependency_1.messages.ContentDetails.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPostResponse {
            return GetPostResponse.deserialize(bytes);
        }
    }
    export class GetCategoriesTeasersRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {}) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data: {}): GetCategoriesTeasersRequest {
            const message = new GetCategoriesTeasersRequest({});
            return message;
        }
        toObject() {
            const data: {} = {};
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetCategoriesTeasersRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetCategoriesTeasersRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetCategoriesTeasersRequest {
            return GetCategoriesTeasersRequest.deserialize(bytes);
        }
    }
    export class GetCategoriesTeasersResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            teasers?: dependency_1.messages.CategoryTeaser[];
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("teasers" in data && data.teasers != undefined) {
                    this.teasers = data.teasers;
                }
            }
        }
        get teasers() {
            return pb_1.Message.getRepeatedWrapperField(this, dependency_1.messages.CategoryTeaser, 1) as dependency_1.messages.CategoryTeaser[];
        }
        set teasers(value: dependency_1.messages.CategoryTeaser[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        static fromObject(data: {
            teasers?: ReturnType<typeof dependency_1.messages.CategoryTeaser.prototype.toObject>[];
        }): GetCategoriesTeasersResponse {
            const message = new GetCategoriesTeasersResponse({});
            if (data.teasers != null) {
                message.teasers = data.teasers.map(item => dependency_1.messages.CategoryTeaser.fromObject(item));
            }
            return message;
        }
        toObject() {
            const data: {
                teasers?: ReturnType<typeof dependency_1.messages.CategoryTeaser.prototype.toObject>[];
            } = {};
            if (this.teasers != null) {
                data.teasers = this.teasers.map((item: dependency_1.messages.CategoryTeaser) => item.toObject());
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.teasers.length)
                writer.writeRepeatedMessage(1, this.teasers, (item: dependency_1.messages.CategoryTeaser) => item.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetCategoriesTeasersResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetCategoriesTeasersResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.teasers, () => pb_1.Message.addToRepeatedWrapperField(message, 1, dependency_1.messages.CategoryTeaser.deserialize(reader), dependency_1.messages.CategoryTeaser));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetCategoriesTeasersResponse {
            return GetCategoriesTeasersResponse.deserialize(bytes);
        }
    }
    export class GetCategoryRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            slug?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("slug" in data && data.slug != undefined) {
                    this.slug = data.slug;
                }
            }
        }
        get slug() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set slug(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            slug?: string;
        }): GetCategoryRequest {
            const message = new GetCategoryRequest({});
            if (data.slug != null) {
                message.slug = data.slug;
            }
            return message;
        }
        toObject() {
            const data: {
                slug?: string;
            } = {};
            if (this.slug != null) {
                data.slug = this.slug;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.slug.length)
                writer.writeString(1, this.slug);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetCategoryRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetCategoryRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.slug = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetCategoryRequest {
            return GetCategoryRequest.deserialize(bytes);
        }
    }
    export class GetCategoryResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            details?: dependency_1.messages.CategoryDetails;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("details" in data && data.details != undefined) {
                    this.details = data.details;
                }
            }
        }
        get details() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.CategoryDetails, 1) as dependency_1.messages.CategoryDetails;
        }
        set details(value: dependency_1.messages.CategoryDetails) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_details() {
            return pb_1.Message.getField(this, 1) != null;
        }
        static fromObject(data: {
            details?: ReturnType<typeof dependency_1.messages.CategoryDetails.prototype.toObject>;
        }): GetCategoryResponse {
            const message = new GetCategoryResponse({});
            if (data.details != null) {
                message.details = dependency_1.messages.CategoryDetails.fromObject(data.details);
            }
            return message;
        }
        toObject() {
            const data: {
                details?: ReturnType<typeof dependency_1.messages.CategoryDetails.prototype.toObject>;
            } = {};
            if (this.details != null) {
                data.details = this.details.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_details)
                writer.writeMessage(1, this.details, () => this.details.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetCategoryResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetCategoryResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.details, () => message.details = dependency_1.messages.CategoryDetails.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetCategoryResponse {
            return GetCategoryResponse.deserialize(bytes);
        }
    }
}
