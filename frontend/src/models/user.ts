import type { ImmerReducer } from 'umi';
import { history } from 'umi';
import jwt_decode from 'jwt-decode';
import type { Effect, EffectsCommandMap } from 'dva';
import type { LoginRequest, LoginReply } from '@/api/user/v1/login';
import type { GetAvatarReply } from '@/api/user/v1/profile';
import { LoginClient, ProfileClient } from '@/api/request';
import { User } from '@/api/user/v1/profile';
export { User };
export interface UserModelType {
  state: User;
  effects: {
    login: (
      action: {
        payload: {
          request: LoginRequest;
          returnTo: string;
        };
      },
      effects: EffectsCommandMap,
    ) => void;
    getAvatar: Effect;
  };
  reducers: {
    setToken: ImmerReducer<
      User,
      {
        type: 'setToken';
        payload: string;
      }
    >;
    setAvatar: ImmerReducer<
      User,
      {
        type: 'setToken';
        payload?: string;
      }
    >;
  };
  subscriptions: unknown;
}
const AUTH_TOKEN_KEY = 'AUTH_TOKEN';
const AVATAR_KEY = 'AVATAR';
const UserModel: UserModelType = {
  state: User.fromJSON({
    token: localStorage.getItem(AUTH_TOKEN_KEY),
    avatar: localStorage.getItem(AVATAR_KEY),
  }),
  effects: {
    *login({ payload }, { call, put }) {
      const { response, error }: { response: LoginReply; error: Error } = yield call(
        LoginClient.Login,
        payload.request,
      );
      if (response) {
        yield put({
          type: 'setToken',
          payload: response.token,
        });
        history.push(payload.returnTo);
        yield put({
          type: 'getAvatar',
        });
      } else {
        history.push(
          `/login?error=${error.name}&return_to=${encodeURIComponent(payload.returnTo)}`,
        );
      }
    },
    *getAvatar(_, { call, put }) {
      const { response, error }: { response: GetAvatarReply; error: Error } = yield call(
        ProfileClient.GetAvatar,
      );
      if (response) {
        yield put({
          type: 'setAvatar',
          payload: response.avatar,
        });
      } else if (error) {
        throw error;
      }
    },
  },
  reducers: {
    setToken(state, { payload }) {
      state.token = payload;
      localStorage.setItem(AUTH_TOKEN_KEY, payload);
      const decoded = jwt_decode(payload);
      console.log(decoded);
      // TODO: 添加idToken
    },
    setAvatar(state, { payload }) {
      if (payload) {
        state.avatar = payload;
        localStorage.setItem(AVATAR_KEY, payload);
      }
    },
  },
  subscriptions: {},
};
export default UserModel;
