// package: api.user.v1
// file: user/v1/profile.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../protoc-gen-openapiv2/options/annotations_pb";

export class GetAvatarReply extends jspb.Message {
  getAvatar(): string;
  setAvatar(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAvatarReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetAvatarReply): GetAvatarReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetAvatarReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAvatarReply;
  static deserializeBinaryFromReader(message: GetAvatarReply, reader: jspb.BinaryReader): GetAvatarReply;
}

export namespace GetAvatarReply {
  export type AsObject = {
    avatar: string,
  }
}

