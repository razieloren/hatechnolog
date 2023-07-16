/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 4.23.4
 * source: stats/api.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./helpers";
import * as pb_1 from "google-protobuf";
export namespace messages {
    export class LatestStatsPushRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            discord_guild?: string;
            youtube_channel?: string;
            github_repo?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("discord_guild" in data && data.discord_guild != undefined) {
                    this.discord_guild = data.discord_guild;
                }
                if ("youtube_channel" in data && data.youtube_channel != undefined) {
                    this.youtube_channel = data.youtube_channel;
                }
                if ("github_repo" in data && data.github_repo != undefined) {
                    this.github_repo = data.github_repo;
                }
            }
        }
        get discord_guild() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set discord_guild(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get youtube_channel() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set youtube_channel(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get github_repo() {
            return pb_1.Message.getFieldWithDefault(this, 3, "") as string;
        }
        set github_repo(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            discord_guild?: string;
            youtube_channel?: string;
            github_repo?: string;
        }): LatestStatsPushRequest {
            const message = new LatestStatsPushRequest({});
            if (data.discord_guild != null) {
                message.discord_guild = data.discord_guild;
            }
            if (data.youtube_channel != null) {
                message.youtube_channel = data.youtube_channel;
            }
            if (data.github_repo != null) {
                message.github_repo = data.github_repo;
            }
            return message;
        }
        toObject() {
            const data: {
                discord_guild?: string;
                youtube_channel?: string;
                github_repo?: string;
            } = {};
            if (this.discord_guild != null) {
                data.discord_guild = this.discord_guild;
            }
            if (this.youtube_channel != null) {
                data.youtube_channel = this.youtube_channel;
            }
            if (this.github_repo != null) {
                data.github_repo = this.github_repo;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.discord_guild.length)
                writer.writeString(1, this.discord_guild);
            if (this.youtube_channel.length)
                writer.writeString(2, this.youtube_channel);
            if (this.github_repo.length)
                writer.writeString(3, this.github_repo);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): LatestStatsPushRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new LatestStatsPushRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.discord_guild = reader.readString();
                        break;
                    case 2:
                        message.youtube_channel = reader.readString();
                        break;
                    case 3:
                        message.github_repo = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): LatestStatsPushRequest {
            return LatestStatsPushRequest.deserialize(bytes);
        }
    }
    export class LatestStatsPushResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            discord_stats?: dependency_1.messages.LatestDiscordStats;
            youtube_stats?: dependency_1.messages.LatestYoutubeStats;
            github_stats?: dependency_1.messages.LatestGithubStats;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("discord_stats" in data && data.discord_stats != undefined) {
                    this.discord_stats = data.discord_stats;
                }
                if ("youtube_stats" in data && data.youtube_stats != undefined) {
                    this.youtube_stats = data.youtube_stats;
                }
                if ("github_stats" in data && data.github_stats != undefined) {
                    this.github_stats = data.github_stats;
                }
            }
        }
        get discord_stats() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.LatestDiscordStats, 1) as dependency_1.messages.LatestDiscordStats;
        }
        set discord_stats(value: dependency_1.messages.LatestDiscordStats) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_discord_stats() {
            return pb_1.Message.getField(this, 1) != null;
        }
        get youtube_stats() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.LatestYoutubeStats, 2) as dependency_1.messages.LatestYoutubeStats;
        }
        set youtube_stats(value: dependency_1.messages.LatestYoutubeStats) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get has_youtube_stats() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get github_stats() {
            return pb_1.Message.getWrapperField(this, dependency_1.messages.LatestGithubStats, 3) as dependency_1.messages.LatestGithubStats;
        }
        set github_stats(value: dependency_1.messages.LatestGithubStats) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get has_github_stats() {
            return pb_1.Message.getField(this, 3) != null;
        }
        static fromObject(data: {
            discord_stats?: ReturnType<typeof dependency_1.messages.LatestDiscordStats.prototype.toObject>;
            youtube_stats?: ReturnType<typeof dependency_1.messages.LatestYoutubeStats.prototype.toObject>;
            github_stats?: ReturnType<typeof dependency_1.messages.LatestGithubStats.prototype.toObject>;
        }): LatestStatsPushResponse {
            const message = new LatestStatsPushResponse({});
            if (data.discord_stats != null) {
                message.discord_stats = dependency_1.messages.LatestDiscordStats.fromObject(data.discord_stats);
            }
            if (data.youtube_stats != null) {
                message.youtube_stats = dependency_1.messages.LatestYoutubeStats.fromObject(data.youtube_stats);
            }
            if (data.github_stats != null) {
                message.github_stats = dependency_1.messages.LatestGithubStats.fromObject(data.github_stats);
            }
            return message;
        }
        toObject() {
            const data: {
                discord_stats?: ReturnType<typeof dependency_1.messages.LatestDiscordStats.prototype.toObject>;
                youtube_stats?: ReturnType<typeof dependency_1.messages.LatestYoutubeStats.prototype.toObject>;
                github_stats?: ReturnType<typeof dependency_1.messages.LatestGithubStats.prototype.toObject>;
            } = {};
            if (this.discord_stats != null) {
                data.discord_stats = this.discord_stats.toObject();
            }
            if (this.youtube_stats != null) {
                data.youtube_stats = this.youtube_stats.toObject();
            }
            if (this.github_stats != null) {
                data.github_stats = this.github_stats.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_discord_stats)
                writer.writeMessage(1, this.discord_stats, () => this.discord_stats.serialize(writer));
            if (this.has_youtube_stats)
                writer.writeMessage(2, this.youtube_stats, () => this.youtube_stats.serialize(writer));
            if (this.has_github_stats)
                writer.writeMessage(3, this.github_stats, () => this.github_stats.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): LatestStatsPushResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new LatestStatsPushResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.discord_stats, () => message.discord_stats = dependency_1.messages.LatestDiscordStats.deserialize(reader));
                        break;
                    case 2:
                        reader.readMessage(message.youtube_stats, () => message.youtube_stats = dependency_1.messages.LatestYoutubeStats.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.github_stats, () => message.github_stats = dependency_1.messages.LatestGithubStats.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): LatestStatsPushResponse {
            return LatestStatsPushResponse.deserialize(bytes);
        }
    }
}