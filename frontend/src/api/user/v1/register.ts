/* eslint-disable */
import * as Long from 'long';
import * as _m0 from 'protobufjs/minimal';
import { Empty } from '../../google/protobuf/empty';

export const protobufPackage = 'api.user.v1';

export interface RegisterAccountRequest {
  /** 账户名(*) */
  username: string;
  /** 密码 */
  password: string;
  /** 昵称 */
  nickname: string;
}

export interface RegisterReply {
  /** 用户ID */
  id: string;
}

export interface RegisterPreEmailRequest {
  /** 邮箱(*) */
  email: string;
}

export interface RegisterEmailRequest {
  /** 邮箱(*) */
  email: string;
  /** 密码 */
  password: string;
  /** 昵称 */
  nickname: string;
  /** 验证码 */
  code: string;
}

export interface RegisterPrePhoneRequest {
  /** 手机号(*) */
  phone: string;
}

export interface RegisterPhoneRequest {
  /** 手机号(*) */
  phone: string;
  /** 密码 */
  password: string;
  /** 昵称 */
  nickname: string;
  /** 验证码 */
  code: string;
}

function createBaseRegisterAccountRequest(): RegisterAccountRequest {
  return { username: '', password: '', nickname: '' };
}

export const RegisterAccountRequest = {
  encode(
    message: RegisterAccountRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.username !== '') {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== '') {
      writer.uint32(18).string(message.password);
    }
    if (message.nickname !== '') {
      writer.uint32(26).string(message.nickname);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): RegisterAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.nickname = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterAccountRequest {
    return {
      username: isSet(object.username) ? String(object.username) : '',
      password: isSet(object.password) ? String(object.password) : '',
      nickname: isSet(object.nickname) ? String(object.nickname) : '',
    };
  },

  toJSON(message: RegisterAccountRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.password !== undefined && (obj.password = message.password);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterAccountRequest>, I>>(
    object: I,
  ): RegisterAccountRequest {
    const message = createBaseRegisterAccountRequest();
    message.username = object.username ?? '';
    message.password = object.password ?? '';
    message.nickname = object.nickname ?? '';
    return message;
  },
};

function createBaseRegisterReply(): RegisterReply {
  return { id: '' };
}

export const RegisterReply = {
  encode(
    message: RegisterReply,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.id !== '') {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterReply();
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

  fromJSON(object: any): RegisterReply {
    return {
      id: isSet(object.id) ? String(object.id) : '',
    };
  },

  toJSON(message: RegisterReply): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterReply>, I>>(
    object: I,
  ): RegisterReply {
    const message = createBaseRegisterReply();
    message.id = object.id ?? '';
    return message;
  },
};

function createBaseRegisterPreEmailRequest(): RegisterPreEmailRequest {
  return { email: '' };
}

export const RegisterPreEmailRequest = {
  encode(
    message: RegisterPreEmailRequest,
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
  ): RegisterPreEmailRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterPreEmailRequest();
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

  fromJSON(object: any): RegisterPreEmailRequest {
    return {
      email: isSet(object.email) ? String(object.email) : '',
    };
  },

  toJSON(message: RegisterPreEmailRequest): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterPreEmailRequest>, I>>(
    object: I,
  ): RegisterPreEmailRequest {
    const message = createBaseRegisterPreEmailRequest();
    message.email = object.email ?? '';
    return message;
  },
};

function createBaseRegisterEmailRequest(): RegisterEmailRequest {
  return { email: '', password: '', nickname: '', code: '' };
}

export const RegisterEmailRequest = {
  encode(
    message: RegisterEmailRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.email !== '') {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== '') {
      writer.uint32(18).string(message.password);
    }
    if (message.nickname !== '') {
      writer.uint32(26).string(message.nickname);
    }
    if (message.code !== '') {
      writer.uint32(34).string(message.code);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): RegisterEmailRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterEmailRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.nickname = reader.string();
          break;
        case 4:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterEmailRequest {
    return {
      email: isSet(object.email) ? String(object.email) : '',
      password: isSet(object.password) ? String(object.password) : '',
      nickname: isSet(object.nickname) ? String(object.nickname) : '',
      code: isSet(object.code) ? String(object.code) : '',
    };
  },

  toJSON(message: RegisterEmailRequest): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterEmailRequest>, I>>(
    object: I,
  ): RegisterEmailRequest {
    const message = createBaseRegisterEmailRequest();
    message.email = object.email ?? '';
    message.password = object.password ?? '';
    message.nickname = object.nickname ?? '';
    message.code = object.code ?? '';
    return message;
  },
};

function createBaseRegisterPrePhoneRequest(): RegisterPrePhoneRequest {
  return { phone: '' };
}

export const RegisterPrePhoneRequest = {
  encode(
    message: RegisterPrePhoneRequest,
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
  ): RegisterPrePhoneRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterPrePhoneRequest();
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

  fromJSON(object: any): RegisterPrePhoneRequest {
    return {
      phone: isSet(object.phone) ? String(object.phone) : '',
    };
  },

  toJSON(message: RegisterPrePhoneRequest): unknown {
    const obj: any = {};
    message.phone !== undefined && (obj.phone = message.phone);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterPrePhoneRequest>, I>>(
    object: I,
  ): RegisterPrePhoneRequest {
    const message = createBaseRegisterPrePhoneRequest();
    message.phone = object.phone ?? '';
    return message;
  },
};

function createBaseRegisterPhoneRequest(): RegisterPhoneRequest {
  return { phone: '', password: '', nickname: '', code: '' };
}

export const RegisterPhoneRequest = {
  encode(
    message: RegisterPhoneRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.phone !== '') {
      writer.uint32(10).string(message.phone);
    }
    if (message.password !== '') {
      writer.uint32(18).string(message.password);
    }
    if (message.nickname !== '') {
      writer.uint32(26).string(message.nickname);
    }
    if (message.code !== '') {
      writer.uint32(34).string(message.code);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): RegisterPhoneRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterPhoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.phone = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.nickname = reader.string();
          break;
        case 4:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterPhoneRequest {
    return {
      phone: isSet(object.phone) ? String(object.phone) : '',
      password: isSet(object.password) ? String(object.password) : '',
      nickname: isSet(object.nickname) ? String(object.nickname) : '',
      code: isSet(object.code) ? String(object.code) : '',
    };
  },

  toJSON(message: RegisterPhoneRequest): unknown {
    const obj: any = {};
    message.phone !== undefined && (obj.phone = message.phone);
    message.password !== undefined && (obj.password = message.password);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterPhoneRequest>, I>>(
    object: I,
  ): RegisterPhoneRequest {
    const message = createBaseRegisterPhoneRequest();
    message.phone = object.phone ?? '';
    message.password = object.password ?? '';
    message.nickname = object.nickname ?? '';
    message.code = object.code ?? '';
    return message;
  },
};

export interface RegisterService {
  Account(request: RegisterAccountRequest): Promise<RegisterReply>;
  PreEmail(request: RegisterPreEmailRequest): Promise<Empty>;
  Email(request: RegisterEmailRequest): Promise<RegisterReply>;
  PrePhone(request: RegisterPrePhoneRequest): Promise<Empty>;
  Phone(request: RegisterPhoneRequest): Promise<RegisterReply>;
}

export class RegisterServiceClientImpl implements RegisterService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Account = this.Account.bind(this);
    this.PreEmail = this.PreEmail.bind(this);
    this.Email = this.Email.bind(this);
    this.PrePhone = this.PrePhone.bind(this);
    this.Phone = this.Phone.bind(this);
  }
  Account(request: RegisterAccountRequest): Promise<RegisterReply> {
    const data = RegisterAccountRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.RegisterService',
      'Account',
      data,
    );
    return promise.then((data) => RegisterReply.decode(new _m0.Reader(data)));
  }

  PreEmail(request: RegisterPreEmailRequest): Promise<Empty> {
    const data = RegisterPreEmailRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.RegisterService',
      'PreEmail',
      data,
    );
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  Email(request: RegisterEmailRequest): Promise<RegisterReply> {
    const data = RegisterEmailRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.RegisterService',
      'Email',
      data,
    );
    return promise.then((data) => RegisterReply.decode(new _m0.Reader(data)));
  }

  PrePhone(request: RegisterPrePhoneRequest): Promise<Empty> {
    const data = RegisterPrePhoneRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.RegisterService',
      'PrePhone',
      data,
    );
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  Phone(request: RegisterPhoneRequest): Promise<RegisterReply> {
    const data = RegisterPhoneRequest.encode(request).finish();
    const promise = this.rpc.request(
      'api.user.v1.RegisterService',
      'Phone',
      data,
    );
    return promise.then((data) => RegisterReply.decode(new _m0.Reader(data)));
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
