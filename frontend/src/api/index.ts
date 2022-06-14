import { Empty } from './google/protobuf/empty';
import { request } from 'umi';
import {
  LoginPrePhoneRequest,
  LoginReply,
  LoginRequest,
} from './user/v1/login';

import { GetAvatarReply } from './user/v1/profile';
import {
  CreateRequest,
  CreateReply,
  RetrieveRequest,
  RetrieveReply,
  GenerateClientSecretRequest,
  GenerateClientSecretReply,
  RevokeClientSecretRequest,
} from './application/v1/application';
import { Token } from '@/util/localStorage';
import { Reject } from '@/util';

export class LoginClient {
  async Login(bm: LoginRequest): Promise<LoginReply> {
    return request('/user/v1/login/login', {
      method: 'POST',
      data: bm,
    });
  }
  async PrePhone(bm: LoginPrePhoneRequest): Promise<Empty> {
    return request('/user/v1/login/pre_phone', {
      method: 'POST',
      data: bm,
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
  async GetAvatar(): Promise<GetAvatarReply> {
    return request('/user/v1/profile/avatar', {
      method: 'GET',
      headers: {
        'Authorization': getBearer(),
      }
    })
  }
}

export class ApplicationClient {
  async Create(bm: CreateRequest): Promise<CreateReply> {
    return request('/application/v1', {
      method: 'POST',
      headers: {
        'Authorization': getBearer(),
      },
      data: bm,
    })
  }
  async Retrieve(bm: RetrieveRequest): Promise<RetrieveReply> {
    return request(`/application/v1/${bm.id}`, {
      method: 'GET',
      headers: {
        'Authorization': getBearer(),
      },
    })
  }
  async GenerateClientSecret(bm: GenerateClientSecretRequest): Promise<GenerateClientSecretReply> {
    return request(`/application/v1/${bm.id}/generateClientSecret`, {
      method: 'POST',
      headers: {
        'Authorization': getBearer(),
      },
      data: bm,
    })
  }
  async RevokeClientSecret(bm: RevokeClientSecretRequest): Promise<Empty> {
    return request(`/application/v1/${bm.id}/${bm.secretId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': getBearer(),
      },
    })
  }
}
