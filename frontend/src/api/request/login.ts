import type { Empty } from '../google/protobuf/empty';
import { go } from '.';
import { request } from 'umi';
import type { LoginPrePhoneRequest, LoginReply, LoginRequest } from '../user/v1/login';

export async function Login(bm: LoginRequest) {
  return go<LoginReply>(
    request('/user/v1/login/login', {
      method: 'POST',
      data: bm,
    }),
  );
}

export async function PrePhone(bm: LoginPrePhoneRequest) {
  return go<Empty>(
    request('/user/v1/login/pre_phone', {
      method: 'POST',
      data: bm,
    }),
  );
}
