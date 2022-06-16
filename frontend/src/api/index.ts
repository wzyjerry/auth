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

const go = async<T>(promise: Promise<T>) => {
  return promise.then(response => ({response})).catch(error => ({error}));
}
const AUTH_TOKEN_KEY = "AUTH_TOKEN";
const getBearer = () => {
  return `Bearer ${localStorage.getItem(AUTH_TOKEN_KEY)}`;
}

export namespace LoginClient {
  export async function Login(bm: LoginRequest) {
    return go<LoginReply>(request('/user/v1/login/login', {
      method: 'POST',
      data: bm,
    }));
  }
  export async function PrePhone(bm: LoginPrePhoneRequest) {
    return go<Empty>(request('/user/v1/login/pre_phone', {
      method: 'POST',
      data: bm,
    }));
  }
}
export namespace ProfileClient {
  export async function GetAvatar() {
    return go<GetAvatarReply>(request('/user/v1/profile/avatar', {
      method: 'GET',
      headers: {
        'Authorization': getBearer(),
      }
    }));
  }
}

export namespace ApplicationClient {
  export async function Create(bm: CreateRequest) {
    return go<CreateReply>(request('/application/v1', {
      method: 'POST',
      headers: {
        'Authorization': getBearer(),
      },
      data: bm,
    }));
  }
  export async function Retrieve(bm: RetrieveRequest) {
    return go<RetrieveReply>(request(`/application/v1/${bm.id}`, {
      method: 'GET',
      headers: {
        'Authorization': getBearer(),
      },
    }));
  }
  export async function GenerateClientSecret(bm: GenerateClientSecretRequest) {
    return go<GenerateClientSecretReply>(request(`/application/v1/${bm.id}/generateClientSecret`, {
      method: 'POST',
      headers: {
        'Authorization': getBearer(),
      },
      data: bm,
    }));
  }
  export async function RevokeClientSecret (bm: RevokeClientSecretRequest) {
    return go<Empty>(request(`/application/v1/${bm.id}/${bm.secretId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': getBearer(),
      },
    }));
  }
}
