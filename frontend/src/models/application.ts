import { ImmerReducer } from 'umi';
import { EffectsCommandMap } from 'dva';
import {
  Application,
  RetrieveRequest,
  RetrieveReply,
  GenerateClientSecretRequest,
  GenerateClientSecretReply,
  RevokeClientSecretRequest,
  Secret,
} from '@/api/application/v1/application';
import { ApplicationClient } from '@/api';
import { Empty } from '@/api/google/protobuf/empty';
export { Application };
export interface ApplicationModelType {
  state: Application;
  effects: {
    setup(action: {payload: RetrieveRequest}, effects: EffectsCommandMap): void,
    generateClientSecret(action: {payload: Omit<GenerateClientSecretRequest, 'id'>}, effects: EffectsCommandMap): void
    revokeClientSecret(action: {payload: Omit<RevokeClientSecretRequest, 'id'>}, effects: EffectsCommandMap): void
  };
  reducers: {
    set: ImmerReducer<Application, {
      type: 'set',
      payload: Application
    }>,
    pushFrontClientSecret: ImmerReducer<Application, {
      type: 'pushFrontClientSecret',
      payload: Secret
    }>,
    removeClientSecret: ImmerReducer<Application, {
      type: 'removeClientSecret',
      payload: string
    }>,
  };
  subscriptions: {};
}

// 如果首个Secret not masked, mask it.
const adjustClientSecrets = (secrets: Secret[]): Secret[] => {
  if (secrets.length > 0) {
    if (!secrets[0].masked) {
      secrets[0].masked = true;
      secrets[0].secret = `*****${secrets[0].secret.slice(secrets[0].secret.length-5)}`
    }
  }
  return secrets;
}

const ApplicationModel: ApplicationModelType = {
  state: Application.fromJSON({}),

  effects: {
    *setup({ payload }, { call, put }) {
      const {response, error}: { response: RetrieveReply, error: Error } = yield call(ApplicationClient.Retrieve, payload);
      const application: Application = {
        id: payload.id,
        ...response
      };
      if (response) {
        yield put({
          type: 'set',
          payload: application
        })
      }
    },
    *generateClientSecret({ payload }, { call, put, select }) {
      const { id } = yield select(state => state.application);
      const request: GenerateClientSecretRequest = { id, ...payload };
      const {response, error}: { response: GenerateClientSecretReply, error: Error } = yield call(ApplicationClient.GenerateClientSecret, request);
      if (response) {
        yield put({
          type: 'pushFrontClientSecret',
          payload: response.secret,
        })
      }
    },
    *revokeClientSecret({ payload }, { call, put, select }) {
      const { id } = yield select(state => state.application);
      const request: RevokeClientSecretRequest = { id, ...payload };
      const {response, error}: { response: Empty, error: Error } = yield call(ApplicationClient.RevokeClientSecret, request);
      if (response) {
        yield put({
          type: 'removeClientSecret',
          payload: payload.secretId,
        })
      }
    },
  },
  reducers: {
    set(state, { payload }) {
      Object.assign(state, payload);
    },
    pushFrontClientSecret(state, { payload }) {
      state.clientSecrets = adjustClientSecrets(state.clientSecrets);
      state.clientSecrets = [payload, ...state.clientSecrets];
    },
    removeClientSecret(state, { payload }) {
      state.clientSecrets = adjustClientSecrets(state.clientSecrets).
        filter(secret => secret.id != payload);
    }
  },
  subscriptions: {},
};

export default ApplicationModel;
