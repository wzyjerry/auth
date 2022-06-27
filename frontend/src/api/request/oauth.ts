import { go, getBearer } from '.';
import { request } from 'umi';
import type {
  PreAuthorizeRequest,
  PreAuthorizeReply,
  AuthorizeRequest,
  AuthorizeReply,
} from '../oauth2/v1/oauth';

export async function PreAuthorize(bm: PreAuthorizeRequest) {
  return go<PreAuthorizeReply>(
    request('/oauth2/v1/pre_authorize', {
      method: 'POST',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}

export async function Authorize(bm: AuthorizeRequest) {
  return go<AuthorizeReply>(
    request('/oauth2/v1/authorize', {
      method: 'POST',
      headers: {
        Authorization: getBearer(),
      },
      data: bm,
    }),
  );
}
