/* eslint-disable */
import * as Long from 'long';
import * as _m0 from 'protobufjs/minimal';
import { Empty } from '../../google/protobuf/empty';

export const protobufPackage = 'api.user.v1';

export enum Type {
  TYPE_UNSET = 0,
  TYPE_ACCOUNT = 1,
  TYPE_EMAIL = 2,
  TYPE_PHONE = 3,
  TYPE_GITHUB = 4,
  TYPE_MICROSOFT = 5,
  UNRECOGNIZED = -1,
}

export function typeFromJSON(object: any): Type {
  switch (object) {
    case 0:
    case 'TYPE_UNSET':
      return Type.TYPE_UNSET;
    case 1:
    case 'TYPE_ACCOUNT':
      return Type.TYPE_ACCOUNT;
    case 2:
    case 'TYPE_EMAIL':
      return Type.TYPE_EMAIL;
    case 3:
    case 'TYPE_PHONE':
      return Type.TYPE_PHONE;
    case 4:
    case 'TYPE_GITHUB':
      return Type.TYPE_GITHUB;
    case 5:
    case 'TYPE_MICROSOFT':
      return Type.TYPE_MICROSOFT;
    case -1:
    case 'UNRECOGNIZED':
    default:
      return Type.UNRECOGNIZED;
  }
}

export function typeToJSON(object: Type): string {
  switch (object) {
    case Type.TYPE_UNSET:
      return 'TYPE_UNSET';
    case Type.TYPE_ACCOUNT:
      return 'TYPE_ACCOUNT';
    case Type.TYPE_EMAIL:
      return 'TYPE_EMAIL';
    case Type.TYPE_PHONE:
      return 'TYPE_PHONE';
    case Type.TYPE_GITHUB:
      return 'TYPE_GITHUB';
    case Type.TYPE_MICROSOFT:
      return 'TYPE_MICROSOFT';
    case Type.UNRECOGNIZED:
    default:
      return 'UNRECOGNIZED';
  }
}

export enum Method {
  METHOD_UNSET = 0,
  METHOD_PASSWORD = 1,
  METHOD_CODE = 2,
  UNRECOGNIZED = -1,
}

export function methodFromJSON(object: any): Method {
  switch (object) {
    case 0:
    case 'METHOD_UNSET':
      return Method.METHOD_UNSET;
    case 1:
    case 'METHOD_PASSWORD':
      return Method.METHOD_PASSWORD;
    case 2:
    case 'METHOD_CODE':
      return Method.METHOD_CODE;
    case -1:
    case 'UNRECOGNIZED':
    default:
      return Method.UNRECOGNIZED;
  }
}

export function methodToJSON(object: Method): string {
  switch (object) {
    case Method.METHOD_UNSET:
      return 'METHOD_UNSET';
    case Method.METHOD_PASSWORD:
      return 'METHOD_PASSWORD';
    case Method.METHOD_CODE:
      return 'METHOD_CODE';
    case Method.UNRECOGNIZED:
    default:
      return 'UNRECOGNIZED';
  }
}

export interface LoginRequest {
  type: Type;
  method: Method;
  unique?: string | undefined;
  secret: string;
}

export interface OAuthLoginReply {
  tokenType: string;
  expiresIn: number;
  scope: string;
  accessToken: string;
  refreshToken?: string | undefined;
  idToken?: string | undefined;
}

export interface LoginReply {
  token: string;
}

export interface LoginPrePhoneRequest {
  phone: string;
}

export interface LoginPreEmailRequest {
  email: string;
}

export interface TrashReply {
  authToken: boolean;
  userToken: boolean;
  clientToken: boolean;
  sub: string;
}

function createBaseLoginRequest(): LoginRequest {
  return { type: 0, method: 0, unique: undefined, secret: '' };
}

export const LoginRequest = {
  encode(
    message: LoginRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.method !== 0) {
      writer.uint32(16).int32(message.method);
    }
    if (message.unique !== undefined) {
      writer.uint32(26).string(message.unique);
    }
    if (message.secret !== '') {
      writer.uint32(34).string(message.secret);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.method = reader.int32() as any;
          break;
        case 3:
          message.unique = reader.string();
          break;
        case 4:
          message.secret = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginRequest {
    return {
      type: isSet(object.type) ? typeFromJSON(object.type) : 0,
      method: isSet(object.method) ? methodFromJSON(object.method) : 0,
      unique: isSet(object.unique) ? String(object.unique) : undefined,
      secret: isSet(object.secret) ? String(object.secret) : '',
    };
  },

  toJSON(message: LoginRequest): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = typeToJSON(message.type));
    message.method !== undefined && (obj.method = methodToJSON(message.method));
    message.unique !== undefined && (obj.unique = message.unique);
    message.secret !== undefined && (obj.secret = message.secret);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginRequest>, I>>(
    object: I,
  ): LoginRequest {
    const message = createBaseLoginRequest();
    message.type = object.type ?? 0;
    message.method = object.method ?? 0;
    message.unique = object.unique ?? undefined;
    message.secret = object.secret ?? '';
    return message;
  },
};

function createBaseOAuthLoginReply(): OAuthLoginReply {
  return {
    tokenType: '',
    expiresIn: 0,
    scope: '',
    accessToken: '',
    refreshToken: undefined,
    idToken: undefined,
  };
}

export const OAuthLoginReply = {
  encode(
    message: OAuthLoginReply,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.tokenType !== '') {
      writer.uint32(10).string(message.tokenType);
    }
    if (message.expiresIn !== 0) {
      writer.uint32(16).int32(message.expiresIn);
    }
    if (message.scope !== '') {
      writer.uint32(26).string(message.scope);
    }
    if (message.accessToken !== '') {
      writer.uint32(34).string(message.accessToken);
    }
    if (message.refreshToken !== undefined) {
      writer.uint32(42).string(message.refreshToken);
    }
    if (message.idToken !== undefined) {
      writer.uint32(50).string(message.idToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OAuthLoginReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuthLoginReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenType = reader.string();
          break;
        case 2:
          message.expiresIn = reader.int32();
          break;
        case 3:
          message.scope = reader.string();
          break;
        case 4:
          message.accessToken = reader.string();
          break;
        case 5:
          message.refreshToken = reader.string();
          break;
        case 6:
          message.idToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OAuthLoginReply {
    return {
      tokenType: isSet(object.tokenType) ? String(object.tokenType) : '',
      expiresIn: isSet(object.expiresIn) ? Number(object.expiresIn) : 0,
      scope: isSet(object.scope) ? String(object.scope) : '',
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : '',
      refreshToken: isSet(object.refreshToken)
        ? String(object.refreshToken)
        : undefined,
      idToken: isSet(object.idToken) ? String(object.idToken) : undefined,
    };
  },

  toJSON(message: OAuthLoginReply): unknown {
    const obj: any = {};
    message.tokenType !== undefined && (obj.tokenType = message.tokenType);
    message.expiresIn !== undefined &&
      (obj.expiresIn = Math.round(message.expiresIn));
    message.scope !== undefined && (obj.scope = message.scope);
    message.accessToken !== undefined &&
      (obj.accessToken = message.accessToken);
    message.refreshToken !== undefined &&
      (obj.refreshToken = message.refreshToken);
    message.idToken !== undefined && (obj.idToken = message.idToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OAuthLoginReply>, I>>(
    object: I,
  ): OAuthLoginReply {
    const message = createBaseOAuthLoginReply();
    message.tokenType = object.tokenType ?? '';
    message.expiresIn = object.expiresIn ?? 0;
    message.scope = object.scope ?? '';
    message.accessToken = object.accessToken ?? '';
    message.refreshToken = object.refreshToken ?? undefined;
    message.idToken = object.idToken ?? undefined;
    return message;
  },
};

function createBaseLoginReply(): LoginReply {
  return { token: '' };
}

export const LoginReply = {
  encode(
    message: LoginReply,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.token !== '') {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginReply {
    return {
      token: isSet(object.token) ? String(object.token) : '',
    };
  },

  toJSON(message: LoginReply): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginReply>, I>>(
    object: I,
  ): LoginReply {
    const message = createBaseLoginReply();
    message.token = object.token ?? '';
    return message;
  },
};

function createBaseLoginPrePhoneRequest(): LoginPrePhoneRequest {
  return { phone: '' };
}

export const LoginPrePhoneRequest = {
  encode(
    message: LoginPrePhoneRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.phone !== '') {
      writer.uint32(10).string(message.phone);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): LoginPrePhoneRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginPrePhoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.phone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginPrePhoneRequest {
    return {
      phone: isSet(object.phone) ? String(object.phone) : '',
    };
  },

  toJSON(message: LoginPrePhoneRequest): unknown {
    const obj: any = {};
    message.phone !== undefined && (obj.phone = message.phone);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginPrePhoneRequest>, I>>(
    object: I,
  ): LoginPrePhoneRequest {
    const message = createBaseLoginPrePhoneRequest();
    message.phone = object.phone ?? '';
    return message;
  },
};

function createBaseLoginPreEmailRequest(): LoginPreEmailRequest {
  return { email: '' };
}

export const LoginPreEmailRequest = {
  encode(
    message: LoginPreEmailRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.email !== '') {
      writer.uint32(10).string(message.email);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): LoginPreEmailRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginPreEmailRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginPreEmailRequest {
    return {
      email: isSet(object.email) ? String(object.email) : '',
    };
  },

  toJSON(message: LoginPreEmailRequest): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginPreEmailRequest>, I>>(
    object: I,
  ): LoginPreEmailRequest {
    const message = createBaseLoginPreEmailRequest();
    message.email = object.email ?? '';
    return message;
  },
};

function createBaseTrashReply(): TrashReply {
  return { authToken: false, userToken: false, clientToken: false, sub: '' };
}

export const TrashReply = {
  encode(
    message: TrashReply,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.authToken === true) {
      writer.uint32(8).bool(message.authToken);
    }
    if (message.userToken === true) {
      writer.uint32(16).bool(message.userToken);
    }
    if (message.clientToken === true) {
      writer.uint32(24).bool(message.clientToken);
    }
    if (message.sub !== '') {
      writer.uint32(34).string(message.sub);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TrashReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTrashReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authToken = reader.bool();
          break;
        case 2:
          message.userToken = reader.bool();
          break;
        case 3:
          message.clientToken = reader.bool();
          break;
        case 4:
          message.sub = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TrashReply {
    return {
      authToken: isSet(object.authToken) ? Boolean(object.authToken) : false,
      userToken: isSet(object.userToken) ? Boolean(object.userToken) : false,
      clientToken: isSet(object.clientToken)
        ? Boolean(object.clientToken)
        : false,
      sub: isSet(object.sub) ? String(object.sub) : '',
    };
  },

  toJSON(message: TrashReply): unknown {
    const obj: any = {};
    message.authToken !== undefined && (obj.authToken = message.authToken);
    message.userToken !== undefined && (obj.userToken = message.userToken);
    message.clientToken !== undefined &&
      (obj.clientToken = message.clientToken);
    message.sub !== undefined && (obj.sub = message.sub);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TrashReply>, I>>(
    object: I,
  ): TrashReply {
    const message = createBaseTrashReply();
    message.authToken = object.authToken ?? false;
    message.userToken = object.userToken ?? false;
    message.clientToken = object.clientToken ?? false;
    message.sub = object.sub ?? '';
    return message;
  },
};

export interface LoginService {
  PrePhone(request: LoginPrePhoneRequest): Promise<Empty>;
  PreEmail(request: LoginPreEmailRequest): Promise<Empty>;
  Login(request: LoginRequest): Promise<LoginReply>;
  Trash(request: Empty): Promise<TrashReply>;
}

export class LoginServiceClientImpl implements LoginService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.PrePhone = this.PrePhone.bind(this);
    this.PreEmail = this.PreEmail.bind(this);
    this.Login = this.Login.bind(this);
    this.Trash = this.Trash.bind(this);
  }
  PrePhone(request: LoginPrePhoneRequest): Promise<Empty> {
    const data = LoginPrePhoneRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.LoginService',
      'PrePhone',
      data,
    );
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  PreEmail(request: LoginPreEmailRequest): Promise<Empty> {
    const data = LoginPreEmailRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.LoginService',
      'PreEmail',
      data,
    );
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  Login(request: LoginRequest): Promise<LoginReply> {
    const data = LoginRequest.encode(request).finish();
    const promise = this.rpc.request('api.user.v1.LoginService', 'Login', data);
    return promise.then((data) => LoginReply.decode(new _m0.Reader(data)));
  }

  Trash(request: Empty): Promise<TrashReply> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request('api.user.v1.LoginService', 'Trash', data);
    return promise.then((data) => TrashReply.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array,
  ): Promise<Uint8Array>;
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

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
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
