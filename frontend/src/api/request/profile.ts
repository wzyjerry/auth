import { go, getBearer } from '.';
import { request } from 'umi';
import type { GetAvatarReply } from '../user/v1/profile';
export async function GetAvatar() {
  return go<GetAvatarReply>(
    request('/user/v1/profile/avatar', {
      method: 'GET',
      headers: {
        Authorization: getBearer(),
      },
    }),
  );
}
