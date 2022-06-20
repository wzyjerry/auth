import { Empty } from '../google/protobuf/empty';
import { request } from 'umi';
import {
  CreateRequest,
  CreateReply,
  RetrieveRequest,
  RetrieveReply,
  GenerateClientSecretRequest,
  GenerateClientSecretReply,
  RevokeClientSecretRequest,
  UploadLogoRequest,
  UpdateRequest,
  DeleteRequest,
  GetAllReply,
} from '../application/v1/application';
import { go, getBearer } from '.';

export async function Create(bm: CreateRequest) {
  return go<CreateReply>(
    request('/application/v1', {
      method: 'POST',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}
export async function Retrieve(bm: RetrieveRequest) {
  return go<RetrieveReply>(
    request(`/application/v1/${bm.id}`, {
      method: 'GET',
      headers: {
        Authorization: getBearer(),
      },
    }),
  );
}

export async function GenerateClientSecret(bm: GenerateClientSecretRequest) {
  return go<GenerateClientSecretReply>(
    request(`/application/v1/${bm.id}/generateClientSecret`, {
      method: 'POST',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}

export async function RevokeClientSecret(bm: RevokeClientSecretRequest) {
  return go<Empty>(
    request(`/application/v1/${bm.id}/${bm.secretId}`, {
      method: 'DELETE',
      headers: {
        Authorization: getBearer(),
      },
    }),
  );
}

export async function UploadLogo(bm: UploadLogoRequest) {
  return go<Empty>(
    request(`/application/v1/${bm.id}/uploadLogo`, {
      method: 'PUT',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}

export async function Update(bm: UpdateRequest) {
  return go<Empty>(
    request(`/application/v1/${bm.id}`, {
      method: 'PUT',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}

export async function Delete(bm: DeleteRequest) {
  return go<Empty>(
    request(`/application/v1/${bm.id}`, {
      method: 'DELETE',
      headers: {
        Authorization: getBearer(),
      },
    }),
  );
}

export async function GetAll() {
  return go<GetAllReply>(
    request(`/application/v1`, {
      method: 'GET',
      headers: {
        Authorization: getBearer(),
      },
    }),
  );
}
