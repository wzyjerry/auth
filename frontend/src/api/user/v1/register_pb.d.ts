// package: api.user.v1
// file: user/v1/register.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class RegisterAccountRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getNickname(): string;
  setNickname(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterAccountRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterAccountRequest): RegisterAccountRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterAccountRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterAccountRequest;
  static deserializeBinaryFromReader(message: RegisterAccountRequest, reader: jspb.BinaryReader): RegisterAccountRequest;
}

export namespace RegisterAccountRequest {
  export type AsObject = {
    username: string,
    password: string,
    nickname: string,
  }
}

export class RegisterReply extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterReply.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterReply): RegisterReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterReply;
  static deserializeBinaryFromReader(message: RegisterReply, reader: jspb.BinaryReader): RegisterReply;
}

export namespace RegisterReply {
  export type AsObject = {
    id: string,
  }
}

export class RegisterPreEmailRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterPreEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterPreEmailRequest): RegisterPreEmailRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterPreEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterPreEmailRequest;
  static deserializeBinaryFromReader(message: RegisterPreEmailRequest, reader: jspb.BinaryReader): RegisterPreEmailRequest;
}

export namespace RegisterPreEmailRequest {
  export type AsObject = {
    email: string,
  }
}

export class RegisterEmailRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getNickname(): string;
  setNickname(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterEmailRequest): RegisterEmailRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterEmailRequest;
  static deserializeBinaryFromReader(message: RegisterEmailRequest, reader: jspb.BinaryReader): RegisterEmailRequest;
}

export namespace RegisterEmailRequest {
  export type AsObject = {
    email: string,
    password: string,
    nickname: string,
    code: string,
  }
}

export class RegisterPrePhoneRequest extends jspb.Message {
  getPhone(): string;
  setPhone(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterPrePhoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterPrePhoneRequest): RegisterPrePhoneRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterPrePhoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterPrePhoneRequest;
  static deserializeBinaryFromReader(message: RegisterPrePhoneRequest, reader: jspb.BinaryReader): RegisterPrePhoneRequest;
}

export namespace RegisterPrePhoneRequest {
  export type AsObject = {
    phone: string,
  }
}

export class RegisterPhoneRequest extends jspb.Message {
  getPhone(): string;
  setPhone(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getNickname(): string;
  setNickname(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterPhoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterPhoneRequest): RegisterPhoneRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterPhoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterPhoneRequest;
  static deserializeBinaryFromReader(message: RegisterPhoneRequest, reader: jspb.BinaryReader): RegisterPhoneRequest;
}

export namespace RegisterPhoneRequest {
  export type AsObject = {
    phone: string,
    password: string,
    nickname: string,
    code: string,
  }
}

