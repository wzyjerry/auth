/* eslint-disable */
import * as Long from 'long';
import * as _m0 from 'protobufjs/minimal';
import { Empty } from '../../google/protobuf/empty';

export const protobufPackage = 'api.user.v1';

export interface GetAvatarReply {
  avatar: string;
}

export interface User {
  token: string;
  id: string;
  nickname: string;
  avatar: string;
}

function createBaseGetAvatarReply(): GetAvatarReply {
  return { avatar: '' };
}

export const GetAvatarReply = {
  encode(message: GetAvatarReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.avatar !== '') {
      writer.uint32(10).string(message.avatar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAvatarReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAvatarReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.avatar = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAvatarReply {
    return {
      avatar: isSet(object.avatar) ? String(object.avatar) : '',
    };
  },

  toJSON(message: GetAvatarReply): unknown {
    const obj: any = {};
    message.avatar !== undefined && (obj.avatar = message.avatar);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAvatarReply>, I>>(object: I): GetAvatarReply {
    const message = createBaseGetAvatarReply();
    message.avatar = object.avatar ?? '';
    return message;
  },
};

function createBaseUser(): User {
  return { token: '', id: '', nickname: '', avatar: '' };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== '') {
      writer.uint32(10).string(message.token);
    }
    if (message.id !== '') {
      writer.uint32(18).string(message.id);
    }
    if (message.nickname !== '') {
      writer.uint32(26).string(message.nickname);
    }
    if (message.avatar !== '') {
      writer.uint32(34).string(message.avatar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.nickname = reader.string();
          break;
        case 4:
          message.avatar = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      token: isSet(object.token) ? String(object.token) : '',
      id: isSet(object.id) ? String(object.id) : '',
      nickname: isSet(object.nickname) ? String(object.nickname) : '',
      avatar: isSet(object.avatar) ? String(object.avatar) : '',
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.id !== undefined && (obj.id = message.id);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    message.avatar !== undefined && (obj.avatar = message.avatar);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.token = object.token ?? '';
    message.id = object.id ?? '';
    message.nickname = object.nickname ?? '';
    message.avatar = object.avatar ?? '';
    return message;
  },
};

export interface ProfileService {
  GetAvatar(request: Empty): Promise<GetAvatarReply>;
}

export class ProfileServiceClientImpl implements ProfileService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetAvatar = this.GetAvatar.bind(this);
  }
  GetAvatar(request: Empty): Promise<GetAvatarReply> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request('api.user.v1.ProfileService', 'GetAvatar', data);
    return promise.then((data) => GetAvatarReply.decode(new _m0.Reader(data)));
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
