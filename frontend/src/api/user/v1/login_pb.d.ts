// package: api.user.v1
// file: user/v1/login.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class LoginRequest extends jspb.Message {
  getType(): TypeMap[keyof TypeMap];
  setType(value: TypeMap[keyof TypeMap]): void;

  getMethod(): MethodMap[keyof MethodMap];
  setMethod(value: MethodMap[keyof MethodMap]): void;

  getUnique(): string;
  setUnique(value: string): void;

  getSecret(): string;
  setSecret(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    type: TypeMap[keyof TypeMap],
    method: MethodMap[keyof MethodMap],
    unique: string,
    secret: string,
  }
}

export class OAuthLoginReply extends jspb.Message {
  getTokenType(): string;
  setTokenType(value: string): void;

  getExpiresIn(): number;
  setExpiresIn(value: number): void;

  getScope(): string;
  setScope(value: string): void;

  getAccessToken(): string;
  setAccessToken(value: string): void;

  hasRefreshToken(): boolean;
  clearRefreshToken(): void;
  getRefreshToken(): string;
  setRefreshToken(value: string): void;

  hasIdToken(): boolean;
  clearIdToken(): void;
  getIdToken(): string;
  setIdToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OAuthLoginReply.AsObject;
  static toObject(includeInstance: boolean, msg: OAuthLoginReply): OAuthLoginReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: OAuthLoginReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OAuthLoginReply;
  static deserializeBinaryFromReader(message: OAuthLoginReply, reader: jspb.BinaryReader): OAuthLoginReply;
}

export namespace OAuthLoginReply {
  export type AsObject = {
    tokenType: string,
    expiresIn: number,
    scope: string,
    accessToken: string,
    refreshToken: string,
    idToken: string,
  }
}

export class LoginReply extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginReply.AsObject;
  static toObject(includeInstance: boolean, msg: LoginReply): LoginReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginReply;
  static deserializeBinaryFromReader(message: LoginReply, reader: jspb.BinaryReader): LoginReply;
}

export namespace LoginReply {
  export type AsObject = {
    token: string,
  }
}

export class LoginPrePhoneRequest extends jspb.Message {
  getPhone(): string;
  setPhone(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginPrePhoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginPrePhoneRequest): LoginPrePhoneRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginPrePhoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginPrePhoneRequest;
  static deserializeBinaryFromReader(message: LoginPrePhoneRequest, reader: jspb.BinaryReader): LoginPrePhoneRequest;
}

export namespace LoginPrePhoneRequest {
  export type AsObject = {
    phone: string,
  }
}

export class LoginPreEmailRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginPreEmailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginPreEmailRequest): LoginPreEmailRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginPreEmailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginPreEmailRequest;
  static deserializeBinaryFromReader(message: LoginPreEmailRequest, reader: jspb.BinaryReader): LoginPreEmailRequest;
}

export namespace LoginPreEmailRequest {
  export type AsObject = {
    email: string,
  }
}

export class TrashReply extends jspb.Message {
  getAuthtoken(): boolean;
  setAuthtoken(value: boolean): void;

  getUsertoken(): boolean;
  setUsertoken(value: boolean): void;

  getClienttoken(): boolean;
  setClienttoken(value: boolean): void;

  getSub(): string;
  setSub(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TrashReply.AsObject;
  static toObject(includeInstance: boolean, msg: TrashReply): TrashReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TrashReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TrashReply;
  static deserializeBinaryFromReader(message: TrashReply, reader: jspb.BinaryReader): TrashReply;
}

export namespace TrashReply {
  export type AsObject = {
    authtoken: boolean,
    usertoken: boolean,
    clienttoken: boolean,
    sub: string,
  }
}

export interface TypeMap {
  TYPE_UNSET: 0;
  TYPE_ACCOUNT: 1;
  TYPE_EMAIL: 2;
  TYPE_PHONE: 3;
  TYPE_GITHUB: 4;
  TYPE_MICROSOFT: 5;
}

export const Type: TypeMap;

export interface MethodMap {
  METHOD_UNSET: 0;
  METHOD_PASSWORD: 1;
  METHOD_CODE: 2;
}

export const Method: MethodMap;

