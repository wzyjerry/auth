import type { ImmerReducer } from 'umi';
import { history } from 'umi';
import type { EffectsCommandMap } from 'dva';
import type { LoginRequest, LoginReply } from '@/api/user/v1/login';
import { LoginClient } from '@/api/request';
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
  };
  reducers: {
    setToken: ImmerReducer<
      User,
      {
        type: 'setToken';
        payload: string;
      }
    >;
  };
  subscriptions: unknown;
}
const AUTH_TOKEN_KEY = 'AUTH_TOKEN';
const UserModel: UserModelType = {
  state: User.fromJSON({
    token: localStorage.getItem(AUTH_TOKEN_KEY),
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
      } else {
        history.push(
          `/user/login?error=${error.name}&return_to=${encodeURIComponent(payload.returnTo)}`,
        );
      }
    },
  },
  reducers: {
    setToken(state, { payload }) {
      state.token = payload;
      localStorage.setItem(AUTH_TOKEN_KEY, payload);
    },
  },
  subscriptions: {},
};
export default UserModel;
