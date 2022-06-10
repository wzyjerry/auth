import { Empty } from 'google-protobuf/google/protobuf/empty_pb';
import { request } from 'umi';
import {
  LoginPrePhoneRequest,
  LoginReply,
  LoginRequest,
} from './user/v1/login_pb';

import { GetAvatarReply } from './user/v1/profile_pb';
import { CreateRequest, CreateReply } from './application/v1/application_pb';
import { Token } from '@/util/localStorage';
import { Reject } from '@/util';

export class LoginClient {
  async Login(bm: LoginRequest): Promise<LoginReply.AsObject> {
    return request('/user/v1/login/login', {
      method: 'POST',
      data: bm.toObject(),
    });
  }
  async PrePhone(bm: LoginPrePhoneRequest): Promise<Empty> {
    return request('/user/v1/login/pre_phone', {
      method: 'POST',
      data: bm.toObject(),
    });
  }
}

const helper = new Token();
const getBearer = (): string => {
  const token = helper.Load()
  if (token instanceof Reject) {
    return '';
  }
  return `Bearer ${token.val}`;
}
export class ProfileClient {
  async GetAvatar(): Promise<GetAvatarReply.AsObject> {
    return request('/user/v1/profile/avatar', {
      method: 'GET',
      headers: {
        'Authorization': getBearer(),
      }
    })
  }
}

export class ApplicationClient {
  async Create(bm: CreateRequest): Promise<CreateReply.AsObject> {
    return request('/application/v1', {
      method: 'POST',
      headers: {
        'Authorization': getBearer(),
      },
      data: bm.toObject(),
    })
  }
}
