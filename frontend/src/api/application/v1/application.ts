/* eslint-disable */
import * as Long from 'long';
import * as _m0 from 'protobufjs/minimal';
import { Timestamp } from '../../google/protobuf/timestamp';
import { Empty } from '../../google/protobuf/empty';

export const protobufPackage = 'api.application.v1';

export interface CreateRequest {
  name: string;
  homepage: string;
  description?: string | undefined;
  callback: string;
}

export interface CreateReply {
  id: string;
}

export interface RetrieveRequest {
  id: string;
}

export interface Secret {
  id: string;
  lastUsed?: Date | undefined;
  description: string;
  masked: boolean;
  secret: string;
}

export interface Application {
  id: string;
  name: string;
  clientId: string;
  clientSecrets: Secret[];
  logo?: string | undefined;
  homepage: string;
  description?: string | undefined;
  callback: string;
}

export interface RetrieveReply {
  name: string;
  clientId: string;
  clientSecrets: Secret[];
  logo?: string | undefined;
  homepage: string;
  description?: string | undefined;
  callback: string;
}

export interface GenerateClientSecretRequest {
  id: string;
  description: string;
}

export interface GenerateClientSecretReply {
  secret: Secret | undefined;
}

export interface RevokeClientSecretRequest {
  id: string;
  secretId: string;
}

export interface UploadLogoRequest {
  id: string;
  logo: string;
}

export interface UpdateRequest {
  id: string;
  name: string;
  homepage: string;
  description?: string | undefined;
  callback: string;
}

export interface DeleteRequest {
  id: string;
}

export interface ApplicationOverview {
  id: string;
  name: string;
  maskedDescription?: string | undefined;
  logo?: string | undefined;
}

export interface GetAllReply {
  applicationOverviews: ApplicationOverview[];
}

function createBaseCreateRequest(): CreateRequest {
  return { name: '', homepage: '', description: undefined, callback: '' };
}

export const CreateRequest = {
  encode(message: CreateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name);
    }
    if (message.homepage !== '') {
      writer.uint32(18).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(26).string(message.description);
    }
    if (message.callback !== '') {
      writer.uint32(34).string(message.callback);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.homepage = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.callback = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateRequest {
    return {
      name: isSet(object.name) ? String(object.name) : '',
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
      callback: isSet(object.callback) ? String(object.callback) : '',
    };
  },

  toJSON(message: CreateRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    message.callback !== undefined && (obj.callback = message.callback);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateRequest>, I>>(object: I): CreateRequest {
    const message = createBaseCreateRequest();
    message.name = object.name ?? '';
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    message.callback = object.callback ?? '';
    return message;
  },
};

function createBaseCreateReply(): CreateReply {
  return { id: '' };
}

export const CreateReply = {
  encode(message: CreateReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateReply {
    return {
      id: isSet(object.id) ? String(object.id) : '',
    };
  },

  toJSON(message: CreateReply): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateReply>, I>>(object: I): CreateReply {
    const message = createBaseCreateReply();
    message.id = object.id ?? '';
    return message;
  },
};

function createBaseRetrieveRequest(): RetrieveRequest {
  return { id: '' };
}

export const RetrieveRequest = {
  encode(message: RetrieveRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RetrieveRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRetrieveRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RetrieveRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
    };
  },

  toJSON(message: RetrieveRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RetrieveRequest>, I>>(object: I): RetrieveRequest {
    const message = createBaseRetrieveRequest();
    message.id = object.id ?? '';
    return message;
  },
};

function createBaseSecret(): Secret {
  return { id: '', lastUsed: undefined, description: '', masked: false, secret: '' };
}

export const Secret = {
  encode(message: Secret, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.lastUsed !== undefined) {
      Timestamp.encode(toTimestamp(message.lastUsed), writer.uint32(18).fork()).ldelim();
    }
    if (message.description !== '') {
      writer.uint32(26).string(message.description);
    }
    if (message.masked === true) {
      writer.uint32(32).bool(message.masked);
    }
    if (message.secret !== '') {
      writer.uint32(42).string(message.secret);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Secret {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSecret();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.lastUsed = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.masked = reader.bool();
          break;
        case 5:
          message.secret = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Secret {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      lastUsed: isSet(object.lastUsed) ? fromJsonTimestamp(object.lastUsed) : undefined,
      description: isSet(object.description) ? String(object.description) : '',
      masked: isSet(object.masked) ? Boolean(object.masked) : false,
      secret: isSet(object.secret) ? String(object.secret) : '',
    };
  },

  toJSON(message: Secret): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.lastUsed !== undefined && (obj.lastUsed = message.lastUsed.toISOString());
    message.description !== undefined && (obj.description = message.description);
    message.masked !== undefined && (obj.masked = message.masked);
    message.secret !== undefined && (obj.secret = message.secret);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Secret>, I>>(object: I): Secret {
    const message = createBaseSecret();
    message.id = object.id ?? '';
    message.lastUsed = object.lastUsed ?? undefined;
    message.description = object.description ?? '';
    message.masked = object.masked ?? false;
    message.secret = object.secret ?? '';
    return message;
  },
};

function createBaseApplication(): Application {
  return {
    id: '',
    name: '',
    clientId: '',
    clientSecrets: [],
    logo: undefined,
    homepage: '',
    description: undefined,
    callback: '',
  };
}

export const Application = {
  encode(message: Application, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name);
    }
    if (message.clientId !== '') {
      writer.uint32(26).string(message.clientId);
    }
    for (const v of message.clientSecrets) {
      Secret.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.logo !== undefined) {
      writer.uint32(42).string(message.logo);
    }
    if (message.homepage !== '') {
      writer.uint32(50).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(58).string(message.description);
    }
    if (message.callback !== '') {
      writer.uint32(66).string(message.callback);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Application {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApplication();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.clientId = reader.string();
          break;
        case 4:
          message.clientSecrets.push(Secret.decode(reader, reader.uint32()));
          break;
        case 5:
          message.logo = reader.string();
          break;
        case 6:
          message.homepage = reader.string();
          break;
        case 7:
          message.description = reader.string();
          break;
        case 8:
          message.callback = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Application {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      name: isSet(object.name) ? String(object.name) : '',
      clientId: isSet(object.clientId) ? String(object.clientId) : '',
      clientSecrets: Array.isArray(object?.clientSecrets)
        ? object.clientSecrets.map((e: any) => Secret.fromJSON(e))
        : [],
      logo: isSet(object.logo) ? String(object.logo) : undefined,
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
      callback: isSet(object.callback) ? String(object.callback) : '',
    };
  },

  toJSON(message: Application): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    if (message.clientSecrets) {
      obj.clientSecrets = message.clientSecrets.map((e) => (e ? Secret.toJSON(e) : undefined));
    } else {
      obj.clientSecrets = [];
    }
    message.logo !== undefined && (obj.logo = message.logo);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    message.callback !== undefined && (obj.callback = message.callback);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Application>, I>>(object: I): Application {
    const message = createBaseApplication();
    message.id = object.id ?? '';
    message.name = object.name ?? '';
    message.clientId = object.clientId ?? '';
    message.clientSecrets = object.clientSecrets?.map((e) => Secret.fromPartial(e)) || [];
    message.logo = object.logo ?? undefined;
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    message.callback = object.callback ?? '';
    return message;
  },
};

function createBaseRetrieveReply(): RetrieveReply {
  return {
    name: '',
    clientId: '',
    clientSecrets: [],
    logo: undefined,
    homepage: '',
    description: undefined,
    callback: '',
  };
}

export const RetrieveReply = {
  encode(message: RetrieveReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name);
    }
    if (message.clientId !== '') {
      writer.uint32(18).string(message.clientId);
    }
    for (const v of message.clientSecrets) {
      Secret.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.logo !== undefined) {
      writer.uint32(34).string(message.logo);
    }
    if (message.homepage !== '') {
      writer.uint32(42).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(50).string(message.description);
    }
    if (message.callback !== '') {
      writer.uint32(58).string(message.callback);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RetrieveReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRetrieveReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.clientId = reader.string();
          break;
        case 3:
          message.clientSecrets.push(Secret.decode(reader, reader.uint32()));
          break;
        case 4:
          message.logo = reader.string();
          break;
        case 5:
          message.homepage = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.callback = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RetrieveReply {
    return {
      name: isSet(object.name) ? String(object.name) : '',
      clientId: isSet(object.clientId) ? String(object.clientId) : '',
      clientSecrets: Array.isArray(object?.clientSecrets)
        ? object.clientSecrets.map((e: any) => Secret.fromJSON(e))
        : [],
      logo: isSet(object.logo) ? String(object.logo) : undefined,
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
      callback: isSet(object.callback) ? String(object.callback) : '',
    };
  },

  toJSON(message: RetrieveReply): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    if (message.clientSecrets) {
      obj.clientSecrets = message.clientSecrets.map((e) => (e ? Secret.toJSON(e) : undefined));
    } else {
      obj.clientSecrets = [];
    }
    message.logo !== undefined && (obj.logo = message.logo);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    message.callback !== undefined && (obj.callback = message.callback);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RetrieveReply>, I>>(object: I): RetrieveReply {
    const message = createBaseRetrieveReply();
    message.name = object.name ?? '';
    message.clientId = object.clientId ?? '';
    message.clientSecrets = object.clientSecrets?.map((e) => Secret.fromPartial(e)) || [];
    message.logo = object.logo ?? undefined;
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    message.callback = object.callback ?? '';
    return message;
  },
};

function createBaseGenerateClientSecretRequest(): GenerateClientSecretRequest {
  return { id: '', description: '' };
}

export const GenerateClientSecretRequest = {
  encode(
    message: GenerateClientSecretRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.description !== '') {
      writer.uint32(18).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenerateClientSecretRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenerateClientSecretRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenerateClientSecretRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      description: isSet(object.description) ? String(object.description) : '',
    };
  },

  toJSON(message: GenerateClientSecretRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.description !== undefined && (obj.description = message.description);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenerateClientSecretRequest>, I>>(
    object: I,
  ): GenerateClientSecretRequest {
    const message = createBaseGenerateClientSecretRequest();
    message.id = object.id ?? '';
    message.description = object.description ?? '';
    return message;
  },
};

function createBaseGenerateClientSecretReply(): GenerateClientSecretReply {
  return { secret: undefined };
}

export const GenerateClientSecretReply = {
  encode(message: GenerateClientSecretReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.secret !== undefined) {
      Secret.encode(message.secret, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenerateClientSecretReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenerateClientSecretReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.secret = Secret.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenerateClientSecretReply {
    return {
      secret: isSet(object.secret) ? Secret.fromJSON(object.secret) : undefined,
    };
  },

  toJSON(message: GenerateClientSecretReply): unknown {
    const obj: any = {};
    message.secret !== undefined &&
      (obj.secret = message.secret ? Secret.toJSON(message.secret) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenerateClientSecretReply>, I>>(
    object: I,
  ): GenerateClientSecretReply {
    const message = createBaseGenerateClientSecretReply();
    message.secret =
      object.secret !== undefined && object.secret !== null
        ? Secret.fromPartial(object.secret)
        : undefined;
    return message;
  },
};

function createBaseRevokeClientSecretRequest(): RevokeClientSecretRequest {
  return { id: '', secretId: '' };
}

export const RevokeClientSecretRequest = {
  encode(message: RevokeClientSecretRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.secretId !== '') {
      writer.uint32(18).string(message.secretId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RevokeClientSecretRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRevokeClientSecretRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.secretId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RevokeClientSecretRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      secretId: isSet(object.secretId) ? String(object.secretId) : '',
    };
  },

  toJSON(message: RevokeClientSecretRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.secretId !== undefined && (obj.secretId = message.secretId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RevokeClientSecretRequest>, I>>(
    object: I,
  ): RevokeClientSecretRequest {
    const message = createBaseRevokeClientSecretRequest();
    message.id = object.id ?? '';
    message.secretId = object.secretId ?? '';
    return message;
  },
};

function createBaseUploadLogoRequest(): UploadLogoRequest {
  return { id: '', logo: '' };
}

export const UploadLogoRequest = {
  encode(message: UploadLogoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.logo !== '') {
      writer.uint32(18).string(message.logo);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UploadLogoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUploadLogoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.logo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UploadLogoRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      logo: isSet(object.logo) ? String(object.logo) : '',
    };
  },

  toJSON(message: UploadLogoRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.logo !== undefined && (obj.logo = message.logo);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UploadLogoRequest>, I>>(object: I): UploadLogoRequest {
    const message = createBaseUploadLogoRequest();
    message.id = object.id ?? '';
    message.logo = object.logo ?? '';
    return message;
  },
};

function createBaseUpdateRequest(): UpdateRequest {
  return { id: '', name: '', homepage: '', description: undefined, callback: '' };
}

export const UpdateRequest = {
  encode(message: UpdateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name);
    }
    if (message.homepage !== '') {
      writer.uint32(26).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(34).string(message.description);
    }
    if (message.callback !== '') {
      writer.uint32(42).string(message.callback);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.homepage = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.callback = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      name: isSet(object.name) ? String(object.name) : '',
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
      callback: isSet(object.callback) ? String(object.callback) : '',
    };
  },

  toJSON(message: UpdateRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    message.callback !== undefined && (obj.callback = message.callback);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateRequest>, I>>(object: I): UpdateRequest {
    const message = createBaseUpdateRequest();
    message.id = object.id ?? '';
    message.name = object.name ?? '';
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    message.callback = object.callback ?? '';
    return message;
  },
};

function createBaseDeleteRequest(): DeleteRequest {
  return { id: '' };
}

export const DeleteRequest = {
  encode(message: DeleteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteRequest {
    return {
      id: isSet(object.id) ? String(object.id) : '',
    };
  },

  toJSON(message: DeleteRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteRequest>, I>>(object: I): DeleteRequest {
    const message = createBaseDeleteRequest();
    message.id = object.id ?? '';
    return message;
  },
};

function createBaseApplicationOverview(): ApplicationOverview {
  return { id: '', name: '', maskedDescription: undefined, logo: undefined };
}

export const ApplicationOverview = {
  encode(message: ApplicationOverview, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name);
    }
    if (message.maskedDescription !== undefined) {
      writer.uint32(26).string(message.maskedDescription);
    }
    if (message.logo !== undefined) {
      writer.uint32(34).string(message.logo);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ApplicationOverview {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApplicationOverview();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.maskedDescription = reader.string();
          break;
        case 4:
          message.logo = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ApplicationOverview {
    return {
      id: isSet(object.id) ? String(object.id) : '',
      name: isSet(object.name) ? String(object.name) : '',
      maskedDescription: isSet(object.maskedDescription)
        ? String(object.maskedDescription)
        : undefined,
      logo: isSet(object.logo) ? String(object.logo) : undefined,
    };
  },

  toJSON(message: ApplicationOverview): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.maskedDescription !== undefined && (obj.maskedDescription = message.maskedDescription);
    message.logo !== undefined && (obj.logo = message.logo);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ApplicationOverview>, I>>(
    object: I,
  ): ApplicationOverview {
    const message = createBaseApplicationOverview();
    message.id = object.id ?? '';
    message.name = object.name ?? '';
    message.maskedDescription = object.maskedDescription ?? undefined;
    message.logo = object.logo ?? undefined;
    return message;
  },
};

function createBaseGetAllReply(): GetAllReply {
  return { applicationOverviews: [] };
}

export const GetAllReply = {
  encode(message: GetAllReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.applicationOverviews) {
      ApplicationOverview.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAllReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAllReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.applicationOverviews.push(ApplicationOverview.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAllReply {
    return {
      applicationOverviews: Array.isArray(object?.applicationOverviews)
        ? object.applicationOverviews.map((e: any) => ApplicationOverview.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetAllReply): unknown {
    const obj: any = {};
    if (message.applicationOverviews) {
      obj.applicationOverviews = message.applicationOverviews.map((e) =>
        e ? ApplicationOverview.toJSON(e) : undefined,
      );
    } else {
      obj.applicationOverviews = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAllReply>, I>>(object: I): GetAllReply {
    const message = createBaseGetAllReply();
    message.applicationOverviews =
      object.applicationOverviews?.map((e) => ApplicationOverview.fromPartial(e)) || [];
    return message;
  },
};

export interface ApplicationService {
  Create(request: CreateRequest): Promise<CreateReply>;
  Retrieve(request: RetrieveRequest): Promise<RetrieveReply>;
  GenerateClientSecret(request: GenerateClientSecretRequest): Promise<GenerateClientSecretReply>;
  RevokeClientSecret(request: RevokeClientSecretRequest): Promise<Empty>;
  UploadLogo(request: UploadLogoRequest): Promise<Empty>;
  Update(request: UpdateRequest): Promise<Empty>;
  Delete(request: DeleteRequest): Promise<Empty>;
  GetAll(request: Empty): Promise<GetAllReply>;
}

export class ApplicationServiceClientImpl implements ApplicationService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Create = this.Create.bind(this);
    this.Retrieve = this.Retrieve.bind(this);
    this.GenerateClientSecret = this.GenerateClientSecret.bind(this);
    this.RevokeClientSecret = this.RevokeClientSecret.bind(this);
    this.UploadLogo = this.UploadLogo.bind(this);
    this.Update = this.Update.bind(this);
    this.Delete = this.Delete.bind(this);
    this.GetAll = this.GetAll.bind(this);
  }
  Create(request: CreateRequest): Promise<CreateReply> {
    const data = CreateRequest.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'Create', data);
    return promise.then((data) => CreateReply.decode(new _m0.Reader(data)));
  }

  Retrieve(request: RetrieveRequest): Promise<RetrieveReply> {
    const data = RetrieveRequest.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'Retrieve', data);
    return promise.then((data) => RetrieveReply.decode(new _m0.Reader(data)));
  }

  GenerateClientSecret(request: GenerateClientSecretRequest): Promise<GenerateClientSecretReply> {
    const data = GenerateClientSecretRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.application.v1.ApplicationService',
      'GenerateClientSecret',
      data,
    );
    return promise.then((data) => GenerateClientSecretReply.decode(new _m0.Reader(data)));
  }

  RevokeClientSecret(request: RevokeClientSecretRequest): Promise<Empty> {
    const data = RevokeClientSecretRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.application.v1.ApplicationService',
      'RevokeClientSecret',
      data,
    );
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  UploadLogo(request: UploadLogoRequest): Promise<Empty> {
    const data = UploadLogoRequest.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'UploadLogo', data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  Update(request: UpdateRequest): Promise<Empty> {
    const data = UpdateRequest.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'Update', data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  Delete(request: DeleteRequest): Promise<Empty> {
    const data = DeleteRequest.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'Delete', data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  GetAll(request: Empty): Promise<GetAllReply> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request('api.application.v1.ApplicationService', 'GetAll', data);
    return promise.then((data) => GetAllReply.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<Exclude<keyof I, KeysOfUnion<P>>, never>;

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === 'string') {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
