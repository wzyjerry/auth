/* eslint-disable */
import Long from 'long';
import * as _m0 from 'protobufjs/minimal';

export const protobufPackage = 'api.oauth2.v1';

export interface AuthorizeRequest {
  responseType: string;
  clientId: string;
  redirectUri: string;
  scope: string;
  nonce?: string | undefined;
}

export interface AuthorizeReply {
  code: string;
  idToken?: string | undefined;
}

export interface PreAuthorizeRequest {
  responseType: string;
  clientId: string;
  redirectUri: string;
  scope: string;
}

export interface PreAuthorizeReply {
  name: string;
  logo?: string | undefined;
  homepage: string;
  description?: string | undefined;
}

export interface OAuth {
  name: string;
  logo?: string | undefined;
  homepage: string;
  description?: string | undefined;
  code: string;
  idToken?: string | undefined;
}

export interface TokenRequest {
  grantType: string;
  clientId: string;
  clientSecret: string;
  code?: string | undefined;
  refreshToken?: string | undefined;
}

export interface TokenReply {
  tokenType: string;
  accessToken: string;
  expiresIn: number;
  scope: string;
  refreshToken?: string | undefined;
  idToken?: string | undefined;
}

function createBaseAuthorizeRequest(): AuthorizeRequest {
  return { responseType: '', clientId: '', redirectUri: '', scope: '', nonce: undefined };
}

export const AuthorizeRequest = {
  encode(message: AuthorizeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.responseType !== '') {
      writer.uint32(10).string(message.responseType);
    }
    if (message.clientId !== '') {
      writer.uint32(18).string(message.clientId);
    }
    if (message.redirectUri !== '') {
      writer.uint32(26).string(message.redirectUri);
    }
    if (message.scope !== '') {
      writer.uint32(34).string(message.scope);
    }
    if (message.nonce !== undefined) {
      writer.uint32(42).string(message.nonce);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthorizeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthorizeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.responseType = reader.string();
          break;
        case 2:
          message.clientId = reader.string();
          break;
        case 3:
          message.redirectUri = reader.string();
          break;
        case 4:
          message.scope = reader.string();
          break;
        case 5:
          message.nonce = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AuthorizeRequest {
    return {
      responseType: isSet(object.responseType) ? String(object.responseType) : '',
      clientId: isSet(object.clientId) ? String(object.clientId) : '',
      redirectUri: isSet(object.redirectUri) ? String(object.redirectUri) : '',
      scope: isSet(object.scope) ? String(object.scope) : '',
      nonce: isSet(object.nonce) ? String(object.nonce) : undefined,
    };
  },

  toJSON(message: AuthorizeRequest): unknown {
    const obj: any = {};
    message.responseType !== undefined && (obj.responseType = message.responseType);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.redirectUri !== undefined && (obj.redirectUri = message.redirectUri);
    message.scope !== undefined && (obj.scope = message.scope);
    message.nonce !== undefined && (obj.nonce = message.nonce);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AuthorizeRequest>, I>>(object: I): AuthorizeRequest {
    const message = createBaseAuthorizeRequest();
    message.responseType = object.responseType ?? '';
    message.clientId = object.clientId ?? '';
    message.redirectUri = object.redirectUri ?? '';
    message.scope = object.scope ?? '';
    message.nonce = object.nonce ?? undefined;
    return message;
  },
};

function createBaseAuthorizeReply(): AuthorizeReply {
  return { code: '', idToken: undefined };
}

export const AuthorizeReply = {
  encode(message: AuthorizeReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== '') {
      writer.uint32(10).string(message.code);
    }
    if (message.idToken !== undefined) {
      writer.uint32(18).string(message.idToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthorizeReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthorizeReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        case 2:
          message.idToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AuthorizeReply {
    return {
      code: isSet(object.code) ? String(object.code) : '',
      idToken: isSet(object.idToken) ? String(object.idToken) : undefined,
    };
  },

  toJSON(message: AuthorizeReply): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.idToken !== undefined && (obj.idToken = message.idToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AuthorizeReply>, I>>(object: I): AuthorizeReply {
    const message = createBaseAuthorizeReply();
    message.code = object.code ?? '';
    message.idToken = object.idToken ?? undefined;
    return message;
  },
};

function createBasePreAuthorizeRequest(): PreAuthorizeRequest {
  return { responseType: '', clientId: '', redirectUri: '', scope: '' };
}

export const PreAuthorizeRequest = {
  encode(message: PreAuthorizeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.responseType !== '') {
      writer.uint32(10).string(message.responseType);
    }
    if (message.clientId !== '') {
      writer.uint32(18).string(message.clientId);
    }
    if (message.redirectUri !== '') {
      writer.uint32(26).string(message.redirectUri);
    }
    if (message.scope !== '') {
      writer.uint32(34).string(message.scope);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PreAuthorizeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreAuthorizeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.responseType = reader.string();
          break;
        case 2:
          message.clientId = reader.string();
          break;
        case 3:
          message.redirectUri = reader.string();
          break;
        case 4:
          message.scope = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreAuthorizeRequest {
    return {
      responseType: isSet(object.responseType) ? String(object.responseType) : '',
      clientId: isSet(object.clientId) ? String(object.clientId) : '',
      redirectUri: isSet(object.redirectUri) ? String(object.redirectUri) : '',
      scope: isSet(object.scope) ? String(object.scope) : '',
    };
  },

  toJSON(message: PreAuthorizeRequest): unknown {
    const obj: any = {};
    message.responseType !== undefined && (obj.responseType = message.responseType);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.redirectUri !== undefined && (obj.redirectUri = message.redirectUri);
    message.scope !== undefined && (obj.scope = message.scope);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PreAuthorizeRequest>, I>>(
    object: I,
  ): PreAuthorizeRequest {
    const message = createBasePreAuthorizeRequest();
    message.responseType = object.responseType ?? '';
    message.clientId = object.clientId ?? '';
    message.redirectUri = object.redirectUri ?? '';
    message.scope = object.scope ?? '';
    return message;
  },
};

function createBasePreAuthorizeReply(): PreAuthorizeReply {
  return { name: '', logo: undefined, homepage: '', description: undefined };
}

export const PreAuthorizeReply = {
  encode(message: PreAuthorizeReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name);
    }
    if (message.logo !== undefined) {
      writer.uint32(18).string(message.logo);
    }
    if (message.homepage !== '') {
      writer.uint32(26).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(34).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PreAuthorizeReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreAuthorizeReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.logo = reader.string();
          break;
        case 3:
          message.homepage = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreAuthorizeReply {
    return {
      name: isSet(object.name) ? String(object.name) : '',
      logo: isSet(object.logo) ? String(object.logo) : undefined,
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
    };
  },

  toJSON(message: PreAuthorizeReply): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.logo !== undefined && (obj.logo = message.logo);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PreAuthorizeReply>, I>>(object: I): PreAuthorizeReply {
    const message = createBasePreAuthorizeReply();
    message.name = object.name ?? '';
    message.logo = object.logo ?? undefined;
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    return message;
  },
};

function createBaseOAuth(): OAuth {
  return {
    name: '',
    logo: undefined,
    homepage: '',
    description: undefined,
    code: '',
    idToken: undefined,
  };
}

export const OAuth = {
  encode(message: OAuth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name);
    }
    if (message.logo !== undefined) {
      writer.uint32(18).string(message.logo);
    }
    if (message.homepage !== '') {
      writer.uint32(26).string(message.homepage);
    }
    if (message.description !== undefined) {
      writer.uint32(34).string(message.description);
    }
    if (message.code !== '') {
      writer.uint32(42).string(message.code);
    }
    if (message.idToken !== undefined) {
      writer.uint32(50).string(message.idToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OAuth {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.logo = reader.string();
          break;
        case 3:
          message.homepage = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.code = reader.string();
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

  fromJSON(object: any): OAuth {
    return {
      name: isSet(object.name) ? String(object.name) : '',
      logo: isSet(object.logo) ? String(object.logo) : undefined,
      homepage: isSet(object.homepage) ? String(object.homepage) : '',
      description: isSet(object.description) ? String(object.description) : undefined,
      code: isSet(object.code) ? String(object.code) : '',
      idToken: isSet(object.idToken) ? String(object.idToken) : undefined,
    };
  },

  toJSON(message: OAuth): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.logo !== undefined && (obj.logo = message.logo);
    message.homepage !== undefined && (obj.homepage = message.homepage);
    message.description !== undefined && (obj.description = message.description);
    message.code !== undefined && (obj.code = message.code);
    message.idToken !== undefined && (obj.idToken = message.idToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OAuth>, I>>(object: I): OAuth {
    const message = createBaseOAuth();
    message.name = object.name ?? '';
    message.logo = object.logo ?? undefined;
    message.homepage = object.homepage ?? '';
    message.description = object.description ?? undefined;
    message.code = object.code ?? '';
    message.idToken = object.idToken ?? undefined;
    return message;
  },
};

function createBaseTokenRequest(): TokenRequest {
  return {
    grantType: '',
    clientId: '',
    clientSecret: '',
    code: undefined,
    refreshToken: undefined,
  };
}

export const TokenRequest = {
  encode(message: TokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.grantType !== '') {
      writer.uint32(10).string(message.grantType);
    }
    if (message.clientId !== '') {
      writer.uint32(18).string(message.clientId);
    }
    if (message.clientSecret !== '') {
      writer.uint32(26).string(message.clientSecret);
    }
    if (message.code !== undefined) {
      writer.uint32(34).string(message.code);
    }
    if (message.refreshToken !== undefined) {
      writer.uint32(42).string(message.refreshToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.grantType = reader.string();
          break;
        case 2:
          message.clientId = reader.string();
          break;
        case 3:
          message.clientSecret = reader.string();
          break;
        case 4:
          message.code = reader.string();
          break;
        case 5:
          message.refreshToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenRequest {
    return {
      grantType: isSet(object.grantType) ? String(object.grantType) : '',
      clientId: isSet(object.clientId) ? String(object.clientId) : '',
      clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : '',
      code: isSet(object.code) ? String(object.code) : undefined,
      refreshToken: isSet(object.refreshToken) ? String(object.refreshToken) : undefined,
    };
  },

  toJSON(message: TokenRequest): unknown {
    const obj: any = {};
    message.grantType !== undefined && (obj.grantType = message.grantType);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    message.code !== undefined && (obj.code = message.code);
    message.refreshToken !== undefined && (obj.refreshToken = message.refreshToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TokenRequest>, I>>(object: I): TokenRequest {
    const message = createBaseTokenRequest();
    message.grantType = object.grantType ?? '';
    message.clientId = object.clientId ?? '';
    message.clientSecret = object.clientSecret ?? '';
    message.code = object.code ?? undefined;
    message.refreshToken = object.refreshToken ?? undefined;
    return message;
  },
};

function createBaseTokenReply(): TokenReply {
  return {
    tokenType: '',
    accessToken: '',
    expiresIn: 0,
    scope: '',
    refreshToken: undefined,
    idToken: undefined,
  };
}

export const TokenReply = {
  encode(message: TokenReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tokenType !== '') {
      writer.uint32(10).string(message.tokenType);
    }
    if (message.accessToken !== '') {
      writer.uint32(18).string(message.accessToken);
    }
    if (message.expiresIn !== 0) {
      writer.uint32(24).int64(message.expiresIn);
    }
    if (message.scope !== '') {
      writer.uint32(34).string(message.scope);
    }
    if (message.refreshToken !== undefined) {
      writer.uint32(42).string(message.refreshToken);
    }
    if (message.idToken !== undefined) {
      writer.uint32(50).string(message.idToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenType = reader.string();
          break;
        case 2:
          message.accessToken = reader.string();
          break;
        case 3:
          message.expiresIn = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.scope = reader.string();
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

  fromJSON(object: any): TokenReply {
    return {
      tokenType: isSet(object.tokenType) ? String(object.tokenType) : '',
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : '',
      expiresIn: isSet(object.expiresIn) ? Number(object.expiresIn) : 0,
      scope: isSet(object.scope) ? String(object.scope) : '',
      refreshToken: isSet(object.refreshToken) ? String(object.refreshToken) : undefined,
      idToken: isSet(object.idToken) ? String(object.idToken) : undefined,
    };
  },

  toJSON(message: TokenReply): unknown {
    const obj: any = {};
    message.tokenType !== undefined && (obj.tokenType = message.tokenType);
    message.accessToken !== undefined && (obj.accessToken = message.accessToken);
    message.expiresIn !== undefined && (obj.expiresIn = Math.round(message.expiresIn));
    message.scope !== undefined && (obj.scope = message.scope);
    message.refreshToken !== undefined && (obj.refreshToken = message.refreshToken);
    message.idToken !== undefined && (obj.idToken = message.idToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TokenReply>, I>>(object: I): TokenReply {
    const message = createBaseTokenReply();
    message.tokenType = object.tokenType ?? '';
    message.accessToken = object.accessToken ?? '';
    message.expiresIn = object.expiresIn ?? 0;
    message.scope = object.scope ?? '';
    message.refreshToken = object.refreshToken ?? undefined;
    message.idToken = object.idToken ?? undefined;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis;
  if (typeof self !== 'undefined') return self;
  if (typeof window !== 'undefined') return window;
  if (typeof global !== 'undefined') return global;
  throw 'Unable to locate global object';
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
