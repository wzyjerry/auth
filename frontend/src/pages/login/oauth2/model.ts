import type { ImmerReducer } from 'umi';
import type { EffectsCommandMap } from 'dva';
import type {
  PreAuthorizeRequest,
  PreAuthorizeReply,
  AuthorizeRequest,
  AuthorizeReply,
} from '@/api/oauth2/v1/oauth';
import { OAuthClient } from '@/api/request';
import { OAuth } from '@/api/oauth2/v1/oauth';
export { OAuth };
export interface LoginOAuth2AuthorizeModelType {
  namespace: 'login_oauth2_authorize';
  state: OAuth;
  effects: {
    setup: (
      action: {
        payload: {
          request: PreAuthorizeRequest;
        };
      },
      effects: EffectsCommandMap,
    ) => void;
    authorize: (
      action: {
        payload: {
          request: AuthorizeRequest;
        };
      },
      effects: EffectsCommandMap,
    ) => void;
  };
  reducers: {
    set: ImmerReducer<
      OAuth,
      {
        type: 'set';
        payload: PreAuthorizeReply | AuthorizeReply;
      }
    >;
  };
  subscriptions: unknown;
}
const LoginOAuth2AuthorizeModel: LoginOAuth2AuthorizeModelType = {
  namespace: 'login_oauth2_authorize',
  state: OAuth.fromJSON({}),
  effects: {
    *setup({ payload }, { call, put }) {
      const { response }: { response: PreAuthorizeReply } = yield call(
        OAuthClient.PreAuthorize,
        payload,
      );
      if (response) {
        yield put({
          type: 'set',
          payload: response,
        });
      }
    },
    *authorize({ payload }, { call, put }) {
      const { response }: { response: AuthorizeReply } = yield call(OAuthClient.Authorize, payload);
      if (response) {
        yield put({
          type: 'set',
          payload: response,
        });
      }
    },
  },
  reducers: {
    set(state, { payload }) {
      Object.assign(state, payload);
    },
  },
  subscriptions: {},
};
export default LoginOAuth2AuthorizeModel;
