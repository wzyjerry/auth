import { Empty } from 'google-protobuf/google/protobuf/empty_pb';
import { request } from 'umi';
import {
  LoginPrePhoneRequest,
  LoginReply,
  LoginRequest,
} from './user/v1/login_pb';

export class LoginClient {
  async Login(bm: LoginRequest): Promise<LoginReply.AsObject> {
    return await request('/user/v1/login/login', {
      method: 'POST',
      data: bm.toObject(),
    });
  }
  async PrePhone(bm: LoginPrePhoneRequest): Promise<Empty> {
    return await request('/user/v1/login/pre_phone', {
      method: 'POST',
      data: bm.toObject(),
    });
  }
}
